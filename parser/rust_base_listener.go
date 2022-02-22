// Code generated from rust.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // rust

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaserustListener is a complete listener for a parse tree produced by rustParser.
type BaserustListener struct{}

var _ rustListener = &BaserustListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaserustListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaserustListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaserustListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaserustListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaserustListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaserustListener) ExitStart(ctx *StartContext) {}

// EnterImpresion is called when production impresion is entered.
func (s *BaserustListener) EnterImpresion(ctx *ImpresionContext) {}

// ExitImpresion is called when production impresion is exited.
func (s *BaserustListener) ExitImpresion(ctx *ImpresionContext) {}
