package ast_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/ast/traverser"
	"github.com/z7zmey/php-parser/ast/visitor"
)

func ExampleTest() {
	stxTree := &ast.Root{
		Stmts: []ast.Vertex{
			&ast.Nullable{
				Expr: &ast.Parameter{
					Type:         nil,
					Var:          nil,
					DefaultValue: nil,
				},
			},
			&ast.Identifier{},
			&ast.ArgumentList{
				Arguments: []ast.Vertex{
					&ast.Argument{},
					&ast.Argument{
						Expr: &ast.ScalarDnumber{},
					},
				},
			},
		},
	}

	v := &testVisitor{}
	t := traverser.NewDFS([]ast.Visitor{v})

	t.Traverse(stxTree)

	//output:
	//=>  Root
	//=>    Nullable
	//=>      Parameter
	//=>    Identifier
	//=>    ArgumentList
	//=>      Argument
	//=>      Argument
	//=>        ScalarDnumber
}

type testVisitor struct {
	visitor.Null
}

func (v *testVisitor) Root(depth int, _ *ast.Root) bool {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", depth), "Root")
	return true
}

func (v *testVisitor) Nullable(depth int, _ *ast.Nullable) bool {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", depth), "Nullable")
	return true
}

func (v *testVisitor) Parameter(depth int, _ *ast.Parameter) bool {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", depth), "Parameter")
	return true
}

func (v *testVisitor) Identifier(depth int, _ *ast.Identifier) bool {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", depth), "Identifier")
	return true
}

func (v *testVisitor) ArgumentList(depth int, _ *ast.ArgumentList) bool {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", depth), "ArgumentList")
	return true
}

func (v *testVisitor) Argument(depth int, _ *ast.Argument) bool {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", depth), "Argument")
	return true
}

func (v *testVisitor) ScalarDnumber(depth int, _ *ast.ScalarDnumber) bool {
	fmt.Fprintln(os.Stdout, "=>", strings.Repeat("  ", depth), "ScalarDnumber")
	return true
}
