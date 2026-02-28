package algorithm

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var treeNode = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	},
	Right: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	},
}

func AddTreeNode() {

	AddDown(treeNode)
}

func AddDown(tn *TreeNode) {

	all := []*TreeNode{tn}
	res := make([][]int, 0)
	for i := 0; len(all) != 0; i++ {
		res = append(res, []int{})
		tmp := make([]*TreeNode, 0)
		for j := 0; j < len(all); j++ {
			res[i] = append(res[i], all[j].Val)
			if all[j].Left != nil {
				tmp = append(tmp, all[j].Left)
			}
			if all[j].Right != nil {
				tmp = append(tmp, all[j].Right)
			}
		}
		all = tmp
	}
	fmt.Println(res)
	lastRes := res[len(res)-1]
	num := 0
	for i := 0; i < len(lastRes); i++ {
		num += lastRes[i]
	}
	fmt.Println(num)
}
