type Link struct {
    val string
    prev *Link
    next *Link
}

type BrowserHistory struct {
    currPage *Link
}


func Constructor(homepage string) BrowserHistory {
    // currPage := &Link{
    //     val: homepage,       
    // }

    // prev := &Link{
    //     next: currPage
    // }
    
    // next := &Link{
    //     prev: currPage
    // }
   return BrowserHistory{
        currPage: &Link{
            val: homepage,
            prev: nil,
            next: nil,
        },
   } 
}


func (this *BrowserHistory) Visit(url string)  {
    newPage := &Link{
        val: url,
        prev: this.currPage,
        next: nil,
    }

    this.currPage.next = newPage
    this.currPage = newPage
}


func (this *BrowserHistory) Back(steps int) string {
    curr := this.currPage
        for range steps {
            if curr.prev == nil {
                break
            }
            curr = curr.prev
        }
    this.currPage = curr
    return curr.val
}


func (this *BrowserHistory) Forward(steps int) string {
    curr := this.currPage
        for range steps {
            if curr.next == nil {
                break
            }
            curr = curr.next
        }
    this.currPage = curr
    return curr.val
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */