package stmt

import (
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/node"
	"github.com/VKCOM/noverify/src/php/parser/position"
	"github.com/VKCOM/noverify/src/php/parser/walker"
)

// Default node
type Default struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Stmts        []node.Node
}

// NewDefault node constructor
func NewDefault(Stmts []node.Node) *Default {
	return &Default{
		FreeFloating: nil,
		Stmts:        Stmts,
	}
}

// SetPosition sets node position
func (n *Default) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Default) GetPosition() *position.Position {
	return n.Position
}

func (n *Default) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Default) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.Stmts != nil {
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(v)
			}
		}
	}

	v.LeaveNode(n)
}
