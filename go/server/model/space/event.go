package space

import "sync"

type (
	Host interface {
		EventHostId() string
	}

	Event struct {
		From       int `json:"-"` // user id
		To         int `json:"-"` // user id, -1 means all
		Topic      string
		Attributes map[string]interface{}
		Payload    interface{}
	}

	Observer struct {
		EventListener chan<- Event
		Done          <-chan bool
	}

	Subscription struct {
		EventListener <-chan Event
		Done          chan<- bool
	}

	Thread struct {
		EventStream    chan Event
		ObserverStream chan Observer
		Done           chan bool
		Attribute      sync.Map
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
					EventStream:    make(chan Event, 5),
					ObserverStream: make(chan Observer, 5),
					Done:           make(chan bool),
					Attribute:      sync.Map{},
				}
				threads[g.Host.EventHostId()] = rt
				go rt.run()
			}
			g.Getter <- rt
		}
	}
}

func SendEvent(r Host, event Event) {
	rt := getThread(r)
	rt.EventStream <- event
}

func Listen(r Host) Subscription {
	rt := getThread(r)
	eventListener := make(chan Event, 5)
	done := make(chan bool, 1)
	rt.ObserverStream <- Observer{
		EventListener: eventListener,
		Done:          done,
	}
	return Subscription{
		EventListener: eventListener,
		Done:          done,
	}
}

func Attribute(r Host, event Event) *sync.Map {
	rt := getThread(r)
	return &rt.Attribute
}

func Done(r Host, event Event) {
	rt := getThread(r)
	rt.Done <- true
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
	observers := make([]Observer, 0)
	for {
		select {
		case event := <-r.EventStream:
			new := make([]Observer, 0)
			for _, v := range observers {
				select {
				case v.EventListener <- event:
					new = append(new, v)
				case <-v.Done:
					close(v.EventListener)
				default:
					close(v.EventListener)
				}
			}
			observers = new
		case obs := <-r.ObserverStream:
			observers = append(observers, obs)
		case <-r.Done:
			for _, v := range observers {
				close(v.EventListener)
			}
			break
		}
	}
}
