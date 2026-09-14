type Node struct {
	Val int
	Next *Node
	Prev *Node
}

type MyLinkedList struct {	
	Head *Node
	Tail *Node
	Size int
}


func Constructor() MyLinkedList {
	h := &Node{
		Val: -1,
		Next: nil,
		Prev: nil,
	}
	t := &Node{
		Val: -1,
		Next: nil,
		Prev: nil,
	}
	h.Next = t
	t.Prev = h
    return MyLinkedList{		
		Head: h,
		Tail: t,
		Size: 0,
	}
}


func (this *MyLinkedList) Get(index int) int {
	// validate index
	if index < 0 || index >= this.Size {
		return -1
	}
    cur := this.Head.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    this.AddAtIndex(0, val)
}


func (this *MyLinkedList) AddAtTail(val int)  {
	this.AddAtIndex(this.Size, val)
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	// validate index
	if index < 0 || index > this.Size {
		return
	}
	prev := this.Head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	next := prev.Next
	newNode := &Node{
		Val: val,
		Next: next,
		Prev: prev,
	}
	prev.Next = newNode
	next.Prev = newNode

	this.Size++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	// validate index
	if index < 0 || index >= this.Size {
		return
	}
	prev := this.Head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	next := prev.Next.Next
	prev.Next = next
	next.Prev = prev
	this.Size--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */