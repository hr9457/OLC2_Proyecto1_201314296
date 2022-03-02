grammar rust;

//impor necesarios para ser utilizados se ponene al inicio del codigo generado
@parser::header{ 
    import "Proyecto1/src/interfaces"
    import "Proyecto1/src/expressiones"
    import "Proyecto1/src/instrucciones"
    import "Proyecto1/src/pruebas"
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

TK_MAIN:        'main';
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


TK_IF:   'if';
TK_ELSE: 'else';



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
    :   funcionmain EOF {}
    ; 

funcionmain
    :   TK_FN TK_MAIN '(' ')' '{'    '}' 
    |   TK_FN TK_MAIN '(' ')' '{'  instrucciones*  '}' 
    ;


// instruccione o varias instrucciones
instrucciones
    :   variable    { fmt.Println($variable.nuevaVariable) }
    |   impresion   { pruebas.Probar($impresion.p)}
    |   expresionIf {}
    ;


// impresion
impresion returns[interfaces.Instruction p]
    :   TK_IMPRESION '(' expresion ')' ';'  {
        fmt.Println("Aca este println")
        fmt.Println($expresion.primate)
        $p = instrucciones.NewImprimir($expresion.primate)
        }
    ;


// instruccion if
expresionIf
    :   sintaxisIf              {} 
    |   sintaxisIf sintaxisElse {}
    ;

sintaxisIf
    :   TK_IF expresion '{' '}' {} 
    ;

sintaxisElse
    :   TK_ELSE '{' '}'
    ;



// declaracion de variables
variable returns[interface{} nuevaVariable]
    :   TK_LET TK_MUT TK_ID '=' expresion ';'           {$nuevaVariable = interfaces.NewSimbolo($TK_ID.text,$expresion.text,true,tipoValor)}
    |   TK_LET TK_ID '=' expresion ';'                  {$nuevaVariable = interfaces.NewSimbolo($TK_ID.text,$expresion.text,false,tipoValor)}
    |   TK_LET TK_MUT TK_ID ':' tipo '=' expresion ';'  {$nuevaVariable = interfaces.NewSimbolo($TK_ID.text,$expresion.text,true,$tipo.tipoExp)}
    |   TK_LET TK_ID ':' tipo '=' expresion ';'         {$nuevaVariable = interfaces.NewSimbolo($TK_ID.text,$expresion.text,false,$tipo.tipoExp)}
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
expresion returns[interfaces.Expresion primate]
    :   expresion ('*'|'/'|'%') expresion
    |   left=expresion op=('+'|'-') right=expresion {$primate = expressiones.NewOperacion($left.primate,$op.text,$right.primate)}
    |   expresion '&' expresion
    |   expresion '^' expresion
    |   expresion '|' 
    |   expresion '==' expresion
    |   expresion '!=' expresion
    |   expresion '>' expresion
    |   expresion '<' expresion
    |   expresion '>=' expresion
    |   expresion '<=' expresion
    |   expresion '&&' expresion
    |   expresion '||' expresion
    |   '(' expresion ')'
    |   valor   {
                    fmt.Println("->>>>")
                    fmt.Println($valor.primate)
                    $primate = $valor.primate
                }
    ;

// valor aceptados en la gramaticas
valor returns[interfaces.Expresion primate]
    :   TK_NUMBER   {$primate = expressiones.NewPrimito(interfaces.ConvertTextInt($TK_NUMBER.text),interfaces.INTEGER)}
    |   TK_DECIMAL  {$primate = expressiones.NewPrimito(interfaces.ConvertTextFloat($TK_DECIMAL.text),interfaces.FLOAT)}
    |   TK_TRUE     {}
    |   TK_FALSE    {}
    |   TK_CADENA   {}
    |   TK_CARACTER {}
    |   TK_ID       {}
    ;
