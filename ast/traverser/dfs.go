package traverser

import "github.com/z7zmey/php-parser/ast"

type DFS struct {
	visitors []ast.Visitor
}

func NewDFS(visitors []ast.Visitor) *DFS {
	return &DFS{
		visitors: visitors,
	}
}

func (t *DFS) Traverse(n ast.Vertex) {
	t.traverse(n, 0)
}

func (t *DFS) traverse(n ast.Vertex, depth int) {
	if n == nil {
		return
	}

	for _, v := range t.visitors {
		r := n.Visit(v, depth)
		if !r {
			return
		}
	}

	for _, cc := range n.Children() {
		for _, c := range cc {
			t.traverse(c, depth+1)
		}
	}
}
