type MinStack struct {
	st  []int
	mst []int
}

func Constructor() MinStack {
	return MinStack{
		st: []int{},
		mst: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.st = append(this.st, val)
	var lastVal int
	if len(this.mst) > 0 {
		lastVal = this.mst[len(this.mst)-1]
	} else {
		lastVal = val
	}
	this.mst = append(this.mst, min(lastVal, val))
}

func (this *MinStack) Pop() {
	this.st = this.st[:len(this.st)-1]
	this.mst = this.mst[:len(this.mst)-1]
}

func (this *MinStack) Top() int {
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	return this.mst[len(this.mst)-1]
}
