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

package assembler

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func PackageAssemblerClass() PackageAssemblerClassLike {
	return packageAssemblerClass()
}

// Constructor Methods

func (c *packageAssemblerClass_) PackageAssembler() PackageAssemblerLike {
	var instance = &packageAssembler_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *packageAssembler_) GetClass() PackageAssemblerClassLike {
	return packageAssemblerClass()
}

func (v *packageAssembler_) AssemblePackage(
	moduleName string,
	wikiPath string,
	packageName string,
	existing string,
	synthesizer PackageTemplateDriven,
) string {
	// Begin with a package template.
	var generated = packageAssemblerClass().packageTemplate_

	// Create the legal notice.
	var legalNotice = synthesizer.CreateLegalNotice()
	generated = uti.ReplaceAll(
		generated,
		"legalNotice",
		legalNotice,
	)

	// Create the warning message.
	var warningMessage = synthesizer.CreateWarningMessage()
	generated = uti.ReplaceAll(
		generated,
		"warningMessage",
		warningMessage,
	)

	// Create the package description.
	var packageDescription = synthesizer.CreatePackageDescription()
	generated = uti.ReplaceAll(
		generated,
		"packageDescription",
		packageDescription,
	)

	// Create the imported packages.
	var importedPackages = synthesizer.CreateImportedPackages()
	generated = uti.ReplaceAll(
		generated,
		"importedPackages",
		importedPackages,
	)

	// Create the type declarations.
	var typeDeclarations = synthesizer.CreateTypeDeclarations()
	generated = uti.ReplaceAll(
		generated,
		"typeDeclarations",
		typeDeclarations,
	)

	// Create the functional declarations.
	var functionalDeclarations = synthesizer.CreateFunctionalDeclarations()
	generated = uti.ReplaceAll(
		generated,
		"functionalDeclarations",
		functionalDeclarations,
	)

	// Create the class declarations.
	var classDeclarations = synthesizer.CreateClassDeclarations()
	generated = uti.ReplaceAll(
		generated,
		"classDeclarations",
		classDeclarations,
	)

	// Create the instance declarations.
	var instanceDeclarations = synthesizer.CreateInstanceDeclarations()
	generated = uti.ReplaceAll(
		generated,
		"instanceDeclarations",
		instanceDeclarations,
	)

	// Create the aspect declarations.
	var aspectDeclarations = synthesizer.CreateAspectDeclarations()
	generated = uti.ReplaceAll(
		generated,
		"aspectDeclarations",
		aspectDeclarations,
	)

	// Perform global updates.
	generated = synthesizer.PerformGlobalUpdates(
		existing,
		generated,
	)
	generated = uti.ReplaceAll(
		generated,
		"moduleName",
		moduleName,
	)
	generated = uti.ReplaceAll(
		generated,
		"wikiPath",
		wikiPath,
	)
	generated = uti.ReplaceAll(
		generated,
		"packageName",
		packageName,
	)

	// Clean up and format the imported packages (must be done last).
	var class = moduleAssemblerClass()
	generated = class.formatImportedPackages(
		existing,
		generated,
	)
	return generated
}

// PROTECTED INTERFACE

// Instance Structure

type packageAssembler_ struct {
	// Declare the instance attributes.
}

// Class Structure

type packageAssemblerClass_ struct {
	// Declare the class constants.
	packageTemplate_ string
	importedPath_    string
}

// Class Reference

func packageAssemblerClass() *packageAssemblerClass_ {
	return packageAssemblerClassReference_
}

var packageAssemblerClassReference_ = &packageAssemblerClass_{
	// Initialize the class constants.
	packageTemplate_: `<LegalNotice>
/*<WarningMessage><PackageDescription>

For detailed documentation on this package refer to the wiki:
  - <WikiPath>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package <PackageName>

import (<ImportedPackages>)

// TYPE DECLARATIONS
<TypeDeclarations>
// FUNCTIONAL DECLARATIONS
<FunctionalDeclarations>
// CLASS DECLARATIONS
<ClassDeclarations>
// INSTANCE DECLARATIONS
<InstanceDeclarations>
// ASPECT DECLARATIONS
<AspectDeclarations>`,

	importedPath_: `
	<~packageAcronym> <packagePath>`,
}
