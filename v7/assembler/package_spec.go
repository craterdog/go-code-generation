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
Package "assembler" provides template-based code assemblers that can assemble
Go code snippets into complete Go files that conform to the Crater Dog
Technologies™ Go Coding Conventions located here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package assembler

import ()

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
ClassAssemblerClassLike declares the set of class constants, constructors and
functions that must be supported by all class-assembler-class-like classes.
*/
type ClassAssemblerClassLike interface {
	// Constructor Methods
	ClassAssembler() ClassAssemblerLike
}

/*
ModuleAssemblerClassLike declares the set of class constants, constructors and
functions that must be supported by all module-assembler-class-like classes.
*/
type ModuleAssemblerClassLike interface {
	// Constructor Methods
	ModuleAssembler() ModuleAssemblerLike
}

/*
PackageAssemblerClassLike declares the set of class constants, constructors and
functions that must be supported by all package-assembler-class-like classes.
*/
type PackageAssemblerClassLike interface {
	// Constructor Methods
	PackageAssembler() PackageAssemblerLike
}

// INSTANCE DECLARATIONS

/*
ClassAssemblerLike declares the set of aspects and methods that must be
supported by all class-assembler-like instances.
*/
type ClassAssemblerLike interface {
	// Principal Methods
	GetClass() ClassAssemblerClassLike
	AssembleClass(
		moduleName string,
		packageName string,
		className string,
		existing string,
		synthesizer ClassTemplateDriven,
	) string
}

/*
ModuleAssemblerLike declares the set of aspects and methods that must be
supported by all module-assembler-like instances.
*/
type ModuleAssemblerLike interface {
	// Principal Methods
	GetClass() ModuleAssemblerClassLike
	AssembleModule(
		moduleName string,
		wikiPath string,
		existing string,
		synthesizer ModuleTemplateDriven,
	) string
}

/*
PackageAssemblerLike declares the set of aspects and methods that must be
supported by all package-assembler-like instances.
*/
type PackageAssemblerLike interface {
	// Principal Methods
	GetClass() PackageAssemblerClassLike
	AssemblePackage(
		moduleName string,
		wikiPath string,
		packageName string,
		existing string,
		synthesizer PackageTemplateDriven,
	) string
}

// ASPECT DECLARATIONS

/*
ClassTemplateDriven declares the set of method signatures that must be
supported by all class-template-driven synthesizers.
*/
type ClassTemplateDriven interface {
	CreateLegalNotice() string
	CreateWarningMessage() string
	CreateImportedPackages() string
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
	CreateClass() string
	PerformGlobalUpdates(
		existing string,
		generated string,
	) string
}

/*
ModuleTemplateDriven declares the set of method signatures that must be
supported by all module-template-driven synthesizers.
*/
type ModuleTemplateDriven interface {
	CreateLegalNotice() string
	CreateWarningMessage() string
	CreateImportedPackages() string
	CreateTypeAliases() string
	CreateClassConstructors() string
	PerformGlobalUpdates(
		existing string,
		generated string,
	) string
}

/*
PackageTemplateDriven declares the set of method signatures that must be
supported by all package-template-driven synthesizers.
*/
type PackageTemplateDriven interface {
	CreateLegalNotice() string
	CreateWarningMessage() string
	CreatePackageDescription() string
	CreateImportedPackages() string
	CreateTypeDeclarations() string
	CreateFunctionalDeclarations() string
	CreateClassDeclarations() string
	CreateInstanceDeclarations() string
	CreateAspectDeclarations() string
	PerformGlobalUpdates(
		existing string,
		generated string,
	) string
}
