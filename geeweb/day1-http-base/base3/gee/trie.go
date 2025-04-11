package gee

type node struct {
	pattern  string  //待匹配路由
	part     string  //路由中的一部分，例如：lang
	children []*node //子节点，例如[doc,tutorial,intro]
	isWild   bool    //是否精确匹配，part含有：或*时为true
}

// 第一个匹配的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
