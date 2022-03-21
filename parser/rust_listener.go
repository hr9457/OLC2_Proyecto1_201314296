// Code generated from rust.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // rust

import "github.com/antlr/antlr4/runtime/Go/antlr"

// rustListener is a complete listener for a parse tree produced by rustParser.
type rustListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterFuncionmain is called when entering the funcionmain production.
	EnterFuncionmain(c *FuncionmainContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterImpresion is called when entering the impresion production.
	EnterImpresion(c *ImpresionContext)

	// EnterListprint is called when entering the listprint production.
	EnterListprint(c *ListprintContext)

	// EnterExpimprimir is called when entering the expimprimir production.
	EnterExpimprimir(c *ExpimprimirContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterAsignacionVariable is called when entering the asignacionVariable production.
	EnterAsignacionVariable(c *AsignacionVariableContext)

	// EnterDeclarcionarreglo is called when entering the declarcionarreglo production.
	EnterDeclarcionarreglo(c *DeclarcionarregloContext)

	// EnterListvalores is called when entering the listvalores production.
	EnterListvalores(c *ListvaloresContext)

	// EnterExpresionIf is called when entering the expresionIf production.
	EnterExpresionIf(c *ExpresionIfContext)

	// EnterListaelif is called when entering the listaelif production.
	EnterListaelif(c *ListaelifContext)

	// EnterElif is called when entering the elif production.
	EnterElif(c *ElifContext)

	// EnterExpresionWhile is called when entering the expresionWhile production.
	EnterExpresionWhile(c *ExpresionWhileContext)

	// EnterExpresionLoop is called when entering the expresionLoop production.
	EnterExpresionLoop(c *ExpresionLoopContext)

	// EnterBreakInst is called when entering the breakInst production.
	EnterBreakInst(c *BreakInstContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterValor is called when entering the valor production.
	EnterValor(c *ValorContext)

	// EnterListaArreglo is called when entering the listaArreglo production.
	EnterListaArreglo(c *ListaArregloContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitFuncionmain is called when exiting the funcionmain production.
	ExitFuncionmain(c *FuncionmainContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitImpresion is called when exiting the impresion production.
	ExitImpresion(c *ImpresionContext)

	// ExitListprint is called when exiting the listprint production.
	ExitListprint(c *ListprintContext)

	// ExitExpimprimir is called when exiting the expimprimir production.
	ExitExpimprimir(c *ExpimprimirContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitAsignacionVariable is called when exiting the asignacionVariable production.
	ExitAsignacionVariable(c *AsignacionVariableContext)

	// ExitDeclarcionarreglo is called when exiting the declarcionarreglo production.
	ExitDeclarcionarreglo(c *DeclarcionarregloContext)

	// ExitListvalores is called when exiting the listvalores production.
	ExitListvalores(c *ListvaloresContext)

	// ExitExpresionIf is called when exiting the expresionIf production.
	ExitExpresionIf(c *ExpresionIfContext)

	// ExitListaelif is called when exiting the listaelif production.
	ExitListaelif(c *ListaelifContext)

	// ExitElif is called when exiting the elif production.
	ExitElif(c *ElifContext)

	// ExitExpresionWhile is called when exiting the expresionWhile production.
	ExitExpresionWhile(c *ExpresionWhileContext)

	// ExitExpresionLoop is called when exiting the expresionLoop production.
	ExitExpresionLoop(c *ExpresionLoopContext)

	// ExitBreakInst is called when exiting the breakInst production.
	ExitBreakInst(c *BreakInstContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitValor is called when exiting the valor production.
	ExitValor(c *ValorContext)

	// ExitListaArreglo is called when exiting the listaArreglo production.
	ExitListaArreglo(c *ListaArregloContext)
}
