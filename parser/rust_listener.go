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

	// EnterImpresion is called when entering the impresion production.
	EnterImpresion(c *ImpresionContext)

	// EnterExpresionIf is called when entering the expresionIf production.
	EnterExpresionIf(c *ExpresionIfContext)

	// EnterSintaxisIf is called when entering the sintaxisIf production.
	EnterSintaxisIf(c *SintaxisIfContext)

	// EnterSintaxisElse is called when entering the sintaxisElse production.
	EnterSintaxisElse(c *SintaxisElseContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterValor is called when entering the valor production.
	EnterValor(c *ValorContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitFuncionmain is called when exiting the funcionmain production.
	ExitFuncionmain(c *FuncionmainContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitImpresion is called when exiting the impresion production.
	ExitImpresion(c *ImpresionContext)

	// ExitExpresionIf is called when exiting the expresionIf production.
	ExitExpresionIf(c *ExpresionIfContext)

	// ExitSintaxisIf is called when exiting the sintaxisIf production.
	ExitSintaxisIf(c *SintaxisIfContext)

	// ExitSintaxisElse is called when exiting the sintaxisElse production.
	ExitSintaxisElse(c *SintaxisElseContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitValor is called when exiting the valor production.
	ExitValor(c *ValorContext)
}
