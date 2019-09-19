package space

import "sync"

func Do(r Host, f func()) {
	getThread(r).ActionStream <- Action{
		Func: f,
		Done: make(chan bool, 1),
	}
}

func DoAndWait(r Host, f func()) {
	done := make(chan bool)
	action := Action{
		Func: f,
		Done: done,
	}
	getThread(r).ActionStream <- action
	<-done
}

func SendEvent(r Host, event Message) {
	rt := getThread(r)
	rt.MessageStream <- event
}

func Publish(r Host) Publisher {
	return Publisher{thread: getThread(r)}
}

func Listen(r Host) Subscription {
	rt := getThread(r)
	eventListener := make(chan Message, 5)
	done := make(chan bool, 1)
	rt.SubscriberStream <- Subscriber{
		MessageListener: eventListener,
		Done:            done,
	}
	return Subscription{
		MessageListener: eventListener,
		Done:            done,
	}
}

func Attribute(r Host, event Message) *sync.Map {
	rt := getThread(r)
	return &rt.Attribute
}

func Done(r Host, event Message) {
	rt := getThread(r)
	rt.Done <- true
}
