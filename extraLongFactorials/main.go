package extraLongFactorials

import "strconv"

const maxValue = 1000000

type Node struct {
	value  int32
	higher *Node
}

func (n *Node) multiply(num int32, carry int32) {
	if num <= 1 {
		return
	}

	n.value = n.value*num + carry

	overflow := n.value / maxValue
	n.value = n.value % maxValue

	if n.higher != nil {
		n.higher.multiply(num, overflow)
	} else {
		if overflow > 0 {
			n.higher = &Node{value: overflow}
		}
	}
}

func (n *Node) toString() string {
	if n.higher == nil {
		return strconv.Itoa(int(n.value))
	}

	return n.higher.toString() + strconv.Itoa(int(n.value + maxValue))[1:]
}

func extraLongFactorials(n int32) string {
	node := Node{value: n}
	for n > 1 {
		n--
		node.multiply(n, 0)
	}

	return node.toString()
}
