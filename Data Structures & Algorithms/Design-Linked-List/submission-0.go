type Node struct {
    Val int
    Next *Node
}

type MyLinkedList struct {
    Head *Node
    Tail *Node
    Length int
}


func Constructor() MyLinkedList {
    return MyLinkedList{}
}


func (this *MyLinkedList) Get(index int) int {
    if index >= this.Length {
        return -1
    }
    cur := this.Head
    for i := 0; i < index; i++ {
        cur = cur.Next
        }
    return cur.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := &Node{
        Val: val, 
        Next: this.Head,
        }
    this.Head = newNode
    this.Length += 1
    if this.Length == 1 {
        this.Tail = newNode
    }
}

func (this *MyLinkedList) AddAtTail(val int)  {
    newNode := &Node{
        Val: val, 
        Next: nil,
        }
    if this.Length == 0 {
        this.Head = newNode
        this.Tail = newNode
    } else {
        this.Tail.Next = newNode
        this.Tail = newNode
    }
    this.Length++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index == this.Length {
        this.AddAtTail(val)
        return
    } else if index > this.Length {
        return
    } else if index == 0 {
        this.AddAtHead(val)
        return
    }else {
        cur := this.Head
        var prev *Node
        for i := 0; i < index; i++ {
        prev = cur
        cur = cur.Next
        }
        newNode := &Node{
            Val: val,
            Next: cur,
        }
        prev.Next = newNode
        this.Length++
    }
}

func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index >= this.Length {
        return
    } else if index == 0 {
        this.Head = this.Head.Next
        this.Length--
        if this.Length == 0 {
            this.Tail = nil
    }
    } else {
        cur := this.Head
        var prev *Node
        for i := 0; i < index; i++ {
        prev = cur
        cur = cur.Next
        }
        prev.Next = cur.Next
        if index == this.Length-1 {
            this.Tail = prev
        }
        this.Length--
    }
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
