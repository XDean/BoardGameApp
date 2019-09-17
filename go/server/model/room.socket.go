package model

type (
	Event struct {
		From       int `json:"-"` // user id
		To         int `json:"-"` // user id, -1 means all
		Topic      string
		Attributes map[string]interface{}
		Payload    interface{}
	}

	RoomObserver struct {
		EventListener chan<- Event
		Done          <-chan bool
	}

	RoomSubscription struct {
		EventListener <-chan Event
		Done          chan<- bool
	}

	RoomThread struct {
		EventStream    chan Event
		ObserverStream chan RoomObserver
		Done           chan bool
	}

	roomThreadGetter struct {
		Room   *Room
		Getter chan<- RoomThread
	}
)

var (
	roomThreadAccessStream = make(chan roomThreadGetter)
	roomThreads            = make(map[uint]RoomThread)
)

func init() {
	go roomThreadAccess(roomThreadAccessStream)
}

func roomThreadAccess(stream chan roomThreadGetter) {
	for {
		select {
		case g := <-stream:
			rt, ok := roomThreads[g.Room.ID]
			if !ok {
				rt = RoomThread{
					EventStream:    make(chan Event, 5),
					ObserverStream: make(chan RoomObserver, 5),
					Done:           make(chan bool),
				}
				roomThreads[g.Room.ID] = rt
				go rt.run()
			}
			g.Getter <- rt
		}
	}
}

func (r *Room) SendEvent(event Event) {
	rt := r.getThread()
	rt.EventStream <- event
}

func (r *Room) Listen() RoomSubscription {
	rt := r.getThread()
	eventListener := make(chan Event, 5)
	done := make(chan bool, 1)
	rt.ObserverStream <- RoomObserver{
		EventListener: eventListener,
		Done:          done,
	}
	return RoomSubscription{
		EventListener: eventListener,
		Done:          done,
	}
}

func (r *Room) ThreadDone(event Event) {
	rt := r.getThread()
	rt.Done <- true
}

func (r *Room) getThread() RoomThread {
	ch := make(chan RoomThread)
	g := roomThreadGetter{
		Room:   r,
		Getter: ch,
	}
	roomThreadAccessStream <- g
	rt := <-ch
	return rt
}

func (r *RoomThread) run() {
	observers := make([]RoomObserver, 0)
	for {
		select {
		case event := <-r.EventStream:
			new := make([]RoomObserver, 0)
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
