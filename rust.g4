grammar rust;

//impor necesarios para ser utilizados se ponene al inicio del codigo generado
@parser::header{ 
    import "Proyecto1/symbol"
}

// insertar atributos en la clase generada
@parser::members{
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
    :   instrucciones EOF {}
    ; 

// instruccione o varias instrucciones
instrucciones
    :   variable    { fmt.Println($variable.v) }
    |   impresion   {}
    ;

// declaracion de variables
variable returns[interface{} v]
    :   TK_LET TK_MUT TK_ID '=' expresion ';'           {}
    |   TK_LET TK_ID '=' expresion ';'                  {}
    |   TK_LET TK_MUT TK_ID ':' tipo '=' expresion ';'  {}
    |   TK_LET TK_ID ':' tipo '=' expresion ';'         {$v = symbol.NewSimbolo($TK_ID.text,$expresion.text,false,0)}
    ;

// tipos aceptados en el lenguaje
tipo
    :   TK_INT
    |   TK_FLOAT
    |   TK_BOOL
    |   TK_CHAR
    |   TK_STRING
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
    :   TK_NUMBER   
    |   TK_DECIMAL
    |   TK_CADENA
    |   TK_CARACTER
    |   TK_ID
    ;

// impreiones
impresion 
    :   TK_IMPRESION '(' expresion ')' ';'  {}
    ;