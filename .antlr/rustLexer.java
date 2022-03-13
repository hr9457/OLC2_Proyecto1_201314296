// Generated from c:\Users\joshu\Documents\USAC\1S2022\COMPI2\LAB\Proyecto\OLC2_Proyecto1_201314296\rust.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class rustLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "TK_PARENTESIS_LEFT", "TK_PARENTESIS_RIGHT", "TK_CORCHETE_LETF", 
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


	public rustLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "rust.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2G\u01d0\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\3\2\3\2\3\3\3\3"+
		"\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f"+
		"\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3"+
		"\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3"+
		"\31\3\32\3\32\3\32\3\33\3\33\3\33\3\34\3\34\3\34\3\35\3\35\3\35\3\36\3"+
		"\36\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3\"\3\"\3\""+
		"\3\"\3#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3"+
		"&\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3)\3)\3)\3)\3*\3*\3*\3*\3*\3*\3"+
		"*\3+\3+\3+\3,\3,\3,\3,\3,\3-\3-\3-\3-\3-\3-\3.\3.\3.\3/\3/\3/\3/\3/\3"+
		"/\3/\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3\62"+
		"\3\62\3\62\3\62\3\62\3\62\3\62\3\63\3\63\3\63\3\63\3\63\3\63\3\64\3\64"+
		"\3\64\3\64\3\65\3\65\3\65\3\65\3\66\3\66\3\66\3\66\3\66\3\67\3\67\3\67"+
		"\3\67\3\67\3\67\3\67\38\38\38\38\38\38\38\38\38\39\39\39\39\39\39\39\3"+
		":\3:\3:\3:\3:\3:\3:\3:\3:\3;\3;\3;\3;\3;\3;\3;\3;\3;\3;\3;\3;\3;\3;\3"+
		"<\3<\3<\3=\3=\3=\3=\3=\3>\3>\3>\3>\3>\3>\3>\3>\3?\3?\3?\3?\3?\3?\3@\6"+
		"@\u0190\n@\r@\16@\u0191\3A\6A\u0195\nA\rA\16A\u0196\3A\3A\6A\u019b\nA"+
		"\rA\16A\u019c\3B\3B\7B\u01a1\nB\fB\16B\u01a4\13B\3B\3B\3C\3C\7C\u01aa"+
		"\nC\fC\16C\u01ad\13C\3C\3C\3D\3D\7D\u01b3\nD\fD\16D\u01b6\13D\3E\3E\3"+
		"E\3E\3E\3E\5E\u01be\nE\3E\7E\u01c1\nE\fE\16E\u01c4\13E\3E\3E\5E\u01c8"+
		"\nE\3F\6F\u01cb\nF\rF\16F\u01cc\3F\3F\2\2G\3\3\5\4\7\5\t\6\13\7\r\b\17"+
		"\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+"+
		"\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+"+
		"U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081"+
		"B\u0083C\u0085D\u0087E\u0089F\u008bG\3\2\t\3\2\62;\3\2$$\5\2C\\aac|\6"+
		"\2\62;C\\aac|\4\2##\61\61\4\2\f\f\17\17\6\2\13\f\17\17\"\"^^\2\u01d9\2"+
		"\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2"+
		"\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2"+
		"\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2"+
		"\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2"+
		"\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2"+
		"\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2"+
		"\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U"+
		"\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2"+
		"\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2"+
		"\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{"+
		"\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085"+
		"\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\3\u008d\3\2\2"+
		"\2\5\u008f\3\2\2\2\7\u0091\3\2\2\2\t\u0093\3\2\2\2\13\u0095\3\2\2\2\r"+
		"\u0097\3\2\2\2\17\u0099\3\2\2\2\21\u009b\3\2\2\2\23\u009d\3\2\2\2\25\u009f"+
		"\3\2\2\2\27\u00a1\3\2\2\2\31\u00a3\3\2\2\2\33\u00a5\3\2\2\2\35\u00a7\3"+
		"\2\2\2\37\u00a9\3\2\2\2!\u00ab\3\2\2\2#\u00ad\3\2\2\2%\u00af\3\2\2\2\'"+
		"\u00b1\3\2\2\2)\u00b3\3\2\2\2+\u00b5\3\2\2\2-\u00b7\3\2\2\2/\u00b9\3\2"+
		"\2\2\61\u00bc\3\2\2\2\63\u00bf\3\2\2\2\65\u00c2\3\2\2\2\67\u00c5\3\2\2"+
		"\29\u00c8\3\2\2\2;\u00cb\3\2\2\2=\u00cd\3\2\2\2?\u00d0\3\2\2\2A\u00d9"+
		"\3\2\2\2C\u00dd\3\2\2\2E\u00e1\3\2\2\2G\u00e6\3\2\2\2I\u00eb\3\2\2\2K"+
		"\u00f2\3\2\2\2M\u00f7\3\2\2\2O\u00fd\3\2\2\2Q\u0101\3\2\2\2S\u0105\3\2"+
		"\2\2U\u010c\3\2\2\2W\u010f\3\2\2\2Y\u0114\3\2\2\2[\u011a\3\2\2\2]\u011d"+
		"\3\2\2\2_\u0124\3\2\2\2a\u0128\3\2\2\2c\u012d\3\2\2\2e\u0137\3\2\2\2g"+
		"\u013d\3\2\2\2i\u0141\3\2\2\2k\u0145\3\2\2\2m\u014a\3\2\2\2o\u0151\3\2"+
		"\2\2q\u015a\3\2\2\2s\u0161\3\2\2\2u\u016a\3\2\2\2w\u0178\3\2\2\2y\u017b"+
		"\3\2\2\2{\u0180\3\2\2\2}\u0188\3\2\2\2\177\u018f\3\2\2\2\u0081\u0194\3"+
		"\2\2\2\u0083\u019e\3\2\2\2\u0085\u01a7\3\2\2\2\u0087\u01b0\3\2\2\2\u0089"+
		"\u01c7\3\2\2\2\u008b\u01ca\3\2\2\2\u008d\u008e\7`\2\2\u008e\4\3\2\2\2"+
		"\u008f\u0090\7*\2\2\u0090\6\3\2\2\2\u0091\u0092\7+\2\2\u0092\b\3\2\2\2"+
		"\u0093\u0094\7]\2\2\u0094\n\3\2\2\2\u0095\u0096\7_\2\2\u0096\f\3\2\2\2"+
		"\u0097\u0098\7}\2\2\u0098\16\3\2\2\2\u0099\u009a\7\177\2\2\u009a\20\3"+
		"\2\2\2\u009b\u009c\7=\2\2\u009c\22\3\2\2\2\u009d\u009e\7<\2\2\u009e\24"+
		"\3\2\2\2\u009f\u00a0\7.\2\2\u00a0\26\3\2\2\2\u00a1\u00a2\7>\2\2\u00a2"+
		"\30\3\2\2\2\u00a3\u00a4\7@\2\2\u00a4\32\3\2\2\2\u00a5\u00a6\7\60\2\2\u00a6"+
		"\34\3\2\2\2\u00a7\u00a8\7?\2\2\u00a8\36\3\2\2\2\u00a9\u00aa\7-\2\2\u00aa"+
		" \3\2\2\2\u00ab\u00ac\7/\2\2\u00ac\"\3\2\2\2\u00ad\u00ae\7,\2\2\u00ae"+
		"$\3\2\2\2\u00af\u00b0\7\61\2\2\u00b0&\3\2\2\2\u00b1\u00b2\7\'\2\2\u00b2"+
		"(\3\2\2\2\u00b3\u00b4\7~\2\2\u00b4*\3\2\2\2\u00b5\u00b6\7(\2\2\u00b6,"+
		"\3\2\2\2\u00b7\u00b8\7#\2\2\u00b8.\3\2\2\2\u00b9\u00ba\7@\2\2\u00ba\u00bb"+
		"\7?\2\2\u00bb\60\3\2\2\2\u00bc\u00bd\7>\2\2\u00bd\u00be\7?\2\2\u00be\62"+
		"\3\2\2\2\u00bf\u00c0\7?\2\2\u00c0\u00c1\7?\2\2\u00c1\64\3\2\2\2\u00c2"+
		"\u00c3\7#\2\2\u00c3\u00c4\7?\2\2\u00c4\66\3\2\2\2\u00c5\u00c6\7~\2\2\u00c6"+
		"\u00c7\7~\2\2\u00c78\3\2\2\2\u00c8\u00c9\7(\2\2\u00c9\u00ca\7(\2\2\u00ca"+
		":\3\2\2\2\u00cb\u00cc\7A\2\2\u00cc<\3\2\2\2\u00cd\u00ce\7/\2\2\u00ce\u00cf"+
		"\7@\2\2\u00cf>\3\2\2\2\u00d0\u00d1\7r\2\2\u00d1\u00d2\7t\2\2\u00d2\u00d3"+
		"\7k\2\2\u00d3\u00d4\7p\2\2\u00d4\u00d5\7v\2\2\u00d5\u00d6\7n\2\2\u00d6"+
		"\u00d7\7p\2\2\u00d7\u00d8\7#\2\2\u00d8@\3\2\2\2\u00d9\u00da\7k\2\2\u00da"+
		"\u00db\78\2\2\u00db\u00dc\7\66\2\2\u00dcB\3\2\2\2\u00dd\u00de\7h\2\2\u00de"+
		"\u00df\78\2\2\u00df\u00e0\7\66\2\2\u00e0D\3\2\2\2\u00e1\u00e2\7d\2\2\u00e2"+
		"\u00e3\7q\2\2\u00e3\u00e4\7q\2\2\u00e4\u00e5\7n\2\2\u00e5F\3\2\2\2\u00e6"+
		"\u00e7\7e\2\2\u00e7\u00e8\7j\2\2\u00e8\u00e9\7c\2\2\u00e9\u00ea\7t\2\2"+
		"\u00eaH\3\2\2\2\u00eb\u00ec\7U\2\2\u00ec\u00ed\7v\2\2\u00ed\u00ee\7t\2"+
		"\2\u00ee\u00ef\7k\2\2\u00ef\u00f0\7p\2\2\u00f0\u00f1\7i\2\2\u00f1J\3\2"+
		"\2\2\u00f2\u00f3\7o\2\2\u00f3\u00f4\7c\2\2\u00f4\u00f5\7k\2\2\u00f5\u00f6"+
		"\7p\2\2\u00f6L\3\2\2\2\u00f7\u00f8\7w\2\2\u00f8\u00f9\7u\2\2\u00f9\u00fa"+
		"\7k\2\2\u00fa\u00fb\7|\2\2\u00fb\u00fc\7g\2\2\u00fcN\3\2\2\2\u00fd\u00fe"+
		"\7n\2\2\u00fe\u00ff\7g\2\2\u00ff\u0100\7v\2\2\u0100P\3\2\2\2\u0101\u0102"+
		"\7o\2\2\u0102\u0103\7w\2\2\u0103\u0104\7v\2\2\u0104R\3\2\2\2\u0105\u0106"+
		"\7u\2\2\u0106\u0107\7v\2\2\u0107\u0108\7t\2\2\u0108\u0109\7w\2\2\u0109"+
		"\u010a\7e\2\2\u010a\u010b\7v\2\2\u010bT\3\2\2\2\u010c\u010d\7c\2\2\u010d"+
		"\u010e\7u\2\2\u010eV\3\2\2\2\u010f\u0110\7v\2\2\u0110\u0111\7t\2\2\u0111"+
		"\u0112\7w\2\2\u0112\u0113\7g\2\2\u0113X\3\2\2\2\u0114\u0115\7h\2\2\u0115"+
		"\u0116\7c\2\2\u0116\u0117\7n\2\2\u0117\u0118\7u\2\2\u0118\u0119\7g\2\2"+
		"\u0119Z\3\2\2\2\u011a\u011b\7h\2\2\u011b\u011c\7p\2\2\u011c\\\3\2\2\2"+
		"\u011d\u011e\7t\2\2\u011e\u011f\7g\2\2\u011f\u0120\7v\2\2\u0120\u0121"+
		"\7w\2\2\u0121\u0122\7t\2\2\u0122\u0123\7p\2\2\u0123^\3\2\2\2\u0124\u0125"+
		"\7c\2\2\u0125\u0126\7d\2\2\u0126\u0127\7u\2\2\u0127`\3\2\2\2\u0128\u0129"+
		"\7u\2\2\u0129\u012a\7s\2\2\u012a\u012b\7t\2\2\u012b\u012c\7v\2\2\u012c"+
		"b\3\2\2\2\u012d\u012e\7v\2\2\u012e\u012f\7q\2\2\u012f\u0130\7a\2\2\u0130"+
		"\u0131\7u\2\2\u0131\u0132\7v\2\2\u0132\u0133\7t\2\2\u0133\u0134\7k\2\2"+
		"\u0134\u0135\7p\2\2\u0135\u0136\7i\2\2\u0136d\3\2\2\2\u0137\u0138\7e\2"+
		"\2\u0138\u0139\7n\2\2\u0139\u013a\7q\2\2\u013a\u013b\7p\2\2\u013b\u013c"+
		"\7g\2\2\u013cf\3\2\2\2\u013d\u013e\7p\2\2\u013e\u013f\7g\2\2\u013f\u0140"+
		"\7y\2\2\u0140h\3\2\2\2\u0141\u0142\7n\2\2\u0142\u0143\7g\2\2\u0143\u0144"+
		"\7p\2\2\u0144j\3\2\2\2\u0145\u0146\7r\2\2\u0146\u0147\7w\2\2\u0147\u0148"+
		"\7u\2\2\u0148\u0149\7j\2\2\u0149l\3\2\2\2\u014a\u014b\7t\2\2\u014b\u014c"+
		"\7g\2\2\u014c\u014d\7o\2\2\u014d\u014e\7q\2\2\u014e\u014f\7x\2\2\u014f"+
		"\u0150\7g\2\2\u0150n\3\2\2\2\u0151\u0152\7e\2\2\u0152\u0153\7q\2\2\u0153"+
		"\u0154\7p\2\2\u0154\u0155\7v\2\2\u0155\u0156\7c\2\2\u0156\u0157\7k\2\2"+
		"\u0157\u0158\7p\2\2\u0158\u0159\7u\2\2\u0159p\3\2\2\2\u015a\u015b\7k\2"+
		"\2\u015b\u015c\7p\2\2\u015c\u015d\7u\2\2\u015d\u015e\7g\2\2\u015e\u015f"+
		"\7t\2\2\u015f\u0160\7v\2\2\u0160r\3\2\2\2\u0161\u0162\7e\2\2\u0162\u0163"+
		"\7c\2\2\u0163\u0164\7r\2\2\u0164\u0165\7c\2\2\u0165\u0166\7e\2\2\u0166"+
		"\u0167\7k\2\2\u0167\u0168\7v\2\2\u0168\u0169\7{\2\2\u0169t\3\2\2\2\u016a"+
		"\u016b\7y\2\2\u016b\u016c\7k\2\2\u016c\u016d\7v\2\2\u016d\u016e\7j\2\2"+
		"\u016e\u016f\7a\2\2\u016f\u0170\7e\2\2\u0170\u0171\7c\2\2\u0171\u0172"+
		"\7r\2\2\u0172\u0173\7c\2\2\u0173\u0174\7e\2\2\u0174\u0175\7k\2\2\u0175"+
		"\u0176\7v\2\2\u0176\u0177\7{\2\2\u0177v\3\2\2\2\u0178\u0179\7k\2\2\u0179"+
		"\u017a\7h\2\2\u017ax\3\2\2\2\u017b\u017c\7g\2\2\u017c\u017d\7n\2\2\u017d"+
		"\u017e\7u\2\2\u017e\u017f\7g\2\2\u017fz\3\2\2\2\u0180\u0181\7g\2\2\u0181"+
		"\u0182\7n\2\2\u0182\u0183\7u\2\2\u0183\u0184\7g\2\2\u0184\u0185\7\"\2"+
		"\2\u0185\u0186\7k\2\2\u0186\u0187\7h\2\2\u0187|\3\2\2\2\u0188\u0189\7"+
		"y\2\2\u0189\u018a\7j\2\2\u018a\u018b\7k\2\2\u018b\u018c\7n\2\2\u018c\u018d"+
		"\7g\2\2\u018d~\3\2\2\2\u018e\u0190\t\2\2\2\u018f\u018e\3\2\2\2\u0190\u0191"+
		"\3\2\2\2\u0191\u018f\3\2\2\2\u0191\u0192\3\2\2\2\u0192\u0080\3\2\2\2\u0193"+
		"\u0195\5\177@\2\u0194\u0193\3\2\2\2\u0195\u0196\3\2\2\2\u0196\u0194\3"+
		"\2\2\2\u0196\u0197\3\2\2\2\u0197\u0198\3\2\2\2\u0198\u019a\7\60\2\2\u0199"+
		"\u019b\5\177@\2\u019a\u0199\3\2\2\2\u019b\u019c\3\2\2\2\u019c\u019a\3"+
		"\2\2\2\u019c\u019d\3\2\2\2\u019d\u0082\3\2\2\2\u019e\u01a2\7$\2\2\u019f"+
		"\u01a1\n\3\2\2\u01a0\u019f\3\2\2\2\u01a1\u01a4\3\2\2\2\u01a2\u01a0\3\2"+
		"\2\2\u01a2\u01a3\3\2\2\2\u01a3\u01a5\3\2\2\2\u01a4\u01a2\3\2\2\2\u01a5"+
		"\u01a6\7$\2\2\u01a6\u0084\3\2\2\2\u01a7\u01ab\7)\2\2\u01a8\u01aa\n\3\2"+
		"\2\u01a9\u01a8\3\2\2\2\u01aa\u01ad\3\2\2\2\u01ab\u01a9\3\2\2\2\u01ab\u01ac"+
		"\3\2\2\2\u01ac\u01ae\3\2\2\2\u01ad\u01ab\3\2\2\2\u01ae\u01af\7)\2\2\u01af"+
		"\u0086\3\2\2\2\u01b0\u01b4\t\4\2\2\u01b1\u01b3\t\5\2\2\u01b2\u01b1\3\2"+
		"\2\2\u01b3\u01b6\3\2\2\2\u01b4\u01b2\3\2\2\2\u01b4\u01b5\3\2\2\2\u01b5"+
		"\u0088\3\2\2\2\u01b6\u01b4\3\2\2\2\u01b7\u01b8\7\61\2\2\u01b8\u01b9\7"+
		"\61\2\2\u01b9\u01bd\3\2\2\2\u01ba\u01be\n\6\2\2\u01bb\u01bc\7\61\2\2\u01bc"+
		"\u01be\7\61\2\2\u01bd\u01ba\3\2\2\2\u01bd\u01bb\3\2\2\2\u01be\u01c2\3"+
		"\2\2\2\u01bf\u01c1\n\7\2\2\u01c0\u01bf\3\2\2\2\u01c1\u01c4\3\2\2\2\u01c2"+
		"\u01c0\3\2\2\2\u01c2\u01c3\3\2\2\2\u01c3\u01c8\3\2\2\2\u01c4\u01c2\3\2"+
		"\2\2\u01c5\u01c6\7\61\2\2\u01c6\u01c8\7\61\2\2\u01c7\u01b7\3\2\2\2\u01c7"+
		"\u01c5\3\2\2\2\u01c8\u008a\3\2\2\2\u01c9\u01cb\t\b\2\2\u01ca\u01c9\3\2"+
		"\2\2\u01cb\u01cc\3\2\2\2\u01cc\u01ca\3\2\2\2\u01cc\u01cd\3\2\2\2\u01cd"+
		"\u01ce\3\2\2\2\u01ce\u01cf\bF\2\2\u01cf\u008c\3\2\2\2\r\2\u0191\u0196"+
		"\u019c\u01a2\u01ab\u01b4\u01bd\u01c2\u01c7\u01cc\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}