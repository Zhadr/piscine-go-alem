package piscine

func ListPushFront(l *List, data interface{}) {
	elm := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = elm
		l.Tail = elm
	} else {
		elm.Next = l.Head

		l.Head = elm
	}
}
