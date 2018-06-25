package queue

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
	// is the queue is full
	if q.full {
		// okay, so we stick it into our back-chained queue
		if q.next == nil {
			// chain on another queue
			q.next = NewQueue(q.size)

			// if we were cool, we might make q.next.size = q.size*2
			// because right now, the enqueue operation is O(n/chainsize) (gotta hop queues to the end)
			// i suppose i could hold a pointer to the tail which would make it O(1)
			// or i could double list sizes and have it be roughly O(log n)
			// but we are not cool enough
		}
		// enqueue into chain instead
		q.next.Enqueue(val)
		return

	}

	q.items[q.pushI] = val
	q.pushI = (q.pushI + 1) % q.size
	if q.pushI == q.popI {
		q.full = true
	}
}

// so basically.. when we dequeue,
//  1. take item[start]
//  2. start++ % list size
//    a. does end - start
//  3. delete item and return
func (q *Queue) Dequeue() (val interface{}, ok bool) {
	// is our queue full?
	if q.full {
		// not anymore.
		q.full = false
	} else if q.popI == q.pushI {
		// do we have a chained queue?
		if q.next != nil {
			// awesome, replace this queue with that one
			*q = *q.next
			return q.Dequeue()
		}
		return
	}

	ok = true
	val = q.items[q.popI]
	q.items[q.popI] = nil
	q.popI = (q.popI + 1) % q.size

	return
}
