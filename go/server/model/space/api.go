package space

import "sync"

func Publish(r Host) Publisher {
	return Publisher{thread: getThread(r)}
}

func Do(r Host, f func()) {
	Publish(r).Do(f)
}

func DoAndWait(r Host, f func()) {
	Publish(r).DoAndWait(f)
}

func SendMessage(r Host, event Message) {
	Publish(r).SendEvent(event)
}

func Listen(r Host) Subscription {
	return Publish(r).Listen()
}

func Attribute(r Host) *sync.Map {
	return Publish(r).Attribute()
}

func Done(r Host) {
	Publish(r).Done()
}

func (p Publisher) Attribute() *sync.Map {
	return &p.thread.Attribute
}

func (p Publisher) Listen() Subscription {
	eventListener := make(chan Message, 5)
	done := make(chan bool, 1)
	p.thread.EventStream <- Subscriber{
		MessageListener: eventListener,
		Done:            done,
	}
	return Subscription{
		MessageListener: eventListener,
		Done:            done,
	}
}

func (p Publisher) SendEvent(message Message) {
	p.thread.EventStream <- message
}

func (p Publisher) Done() {
	p.thread.Done <- true
}

func (p Publisher) Do(f func()) {
	p.thread.EventStream <- Action{
		Func: f,
		Done: make(chan bool, 1),
	}
}

func (p Publisher) DoAndWait(f func()) {
	done := make(chan bool)
	action := Action{
		Func: f,
		Done: done,
	}
	p.thread.EventStream <- action
	<-done
}
