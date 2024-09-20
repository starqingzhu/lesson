package algorithm

type (
	any     = interface{}
	Element struct {
		//pre  *ListNode
		Next *Element
		Data any
	}

	List struct {
		Head *Element
		Size int
	}
)

func (l *List) Add(d any) {
	newE := &Element{Data: d}
	if l.Head == nil {
		l.Head = newE
		l.Size++
		return
	}

	last := l.Head
	for last.Next != nil {
		last = last.Next
	}

	last.Next = newE
	l.Size++
}

func (l *List) Remove(data int) bool {
	return true
}
