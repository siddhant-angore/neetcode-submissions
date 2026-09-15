type Node struct {
	Val int
	Prev *Node
	Next *Node
}

type MyStack struct {
	Head *Node
	Tail *Node
	Size int
}

func Constructor() MyStack {
	dh := &Node{Val: -1}
	dt := &Node{Val: -1}
	dh.Next = dt
	dt.Prev = dh
	return MyStack{
		Head: dh,
		Tail: dt,
		Size: 0,
	}
}

func (this *MyStack) Push(x int) {
	newNode := &Node{Val: x}
	lastNode := this.Tail.Prev
	lastNode.Next = newNode
	newNode.Prev = lastNode
	newNode.Next = this.Tail
	this.Tail.Prev = newNode
	this.Size++
}

func (this *MyStack) Pop() int {	
	if this.Empty() {
		return -1
	}
	lastNode := this.Tail.Prev
	val := lastNode.Val
	prevNode := lastNode.Prev
	prevNode.Next = this.Tail
	this.Tail.Prev = prevNode
	this.Size--
	return val	
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}
	return this.Tail.Prev.Val
}

func (this *MyStack) Empty() bool {
	return this.Size == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
