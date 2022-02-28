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
		TK_INSERT=55, TK_CAPACITY=56, TK_WITHCAPACITY=57, TK_NUMBER=58, TK_DECIMAL=59, 
		TK_CADENA=60, TK_CARACTER=61, TK_ID=62, TK_COMMET=63, SPACES=64;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"TK_PARENTESIS_LEFT", "TK_PARENTESIS_RIGHT", "TK_CORCHETE_LETF", "TK_CORCHETE_RIGHT", 
			"TK_LLAVE_LEFT", "TK_LLAVE_RIGHT", "TK_PUNTO_COMA", "TK_DOSPUNTOS", "TK_COMA", 
			"TK_MENOR", "TK_MAYOR", "TK_PUNTO", "TK_IGUAL", "TK_MAS", "TK_MENOS", 
			"TK_POR", "TK_DIVISION", "TK_PORCENTAJE", "TK_BARRA", "TK_AMPERSAND", 
			"TK_ADMIRACION", "TK_MAYORIGULA", "TK_MENOIGUAL", "TK_IGUALIGUAL", "TK_DIFERENTE", 
			"TK_OR", "TK_AND", "TK_WHAT", "TK_TIPORETURN", "TK_IMPRESION", "TK_INT", 
			"TK_FLOAT", "TK_BOOL", "TK_CHAR", "TK_STRING", "TK_MAIN", "TK_USIZE", 
			"TK_LET", "TK_MUT", "TK_STRUCT", "TK_AS", "TK_TRUE", "TK_FALSE", "TK_FN", 
			"TK_RETURN", "TK_ABS", "TK_SQRT", "TK_TOSTRING", "TK_CLONE", "TK_NEW", 
			"TK_LEN", "TK_PUSH", "TK_REMOVED", "TK_CONTAINS", "TK_INSERT", "TK_CAPACITY", 
			"TK_WITHCAPACITY", "TK_NUMBER", "TK_DECIMAL", "TK_CADENA", "TK_CARACTER", 
			"TK_ID", "TK_COMMET", "SPACES"
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
			"'with_capacity'"
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
			"TK_WITHCAPACITY", "TK_NUMBER", "TK_DECIMAL", "TK_CADENA", "TK_CARACTER", 
			"TK_ID", "TK_COMMET", "SPACES"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2B\u01ae\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3"+
		"\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3"+
		"\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3"+
		"\26\3\27\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\32\3\32\3\32\3\33\3"+
		"\33\3\33\3\34\3\34\3\34\3\35\3\35\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3"+
		"\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3"+
		"#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3$\3$\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3&\3"+
		"\'\3\'\3\'\3\'\3(\3(\3(\3(\3)\3)\3)\3)\3)\3)\3)\3*\3*\3*\3+\3+\3+\3+\3"+
		"+\3,\3,\3,\3,\3,\3,\3-\3-\3-\3.\3.\3.\3.\3.\3.\3.\3/\3/\3/\3/\3\60\3\60"+
		"\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\62"+
		"\3\62\3\62\3\62\3\62\3\62\3\63\3\63\3\63\3\63\3\64\3\64\3\64\3\64\3\65"+
		"\3\65\3\65\3\65\3\65\3\66\3\66\3\66\3\66\3\66\3\66\3\66\3\67\3\67\3\67"+
		"\3\67\3\67\3\67\3\67\3\67\3\67\38\38\38\38\38\38\38\39\39\39\39\39\39"+
		"\39\39\39\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3:\3;\6;\u016e\n;\r;"+
		"\16;\u016f\3<\6<\u0173\n<\r<\16<\u0174\3<\3<\6<\u0179\n<\r<\16<\u017a"+
		"\3=\3=\7=\u017f\n=\f=\16=\u0182\13=\3=\3=\3>\3>\7>\u0188\n>\f>\16>\u018b"+
		"\13>\3>\3>\3?\3?\7?\u0191\n?\f?\16?\u0194\13?\3@\3@\3@\3@\3@\3@\5@\u019c"+
		"\n@\3@\7@\u019f\n@\f@\16@\u01a2\13@\3@\3@\5@\u01a6\n@\3A\6A\u01a9\nA\r"+
		"A\16A\u01aa\3A\3A\2\2B\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27"+
		"\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33"+
		"\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63"+
		"e\64g\65i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\3\2\t\3\2\62;\3\2$$\5\2"+
		"C\\aac|\6\2\62;C\\aac|\4\2##\61\61\4\2\f\f\17\17\6\2\13\f\17\17\"\"^^"+
		"\2\u01b7\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2"+
		"\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27"+
		"\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2"+
		"\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2"+
		"\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2"+
		"\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2"+
		"\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S"+
		"\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2"+
		"\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2"+
		"\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y"+
		"\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\3\u0083\3"+
		"\2\2\2\5\u0085\3\2\2\2\7\u0087\3\2\2\2\t\u0089\3\2\2\2\13\u008b\3\2\2"+
		"\2\r\u008d\3\2\2\2\17\u008f\3\2\2\2\21\u0091\3\2\2\2\23\u0093\3\2\2\2"+
		"\25\u0095\3\2\2\2\27\u0097\3\2\2\2\31\u0099\3\2\2\2\33\u009b\3\2\2\2\35"+
		"\u009d\3\2\2\2\37\u009f\3\2\2\2!\u00a1\3\2\2\2#\u00a3\3\2\2\2%\u00a5\3"+
		"\2\2\2\'\u00a7\3\2\2\2)\u00a9\3\2\2\2+\u00ab\3\2\2\2-\u00ad\3\2\2\2/\u00b0"+
		"\3\2\2\2\61\u00b3\3\2\2\2\63\u00b6\3\2\2\2\65\u00b9\3\2\2\2\67\u00bc\3"+
		"\2\2\29\u00bf\3\2\2\2;\u00c1\3\2\2\2=\u00c4\3\2\2\2?\u00cd\3\2\2\2A\u00d1"+
		"\3\2\2\2C\u00d5\3\2\2\2E\u00da\3\2\2\2G\u00df\3\2\2\2I\u00e6\3\2\2\2K"+
		"\u00eb\3\2\2\2M\u00f1\3\2\2\2O\u00f5\3\2\2\2Q\u00f9\3\2\2\2S\u0100\3\2"+
		"\2\2U\u0103\3\2\2\2W\u0108\3\2\2\2Y\u010e\3\2\2\2[\u0111\3\2\2\2]\u0118"+
		"\3\2\2\2_\u011c\3\2\2\2a\u0121\3\2\2\2c\u012b\3\2\2\2e\u0131\3\2\2\2g"+
		"\u0135\3\2\2\2i\u0139\3\2\2\2k\u013e\3\2\2\2m\u0145\3\2\2\2o\u014e\3\2"+
		"\2\2q\u0155\3\2\2\2s\u015e\3\2\2\2u\u016d\3\2\2\2w\u0172\3\2\2\2y\u017c"+
		"\3\2\2\2{\u0185\3\2\2\2}\u018e\3\2\2\2\177\u01a5\3\2\2\2\u0081\u01a8\3"+
		"\2\2\2\u0083\u0084\7*\2\2\u0084\4\3\2\2\2\u0085\u0086\7+\2\2\u0086\6\3"+
		"\2\2\2\u0087\u0088\7]\2\2\u0088\b\3\2\2\2\u0089\u008a\7_\2\2\u008a\n\3"+
		"\2\2\2\u008b\u008c\7}\2\2\u008c\f\3\2\2\2\u008d\u008e\7\177\2\2\u008e"+
		"\16\3\2\2\2\u008f\u0090\7=\2\2\u0090\20\3\2\2\2\u0091\u0092\7<\2\2\u0092"+
		"\22\3\2\2\2\u0093\u0094\7.\2\2\u0094\24\3\2\2\2\u0095\u0096\7>\2\2\u0096"+
		"\26\3\2\2\2\u0097\u0098\7@\2\2\u0098\30\3\2\2\2\u0099\u009a\7\60\2\2\u009a"+
		"\32\3\2\2\2\u009b\u009c\7?\2\2\u009c\34\3\2\2\2\u009d\u009e\7-\2\2\u009e"+
		"\36\3\2\2\2\u009f\u00a0\7/\2\2\u00a0 \3\2\2\2\u00a1\u00a2\7,\2\2\u00a2"+
		"\"\3\2\2\2\u00a3\u00a4\7\61\2\2\u00a4$\3\2\2\2\u00a5\u00a6\7\'\2\2\u00a6"+
		"&\3\2\2\2\u00a7\u00a8\7~\2\2\u00a8(\3\2\2\2\u00a9\u00aa\7(\2\2\u00aa*"+
		"\3\2\2\2\u00ab\u00ac\7#\2\2\u00ac,\3\2\2\2\u00ad\u00ae\7@\2\2\u00ae\u00af"+
		"\7?\2\2\u00af.\3\2\2\2\u00b0\u00b1\7>\2\2\u00b1\u00b2\7?\2\2\u00b2\60"+
		"\3\2\2\2\u00b3\u00b4\7?\2\2\u00b4\u00b5\7?\2\2\u00b5\62\3\2\2\2\u00b6"+
		"\u00b7\7#\2\2\u00b7\u00b8\7?\2\2\u00b8\64\3\2\2\2\u00b9\u00ba\7~\2\2\u00ba"+
		"\u00bb\7~\2\2\u00bb\66\3\2\2\2\u00bc\u00bd\7(\2\2\u00bd\u00be\7(\2\2\u00be"+
		"8\3\2\2\2\u00bf\u00c0\7A\2\2\u00c0:\3\2\2\2\u00c1\u00c2\7/\2\2\u00c2\u00c3"+
		"\7@\2\2\u00c3<\3\2\2\2\u00c4\u00c5\7r\2\2\u00c5\u00c6\7t\2\2\u00c6\u00c7"+
		"\7k\2\2\u00c7\u00c8\7p\2\2\u00c8\u00c9\7v\2\2\u00c9\u00ca\7n\2\2\u00ca"+
		"\u00cb\7p\2\2\u00cb\u00cc\7#\2\2\u00cc>\3\2\2\2\u00cd\u00ce\7k\2\2\u00ce"+
		"\u00cf\78\2\2\u00cf\u00d0\7\66\2\2\u00d0@\3\2\2\2\u00d1\u00d2\7h\2\2\u00d2"+
		"\u00d3\78\2\2\u00d3\u00d4\7\66\2\2\u00d4B\3\2\2\2\u00d5\u00d6\7d\2\2\u00d6"+
		"\u00d7\7q\2\2\u00d7\u00d8\7q\2\2\u00d8\u00d9\7n\2\2\u00d9D\3\2\2\2\u00da"+
		"\u00db\7e\2\2\u00db\u00dc\7j\2\2\u00dc\u00dd\7c\2\2\u00dd\u00de\7t\2\2"+
		"\u00deF\3\2\2\2\u00df\u00e0\7U\2\2\u00e0\u00e1\7v\2\2\u00e1\u00e2\7t\2"+
		"\2\u00e2\u00e3\7k\2\2\u00e3\u00e4\7p\2\2\u00e4\u00e5\7i\2\2\u00e5H\3\2"+
		"\2\2\u00e6\u00e7\7o\2\2\u00e7\u00e8\7c\2\2\u00e8\u00e9\7k\2\2\u00e9\u00ea"+
		"\7p\2\2\u00eaJ\3\2\2\2\u00eb\u00ec\7w\2\2\u00ec\u00ed\7u\2\2\u00ed\u00ee"+
		"\7k\2\2\u00ee\u00ef\7|\2\2\u00ef\u00f0\7g\2\2\u00f0L\3\2\2\2\u00f1\u00f2"+
		"\7n\2\2\u00f2\u00f3\7g\2\2\u00f3\u00f4\7v\2\2\u00f4N\3\2\2\2\u00f5\u00f6"+
		"\7o\2\2\u00f6\u00f7\7w\2\2\u00f7\u00f8\7v\2\2\u00f8P\3\2\2\2\u00f9\u00fa"+
		"\7u\2\2\u00fa\u00fb\7v\2\2\u00fb\u00fc\7t\2\2\u00fc\u00fd\7w\2\2\u00fd"+
		"\u00fe\7e\2\2\u00fe\u00ff\7v\2\2\u00ffR\3\2\2\2\u0100\u0101\7c\2\2\u0101"+
		"\u0102\7u\2\2\u0102T\3\2\2\2\u0103\u0104\7v\2\2\u0104\u0105\7t\2\2\u0105"+
		"\u0106\7w\2\2\u0106\u0107\7g\2\2\u0107V\3\2\2\2\u0108\u0109\7h\2\2\u0109"+
		"\u010a\7c\2\2\u010a\u010b\7n\2\2\u010b\u010c\7u\2\2\u010c\u010d\7g\2\2"+
		"\u010dX\3\2\2\2\u010e\u010f\7h\2\2\u010f\u0110\7p\2\2\u0110Z\3\2\2\2\u0111"+
		"\u0112\7t\2\2\u0112\u0113\7g\2\2\u0113\u0114\7v\2\2\u0114\u0115\7w\2\2"+
		"\u0115\u0116\7t\2\2\u0116\u0117\7p\2\2\u0117\\\3\2\2\2\u0118\u0119\7c"+
		"\2\2\u0119\u011a\7d\2\2\u011a\u011b\7u\2\2\u011b^\3\2\2\2\u011c\u011d"+
		"\7u\2\2\u011d\u011e\7s\2\2\u011e\u011f\7t\2\2\u011f\u0120\7v\2\2\u0120"+
		"`\3\2\2\2\u0121\u0122\7v\2\2\u0122\u0123\7q\2\2\u0123\u0124\7a\2\2\u0124"+
		"\u0125\7u\2\2\u0125\u0126\7v\2\2\u0126\u0127\7t\2\2\u0127\u0128\7k\2\2"+
		"\u0128\u0129\7p\2\2\u0129\u012a\7i\2\2\u012ab\3\2\2\2\u012b\u012c\7e\2"+
		"\2\u012c\u012d\7n\2\2\u012d\u012e\7q\2\2\u012e\u012f\7p\2\2\u012f\u0130"+
		"\7g\2\2\u0130d\3\2\2\2\u0131\u0132\7p\2\2\u0132\u0133\7g\2\2\u0133\u0134"+
		"\7y\2\2\u0134f\3\2\2\2\u0135\u0136\7n\2\2\u0136\u0137\7g\2\2\u0137\u0138"+
		"\7p\2\2\u0138h\3\2\2\2\u0139\u013a\7r\2\2\u013a\u013b\7w\2\2\u013b\u013c"+
		"\7u\2\2\u013c\u013d\7j\2\2\u013dj\3\2\2\2\u013e\u013f\7t\2\2\u013f\u0140"+
		"\7g\2\2\u0140\u0141\7o\2\2\u0141\u0142\7q\2\2\u0142\u0143\7x\2\2\u0143"+
		"\u0144\7g\2\2\u0144l\3\2\2\2\u0145\u0146\7e\2\2\u0146\u0147\7q\2\2\u0147"+
		"\u0148\7p\2\2\u0148\u0149\7v\2\2\u0149\u014a\7c\2\2\u014a\u014b\7k\2\2"+
		"\u014b\u014c\7p\2\2\u014c\u014d\7u\2\2\u014dn\3\2\2\2\u014e\u014f\7k\2"+
		"\2\u014f\u0150\7p\2\2\u0150\u0151\7u\2\2\u0151\u0152\7g\2\2\u0152\u0153"+
		"\7t\2\2\u0153\u0154\7v\2\2\u0154p\3\2\2\2\u0155\u0156\7e\2\2\u0156\u0157"+
		"\7c\2\2\u0157\u0158\7r\2\2\u0158\u0159\7c\2\2\u0159\u015a\7e\2\2\u015a"+
		"\u015b\7k\2\2\u015b\u015c\7v\2\2\u015c\u015d\7{\2\2\u015dr\3\2\2\2\u015e"+
		"\u015f\7y\2\2\u015f\u0160\7k\2\2\u0160\u0161\7v\2\2\u0161\u0162\7j\2\2"+
		"\u0162\u0163\7a\2\2\u0163\u0164\7e\2\2\u0164\u0165\7c\2\2\u0165\u0166"+
		"\7r\2\2\u0166\u0167\7c\2\2\u0167\u0168\7e\2\2\u0168\u0169\7k\2\2\u0169"+
		"\u016a\7v\2\2\u016a\u016b\7{\2\2\u016bt\3\2\2\2\u016c\u016e\t\2\2\2\u016d"+
		"\u016c\3\2\2\2\u016e\u016f\3\2\2\2\u016f\u016d\3\2\2\2\u016f\u0170\3\2"+
		"\2\2\u0170v\3\2\2\2\u0171\u0173\5u;\2\u0172\u0171\3\2\2\2\u0173\u0174"+
		"\3\2\2\2\u0174\u0172\3\2\2\2\u0174\u0175\3\2\2\2\u0175\u0176\3\2\2\2\u0176"+
		"\u0178\7\60\2\2\u0177\u0179\5u;\2\u0178\u0177\3\2\2\2\u0179\u017a\3\2"+
		"\2\2\u017a\u0178\3\2\2\2\u017a\u017b\3\2\2\2\u017bx\3\2\2\2\u017c\u0180"+
		"\7$\2\2\u017d\u017f\n\3\2\2\u017e\u017d\3\2\2\2\u017f\u0182\3\2\2\2\u0180"+
		"\u017e\3\2\2\2\u0180\u0181\3\2\2\2\u0181\u0183\3\2\2\2\u0182\u0180\3\2"+
		"\2\2\u0183\u0184\7$\2\2\u0184z\3\2\2\2\u0185\u0189\7)\2\2\u0186\u0188"+
		"\n\3\2\2\u0187\u0186\3\2\2\2\u0188\u018b\3\2\2\2\u0189\u0187\3\2\2\2\u0189"+
		"\u018a\3\2\2\2\u018a\u018c\3\2\2\2\u018b\u0189\3\2\2\2\u018c\u018d\7)"+
		"\2\2\u018d|\3\2\2\2\u018e\u0192\t\4\2\2\u018f\u0191\t\5\2\2\u0190\u018f"+
		"\3\2\2\2\u0191\u0194\3\2\2\2\u0192\u0190\3\2\2\2\u0192\u0193\3\2\2\2\u0193"+
		"~\3\2\2\2\u0194\u0192\3\2\2\2\u0195\u0196\7\61\2\2\u0196\u0197\7\61\2"+
		"\2\u0197\u019b\3\2\2\2\u0198\u019c\n\6\2\2\u0199\u019a\7\61\2\2\u019a"+
		"\u019c\7\61\2\2\u019b\u0198\3\2\2\2\u019b\u0199\3\2\2\2\u019c\u01a0\3"+
		"\2\2\2\u019d\u019f\n\7\2\2\u019e\u019d\3\2\2\2\u019f\u01a2\3\2\2\2\u01a0"+
		"\u019e\3\2\2\2\u01a0\u01a1\3\2\2\2\u01a1\u01a6\3\2\2\2\u01a2\u01a0\3\2"+
		"\2\2\u01a3\u01a4\7\61\2\2\u01a4\u01a6\7\61\2\2\u01a5\u0195\3\2\2\2\u01a5"+
		"\u01a3\3\2\2\2\u01a6\u0080\3\2\2\2\u01a7\u01a9\t\b\2\2\u01a8\u01a7\3\2"+
		"\2\2\u01a9\u01aa\3\2\2\2\u01aa\u01a8\3\2\2\2\u01aa\u01ab\3\2\2\2\u01ab"+
		"\u01ac\3\2\2\2\u01ac\u01ad\bA\2\2\u01ad\u0082\3\2\2\2\r\2\u016f\u0174"+
		"\u017a\u0180\u0189\u0192\u019b\u01a0\u01a5\u01aa\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}