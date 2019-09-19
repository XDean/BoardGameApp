package space

import "sync"

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
		MessageStream    chan Message
		ActionStream     chan Action
		SubscriberStream chan Subscriber
		Done             chan bool
		Attribute        sync.Map
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
					ActionStream:     make(chan Action, 5),
					MessageStream:    make(chan Message, 5),
					SubscriberStream: make(chan Subscriber, 5),
					Done:             make(chan bool),
					Attribute:        sync.Map{},
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

func (p Publisher) SendEvent(message Message) {
	p.thread.MessageStream <- message
}

func (r *Thread) run() {
	subscribers := make([]Subscriber, 0)
	for {
		select {
		case message := <-r.MessageStream:
			new := make([]Subscriber, 0)
			for _, v := range subscribers {
				select {
				case v.MessageListener <- message:
					new = append(new, v)
				case <-v.Done:
					close(v.MessageListener)
				default:
					close(v.MessageListener)
				}
			}
			subscribers = new
		case action := <-r.ActionStream:
			action.Func()
			action.Done <- true
			close(action.Done)
		case obs := <-r.SubscriberStream:
			subscribers = append(subscribers, obs)
		case <-r.Done:
			for _, v := range subscribers {
				close(v.MessageListener)
			}
			break
		}
	}
}
