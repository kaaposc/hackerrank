// Package extraLongFactorials solves Hackerrank challenge "Extra long factorials"[1].
//
// [1] https://www.hackerrank.com/challenges/extra-long-factorials/problem
package extralongfactorials

import "strconv"

const maxValue = 1000000

type Node struct {
	value  int32
	higher *Node
}

// multiply node value by num and add carry from previous node.
func (n *Node) multiply(num int32, carry int32) {
	if num <= 1 {
		return
	}

	n.value = n.value*num + carry

	overflow := n.value / maxValue
	n.value %= maxValue

	if n.higher != nil {
		n.higher.multiply(num, overflow)

		return
	}

	if overflow > 0 {
		n.higher = &Node{value: overflow, higher: nil}
	}
}

// toString returns string representation of the node value.
func (n *Node) toString() string {
	if n.higher == nil {
		return strconv.Itoa(int(n.value))
	}

	return n.higher.toString() + strconv.Itoa(int(n.value + maxValue))[1:]
}

// extraLongFactorials takes number and returns string representation of number's factorial.
func extraLongFactorials(n int32) string {
	node := Node{value: n, higher: nil}

	for n > 2 {
		n--
		node.multiply(n, 0)
	}

	return node.toString()
}
