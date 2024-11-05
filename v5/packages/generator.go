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

func (c *generatorClass_) Make(
	synthesizer TemplateDriven,
) GeneratorLike {
	var instance = &generator_{
		// Initialize the instance attributes.
		synthesizer_: synthesizer,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *generator_) GetClass() GeneratorClassLike {
	return generatorReference()
}

func (v *generator_) GeneratePackage() string {
	// Begin with a package template.
	var result = generatorReference().packageTemplate_

	// Insert the legal notice.
	var legalNotice = v.synthesizer_.GenerateLegalNotice()
	result = uti.ReplaceAll(result, "legalNotice", legalNotice)

	// Insert the package declaration.
	var packageDeclaration = v.synthesizer_.GeneratePackageDeclaration()
	result = uti.ReplaceAll(result, "packageDeclaration", packageDeclaration)

	// Insert the type declarations.
	var typeDeclarations = v.synthesizer_.GenerateTypeDeclarations()
	result = uti.ReplaceAll(result, "typeDeclarations", typeDeclarations)

	// Insert the class declarations.
	var classDeclarations = v.synthesizer_.GenerateClassDeclarations()
	result = uti.ReplaceAll(result, "classDeclarations", classDeclarations)

	// Insert the instance declarations.
	var instanceDeclarations = v.synthesizer_.GenerateInstanceDeclarations()
	result = uti.ReplaceAll(result, "instanceDeclarations", instanceDeclarations)

	// Insert the aspect declarations.
	var aspectDeclarations = v.synthesizer_.GenerateAspectDeclarations()
	result = uti.ReplaceAll(result, "aspectDeclarations", aspectDeclarations)

	// Insert the module imports (this must be done last).
	var moduleImports = v.synthesizer_.GenerateModuleImports()
	result = uti.ReplaceAll(result, "moduleImports", moduleImports)

	return result
}

// PROTECTED INTERFACE

// Instance Structure

type generator_ struct {
	// Declare the instance attributes.
	synthesizer_ TemplateDriven
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
