grammar rust;

//impor necesarios para ser utilizados se ponene al inicio del codigo generado
@parser::header{ 
    import "Proyecto1/src/symbol"
    import "Proyecto1/src/interfaces"
}

// insertar atributos en la clase generada
@parser::members{
    var tipoValor = interfaces.NULL
}


// tokens -> simbolos reservados  
TK_PARENTESIS_LEFT:    '(';
TK_PARENTESIS_RIGHT:   ')';
TK_CORCHETE_LETF:      '[';
TK_CORCHETE_RIGHT:     ']';
TK_LLAVE_LEFT:         '{';
TK_LLAVE_RIGHT:        '}';
TK_PUNTO_COMA:         ';';
TK_DOSPUNTOS :         ':';
TK_COMA:               ',';
TK_MENOR:              '<';
TK_MAYOR:              '>';
TK_PUNTO:              '.';
TK_IGUAL:              '=';
TK_MAS:                '+';
TK_MENOS:              '-';
TK_POR:                '*';
TK_DIVISION:           '/';
TK_PORCENTAJE:         '%';
TK_BARRA:              '|';
TK_AMPERSAND:          '&';
TK_ADMIRACION:         '!';
TK_MAYORIGULA:         '>=';
TK_MENOIGUAL:          '<=';
TK_IGUALIGUAL:         '==';
TK_DIFERENTE:          '!='; 
TK_OR:                 '||'; 
TK_AND:                '&&';
TK_WHAT:               '?';
TK_TIPORETURN:         '->';



// token reservados por el lenguaje
TK_IMPRESION:   'println!';

TK_INT:         'i64';
TK_FLOAT:       'f64';
TK_BOOL:        'bool'; 
TK_CHAR:        'char'; 
TK_STRING:      'String';

TK_USIZE:       'usize';
TK_LET:         'let';
TK_MUT:         'mut';
TK_STRUCT:      'struct';
TK_AS:          'as';
TK_TRUE:        'true';
TK_FALSE:       'false';
TK_FN:          'fn'; 
TK_RETURN:      'return';
TK_ABS:         'abs';
TK_SQRT:        'sqrt';
TK_TOSTRING:    'to_string';
TK_CLONE:       'clone'; 
TK_NEW:         'new';
TK_LEN:         'len';
TK_PUSH:        'push';
TK_REMOVED:     'remove';
TK_CONTAINS:    'contains';
TK_INSERT:      'insert';
TK_CAPACITY:    'capacity';
TK_WITHCAPACITY:'with_capacity';





// identificadores
TK_NUMBER:  [0-9]+;
TK_DECIMAL: TK_NUMBER+'.'TK_NUMBER+;
TK_CADENA:  '"'~["]*'"';
TK_CARACTER:'\''~["]*'\'';
TK_ID:      ([a-zA-Z_])[a-zA-Z0-9_]*;
TK_COMMET:  ('//' (~[/!] | '//') ~[\r\n]* | '//');


//ignoracon de espacios
SPACES: [ \\\r\n\t]+ -> skip;


// gramatica
start 
    :   instrucciones* EOF {}
    ; 

// instruccione o varias instrucciones
instrucciones
    :   variable    { fmt.Println($variable.nuevaVariable) }
    |   impresion   {}
    ;

// declaracion de variables
variable returns[interface{} nuevaVariable]
    :   TK_LET TK_MUT TK_ID '=' expresion ';'           {$nuevaVariable = symbol.NewSimbolo($TK_ID.text,$expresion.text,true,tipoValor)}
    |   TK_LET TK_ID '=' expresion ';'                  {$nuevaVariable = symbol.NewSimbolo($TK_ID.text,$expresion.text,false,tipoValor)}
    |   TK_LET TK_MUT TK_ID ':' tipo '=' expresion ';'  {$nuevaVariable = symbol.NewSimbolo($TK_ID.text,$expresion.text,true,$tipo.tipoExp)}
    |   TK_LET TK_ID ':' tipo '=' expresion ';'         {$nuevaVariable = symbol.NewSimbolo($TK_ID.text,$expresion.text,false,$tipo.tipoExp)}
    ;

// tipos aceptados en el lenguaje
tipo returns[interfaces.TipoExpression tipoExp]
    :   TK_INT      {$tipoExp = interfaces.INTEGER}
    |   TK_FLOAT    {$tipoExp = interfaces.FLOAT}
    |   TK_BOOL     {$tipoExp = interfaces.BOOLEAN}
    |   TK_CHAR     {$tipoExp = interfaces.CHAR}
    |   TK_STRING   {$tipoExp = interfaces.STRING}
    ;

// expresiones aceptadas en el lenguaje
expresion 
    :   expresion '*' expresion
    |   expresion '/' expresion
    |   expresion '+' expresion
    |   expresion '-' expresion
    |   valor   
    ;

// valor aceptados en la gramaticas
valor 
    :   TK_NUMBER   {tipoValor = interfaces.INTEGER}
    |   TK_DECIMAL  {tipoValor = interfaces.FLOAT}
    |   TK_TRUE     {tipoValor = interfaces.BOOLEAN}
    |   TK_FALSE    {tipoValor = interfaces.BOOLEAN}
    |   TK_CADENA   {tipoValor = interfaces.STRING}
    |   TK_CARACTER {tipoValor = interfaces.CHAR}
    |   TK_ID       {tipoValor = interfaces.IDENTIFICADOR}
    ;

// impreiones
impresion 
    :   TK_IMPRESION '(' expresion ')' ';'  {}
    ;