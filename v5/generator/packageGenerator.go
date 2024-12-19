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

package generator

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func PackageGeneratorClass() PackageGeneratorClassLike {
	return packageGeneratorClassReference()
}

// Constructor Methods

func (c *packageGeneratorClass_) PackageGenerator() PackageGeneratorLike {
	var instance = &packageGenerator_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *packageGenerator_) GetClass() PackageGeneratorClassLike {
	return packageGeneratorClassReference()
}

func (v *packageGenerator_) GeneratePackage(
	moduleName string,
	wikiPath string,
	packageName string,
	existing string,
	synthesizer PackageTemplateDriven,
) string {
	// Begin with a package template.
	var generated = packageGeneratorClassReference().packageTemplate_

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

	// Perform global updates (this must be done last).
	generated = synthesizer.PerformGlobalUpdates(existing, generated)
	generated = uti.ReplaceAll(generated, "moduleName", moduleName)
	generated = uti.ReplaceAll(generated, "wikiPath", wikiPath)
	generated = uti.ReplaceAll(generated, "packageName", packageName)

	return generated
}

// PROTECTED INTERFACE

// Instance Structure

type packageGenerator_ struct {
	// Declare the instance attributes.
}

// Class Structure

type packageGeneratorClass_ struct {
	// Declare the class constants.
	packageTemplate_ string
}

// Class Reference

func packageGeneratorClassReference() *packageGeneratorClass_ {
	return packageGeneratorClassReference_
}

var packageGeneratorClassReference_ = &packageGeneratorClass_{
	// Initialize the class constants.
	packageTemplate_: `<LegalNotice>
/*<WarningMessage><PackageDescription>

For detailed documentation on this package refer to the wiki:
  - <WikiPath>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

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
}
