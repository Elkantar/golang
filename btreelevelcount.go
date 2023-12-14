package piscine

/*
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else if data > root.Data {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}*/

func BTreeLevelCount(root *TreeNode) int {
	// count := 0
	if root == nil {
		return 0
	}
	Right := BTreeLevelCount(root.Right)
	Left := BTreeLevelCount(root.Left)

	if Left > Right {
		return Left + 1
	}
	return Right + 1
}

/*
func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	fmt.Println(BTreeLevelCount(root))
}*/
