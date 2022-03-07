// Generated from c:\Users\joshu\Documents\USAC\1S2022\COMPI2\LAB\Proyecto\OLC2_Proyecto1_201314296\rust.g4 by ANTLR 4.8
 
    import "Proyecto1/src/interfaces"
    import "Proyecto1/src/expressiones"
    import "Proyecto1/src/instrucciones"
    import arrayList "github.com/colegno/arraylist"
    import "Proyecto1/src/pruebas"

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
		RULE_asignacionVariable = 4, RULE_expresionIf = 5, RULE_sintaxisIf = 6, 
		RULE_sintaxisElse = 7, RULE_variable = 8, RULE_tipo = 9, RULE_expresion = 10, 
		RULE_valor = 11;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "funcionmain", "instrucciones", "impresion", "asignacionVariable", 
			"expresionIf", "sintaxisIf", "sintaxisElse", "variable", "tipo", "expresion", 
			"valor"
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
			setState(24);
			((StartContext)_localctx).funcionmain = funcionmain();
			setState(25);
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
		public List<InstruccionesContext> e = new ArrayList<InstruccionesContext>();
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
		_localctx.lista = arrayList.New()
		int _la;
		try {
			setState(47);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(28);
				match(TK_FN);
				setState(29);
				match(TK_MAIN);
				setState(30);
				match(TK_PARENTESIS_LEFT);
				setState(31);
				match(TK_PARENTESIS_RIGHT);
				setState(32);
				match(TK_LLAVE_LEFT);
				setState(33);
				match(TK_LLAVE_RIGHT);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(34);
				match(TK_FN);
				setState(35);
				match(TK_MAIN);
				setState(36);
				match(TK_PARENTESIS_LEFT);
				setState(37);
				match(TK_PARENTESIS_RIGHT);
				setState(38);
				match(TK_LLAVE_LEFT);
				setState(42);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 31)) & ~0x3f) == 0 && ((1L << (_la - 31)) & ((1L << (TK_IMPRESION - 31)) | (1L << (TK_LET - 31)) | (1L << (TK_IF - 31)) | (1L << (TK_ID - 31)))) != 0)) {
					{
					{
					setState(39);
					((FuncionmainContext)_localctx).instrucciones = instrucciones();
					((FuncionmainContext)_localctx).e.add(((FuncionmainContext)_localctx).instrucciones);
					}
					}
					setState(44);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(45);
				match(TK_LLAVE_RIGHT);

				            listInt := localctx.(*FuncionmainContext).GetE()
				            for _,e:= range listInt{
				                _localctx.lista.Add(e.GetInst())
				            }
				        
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
		public interfaces.Instruction inst;
		public VariableContext variable;
		public ImpresionContext impresion;
		public AsignacionVariableContext asignacionVariable;
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
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instrucciones);
		try {
			setState(61);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_LET:
				enterOuterAlt(_localctx, 1);
				{
				setState(49);
				((InstruccionesContext)_localctx).variable = variable();
				 _localctx.inst = ((InstruccionesContext)_localctx).variable.inst 
				}
				break;
			case TK_IMPRESION:
				enterOuterAlt(_localctx, 2);
				{
				setState(52);
				((InstruccionesContext)_localctx).impresion = impresion();
				 _localctx.inst = ((InstruccionesContext)_localctx).impresion.inst 
				}
				break;
			case TK_ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(55);
				((InstruccionesContext)_localctx).asignacionVariable = asignacionVariable();
				 _localctx.inst = ((InstruccionesContext)_localctx).asignacionVariable.inst 
				}
				break;
			case TK_IF:
				enterOuterAlt(_localctx, 4);
				{
				setState(58);
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
		enterRule(_localctx, 6, RULE_impresion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(63);
			match(TK_IMPRESION);
			setState(64);
			match(TK_PARENTESIS_LEFT);
			setState(65);
			((ImpresionContext)_localctx).expresion = expresion(0);
			setState(66);
			match(TK_PARENTESIS_RIGHT);
			setState(67);
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
		enterRule(_localctx, 8, RULE_asignacionVariable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			((AsignacionVariableContext)_localctx).TK_ID = match(TK_ID);
			setState(71);
			match(TK_IGUAL);
			setState(72);
			((AsignacionVariableContext)_localctx).expresion = expresion(0);
			setState(73);
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
		enterRule(_localctx, 10, RULE_expresionIf);
		try {
			setState(83);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(76);
				sintaxisIf();

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(79);
				sintaxisIf();
				setState(80);
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
		enterRule(_localctx, 12, RULE_sintaxisIf);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(85);
			match(TK_IF);
			setState(86);
			expresion(0);
			setState(87);
			match(TK_LLAVE_LEFT);
			setState(88);
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
		enterRule(_localctx, 14, RULE_sintaxisElse);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(91);
			match(TK_ELSE);
			setState(92);
			match(TK_LLAVE_LEFT);
			setState(93);
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
		public interfaces.Instruction inst;
		public Token TK_ID;
		public TipoContext tipo;
		public ExpresionContext expresion;
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
		enterRule(_localctx, 16, RULE_variable);
		try {
			setState(130);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(96);
				match(TK_LET);
				setState(97);
				match(TK_MUT);
				setState(98);
				match(TK_ID);
				setState(99);
				match(TK_IGUAL);
				setState(100);
				expresion(0);
				setState(101);
				match(TK_PUNTO_COMA);

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(104);
				match(TK_LET);
				setState(105);
				match(TK_ID);
				setState(106);
				match(TK_IGUAL);
				setState(107);
				expresion(0);
				setState(108);
				match(TK_PUNTO_COMA);

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(111);
				match(TK_LET);
				setState(112);
				match(TK_MUT);
				setState(113);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(114);
				match(TK_DOSPUNTOS);
				setState(115);
				((VariableContext)_localctx).tipo = tipo();
				setState(116);
				match(TK_IGUAL);
				setState(117);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(118);
				match(TK_PUNTO_COMA);
				_localctx.inst = instrucciones.NewDeclaracion((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),false,((VariableContext)_localctx).tipo.tipoExp,((VariableContext)_localctx).expresion.primate)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(121);
				match(TK_LET);
				setState(122);
				((VariableContext)_localctx).TK_ID = match(TK_ID);
				setState(123);
				match(TK_DOSPUNTOS);
				setState(124);
				((VariableContext)_localctx).tipo = tipo();
				setState(125);
				match(TK_IGUAL);
				setState(126);
				((VariableContext)_localctx).expresion = expresion(0);
				setState(127);
				match(TK_PUNTO_COMA);
				_localctx.inst = instrucciones.NewDeclaracion((((VariableContext)_localctx).TK_ID!=null?((VariableContext)_localctx).TK_ID.getText():null),true,((VariableContext)_localctx).tipo.tipoExp,((VariableContext)_localctx).expresion.primate)
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
		enterRule(_localctx, 18, RULE_tipo);
		try {
			setState(142);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(132);
				match(TK_INT);
				_localctx.tipoExp = interfaces.INTEGER
				}
				break;
			case TK_FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(134);
				match(TK_FLOAT);
				_localctx.tipoExp = interfaces.FLOAT
				}
				break;
			case TK_BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(136);
				match(TK_BOOL);
				_localctx.tipoExp = interfaces.BOOLEAN
				}
				break;
			case TK_CHAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(138);
				match(TK_CHAR);
				_localctx.tipoExp = interfaces.CHAR
				}
				break;
			case TK_STRING:
				enterOuterAlt(_localctx, 5);
				{
				setState(140);
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
		public ValorContext valor;
		public Token op;
		public ExpresionContext right;
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
		int _startState = 20;
		enterRecursionRule(_localctx, 20, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_PARENTESIS_LEFT:
				{
				setState(145);
				match(TK_PARENTESIS_LEFT);
				setState(146);
				expresion(0);
				setState(147);
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
				setState(149);
				((ExpresionContext)_localctx).valor = valor();
				_localctx.primate = ((ExpresionContext)_localctx).valor.primate
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(214);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(212);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(154);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(155);
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
						setState(156);
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
						setState(159);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(160);
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
						setState(161);
						((ExpresionContext)_localctx).right = expresion(15);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(164);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(165);
						match(TK_AMPERSAND);
						setState(166);
						expresion(14);
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(167);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(168);
						match(T__0);
						setState(169);
						expresion(13);
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(170);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(171);
						((ExpresionContext)_localctx).op = match(TK_IGUALIGUAL);
						setState(172);
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
						setState(175);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(176);
						((ExpresionContext)_localctx).op = match(TK_DIFERENTE);
						setState(177);
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
						setState(180);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(181);
						((ExpresionContext)_localctx).op = match(TK_MAYOR);
						setState(182);
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
						setState(185);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(186);
						((ExpresionContext)_localctx).op = match(TK_MENOR);
						setState(187);
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
						setState(190);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(191);
						((ExpresionContext)_localctx).op = match(TK_MAYORIGULA);
						setState(192);
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
						setState(195);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(196);
						((ExpresionContext)_localctx).op = match(TK_MENOIGUAL);
						setState(197);
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
						setState(200);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(201);
						((ExpresionContext)_localctx).op = match(TK_AND);
						setState(202);
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
						setState(205);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(206);
						((ExpresionContext)_localctx).op = match(TK_OR);
						setState(207);
						((ExpresionContext)_localctx).right = expresion(4);
						 _localctx.primate = expressiones.NewOperacion(((ExpresionContext)_localctx).left.primate,(((ExpresionContext)_localctx).op!=null?((ExpresionContext)_localctx).op.getText():null),((ExpresionContext)_localctx).right.primate) 
						}
						break;
					case 13:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(210);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(211);
						match(TK_BARRA);
						}
						break;
					}
					} 
				}
				setState(216);
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
		enterRule(_localctx, 22, RULE_valor);
		try {
			setState(231);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TK_NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(217);
				((ValorContext)_localctx).TK_NUMBER = match(TK_NUMBER);
				_localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextInt((((ValorContext)_localctx).TK_NUMBER!=null?((ValorContext)_localctx).TK_NUMBER.getText():null)),interfaces.INTEGER)
				}
				break;
			case TK_DECIMAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(219);
				((ValorContext)_localctx).TK_DECIMAL = match(TK_DECIMAL);
				_localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextFloat((((ValorContext)_localctx).TK_DECIMAL!=null?((ValorContext)_localctx).TK_DECIMAL.getText():null)),interfaces.FLOAT)
				}
				break;
			case TK_TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(221);
				((ValorContext)_localctx).TK_TRUE = match(TK_TRUE);
				_localctx.primate = expressiones.NewPrimito((((ValorContext)_localctx).TK_TRUE!=null?((ValorContext)_localctx).TK_TRUE.getText():null),interfaces.BOOLEAN)
				}
				break;
			case TK_FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(223);
				((ValorContext)_localctx).TK_FALSE = match(TK_FALSE);
				_localctx.primate = expressiones.NewPrimito((((ValorContext)_localctx).TK_FALSE!=null?((ValorContext)_localctx).TK_FALSE.getText():null),interfaces.BOOLEAN)
				}
				break;
			case TK_CADENA:
				enterOuterAlt(_localctx, 5);
				{
				setState(225);
				((ValorContext)_localctx).TK_CADENA = match(TK_CADENA);
				_localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextString((((ValorContext)_localctx).TK_CADENA!=null?((ValorContext)_localctx).TK_CADENA.getText():null)),interfaces.STRING)
				}
				break;
			case TK_CARACTER:
				enterOuterAlt(_localctx, 6);
				{
				setState(227);
				((ValorContext)_localctx).TK_CARACTER = match(TK_CARACTER);
				_localctx.primate = expressiones.NewPrimito(interfaces.ConvertTextString((((ValorContext)_localctx).TK_CARACTER!=null?((ValorContext)_localctx).TK_CARACTER.getText():null)),interfaces.CHAR)
				}
				break;
			case TK_ID:
				enterOuterAlt(_localctx, 7);
				{
				setState(229);
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
		case 10:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3E\u00ec\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\3\7\3+\n\3\f\3\16\3.\13\3\3\3\3\3\5\3\62\n\3\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4@\n\4\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7V\n\7\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u0085\n\n\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u0091\n\13\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\5\f\u009b\n\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\7\f\u00d7\n\f\f\f\16\f\u00da"+
		"\13\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00ea"+
		"\n\r\3\r\2\3\26\16\2\4\6\b\n\f\16\20\22\24\26\30\2\4\3\2\23\25\3\2\21"+
		"\22\2\u0100\2\32\3\2\2\2\4\61\3\2\2\2\6?\3\2\2\2\bA\3\2\2\2\nH\3\2\2\2"+
		"\fU\3\2\2\2\16W\3\2\2\2\20]\3\2\2\2\22\u0084\3\2\2\2\24\u0090\3\2\2\2"+
		"\26\u009a\3\2\2\2\30\u00e9\3\2\2\2\32\33\5\4\3\2\33\34\7\2\2\3\34\35\b"+
		"\2\1\2\35\3\3\2\2\2\36\37\7/\2\2\37 \7\'\2\2 !\7\4\2\2!\"\7\5\2\2\"#\7"+
		"\b\2\2#\62\7\t\2\2$%\7/\2\2%&\7\'\2\2&\'\7\4\2\2\'(\7\5\2\2(,\7\b\2\2"+
		")+\5\6\4\2*)\3\2\2\2+.\3\2\2\2,*\3\2\2\2,-\3\2\2\2-/\3\2\2\2.,\3\2\2\2"+
		"/\60\7\t\2\2\60\62\b\3\1\2\61\36\3\2\2\2\61$\3\2\2\2\62\5\3\2\2\2\63\64"+
		"\5\22\n\2\64\65\b\4\1\2\65@\3\2\2\2\66\67\5\b\5\2\678\b\4\1\28@\3\2\2"+
		"\29:\5\n\6\2:;\b\4\1\2;@\3\2\2\2<=\5\f\7\2=>\b\4\1\2>@\3\2\2\2?\63\3\2"+
		"\2\2?\66\3\2\2\2?9\3\2\2\2?<\3\2\2\2@\7\3\2\2\2AB\7!\2\2BC\7\4\2\2CD\5"+
		"\26\f\2DE\7\5\2\2EF\7\n\2\2FG\b\5\1\2G\t\3\2\2\2HI\7C\2\2IJ\7\20\2\2J"+
		"K\5\26\f\2KL\7\n\2\2LM\b\6\1\2M\13\3\2\2\2NO\5\16\b\2OP\b\7\1\2PV\3\2"+
		"\2\2QR\5\16\b\2RS\5\20\t\2ST\b\7\1\2TV\3\2\2\2UN\3\2\2\2UQ\3\2\2\2V\r"+
		"\3\2\2\2WX\7=\2\2XY\5\26\f\2YZ\7\b\2\2Z[\7\t\2\2[\\\b\b\1\2\\\17\3\2\2"+
		"\2]^\7>\2\2^_\7\b\2\2_`\7\t\2\2`a\b\t\1\2a\21\3\2\2\2bc\7)\2\2cd\7*\2"+
		"\2de\7C\2\2ef\7\20\2\2fg\5\26\f\2gh\7\n\2\2hi\b\n\1\2i\u0085\3\2\2\2j"+
		"k\7)\2\2kl\7C\2\2lm\7\20\2\2mn\5\26\f\2no\7\n\2\2op\b\n\1\2p\u0085\3\2"+
		"\2\2qr\7)\2\2rs\7*\2\2st\7C\2\2tu\7\13\2\2uv\5\24\13\2vw\7\20\2\2wx\5"+
		"\26\f\2xy\7\n\2\2yz\b\n\1\2z\u0085\3\2\2\2{|\7)\2\2|}\7C\2\2}~\7\13\2"+
		"\2~\177\5\24\13\2\177\u0080\7\20\2\2\u0080\u0081\5\26\f\2\u0081\u0082"+
		"\7\n\2\2\u0082\u0083\b\n\1\2\u0083\u0085\3\2\2\2\u0084b\3\2\2\2\u0084"+
		"j\3\2\2\2\u0084q\3\2\2\2\u0084{\3\2\2\2\u0085\23\3\2\2\2\u0086\u0087\7"+
		"\"\2\2\u0087\u0091\b\13\1\2\u0088\u0089\7#\2\2\u0089\u0091\b\13\1\2\u008a"+
		"\u008b\7$\2\2\u008b\u0091\b\13\1\2\u008c\u008d\7%\2\2\u008d\u0091\b\13"+
		"\1\2\u008e\u008f\7&\2\2\u008f\u0091\b\13\1\2\u0090\u0086\3\2\2\2\u0090"+
		"\u0088\3\2\2\2\u0090\u008a\3\2\2\2\u0090\u008c\3\2\2\2\u0090\u008e\3\2"+
		"\2\2\u0091\25\3\2\2\2\u0092\u0093\b\f\1\2\u0093\u0094\7\4\2\2\u0094\u0095"+
		"\5\26\f\2\u0095\u0096\7\5\2\2\u0096\u009b\3\2\2\2\u0097\u0098\5\30\r\2"+
		"\u0098\u0099\b\f\1\2\u0099\u009b\3\2\2\2\u009a\u0092\3\2\2\2\u009a\u0097"+
		"\3\2\2\2\u009b\u00d8\3\2\2\2\u009c\u009d\f\21\2\2\u009d\u009e\t\2\2\2"+
		"\u009e\u009f\5\26\f\22\u009f\u00a0\b\f\1\2\u00a0\u00d7\3\2\2\2\u00a1\u00a2"+
		"\f\20\2\2\u00a2\u00a3\t\3\2\2\u00a3\u00a4\5\26\f\21\u00a4\u00a5\b\f\1"+
		"\2\u00a5\u00d7\3\2\2\2\u00a6\u00a7\f\17\2\2\u00a7\u00a8\7\27\2\2\u00a8"+
		"\u00d7\5\26\f\20\u00a9\u00aa\f\16\2\2\u00aa\u00ab\7\3\2\2\u00ab\u00d7"+
		"\5\26\f\17\u00ac\u00ad\f\f\2\2\u00ad\u00ae\7\33\2\2\u00ae\u00af\5\26\f"+
		"\r\u00af\u00b0\b\f\1\2\u00b0\u00d7\3\2\2\2\u00b1\u00b2\f\13\2\2\u00b2"+
		"\u00b3\7\34\2\2\u00b3\u00b4\5\26\f\f\u00b4\u00b5\b\f\1\2\u00b5\u00d7\3"+
		"\2\2\2\u00b6\u00b7\f\n\2\2\u00b7\u00b8\7\16\2\2\u00b8\u00b9\5\26\f\13"+
		"\u00b9\u00ba\b\f\1\2\u00ba\u00d7\3\2\2\2\u00bb\u00bc\f\t\2\2\u00bc\u00bd"+
		"\7\r\2\2\u00bd\u00be\5\26\f\n\u00be\u00bf\b\f\1\2\u00bf\u00d7\3\2\2\2"+
		"\u00c0\u00c1\f\b\2\2\u00c1\u00c2\7\31\2\2\u00c2\u00c3\5\26\f\t\u00c3\u00c4"+
		"\b\f\1\2\u00c4\u00d7\3\2\2\2\u00c5\u00c6\f\7\2\2\u00c6\u00c7\7\32\2\2"+
		"\u00c7\u00c8\5\26\f\b\u00c8\u00c9\b\f\1\2\u00c9\u00d7\3\2\2\2\u00ca\u00cb"+
		"\f\6\2\2\u00cb\u00cc\7\36\2\2\u00cc\u00cd\5\26\f\7\u00cd\u00ce\b\f\1\2"+
		"\u00ce\u00d7\3\2\2\2\u00cf\u00d0\f\5\2\2\u00d0\u00d1\7\35\2\2\u00d1\u00d2"+
		"\5\26\f\6\u00d2\u00d3\b\f\1\2\u00d3\u00d7\3\2\2\2\u00d4\u00d5\f\r\2\2"+
		"\u00d5\u00d7\7\26\2\2\u00d6\u009c\3\2\2\2\u00d6\u00a1\3\2\2\2\u00d6\u00a6"+
		"\3\2\2\2\u00d6\u00a9\3\2\2\2\u00d6\u00ac\3\2\2\2\u00d6\u00b1\3\2\2\2\u00d6"+
		"\u00b6\3\2\2\2\u00d6\u00bb\3\2\2\2\u00d6\u00c0\3\2\2\2\u00d6\u00c5\3\2"+
		"\2\2\u00d6\u00ca\3\2\2\2\u00d6\u00cf\3\2\2\2\u00d6\u00d4\3\2\2\2\u00d7"+
		"\u00da\3\2\2\2\u00d8\u00d6\3\2\2\2\u00d8\u00d9\3\2\2\2\u00d9\27\3\2\2"+
		"\2\u00da\u00d8\3\2\2\2\u00db\u00dc\7?\2\2\u00dc\u00ea\b\r\1\2\u00dd\u00de"+
		"\7@\2\2\u00de\u00ea\b\r\1\2\u00df\u00e0\7-\2\2\u00e0\u00ea\b\r\1\2\u00e1"+
		"\u00e2\7.\2\2\u00e2\u00ea\b\r\1\2\u00e3\u00e4\7A\2\2\u00e4\u00ea\b\r\1"+
		"\2\u00e5\u00e6\7B\2\2\u00e6\u00ea\b\r\1\2\u00e7\u00e8\7C\2\2\u00e8\u00ea"+
		"\b\r\1\2\u00e9\u00db\3\2\2\2\u00e9\u00dd\3\2\2\2\u00e9\u00df\3\2\2\2\u00e9"+
		"\u00e1\3\2\2\2\u00e9\u00e3\3\2\2\2\u00e9\u00e5\3\2\2\2\u00e9\u00e7\3\2"+
		"\2\2\u00ea\31\3\2\2\2\f,\61?U\u0084\u0090\u009a\u00d6\u00d8\u00e9";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}