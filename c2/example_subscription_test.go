package c2

import (
	"container/list"
	"fmt"
)

func ExampleSubscription() {
	listeners := list.New()
	listener := func() {}
	sub := NewSubscription(listeners, listeners.PushBack(listener))
	fmt.Printf("%d listeners before close\n", listeners.Len())
	sub.Close()
	fmt.Printf("%d listeners after close\n", listeners.Len())
	// Output:
	// 1 listeners before close
	// 0 listeners after close
}
