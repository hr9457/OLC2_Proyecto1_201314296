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

// EnterInstruccion is called when production instruccion is entered.
func (s *BaserustListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaserustListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterImpresion is called when production impresion is entered.
func (s *BaserustListener) EnterImpresion(ctx *ImpresionContext) {}

// ExitImpresion is called when production impresion is exited.
func (s *BaserustListener) ExitImpresion(ctx *ImpresionContext) {}

// EnterListprint is called when production listprint is entered.
func (s *BaserustListener) EnterListprint(ctx *ListprintContext) {}

// ExitListprint is called when production listprint is exited.
func (s *BaserustListener) ExitListprint(ctx *ListprintContext) {}

// EnterExpimprimir is called when production expimprimir is entered.
func (s *BaserustListener) EnterExpimprimir(ctx *ExpimprimirContext) {}

// ExitExpimprimir is called when production expimprimir is exited.
func (s *BaserustListener) ExitExpimprimir(ctx *ExpimprimirContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaserustListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaserustListener) ExitVariable(ctx *VariableContext) {}

// EnterAsignacionVariable is called when production asignacionVariable is entered.
func (s *BaserustListener) EnterAsignacionVariable(ctx *AsignacionVariableContext) {}

// ExitAsignacionVariable is called when production asignacionVariable is exited.
func (s *BaserustListener) ExitAsignacionVariable(ctx *AsignacionVariableContext) {}

// EnterDeclarcionarreglo is called when production declarcionarreglo is entered.
func (s *BaserustListener) EnterDeclarcionarreglo(ctx *DeclarcionarregloContext) {}

// ExitDeclarcionarreglo is called when production declarcionarreglo is exited.
func (s *BaserustListener) ExitDeclarcionarreglo(ctx *DeclarcionarregloContext) {}

// EnterListvalores is called when production listvalores is entered.
func (s *BaserustListener) EnterListvalores(ctx *ListvaloresContext) {}

// ExitListvalores is called when production listvalores is exited.
func (s *BaserustListener) ExitListvalores(ctx *ListvaloresContext) {}

// EnterExpresionIf is called when production expresionIf is entered.
func (s *BaserustListener) EnterExpresionIf(ctx *ExpresionIfContext) {}

// ExitExpresionIf is called when production expresionIf is exited.
func (s *BaserustListener) ExitExpresionIf(ctx *ExpresionIfContext) {}

// EnterListaelif is called when production listaelif is entered.
func (s *BaserustListener) EnterListaelif(ctx *ListaelifContext) {}

// ExitListaelif is called when production listaelif is exited.
func (s *BaserustListener) ExitListaelif(ctx *ListaelifContext) {}

// EnterElif is called when production elif is entered.
func (s *BaserustListener) EnterElif(ctx *ElifContext) {}

// ExitElif is called when production elif is exited.
func (s *BaserustListener) ExitElif(ctx *ElifContext) {}

// EnterExpresionWhile is called when production expresionWhile is entered.
func (s *BaserustListener) EnterExpresionWhile(ctx *ExpresionWhileContext) {}

// ExitExpresionWhile is called when production expresionWhile is exited.
func (s *BaserustListener) ExitExpresionWhile(ctx *ExpresionWhileContext) {}

// EnterExpresionLoop is called when production expresionLoop is entered.
func (s *BaserustListener) EnterExpresionLoop(ctx *ExpresionLoopContext) {}

// ExitExpresionLoop is called when production expresionLoop is exited.
func (s *BaserustListener) ExitExpresionLoop(ctx *ExpresionLoopContext) {}

// EnterBreakInst is called when production breakInst is entered.
func (s *BaserustListener) EnterBreakInst(ctx *BreakInstContext) {}

// ExitBreakInst is called when production breakInst is exited.
func (s *BaserustListener) ExitBreakInst(ctx *BreakInstContext) {}

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

// EnterListaArreglo is called when production listaArreglo is entered.
func (s *BaserustListener) EnterListaArreglo(ctx *ListaArregloContext) {}

// ExitListaArreglo is called when production listaArreglo is exited.
func (s *BaserustListener) ExitListaArreglo(ctx *ListaArregloContext) {}
