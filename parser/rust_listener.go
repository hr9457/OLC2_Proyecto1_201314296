// Code generated from rust.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // rust

import "github.com/antlr/antlr4/runtime/Go/antlr"

// rustListener is a complete listener for a parse tree produced by rustParser.
type rustListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterImpresion is called when entering the impresion production.
	EnterImpresion(c *ImpresionContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitImpresion is called when exiting the impresion production.
	ExitImpresion(c *ImpresionContext)
}
