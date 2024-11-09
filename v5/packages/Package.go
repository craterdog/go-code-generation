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
Package "packages" provides a template-based code generator that can generate
the class model packages based on any language defined using Crater Dog Syntax
Notation™ (CDSN).

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-code-generation/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package packages

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	not "github.com/craterdog/go-syntax-notation/v5"
)

// Type Declarations

// Functional Declarations

// Class Declarations

/*
AnalyzerClassLike defines the set of class constants, constructors and
functions that must be supported by all analyzer-class-like classes.
*/
type AnalyzerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) AnalyzerLike
}

/*
AstSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all ast-synthesizer-class-like classes.
*/
type AstSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) AstSynthesizerLike
}

/*
ExampleSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all example-synthesizer-class-like classes.
*/
type ExampleSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) ExampleSynthesizerLike
}

/*
GeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all generator-class-like classes.
*/
type GeneratorClassLike interface {
	// Constructor Methods
	Make() GeneratorLike
}

/*
GrammarSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all grammar-synthesizer-class-like classes.
*/
type GrammarSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) GrammarSynthesizerLike
}

// Instance Declarations

/*
AnalyzerLike defines the set of aspects and methods that must be supported by
all analyzer-like instances.
*/
type AnalyzerLike interface {
	// Primary Methods
	GetClass() AnalyzerClassLike
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

/*
AstSynthesizerLike defines the set of aspects and methods that must be supported
by all ast-synthesizer-like instances.
*/
type AstSynthesizerLike interface {
	// Primary Methods
	GetClass() AstSynthesizerClassLike

	// Aspect Interfaces
	TemplateDriven
}

/*
ExampleSynthesizerLike defines the set of aspects and methods that must be
supported by all example-synthesizer-like instances.
*/
type ExampleSynthesizerLike interface {
	// Primary Methods
	GetClass() ExampleSynthesizerClassLike

	// Aspect Interfaces
	TemplateDriven
}

/*
GeneratorLike defines the set of aspects and methods that must be supported by
all generator-like instances.
*/
type GeneratorLike interface {
	// Primary Methods
	GetClass() GeneratorClassLike
	GeneratePackage(
		moduleName string,
		wikiPath string,
		packageName string,
		synthesizer TemplateDriven,
	) string
}

/*
GrammarSynthesizerLike defines the set of aspects and methods that must be
supported by all grammar-synthesizer-like instances.
*/
type GrammarSynthesizerLike interface {
	// Primary Methods
	GetClass() GrammarSynthesizerClassLike

	// Aspect Interfaces
	TemplateDriven
}

// Aspect Declarations

/*
TemplateDriven defines the set of method signatures that must be supported by
all template-driven synthesizers.
*/
type TemplateDriven interface {
	CreateLegalNotice() string
	CreatePackageDescription() string
	CreateModuleImports() string
	CreateTypeDeclarations() string
	CreateFunctionalDeclarations() string
	CreateClassDeclarations() string
	CreateInstanceDeclarations() string
	CreateAspectDeclarations() string
}
