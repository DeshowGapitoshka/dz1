package rbt

const (
	red       = true
	black     = false
	NEW       = "NEW"
	UPDATED   = "UPDATED"
	NOT_FOUND = "NOT_FOUND"
	ABSENT    = "ABSENT"
	ERASED    = "ERASED"
)

type Outdump struct {
	Key   string
	Value int64
}

type Node struct {
	key         string
	value       int64
	color       bool
	left, right *Node
	parent      *Node
}

type RBT struct {
	data *Node
	size int
}

func (b *RBT) insertFix(child *Node) {
	for child.parent != nil && child.parent.color == red {
		p, g := child.parent, child.parent.parent
		if p == g.left {
			u := g.right
			if u != nil && u.color == red {
				p.color, u.color, g.color = black, black, red
				child = g
			} else {
				if child == p.right {
					child = p
					b.rotateLeft(child)
				}
				child.parent.color, g.color = black, red
				b.rotateRight(g)
			}
		} else {
			u := g.left
			if u != nil && u.color == red {
				p.color, u.color, g.color = black, black, red
				child = g
			} else {
				if child == p.left {
					child = p
					b.rotateRight(child)
				}
				child.parent.color, g.color = black, red
				b.rotateLeft(g)
			}
		}
	}
	b.data.color = black
}

func (b *RBT) rotateLeft(x *Node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		b.data = y
	} else if x.parent.left == x {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (b *RBT) rotateRight(x *Node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		b.data = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y

}

func (b *RBT) Insert(key string, value int64) string {
	if b.data == nil {
		b.data = &Node{key: key, value: value, color: black}
		b.size++
		return NEW
	}
	current := b.data
	for {
		switch {
		case key == current.key:
			current.value = value
			return UPDATED
		case key < current.key:
			if current.left == nil {
				current.left = &Node{key: key, value: value, color: red, parent: current}
				b.size++
				b.insertFix(current.left)
				return NEW
			}
			current.left.parent = current
			current = current.left
		default:
			if current.right == nil {
				current.right = &Node{key: key, value: value, color: red, parent: current}
				b.size++
				b.insertFix(current.right)
				return NEW
			}
			current.right.parent = current
			current = current.right
		}
	}
}

func (b *RBT) Find(key string) (int64, bool) {
	current := b.data
	for current != nil {
		switch {
		case current.key == key:
			return current.value, true
		case key < current.key:
			current = current.left
		default:
			current = current.right
		}
	}
	return 0, false
}

func clr(n *Node) bool {
	if n == nil {
		return black
	}
	return n.color
}

func (r *RBT) eraseFix(x, parent *Node) {
	for x != r.data && clr(x) == black {
		if x == parent.left {
			b := parent.right
			if clr(b) == red { // брат ребёнка красный
				b.color = black
				parent.color = red
				r.rotateLeft(parent)
				b = parent.right
			}
			if b != nil && clr(b.left) == black && clr(b.right) == black && clr(b) == black { // оба ребёнка у брата чёрные
				b.color = red
				x = parent
				parent = x.parent
			} else {
				if b != nil && clr(b.left) == red && clr(b.right) == black { // у брата правый ребёнок чёрный,а левый красный
					b.color = red
					b.left.color = black
					r.rotateRight(b)
					b = parent.right
				} else if b != nil && clr(b.right) == red {
					if b.parent != nil {
						b.color = b.parent.color
						b.parent.color = black
					}
					if b.right != nil {
						b.right.color = black
					}
					r.rotateLeft(parent)
					x = r.data
				}
			}
		} else {
			b := parent.left
			if clr(b) == red {
				b.color = black
				parent.color = red
				r.rotateRight(parent)
				b = parent.left
			}
			if b != nil && clr(b.left) == black && clr(b.right) == black && clr(b) == black {
				b.color = red
				x = parent
				parent = x.parent
			} else {
				if b != nil && clr(b.left) == black && clr(b.right) == red {
					b.color = red
					b.right.color = black
					r.rotateLeft(b)
					b = parent.left
				} else if b != nil && clr(b.left) == red {
					if b.parent != nil {
						b.color = b.parent.color
						b.parent.color = black
					}
					if b.left != nil {
						b.left.color = black
					}
					r.rotateRight(parent)
					x = r.data
				}
			}
		}
	}
	if x != nil {
		x.color = black
	}

}

func (b *RBT) Erase(key string) string {
	current := b.data

	for current != nil && current.key != key {
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}

	if current == nil {
		return ABSENT
	}

	y := current // физически удаляемый узел
	originalColor := y.color

	var x *Node // ребёнок, который займёт место y
	var xParent *Node

	if current.left == nil || current.right == nil { // один ребенок
		if current.left != nil {
			x = current.left
		} else {
			x = current.right
		}

		xParent = current.parent

		if x != nil {
			x.parent = xParent
		}

		if xParent == nil {
			b.data = x
		} else if xParent.left == current {
			xParent.left = x
		} else {
			xParent.right = x
		}

	} else { // два ребенка
		y = current.right

		for y.left != nil {
			y = y.left
		}

		originalColor = y.color
		x = y.right

		if y.parent == current {
			xParent = y

			if x != nil {
				x.parent = y
			}
		} else {
			xParent = y.parent

			if x != nil { // убир приемника с его старого места
				x.parent = y.parent
			}

			y.parent.left = x

			y.right = current.right
			y.right.parent = y
		}

		y.parent = current.parent
		// cтавим приемника на место current
		if current.parent == nil {
			b.data = y
		} else if current.parent.left == current {
			current.parent.left = y
		} else {
			current.parent.right = y
		}

		y.left = current.left
		y.left.parent = y

		y.color = current.color
	}

	b.size--

	// если удалённый узел чёрный
	if originalColor == black {
		b.eraseFix(x, xParent)
	}

	return ERASED
}

func (b *RBT) Scale() int {
	return b.size
}

func (b *RBT) Dump() []Outdump {
	var data []*Node
	current := b.data
	var out []Outdump
	var temp Outdump
	for current != nil || len(data) > 0 {
		for current != nil {
			data = append(data, current)
			current = current.left
		}
		current = data[len(data)-1]
		data = data[:len(data)-1]

		temp.Key = current.key
		temp.Value = current.value
		out = append(out, temp)

		current = current.right
	}
	return out
}
