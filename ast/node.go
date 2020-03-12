package ast

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/position"
)

type Node struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
}

// SetPosition sets node position
func (n *Node) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Node) GetPosition() *position.Position {
	return n.Position
}

func (n *Node) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Root node
type Root struct {
	Node
	Stmts []Vertex
}

func (n *Root) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Stmts": n.Stmts,
	}
}

func (n *Root) Visit(v Visitor, depth int) bool {
	return v.Root(depth, n)
}

// Nullable node
type Nullable struct {
	Node
	Expr Vertex
}

func (n *Nullable) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *Nullable) Visit(v Visitor, depth int) bool {
	return v.Nullable(depth, n)
}

// Parameter node
type Parameter struct {
	Node
	ByRef        bool
	Variadic     bool
	Type Vertex
	Var     Vertex
	DefaultValue Vertex
}

func (n *Parameter) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Type": {n.Type},
		"Var":     {n.Var},
		"DefaultValue": {n.DefaultValue},
	}
}

func (n *Parameter) Visit(v Visitor, depth int) bool {
	return v.Parameter(depth, n)
}

// Identifier node
type Identifier struct {
	Node
	value string
}

func (n *Identifier) Children() map[string][]Vertex {
	return nil
}

func (n *Identifier) Visit(v Visitor, depth int) bool {
	return v.Identifier(depth, n)
}

// ArgumentList node
type ArgumentList struct {
	Node
	Arguments []Vertex
}

func (n *ArgumentList) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Arguments": n.Arguments,
	}
}

func (n *ArgumentList) Visit(v Visitor, depth int) bool {
	return v.ArgumentList(depth, n)
}

// Argument node
type Argument struct {
	Node
	Variadic    bool
	IsReference bool
	Expr        Vertex
}

func (n *Argument) Children() map[string][]Vertex {
	return map[string][]Vertex{
		"Expr": {n.Expr},
	}
}

func (n *Argument) Visit(v Visitor, depth int) bool {
	return v.Argument(depth, n)
}
