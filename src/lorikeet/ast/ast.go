package ast

import (
	"bytes"
	"fmt"
	"lorikeet/token"
	"strings"
)

// Node base interface
type Node interface {
	TokenLiteral() string
	String() string
	Line() int
}

// Statement nodes implement this
type Statement interface {
	Node
	statementNode()
}

// Expression nodes implement this
type Expression interface {
	Node
	expressionNode()
}

// Program struct
type Program struct {
	Statements []Statement
}

// TokenLiteral return literal for program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Line return line number
func (p *Program) Line() int { return 0 }

// Statements

// LetStatement node
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
	Mut   bool
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral return literal for let statement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Line return line number
func (ls *LetStatement) Line() int { return ls.Token.Line }

// ReturnStatement node
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral return literal for return statement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// Line return line number
func (rs *ReturnStatement) Line() int { return rs.Token.Line }

// ExpressionStatement node
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral return literal for expression Statement
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Line return line number
func (es *ExpressionStatement) Line() int { return es.Token.Line }

// BlockStatement node
type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral return literal for block statement
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Line return line number
func (bs *BlockStatement) Line() int { return bs.Token.Line }

// MutStatement node
type MutStatement struct {
	Token token.Token // the ASSIGN token
	Name  *Identifier
	Value Expression
}

func (ms *MutStatement) statementNode() {}

// TokenLiteral return literal for mut statement
func (ms *MutStatement) TokenLiteral() string { return ms.Token.Literal }
func (ms *MutStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ms.Name.String())
	out.WriteString(" = ")
	out.WriteString(ms.TokenLiteral() + " ")

	if ms.Value != nil {
		out.WriteString(ms.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Line return line number
func (ms *MutStatement) Line() int { return ms.Token.Line }

// Expressions

// Identifier node
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral return literal for identifier
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// Line return line number
func (i *Identifier) Line() int { return i.Token.Line }

// Boolean node
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral return literal for boolean
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// Line return line number
func (b *Boolean) Line() int { return b.Token.Line }

// IntegerLiteral node
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral return literal for integer literal
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// Line return line number
func (il *IntegerLiteral) Line() int { return il.Token.Line }

// FloatLiteral node
type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (il *FloatLiteral) expressionNode() {}

// TokenLiteral return literal for integer literal
func (il *FloatLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *FloatLiteral) String() string       { return il.Token.Literal }

// Line return line number
func (il *FloatLiteral) Line() int { return il.Token.Line }

// PrefixExpression node
type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral return literal for prefix expression
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// Line return line number
func (pe *PrefixExpression) Line() int { return pe.Token.Line }

// InfixExpression node
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral return literal for Infix expression
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Line return line number
func (ie *InfixExpression) Line() int { return ie.Token.Line }

// IfExpression node
type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral return literal for if expression
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// Line return line number
func (ie *IfExpression) Line() int { return ie.Token.Line }

// FunctionLiteral node
type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
	Name       string
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral return literal for function Literal
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	if fl.Name != "" {
		out.WriteString(fmt.Sprintf("<%s>", fl.Name))
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// Line return line number
func (fl *FunctionLiteral) Line() int { return fl.Token.Line }

// CallExpression node
type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

// TokenLiteral return literal for call expression
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// Line return line number
func (ce *CallExpression) Line() int { return ce.Token.Line }

// StringLiteral node
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral return literal for string
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

// Line return line number
func (sl *StringLiteral) Line() int { return sl.Token.Line }

// ArrayLiteral node
type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

// TokenLiteral return literal for array
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// Line return line number
func (al *ArrayLiteral) Line() int { return al.Token.Line }

// IndexExpression node
type IndexExpression struct {
	Token token.Token // the [ token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}

// TokenLiteral return literal for index
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

// Line return line number
func (ie *IndexExpression) Line() int { return ie.Token.Line }

// HashLiteral node
type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode() {}

// TokenLiteral return literal for hash
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// Line return line number
func (hl *HashLiteral) Line() int { return hl.Token.Line }

// MacroLiteral node
type MacroLiteral struct {
	Token      token.Token // The 'macro' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (ml *MacroLiteral) expressionNode() {}

// TokenLiteral return literal for macro
func (ml *MacroLiteral) TokenLiteral() string { return ml.Token.Literal }
func (ml *MacroLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range ml.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(ml.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(ml.Body.String())

	return out.String()
}

// Line return line number
func (ml *MacroLiteral) Line() int { return ml.Token.Line }
