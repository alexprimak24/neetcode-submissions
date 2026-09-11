type ListNode struct {
    val int
	next *ListNode
}

type MyLinkedList struct {
    head *ListNode
    length int
}


func Constructor() MyLinkedList {
    return MyLinkedList{}
}


func (this *MyLinkedList) Get(index int) int {
   if index > this.length - 1 {
    return -1 
   }

   curr := this.head
   for i := 0; i < index; i++{
        curr = curr.next
   }

   return curr.val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    oldHead := this.head 

    newHead := &ListNode{
        val: val,
        next: oldHead,
    }
    // update head and length
    this.head = newHead
    this.length++
}


func (this *MyLinkedList) AddAtTail(val int)  {
    curr := this.head
    // grab tail
    for curr.next != nil {
        curr = curr.next
    }

    curr.next = &ListNode{
        val:val,
        next:nil,
    }
    this.length++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    curr := this.head
    for i := 0; i < index -1; i++{
            curr = curr.next
    }

    next := curr.next

    curr.next = &ListNode{
        val:val,
        next: next,
    }
    this.length++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    curr := this.head
    
    if index >= this.length {
        return
    }

    for i := 0; i < index -1; i++{
            curr = curr.next
    }

    curr.next = curr.next.next

    this.length--
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