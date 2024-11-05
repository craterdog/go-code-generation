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
	packageName string,
	synthesizer TemplateDriven,
) string {
	// Begin with a package template.
	var result = generatorReference().packageTemplate_

	// Insert the legal notice.
	var legalNotice = synthesizer.CreateLegalNotice()
	result = uti.ReplaceAll(result, "legalNotice", legalNotice)

	// Insert the package declaration.
	var packageDeclaration = synthesizer.CreatePackageDeclaration()
	result = uti.ReplaceAll(result, "packageDeclaration", packageDeclaration)

	// Insert the type declarations.
	var typeDeclarations = synthesizer.CreateTypeDeclarations()
	result = uti.ReplaceAll(result, "typeDeclarations", typeDeclarations)

	// Insert the class declarations.
	var classDeclarations = synthesizer.CreateClassDeclarations()
	result = uti.ReplaceAll(result, "classDeclarations", classDeclarations)

	// Insert the instance declarations.
	var instanceDeclarations = synthesizer.CreateInstanceDeclarations()
	result = uti.ReplaceAll(result, "instanceDeclarations", instanceDeclarations)

	// Insert the aspect declarations.
	var aspectDeclarations = synthesizer.CreateAspectDeclarations()
	result = uti.ReplaceAll(result, "aspectDeclarations", aspectDeclarations)

	// Insert the module imports (this must be done last).
	var moduleImports = synthesizer.CreateModuleImports()
	result = uti.ReplaceAll(result, "moduleImports", moduleImports)
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
	packageTemplate_: `<Notice><PackageDeclaration><ModuleImports>

// Type Declarations<TypeDeclarations>

// Class Declarations<ClassDeclarations>

// Instance Declarations<InstanceDeclarations>

// Aspect Declarations<AspectDeclarations>
`,
}
