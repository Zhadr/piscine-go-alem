package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	elm := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = elm
		l.Tail = elm
	} else {
		l.Tail.Next = elm
		l.Tail = elm
	}
}
