// Code generated from ADL.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // ADL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseADLListener is a complete listener for a parse tree produced by ADLParser.
type BaseADLListener struct{}

var _ ADLListener = &BaseADLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseADLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseADLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseADLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseADLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSourceFile is called when production sourceFile is entered.
func (s *BaseADLListener) EnterSourceFile(ctx *SourceFileContext) {}

// ExitSourceFile is called when production sourceFile is exited.
func (s *BaseADLListener) ExitSourceFile(ctx *SourceFileContext) {}

// EnterPackageClause is called when production packageClause is entered.
func (s *BaseADLListener) EnterPackageClause(ctx *PackageClauseContext) {}

// ExitPackageClause is called when production packageClause is exited.
func (s *BaseADLListener) ExitPackageClause(ctx *PackageClauseContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseADLListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseADLListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *BaseADLListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *BaseADLListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BaseADLListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BaseADLListener) ExitString_(ctx *String_Context) {}

// EnterImportPath is called when production importPath is entered.
func (s *BaseADLListener) EnterImportPath(ctx *ImportPathContext) {}

// ExitImportPath is called when production importPath is exited.
func (s *BaseADLListener) ExitImportPath(ctx *ImportPathContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseADLListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseADLListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseADLListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseADLListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseADLListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseADLListener) ExitType_(ctx *Type_Context) {}

// EnterTypeLit is called when production typeLit is entered.
func (s *BaseADLListener) EnterTypeLit(ctx *TypeLitContext) {}

// ExitTypeLit is called when production typeLit is exited.
func (s *BaseADLListener) ExitTypeLit(ctx *TypeLitContext) {}

// EnterStructType is called when production structType is entered.
func (s *BaseADLListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *BaseADLListener) ExitStructType(ctx *StructTypeContext) {}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *BaseADLListener) EnterFieldDecl(ctx *FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *BaseADLListener) ExitFieldDecl(ctx *FieldDeclContext) {}

// EnterInterfaceType is called when production interfaceType is entered.
func (s *BaseADLListener) EnterInterfaceType(ctx *InterfaceTypeContext) {}

// ExitInterfaceType is called when production interfaceType is exited.
func (s *BaseADLListener) ExitInterfaceType(ctx *InterfaceTypeContext) {}

// EnterMethodSpec is called when production methodSpec is entered.
func (s *BaseADLListener) EnterMethodSpec(ctx *MethodSpecContext) {}

// ExitMethodSpec is called when production methodSpec is exited.
func (s *BaseADLListener) ExitMethodSpec(ctx *MethodSpecContext) {}

// EnterResult is called when production result is entered.
func (s *BaseADLListener) EnterResult(ctx *ResultContext) {}

// ExitResult is called when production result is exited.
func (s *BaseADLListener) ExitResult(ctx *ResultContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseADLListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseADLListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameterDecl is called when production parameterDecl is entered.
func (s *BaseADLListener) EnterParameterDecl(ctx *ParameterDeclContext) {}

// ExitParameterDecl is called when production parameterDecl is exited.
func (s *BaseADLListener) ExitParameterDecl(ctx *ParameterDeclContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseADLListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseADLListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BaseADLListener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BaseADLListener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseADLListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseADLListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterQualifiedIdent is called when production qualifiedIdent is entered.
func (s *BaseADLListener) EnterQualifiedIdent(ctx *QualifiedIdentContext) {}

// ExitQualifiedIdent is called when production qualifiedIdent is exited.
func (s *BaseADLListener) ExitQualifiedIdent(ctx *QualifiedIdentContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseADLListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseADLListener) ExitEos(ctx *EosContext) {}

// EnterEos2 is called when production eos2 is entered.
func (s *BaseADLListener) EnterEos2(ctx *Eos2Context) {}

// ExitEos2 is called when production eos2 is exited.
func (s *BaseADLListener) ExitEos2(ctx *Eos2Context) {}
