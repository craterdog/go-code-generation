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

package packages

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Generator() GeneratorClassLike {
	return generatorReference()
}

// Constructor Methods

func (c *generatorClass_) Make() GeneratorLike {
	var instance = &generator_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *generator_) GetClass() GeneratorClassLike {
	return generatorReference()
}

func (v *generator_) GeneratePackage(
	moduleName string,
	wikiPath string,
	packageName string,
	synthesizer TemplateDriven,
) string {
	// Begin with a package template.
	var result = generatorReference().packageTemplate_

	// Create the legal notice.
	var legalNotice = synthesizer.CreateLegalNotice()
	result = uti.ReplaceAll(result, "legalNotice", legalNotice)

	// Create the package description.
	var packageDescription = synthesizer.CreatePackageDescription()
	result = uti.ReplaceAll(result, "packageDescription", packageDescription)

	// Create the type declarations.
	var typeDeclarations = synthesizer.CreateTypeDeclarations()
	result = uti.ReplaceAll(result, "typeDeclarations", typeDeclarations)

	// Create the functional declarations.
	var functionalDeclarations = synthesizer.CreateFunctionalDeclarations()
	result = uti.ReplaceAll(result, "functionalDeclarations", functionalDeclarations)

	// Create the class declarations.
	var classDeclarations = synthesizer.CreateClassDeclarations()
	result = uti.ReplaceAll(result, "classDeclarations", classDeclarations)

	// Create the instance declarations.
	var instanceDeclarations = synthesizer.CreateInstanceDeclarations()
	result = uti.ReplaceAll(result, "instanceDeclarations", instanceDeclarations)

	// Create the aspect declarations.
	var aspectDeclarations = synthesizer.CreateAspectDeclarations()
	result = uti.ReplaceAll(result, "aspectDeclarations", aspectDeclarations)

	// Create the module imports (this must be done last).
	var moduleImports = synthesizer.CreateModuleImports()
	result = uti.ReplaceAll(result, "moduleImports", moduleImports)

	// Update the package information.
	result = uti.ReplaceAll(result, "moduleName", moduleName)
	result = uti.ReplaceAll(result, "wikiPath", wikiPath)
	result = uti.ReplaceAll(result, "packageName", packageName)

	return result
}

// PROTECTED INTERFACE

// Instance Structure

type generator_ struct {
	// Declare the instance attributes.
}

// Class Structure

type generatorClass_ struct {
	// Declare the class constants.
	packageTemplate_ string
}

// Class Reference

func generatorReference() *generatorClass_ {
	return generatorReference_
}

var generatorReference_ = &generatorClass_{
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
package <PackageName><ModuleImports>

// Type Declarations<TypeDeclarations>

// Functional Declarations<FunctionalDeclarations>

// Class Declarations<ClassDeclarations>

// Instance Declarations<InstanceDeclarations>

// Aspect Declarations<AspectDeclarations>
`,
}
