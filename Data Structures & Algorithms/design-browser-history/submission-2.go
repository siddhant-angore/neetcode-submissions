type Page struct {
	Url  string
	Prev *Page
	Next *Page
}

type BrowserHistory struct {
	Head    *Page
	Tail    *Page
	Current *Page
}

func Constructor(homepage string) BrowserHistory {	
	dh := &Page{Url: ""}
	dt := &Page{Url: ""}
	page := &Page{
		Url:  homepage,
		Prev: dh,
		Next: dt,
	}
	dh.Next = page
	dt.Prev = page	
	return BrowserHistory{
		Head:    dh,
		Tail:    dt,
		Current: page,
	}
}

func (this *BrowserHistory) Visit(url string)  {
	// 1. Clear forward history by linking directly to Tail
	newPage := &Page{
		Url:  url,
		Prev: this.Current,
		Next: this.Tail,
	}
	this.Current.Next = newPage
	this.Tail.Prev = newPage
	
	// 2. Move current pointer to the new page
	this.Current = newPage
}

func (this *BrowserHistory) Back(steps int) string {
	// Use a for loop to move backward until steps run out OR we hit the first real page
	for steps > 0 && this.Current.Prev != this.Head {
		this.Current = this.Current.Prev
		steps--
	}
	return this.Current.Url
}

func (this *BrowserHistory) Forward(steps int) string {
	// Use a for loop to move forward until steps run out OR we hit the last real page
	for steps > 0 && this.Current.Next != this.Tail {
		this.Current = this.Current.Next
		steps--
	}
	return this.Current.Url
}
