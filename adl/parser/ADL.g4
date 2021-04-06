grammar ADL;

/*
 * Parser Rules
 */

sourceFile
    : packageClause eos2 (importDecl eos)* ((declaration) eos)* EOF
    ;

    packageClause
    : 'package' IDENTIFIER
    ;

importDecl
    : 'import' (importSpec | '(' (importSpec eos)* ')')
    ;

importSpec
    : ('.' | IDENTIFIER)? importPath
    ;

string_
    : INTERPRETED_STRING_LIT
    ;

importPath
    : string_
    ;

declaration
    : typeDecl
    ;

typeDecl
    : 'type' (typeSpec | '(' (typeSpec eos)* ')')
    ;

type_
    : typeName
    | typeLit
    | '(' type_ ')'
    ;

typeLit
    : structType
    | interfaceType
    ;


structType
    : 'struct' '{' (fieldDecl eos)* '}'
    ;

fieldDecl
    : ({p.noTerminatorBetween(2)}? identifierList type_) string_?
    ;


interfaceType
    : 'interface' '{' (methodSpec eos)* '}'
    ;

methodSpec
    : {p.noTerminatorAfterParams(2)}? IDENTIFIER parameters result
    | IDENTIFIER parameters
    ;

result
    : parameters
    | type_
    ;

parameters
    : '(' (parameterDecl (COMMA parameterDecl)* COMMA?)? ')'
    ;

parameterDecl
    : identifierList? '...'? type_
    ;

identifierList
    : IDENTIFIER (',' IDENTIFIER)*
    ;

typeSpec
    : IDENTIFIER type_
    ;

typeName
    : IDENTIFIER
    | qualifiedIdent
    ;

qualifiedIdent
    : IDENTIFIER '.' IDENTIFIER
    ;

eos
    : ';'
    | EOF
    | {p.lineTerminatorAhead()}?
    | {p.checkPreviousTokenText("}")}?
    ;

eos2
    : ';'
    | {p.lineTerminatorAhead()}?
    | {p.checkPreviousTokenText("}")}?
    ;

/*
 * Lexer Rules
 */

FUNC                   : 'func';
INTERFACE              : 'interface';
STRUCT                 : 'struct';
PACKAGE                : 'package';

NIL_LIT                : 'nil';

IDENTIFIER             : [_\p{L}] [_\p{L}\p{Nd}]*;
IMPORT                 : 'import';

TYPE                   : 'type';

STAR                   : '*';
AMPERSAND              : '&';

WS                     : [ \t]+             -> channel(HIDDEN);
COMMENT                : '/*' .*? '*/'      -> channel(HIDDEN);
TERMINATOR             : [\r\n]+            -> channel(HIDDEN);
LINE_COMMENT           : '//' ~[\r\n]*      -> channel(HIDDEN);

INTERPRETED_STRING_LIT : '"' (~["\\] | ESCAPED_VALUE)*  '"';

OCTAL_LIT              : '0' OCTAL_DIGIT*;
HEX_LIT                : '0' [xX] HEX_DIGIT+;

COMMA                  : ',';

L_PAREN                : '(';
R_PAREN                : ')';

fragment ESCAPED_VALUE
    : '\\' ('u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | [abfnrtv\\'"]
           | OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT
           | 'x' HEX_DIGIT HEX_DIGIT)
    ;

fragment OCTAL_DIGIT
    : [0-7]
    ;

fragment HEX_DIGIT
    : [0-9a-fA-F]
    ;