grammar rust;

//impor necesarios para ser utilizados se ponene al inicio del codigo generado
@parser::header{ 
    import "Proyecto1/src/interfaces"
    import "Proyecto1/src/expressiones"
    import "Proyecto1/src/instrucciones"
    import "Proyecto1/src/arreglos"
    import arrayList "github.com/colegno/arraylist"
    // import "Proyecto1/src/pruebas" 
    // import "reflect"
}

// insertar atributos en la clase generada
@parser::members{
    var tipoValor = interfaces.NULL
}




// ********************************
// Inicio de la gramatica - sintactica
// ********************************
start returns[*arrayList.List lista]
    :   funcionmain EOF     { $lista = $funcionmain.lista }
    // |   listafn funcionmain EOF     { $lista = $funcionmain.lista }
    ; 


// funcionmain EOF     {pruebas.Probar($funcionmain.lista)}







// ********************************
// Inicio de la gramatica - sintactica
// ********************************
funcionmain returns [*arrayList.List lista]    
    :   TK_FN TK_MAIN '(' ')' '{'    '}' 
    |   TK_FN TK_MAIN '(' ')' '{'  instrucciones  '}'   { $lista = $instrucciones.lista }    
    ;




// ********************************
// lista de de funciones 
// ********************************

// listafn returns[*arrayList.List lista]
//     @init{$lista = arrayList.New()}
//     :   e += funcion+
//         {
//             listInt := localctx.(*ListafnContext).GetE()
//             for _,e:= range listInt{
//                 $lista.Add(e.GetInst())
//             }
//         }
//     ;




// ********************************
// gramatica para creacion de funciones del programa
// ********************************

// funcion returns[interfaces.Instruction inst]
//     :   TK_FN TK_ID '(' ')' '{' instrucciones '}'     { funciones.NewFuncion($TK_ID.text,$instrucciones.lista) }
//     ;







// ********************************
// concatenacion y conjutos de todas las intrucciones juntas
// ********************************
instrucciones returns[*arrayList.List lista]
    @init{$lista = arrayList.New()}
    : e += instruccion*
        {
            listInt := localctx.(*InstruccionesContext).GetE()
            for _,e:= range listInt{
                $lista.Add(e.GetInst())
            }
        }  
    ;


// ********************************
// instrucciones aceptadas por el lenguaje
// ********************************
instruccion returns[interfaces.Instruction inst]
    :   variable            { $inst = $variable.inst }
    |   impresion           { $inst = $impresion.inst }
    |   asignacionVariable  { $inst = $asignacionVariable.inst }
    |   expresionIf         { $inst = $expresionIf.inst }
    |   expresionWhile      { $inst = $expresionWhile.inst }
    |   expresionLoop       { $inst = $expresionLoop.inst  }
    |   breakInst           { $inst = $breakInst.inst }
    |   declarcionarreglo   { $inst = $declarcionarreglo.inst }
    ;






// ********************************
// Instruccion para llamado a funciones
// ********************************

// llamadofn returns[interfaces.Instruction inst]
//     :   TK_ID '(' ')' ';'   { $inst = funciones.NewLlamdoFn($TK_ID.text,nil) } 
//     ;






// ********************************
// Instruccion para imprimr
// ********************************
impresion returns[interfaces.Instruction inst]
    :   TK_IMPRESION '(' expresion ')' ';'  {$inst = instrucciones.NewImprimir($expresion.primate,nil)}
    |   TK_IMPRESION '(' expresion  listprint ')'  ';'  {$inst = instrucciones.NewImprimir($expresion.primate,$listprint.lista)}
    ;



listprint returns[*arrayList.List lista]
    @init{$lista = arrayList.New()}
    : list += expimprimir+
        {
            localList := localctx.(*ListprintContext).GetList()
            for _, list := range localList{
                $lista.Add(list.GetPrimate())
            }
        }
    ;

expimprimir returns[interfaces.Expresion primate]
    :   ',' expresion   {$primate = $expresion.primate}
    ;






// ********************************
// declaracion de variables
// ********************************
variable returns[interfaces.Instruction inst]
    :   TK_LET TK_MUT TK_ID '=' expresion ';'           { $inst = instrucciones.NewDeclaracion($TK_ID.text,true,interfaces.NULL,$expresion.primate)  }
    |   TK_LET TK_ID '=' expresion ';'                  { $inst = instrucciones.NewDeclaracion($TK_ID.text,false,interfaces.NULL,$expresion.primate) }
    |   TK_LET TK_MUT TK_ID ':' tipo '=' expresion ';'  { $inst = instrucciones.NewDeclaracion($TK_ID.text,true,$tipo.tipoExp,$expresion.primate)    }
    |   TK_LET TK_ID ':' tipo '=' expresion ';'         { $inst = instrucciones.NewDeclaracion($TK_ID.text,false,$tipo.tipoExp,$expresion.primate)   }
    ;





