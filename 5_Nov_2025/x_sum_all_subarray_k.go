package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	cnt, val    int
	priority    int64
	left, right *Node
	size        int
	sum         int64
}

func newNode(cnt, val int) *Node {
	n := &Node{cnt: cnt, val: val, priority: rand.Int63()}
	n.size = 1
	n.sum = int64(cnt) * int64(val)
	return n
}

func update(n *Node) {
	if n == nil {
		return
	}
	n.size = 1
	n.sum = int64(n.cnt) * int64(n.val)
	if n.left != nil {
		n.size += n.left.size
		n.sum += n.left.sum
	}
	if n.right != nil {
		n.size += n.right.size
		n.sum += n.right.sum
	}
}

func lessKey(aCnt, aVal, bCnt, bVal int) bool {
	if aCnt != bCnt {
		return aCnt < bCnt
	}
	return aVal < bVal
}
func equalKey(aCnt, aVal, bCnt, bVal int) bool {
	return aCnt == bCnt && aVal == bVal
}

func rotateRight(y *Node) *Node {
	x := y.left
	y.left = x.right
	x.right = y
	update(y)
	update(x)
	return x
}
func rotateLeft(x *Node) *Node {
	y := x.right
	x.right = y.left
	y.left = x
	update(x)
	update(y)
	return y
}

func insert(root *Node, cnt, val int) *Node {
	if root == nil {
		return newNode(cnt, val)
	}
	if equalKey(cnt, val, root.cnt, root.val) {
		root.cnt = cnt
		root.val = val
	} else if lessKey(cnt, val, root.cnt, root.val) {
		root.left = insert(root.left, cnt, val)
		if root.left.priority > root.priority {
			root = rotateRight(root)
		}
	} else {
		root.right = insert(root.right, cnt, val)
		if root.right.priority > root.priority {
			root = rotateLeft(root)
		}
	}
	update(root)
	return root
}

func merge(left, right *Node) *Node {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	if left.priority > right.priority {
		left.right = merge(left.right, right)
		update(left)
		return left
	} else {
		right.left = merge(left, right.left)
		update(right)
		return right
	}
}

func erase(root *Node, cnt, val int) *Node {
	if root == nil {
		return nil
	}
	if equalKey(cnt, val, root.cnt, root.val) {
		return merge(root.left, root.right)
	}
	if lessKey(cnt, val, root.cnt, root.val) {
		root.left = erase(root.left, cnt, val)
	} else {
		root.right = erase(root.right, cnt, val)
	}
	update(root)
	return root
}

func sumLargestX(root *Node, x int) int64 {
	if root == nil || x <= 0 {
		return 0
	}
	if root.size <= x {
		return root.sum
	}
	if root.right != nil {
		if root.right.size >= x {
			return sumLargestX(root.right, x)
		}
		res := root.right.sum
		if root.right.size+1 == x {
			res += int64(root.cnt) * int64(root.val)
			return res
		}
		res += int64(root.cnt) * int64(root.val)
		return res + sumLargestX(root.left, x-root.right.size-1)
	} else {
		if x == 1 {
			return int64(root.cnt) * int64(root.val)
		}
		res := int64(root.cnt) * int64(root.val)
		return res + sumLargestX(root.left, x-1)
	}
}

func findXSum(nums []int, k int, x int) []int64 {
	rand.Seed(time.Now().UnixNano())

	n := len(nums)
	if n == 0 || k == 0 {
		return nil
	}

	cnt := make(map[int]int)
	var root *Node

	push := func(v int) {
		prev := cnt[v]
		if prev > 0 {
			root = erase(root, prev, v)
		}
		cnt[v] = prev + 1
		root = insert(root, prev+1, v)
	}
	pop := func(v int) {
		prev := cnt[v]
		if prev > 0 {
			root = erase(root, prev, v)
			if prev-1 > 0 {
				cnt[v] = prev - 1
				root = insert(root, prev-1, v)
			} else {
				delete(cnt, v)
			}
		}
	}

	res := make([]int64, 0, n-k+1)
	for i := 0; i < n; i++ {
		push(nums[i])
		if i >= k-1 {
			if root == nil {
				res = append(res, 0)
			} else {
				xx := x
				if xx > root.size {
					xx = root.size
				}
				res = append(res, sumLargestX(root, xx))
			}
			out := nums[i-k+1]
			pop(out)
		}
	}
	return res
}

func main() {
	fmt.Println(findXSum([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2))
}
