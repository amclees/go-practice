package list

type List struct {
  first *Node
  last *Node
  length int
}

type Node struct {
  previous *Node
  next *Node
  val int64
}

func (list *List) init(value int64) {
  node := Node{val: value}
  list.first = &node
  list.last = &node
}

func (list *List) Append(value int64) {
  list.length += 1
  if list.first == nil {
    list.init(value)
    return
  }

  node := Node{val: value, previous: list.last}

  list.last.next = &node
  list.last = &node
}

func (list *List) Prepend(value int64) {
  list.length += 1
  if list.first == nil {
    list.init(value)
    return
  }

  node := Node{val: value}

  node.next = list.first
  list.first.previous = &node
  list.first = &node
}

func (list *List) IndexOf(value int64) int {
  i := 0
  for current := list.first; current != nil; current = current.next {
    if current.val == value {
      return i
    }
    i += 1
  }
  return -1
}

func (list *List) IndexOfReverse(value int64) int {
  i := list.length - 1
  for current := list.last; current != nil; current = current.previous {
    if current.val == value {
      return i
    }
    i -= 1
  }
  return -1
}

func (list *List) Get(index int) (int64, bool) {
  i := 0
  for current := list.first; current != nil; current = current.next {
    if i == index {
      return current.val, true
    }
    i += 1
  }
  return 0, false
}

func (list *List) GetReverse(index int) (int64, bool) {
  i := list.length - 1
  for current := list.last; current != nil; current = current.previous {
    if i == index {
      return current.val, true
    }
    i -= 1
  }
  return 0, false
}

func (list *List) Remove(index int) bool {
  if list.first == nil {
    return false
  }

  if list.length == 1 {
    if index == 0 {
      list.first = nil
      list.last = nil
      list.length = 0
      return true
    } else {
      return false
    }
  }

  i := 0

  if i == index {
    list.first = list.first.next
    list.first.previous = nil
    list.length -= 1
    return true
  }

  for current := list.first; current.next != nil; current = current.next {
    if i == index - 1 && current.next != nil {
      list.length -= 1
      if current.next.next == nil {
        current.next = nil
        list.last = current
      } else {
        current.next = current.next.next
        current.next.previous = current
      }
      return true
    }
    i += 1
  }

  return false
}

func (list *List) RemoveReverse(index int) bool {
  if list.first == nil {
    return false
  }

  if list.length == 1 {
    if index == 0 {
      list.first = nil
      list.last = nil
      list.length = 0
      return true
    } else {
      return false
    }
  }

  i := list.length - 1

  if i == index {
    list.last = list.last.previous
    list.last.next = nil
    list.length -= 1
    return true
  }

  for current := list.last; current.previous != nil; current = current.previous {
    if i == index + 1 && current.previous != nil {
      list.length -= 1
      if current.previous.previous == nil {
        current.previous = nil
        list.first = current
      } else {
        current.previous = current.previous.previous
        current.previous.next = current
      }
      return true
    }
    i -= 1
  }

  return false
}

func (list *List) First() int64 {
  return list.first.val
}

func (list *List) Last() int64 {
  return list.last.val
}

func (list *List) Length() int {
  return list.length
}
