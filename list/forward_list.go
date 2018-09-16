package list

type ForwardList struct {
  first *ForwardNode
  length int
}

type ForwardNode struct {
  next *ForwardNode
  val int64
}

func (list *ForwardList) last() *ForwardNode {
  current := list.first
  for ; current.next != nil; current = current.next {}
  return current
}

func (list *ForwardList) Append(value int64) {
  list.length += 1
  node := ForwardNode{val: value}
  if list.first == nil {
    list.first = &node
    return
  }
  list.last().next = &node
}

func (list *ForwardList) Prepend(value int64) {
  list.length += 1
  node := ForwardNode{val: value, next: list.first}
  list.first = &node
}

func (list *ForwardList) IndexOf(value int64) int {
  i := 0
  for current := list.first; current != nil; current = current.next {
    if current.val == value {
      return i
    }
    i += 1
  }
  return -1
}

func (list *ForwardList) Get(index int) (int64, bool) {
  i := 0
  for current := list.first; current != nil; current = current.next {
    if i == index {
      return current.val, true
    }
    i += 1
  }
  return 0, false
}

func (list *ForwardList) Remove(index int) bool {
  if list.first == nil {
    return false
  }
  i := 0
  current := list.first
  for ; i < index - 1; current = current.next {
    if current.next == nil {
      return false
    }
    i += 1
  }
  current.next = current.next.next
  list.length -= 1
  return true
}

func (list *ForwardList) Length() int {
  return list.length
}
