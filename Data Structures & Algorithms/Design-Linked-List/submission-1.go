type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Length {
		return -1
	}

	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		Val:  val,
		Next: this.Head,
	}

	this.Head = newNode
	if this.Length == 0 {
		this.Tail = newNode
	}

	this.Length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{Val: val}

	if this.Length == 0 {
		this.Head = newNode
		this.Tail = newNode
	} else {
		this.Tail.Next = newNode
		this.Tail = newNode
	}

	this.Length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Length {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.Length {
		this.AddAtTail(val)
		return
	}

	prev := this.Head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}

	newNode := &Node{
		Val:  val,
		Next: prev.Next,
	}

	prev.Next = newNode
	this.Length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Length {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
		this.Length--

		if this.Length == 0 {
			this.Tail = nil
		}
		return
	}

	prev := this.Head
	for i := 0; i < index-1; i++ {
		prev = prev.Next
	}

	toDelete := prev.Next
	prev.Next = toDelete.Next

	if index == this.Length-1 {
		this.Tail = prev
	}

	this.Length--
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
