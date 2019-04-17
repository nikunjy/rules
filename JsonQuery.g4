grammar JsonQuery;

query
   : NOT? SP? '(' query ')'                                                                         #parenExp
   | query SP LOGICAL_OPERATOR SP query                                                             #logicalExp
   | attrPath SP 'pr'                                                                               #presentExp
   | attrPath SP op=( 'eq' | 'ne' | 'gt' | 'lt' | 'ge' | 'le' | 'co' | 'sw' | 'ew' ) SP value       #compareExp

   ;

NOT
   : 'not'
   ;

LOGICAL_OPERATOR
   : 'and' | 'or'
   ;

BOOLEAN
   : 'true' | 'false'
   ;

NULL
   : 'null'
   ;

EQ : 'eq' ;
NE : 'ne' ;
GT : 'gt' ;
LT : 'lt' ;
GE : 'ge' ;
LE : 'le' ;
CO : 'co' ;
SW : 'sw' ;
EW : 'ew' ;

attrPath
   : ATTRNAME subAttr?
   ;

subAttr
   : '.' attrPath
   ;

ATTRNAME
   : ALPHA ATTR_NAME_CHAR* ;

fragment ATTR_NAME_CHAR
   : '-' | '_' | ':' | DIGIT | ALPHA
   ;

fragment DIGIT
   : ('0'..'9')
   ;

fragment ALPHA
   : ( 'A'..'Z' | 'a'..'z' )
   ;

value
   : BOOLEAN           #boolean
   | NULL              #null
   | STRING            #string
   | DOUBLE            #double
   | '-'? INT EXP?     #long
   ;

STRING
   : '"' (ESC | ~ ["\\])* '"'
   ;

fragment ESC
   : '\\' (["\\/bfnrt] | UNICODE)
   ;

fragment UNICODE
   : 'u' HEX HEX HEX HEX
   ;

fragment HEX
   : [0-9a-fA-F]
   ;

DOUBLE
   : '-'? INT '.' [0-9] + EXP?
   ;

// INT no leading zeros.
INT
   : '0' | [1-9] [0-9]*
   ;

// EXP we use "\-" since "-" means "range" inside [...]
EXP
   : [Ee] [+\-]? INT
   ;

NEWLINE
   : '\n' ;

SP
   : ' ' NEWLINE*
   ;
