package dfsbfs

import (
	"fmt"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//bfs实现
	queue := []*TreeNode{}
	nLqueue := []*TreeNode{}
	minDepth := 1
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:len(queue)]
		//有节点为叶子结点 返回最小深度
		if node.Left == nil && node.Right == nil {
			return minDepth
		}
		if node.Left != nil {
			nLqueue = append(nLqueue, node.Left)
		}
		if node.Right != nil {
			nLqueue = append(nLqueue, node.Right)
		}
		//当前层遍历完成 开始遍历下一层节点
		if len(queue) == 0 {
			queue = nLqueue
			nLqueue = []*TreeNode{}
			minDepth++
		}
	}
	return minDepth

}

func TestMinDepth(t *testing.T) {
	root := &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}}}
	t.Log(minDepth(root))
	fmt.Println(minDepth(root))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//dfs实现
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depth := 0
	if root.Left != nil {
		depth = minDepth(root.Left)
	}
	if root.Right != nil {
		a := minDepth(root.Right)
		if (depth != 0 && a < depth) || depth == 0 {
			depth = a
		}
	}
	return 1 + depth
}
