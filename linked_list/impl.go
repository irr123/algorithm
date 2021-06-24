// there are exists "container"-package in std lib

package linked_list

type Node struct {
	// left, right
	l, r *Node
	data int
}

type List struct {
	// head, tail
	h, t *Node
}

type LTRIterator struct {
	// current
	c *Node
}

type RTLIterator struct {
	c *Node
}

func New() *List {
	return new(List)
}

func (h *List) Push(d int) {
	n := &Node{data: d}

	if h.h == nil {
		h.h = n
	}

	if h.t == nil {
		h.t = n
		return
	}

	h.t.r = n
	n.l = h.t

	h.t = n
}

func (h *List) Pop() int {
	if h.t == nil {
		return 0
	}

	n := h.t
	h.t = n.l
	if h.t == nil {
		h.h = nil
	}

	return n.data
}

func (h *List) Queue(d int) {
	n := &Node{data: d}

	if h.h == nil {
		h.h = n
	}

	if h.t == nil {
		h.t = n
		return
	}

	h.h.l = n
	n.r = h.h

	h.h = n

}

func (h *List) Dequeue() int {
	if h.h == nil {
		return 0
	}

	n := h.h
	h.h = n.r
	if h.h == nil {
		h.t = nil
	}

	return n.data
}

func (i *LTRIterator) Next() *Node {
	if i.c == nil {
		return nil
	}

	c := i.c
	i.c = i.c.l
	return c
}

func (i *RTLIterator) Next() *Node {
	if i.c == nil {
		return nil
	}

	c := i.c
	i.c = i.c.r
	return c
}
