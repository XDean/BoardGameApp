package space

import (
	"github.com/xdean/miniboardgame/go/server/log"
	"sync"
)

type (
	Host interface {
		EventHostId() string
	}

	Message struct {
		From       int `json:"-"` // user id
		To         int `json:"-"` // user id, -1 means all
		Topic      string
		Attributes map[string]interface{}
		Payload    interface{}
	}

	Action struct {
		Func func()
		Done chan<- bool
	}

	Publisher struct {
		thread *Thread
	}

	Subscriber struct {
		MessageListener chan<- Message
		Done            <-chan bool
	}

	Subscription struct {
		MessageListener <-chan Message
		Done            chan<- bool
	}

	Thread struct {
		EventStream chan interface{} // Message or Action or Subscriber
		Done        chan bool
		Attribute   sync.Map
	}

	threadGetter struct {
		Host   Host
		Getter chan<- *Thread
	}
)

var (
	threadAccessStream = make(chan threadGetter)
	threads            = make(map[string]*Thread)
)

func init() {
	go threadAccess(threadAccessStream)
}

func threadAccess(stream chan threadGetter) {
	for {
		select {
		case g := <-stream:
			rt, ok := threads[g.Host.EventHostId()]
			if !ok {
				rt = &Thread{
					EventStream: make(chan interface{}, 10),
					Done:        make(chan bool),
					Attribute:   sync.Map{},
				}
				threads[g.Host.EventHostId()] = rt
				go rt.run()
			}
			g.Getter <- rt
		}
	}
}

func getThread(r Host) *Thread {
	ch := make(chan *Thread)
	g := threadGetter{
		Host:   r,
		Getter: ch,
	}
	threadAccessStream <- g
	rt := <-ch
	return rt
}

func (r *Thread) run() {
	subscribers := make([]Subscriber, 0)
	for {
		select {
		case event := <-r.EventStream:
			switch t := event.(type) {
			case Message:
				new := make([]Subscriber, 0)
				for _, v := range subscribers {
					select {
					case <-v.Done:
						close(v.MessageListener)
					default:
						select {
						case v.MessageListener <- t:
							new = append(new, v)
						default:
							// downstream is block
							close(v.MessageListener)
						}
					}
				}
				subscribers = new
			case Action:
				t.Func()
				t.Done <- true
				close(t.Done)
			case Subscriber:
				subscribers = append(subscribers, t)
			default:
				log.Global.Warn("Unknown event:", t)
			}
		case <-r.Done:
			for _, v := range subscribers {
				close(v.MessageListener)
			}
			return
		}
	}
}
