// Generated from c:\Users\joshu\Documents\USAC\1S2022\COMPI2\LAB\Proyecto\OLC2_Proyecto1_201314296\rust.g4 by ANTLR 4.8
 
    import "Proyecto1/src/interfaces"
    import "Proyecto1/src/expressiones"
    import "Proyecto1/src/instrucciones"
    // import "Proyecto1/src/funciones"
    import arrayList "github.com/colegno/arraylist"
    // import "Proyecto1/src/pruebas" 
    // import "reflect"

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
		TK_PARENTESIS_LEFT=1, TK_PARENTESIS_RIGHT=2, TK_CORCHETE_LETF=3, TK_CORCHETE_RIGHT=4, 
		TK_LLAVE_LEFT=5, TK_LLAVE_RIGHT=6, TK_PUNTO_COMA=7, TK_DOSPUNTOS=8, TK_COMA=9, 
		TK_MENOR=10, TK_MAYOR=11, TK_PUNTO=12, TK_IGUAL=13, TK_MAS=14, TK_MENOS=15, 
		TK_POR=16, TK_DIVISION=17, TK_PORCENTAJE=18, TK_BARRA=19, TK_AMPERSAND=20, 
		TK_ADMIRACION=21, TK_MAYORIGULA=22, TK_MENOIGUAL=23, TK_IGUALIGUAL=24, 
		TK_DIFERENTE=25, TK_OR=26, TK_AND=27, TK_WHAT=28, TK_TIPORETURN=29, TK_IMPRESION=30, 
		TK_INT=31, TK_FLOAT=32, TK_BOOL=33, TK_CHAR=34, TK_STRING=35, TK_MAIN=36, 
		TK_USIZE=37, TK_LET=38, TK_MUT=39, TK_STRUCT=40, TK_AS=41, TK_TRUE=42, 
		TK_FALSE=43, TK_FN=44, TK_RETURN=45, TK_ABS=46, TK_SQRT=47, TK_TOSTRING=48, 
		TK_CLONE=49, TK_NEW=50, TK_LEN=51, TK_PUSH=52, TK_REMOVED=53, TK_CONTAINS=54, 
		TK_INSERT=55, TK_CAPACITY=56, TK_WITHCAPACITY=57, TK_IF=58, TK_ELSE=59, 
		TK_ELSE_IF=60, TK_WHILE=61, TK_LOOP=62, TK_BREAK=63, TK_POW_I64=64, TK_POW_F64=65, 
		TK_NUMBER=66, TK_DECIMAL=67, TK_CADENA=68, TK_CARACTER=69, TK_ID=70, TK_COMMET=71, 
		SPACES=72;
	public static final int
		RULE_start = 0, RULE_funcionmain = 1, RULE_instrucciones = 2, RULE_instruccion = 3, 
		RULE_impresion = 4, RULE_listprint = 5, RULE_expimprimir = 6, RULE_asignacionVariable = 7, 
		RULE_expresionIf = 8, RULE_listaelif = 9, RULE_elif = 10, RULE_expresionWhile = 11, 
		RULE_expresionLoop = 12, RULE_breakInst = 13, RULE_variable = 14, RULE_tipo = 15, 
		RULE_expresion = 16, RULE_valor = 17;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "funcionmain", "instrucciones", "instruccion", "impresion", 
			"listprint", "expimprimir", "asignacionVariable", "expresionIf", "listaelif", 
			"elif", "expresionWhile", "expresionLoop", "breakInst", "variable", "tipo", 
			"expresion", "valor"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'['", "']'", "'{'", "'}'", "';'", "':'", "','", 
			"'<'", "'>'", "'.'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'|'", 
			"'&'", "'!'", "'>='", "'<='", "'=='", "'!='", "'||'", "'&&'", "'?'", 
			"'->'", "'println!'", "'i64'", "'f64'", "'bool'", "'char'", "'String'", 
			"'main'", "'usize'", "'let'", "'mut'", "'struct'", "'as'", "'true'", 
			"'false'", "'fn'", "'return'", "'abs'", "'sqrt'", "'to_string'", "'clone'", 
			"'new'", "'len'", "'push'", "'remove'", "'contains'", "'insert'", "'capacity'", 
			"'with_capacity'", "'if'", "'else'", "'else if'", "'while'", "'loop'", 
			"'break'", "'i64::pow'", "'f64::powf'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TK_PARENTESIS_LEFT", "TK_PARENTESIS_RIGHT", "TK_CORCHETE_LETF", 
			"TK_CORCHETE_RIGHT", "TK_LLAVE_LEFT", "TK_LLAVE_RIGHT", "TK_PUNTO_COMA", 
			"TK_DOSPUNTOS", "TK_COMA", "TK_MENOR", "TK_MAYOR", "TK_PUNTO", "TK_IGUAL", 
			"TK_MAS", "TK_MENOS", "TK_POR", "TK_DIVISION", "TK_PORCENTAJE", "TK_BARRA", 
			"TK_AMPERSAND", "TK_ADMIRACION", "TK_MAYORIGULA", "TK_MENOIGUAL", "TK_IGUALIGUAL", 
			"TK_DIFERENTE", "TK_OR", "TK_AND", "TK_WHAT", "TK_TIPORETURN", "TK_IMPRESION", 
			"TK_INT", "TK_FLOAT", "TK_BOOL", "TK_CHAR", "TK_STRING", "TK_MAIN", "TK_USIZE", 
			"TK_LET", "TK_MUT", "TK_STRUCT", "TK_AS", "TK_TRUE", "TK_FALSE", "TK_FN", 
			"TK_RETURN", "TK_ABS", "TK_SQRT", "TK_TOSTRING", "TK_CLONE", "TK_NEW", 
			"TK_LEN", "TK_PUSH", "TK_REMOVED", "TK_CONTAINS", "TK_INSERT", "TK_CAPACITY", 
			"TK_WITHCAPACITY", "TK_IF", "TK_ELSE", "TK_ELSE_IF", "TK_WHILE", "TK_LOOP", 
			"TK_BREAK", "TK_POW_I64", "TK_POW_F64", "TK_NUMBER", "TK_DECIMAL", "TK_CADENA", 
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
		public *arrayList.List lista;
		public FuncionmainContext funcionmain;
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
			setState(36);
			((StartContext)_localctx).funcionmain = funcionmain();
			setState(37);
			match(EOF);
			 _localctx.lista = ((StartContext)_localctx).funcionmain.lista 
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
		public *arrayList.List lista;
		public InstruccionesContext instrucciones;
		public TerminalNode TK_FN() { return getToken(rustParser.TK_FN, 0); }
		public TerminalNode TK_MAIN() { return getToken(rustParser.TK_MAIN, 0); }
		public TerminalNode TK_PARENTESIS_LEFT() { return getToken(rustParser.TK_PARENTESIS_LEFT, 0); }
		public TerminalNode TK_PARENTESIS_RIGHT() { return getToken(rustParser.TK_PARENTESIS_RIGHT, 0); }
		public TerminalNode TK_LLAVE_LEFT() { return getToken(rustParser.TK_LLAVE_LEFT, 0); }
		public TerminalNode TK_LLAVE_RIGHT() { return getToken(rustParser.TK_LLAVE_RIGHT, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public FuncionmainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcionmain; }
	}

	public final FuncionmainContext funcionmain() throws RecognitionException {
		FuncionmainContext _localctx = new FuncionmainContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_funcionmain);
		try {
			setState(55);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(40);
				match(TK_FN);
				setState(41);
				match(TK_MAIN);
				setState(42);
				match(TK_PARENTESIS_LEFT);
				setState(43);
				match(TK_PARENTESIS_RIGHT);
				setState(44);
				match(TK_LLAVE_LEFT);
				setState(45);
				match(TK_LLAVE_RIGHT);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(46);
				match(TK_FN);
				setState(47);
				match(TK_MAIN);
				setState(48);
				match(TK_PARENTESIS_LEFT);
				setState(49);
				match(TK_PARENTESIS_RIGHT);
				setState(50);
				match(TK_LLAVE_LEFT);
				setState(51);
				((FuncionmainContext)_localctx).instrucciones = instrucciones();
				setState(52);
				match(TK_LLAVE_RIGHT);
				 _localctx.lista = ((FuncionmainContext)_localctx).instrucciones.lista 
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
		public *arrayList.List lista;
		public InstruccionContext instruccion;
		public List<InstruccionContext> e = new ArrayList<InstruccionContext>();
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instrucciones);
		_localctx.lista = arrayList.New()
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(60);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 30)) & ~0x3f) == 0 && ((1L << (_la - 30)) & ((1L << (TK_IMPRESION - 30)) | (1L << (TK_LET - 30)) | (1L << (TK_IF - 30)) | (1L << (TK_WHILE - 30)) | (1L << (TK_LOOP - 30)) | (1L << (TK_BREAK - 30)) | (1L << (TK_ID - 30)))) != 0)) {
				{
				{
				setState(57);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(62);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

			            listInt := localctx.(*InstruccionesContext).GetE()
			            for _,e:= range listInt{
			                _localctx.lista.Add(e.GetInst())
			            }
			        
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

	public static class InstruccionContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public VariableContext variable;
		public ImpresionContext impresion;
		public AsignacionVariableContext asignacionVariable;
		public ExpresionIfContext expresionIf;
		public ExpresionWhileContext expresionWhile;
		public ExpresionLoopContext expresionLoop;
		public BreakInstContext breakInst;
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public ImpresionContext impresion() {
			return getRuleContext(ImpresionContext.class,0);
		}
		public AsignacionVariableContext asignacionVariable() {
			return getRuleContext(AsignacionVariableContext.class,0);
		}
		public ExpresionIfContext expresionIf() {
			return getRuleContext(ExpresionIfContext.class,0);
		}
		public ExpresionWhileContext expresionWhile() {
			return getRuleContext(ExpresionWhileContext.class,0);
		}
		public ExpresionLoopContext expresionLoop() {
			return getRuleContext(ExpresionLoopContext.class,0);
		}
		public BreakInstContext breakInst() {
			return getRuleContext(BreakInstContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_instruccion);
		try {
			setState(86);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_LET:
				enterOuterAlt(_localctx, 1);
				{
				setState(65);
				((InstruccionContext)_localctx).variable = variable();
				 _localctx.inst = ((InstruccionContext)_localctx).variable.inst 
				}
				break;
			case TK_IMPRESION:
				enterOuterAlt(_localctx, 2);
				{
				setState(68);
				((InstruccionContext)_localctx).impresion = impresion();
				 _localctx.inst = ((InstruccionContext)_localctx).impresion.inst 
				}
				break;
			case TK_ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(71);
				((InstruccionContext)_localctx).asignacionVariable = asignacionVariable();
				 _localctx.inst = ((InstruccionContext)_localctx).asignacionVariable.inst 
				}
				break;
			case TK_IF:
				enterOuterAlt(_localctx, 4);
				{
				setState(74);
				((InstruccionContext)_localctx).expresionIf = expresionIf();
				 _localctx.inst = ((InstruccionContext)_localctx).expresionIf.inst 
				}
				break;
			case TK_WHILE:
				enterOuterAlt(_localctx, 5);
				{
				setState(77);
				((InstruccionContext)_localctx).expresionWhile = expresionWhile();
				 _localctx.inst = ((InstruccionContext)_localctx).expresionWhile.inst 
				}
				break;
			case TK_LOOP:
				enterOuterAlt(_localctx, 6);
				{
				setState(80);
				((InstruccionContext)_localctx).expresionLoop = expresionLoop();
				 _localctx.inst = ((InstruccionContext)_localctx).expresionLoop.inst  
				}
				break;
			case TK_BREAK:
				enterOuterAlt(_localctx, 7);
				{
				setState(83);
				((InstruccionContext)_localctx).breakInst = breakInst();
				 _localctx.inst = ((InstruccionContext)_localctx).breakInst.inst 
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
		public interfaces.Instruction inst;
		public ExpresionContext expresion;
		public ListprintContext listprint;
		public TerminalNode TK_IMPRESION() { return getToken(rustParser.TK_IMPRESION, 0); }
		public TerminalNode TK_PARENTESIS_LEFT() { return getToken(rustParser.TK_PARENTESIS_LEFT, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode TK_PARENTESIS_RIGHT() { return getToken(rustParser.TK_PARENTESIS_RIGHT, 0); }
		public TerminalNode TK_PUNTO_COMA() { return getToken(rustParser.TK_PUNTO_COMA, 0); }
		public ListprintContext listprint() {
			return getRuleContext(ListprintContext.class,0);
		}
		public ImpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_impresion; }
	}

	public final ImpresionContext impresion() throws RecognitionException {
		ImpresionContext _localctx = new ImpresionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_impresion);
		try {
			setState(103);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(88);
				match(TK_IMPRESION);
				setState(89);
				match(TK_PARENTESIS_LEFT);
				setState(90);
				((ImpresionContext)_localctx).expresion = expresion(0);
				setState(91);
				match(TK_PARENTESIS_RIGHT);
				setState(92);
				match(TK_PUNTO_COMA);
				_localctx.inst = instrucciones.NewImprimir(((ImpresionContext)_localctx).expresion.primate,nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(95);
				match(TK_IMPRESION);
				setState(96);
				match(TK_PARENTESIS_LEFT);
				setState(97);
				((ImpresionContext)_localctx).expresion = expresion(0);
				setState(98);
				((ImpresionContext)_localctx).listprint = listprint();
				setState(99);
				match(TK_PARENTESIS_RIGHT);
				setState(100);
				match(TK_PUNTO_COMA);
				_localctx.inst = instrucciones.NewImprimir(((ImpresionContext)_localctx).expresion.primate,((ImpresionContext)_localctx).listprint.lista)
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

	public static class ListprintContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ExpimprimirContext expimprimir;
		public List<ExpimprimirContext> list = new ArrayList<ExpimprimirContext>();
		public List<ExpimprimirContext> expimprimir() {
			return getRuleContexts(ExpimprimirContext.class);
		}
		public ExpimprimirContext expimprimir(int i) {
			return getRuleContext(ExpimprimirContext.class,i);
		}
		public ListprintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listprint; }
	}

	public final ListprintContext listprint() throws RecognitionException {
		ListprintContext _localctx = new ListprintContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_listprint);
		_localctx.lista = arrayList.New()
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(106); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(105);
				((ListprintContext)_localctx).expimprimir = expimprimir();
				((ListprintContext)_localctx).list.add(((ListprintContext)_localctx).expimprimir);
				}
				}
				setState(108); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==TK_COMA );

			            localList := localctx.(*ListprintContext).GetList()
			            for _, list := range localList{
			                _localctx.lista.Add(list.GetPrimate())
			            }
			        
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

	public static class ExpimprimirContext extends ParserRuleContext {
		public interfaces.Expresion primate;
		public ExpresionContext expresion;
		public TerminalNode TK_COMA() { return getToken(rustParser.TK_COMA, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ExpimprimirContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expimprimir; }
	}

	public final ExpimprimirContext expimprimir() throws RecognitionException {
		ExpimprimirContext _localctx = new ExpimprimirContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_expimprimir);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(112);
			match(TK_COMA);
			setState(113);
			((ExpimprimirContext)_localctx).expresion = expresion(0);
			_localctx.primate = ((ExpimprimirContext)_localctx).expresion.primate
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

	public static class AsignacionVariableContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public Token TK_ID;
		public ExpresionContext expresion;
		public TerminalNode TK_ID() { return getToken(rustParser.TK_ID, 0); }
		public TerminalNode TK_IGUAL() { return getToken(rustParser.TK_IGUAL, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode TK_PUNTO_COMA() { return getToken(rustParser.TK_PUNTO_COMA, 0); }
		public AsignacionVariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacionVariable; }
	}

	public final AsignacionVariableContext asignacionVariable() throws RecognitionException {
		AsignacionVariableContext _localctx = new AsignacionVariableContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_asignacionVariable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(116);
			((AsignacionVariableContext)_localctx).TK_ID = match(TK_ID);
			setState(117);
			match(TK_IGUAL);
			setState(118);
			((AsignacionVariableContext)_localctx).expresion = expresion(0);
			setState(119);
			match(TK_PUNTO_COMA);
			 _localctx.inst = instrucciones.NewAsignacion((((AsignacionVariableContext)_localctx).TK_ID!=null?((AsignacionVariableContext)_localctx).TK_ID.getText():null),((AsignacionVariableContext)_localctx).expresion.primate) 
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
		public interfaces.Instruction inst;
		public ExpresionContext expresion;
		public InstruccionesContext bloqueif;
		public InstruccionesContext bloqueelse;
		public InstruccionesContext ifblock;
		public ListaelifContext listaelif;
		public TerminalNode TK_IF() { return getToken(rustParser.TK_IF, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public List<TerminalNode> TK_LLAVE_LEFT() { return getTokens(rustParser.TK_LLAVE_LEFT); }
		public TerminalNode TK_LLAVE_LEFT(int i) {
			return getToken(rustParser.TK_LLAVE_LEFT, i);
		}
		public List<TerminalNode> TK_LLAVE_RIGHT() { return getTokens(rustParser.TK_LLAVE_RIGHT); }
		public TerminalNode TK_LLAVE_RIGHT(int i) {
			return getToken(rustParser.TK_LLAVE_RIGHT, i);
		}
		public List<InstruccionesContext> instrucciones() {
			return getRuleContexts(InstruccionesContext.class);
		}
		public InstruccionesContext instrucciones(int i) {
			return getRuleContext(InstruccionesContext.class,i);
		}
		public TerminalNode TK_ELSE() { return getToken(rustParser.TK_ELSE, 0); }
		public ListaelifContext listaelif() {
			return getRuleContext(ListaelifContext.class,0);
		}
		public ExpresionIfContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresionIf; }
	}

	public final ExpresionIfContext expresionIf() throws RecognitionException {
		ExpresionIfContext _localctx = new ExpresionIfContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_expresionIf);
		try {
			setState(160);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(122);
				match(TK_IF);
				setState(123);
				((ExpresionIfContext)_localctx).expresion = expresion(0);
				setState(124);
				match(TK_LLAVE_LEFT);
				setState(125);
				((ExpresionIfContext)_localctx).bloqueif = instrucciones();
				setState(126);
				match(TK_LLAVE_RIGHT);
				 _localctx.inst = instrucciones.NewIf(((ExpresionIfContext)_localctx).expresion.primate,((ExpresionIfContext)_localctx).bloqueif.lista,nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(129);
				match(TK_IF);
				setState(130);
				((ExpresionIfContext)_localctx).expresion = expresion(0);
				setState(131);
				match(TK_LLAVE_LEFT);
				setState(132);
				((ExpresionIfContext)_localctx).bloqueif = instrucciones();
				setState(133);
				match(TK_LLAVE_RIGHT);
				setState(134);
				match(TK_ELSE);
				setState(135);
				match(TK_LLAVE_LEFT);
				setState(136);
				((ExpresionIfContext)_localctx).bloqueelse = instrucciones();
				setState(137);
				match(TK_LLAVE_RIGHT);
				 _localctx.inst = instrucciones.NewIf(((ExpresionIfContext)_localctx).expresion.primate,((ExpresionIfContext)_localctx).bloqueif.lista,((ExpresionIfContext)_localctx).bloqueelse.lista) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(140);
				match(TK_IF);
				setState(141);
				((ExpresionIfContext)_localctx).expresion = expresion(0);
				setState(142);
				match(TK_LLAVE_LEFT);
				setState(143);
				((ExpresionIfContext)_localctx).ifblock = instrucciones();
				setState(144);
				match(TK_LLAVE_RIGHT);
				setState(145);
				((ExpresionIfContext)_localctx).listaelif = listaelif();
				 _localctx.inst = instrucciones.NewIf2(((ExpresionIfContext)_localctx).expresion.primate,((ExpresionIfContext)_localctx).ifblock.lista,nil,((ExpresionIfContext)_localctx).listaelif.lista)  
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(148);
				match(TK_IF);
				setState(149);
				((ExpresionIfContext)_localctx).expresion = expresion(0);
				setState(150);
				match(TK_LLAVE_LEFT);
				setState(151);
				((ExpresionIfContext)_localctx).ifblock = instrucciones();
				setState(152);
				match(TK_LLAVE_RIGHT);
				setState(153);
				((ExpresionIfContext)_localctx).listaelif = listaelif();
				setState(154);
				match(TK_ELSE);
				setState(155);
				match(TK_LLAVE_LEFT);
				setState(156);
				((ExpresionIfContext)_localctx).bloqueelse = instrucciones();
				setState(157);
				match(TK_LLAVE_RIGHT);
				 _localctx.inst = instrucciones.NewIf2(((ExpresionIfContext)_localctx).expresion.primate,((ExpresionIfContext)_localctx).ifblock.lista,((ExpresionIfContext)_localctx).bloqueelse.lista,((ExpresionIfContext)_localctx).listaelif.lista)  
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

	public static class ListaelifContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ElifContext elif;
		public List<ElifContext> l = new ArrayList<ElifContext>();
		public List<ElifContext> elif() {
			return getRuleContexts(ElifContext.class);
		}
		public ElifContext elif(int i) {
			return getRuleContext(ElifContext.class,i);
		}
		public ListaelifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaelif; }
	}

	public final ListaelifContext listaelif() throws RecognitionException {
		ListaelifContext _localctx = new ListaelifContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_listaelif);
		 _localctx.lista = arrayList.New() 
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(163); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(162);
				((ListaelifContext)_localctx).elif = elif();
				((ListaelifContext)_localctx).l.add(((ListaelifContext)_localctx).elif);
				}
				}
				setState(165); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==TK_ELSE_IF );

			            listLocal := localctx.(*ListaelifContext).GetL()
			            for _,l := range listLocal{
			                _localctx.lista.Add(l.GetInst())
			            }
			        
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

	public static class ElifContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public ExpresionContext expresion;
		public InstruccionesContext elifblock;
		public TerminalNode TK_ELSE_IF() { return getToken(rustParser.TK_ELSE_IF, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode TK_LLAVE_LEFT() { return getToken(rustParser.TK_LLAVE_LEFT, 0); }
		public TerminalNode TK_LLAVE_RIGHT() { return getToken(rustParser.TK_LLAVE_RIGHT, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public ElifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elif; }
	}

	public final ElifContext elif() throws RecognitionException {
		ElifContext _localctx = new ElifContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_elif);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			match(TK_ELSE_IF);
			setState(170);
			((ElifContext)_localctx).expresion = expresion(0);
			setState(171);
			match(TK_LLAVE_LEFT);
			setState(172);
			((ElifContext)_localctx).elifblock = instrucciones();
			setState(173);
			match(TK_LLAVE_RIGHT);
			 _localctx.inst = instrucciones.NewIf(((ElifContext)_localctx).expresion.primate,((ElifContext)_localctx).elifblock.lista,nil)  
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

	public static class ExpresionWhileContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public ExpresionContext exp;
		public InstruccionesContext instrucciones;
		public TerminalNode TK_WHILE() { return getToken(rustParser.TK_WHILE, 0); }
		public TerminalNode TK_LLAVE_LEFT() { return getToken(rustParser.TK_LLAVE_LEFT, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_LLAVE_RIGHT() { return getToken(rustParser.TK_LLAVE_RIGHT, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ExpresionWhileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresionWhile; }
	}

	public final ExpresionWhileContext expresionWhile() throws RecognitionException {
		ExpresionWhileContext _localctx = new ExpresionWhileContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_expresionWhile);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(176);
			match(TK_WHILE);
			setState(177);
			((ExpresionWhileContext)_localctx).exp = expresion(0);
			setState(178);
			match(TK_LLAVE_LEFT);
			setState(179);
			((ExpresionWhileContext)_localctx).instrucciones = instrucciones();
			setState(180);
			match(TK_LLAVE_RIGHT);
			 _localctx.inst = instrucciones.NewWhile(((ExpresionWhileContext)_localctx).exp.primate,((ExpresionWhileContext)_localctx).instrucciones.lista) 
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

	public static class ExpresionLoopContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public InstruccionesContext instrucciones;
		public TerminalNode TK_LOOP() { return getToken(rustParser.TK_LOOP, 0); }
		public TerminalNode TK_LLAVE_LEFT() { return getToken(rustParser.TK_LLAVE_LEFT, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode TK_LLAVE_RIGHT() { return getToken(rustParser.TK_LLAVE_RIGHT, 0); }
		public ExpresionLoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresionLoop; }
	}

	public final ExpresionLoopContext expresionLoop() throws RecognitionException {
		ExpresionLoopContext _localctx = new ExpresionLoopContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_expresionLoop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			match(TK_LOOP);
			setState(184);
			match(TK_LLAVE_LEFT);
			setState(185);
			((ExpresionLoopContext)_localctx).instrucciones = instrucciones();
			setState(186);
			match(TK_LLAVE_RIGHT);
			 _localctx.inst = instrucciones.NewLoop(((ExpresionLoopContext)_localctx).instrucciones.lista) 
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

	public static class BreakInstContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public TerminalNode TK_BREAK() { return getToken(rustParser.TK_BREAK, 0); }
		public TerminalNode TK_PUNTO_COMA() { return getToken(rustParser.TK_PUNTO_COMA, 0); }
		public BreakInstContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakInst; }
	}

	public final BreakInstContext breakInst() throws RecognitionException {
		BreakInstContext _localctx = new BreakInstContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_breakInst);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(189);
			match(TK_BREAK);
			setState(190);
			match(TK_PUNTO_COMA);
			   _localctx.inst = instrucciones.NewBreak(interfaces.BREAK)    
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
		public interfaces.Instruction inst;
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
		enterRule(_localctx, 28, RULE_variable);
		try {
			setState(227);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(193);
				match(TK_LET);
				setState(194);
				match(TK_MUT);
				setState(195);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(196);
				match(TK_IGUAL);
				setState(197);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(198);
				match(TK_PUNTO_COMA);
				 _localctx.inst = instrucciones.NewDeclaracion((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),true,interfaces.NULL,((VariableContext)_localctx).expresion.primate)  
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(201);
				match(TK_LET);
				setState(202);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(203);
				match(TK_IGUAL);
				setState(204);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(205);
				match(TK_PUNTO_COMA);
				 _localctx.inst = instrucciones.NewDeclaracion((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),false,interfaces.NULL,((VariableContext)_localctx).expresion.primate) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(208);
				match(TK_LET);
				setState(209);
				match(TK_MUT);
				setState(210);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(211);
				match(TK_DOSPUNTOS);
				setState(212);
				((VariableContext)_localctx).tipo = tipo();
				setState(213);
				match(TK_IGUAL);
				setState(214);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(215);
				match(TK_PUNTO_COMA);
				 _localctx.inst = instrucciones.NewDeclaracion((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),true,((VariableContext)_localctx).tipo.tipoExp,((VariableContext)_localctx).expresion.primate)    
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(218);
				match(TK_LET);
				setState(219);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(220);
				match(TK_DOSPUNTOS);
				setState(221);
				((VariableContext)_localctx).tipo = tipo();
				setState(222);
				match(TK_IGUAL);
				setState(223);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(224);
				match(TK_PUNTO_COMA);
				 _localctx.inst = instrucciones.NewDeclaracion((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),false,((VariableContext)_localctx).tipo.tipoExp,((VariableContext)_localctx).expresion.primate)   
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
		enterRule(_localctx, 30, RULE_tipo);
		try {
			setState(239);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(229);
				match(TK_INT);
				_localctx.tipoExp = interfaces.INTEGER
				}
				break;
			case TK_FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(231);
				match(TK_FLOAT);
				_localctx.tipoExp = interfaces.FLOAT
				}
				break;
			case TK_BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(233);
				match(TK_BOOL);
				_localctx.tipoExp = interfaces.BOOLEAN
				}
				break;
			case TK_CHAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(235);
				match(TK_CHAR);
				_localctx.tipoExp = interfaces.CHAR
				}
				break;
			case TK_STRING:
				enterOuterAlt(_localctx, 5);
				{
				setState(237);
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
		public interfaces.Expresion primate;
		public ExpresionContext left;
		public Token op;
		public ExpresionContext right;
		public Token TK_POW_I64;
		public Token TK_POW_F64;
		public ValorContext valor;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode TK_ADMIRACION() { return getToken(rustParser.TK_ADMIRACION, 0); }
		public TerminalNode TK_MENOS() { return getToken(rustParser.TK_MENOS, 0); }
		public TerminalNode TK_POW_I64() { return getToken(rustParser.TK_POW_I64, 0); }
		public TerminalNode TK_PARENTESIS_LEFT() { return getToken(rustParser.TK_PARENTESIS_LEFT, 0); }
		public TerminalNode TK_COMA() { return getToken(rustParser.TK_COMA, 0); }
		public TerminalNode TK_PARENTESIS_RIGHT() { return getToken(rustParser.TK_PARENTESIS_RIGHT, 0); }
		public TerminalNode TK_POW_F64() { return getToken(rustParser.TK_POW_F64, 0); }
		public ValorContext valor() {
			return getRuleContext(ValorContext.class,0);
		}
		public TerminalNode TK_POR() { return getToken(rustParser.TK_POR, 0); }
		public TerminalNode TK_DIVISION() { return getToken(rustParser.TK_DIVISION, 0); }
		public TerminalNode TK_PORCENTAJE() { return getToken(rustParser.TK_PORCENTAJE, 0); }
		public TerminalNode TK_MAS() { return getToken(rustParser.TK_MAS, 0); }
		public TerminalNode TK_IGUALIGUAL() { return getToken(rustParser.TK_IGUALIGUAL, 0); }
		public TerminalNode TK_DIFERENTE() { return getToken(rustParser.TK_DIFERENTE, 0); }
		public TerminalNode TK_MAYOR() { return getToken(rustParser.TK_MAYOR, 0); }
		public TerminalNode TK_MENOR() { return getToken(rustParser.TK_MENOR, 0); }
		public TerminalNode TK_MAYORIGULA() { return getToken(rustParser.TK_MAYORIGULA, 0); }
		public TerminalNode TK_MENOIGUAL() { return getToken(rustParser.TK_MENOIGUAL, 0); }
		public TerminalNode TK_AND() { return getToken(rustParser.TK_AND, 0); }
		public TerminalNode TK_OR() { return getToken(rustParser.TK_OR, 0); }
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
		int _startState = 32;
		enterRecursionRule(_localctx, 32, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(270);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_MENOS:
			case TK_ADMIRACION:
				{
				setState(242);
				((ExpresionContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==TK_MENOS || _la==TK_ADMIRACION) ) {
					((ExpresionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(243);
				((ExpresionContext)_localctx).right = expresion(15);
				 _localctx.primate = expressiones.NewOperacion(nil,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
				}
				break;
			case TK_POW_I64:
				{
				setState(246);
				((ExpresionContext)_localctx).TK_POW_I64 = match(TK_POW_I64);
				setState(247);
				match(TK_PARENTESIS_LEFT);
				setState(248);
				((ExpresionContext)_localctx).left = expresion(0);
				setState(249);
				match(TK_COMA);
				setState(250);
				((ExpresionContext)_localctx).right = expresion(0);
				setState(251);
				match(TK_PARENTESIS_RIGHT);
				 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).TK_POW_I64!=null?((ExpresionContext)_localctx).TK_POW_I64.getText():null),((ExpresionContext)_localctx).right.primate) 
				}
				break;
			case TK_POW_F64:
				{
				setState(254);
				((ExpresionContext)_localctx).TK_POW_F64 = match(TK_POW_F64);
				setState(255);
				match(TK_PARENTESIS_LEFT);
				setState(256);
				((ExpresionContext)_localctx).left = expresion(0);
				setState(257);
				match(TK_COMA);
				setState(258);
				((ExpresionContext)_localctx).right = expresion(0);
				setState(259);
				match(TK_PARENTESIS_RIGHT);
				 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).TK_POW_F64!=null?((ExpresionContext)_localctx).TK_POW_F64.getText():null),((ExpresionContext)_localctx).right.primate) 
				}
				break;
			case TK_PARENTESIS_LEFT:
				{
				setState(262);
				match(TK_PARENTESIS_LEFT);
				setState(263);
				expresion(0);
				setState(264);
				match(TK_PARENTESIS_RIGHT);
				 _localctx.primate = ((ExpresionContext)_localctx).valor.primate 
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
				setState(267);
				((ExpresionContext)_localctx).valor = valor();
				 _localctx.primate = ((ExpresionContext)_localctx).valor.primate 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(324);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(322);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(272);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(273);
						((ExpresionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TK_POR) | (1L << TK_DIVISION) | (1L << TK_PORCENTAJE))) != 0)) ) {
							((ExpresionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(274);
						((ExpresionContext)_localctx).right = expresion(15);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(277);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(278);
						((ExpresionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==TK_MAS || _la==TK_MENOS) ) {
							((ExpresionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(279);
						((ExpresionContext)_localctx).right = expresion(14);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(282);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(283);
						((ExpresionContext)_localctx).op = match(TK_IGUALIGUAL);
						setState(284);
						((ExpresionContext)_localctx).right = expresion(11);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(287);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(288);
						((ExpresionContext)_localctx).op = match(TK_DIFERENTE);
						setState(289);
						((ExpresionContext)_localctx).right = expresion(10);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(292);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(293);
						((ExpresionContext)_localctx).op = match(TK_MAYOR);
						setState(294);
						((ExpresionContext)_localctx).right = expresion(9);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(297);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(298);
						((ExpresionContext)_localctx).op = match(TK_MENOR);
						setState(299);
						((ExpresionContext)_localctx).right = expresion(8);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(302);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(303);
						((ExpresionContext)_localctx).op = match(TK_MAYORIGULA);
						setState(304);
						((ExpresionContext)_localctx).right = expresion(7);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(307);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(308);
						((ExpresionContext)_localctx).op = match(TK_MENOIGUAL);
						setState(309);
						((ExpresionContext)_localctx).right = expresion(6);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(312);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(313);
						((ExpresionContext)_localctx).op = match(TK_AND);
						setState(314);
						((ExpresionContext)_localctx).right = expresion(5);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(317);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(318);
						((ExpresionContext)_localctx).op = match(TK_OR);
						setState(319);
						((ExpresionContext)_localctx).right = expresion(4);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					}
					} 
				}
				setState(326);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
		public Token TK_DECIMAL;
		public Token TK_TRUE;
		public Token TK_FALSE;
		public Token TK_CADENA;
		public Token TK_CARACTER;
		public Token TK_ID;
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
		enterRule(_localctx, 34, RULE_valor);
		try {
			setState(341);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(327);
				((ValorContext)_localctx).TK_NUMBER = match(TK_NUMBER);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextInt((((ValorContext)_localctx).TK_NUMBER!=null?((ValorContext)_localctx).TK_NUMBER.getText():null)),interfaces.INTEGER) 
				}
				break;
			case TK_DECIMAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(329);
				((ValorContext)_localctx).TK_DECIMAL = match(TK_DECIMAL);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextFloat((((ValorContext)_localctx).TK_DECIMAL!=null?((ValorContext)_localctx).TK_DECIMAL.getText():null)),interfaces.FLOAT) 
				}
				break;
			case TK_TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(331);
				((ValorContext)_localctx).TK_TRUE = match(TK_TRUE);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextBool((((ValorContext)_localctx).TK_TRUE!=null?((ValorContext)_localctx).TK_TRUE.getText():null)),interfaces.BOOLEAN) 
				}
				break;
			case TK_FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(333);
				((ValorContext)_localctx).TK_FALSE = match(TK_FALSE);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextBool((((ValorContext)_localctx).TK_FALSE!=null?((ValorContext)_localctx).TK_FALSE.getText():null)),interfaces.BOOLEAN) 
				}
				break;
			case TK_CADENA:
				enterOuterAlt(_localctx, 5);
				{
				setState(335);
				((ValorContext)_localctx).TK_CADENA = match(TK_CADENA);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextString((((ValorContext)_localctx).TK_CADENA!=null?((ValorContext)_localctx).TK_CADENA.getText():null)),interfaces.STRING) 
				}
				break;
			case TK_CARACTER:
				enterOuterAlt(_localctx, 6);
				{
				setState(337);
				((ValorContext)_localctx).TK_CARACTER = match(TK_CARACTER);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextString((((ValorContext)_localctx).TK_CARACTER!=null?((ValorContext)_localctx).TK_CARACTER.getText():null)),interfaces.CHAR) 
				}
				break;
			case TK_ID:
				enterOuterAlt(_localctx, 7);
				{
				setState(339);
				((ValorContext)_localctx).TK_ID = match(TK_ID);
				 _localctx.primate = expressiones.NewLLamadoVariable((((ValorContext)_localctx).TK_ID!=null?((ValorContext)_localctx).TK_ID.getText():null)) 
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
		case 16:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 14);
		case 1:
			return precpred(_ctx, 13);
		case 2:
			return precpred(_ctx, 10);
		case 3:
			return precpred(_ctx, 9);
		case 4:
			return precpred(_ctx, 8);
		case 5:
			return precpred(_ctx, 7);
		case 6:
			return precpred(_ctx, 6);
		case 7:
			return precpred(_ctx, 5);
		case 8:
			return precpred(_ctx, 4);
		case 9:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3J\u015a\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\5\3:\n\3\3\4\7\4=\n\4\f\4\16\4@\13\4\3\4\3\4\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\5\5Y\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\5\6j\n\6\3\7\6\7m\n\7\r\7\16\7n\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u00a3\n\n\3\13\6\13\u00a6\n\13\r\13\16"+
		"\13\u00a7\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\5\20\u00e6\n\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\5\21\u00f2\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0111\n\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\7\22\u0145\n\22\f\22\16\22\u0148\13\22\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u0158\n\23"+
		"\3\23\2\3\"\24\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$\2\5\4\2\21\21"+
		"\27\27\3\2\22\24\3\2\20\21\2\u0170\2&\3\2\2\2\49\3\2\2\2\6>\3\2\2\2\b"+
		"X\3\2\2\2\ni\3\2\2\2\fl\3\2\2\2\16r\3\2\2\2\20v\3\2\2\2\22\u00a2\3\2\2"+
		"\2\24\u00a5\3\2\2\2\26\u00ab\3\2\2\2\30\u00b2\3\2\2\2\32\u00b9\3\2\2\2"+
		"\34\u00bf\3\2\2\2\36\u00e5\3\2\2\2 \u00f1\3\2\2\2\"\u0110\3\2\2\2$\u0157"+
		"\3\2\2\2&\'\5\4\3\2\'(\7\2\2\3()\b\2\1\2)\3\3\2\2\2*+\7.\2\2+,\7&\2\2"+
		",-\7\3\2\2-.\7\4\2\2./\7\7\2\2/:\7\b\2\2\60\61\7.\2\2\61\62\7&\2\2\62"+
		"\63\7\3\2\2\63\64\7\4\2\2\64\65\7\7\2\2\65\66\5\6\4\2\66\67\7\b\2\2\67"+
		"8\b\3\1\28:\3\2\2\29*\3\2\2\29\60\3\2\2\2:\5\3\2\2\2;=\5\b\5\2<;\3\2\2"+
		"\2=@\3\2\2\2><\3\2\2\2>?\3\2\2\2?A\3\2\2\2@>\3\2\2\2AB\b\4\1\2B\7\3\2"+
		"\2\2CD\5\36\20\2DE\b\5\1\2EY\3\2\2\2FG\5\n\6\2GH\b\5\1\2HY\3\2\2\2IJ\5"+
		"\20\t\2JK\b\5\1\2KY\3\2\2\2LM\5\22\n\2MN\b\5\1\2NY\3\2\2\2OP\5\30\r\2"+
		"PQ\b\5\1\2QY\3\2\2\2RS\5\32\16\2ST\b\5\1\2TY\3\2\2\2UV\5\34\17\2VW\b\5"+
		"\1\2WY\3\2\2\2XC\3\2\2\2XF\3\2\2\2XI\3\2\2\2XL\3\2\2\2XO\3\2\2\2XR\3\2"+
		"\2\2XU\3\2\2\2Y\t\3\2\2\2Z[\7 \2\2[\\\7\3\2\2\\]\5\"\22\2]^\7\4\2\2^_"+
		"\7\t\2\2_`\b\6\1\2`j\3\2\2\2ab\7 \2\2bc\7\3\2\2cd\5\"\22\2de\5\f\7\2e"+
		"f\7\4\2\2fg\7\t\2\2gh\b\6\1\2hj\3\2\2\2iZ\3\2\2\2ia\3\2\2\2j\13\3\2\2"+
		"\2km\5\16\b\2lk\3\2\2\2mn\3\2\2\2nl\3\2\2\2no\3\2\2\2op\3\2\2\2pq\b\7"+
		"\1\2q\r\3\2\2\2rs\7\13\2\2st\5\"\22\2tu\b\b\1\2u\17\3\2\2\2vw\7H\2\2w"+
		"x\7\17\2\2xy\5\"\22\2yz\7\t\2\2z{\b\t\1\2{\21\3\2\2\2|}\7<\2\2}~\5\"\22"+
		"\2~\177\7\7\2\2\177\u0080\5\6\4\2\u0080\u0081\7\b\2\2\u0081\u0082\b\n"+
		"\1\2\u0082\u00a3\3\2\2\2\u0083\u0084\7<\2\2\u0084\u0085\5\"\22\2\u0085"+
		"\u0086\7\7\2\2\u0086\u0087\5\6\4\2\u0087\u0088\7\b\2\2\u0088\u0089\7="+
		"\2\2\u0089\u008a\7\7\2\2\u008a\u008b\5\6\4\2\u008b\u008c\7\b\2\2\u008c"+
		"\u008d\b\n\1\2\u008d\u00a3\3\2\2\2\u008e\u008f\7<\2\2\u008f\u0090\5\""+
		"\22\2\u0090\u0091\7\7\2\2\u0091\u0092\5\6\4\2\u0092\u0093\7\b\2\2\u0093"+
		"\u0094\5\24\13\2\u0094\u0095\b\n\1\2\u0095\u00a3\3\2\2\2\u0096\u0097\7"+
		"<\2\2\u0097\u0098\5\"\22\2\u0098\u0099\7\7\2\2\u0099\u009a\5\6\4\2\u009a"+
		"\u009b\7\b\2\2\u009b\u009c\5\24\13\2\u009c\u009d\7=\2\2\u009d\u009e\7"+
		"\7\2\2\u009e\u009f\5\6\4\2\u009f\u00a0\7\b\2\2\u00a0\u00a1\b\n\1\2\u00a1"+
		"\u00a3\3\2\2\2\u00a2|\3\2\2\2\u00a2\u0083\3\2\2\2\u00a2\u008e\3\2\2\2"+
		"\u00a2\u0096\3\2\2\2\u00a3\23\3\2\2\2\u00a4\u00a6\5\26\f\2\u00a5\u00a4"+
		"\3\2\2\2\u00a6\u00a7\3\2\2\2\u00a7\u00a5\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8"+
		"\u00a9\3\2\2\2\u00a9\u00aa\b\13\1\2\u00aa\25\3\2\2\2\u00ab\u00ac\7>\2"+
		"\2\u00ac\u00ad\5\"\22\2\u00ad\u00ae\7\7\2\2\u00ae\u00af\5\6\4\2\u00af"+
		"\u00b0\7\b\2\2\u00b0\u00b1\b\f\1\2\u00b1\27\3\2\2\2\u00b2\u00b3\7?\2\2"+
		"\u00b3\u00b4\5\"\22\2\u00b4\u00b5\7\7\2\2\u00b5\u00b6\5\6\4\2\u00b6\u00b7"+
		"\7\b\2\2\u00b7\u00b8\b\r\1\2\u00b8\31\3\2\2\2\u00b9\u00ba\7@\2\2\u00ba"+
		"\u00bb\7\7\2\2\u00bb\u00bc\5\6\4\2\u00bc\u00bd\7\b\2\2\u00bd\u00be\b\16"+
		"\1\2\u00be\33\3\2\2\2\u00bf\u00c0\7A\2\2\u00c0\u00c1\7\t\2\2\u00c1\u00c2"+
		"\b\17\1\2\u00c2\35\3\2\2\2\u00c3\u00c4\7(\2\2\u00c4\u00c5\7)\2\2\u00c5"+
		"\u00c6\7H\2\2\u00c6\u00c7\7\17\2\2\u00c7\u00c8\5\"\22\2\u00c8\u00c9\7"+
		"\t\2\2\u00c9\u00ca\b\20\1\2\u00ca\u00e6\3\2\2\2\u00cb\u00cc\7(\2\2\u00cc"+
		"\u00cd\7H\2\2\u00cd\u00ce\7\17\2\2\u00ce\u00cf\5\"\22\2\u00cf\u00d0\7"+
		"\t\2\2\u00d0\u00d1\b\20\1\2\u00d1\u00e6\3\2\2\2\u00d2\u00d3\7(\2\2\u00d3"+
		"\u00d4\7)\2\2\u00d4\u00d5\7H\2\2\u00d5\u00d6\7\n\2\2\u00d6\u00d7\5 \21"+
		"\2\u00d7\u00d8\7\17\2\2\u00d8\u00d9\5\"\22\2\u00d9\u00da\7\t\2\2\u00da"+
		"\u00db\b\20\1\2\u00db\u00e6\3\2\2\2\u00dc\u00dd\7(\2\2\u00dd\u00de\7H"+
		"\2\2\u00de\u00df\7\n\2\2\u00df\u00e0\5 \21\2\u00e0\u00e1\7\17\2\2\u00e1"+
		"\u00e2\5\"\22\2\u00e2\u00e3\7\t\2\2\u00e3\u00e4\b\20\1\2\u00e4\u00e6\3"+
		"\2\2\2\u00e5\u00c3\3\2\2\2\u00e5\u00cb\3\2\2\2\u00e5\u00d2\3\2\2\2\u00e5"+
		"\u00dc\3\2\2\2\u00e6\37\3\2\2\2\u00e7\u00e8\7!\2\2\u00e8\u00f2\b\21\1"+
		"\2\u00e9\u00ea\7\"\2\2\u00ea\u00f2\b\21\1\2\u00eb\u00ec\7#\2\2\u00ec\u00f2"+
		"\b\21\1\2\u00ed\u00ee\7$\2\2\u00ee\u00f2\b\21\1\2\u00ef\u00f0\7%\2\2\u00f0"+
		"\u00f2\b\21\1\2\u00f1\u00e7\3\2\2\2\u00f1\u00e9\3\2\2\2\u00f1\u00eb\3"+
		"\2\2\2\u00f1\u00ed\3\2\2\2\u00f1\u00ef\3\2\2\2\u00f2!\3\2\2\2\u00f3\u00f4"+
		"\b\22\1\2\u00f4\u00f5\t\2\2\2\u00f5\u00f6\5\"\22\21\u00f6\u00f7\b\22\1"+
		"\2\u00f7\u0111\3\2\2\2\u00f8\u00f9\7B\2\2\u00f9\u00fa\7\3\2\2\u00fa\u00fb"+
		"\5\"\22\2\u00fb\u00fc\7\13\2\2\u00fc\u00fd\5\"\22\2\u00fd\u00fe\7\4\2"+
		"\2\u00fe\u00ff\b\22\1\2\u00ff\u0111\3\2\2\2\u0100\u0101\7C\2\2\u0101\u0102"+
		"\7\3\2\2\u0102\u0103\5\"\22\2\u0103\u0104\7\13\2\2\u0104\u0105\5\"\22"+
		"\2\u0105\u0106\7\4\2\2\u0106\u0107\b\22\1\2\u0107\u0111\3\2\2\2\u0108"+
		"\u0109\7\3\2\2\u0109\u010a\5\"\22\2\u010a\u010b\7\4\2\2\u010b\u010c\b"+
		"\22\1\2\u010c\u0111\3\2\2\2\u010d\u010e\5$\23\2\u010e\u010f\b\22\1\2\u010f"+
		"\u0111\3\2\2\2\u0110\u00f3\3\2\2\2\u0110\u00f8\3\2\2\2\u0110\u0100\3\2"+
		"\2\2\u0110\u0108\3\2\2\2\u0110\u010d\3\2\2\2\u0111\u0146\3\2\2\2\u0112"+
		"\u0113\f\20\2\2\u0113\u0114\t\3\2\2\u0114\u0115\5\"\22\21\u0115\u0116"+
		"\b\22\1\2\u0116\u0145\3\2\2\2\u0117\u0118\f\17\2\2\u0118\u0119\t\4\2\2"+
		"\u0119\u011a\5\"\22\20\u011a\u011b\b\22\1\2\u011b\u0145\3\2\2\2\u011c"+
		"\u011d\f\f\2\2\u011d\u011e\7\32\2\2\u011e\u011f\5\"\22\r\u011f\u0120\b"+
		"\22\1\2\u0120\u0145\3\2\2\2\u0121\u0122\f\13\2\2\u0122\u0123\7\33\2\2"+
		"\u0123\u0124\5\"\22\f\u0124\u0125\b\22\1\2\u0125\u0145\3\2\2\2\u0126\u0127"+
		"\f\n\2\2\u0127\u0128\7\r\2\2\u0128\u0129\5\"\22\13\u0129\u012a\b\22\1"+
		"\2\u012a\u0145\3\2\2\2\u012b\u012c\f\t\2\2\u012c\u012d\7\f\2\2\u012d\u012e"+
		"\5\"\22\n\u012e\u012f\b\22\1\2\u012f\u0145\3\2\2\2\u0130\u0131\f\b\2\2"+
		"\u0131\u0132\7\30\2\2\u0132\u0133\5\"\22\t\u0133\u0134\b\22\1\2\u0134"+
		"\u0145\3\2\2\2\u0135\u0136\f\7\2\2\u0136\u0137\7\31\2\2\u0137\u0138\5"+
		"\"\22\b\u0138\u0139\b\22\1\2\u0139\u0145\3\2\2\2\u013a\u013b\f\6\2\2\u013b"+
		"\u013c\7\35\2\2\u013c\u013d\5\"\22\7\u013d\u013e\b\22\1\2\u013e\u0145"+
		"\3\2\2\2\u013f\u0140\f\5\2\2\u0140\u0141\7\34\2\2\u0141\u0142\5\"\22\6"+
		"\u0142\u0143\b\22\1\2\u0143\u0145\3\2\2\2\u0144\u0112\3\2\2\2\u0144\u0117"+
		"\3\2\2\2\u0144\u011c\3\2\2\2\u0144\u0121\3\2\2\2\u0144\u0126\3\2\2\2\u0144"+
		"\u012b\3\2\2\2\u0144\u0130\3\2\2\2\u0144\u0135\3\2\2\2\u0144\u013a\3\2"+
		"\2\2\u0144\u013f\3\2\2\2\u0145\u0148\3\2\2\2\u0146\u0144\3\2\2\2\u0146"+
		"\u0147\3\2\2\2\u0147#\3\2\2\2\u0148\u0146\3\2\2\2\u0149\u014a\7D\2\2\u014a"+
		"\u0158\b\23\1\2\u014b\u014c\7E\2\2\u014c\u0158\b\23\1\2\u014d\u014e\7"+
		",\2\2\u014e\u0158\b\23\1\2\u014f\u0150\7-\2\2\u0150\u0158\b\23\1\2\u0151"+
		"\u0152\7F\2\2\u0152\u0158\b\23\1\2\u0153\u0154\7G\2\2\u0154\u0158\b\23"+
		"\1\2\u0155\u0156\7H\2\2\u0156\u0158\b\23\1\2\u0157\u0149\3\2\2\2\u0157"+
		"\u014b\3\2\2\2\u0157\u014d\3\2\2\2\u0157\u014f\3\2\2\2\u0157\u0151\3\2"+
		"\2\2\u0157\u0153\3\2\2\2\u0157\u0155\3\2\2\2\u0158%\3\2\2\2\179>Xin\u00a2"+
		"\u00a7\u00e5\u00f1\u0110\u0144\u0146\u0157";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}