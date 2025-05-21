// Code generated from ApiParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package api // ApiParser
import "github.com/antlr4-go/antlr/v4"

// BaseApiParserListener is a complete listener for a parse tree produced by ApiParserParser.
type BaseApiParserListener struct{}

var _ ApiParserListener = &BaseApiParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseApiParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseApiParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseApiParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseApiParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterApi is called when production api is entered.
func (s *BaseApiParserListener) EnterApi(ctx *ApiContext) {}

// ExitApi is called when production api is exited.
func (s *BaseApiParserListener) ExitApi(ctx *ApiContext) {}

// EnterSpec is called when production spec is entered.
func (s *BaseApiParserListener) EnterSpec(ctx *SpecContext) {}

// ExitSpec is called when production spec is exited.
func (s *BaseApiParserListener) ExitSpec(ctx *SpecContext) {}

// EnterSyntaxLit is called when production syntaxLit is entered.
func (s *BaseApiParserListener) EnterSyntaxLit(ctx *SyntaxLitContext) {}

// ExitSyntaxLit is called when production syntaxLit is exited.
func (s *BaseApiParserListener) ExitSyntaxLit(ctx *SyntaxLitContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *BaseApiParserListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *BaseApiParserListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterImportLit is called when production importLit is entered.
func (s *BaseApiParserListener) EnterImportLit(ctx *ImportLitContext) {}

// ExitImportLit is called when production importLit is exited.
func (s *BaseApiParserListener) ExitImportLit(ctx *ImportLitContext) {}

// EnterImportBlock is called when production importBlock is entered.
func (s *BaseApiParserListener) EnterImportBlock(ctx *ImportBlockContext) {}

// ExitImportBlock is called when production importBlock is exited.
func (s *BaseApiParserListener) ExitImportBlock(ctx *ImportBlockContext) {}

// EnterImportBlockValue is called when production importBlockValue is entered.
func (s *BaseApiParserListener) EnterImportBlockValue(ctx *ImportBlockValueContext) {}

// ExitImportBlockValue is called when production importBlockValue is exited.
func (s *BaseApiParserListener) ExitImportBlockValue(ctx *ImportBlockValueContext) {}

// EnterImportValue is called when production importValue is entered.
func (s *BaseApiParserListener) EnterImportValue(ctx *ImportValueContext) {}

// ExitImportValue is called when production importValue is exited.
func (s *BaseApiParserListener) ExitImportValue(ctx *ImportValueContext) {}

// EnterInfoSpec is called when production infoSpec is entered.
func (s *BaseApiParserListener) EnterInfoSpec(ctx *InfoSpecContext) {}

// ExitInfoSpec is called when production infoSpec is exited.
func (s *BaseApiParserListener) ExitInfoSpec(ctx *InfoSpecContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BaseApiParserListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BaseApiParserListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterTypeLit is called when production typeLit is entered.
func (s *BaseApiParserListener) EnterTypeLit(ctx *TypeLitContext) {}

// ExitTypeLit is called when production typeLit is exited.
func (s *BaseApiParserListener) ExitTypeLit(ctx *TypeLitContext) {}

// EnterTypeBlock is called when production typeBlock is entered.
func (s *BaseApiParserListener) EnterTypeBlock(ctx *TypeBlockContext) {}

// ExitTypeBlock is called when production typeBlock is exited.
func (s *BaseApiParserListener) ExitTypeBlock(ctx *TypeBlockContext) {}

// EnterTypeLitBody is called when production typeLitBody is entered.
func (s *BaseApiParserListener) EnterTypeLitBody(ctx *TypeLitBodyContext) {}

// ExitTypeLitBody is called when production typeLitBody is exited.
func (s *BaseApiParserListener) ExitTypeLitBody(ctx *TypeLitBodyContext) {}

// EnterTypeBlockBody is called when production typeBlockBody is entered.
func (s *BaseApiParserListener) EnterTypeBlockBody(ctx *TypeBlockBodyContext) {}

// ExitTypeBlockBody is called when production typeBlockBody is exited.
func (s *BaseApiParserListener) ExitTypeBlockBody(ctx *TypeBlockBodyContext) {}

// EnterTypeStruct is called when production typeStruct is entered.
func (s *BaseApiParserListener) EnterTypeStruct(ctx *TypeStructContext) {}

// ExitTypeStruct is called when production typeStruct is exited.
func (s *BaseApiParserListener) ExitTypeStruct(ctx *TypeStructContext) {}

// EnterTypeAlias is called when production typeAlias is entered.
func (s *BaseApiParserListener) EnterTypeAlias(ctx *TypeAliasContext) {}

// ExitTypeAlias is called when production typeAlias is exited.
func (s *BaseApiParserListener) ExitTypeAlias(ctx *TypeAliasContext) {}

// EnterTypeBlockStruct is called when production typeBlockStruct is entered.
func (s *BaseApiParserListener) EnterTypeBlockStruct(ctx *TypeBlockStructContext) {}

// ExitTypeBlockStruct is called when production typeBlockStruct is exited.
func (s *BaseApiParserListener) ExitTypeBlockStruct(ctx *TypeBlockStructContext) {}

// EnterTypeBlockAlias is called when production typeBlockAlias is entered.
func (s *BaseApiParserListener) EnterTypeBlockAlias(ctx *TypeBlockAliasContext) {}

// ExitTypeBlockAlias is called when production typeBlockAlias is exited.
func (s *BaseApiParserListener) ExitTypeBlockAlias(ctx *TypeBlockAliasContext) {}

// EnterField is called when production field is entered.
func (s *BaseApiParserListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseApiParserListener) ExitField(ctx *FieldContext) {}

// EnterNormalField is called when production normalField is entered.
func (s *BaseApiParserListener) EnterNormalField(ctx *NormalFieldContext) {}

// ExitNormalField is called when production normalField is exited.
func (s *BaseApiParserListener) ExitNormalField(ctx *NormalFieldContext) {}

// EnterAnonymousFiled is called when production anonymousFiled is entered.
func (s *BaseApiParserListener) EnterAnonymousFiled(ctx *AnonymousFiledContext) {}

// ExitAnonymousFiled is called when production anonymousFiled is exited.
func (s *BaseApiParserListener) ExitAnonymousFiled(ctx *AnonymousFiledContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseApiParserListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseApiParserListener) ExitDataType(ctx *DataTypeContext) {}

// EnterQualifiedType is called when production qualifiedType is entered.
func (s *BaseApiParserListener) EnterQualifiedType(ctx *QualifiedTypeContext) {}

// ExitQualifiedType is called when production qualifiedType is exited.
func (s *BaseApiParserListener) ExitQualifiedType(ctx *QualifiedTypeContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *BaseApiParserListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *BaseApiParserListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseApiParserListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseApiParserListener) ExitMapType(ctx *MapTypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseApiParserListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseApiParserListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterServiceSpec is called when production serviceSpec is entered.
func (s *BaseApiParserListener) EnterServiceSpec(ctx *ServiceSpecContext) {}

// ExitServiceSpec is called when production serviceSpec is exited.
func (s *BaseApiParserListener) ExitServiceSpec(ctx *ServiceSpecContext) {}

// EnterAtServer is called when production atServer is entered.
func (s *BaseApiParserListener) EnterAtServer(ctx *AtServerContext) {}

// ExitAtServer is called when production atServer is exited.
func (s *BaseApiParserListener) ExitAtServer(ctx *AtServerContext) {}

// EnterServiceApi is called when production serviceApi is entered.
func (s *BaseApiParserListener) EnterServiceApi(ctx *ServiceApiContext) {}

// ExitServiceApi is called when production serviceApi is exited.
func (s *BaseApiParserListener) ExitServiceApi(ctx *ServiceApiContext) {}

// EnterServiceRoute is called when production serviceRoute is entered.
func (s *BaseApiParserListener) EnterServiceRoute(ctx *ServiceRouteContext) {}

// ExitServiceRoute is called when production serviceRoute is exited.
func (s *BaseApiParserListener) ExitServiceRoute(ctx *ServiceRouteContext) {}

// EnterAtDoc is called when production atDoc is entered.
func (s *BaseApiParserListener) EnterAtDoc(ctx *AtDocContext) {}

// ExitAtDoc is called when production atDoc is exited.
func (s *BaseApiParserListener) ExitAtDoc(ctx *AtDocContext) {}

// EnterAtHandler is called when production atHandler is entered.
func (s *BaseApiParserListener) EnterAtHandler(ctx *AtHandlerContext) {}

// ExitAtHandler is called when production atHandler is exited.
func (s *BaseApiParserListener) ExitAtHandler(ctx *AtHandlerContext) {}

// EnterAtRoles is called when production atRoles is entered.
func (s *BaseApiParserListener) EnterAtRoles(ctx *AtRolesContext) {}

// ExitAtRoles is called when production atRoles is exited.
func (s *BaseApiParserListener) ExitAtRoles(ctx *AtRolesContext) {}

// EnterAtFeatures is called when production atFeatures is entered.
func (s *BaseApiParserListener) EnterAtFeatures(ctx *AtFeaturesContext) {}

// ExitAtFeatures is called when production atFeatures is exited.
func (s *BaseApiParserListener) ExitAtFeatures(ctx *AtFeaturesContext) {}

// EnterRoute is called when production route is entered.
func (s *BaseApiParserListener) EnterRoute(ctx *RouteContext) {}

// ExitRoute is called when production route is exited.
func (s *BaseApiParserListener) ExitRoute(ctx *RouteContext) {}

// EnterBody is called when production body is entered.
func (s *BaseApiParserListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseApiParserListener) ExitBody(ctx *BodyContext) {}

// EnterReplybody is called when production replybody is entered.
func (s *BaseApiParserListener) EnterReplybody(ctx *ReplybodyContext) {}

// ExitReplybody is called when production replybody is exited.
func (s *BaseApiParserListener) ExitReplybody(ctx *ReplybodyContext) {}

// EnterKvLit is called when production kvLit is entered.
func (s *BaseApiParserListener) EnterKvLit(ctx *KvLitContext) {}

// ExitKvLit is called when production kvLit is exited.
func (s *BaseApiParserListener) ExitKvLit(ctx *KvLitContext) {}

// EnterServiceName is called when production serviceName is entered.
func (s *BaseApiParserListener) EnterServiceName(ctx *ServiceNameContext) {}

// ExitServiceName is called when production serviceName is exited.
func (s *BaseApiParserListener) ExitServiceName(ctx *ServiceNameContext) {}

// EnterPath is called when production path is entered.
func (s *BaseApiParserListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseApiParserListener) ExitPath(ctx *PathContext) {}

// EnterPathItem is called when production pathItem is entered.
func (s *BaseApiParserListener) EnterPathItem(ctx *PathItemContext) {}

// ExitPathItem is called when production pathItem is exited.
func (s *BaseApiParserListener) ExitPathItem(ctx *PathItemContext) {}
