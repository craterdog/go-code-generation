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

package classes

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

func (v *generator_) GenerateClass(
	className string,
	synthesizer TemplateDriven,
) string {
	// Begin with a class template.
	var result = generatorReference().classTemplate_

	// Insert the legal notice.
	var legalNotice = synthesizer.CreateLegalNotice()
	result = uti.ReplaceAll(result, "legalNotice", legalNotice)

	// Insert the package declaration.
	var packageDeclaration = synthesizer.CreatePackageDeclaration()
	result = uti.ReplaceAll(result, "packageDeclaration", packageDeclaration)

	// Insert the access function.
	var accessFunction = synthesizer.CreateAccessFunction()
	result = uti.ReplaceAll(result, "accessFunction", accessFunction)

	// Insert the constructor methods.
	var constructorMethods = synthesizer.CreateConstructorMethods()
	result = uti.ReplaceAll(result, "constructorMethods", constructorMethods)

	// Insert the function methods.
	var functionMethods = synthesizer.CreateFunctionMethods()
	result = uti.ReplaceAll(result, "functionMethods", functionMethods)

	// Insert the primary methods.
	var primaryMethods = synthesizer.CreatePrimaryMethods()
	result = uti.ReplaceAll(result, "primaryMethods", primaryMethods)

	// Insert the attribute methods.
	var attributeMethods = synthesizer.CreateAttributeMethods()
	result = uti.ReplaceAll(result, "attributeMethods", attributeMethods)

	// Insert the aspect methods.
	var aspectMethods = synthesizer.CreateAspectMethods()
	result = uti.ReplaceAll(result, "aspectMethods", aspectMethods)

	// Insert the private methods.
	var privateMethods = synthesizer.CreatePrivateMethods()
	result = uti.ReplaceAll(result, "privateMethods", privateMethods)

	// Insert the instance structure.
	var instanceStructure = synthesizer.CreateInstanceStructure()
	result = uti.ReplaceAll(result, "instanceStructure", instanceStructure)

	// Insert the class structure.
	var classStructure = synthesizer.CreateClassStructure()
	result = uti.ReplaceAll(result, "classStructure", classStructure)

	// Insert the class reference.
	var classReference = synthesizer.CreateClassReference()
	result = uti.ReplaceAll(result, "classReference", classReference)

	// Insert the module imports (this must be done last).
	var moduleImports = synthesizer.CreateModuleImports()
	result = uti.ReplaceAll(result, "moduleImports", moduleImports)
	result = uti.ReplaceAll(result, "className", className)

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
	classTemplate_ string
}

// Class Reference

func generatorReference() *generatorClass_ {
	return generatorReference_
}

var generatorReference_ = &generatorClass_{
	// Initialize the class constants.
	classTemplate_: `<Notice><PackageDeclaration><ModuleImports>

// CLASS INTERFACE<AccessFunction><ConstructorMethods><FunctionMethods>

// INSTANCE INTERFACE<PrimaryMethods><AttributeMethods><AspectMethods>

// PROTECTED INTERFACE<PrivateMethods><InstanceStructure><ClassStructure><ClassReference>
`,
}
