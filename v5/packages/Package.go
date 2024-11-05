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

// Class Declarations

/*
GeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all generator-class-like classes.
*/
type GeneratorClassLike interface {
	// Constructor Methods
	Make(
		synthesizer TemplateDriven,
	) GeneratorLike
}

// Instance Declarations

/*
GeneratorLike defines the set of aspects and methods that must be supported by
all generator-like instances.
*/
type GeneratorLike interface {
	// Primary Methods
	GetClass() GeneratorClassLike
	GeneratePackage() string
}

// Aspect Declarations

/*
TemplateDriven defines the set of method signatures that must be supported by
all template-driven generators.
*/
type TemplateDriven interface {
	GenerateLegalNotice() string
	GeneratePackageDeclaration() string
	GenerateModuleImports() string
	GenerateTypeDeclarations() string
	GenerateClassDeclarations() string
	GenerateInstanceDeclarations() string
	GenerateAspectDeclarations() string
}
