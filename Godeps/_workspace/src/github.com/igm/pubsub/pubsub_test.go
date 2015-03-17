package pubsub_test

import (
	"github.com/igm/pubsub"
	"fmt"
	"testing"
	"time"
)

func TestSubReader(t *testing.T) {
	pub := pubsub.Publisher{}
	r, last := pub.SubReader()
	if last != nil {
		t.Fatal("last message shuold be nil")
	}
	pub.Publish("msg 1")
	pub.Publish("msg 2")
	if r.Read() != "msg 1" {
		t.Fatal("incorrect message received")
	}
	if r.Read() != "msg 2" {
		t.Fatal("incorrect message received")
	}
}

func TestMultipleSubReaders(t *testing.T) {
	pub := pubsub.Publisher{}
	r, last := pub.SubReader()
	if last != nil {
		t.Fatal("last message shuold be nil")
	}
	pub.Publish("msg 1")

	r2, last2 := pub.SubReader()
	if last2 != "msg 1" {
		t.Fatal("last message shuold be nil")
	}

	pub.Publish("msg 2")
	if r.Read() != "msg 1" {
		t.Fatal("incorrect message received")
	}
	if r.Read() != "msg 2" {
		t.Fatal("incorrect message received")
	}
	if r2.Read() != "msg 2" {
		t.Fatal("incorrect message received")
	}
}

func TestSubChannel(t *testing.T) {
	pub := pubsub.Publisher{}
	ch, last := pub.SubChannel(nil)
	if last != nil {
		t.Fatal("last message shuold be nil")
	}
	pub.Publish("msg 1")
	pub.Publish("msg 2")
	pub.Publish(nil)

	if "msg 1" != <-ch {
		t.Fatal("incorrect message received")
	}
	if "msg 2" != <-ch {
		t.Fatal("incorrect message received")
	}
	if nil != <-ch {
		t.Fatal("incorrect message received")
	}
	if _, ok := <-ch; ok {
		t.Fatal("channel should be closed")
	}
}

func ExamplePublisher() {
	pub := pubsub.Publisher{}
	ch, _ := pub.SubChannel(nil)
	pub.Publish("msg 1")
	pub.Publish("msg 2")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// Output:
	// msg 1
	// msg 2
}

func ExamplePublisher_SubChannel() {
	pub := new(pubsub.Publisher)
	ch, _ := pub.SubChannel("close")
	pub.Publish("msg 1")
	pub.Publish("msg 2")
	pub.Publish("close")

	for msg := range ch {
		fmt.Println(msg)
	}
	// Output:
	// msg 1
	// msg 2
	// close
}

func ExamplePublisher_SubReader() {
	type Timer struct {
		pubsub.Publisher
	}

	timer := new(Timer)

	go func() {
		for /*...*/ {
			time.Sleep(time.Second)
			timer.Publish(time.Now())
			// ...
		}
	}()

	reader, _ := timer.SubReader()
	for {
		fmt.Println(reader.Read())
	}
}

func ExamplePublisher_composition() {
	type StockExchange struct {
		pubsub.Publisher
	}
	se := new(StockExchange)

	ch, _ := se.SubChannel(nil)
	se.Publish("msg 1")
	se.Publish("msg 2")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// Output:
	// msg 1
	// msg 2
}

func BenchmarkPubSub(b *testing.B) {
	pub := new(pubsub.Publisher)
	r1, _ := pub.SubReader()
	r2, _ := pub.SubReader()
	for i := 0; i < b.N; i++ {
		pub.Publish(i)
		v1 := r1.Read()
		v2 := r2.Read()
		if v1 != v2 || v1 != i {
			b.Fatal("incorrect value")
		}
	}
}
