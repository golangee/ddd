// Code generated from ADL.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // ADL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ADLParser.
type ADLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ADLParser#sourceFile.
	VisitSourceFile(ctx *SourceFileContext) interface{}

	// Visit a parse tree produced by ADLParser#packageClause.
	VisitPackageClause(ctx *PackageClauseContext) interface{}

	// Visit a parse tree produced by ADLParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by ADLParser#importSpec.
	VisitImportSpec(ctx *ImportSpecContext) interface{}

	// Visit a parse tree produced by ADLParser#string_.
	VisitString_(ctx *String_Context) interface{}

	// Visit a parse tree produced by ADLParser#importPath.
	VisitImportPath(ctx *ImportPathContext) interface{}

	// Visit a parse tree produced by ADLParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by ADLParser#typeDecl.
	VisitTypeDecl(ctx *TypeDeclContext) interface{}

	// Visit a parse tree produced by ADLParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by ADLParser#typeLit.
	VisitTypeLit(ctx *TypeLitContext) interface{}

	// Visit a parse tree produced by ADLParser#structType.
	VisitStructType(ctx *StructTypeContext) interface{}

	// Visit a parse tree produced by ADLParser#fieldDecl.
	VisitFieldDecl(ctx *FieldDeclContext) interface{}

	// Visit a parse tree produced by ADLParser#interfaceType.
	VisitInterfaceType(ctx *InterfaceTypeContext) interface{}

	// Visit a parse tree produced by ADLParser#methodSpec.
	VisitMethodSpec(ctx *MethodSpecContext) interface{}

	// Visit a parse tree produced by ADLParser#result.
	VisitResult(ctx *ResultContext) interface{}

	// Visit a parse tree produced by ADLParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by ADLParser#parameterDecl.
	VisitParameterDecl(ctx *ParameterDeclContext) interface{}

	// Visit a parse tree produced by ADLParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by ADLParser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by ADLParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by ADLParser#qualifiedIdent.
	VisitQualifiedIdent(ctx *QualifiedIdentContext) interface{}

	// Visit a parse tree produced by ADLParser#eos.
	VisitEos(ctx *EosContext) interface{}

	// Visit a parse tree produced by ADLParser#eos2.
	VisitEos2(ctx *Eos2Context) interface{}
}
