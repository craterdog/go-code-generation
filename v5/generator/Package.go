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

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package generator

import ()

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
ClassGeneratorClassLike declares the set of class constants, constructors and
functions that must be supported by all class-generator-class-like classes.
*/
type ClassGeneratorClassLike interface {
	// Constructor Methods
	Make() ClassGeneratorLike
}

/*
ModuleGeneratorClassLike declares the set of class constants, constructors and
functions that must be supported by all module-generator-class-like classes.
*/
type ModuleGeneratorClassLike interface {
	// Constructor Methods
	Make() ModuleGeneratorLike
}

/*
PackageGeneratorClassLike declares the set of class constants, constructors and
functions that must be supported by all package-generator-class-like classes.
*/
type PackageGeneratorClassLike interface {
	// Constructor Methods
	Make() PackageGeneratorLike
}

/*
ProcessorGeneratorClassLike declares the set of class constants, constructors and
functions that must be supported by all processor-generator-class-like classes.
*/
type ProcessorGeneratorClassLike interface {
	// Constructor Methods
	Make() ProcessorGeneratorLike
}

/*
TemplateGeneratorClassLike declares the set of class constants, constructors and
functions that must be supported by all template-generator-class-like classes.
*/
type TemplateGeneratorClassLike interface {
	// Constructor Methods
	Make() TemplateGeneratorLike
}

// INSTANCE DECLARATIONS

/*
ClassGeneratorLike declares the set of aspects and methods that must be
supported by all class-generator-like instances.
*/
type ClassGeneratorLike interface {
	// Principal Methods
	GetClass() ClassGeneratorClassLike
	GenerateClass(
		moduleName string,
		packageName string,
		className string,
		synthesizer ClassTemplateDriven,
	) string
}

/*
ModuleGeneratorLike declares the set of aspects and methods that must be
supported by all module-generator-like instances.
*/
type ModuleGeneratorLike interface {
	// Principal Methods
	GetClass() ModuleGeneratorClassLike
	GenerateModule(
		moduleName string,
		wikiPath string,
		synthesizer ModuleTemplateDriven,
	) string
}

/*
PackageGeneratorLike declares the set of aspects and methods that must be
supported by all package-generator-like instances.
*/
type PackageGeneratorLike interface {
	// Principal Methods
	GetClass() PackageGeneratorClassLike
	GeneratePackage(
		moduleName string,
		wikiPath string,
		packageName string,
		synthesizer PackageTemplateDriven,
	) string
}

/*
ProcessorGeneratorLike declares the set of aspects and methods that must be
supported by all processor-generator-like instances.
*/
type ProcessorGeneratorLike interface {
	// Principal Methods
	GetClass() ProcessorGeneratorClassLike
	GenerateClass(
		moduleName string,
		packageName string,
		className string,
		synthesizer ClassTemplateDriven,
	) string
}

/*
TemplateGeneratorLike declares the set of aspects and methods that must be
supported by all template-generator-like instances.
*/
type TemplateGeneratorLike interface {
	// Principal Methods
	GetClass() TemplateGeneratorClassLike
	GenerateClass(
		moduleName string,
		packageName string,
		className string,
		synthesizer ClassTemplateDriven,
	) string
}

// ASPECT DECLARATIONS

/*
ClassTemplateDriven declares the set of method signatures that must be
supported by all class-template-driven synthesizers.
*/
type ClassTemplateDriven interface {
	CreateLegalNotice() string
	CreateAccessFunction() string
	CreateConstructorMethods() string
	CreateConstantMethods() string
	CreateFunctionMethods() string
	CreatePrincipalMethods() string
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
ModuleTemplateDriven declares the set of method signatures that must be
supported by all module-template-driven synthesizers.
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
PackageTemplateDriven declares the set of method signatures that must be
supported by all package-template-driven synthesizers.
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
