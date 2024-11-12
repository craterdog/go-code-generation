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
Package "generator" provides template-based code generators that can generate
Go files that conform to the Crater Dog Technologies™ Go Coding Conventions
located here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package generator

// Class Declarations

/*
ClassGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all class-generator-class-like classes.
*/
type ClassGeneratorClassLike interface {
	// Constructor Methods
	Make() ClassGeneratorLike
}

/*
ModuleGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all module-generator-class-like classes.
*/
type ModuleGeneratorClassLike interface {
	// Constructor Methods
	Make() ModuleGeneratorLike
}

/*
PackageGeneratorClassLike defines the set of class constants, constructors and
functions that must be supported by all package-generator-class-like classes.
*/
type PackageGeneratorClassLike interface {
	// Constructor Methods
	Make() PackageGeneratorLike
}

// Instance Declarations

/*
ClassGeneratorLike defines the set of aspects and methods that must be supported by
all class-generator-like instances.
*/
type ClassGeneratorLike interface {
	// Primary Methods
	GetClass() ClassGeneratorClassLike
	GenerateClass(
		moduleName string,
		wikiPath string,
		packageName string,
		className string,
		synthesizer ClassTemplateDriven,
	) string
}

/*
ModuleGeneratorLike defines the set of aspects and methods that must be supported by
all module-generator-like instances.
*/
type ModuleGeneratorLike interface {
	// Primary Methods
	GetClass() ModuleGeneratorClassLike
	GenerateModule(
		wikiPath string,
		synthesizer ModuleTemplateDriven,
	) string
}

/*
PackageGeneratorLike defines the set of aspects and methods that must be supported by
all package-generator-like instances.
*/
type PackageGeneratorLike interface {
	// Primary Methods
	GetClass() PackageGeneratorClassLike
	GeneratePackage(
		moduleName string,
		wikiPath string,
		packageName string,
		synthesizer PackageTemplateDriven,
	) string
}

// Aspect Declarations

/*
ClassTemplateDriven defines the set of method signatures that must be supported by
all class-template-driven synthesizers.
*/
type ClassTemplateDriven interface {
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

/*
ModuleTemplateDriven defines the set of method signatures that must be supported by
all module-template-driven synthesizers.
*/
type ModuleTemplateDriven interface {
	CreateLegalNotice() string
	CreateTypeAliases() string
	CreateUniversalConstructors() string
	PerformGlobalUpdates(
		source string,
	) string
}

/*
PackageTemplateDriven defines the set of method signatures that must be supported by
all package-template-driven synthesizers.
*/
type PackageTemplateDriven interface {
	CreateLegalNotice() string
	CreatePackageDescription() string
	CreateTypeDeclarations() string
	CreateFunctionalDeclarations() string
	CreateClassDeclarations() string
	CreateInstanceDeclarations() string
	CreateAspectDeclarations() string
	PerformGlobalUpdates(
		source string,
	) string
}
