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
Package "classes" provides a template-based code generator that can generate
a concrete class definition based on the class model declared in a Package.go
file.

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
package classes

import (
	mod "github.com/craterdog/go-class-model/v5"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
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
		model mod.ModelLike,
		className string,
	) AnalyzerLike
}

/*
ClassSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all example-synthesizer-class-like classes.
*/
type ClassSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		model mod.ModelLike,
		className string,
	) ClassSynthesizerLike
}

/*
GeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all generator-class-like classes.
*/
type GeneratorClassLike interface {
	// Constructor Methods
	Make() GeneratorLike
}

// Instance Declarations

/*
AnalyzerLike defines the set of aspects and methods that must be supported by
all analyzer-like instances.
*/
type AnalyzerLike interface {
	// Primary Methods
	GetClass() AnalyzerClassLike
	GetLegalNotice() string
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
	GetPrimaryMethods() abs.ListLike[mod.PrimaryMethodLike]
	GetAttributeMethods() abs.ListLike[mod.AttributeMethodLike]
	GetAspectInterfaces() abs.ListLike[mod.AspectInterfaceLike]
	GetAspectDeclarations() abs.ListLike[mod.AspectDeclarationLike]
	GetModuleImports() mod.ModuleImportsLike
}

/*
ClassSynthesizerLike defines the set of aspects and methods that must be
supported by all example-synthesizer-like instances.
*/
type ClassSynthesizerLike interface {
	// Primary Methods
	GetClass() ClassSynthesizerClassLike

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
	GenerateClass(
		moduleName string,
		wikiPath string,
		packageName string,
		className string,
		synthesizer TemplateDriven,
	) string
}

// Aspect Declarations

/*
TemplateDriven defines the set of method signatures that must be supported by
all template-driven synthesizers.
*/
type TemplateDriven interface {
	CreateLegalNotice() string
	CreateAccessFunction() string
	CreateConstructorMethods() string
	CreateConstantMethods() string
	CreateFunctionMethods() string
	CreatePrimaryMethods() string
	CreateAttributeMethods() string
	CreateAspectMethods() string
	CreatePrivateMethods() string
	CreateInstanceStructure() string
	CreateClassStructure() string
	CreateClassReference() string
	PerformGlobalUpdates(
		source string,
	) string
}
