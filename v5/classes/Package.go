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

// Class Declarations

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
GeneratorLike defines the set of aspects and methods that must be supported by
all generator-like instances.
*/
type GeneratorLike interface {
	// Primary Methods
	GetClass() GeneratorClassLike
	GenerateClass(
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
	CreatePackageDeclaration() string
	CreateModuleImports() string
	CreateAccessFunction() string
	CreateConstructorMethods() string
	CreateFunctionMethods() string
	CreatePrimaryMethods() string
	CreateAttributeMethods() string
	CreateAspectMethods() string
	CreatePrivateMethods() string
	CreateInstanceStructure() string
	CreateClassStructure() string
	CreateClassReference() string
}
