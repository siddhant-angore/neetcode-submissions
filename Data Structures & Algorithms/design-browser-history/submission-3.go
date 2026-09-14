type Page struct {
	Url string
	Prev *Page
	Next *Page
}

type BrowserHistory struct {
	Head *Page
	Tail *Page
	Current *Page
}


func Constructor(homepage string) BrowserHistory {	
    dh := &Page{Url: ""}
	dt := &Page{Url: ""}
	page := &Page{
		Url: homepage,
		Prev: dh,
		Next: dt,
	}
	dh.Next = page
	dt.Prev = page	
	return BrowserHistory{
		Head: dh,
		Tail: dt,
		Current: page,
	}
}


func (this *BrowserHistory) Visit(url string)  {
	newPage := &Page{
		Url: url,
		Prev: this.Current,
		Next: this.Tail,
	}
	this.Current.Next = newPage
	this.Tail.Prev = newPage

	this.Current = newPage
}


func (this *BrowserHistory) Back(steps int) string {
	for steps > 0 && this.Current.Prev != this.Head {
		this.Current = this.Current.Prev
		steps--
	}
	return this.Current.Url
}


func (this *BrowserHistory) Forward(steps int) string {
	for steps > 0 && this.Current.Next != this.Tail {
		this.Current = this.Current.Next
		steps--
	}
	return this.Current.Url
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */