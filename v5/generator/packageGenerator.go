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

func PackageGenerator() PackageGeneratorClassLike {
	return packageGeneratorReference()
}

// Constructor Methods

func (c *packageGeneratorClass_) Make() PackageGeneratorLike {
	var instance = &packageGenerator_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *packageGenerator_) GetClass() PackageGeneratorClassLike {
	return packageGeneratorReference()
}

func (v *packageGenerator_) GeneratePackage(
	moduleName string,
	wikiPath string,
	packageName string,
	synthesizer PackageTemplateDriven,
) string {
	// Begin with a package template.
	var source = packageGeneratorReference().packageTemplate_

	// Create the legal notice.
	var legalNotice = synthesizer.CreateLegalNotice()
	source = uti.ReplaceAll(source, "legalNotice", legalNotice)

	// Create the package description.
	var packageDescription = synthesizer.CreatePackageDescription()
	source = uti.ReplaceAll(source, "packageDescription", packageDescription)

	// Create the type declarations.
	var typeDeclarations = synthesizer.CreateTypeDeclarations()
	source = uti.ReplaceAll(source, "typeDeclarations", typeDeclarations)

	// Create the functional declarations.
	var functionalDeclarations = synthesizer.CreateFunctionalDeclarations()
	source = uti.ReplaceAll(source, "functionalDeclarations", functionalDeclarations)

	// Create the class declarations.
	var classDeclarations = synthesizer.CreateClassDeclarations()
	source = uti.ReplaceAll(source, "classDeclarations", classDeclarations)

	// Create the instance declarations.
	var instanceDeclarations = synthesizer.CreateInstanceDeclarations()
	source = uti.ReplaceAll(source, "instanceDeclarations", instanceDeclarations)

	// Create the aspect declarations.
	var aspectDeclarations = synthesizer.CreateAspectDeclarations()
	source = uti.ReplaceAll(source, "aspectDeclarations", aspectDeclarations)

	// Perform global updates (this must be done last).
	source = synthesizer.PerformGlobalUpdates(source)
	source = uti.ReplaceAll(source, "moduleName", moduleName)
	source = uti.ReplaceAll(source, "wikiPath", wikiPath)
	source = uti.ReplaceAll(source, "packageName", packageName)

	return source
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

func packageGeneratorReference() *packageGeneratorClass_ {
	return packageGeneratorReference_
}

var packageGeneratorReference_ = &packageGeneratorClass_{
	// Initialize the class constants.
	packageTemplate_: `<LegalNotice>
/*<PackageDescription>

For detailed documentation on this package refer to the wiki:
  - https://<WikiPath>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package <PackageName>
<PackageImports>
// Type Declarations
<TypeDeclarations>
// Functional Declarations
<FunctionalDeclarations>
// Class Declarations
<ClassDeclarations>
// Instance Declarations
<InstanceDeclarations>
// Aspect Declarations
<AspectDeclarations>`,
}