// ********************************
// Instruccion para asignar valores a una variable
// ********************************
asignacionVariable returns[interfaces.Instruction inst]
    : TK_ID '=' expresion ';'  { $inst = instrucciones.NewAsignacion($TK_ID.text,$expresion.primate) }
    ;





// ********************************
// Instruccion para declara una variable
// ********************************
declarcionarreglo returns[interfaces.Instruction inst]
    :   TK_LET TK_ID ':' '[' tipo ';' valor ']' '=' expresion ';'
    { $inst = arreglos.NewArreglo($TK_ID.text, false, $tipo.tipoExp, $valor.text, true, $expresion.primate) }
    ;




// **********************************
// expresiones listado de expresiones
// **********************************
listvalores returns[*arrayList.List lista]
    :   l = listvalores ',' expresion
        {
            $l.lista.Add($expresion.primate)
            $lista = $l.lista
        }

    | expresion
        {
            $lista = arrayList.New()
            $lista.Add($expresion.primate)
        }
    ;






// ********************************
// Instruccion para la creacion de un if y else
// ********************************
expresionIf returns[interfaces.Instruction inst]
    :   TK_IF expresion '{' bloqueif=instrucciones '}'      
        { $inst = instrucciones.NewIf($expresion.primate,$bloqueif.lista,nil) } 

    |   TK_IF expresion '{' bloqueif=instrucciones '}' TK_ELSE '{' bloqueelse=instrucciones '}'     
        { $inst = instrucciones.NewIf($expresion.primate,$bloqueif.lista,$bloqueelse.lista) }


    |   TK_IF expresion '{' ifblock=instrucciones '}' listaelif
        { $inst = instrucciones.NewIf2($expresion.primate,$ifblock.lista,nil,$listaelif.lista)  }

    |   TK_IF expresion '{' ifblock=instrucciones '}' listaelif TK_ELSE '{' bloqueelse=instrucciones '}'
        { $inst = instrucciones.NewIf2($expresion.primate,$ifblock.lista,$bloqueelse.lista,$listaelif.lista)  }
    ;



// concatenacion de la lista de else if
listaelif returns[*arrayList.List lista]
    @init{ $lista = arrayList.New() }
    :   l += elif+
        {
            listLocal := localctx.(*ListaelifContext).GetL()
            for _,l := range listLocal{
                $lista.Add(l.GetInst())
            }
        }
    ;


// else if
elif returns[interfaces.Instruction inst]
    :   TK_ELSE_IF expresion '{' elifblock=instrucciones '}'
        { $inst = instrucciones.NewIf($expresion.primate,$elifblock.lista,nil)  }
    ;






// ********************************
// Instruccion para la creacion de un if y else
// ********************************
expresionWhile returns[interfaces.Instruction inst]
    :   TK_WHILE exp=expresion '{' instrucciones '}'    { $inst = instrucciones.NewWhile($exp.primate,$instrucciones.lista) } 
    ;






// ********************************
// Instruccion para la creacion de un loop
// ********************************
expresionLoop returns[interfaces.Instruction inst]
    :   TK_LOOP '{' instrucciones '}'   { $inst = instrucciones.NewLoop($instrucciones.lista) }
    ;







// ********************************
// Sentencia de control break
// ********************************
breakInst returns[interfaces.Instruction inst]
    :   TK_BREAK ';'    {   $inst = instrucciones.NewBreak(interfaces.BREAK)    }
    ;














// ********************************
// tipos aceptados en el lenguaje
// ********************************
tipo returns[interfaces.TipoExpression tipoExp]
    :   TK_INT      {$tipoExp = interfaces.INTEGER}
    |   TK_FLOAT    {$tipoExp = interfaces.FLOAT}
    |   TK_BOOL     {$tipoExp = interfaces.BOOLEAN}
    |   TK_CHAR     {$tipoExp = interfaces.CHAR}
    |   TK_STRING   {$tipoExp = interfaces.STRING}
    ;















