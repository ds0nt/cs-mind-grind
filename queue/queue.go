package queue

import "fmt"

// Queue implementation off the top of my head:
// Okay so this is going to be my little cyclic chained queue invention..
type Queue struct {
	items       []interface{}
	popI, pushI int
	size        int
	next        *Queue
	full        bool
}

func NewQueue(chainsize int) *Queue {
	return &Queue{
		popI:  0,
		pushI: 0,
		items: make([]interface{}, chainsize),
		size:  chainsize,
		full:  false,
	}
}

func (q *Queue) Enqueue(val interface{}) {
	fmt.Println("\nEnqueue val, push, pop:", val, q.pushI, q.popI, q.next)
	// is the queue is full
	if q.full {
		// okay, so we stick it into our back-chained queue
		if q.next == nil {
			// chain on another queue
			q.next = NewQueue(q.size)
		}
		// enqueue into chain instead
		q.next.Enqueue(val)
		fmt.Println("Enqueue (in next)", q.pushI, q.popI, q.next)
		return

	}

	q.items[q.pushI] = val
	q.pushI = (q.pushI + 1) % q.size
	if q.pushI == q.popI {
		q.full = true
	}
	fmt.Println("Enqueue val, push, pop:", val, q.pushI, q.popI, q.next)
}

// so basically.. when we dequeue,
//  1. take item[start]
//  2. start++ % list size
//    a. does end - start
//  3. delete item and return
func (q *Queue) Dequeue() (val interface{}, ok bool) {
	fmt.Println("\nDequeue val, push, pop:", val, q.pushI, q.popI, q.next)
	// is our queue full?
	if q.full {
		// not anymore.
		q.full = false
	} else if q.popI == q.pushI {
		// do we have a chained queue?
		if q.next != nil {
			// awesome, replace this queue with that one
			*q = *q.next
			fmt.Println("Dequeue (from next)", q.pushI, q.popI, q.next)
			return q.Dequeue()
		}
		fmt.Println("Dequeue (no next)", q.pushI, q.popI, q.next)
		return
	}

	ok = true
	val = q.items[q.popI]
	q.items[q.popI] = nil
	q.popI = (q.popI + 1) % q.size
	fmt.Println("Dequeue val, push, pop:", val, q.pushI, q.popI, q.next)

	return
}
