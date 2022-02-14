package tree

import "fmt"

//二叉树按层便利

type Node struct {
	Val         int
	Left, Right *Node
}

// 把元素放到队列中然后遍历每一个节点。拿到值并把节点继续放到队列后面
func Bfs(root *Node) []int {
	var ret []int
	var list []*Node
	list = append(list, root)

	var index = 0
	for index < len(list) {
		ret = append(ret, list[index].Val)
		if list[index].Left != nil {
			list = append(list, list[index].Left)
		}
		if list[index].Right != nil {
			list = append(list, list[index].Right)
		}

		index++
	}

	fmt.Println(ret)
	return ret
}
