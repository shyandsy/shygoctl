grammar ApiParser;

import ApiLexer;

@lexer::members{
    const COMEMNTS = 88
}

api:            spec*;
spec:           syntaxLit
                |importSpec
                |infoSpec
                |typeSpec
                |serviceSpec
                ;

// syntax
syntaxLit:      {match(p,"syntax")}syntaxToken=ID assign='=' {checkVersion(p)}version=STRING;

// import
importSpec:     importLit|importBlock;
importLit:      {match(p,"import")}importToken=ID importValue ;
importBlock:    {match(p,"import")}importToken=ID '(' importBlockValue+ ')';
importBlockValue:   importValue;
importValue:    {checkImportValue(p)}STRING;

// info
infoSpec:       {match(p,"info")}infoToken=ID lp='(' kvLit+ rp=')';

// type
typeSpec:       typeLit
                |typeBlock;

// eg: type Foo int
typeLit:        {match(p,"type")}typeToken=ID  typeLitBody;
// eg: type (...)
typeBlock:      {match(p,"type")}typeToken=ID lp='(' typeBlockBody* rp=')';
typeLitBody:    typeStruct|typeAlias;
typeBlockBody:  typeBlockStruct|typeBlockAlias;
typeStruct:     {checkKeyword(p)}structName=ID structToken=ID? lbrace='{'  field* rbrace='}';
typeAlias:      {checkKeyword(p)}alias=ID assign='='? dataType;
typeBlockStruct: {checkKeyword(p)}structName=ID structToken=ID? lbrace='{'  field* rbrace='}';
typeBlockAlias: {checkKeyword(p)}alias=ID assign='='? dataType;
field:          {isNormal(p)}? normalField|anonymousFiled ;
normalField:    {checkKeyword(p)}fieldName=ID dataType tag=RAW_STRING?;
anonymousFiled: star='*'? qualifiedType | ID;
dataType:       qualifiedType //用于支持包类型，包括 time='time.Time'
                |ID
                |mapType
                |arrayType
                |inter='interface{}'
                |pointerType
                |typeStruct
                ;
qualifiedType: pkg=ID dot=DOT name=ID;
pointerType:    star='*' dataType;
mapType:        {match(p,"map")}mapToken=ID lbrack='[' {checkKey(p)}key=ID rbrack=']' value=dataType;
arrayType:      lbrack='[' rbrack=']' dataType;

// service
serviceSpec:    atServer? serviceApi;
atServer:       ATSERVER lp='(' kvLit+ rp=')';
serviceApi:     {match(p,"service")}serviceToken=ID serviceName lbrace='{' serviceRoute* rbrace='}';
serviceRoute:   atDoc? atRoles? atFeatures? (atServer|atHandler) route;
atDoc:          ATDOC lp='('? ((kvLit+)|STRING) rp=')'?;
atHandler:      ATHANDLER ID;
atRoles:        ATROLES lp='['? ((kvLit+)|STRING) rp=']'?;
atFeatures:     ATFEATURES lp='['? ((kvLit+)|STRING) rp=']'?;
route:          {checkHTTPMethod(p)}httpMethod=ID path request=body? response=replybody?;
body:           lp='(' (ID)? rp=')';
replybody:      returnToken='returns' lp='(' dataType? rp=')';
// kv
kvLit:          key=ID {checkKeyValue(p)}value=LINE_VALUE;

serviceName:    (ID '-'?)+;
path:           (('/' (pathItem ('-' pathItem)*))|('/:' (pathItem ('-' pathItem)?)))+ | '/';
pathItem:       (ID|LetterOrDigit)+;