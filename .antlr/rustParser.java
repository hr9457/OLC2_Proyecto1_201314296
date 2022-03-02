// Generated from c:\Users\joshu\Documents\USAC\1S2022\COMPI2\LAB\Proyecto\OLC2_Proyecto1_201314296\rust.g4 by ANTLR 4.8
 
    import "Proyecto1/src/interfaces"
    import "Proyecto1/src/expressiones"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class rustParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, TK_PARENTESIS_LEFT=2, TK_PARENTESIS_RIGHT=3, TK_CORCHETE_LETF=4, 
		TK_CORCHETE_RIGHT=5, TK_LLAVE_LEFT=6, TK_LLAVE_RIGHT=7, TK_PUNTO_COMA=8, 
		TK_DOSPUNTOS=9, TK_COMA=10, TK_MENOR=11, TK_MAYOR=12, TK_PUNTO=13, TK_IGUAL=14, 
		TK_MAS=15, TK_MENOS=16, TK_POR=17, TK_DIVISION=18, TK_PORCENTAJE=19, TK_BARRA=20, 
		TK_AMPERSAND=21, TK_ADMIRACION=22, TK_MAYORIGULA=23, TK_MENOIGUAL=24, 
		TK_IGUALIGUAL=25, TK_DIFERENTE=26, TK_OR=27, TK_AND=28, TK_WHAT=29, TK_TIPORETURN=30, 
		TK_IMPRESION=31, TK_INT=32, TK_FLOAT=33, TK_BOOL=34, TK_CHAR=35, TK_STRING=36, 
		TK_MAIN=37, TK_USIZE=38, TK_LET=39, TK_MUT=40, TK_STRUCT=41, TK_AS=42, 
		TK_TRUE=43, TK_FALSE=44, TK_FN=45, TK_RETURN=46, TK_ABS=47, TK_SQRT=48, 
		TK_TOSTRING=49, TK_CLONE=50, TK_NEW=51, TK_LEN=52, TK_PUSH=53, TK_REMOVED=54, 
		TK_CONTAINS=55, TK_INSERT=56, TK_CAPACITY=57, TK_WITHCAPACITY=58, TK_IF=59, 
		TK_ELSE=60, TK_NUMBER=61, TK_DECIMAL=62, TK_CADENA=63, TK_CARACTER=64, 
		TK_ID=65, TK_COMMET=66, SPACES=67;
	public static final int
		RULE_start = 0, RULE_funcionmain = 1, RULE_instrucciones = 2, RULE_impresion = 3, 
		RULE_expresionIf = 4, RULE_sintaxisIf = 5, RULE_sintaxisElse = 6, RULE_variable = 7, 
		RULE_tipo = 8, RULE_expresion = 9, RULE_valor = 10;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "funcionmain", "instrucciones", "impresion", "expresionIf", 
			"sintaxisIf", "sintaxisElse", "variable", "tipo", "expresion", "valor"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'^'", "'('", "')'", "'['", "']'", "'{'", "'}'", "';'", "':'", 
			"','", "'<'", "'>'", "'.'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", 
			"'|'", "'&'", "'!'", "'>='", "'<='", "'=='", "'!='", "'||'", "'&&'", 
			"'?'", "'->'", "'println!'", "'i64'", "'f64'", "'bool'", "'char'", "'String'", 
			"'main'", "'usize'", "'let'", "'mut'", "'struct'", "'as'", "'true'", 
			"'false'", "'fn'", "'return'", "'abs'", "'sqrt'", "'to_string'", "'clone'", 
			"'new'", "'len'", "'push'", "'remove'", "'contains'", "'insert'", "'capacity'", 
			"'with_capacity'", "'if'", "'else'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "TK_PARENTESIS_LEFT", "TK_PARENTESIS_RIGHT", "TK_CORCHETE_LETF", 
			"TK_CORCHETE_RIGHT", "TK_LLAVE_LEFT", "TK_LLAVE_RIGHT", "TK_PUNTO_COMA", 
			"TK_DOSPUNTOS", "TK_COMA", "TK_MENOR", "TK_MAYOR", "TK_PUNTO", "TK_IGUAL", 
			"TK_MAS", "TK_MENOS", "TK_POR", "TK_DIVISION", "TK_PORCENTAJE", "TK_BARRA", 
			"TK_AMPERSAND", "TK_ADMIRACION", "TK_MAYORIGULA", "TK_MENOIGUAL", "TK_IGUALIGUAL", 
			"TK_DIFERENTE", "TK_OR", "TK_AND", "TK_WHAT", "TK_TIPORETURN", "TK_IMPRESION", 
			"TK_INT", "TK_FLOAT", "TK_BOOL", "TK_CHAR", "TK_STRING", "TK_MAIN", "TK_USIZE", 
			"TK_LET", "TK_MUT", "TK_STRUCT", "TK_AS", "TK_TRUE", "TK_FALSE", "TK_FN", 
			"TK_RETURN", "TK_ABS", "TK_SQRT", "TK_TOSTRING", "TK_CLONE", "TK_NEW", 
			"TK_LEN", "TK_PUSH", "TK_REMOVED", "TK_CONTAINS", "TK_INSERT", "TK_CAPACITY", 
			"TK_WITHCAPACITY", "TK_IF", "TK_ELSE", "TK_NUMBER", "TK_DECIMAL", "TK_CADENA", 
			"TK_CARACTER", "TK_ID", "TK_COMMET", "SPACES"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "rust.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }


	    var tipoValor = interfaces.NULL

	public rustParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public FuncionmainContext funcionmain() {
			return getRuleContext(FuncionmainContext.class,0);
		}
		public TerminalNode EOF() { return getToken(rustParser.EOF, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(22);
			funcionmain();
			setState(23);
			match(EOF);

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FuncionmainContext extends ParserRuleContext {
		public TerminalNode TK_FN() { return getToken(rustParser.TK_FN, 0); }
		public TerminalNode TK_MAIN() { return getToken(rustParser.TK_MAIN, 0); }
		public TerminalNode TK_PARENTESIS_LEFT() { return getToken(rustParser.TK_PARENTESIS_LEFT, 0); }
		public TerminalNode TK_PARENTESIS_RIGHT() { return getToken(rustParser.TK_PARENTESIS_RIGHT, 0); }
		public TerminalNode TK_LLAVE_LEFT() { return getToken(rustParser.TK_LLAVE_LEFT, 0); }
		public TerminalNode TK_LLAVE_RIGHT() { return getToken(rustParser.TK_LLAVE_RIGHT, 0); }
		public List<InstruccionesContext> instrucciones() {
			return getRuleContexts(InstruccionesContext.class);
		}
		public InstruccionesContext instrucciones(int i) {
			return getRuleContext(InstruccionesContext.class,i);
		}
		public FuncionmainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcionmain; }
	}

	public final FuncionmainContext funcionmain() throws RecognitionException {
		FuncionmainContext _localctx = new FuncionmainContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_funcionmain);
		int _la;
		try {
			setState(44);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(26);
				match(TK_FN);
				setState(27);
				match(TK_MAIN);
				setState(28);
				match(TK_PARENTESIS_LEFT);
				setState(29);
				match(TK_PARENTESIS_RIGHT);
				setState(30);
				match(TK_LLAVE_LEFT);
				setState(31);
				match(TK_LLAVE_RIGHT);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(32);
				match(TK_FN);
				setState(33);
				match(TK_MAIN);
				setState(34);
				match(TK_PARENTESIS_LEFT);
				setState(35);
				match(TK_PARENTESIS_RIGHT);
				setState(36);
				match(TK_LLAVE_LEFT);
				setState(40);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TK_IMPRESION) | (1L << TK_LET) | (1L << TK_IF))) != 0)) {
					{
					{
					setState(37);
					instrucciones();
					}
					}
					setState(42);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(43);
				match(TK_LLAVE_RIGHT);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionesContext extends ParserRuleContext {
		public VariableContext variable;
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public ImpresionContext impresion() {
			return getRuleContext(ImpresionContext.class,0);
		}
		public ExpresionIfContext expresionIf() {
			return getRuleContext(ExpresionIfContext.class,0);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instrucciones);
		try {
			setState(55);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_LET:
				enterOuterAlt(_localctx, 1);
				{
				setState(46);
				((InstruccionesContext)_localctx).variable = variable();
				 fmt.Println(((InstruccionesContext)_localctx).variable.nuevaVariable) 
				}
				break;
			case TK_IMPRESION:
				enterOuterAlt(_localctx, 2);
				{
				setState(49);
				impresion();

				}
				break;
			case TK_IF:
				enterOuterAlt(_localctx, 3);
				{
				setState(52);
				expresionIf();

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImpresionContext extends ParserRuleContext {
		public ExpresionContext expresion;
		public TerminalNode TK_IMPRESION() { return getToken(rustParser.TK_IMPRESION, 0); }
		public TerminalNode TK_PARENTESIS_LEFT() { return getToken(rustParser.TK_PARENTESIS_LEFT, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode TK_PARENTESIS_RIGHT() { return getToken(rustParser.TK_PARENTESIS_RIGHT, 0); }
		public TerminalNode TK_PUNTO_COMA() { return getToken(rustParser.TK_PUNTO_COMA, 0); }
		public ImpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_impresion; }
	}

	public final ImpresionContext impresion() throws RecognitionException {
		ImpresionContext _localctx = new ImpresionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_impresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			match(TK_IMPRESION);
			setState(58);
			match(TK_PARENTESIS_LEFT);
			setState(59);
			((ImpresionContext)_localctx).expresion = expresion(0);
			setState(60);
			match(TK_PARENTESIS_RIGHT);
			setState(61);
			match(TK_PUNTO_COMA);
			fmt.Println(((ImpresionContext)_localctx).expresion.primate)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpresionIfContext extends ParserRuleContext {
		public SintaxisIfContext sintaxisIf() {
			return getRuleContext(SintaxisIfContext.class,0);
		}
		public SintaxisElseContext sintaxisElse() {
			return getRuleContext(SintaxisElseContext.class,0);
		}
		public ExpresionIfContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresionIf; }
	}

	public final ExpresionIfContext expresionIf() throws RecognitionException {
		ExpresionIfContext _localctx = new ExpresionIfContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_expresionIf);
		try {
			setState(71);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(64);
				sintaxisIf();

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(67);
				sintaxisIf();
				setState(68);
				sintaxisElse();

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SintaxisIfContext extends ParserRuleContext {
		public TerminalNode TK_IF() { return getToken(rustParser.TK_IF, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode TK_LLAVE_LEFT() { return getToken(rustParser.TK_LLAVE_LEFT, 0); }
		public TerminalNode TK_LLAVE_RIGHT() { return getToken(rustParser.TK_LLAVE_RIGHT, 0); }
		public SintaxisIfContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sintaxisIf; }
	}

	public final SintaxisIfContext sintaxisIf() throws RecognitionException {
		SintaxisIfContext _localctx = new SintaxisIfContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_sintaxisIf);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			match(TK_IF);
			setState(74);
			expresion(0);
			setState(75);
			match(TK_LLAVE_LEFT);
			setState(76);
			match(TK_LLAVE_RIGHT);

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SintaxisElseContext extends ParserRuleContext {
		public TerminalNode TK_ELSE() { return getToken(rustParser.TK_ELSE, 0); }
		public TerminalNode TK_LLAVE_LEFT() { return getToken(rustParser.TK_LLAVE_LEFT, 0); }
		public TerminalNode TK_LLAVE_RIGHT() { return getToken(rustParser.TK_LLAVE_RIGHT, 0); }
		public SintaxisElseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sintaxisElse; }
	}

	public final SintaxisElseContext sintaxisElse() throws RecognitionException {
		SintaxisElseContext _localctx = new SintaxisElseContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_sintaxisElse);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(79);
			match(TK_ELSE);
			setState(80);
			match(TK_LLAVE_LEFT);
			setState(81);
			match(TK_LLAVE_RIGHT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VariableContext extends ParserRuleContext {
		public interface{} nuevaVariable;
		public Token TK_ID;
		public ExpresionContext expresion;
		public TipoContext tipo;
		public TerminalNode TK_LET() { return getToken(rustParser.TK_LET, 0); }
		public TerminalNode TK_MUT() { return getToken(rustParser.TK_MUT, 0); }
		public TerminalNode TK_ID() { return getToken(rustParser.TK_ID, 0); }
		public TerminalNode TK_IGUAL() { return getToken(rustParser.TK_IGUAL, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode TK_PUNTO_COMA() { return getToken(rustParser.TK_PUNTO_COMA, 0); }
		public TerminalNode TK_DOSPUNTOS() { return getToken(rustParser.TK_DOSPUNTOS, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public VariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable; }
	}

	public final VariableContext variable() throws RecognitionException {
		VariableContext _localctx = new VariableContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_variable);
		try {
			setState(117);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(83);
				match(TK_LET);
				setState(84);
				match(TK_MUT);
				setState(85);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(86);
				match(TK_IGUAL);
				setState(87);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(88);
				match(TK_PUNTO_COMA);
				_localctx.nuevaVariable = interfaces.NewSimbolo((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),(((VariableContext)_localctx).expresion!=null?_input.getText(((VariableContext)_localctx).expresion.start,((VariableContext)_localctx).expresion.stop):null),true,tipoValor)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(91);
				match(TK_LET);
				setState(92);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(93);
				match(TK_IGUAL);
				setState(94);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(95);
				match(TK_PUNTO_COMA);
				_localctx.nuevaVariable = interfaces.NewSimbolo((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),(((VariableContext)_localctx).expresion!=null?_input.getText(((VariableContext)_localctx).expresion.start,((VariableContext)_localctx).expresion.stop):null),false,tipoValor)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(98);
				match(TK_LET);
				setState(99);
				match(TK_MUT);
				setState(100);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(101);
				match(TK_DOSPUNTOS);
				setState(102);
				((VariableContext)_localctx).tipo = tipo();
				setState(103);
				match(TK_IGUAL);
				setState(104);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(105);
				match(TK_PUNTO_COMA);
				_localctx.nuevaVariable = interfaces.NewSimbolo((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),(((VariableContext)_localctx).expresion!=null?_input.getText(((VariableContext)_localctx).expresion.start,((VariableContext)_localctx).expresion.stop):null),true,((VariableContext)_localctx).tipo.tipoExp)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(108);
				match(TK_LET);
				setState(109);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(110);
				match(TK_DOSPUNTOS);
				setState(111);
				((VariableContext)_localctx).tipo = tipo();
				setState(112);
				match(TK_IGUAL);
				setState(113);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(114);
				match(TK_PUNTO_COMA);
				_localctx.nuevaVariable = interfaces.NewSimbolo((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),(((VariableContext)_localctx).expresion!=null?_input.getText(((VariableContext)_localctx).expresion.start,((VariableContext)_localctx).expresion.stop):null),false,((VariableContext)_localctx).tipo.tipoExp)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TipoContext extends ParserRuleContext {
		public interfaces.TipoExpression tipoExp;
		public TerminalNode TK_INT() { return getToken(rustParser.TK_INT, 0); }
		public TerminalNode TK_FLOAT() { return getToken(rustParser.TK_FLOAT, 0); }
		public TerminalNode TK_BOOL() { return getToken(rustParser.TK_BOOL, 0); }
		public TerminalNode TK_CHAR() { return getToken(rustParser.TK_CHAR, 0); }
		public TerminalNode TK_STRING() { return getToken(rustParser.TK_STRING, 0); }
		public TipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo; }
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_tipo);
		try {
			setState(129);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(119);
				match(TK_INT);
				_localctx.tipoExp = interfaces.INTEGER
				}
				break;
			case TK_FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(121);
				match(TK_FLOAT);
				_localctx.tipoExp = interfaces.FLOAT
				}
				break;
			case TK_BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(123);
				match(TK_BOOL);
				_localctx.tipoExp = interfaces.BOOLEAN
				}
				break;
			case TK_CHAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(125);
				match(TK_CHAR);
				_localctx.tipoExp = interfaces.CHAR
				}
				break;
			case TK_STRING:
				enterOuterAlt(_localctx, 5);
				{
				setState(127);
				match(TK_STRING);
				_localctx.tipoExp = interfaces.STRING
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpresionContext extends ParserRuleContext {
		public interface{} primate;
		public ValorContext valor;
		public TerminalNode TK_PARENTESIS_LEFT() { return getToken(rustParser.TK_PARENTESIS_LEFT, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode TK_PARENTESIS_RIGHT() { return getToken(rustParser.TK_PARENTESIS_RIGHT, 0); }
		public ValorContext valor() {
			return getRuleContext(ValorContext.class,0);
		}
		public TerminalNode TK_POR() { return getToken(rustParser.TK_POR, 0); }
		public TerminalNode TK_DIVISION() { return getToken(rustParser.TK_DIVISION, 0); }
		public TerminalNode TK_PORCENTAJE() { return getToken(rustParser.TK_PORCENTAJE, 0); }
		public TerminalNode TK_MAS() { return getToken(rustParser.TK_MAS, 0); }
		public TerminalNode TK_MENOS() { return getToken(rustParser.TK_MENOS, 0); }
		public TerminalNode TK_AMPERSAND() { return getToken(rustParser.TK_AMPERSAND, 0); }
		public TerminalNode TK_IGUALIGUAL() { return getToken(rustParser.TK_IGUALIGUAL, 0); }
		public TerminalNode TK_DIFERENTE() { return getToken(rustParser.TK_DIFERENTE, 0); }
		public TerminalNode TK_MAYOR() { return getToken(rustParser.TK_MAYOR, 0); }
		public TerminalNode TK_MENOR() { return getToken(rustParser.TK_MENOR, 0); }
		public TerminalNode TK_MAYORIGULA() { return getToken(rustParser.TK_MAYORIGULA, 0); }
		public TerminalNode TK_MENOIGUAL() { return getToken(rustParser.TK_MENOIGUAL, 0); }
		public TerminalNode TK_AND() { return getToken(rustParser.TK_AND, 0); }
		public TerminalNode TK_OR() { return getToken(rustParser.TK_OR, 0); }
		public TerminalNode TK_BARRA() { return getToken(rustParser.TK_BARRA, 0); }
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	}

	public final ExpresionContext expresion() throws RecognitionException {
		return expresion(0);
	}

	private ExpresionContext expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpresionContext _localctx = new ExpresionContext(_ctx, _parentState);
		ExpresionContext _prevctx = _localctx;
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_PARENTESIS_LEFT:
				{
				setState(132);
				match(TK_PARENTESIS_LEFT);
				setState(133);
				expresion(0);
				setState(134);
				match(TK_PARENTESIS_RIGHT);
				}
				break;
			case TK_TRUE:
			case TK_FALSE:
			case TK_NUMBER:
			case TK_DECIMAL:
			case TK_CADENA:
			case TK_CARACTER:
			case TK_ID:
				{
				setState(136);
				((ExpresionContext)_localctx).valor = valor();

				                    fmt.Println("->>>>")
				                    fmt.Println(((ExpresionContext)_localctx).valor.primate)
				                    _localctx.primate = ((ExpresionContext)_localctx).valor.primate
				                
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(181);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(179);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(141);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(142);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TK_POR) | (1L << TK_DIVISION) | (1L << TK_PORCENTAJE))) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(143);
						expresion(16);
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(144);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(145);
						_la = _input.LA(1);
						if ( !(_la==TK_MAS || _la==TK_MENOS) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(146);
						expresion(15);
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(147);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(148);
						match(TK_AMPERSAND);
						setState(149);
						expresion(14);
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(150);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(151);
						match(T__0);
						setState(152);
						expresion(13);
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(153);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(154);
						match(TK_IGUALIGUAL);
						setState(155);
						expresion(11);
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(156);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(157);
						match(TK_DIFERENTE);
						setState(158);
						expresion(10);
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(159);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(160);
						match(TK_MAYOR);
						setState(161);
						expresion(9);
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(162);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(163);
						match(TK_MENOR);
						setState(164);
						expresion(8);
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(165);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(166);
						match(TK_MAYORIGULA);
						setState(167);
						expresion(7);
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(168);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(169);
						match(TK_MENOIGUAL);
						setState(170);
						expresion(6);
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(171);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(172);
						match(TK_AND);
						setState(173);
						expresion(5);
						}
						break;
					case 12:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(174);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(175);
						match(TK_OR);
						setState(176);
						expresion(4);
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(177);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(178);
						match(TK_BARRA);
						}
						break;
					}
					} 
				}
				setState(183);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ValorContext extends ParserRuleContext {
		public interfaces.Expresion primate;
		public Token TK_NUMBER;
		public TerminalNode TK_NUMBER() { return getToken(rustParser.TK_NUMBER, 0); }
		public TerminalNode TK_DECIMAL() { return getToken(rustParser.TK_DECIMAL, 0); }
		public TerminalNode TK_TRUE() { return getToken(rustParser.TK_TRUE, 0); }
		public TerminalNode TK_FALSE() { return getToken(rustParser.TK_FALSE, 0); }
		public TerminalNode TK_CADENA() { return getToken(rustParser.TK_CADENA, 0); }
		public TerminalNode TK_CARACTER() { return getToken(rustParser.TK_CARACTER, 0); }
		public TerminalNode TK_ID() { return getToken(rustParser.TK_ID, 0); }
		public ValorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valor; }
	}

	public final ValorContext valor() throws RecognitionException {
		ValorContext _localctx = new ValorContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_valor);
		try {
			setState(198);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(184);
				((ValorContext)_localctx).TK_NUMBER = match(TK_NUMBER);
				_localctx.primate = expressiones.NewPrimito((((ValorContext)_localctx).TK_NUMBER!=null?((ValorContext)_localctx).TK_NUMBER.getText():null),interfaces.INTEGER)
				}
				break;
			case TK_DECIMAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(186);
				match(TK_DECIMAL);

				}
				break;
			case TK_TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(188);
				match(TK_TRUE);

				}
				break;
			case TK_FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(190);
				match(TK_FALSE);

				}
				break;
			case TK_CADENA:
				enterOuterAlt(_localctx, 5);
				{
				setState(192);
				match(TK_CADENA);

				}
				break;
			case TK_CARACTER:
				enterOuterAlt(_localctx, 6);
				{
				setState(194);
				match(TK_CARACTER);

				}
				break;
			case TK_ID:
				enterOuterAlt(_localctx, 7);
				{
				setState(196);
				match(TK_ID);

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 9:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 15);
		case 1:
			return precpred(_ctx, 14);
		case 2:
			return precpred(_ctx, 13);
		case 3:
			return precpred(_ctx, 12);
		case 4:
			return precpred(_ctx, 10);
		case 5:
			return precpred(_ctx, 9);
		case 6:
			return precpred(_ctx, 8);
		case 7:
			return precpred(_ctx, 7);
		case 8:
			return precpred(_ctx, 6);
		case 9:
			return precpred(_ctx, 5);
		case 10:
			return precpred(_ctx, 4);
		case 11:
			return precpred(_ctx, 3);
		case 12:
			return precpred(_ctx, 11);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3E\u00cb\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\7\3)\n\3\f\3\16\3,\13\3\3\3\5\3/\n\3\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\5\4:\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\5\6J\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\tx\n\t\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u0084\n\n\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\5\13\u008e\n\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\7\13\u00b6\n\13\f\13\16\13\u00b9\13\13\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u00c9\n\f\3\f\2\3\24\r\2\4\6"+
		"\b\n\f\16\20\22\24\26\2\4\3\2\23\25\3\2\21\22\2\u00df\2\30\3\2\2\2\4."+
		"\3\2\2\2\69\3\2\2\2\b;\3\2\2\2\nI\3\2\2\2\fK\3\2\2\2\16Q\3\2\2\2\20w\3"+
		"\2\2\2\22\u0083\3\2\2\2\24\u008d\3\2\2\2\26\u00c8\3\2\2\2\30\31\5\4\3"+
		"\2\31\32\7\2\2\3\32\33\b\2\1\2\33\3\3\2\2\2\34\35\7/\2\2\35\36\7\'\2\2"+
		"\36\37\7\4\2\2\37 \7\5\2\2 !\7\b\2\2!/\7\t\2\2\"#\7/\2\2#$\7\'\2\2$%\7"+
		"\4\2\2%&\7\5\2\2&*\7\b\2\2\')\5\6\4\2(\'\3\2\2\2),\3\2\2\2*(\3\2\2\2*"+
		"+\3\2\2\2+-\3\2\2\2,*\3\2\2\2-/\7\t\2\2.\34\3\2\2\2.\"\3\2\2\2/\5\3\2"+
		"\2\2\60\61\5\20\t\2\61\62\b\4\1\2\62:\3\2\2\2\63\64\5\b\5\2\64\65\b\4"+
		"\1\2\65:\3\2\2\2\66\67\5\n\6\2\678\b\4\1\28:\3\2\2\29\60\3\2\2\29\63\3"+
		"\2\2\29\66\3\2\2\2:\7\3\2\2\2;<\7!\2\2<=\7\4\2\2=>\5\24\13\2>?\7\5\2\2"+
		"?@\7\n\2\2@A\b\5\1\2A\t\3\2\2\2BC\5\f\7\2CD\b\6\1\2DJ\3\2\2\2EF\5\f\7"+
		"\2FG\5\16\b\2GH\b\6\1\2HJ\3\2\2\2IB\3\2\2\2IE\3\2\2\2J\13\3\2\2\2KL\7"+
		"=\2\2LM\5\24\13\2MN\7\b\2\2NO\7\t\2\2OP\b\7\1\2P\r\3\2\2\2QR\7>\2\2RS"+
		"\7\b\2\2ST\7\t\2\2T\17\3\2\2\2UV\7)\2\2VW\7*\2\2WX\7C\2\2XY\7\20\2\2Y"+
		"Z\5\24\13\2Z[\7\n\2\2[\\\b\t\1\2\\x\3\2\2\2]^\7)\2\2^_\7C\2\2_`\7\20\2"+
		"\2`a\5\24\13\2ab\7\n\2\2bc\b\t\1\2cx\3\2\2\2de\7)\2\2ef\7*\2\2fg\7C\2"+
		"\2gh\7\13\2\2hi\5\22\n\2ij\7\20\2\2jk\5\24\13\2kl\7\n\2\2lm\b\t\1\2mx"+
		"\3\2\2\2no\7)\2\2op\7C\2\2pq\7\13\2\2qr\5\22\n\2rs\7\20\2\2st\5\24\13"+
		"\2tu\7\n\2\2uv\b\t\1\2vx\3\2\2\2wU\3\2\2\2w]\3\2\2\2wd\3\2\2\2wn\3\2\2"+
		"\2x\21\3\2\2\2yz\7\"\2\2z\u0084\b\n\1\2{|\7#\2\2|\u0084\b\n\1\2}~\7$\2"+
		"\2~\u0084\b\n\1\2\177\u0080\7%\2\2\u0080\u0084\b\n\1\2\u0081\u0082\7&"+
		"\2\2\u0082\u0084\b\n\1\2\u0083y\3\2\2\2\u0083{\3\2\2\2\u0083}\3\2\2\2"+
		"\u0083\177\3\2\2\2\u0083\u0081\3\2\2\2\u0084\23\3\2\2\2\u0085\u0086\b"+
		"\13\1\2\u0086\u0087\7\4\2\2\u0087\u0088\5\24\13\2\u0088\u0089\7\5\2\2"+
		"\u0089\u008e\3\2\2\2\u008a\u008b\5\26\f\2\u008b\u008c\b\13\1\2\u008c\u008e"+
		"\3\2\2\2\u008d\u0085\3\2\2\2\u008d\u008a\3\2\2\2\u008e\u00b7\3\2\2\2\u008f"+
		"\u0090\f\21\2\2\u0090\u0091\t\2\2\2\u0091\u00b6\5\24\13\22\u0092\u0093"+
		"\f\20\2\2\u0093\u0094\t\3\2\2\u0094\u00b6\5\24\13\21\u0095\u0096\f\17"+
		"\2\2\u0096\u0097\7\27\2\2\u0097\u00b6\5\24\13\20\u0098\u0099\f\16\2\2"+
		"\u0099\u009a\7\3\2\2\u009a\u00b6\5\24\13\17\u009b\u009c\f\f\2\2\u009c"+
		"\u009d\7\33\2\2\u009d\u00b6\5\24\13\r\u009e\u009f\f\13\2\2\u009f\u00a0"+
		"\7\34\2\2\u00a0\u00b6\5\24\13\f\u00a1\u00a2\f\n\2\2\u00a2\u00a3\7\16\2"+
		"\2\u00a3\u00b6\5\24\13\13\u00a4\u00a5\f\t\2\2\u00a5\u00a6\7\r\2\2\u00a6"+
		"\u00b6\5\24\13\n\u00a7\u00a8\f\b\2\2\u00a8\u00a9\7\31\2\2\u00a9\u00b6"+
		"\5\24\13\t\u00aa\u00ab\f\7\2\2\u00ab\u00ac\7\32\2\2\u00ac\u00b6\5\24\13"+
		"\b\u00ad\u00ae\f\6\2\2\u00ae\u00af\7\36\2\2\u00af\u00b6\5\24\13\7\u00b0"+
		"\u00b1\f\5\2\2\u00b1\u00b2\7\35\2\2\u00b2\u00b6\5\24\13\6\u00b3\u00b4"+
		"\f\r\2\2\u00b4\u00b6\7\26\2\2\u00b5\u008f\3\2\2\2\u00b5\u0092\3\2\2\2"+
		"\u00b5\u0095\3\2\2\2\u00b5\u0098\3\2\2\2\u00b5\u009b\3\2\2\2\u00b5\u009e"+
		"\3\2\2\2\u00b5\u00a1\3\2\2\2\u00b5\u00a4\3\2\2\2\u00b5\u00a7\3\2\2\2\u00b5"+
		"\u00aa\3\2\2\2\u00b5\u00ad\3\2\2\2\u00b5\u00b0\3\2\2\2\u00b5\u00b3\3\2"+
		"\2\2\u00b6\u00b9\3\2\2\2\u00b7\u00b5\3\2\2\2\u00b7\u00b8\3\2\2\2\u00b8"+
		"\25\3\2\2\2\u00b9\u00b7\3\2\2\2\u00ba\u00bb\7?\2\2\u00bb\u00c9\b\f\1\2"+
		"\u00bc\u00bd\7@\2\2\u00bd\u00c9\b\f\1\2\u00be\u00bf\7-\2\2\u00bf\u00c9"+
		"\b\f\1\2\u00c0\u00c1\7.\2\2\u00c1\u00c9\b\f\1\2\u00c2\u00c3\7A\2\2\u00c3"+
		"\u00c9\b\f\1\2\u00c4\u00c5\7B\2\2\u00c5\u00c9\b\f\1\2\u00c6\u00c7\7C\2"+
		"\2\u00c7\u00c9\b\f\1\2\u00c8\u00ba\3\2\2\2\u00c8\u00bc\3\2\2\2\u00c8\u00be"+
		"\3\2\2\2\u00c8\u00c0\3\2\2\2\u00c8\u00c2\3\2\2\2\u00c8\u00c4\3\2\2\2\u00c8"+
		"\u00c6\3\2\2\2\u00c9\27\3\2\2\2\f*.9Iw\u0083\u008d\u00b5\u00b7\u00c8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}