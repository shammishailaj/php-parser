package graph

import (
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/traverser"
)

type AST struct {
	FileData  []byte
	Positions PositionStorage
	Nodes     NodeStorage
	Edges     EdgeStorage
	Tokens    TokenStorage
	RootNode  NodeID

	queue []queueItem
}

type queueItem struct {
	id    NodeID
	depth int
}

func (g *AST) Reset() {
	g.FileData = g.FileData[:0]
	g.Nodes = g.Nodes[:0]
	g.Edges = g.Edges[:0]
	g.Positions = g.Positions[:0]
	g.Tokens = g.Tokens[:0]
	g.RootNode = 0
}

func (g *AST) Link(nodeID NodeID, edgeType EdgeType, target uint32) {
	edge := Edge{
		Type:   edgeType,
		Target: target,
	}

	edgeID := g.Edges.Put(edge)

	nodeEdges := &g.Nodes[nodeID-1].Edges

	if nodeEdges.First == 0 {
		nodeEdges.First = edgeID
		nodeEdges.Last = edgeID
	} else {
		g.Edges[nodeEdges.Last-1].next = edgeID
		nodeEdges.Last = edgeID
	}
}

func (g *AST) AppendEdges(src EdgeList, edges EdgeList) EdgeList {
	if edges.First == 0 {
		return src
	}

	if src.First == 0 {
		return edges
	}

	g.Edges[src.Last-1].next = edges.First
	src.Last = edges.Last

	return src
}

func (g *AST) RemoveEdges(nodeID NodeID, f EdgeFilter) EdgeList {
	nodeEdges := &g.Nodes[nodeID-1].Edges

	var removedEdges EdgeList
	var prevEdgeID EdgeID

	edgeID := nodeEdges.First
	for edgeID != 0 {
		edge := &g.Edges[edgeID-1]

		if f(*edge) {
			if prevEdgeID == 0 {
				nodeEdges.First = edge.next
			} else {
				g.Edges[prevEdgeID-1].next = edge.next
			}

			removedEdges = g.AppendEdges(removedEdges, EdgeList{edgeID, edgeID})
		} else {
			prevEdgeID = edgeID
		}

		edgeID = edge.next
	}

	return removedEdges
}

func (g *AST) EachEdge(edges EdgeList, callback func(e Edge) bool) {
	edgeID := edges.First
	for edgeID != 0 {
		edge := g.Edges[edgeID-1]

		if callback(edge) {
			return
		}

		edgeID = edge.next
	}
}

func (g *AST) Traverse(v traverser.Visitor) {
	g.queue = g.queue[:0]
	g.queue = append(g.queue, queueItem{
		id:    g.RootNode,
		depth: 0,
	})

	for {
		if len(g.queue) == 0 {
			break
		}

		item := g.queue[len(g.queue)-1]
		g.queue = g.queue[:len(g.queue)-1]

		graphNode := g.Nodes.Get(item.id)
		depth := item.depth

		astNode := ast.Node{
			Type:   graphNode.Type,
			Flags:  graphNode.Flag,
			Tokens: make(map[ast.TokenGroup][]ast.Token),
		}

		g.EachEdge(graphNode.Edges, func(e Edge) bool {
			if e.Type != EdgeTypePosition {
				return false
			}

			posID := PositionID(e.Target)
			astNode.Position = g.Positions.Get(posID)

			return true
		})

		g.EachEdge(graphNode.Edges, func(e Edge) bool {
			if e.Type != EdgeTypeToken {
				return false
			}

			tokenID := TokenID(e.Target)

			token := g.Tokens.Get(tokenID)
			tokenPos := g.Positions.Get(token.Pos)

			nestedToken := ast.Token{
				Type:  token.Type,
				Value: string(g.FileData[tokenPos.PS:tokenPos.PE]),
			}

			astNode.Tokens[token.Group] = append(astNode.Tokens[token.Group], nestedToken)

			return false
		})

		visitChild := v.VisitNode(astNode, graphNode.Group, depth)

		if visitChild {
			depth++
			g.EachEdge(graphNode.Edges, func(e Edge) bool {
				if e.Type != EdgeTypeNode {
					return false
				}

				g.queue = append(g.queue, queueItem{
					id:    NodeID(e.Target),
					depth: depth,
				})

				return false
			})
		}

	}
}
