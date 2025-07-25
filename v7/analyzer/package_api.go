/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
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
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package analyzer

import (
	mod "github.com/craterdog/go-class-model/v7"
	fra "github.com/craterdog/go-component-framework/v7"
	not "github.com/craterdog/go-syntax-notation/v7"
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
	GetConstants() fra.CatalogLike[string, string]
	GetAttributes() fra.CatalogLike[string, string]
	GetConstructorMethods() fra.ListLike[mod.ConstructorMethodLike]
	GetConstantMethods() fra.ListLike[mod.ConstantMethodLike]
	GetFunctionMethods() fra.ListLike[mod.FunctionMethodLike]
	GetPrincipalMethods() fra.ListLike[mod.PrincipalMethodLike]
	GetAttributeMethods() fra.ListLike[mod.AttributeMethodLike]
	GetAspectInterfaces() fra.ListLike[mod.AspectInterfaceLike]
}

/*
PackageAnalyzerLike declares the set of aspects and methods that must be
supported by all package-analyzer-like instances.
*/
type PackageAnalyzerLike interface {
	// Principal Methods
	GetClass() PackageAnalyzerClassLike
	GetLegalNotice() string
	GetPackageName() string
	GetImportedPackages() fra.CatalogLike[string, string]
	GetTypeDeclarations() fra.ListLike[mod.TypeDeclarationLike]
	GetEnumeratedValues() fra.ListLike[string]
	GetFunctionalDeclarations() fra.ListLike[mod.FunctionalDeclarationLike]
	GetClassDeclarations() fra.ListLike[mod.ClassDeclarationLike]
	GetInstanceDeclarations() fra.ListLike[mod.InstanceDeclarationLike]
	GetAspectDeclarations() fra.ListLike[mod.AspectDeclarationLike]

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
	GetRules() fra.SetLike[string]
	GetExpressions() fra.SetLike[string]
	GetLiteralValues(
		ruleName string,
	) fra.ListLike[not.LiteralValueLike]
	GetExpressionNames(
		ruleName string,
	) fra.ListLike[not.ExpressionNameLike]
	GetRuleNames(
		ruleName string,
	) fra.ListLike[not.RuleNameLike]
	GetRuleTerms(
		ruleName string,
	) fra.ListLike[not.RuleTermLike]
	GetVariables(
		ruleName string,
	) fra.ListLike[string]
	GetVariableType(
		component not.ComponentLike,
	) string
	GetDefinitions() fra.CatalogLike[string, not.DefinitionLike]
	GetPatterns() fra.CatalogLike[string, string]
	GetSyntaxMap() string

	// Aspect Interfaces
	not.Methodical
}

// ASPECT DECLARATIONS