// **********************************
// expresiones aceptadas en el lenguaje
// *************************************
expresion returns[interfaces.Expresion primate]
    :   op=('!'|'-') right=expresion                    { $primate = expressiones.NewOperacion(nil,$op.text,$right.primate) }
    |   left=expresion op=('*'|'/'|'%') right=expresion { $primate = expressiones.NewOperacion($left.primate,$op.text,$right.primate) }
    |   left=expresion op=('+'|'-') right=expresion     { $primate = expressiones.NewOperacion($left.primate,$op.text,$right.primate) }
    |   TK_POW_I64 '(' left=expresion ',' right=expresion ')'      { $primate = expressiones.NewOperacion($left.primate,$TK_POW_I64.text,$right.primate) }
    |   TK_POW_F64 '(' left=expresion ',' right=expresion ')'      { $primate = expressiones.NewOperacion($left.primate,$TK_POW_F64.text,$right.primate) }
    |   left=expresion op='==' right=expresion          { $primate = expressiones.NewOperacion($left.primate,$op.text,$right.primate) }
    |   left=expresion op='!=' right=expresion          { $primate = expressiones.NewOperacion($left.primate,$op.text,$right.primate) }
    |   left=expresion op='>'  right=expresion          { $primate = expressiones.NewOperacion($left.primate,$op.text,$right.primate) }
    |   left=expresion op='<'  right=expresion          { $primate = expressiones.NewOperacion($left.primate,$op.text,$right.primate) }
    |   left=expresion op='>=' right=expresion          { $primate = expressiones.NewOperacion($left.primate,$op.text,$right.primate) }
    |   left=expresion op='<=' right=expresion          { $primate = expressiones.NewOperacion($left.primate,$op.text,$right.primate) }
    |   left=expresion op='&&' right=expresion          { $primate = expressiones.NewOperacion($left.primate,$op.text,$right.primate) }
    |   left=expresion op='||' right=expresion          { $primate = expressiones.NewOperacion($left.primate,$op.text,$right.primate) }
    |   '(' expresion ')'                               { $primate = $valor.primate }
    |   '[' listvalores ']'                             { $primate = arreglos.NewArray($listvalores.lista) }
    |   valor                                           { $primate = $valor.primate }
    |   
    ;








// ********************************
// valor aceptados en la gramaticas
// ********************************
valor returns[interfaces.Expresion primate]
    :   TK_NUMBER   { $primate = expressiones.NewPrimito(interfaces.ConvertTextInt($TK_NUMBER.text),interfaces.INTEGER) }
    |   TK_DECIMAL  { $primate = expressiones.NewPrimito(interfaces.ConvertTextFloat($TK_DECIMAL.text),interfaces.FLOAT) }
    |   TK_TRUE     { $primate = expressiones.NewPrimito(interfaces.ConvertTextBool($TK_TRUE.text),interfaces.BOOLEAN) }
    |   TK_FALSE    { $primate = expressiones.NewPrimito(interfaces.ConvertTextBool($TK_FALSE.text),interfaces.BOOLEAN) }
    |   TK_CADENA   { $primate = expressiones.NewPrimito(interfaces.ConvertTextString($TK_CADENA.text),interfaces.STRING) }
    |   TK_CARACTER { $primate = expressiones.NewPrimito(interfaces.ConvertTextString($TK_CARACTER.text),interfaces.CHAR) }
    |   TK_ID       { $primate = expressiones.NewLLamadoVariable($TK_ID.text) }
    |   lista=listaArreglo { $primate = $lista.primate }
    ;




listaArreglo returns[interfaces.Expresion primate]
    : lista = listaArreglo '[' expresion ']'
        {
            $primate = arreglos.NewArrayAccess($lista.primate,$expresion.primate)
        }
    |   TK_ID
        {       
            $primate = expressiones.NewLLamadoVariable($TK_ID.text)     
        }
;







// ********************************
// GRAMATICA LEXICA
// ********************************
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


TK_IF:      'if';
TK_ELSE:    'else';
TK_ELSE_IF: 'else if';
TK_WHILE:   'while';
TK_LOOP:    'loop';


TK_BREAK:   'break';
TK_POW_I64: 'i64::pow';
TK_POW_F64: 'f64::powf';





// identificadores
TK_NUMBER:  [0-9]+;
TK_DECIMAL: TK_NUMBER+'.'TK_NUMBER+;
TK_CADENA:  '"'~["]*'"';
TK_CARACTER:'\''~["]*'\'';
TK_ID:      ([a-zA-Z_])[a-zA-Z0-9_]*;
TK_COMMET:  ('//' (~[/!] | '//') ~[\r\n]* | '//');


//ignoracon de espacios
SPACES: [ \\\r\n\t]+ -> skip;