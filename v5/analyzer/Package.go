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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	not "github.com/craterdog/go-syntax-notation/v5"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
ModelAnalyzerClassLike declares the set of class constants, constructors and
functions that must be supported by all model-analyzer-class-like classes.
*/
type ModelAnalyzerClassLike interface {
	// Constructor Methods
	Make(
		model mod.ModelLike,
		className string,
	) ModelAnalyzerLike
}

/*
SyntaxAnalyzerClassLike declares the set of class constants, constructors and
functions that must be supported by all syntax-analyzer-class-like classes.
*/
type SyntaxAnalyzerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) SyntaxAnalyzerLike
}

// INSTANCE DECLARATIONS

/*
ModelAnalyzerLike declares the set of aspects and methods that must be
supported by all model-analyzer-like instances.
*/
type ModelAnalyzerLike interface {
	// Principal Methods
	GetClass() ModelAnalyzerClassLike
	GetLegalNotice() string
	GetImportedPackages() abs.CatalogLike[string, string]
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
	GetAspectDeclarations() abs.ListLike[mod.AspectDeclarationLike]
}

/*
SyntaxAnalyzerLike declares the set of aspects and methods that must be
supported by all syntax-analyzer-like instances.
*/
type SyntaxAnalyzerLike interface {
	// Principal Methods
	GetClass() SyntaxAnalyzerClassLike
	GetExpressions() abs.Sequential[abs.AssociationLike[string, string]]
	GetIdentifiers(
		ruleName string,
	) abs.Sequential[not.IdentifierLike]
	GetLegalNotice() string
	GetReferences(
		ruleName string,
	) abs.Sequential[not.ReferenceLike]
	GetRuleNames() abs.Sequential[string]
	GetSyntaxMap() string
	GetSyntaxName() string
	GetTerms(
		ruleName string,
	) abs.Sequential[not.TermLike]
	GetTokenNames() abs.Sequential[string]
	GetVariableType(
		reference not.ReferenceLike,
	) string
	GetVariables(
		ruleName string,
	) abs.Sequential[string]
	HasPlurals() bool
	IsDelimited(
		ruleName string,
	) bool
	IsPlural(
		name string,
	) bool

	// Aspect Interfaces
	not.Methodical
}

// ASPECT DECLARATIONS
