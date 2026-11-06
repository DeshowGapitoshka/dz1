package bst

const (
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
	left, right *Node
}

type BST struct {
	data *Node
	size int
}

func (b *BST) Insert(key string, value int64) string {
	if b.data == nil {
		b.data = &Node{key: key, value: value}
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
				current.left = &Node{key: key, value: value}
				b.size++
				return NEW
			}
			current = current.left

		default:
			if current.right == nil {
				current.right = &Node{key: key, value: value}
				b.size++
				return NEW
			}
			current = current.right
		}
	}
}

func (b *BST) Find(key string) (int64, bool) {
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

func (b *BST) Erase(key string) string {
	var parent *Node
	current := b.data
	var count uint8
	for current != nil && current.key != key {
		parent = current
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}
	if current == nil {
		return ABSENT
	}
	if current.left != nil {
		count++
	}
	if current.right != nil {
		count++
	}
	switch {

	case current.left == nil && current.right == nil:
		if parent == nil {
			b.data = nil
		} else if parent.left == current {
			parent.left = nil
		} else {
			parent.right = nil
		}
		b.size--
		return ERASED

	case count == 1:
		var child *Node
		if current.left != nil {
			child = current.left
		} else {
			child = current.right
		}
		if parent == nil {
			b.data = child
		} else if parent.left == current {
			parent.left = child
		} else {
			parent.right = child
		}
		b.size--
		return ERASED
	case count == 2:
		temp := current
		temp_pr := current.right
		for temp_pr.left != nil {
			temp = temp_pr
			temp_pr = temp_pr.left
		}
		current.key = temp_pr.key
		current.value = temp_pr.value
		if temp.left == temp_pr {
			temp.left = temp_pr.right
		} else {
			temp.right = temp_pr.right
		}
		b.size--
		return ERASED

	}
	return ABSENT
}

func (b *BST) Scale() int {
	return b.size
}

func (b *BST) Dump() []Outdump {
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
