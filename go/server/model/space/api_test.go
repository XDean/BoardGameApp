package space

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

type Mock int

func (m Mock) EventHostId() string {
	return string(m)
}

func TestListen(t *testing.T) {
	host := Mock(1)

	sub1 := Listen(host)
	msg1 := sub1.toSlice()

	publisher := Publish(host)
	publisher.SendEvent(Message{To: 1})

	sub2 := Listen(host)
	msg2 := sub2.toSlice()

	publisher.SendEvent(Message{To: 2})
	publisher.DoAndWait(func() {})

	sub1.Done <- true
	publisher.DoAndWait(func() {})

	publisher.SendEvent(Message{To: 3})

	publisher.DoAndWait(func() {})
	publisher.Done()

	result1 := <-msg1
	result2 := <-msg2

	assert.Equal(t, []Message{{To: 1}, {To: 2}}, result1)
	assert.Equal(t, []Message{{To: 2}, {To: 3}}, result2)
}

func (sub Subscription) toSlice() chan []Message {
	resultStream := make(chan []Message)
	go func() {
		result := []Message(nil)
		for {
			m, ok := <-sub.MessageListener
			if !ok {
				resultStream <- result
				break
			}
			result = append(result, m)
		}
	}()
	return resultStream
}
