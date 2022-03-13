// Generated from c:\Users\joshu\Documents\USAC\1S2022\COMPI2\LAB\Proyecto\OLC2_Proyecto1_201314296\rust.g4 by ANTLR 4.8
 
    import "Proyecto1/src/interfaces"
    import "Proyecto1/src/expressiones"
    import "Proyecto1/src/instrucciones"
    import arrayList "github.com/colegno/arraylist"
    import "Proyecto1/src/pruebas" 
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
		TK_ELSE=60, TK_ELSE_IF=61, TK_WHILE=62, TK_NUMBER=63, TK_DECIMAL=64, TK_CADENA=65, 
		TK_CARACTER=66, TK_ID=67, TK_COMMET=68, SPACES=69;
	public static final int
		RULE_start = 0, RULE_funcionmain = 1, RULE_instrucciones = 2, RULE_instruccion = 3, 
		RULE_impresion = 4, RULE_asignacionVariable = 5, RULE_expresionIf = 6, 
		RULE_listaelif = 7, RULE_elif = 8, RULE_expresionWhile = 9, RULE_variable = 10, 
		RULE_tipo = 11, RULE_expresion = 12, RULE_valor = 13;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "funcionmain", "instrucciones", "instruccion", "impresion", 
			"asignacionVariable", "expresionIf", "listaelif", "elif", "expresionWhile", 
			"variable", "tipo", "expresion", "valor"
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
			"'with_capacity'", "'if'", "'else'", "'else if'", "'while'"
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
			"TK_WITHCAPACITY", "TK_IF", "TK_ELSE", "TK_ELSE_IF", "TK_WHILE", "TK_NUMBER", 
			"TK_DECIMAL", "TK_CADENA", "TK_CARACTER", "TK_ID", "TK_COMMET", "SPACES"
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
			setState(28);
			((StartContext)_localctx).funcionmain = funcionmain();
			setState(29);
			match(EOF);
			pruebas.Probar(((StartContext)_localctx).funcionmain.lista)
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
			setState(47);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
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
				setState(37);
				match(TK_LLAVE_RIGHT);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(38);
				match(TK_FN);
				setState(39);
				match(TK_MAIN);
				setState(40);
				match(TK_PARENTESIS_LEFT);
				setState(41);
				match(TK_PARENTESIS_RIGHT);
				setState(42);
				match(TK_LLAVE_LEFT);
				setState(43);
				((FuncionmainContext)_localctx).instrucciones = instrucciones();
				setState(44);
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
			setState(52);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 31)) & ~0x3f) == 0 && ((1L << (_la - 31)) & ((1L << (TK_IMPRESION - 31)) | (1L << (TK_LET - 31)) | (1L << (TK_IF - 31)) | (1L << (TK_WHILE - 31)) | (1L << (TK_ID - 31)))) != 0)) {
				{
				{
				setState(49);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(54);
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
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_instruccion);
		try {
			setState(72);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_LET:
				enterOuterAlt(_localctx, 1);
				{
				setState(57);
				((InstruccionContext)_localctx).variable = variable();
				 _localctx.inst = ((InstruccionContext)_localctx).variable.inst 
				}
				break;
			case TK_IMPRESION:
				enterOuterAlt(_localctx, 2);
				{
				setState(60);
				((InstruccionContext)_localctx).impresion = impresion();
				 _localctx.inst = ((InstruccionContext)_localctx).impresion.inst 
				}
				break;
			case TK_ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(63);
				((InstruccionContext)_localctx).asignacionVariable = asignacionVariable();
				 _localctx.inst = ((InstruccionContext)_localctx).asignacionVariable.inst 
				}
				break;
			case TK_IF:
				enterOuterAlt(_localctx, 4);
				{
				setState(66);
				((InstruccionContext)_localctx).expresionIf = expresionIf();
				 _localctx.inst = ((InstruccionContext)_localctx).expresionIf.inst 
				}
				break;
			case TK_WHILE:
				enterOuterAlt(_localctx, 5);
				{
				setState(69);
				((InstruccionContext)_localctx).expresionWhile = expresionWhile();
				 _localctx.inst = ((InstruccionContext)_localctx).expresionWhile.inst 
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
		enterRule(_localctx, 8, RULE_impresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(74);
			match(TK_IMPRESION);
			setState(75);
			match(TK_PARENTESIS_LEFT);
			setState(76);
			((ImpresionContext)_localctx).expresion = expresion(0);
			setState(77);
			match(TK_PARENTESIS_RIGHT);
			setState(78);
			match(TK_PUNTO_COMA);
			_localctx.inst = instrucciones.NewImprimir(((ImpresionContext)_localctx).expresion.primate)
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
		enterRule(_localctx, 10, RULE_asignacionVariable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(81);
			((AsignacionVariableContext)_localctx).TK_ID = match(TK_ID);
			setState(82);
			match(TK_IGUAL);
			setState(83);
			((AsignacionVariableContext)_localctx).expresion = expresion(0);
			setState(84);
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
		enterRule(_localctx, 12, RULE_expresionIf);
		try {
			setState(125);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(87);
				match(TK_IF);
				setState(88);
				((ExpresionIfContext)_localctx).expresion = expresion(0);
				setState(89);
				match(TK_LLAVE_LEFT);
				setState(90);
				((ExpresionIfContext)_localctx).bloqueif = instrucciones();
				setState(91);
				match(TK_LLAVE_RIGHT);
				 _localctx.inst = instrucciones.NewIf(((ExpresionIfContext)_localctx).expresion.primate,((ExpresionIfContext)_localctx).bloqueif.lista,nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(94);
				match(TK_IF);
				setState(95);
				((ExpresionIfContext)_localctx).expresion = expresion(0);
				setState(96);
				match(TK_LLAVE_LEFT);
				setState(97);
				((ExpresionIfContext)_localctx).bloqueif = instrucciones();
				setState(98);
				match(TK_LLAVE_RIGHT);
				setState(99);
				match(TK_ELSE);
				setState(100);
				match(TK_LLAVE_LEFT);
				setState(101);
				((ExpresionIfContext)_localctx).bloqueelse = instrucciones();
				setState(102);
				match(TK_LLAVE_RIGHT);
				 _localctx.inst = instrucciones.NewIf(((ExpresionIfContext)_localctx).expresion.primate,((ExpresionIfContext)_localctx).bloqueif.lista,((ExpresionIfContext)_localctx).bloqueelse.lista) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(105);
				match(TK_IF);
				setState(106);
				((ExpresionIfContext)_localctx).expresion = expresion(0);
				setState(107);
				match(TK_LLAVE_LEFT);
				setState(108);
				((ExpresionIfContext)_localctx).ifblock = instrucciones();
				setState(109);
				match(TK_LLAVE_RIGHT);
				setState(110);
				((ExpresionIfContext)_localctx).listaelif = listaelif();
				 _localctx.inst = instrucciones.NewIf2(((ExpresionIfContext)_localctx).expresion.primate,((ExpresionIfContext)_localctx).ifblock.lista,nil,((ExpresionIfContext)_localctx).listaelif.lista)  
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(113);
				match(TK_IF);
				setState(114);
				((ExpresionIfContext)_localctx).expresion = expresion(0);
				setState(115);
				match(TK_LLAVE_LEFT);
				setState(116);
				((ExpresionIfContext)_localctx).ifblock = instrucciones();
				setState(117);
				match(TK_LLAVE_RIGHT);
				setState(118);
				((ExpresionIfContext)_localctx).listaelif = listaelif();
				setState(119);
				match(TK_ELSE);
				setState(120);
				match(TK_LLAVE_LEFT);
				setState(121);
				((ExpresionIfContext)_localctx).bloqueelse = instrucciones();
				setState(122);
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
		enterRule(_localctx, 14, RULE_listaelif);
		 _localctx.lista = arrayList.New() 
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(127);
				((ListaelifContext)_localctx).elif = elif();
				((ListaelifContext)_localctx).l.add(((ListaelifContext)_localctx).elif);
				}
				}
				setState(130); 
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
		enterRule(_localctx, 16, RULE_elif);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(TK_ELSE_IF);
			setState(135);
			((ElifContext)_localctx).expresion = expresion(0);
			setState(136);
			match(TK_LLAVE_LEFT);
			setState(137);
			((ElifContext)_localctx).elifblock = instrucciones();
			setState(138);
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
		enterRule(_localctx, 18, RULE_expresionWhile);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			match(TK_WHILE);
			setState(142);
			((ExpresionWhileContext)_localctx).exp = expresion(0);
			setState(143);
			match(TK_LLAVE_LEFT);
			setState(144);
			((ExpresionWhileContext)_localctx).instrucciones = instrucciones();
			setState(145);
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
		enterRule(_localctx, 20, RULE_variable);
		try {
			setState(182);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(148);
				match(TK_LET);
				setState(149);
				match(TK_MUT);
				setState(150);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(151);
				match(TK_IGUAL);
				setState(152);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(153);
				match(TK_PUNTO_COMA);
				 _localctx.inst = instrucciones.NewDeclaracion((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),true,interfaces.NULL,((VariableContext)_localctx).expresion.primate)  
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(156);
				match(TK_LET);
				setState(157);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(158);
				match(TK_IGUAL);
				setState(159);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(160);
				match(TK_PUNTO_COMA);
				 _localctx.inst = instrucciones.NewDeclaracion((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),false,interfaces.NULL,((VariableContext)_localctx).expresion.primate) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(163);
				match(TK_LET);
				setState(164);
				match(TK_MUT);
				setState(165);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(166);
				match(TK_DOSPUNTOS);
				setState(167);
				((VariableContext)_localctx).tipo = tipo();
				setState(168);
				match(TK_IGUAL);
				setState(169);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(170);
				match(TK_PUNTO_COMA);
				 _localctx.inst = instrucciones.NewDeclaracion((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),true,((VariableContext)_localctx).tipo.tipoExp,((VariableContext)_localctx).expresion.primate)    
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(173);
				match(TK_LET);
				setState(174);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(175);
				match(TK_DOSPUNTOS);
				setState(176);
				((VariableContext)_localctx).tipo = tipo();
				setState(177);
				match(TK_IGUAL);
				setState(178);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(179);
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
		enterRule(_localctx, 22, RULE_tipo);
		try {
			setState(194);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(184);
				match(TK_INT);
				_localctx.tipoExp = interfaces.INTEGER
				}
				break;
			case TK_FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(186);
				match(TK_FLOAT);
				_localctx.tipoExp = interfaces.FLOAT
				}
				break;
			case TK_BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(188);
				match(TK_BOOL);
				_localctx.tipoExp = interfaces.BOOLEAN
				}
				break;
			case TK_CHAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(190);
				match(TK_CHAR);
				_localctx.tipoExp = interfaces.CHAR
				}
				break;
			case TK_STRING:
				enterOuterAlt(_localctx, 5);
				{
				setState(192);
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
		public ValorContext valor;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode TK_ADMIRACION() { return getToken(rustParser.TK_ADMIRACION, 0); }
		public TerminalNode TK_MENOS() { return getToken(rustParser.TK_MENOS, 0); }
		public TerminalNode TK_PARENTESIS_LEFT() { return getToken(rustParser.TK_PARENTESIS_LEFT, 0); }
		public TerminalNode TK_PARENTESIS_RIGHT() { return getToken(rustParser.TK_PARENTESIS_RIGHT, 0); }
		public ValorContext valor() {
			return getRuleContext(ValorContext.class,0);
		}
		public TerminalNode TK_POR() { return getToken(rustParser.TK_POR, 0); }
		public TerminalNode TK_DIVISION() { return getToken(rustParser.TK_DIVISION, 0); }
		public TerminalNode TK_PORCENTAJE() { return getToken(rustParser.TK_PORCENTAJE, 0); }
		public TerminalNode TK_MAS() { return getToken(rustParser.TK_MAS, 0); }
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
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(209);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_MENOS:
			case TK_ADMIRACION:
				{
				setState(197);
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
				setState(198);
				((ExpresionContext)_localctx).right = expresion(16);
				 _localctx.primate = expressiones.NewOperacion(nil,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
				}
				break;
			case TK_PARENTESIS_LEFT:
				{
				setState(201);
				match(TK_PARENTESIS_LEFT);
				setState(202);
				expresion(0);
				setState(203);
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
				setState(206);
				((ExpresionContext)_localctx).valor = valor();
				 _localctx.primate = ((ExpresionContext)_localctx).valor.primate 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(271);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(269);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(211);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(212);
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
						setState(213);
						((ExpresionContext)_localctx).right = expresion(16);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(216);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(217);
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
						setState(218);
						((ExpresionContext)_localctx).right = expresion(15);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(221);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(222);
						match(TK_AMPERSAND);
						setState(223);
						expresion(14);
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(224);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(225);
						match(T__0);
						setState(226);
						expresion(13);
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(227);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(228);
						((ExpresionContext)_localctx).op = match(TK_IGUALIGUAL);
						setState(229);
						((ExpresionContext)_localctx).right = expresion(11);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 6:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(232);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(233);
						((ExpresionContext)_localctx).op = match(TK_DIFERENTE);
						setState(234);
						((ExpresionContext)_localctx).right = expresion(10);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 7:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(237);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(238);
						((ExpresionContext)_localctx).op = match(TK_MAYOR);
						setState(239);
						((ExpresionContext)_localctx).right = expresion(9);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 8:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(242);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(243);
						((ExpresionContext)_localctx).op = match(TK_MENOR);
						setState(244);
						((ExpresionContext)_localctx).right = expresion(8);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 9:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(247);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(248);
						((ExpresionContext)_localctx).op = match(TK_MAYORIGULA);
						setState(249);
						((ExpresionContext)_localctx).right = expresion(7);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 10:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(252);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(253);
						((ExpresionContext)_localctx).op = match(TK_MENOIGUAL);
						setState(254);
						((ExpresionContext)_localctx).right = expresion(6);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 11:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(257);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(258);
						((ExpresionContext)_localctx).op = match(TK_AND);
						setState(259);
						((ExpresionContext)_localctx).right = expresion(5);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 12:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(262);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(263);
						((ExpresionContext)_localctx).op = match(TK_OR);
						setState(264);
						((ExpresionContext)_localctx).right = expresion(4);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(267);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(268);
						match(TK_BARRA);
						}
						break;
					}
					} 
				}
				setState(273);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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
		enterRule(_localctx, 26, RULE_valor);
		try {
			setState(288);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(274);
				((ValorContext)_localctx).TK_NUMBER = match(TK_NUMBER);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextInt((((ValorContext)_localctx).TK_NUMBER!=null?((ValorContext)_localctx).TK_NUMBER.getText():null)),interfaces.INTEGER) 
				}
				break;
			case TK_DECIMAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(276);
				((ValorContext)_localctx).TK_DECIMAL = match(TK_DECIMAL);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextFloat((((ValorContext)_localctx).TK_DECIMAL!=null?((ValorContext)_localctx).TK_DECIMAL.getText():null)),interfaces.FLOAT) 
				}
				break;
			case TK_TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(278);
				((ValorContext)_localctx).TK_TRUE = match(TK_TRUE);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextBool((((ValorContext)_localctx).TK_TRUE!=null?((ValorContext)_localctx).TK_TRUE.getText():null)),interfaces.BOOLEAN) 
				}
				break;
			case TK_FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(280);
				((ValorContext)_localctx).TK_FALSE = match(TK_FALSE);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextBool((((ValorContext)_localctx).TK_FALSE!=null?((ValorContext)_localctx).TK_FALSE.getText():null)),interfaces.BOOLEAN) 
				}
				break;
			case TK_CADENA:
				enterOuterAlt(_localctx, 5);
				{
				setState(282);
				((ValorContext)_localctx).TK_CADENA = match(TK_CADENA);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextString((((ValorContext)_localctx).TK_CADENA!=null?((ValorContext)_localctx).TK_CADENA.getText():null)),interfaces.STRING) 
				}
				break;
			case TK_CARACTER:
				enterOuterAlt(_localctx, 6);
				{
				setState(284);
				((ValorContext)_localctx).TK_CARACTER = match(TK_CARACTER);
				 _localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextString((((ValorContext)_localctx).TK_CARACTER!=null?((ValorContext)_localctx).TK_CARACTER.getText():null)),interfaces.CHAR) 
				}
				break;
			case TK_ID:
				enterOuterAlt(_localctx, 7);
				{
				setState(286);
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
		case 12:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3G\u0125\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\5\3\62\n\3\3\4\7\4\65\n"+
		"\4\f\4\16\48\13\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\5\5K\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\5\b\u0080\n\b\3\t\6\t\u0083\n\t\r\t\16\t\u0084\3"+
		"\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f"+
		"\u00b9\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00c5\n\r\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u00d4"+
		"\n\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\7\16\u0110\n\16\f\16\16\16\u0113\13\16\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u0123\n\17"+
		"\3\17\2\3\32\20\2\4\6\b\n\f\16\20\22\24\26\30\32\34\2\5\4\2\22\22\30\30"+
		"\3\2\23\25\3\2\21\22\2\u013c\2\36\3\2\2\2\4\61\3\2\2\2\6\66\3\2\2\2\b"+
		"J\3\2\2\2\nL\3\2\2\2\fS\3\2\2\2\16\177\3\2\2\2\20\u0082\3\2\2\2\22\u0088"+
		"\3\2\2\2\24\u008f\3\2\2\2\26\u00b8\3\2\2\2\30\u00c4\3\2\2\2\32\u00d3\3"+
		"\2\2\2\34\u0122\3\2\2\2\36\37\5\4\3\2\37 \7\2\2\3 !\b\2\1\2!\3\3\2\2\2"+
		"\"#\7/\2\2#$\7\'\2\2$%\7\4\2\2%&\7\5\2\2&\'\7\b\2\2\'\62\7\t\2\2()\7/"+
		"\2\2)*\7\'\2\2*+\7\4\2\2+,\7\5\2\2,-\7\b\2\2-.\5\6\4\2./\7\t\2\2/\60\b"+
		"\3\1\2\60\62\3\2\2\2\61\"\3\2\2\2\61(\3\2\2\2\62\5\3\2\2\2\63\65\5\b\5"+
		"\2\64\63\3\2\2\2\658\3\2\2\2\66\64\3\2\2\2\66\67\3\2\2\2\679\3\2\2\28"+
		"\66\3\2\2\29:\b\4\1\2:\7\3\2\2\2;<\5\26\f\2<=\b\5\1\2=K\3\2\2\2>?\5\n"+
		"\6\2?@\b\5\1\2@K\3\2\2\2AB\5\f\7\2BC\b\5\1\2CK\3\2\2\2DE\5\16\b\2EF\b"+
		"\5\1\2FK\3\2\2\2GH\5\24\13\2HI\b\5\1\2IK\3\2\2\2J;\3\2\2\2J>\3\2\2\2J"+
		"A\3\2\2\2JD\3\2\2\2JG\3\2\2\2K\t\3\2\2\2LM\7!\2\2MN\7\4\2\2NO\5\32\16"+
		"\2OP\7\5\2\2PQ\7\n\2\2QR\b\6\1\2R\13\3\2\2\2ST\7E\2\2TU\7\20\2\2UV\5\32"+
		"\16\2VW\7\n\2\2WX\b\7\1\2X\r\3\2\2\2YZ\7=\2\2Z[\5\32\16\2[\\\7\b\2\2\\"+
		"]\5\6\4\2]^\7\t\2\2^_\b\b\1\2_\u0080\3\2\2\2`a\7=\2\2ab\5\32\16\2bc\7"+
		"\b\2\2cd\5\6\4\2de\7\t\2\2ef\7>\2\2fg\7\b\2\2gh\5\6\4\2hi\7\t\2\2ij\b"+
		"\b\1\2j\u0080\3\2\2\2kl\7=\2\2lm\5\32\16\2mn\7\b\2\2no\5\6\4\2op\7\t\2"+
		"\2pq\5\20\t\2qr\b\b\1\2r\u0080\3\2\2\2st\7=\2\2tu\5\32\16\2uv\7\b\2\2"+
		"vw\5\6\4\2wx\7\t\2\2xy\5\20\t\2yz\7>\2\2z{\7\b\2\2{|\5\6\4\2|}\7\t\2\2"+
		"}~\b\b\1\2~\u0080\3\2\2\2\177Y\3\2\2\2\177`\3\2\2\2\177k\3\2\2\2\177s"+
		"\3\2\2\2\u0080\17\3\2\2\2\u0081\u0083\5\22\n\2\u0082\u0081\3\2\2\2\u0083"+
		"\u0084\3\2\2\2\u0084\u0082\3\2\2\2\u0084\u0085\3\2\2\2\u0085\u0086\3\2"+
		"\2\2\u0086\u0087\b\t\1\2\u0087\21\3\2\2\2\u0088\u0089\7?\2\2\u0089\u008a"+
		"\5\32\16\2\u008a\u008b\7\b\2\2\u008b\u008c\5\6\4\2\u008c\u008d\7\t\2\2"+
		"\u008d\u008e\b\n\1\2\u008e\23\3\2\2\2\u008f\u0090\7@\2\2\u0090\u0091\5"+
		"\32\16\2\u0091\u0092\7\b\2\2\u0092\u0093\5\6\4\2\u0093\u0094\7\t\2\2\u0094"+
		"\u0095\b\13\1\2\u0095\25\3\2\2\2\u0096\u0097\7)\2\2\u0097\u0098\7*\2\2"+
		"\u0098\u0099\7E\2\2\u0099\u009a\7\20\2\2\u009a\u009b\5\32\16\2\u009b\u009c"+
		"\7\n\2\2\u009c\u009d\b\f\1\2\u009d\u00b9\3\2\2\2\u009e\u009f\7)\2\2\u009f"+
		"\u00a0\7E\2\2\u00a0\u00a1\7\20\2\2\u00a1\u00a2\5\32\16\2\u00a2\u00a3\7"+
		"\n\2\2\u00a3\u00a4\b\f\1\2\u00a4\u00b9\3\2\2\2\u00a5\u00a6\7)\2\2\u00a6"+
		"\u00a7\7*\2\2\u00a7\u00a8\7E\2\2\u00a8\u00a9\7\13\2\2\u00a9\u00aa\5\30"+
		"\r\2\u00aa\u00ab\7\20\2\2\u00ab\u00ac\5\32\16\2\u00ac\u00ad\7\n\2\2\u00ad"+
		"\u00ae\b\f\1\2\u00ae\u00b9\3\2\2\2\u00af\u00b0\7)\2\2\u00b0\u00b1\7E\2"+
		"\2\u00b1\u00b2\7\13\2\2\u00b2\u00b3\5\30\r\2\u00b3\u00b4\7\20\2\2\u00b4"+
		"\u00b5\5\32\16\2\u00b5\u00b6\7\n\2\2\u00b6\u00b7\b\f\1\2\u00b7\u00b9\3"+
		"\2\2\2\u00b8\u0096\3\2\2\2\u00b8\u009e\3\2\2\2\u00b8\u00a5\3\2\2\2\u00b8"+
		"\u00af\3\2\2\2\u00b9\27\3\2\2\2\u00ba\u00bb\7\"\2\2\u00bb\u00c5\b\r\1"+
		"\2\u00bc\u00bd\7#\2\2\u00bd\u00c5\b\r\1\2\u00be\u00bf\7$\2\2\u00bf\u00c5"+
		"\b\r\1\2\u00c0\u00c1\7%\2\2\u00c1\u00c5\b\r\1\2\u00c2\u00c3\7&\2\2\u00c3"+
		"\u00c5\b\r\1\2\u00c4\u00ba\3\2\2\2\u00c4\u00bc\3\2\2\2\u00c4\u00be\3\2"+
		"\2\2\u00c4\u00c0\3\2\2\2\u00c4\u00c2\3\2\2\2\u00c5\31\3\2\2\2\u00c6\u00c7"+
		"\b\16\1\2\u00c7\u00c8\t\2\2\2\u00c8\u00c9\5\32\16\22\u00c9\u00ca\b\16"+
		"\1\2\u00ca\u00d4\3\2\2\2\u00cb\u00cc\7\4\2\2\u00cc\u00cd\5\32\16\2\u00cd"+
		"\u00ce\7\5\2\2\u00ce\u00cf\b\16\1\2\u00cf\u00d4\3\2\2\2\u00d0\u00d1\5"+
		"\34\17\2\u00d1\u00d2\b\16\1\2\u00d2\u00d4\3\2\2\2\u00d3\u00c6\3\2\2\2"+
		"\u00d3\u00cb\3\2\2\2\u00d3\u00d0\3\2\2\2\u00d4\u0111\3\2\2\2\u00d5\u00d6"+
		"\f\21\2\2\u00d6\u00d7\t\3\2\2\u00d7\u00d8\5\32\16\22\u00d8\u00d9\b\16"+
		"\1\2\u00d9\u0110\3\2\2\2\u00da\u00db\f\20\2\2\u00db\u00dc\t\4\2\2\u00dc"+
		"\u00dd\5\32\16\21\u00dd\u00de\b\16\1\2\u00de\u0110\3\2\2\2\u00df\u00e0"+
		"\f\17\2\2\u00e0\u00e1\7\27\2\2\u00e1\u0110\5\32\16\20\u00e2\u00e3\f\16"+
		"\2\2\u00e3\u00e4\7\3\2\2\u00e4\u0110\5\32\16\17\u00e5\u00e6\f\f\2\2\u00e6"+
		"\u00e7\7\33\2\2\u00e7\u00e8\5\32\16\r\u00e8\u00e9\b\16\1\2\u00e9\u0110"+
		"\3\2\2\2\u00ea\u00eb\f\13\2\2\u00eb\u00ec\7\34\2\2\u00ec\u00ed\5\32\16"+
		"\f\u00ed\u00ee\b\16\1\2\u00ee\u0110\3\2\2\2\u00ef\u00f0\f\n\2\2\u00f0"+
		"\u00f1\7\16\2\2\u00f1\u00f2\5\32\16\13\u00f2\u00f3\b\16\1\2\u00f3\u0110"+
		"\3\2\2\2\u00f4\u00f5\f\t\2\2\u00f5\u00f6\7\r\2\2\u00f6\u00f7\5\32\16\n"+
		"\u00f7\u00f8\b\16\1\2\u00f8\u0110\3\2\2\2\u00f9\u00fa\f\b\2\2\u00fa\u00fb"+
		"\7\31\2\2\u00fb\u00fc\5\32\16\t\u00fc\u00fd\b\16\1\2\u00fd\u0110\3\2\2"+
		"\2\u00fe\u00ff\f\7\2\2\u00ff\u0100\7\32\2\2\u0100\u0101\5\32\16\b\u0101"+
		"\u0102\b\16\1\2\u0102\u0110\3\2\2\2\u0103\u0104\f\6\2\2\u0104\u0105\7"+
		"\36\2\2\u0105\u0106\5\32\16\7\u0106\u0107\b\16\1\2\u0107\u0110\3\2\2\2"+
		"\u0108\u0109\f\5\2\2\u0109\u010a\7\35\2\2\u010a\u010b\5\32\16\6\u010b"+
		"\u010c\b\16\1\2\u010c\u0110\3\2\2\2\u010d\u010e\f\r\2\2\u010e\u0110\7"+
		"\26\2\2\u010f\u00d5\3\2\2\2\u010f\u00da\3\2\2\2\u010f\u00df\3\2\2\2\u010f"+
		"\u00e2\3\2\2\2\u010f\u00e5\3\2\2\2\u010f\u00ea\3\2\2\2\u010f\u00ef\3\2"+
		"\2\2\u010f\u00f4\3\2\2\2\u010f\u00f9\3\2\2\2\u010f\u00fe\3\2\2\2\u010f"+
		"\u0103\3\2\2\2\u010f\u0108\3\2\2\2\u010f\u010d\3\2\2\2\u0110\u0113\3\2"+
		"\2\2\u0111\u010f\3\2\2\2\u0111\u0112\3\2\2\2\u0112\33\3\2\2\2\u0113\u0111"+
		"\3\2\2\2\u0114\u0115\7A\2\2\u0115\u0123\b\17\1\2\u0116\u0117\7B\2\2\u0117"+
		"\u0123\b\17\1\2\u0118\u0119\7-\2\2\u0119\u0123\b\17\1\2\u011a\u011b\7"+
		".\2\2\u011b\u0123\b\17\1\2\u011c\u011d\7C\2\2\u011d\u0123\b\17\1\2\u011e"+
		"\u011f\7D\2\2\u011f\u0123\b\17\1\2\u0120\u0121\7E\2\2\u0121\u0123\b\17"+
		"\1\2\u0122\u0114\3\2\2\2\u0122\u0116\3\2\2\2\u0122\u0118\3\2\2\2\u0122"+
		"\u011a\3\2\2\2\u0122\u011c\3\2\2\2\u0122\u011e\3\2\2\2\u0122\u0120\3\2"+
		"\2\2\u0123\35\3\2\2\2\r\61\66J\177\u0084\u00b8\u00c4\u00d3\u010f\u0111"+
		"\u0122";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}