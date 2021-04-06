// Code generated from ADL.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // ADL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ADLListener is a complete listener for a parse tree produced by ADLParser.
type ADLListener interface {
	antlr.ParseTreeListener

	// EnterSourceFile is called when entering the sourceFile production.
	EnterSourceFile(c *SourceFileContext)

	// EnterPackageClause is called when entering the packageClause production.
	EnterPackageClause(c *PackageClauseContext)

	// EnterImportDecl is called when entering the importDecl production.
	EnterImportDecl(c *ImportDeclContext)

	// EnterImportSpec is called when entering the importSpec production.
	EnterImportSpec(c *ImportSpecContext)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterImportPath is called when entering the importPath production.
	EnterImportPath(c *ImportPathContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterTypeLit is called when entering the typeLit production.
	EnterTypeLit(c *TypeLitContext)

	// EnterStructType is called when entering the structType production.
	EnterStructType(c *StructTypeContext)

	// EnterFieldDecl is called when entering the fieldDecl production.
	EnterFieldDecl(c *FieldDeclContext)

	// EnterInterfaceType is called when entering the interfaceType production.
	EnterInterfaceType(c *InterfaceTypeContext)

	// EnterMethodSpec is called when entering the methodSpec production.
	EnterMethodSpec(c *MethodSpecContext)

	// EnterResult is called when entering the result production.
	EnterResult(c *ResultContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParameterDecl is called when entering the parameterDecl production.
	EnterParameterDecl(c *ParameterDeclContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterTypeSpec is called when entering the typeSpec production.
	EnterTypeSpec(c *TypeSpecContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterQualifiedIdent is called when entering the qualifiedIdent production.
	EnterQualifiedIdent(c *QualifiedIdentContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// EnterEos2 is called when entering the eos2 production.
	EnterEos2(c *Eos2Context)

	// ExitSourceFile is called when exiting the sourceFile production.
	ExitSourceFile(c *SourceFileContext)

	// ExitPackageClause is called when exiting the packageClause production.
	ExitPackageClause(c *PackageClauseContext)

	// ExitImportDecl is called when exiting the importDecl production.
	ExitImportDecl(c *ImportDeclContext)

	// ExitImportSpec is called when exiting the importSpec production.
	ExitImportSpec(c *ImportSpecContext)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitImportPath is called when exiting the importPath production.
	ExitImportPath(c *ImportPathContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitTypeLit is called when exiting the typeLit production.
	ExitTypeLit(c *TypeLitContext)

	// ExitStructType is called when exiting the structType production.
	ExitStructType(c *StructTypeContext)

	// ExitFieldDecl is called when exiting the fieldDecl production.
	ExitFieldDecl(c *FieldDeclContext)

	// ExitInterfaceType is called when exiting the interfaceType production.
	ExitInterfaceType(c *InterfaceTypeContext)

	// ExitMethodSpec is called when exiting the methodSpec production.
	ExitMethodSpec(c *MethodSpecContext)

	// ExitResult is called when exiting the result production.
	ExitResult(c *ResultContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParameterDecl is called when exiting the parameterDecl production.
	ExitParameterDecl(c *ParameterDeclContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitTypeSpec is called when exiting the typeSpec production.
	ExitTypeSpec(c *TypeSpecContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitQualifiedIdent is called when exiting the qualifiedIdent production.
	ExitQualifiedIdent(c *QualifiedIdentContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)

	// ExitEos2 is called when exiting the eos2 production.
	ExitEos2(c *Eos2Context)
}
