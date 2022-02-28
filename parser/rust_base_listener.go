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

// EnterFuncionmain is called when production funcionmain is entered.
func (s *BaserustListener) EnterFuncionmain(ctx *FuncionmainContext) {}

// ExitFuncionmain is called when production funcionmain is exited.
func (s *BaserustListener) ExitFuncionmain(ctx *FuncionmainContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaserustListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaserustListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaserustListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaserustListener) ExitVariable(ctx *VariableContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaserustListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaserustListener) ExitTipo(ctx *TipoContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BaserustListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BaserustListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterValor is called when production valor is entered.
func (s *BaserustListener) EnterValor(ctx *ValorContext) {}

// ExitValor is called when production valor is exited.
func (s *BaserustListener) ExitValor(ctx *ValorContext) {}

// EnterImpresion is called when production impresion is entered.
func (s *BaserustListener) EnterImpresion(ctx *ImpresionContext) {}

// ExitImpresion is called when production impresion is exited.
func (s *BaserustListener) ExitImpresion(ctx *ImpresionContext) {}
