/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
Package "analyzer" provides classes that can analyze the abstract syntax trees
(ASTs) for different languages and pull out the parts that are needed by each
template-driven code synthesizer.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-code-generation/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package analyzer

import (
	mod "github.com/craterdog/go-class-model/v5"
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	not "github.com/craterdog/go-syntax-notation/v5"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
ClassAnalyzerClassLike declares the set of class constants, constructors and
functions that must be supported by all class-analyzer-class-like classes.
*/
type ClassAnalyzerClassLike interface {
	// Constructor Methods
	ClassAnalyzer(
		model mod.ModelLike,
		className string,
	) ClassAnalyzerLike
}

/*
PackageAnalyzerClassLike declares the set of class constants, constructors and
functions that must be supported by all package-analyzer-class-like classes.
*/
type PackageAnalyzerClassLike interface {
	// Constructor Methods
	PackageAnalyzer(
		model mod.ModelLike,
	) PackageAnalyzerLike
}

/*
SyntaxAnalyzerClassLike declares the set of class constants, constructors and
functions that must be supported by all syntax-analyzer-class-like classes.
*/
type SyntaxAnalyzerClassLike interface {
	// Constructor Methods
	SyntaxAnalyzer(
		syntax not.SyntaxLike,
	) SyntaxAnalyzerLike
}

// INSTANCE DECLARATIONS

/*
ClassAnalyzerLike declares the set of aspects and methods that must be
supported by all class-analyzer-like instances.
*/
type ClassAnalyzerLike interface {
	// Principal Methods
	GetClass() ClassAnalyzerClassLike
	IsGeneric() bool
	GetTypeConstraints() string
	GetTypeArguments() string
	IsIntrinsic() bool
	GetIntrinsicType() mod.AbstractionLike
	GetConstants() abs.CatalogLike[string, string]
	GetAttributes() abs.CatalogLike[string, string]
	GetConstructorMethods() abs.ListLike[mod.ConstructorMethodLike]
	GetConstantMethods() abs.ListLike[mod.ConstantMethodLike]
	GetFunctionMethods() abs.ListLike[mod.FunctionMethodLike]
	GetPrincipalMethods() abs.ListLike[mod.PrincipalMethodLike]
	GetAttributeMethods() abs.ListLike[mod.AttributeMethodLike]
	GetAspectInterfaces() abs.ListLike[mod.AspectInterfaceLike]
}

/*
PackageAnalyzerLike declares the set of aspects and methods that must be
supported by all package-analyzer-like instances.
*/
type PackageAnalyzerLike interface {
	// Principal Methods
	GetClass() PackageAnalyzerClassLike
	GetLegalNotice() string
	GetImportedPackages() abs.CatalogLike[string, string]
	GetTypeDeclarations() abs.ListLike[mod.TypeDeclarationLike]
	GetEnumeratedValues() abs.ListLike[string]
	GetFunctionalDeclarations() abs.ListLike[mod.FunctionalDeclarationLike]
	GetClassDeclarations() abs.ListLike[mod.ClassDeclarationLike]
	GetInstanceDeclarations() abs.ListLike[mod.InstanceDeclarationLike]
	GetAspectDeclarations() abs.ListLike[mod.AspectDeclarationLike]

	// Aspect Interfaces
	mod.Methodical
}

/*
SyntaxAnalyzerLike declares the set of aspects and methods that must be
supported by all syntax-analyzer-like instances.
*/
type SyntaxAnalyzerLike interface {
	// Principal Methods
	GetClass() SyntaxAnalyzerClassLike
	GetLegalNotice() string
	GetSyntaxName() string
	HasPlurals() bool
	GetRuleNames() abs.SetLike[string]
	GetTokenNames() abs.SetLike[string]
	IsDelimited(
		ruleName string,
	) bool
	GetTerms(
		ruleName string,
	) abs.ListLike[not.TermLike]
	GetReferences(
		ruleName string,
	) abs.ListLike[not.ReferenceLike]
	GetVariables(
		ruleName string,
	) abs.ListLike[string]
	GetVariableType(
		reference not.ReferenceLike,
	) string
	GetIdentifiers(
		ruleName string,
	) abs.ListLike[not.IdentifierLike]
	IsPlural(
		identifierName string,
	) bool
	GetExpressions() abs.CatalogLike[string, string]
	GetSyntaxMap() string

	// Aspect Interfaces
	not.Methodical
}

// ASPECT DECLARATIONS
