package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	repl := node.Parent
	if node == repl.Left {
		repl.Left = rplc
		rplc.Parent = repl
	} else if node == repl.Right {
		repl.Right = rplc
		rplc.Parent = repl
	}
	return root
}
