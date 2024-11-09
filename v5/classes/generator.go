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
	moduleName string,
	wikiPath string,
	packageName string,
	className string,
	synthesizer TemplateDriven,
) string {
	// Begin with a class template.
	var class = generatorReference().classTemplate_

	// Create the legal notice.
	var legalNotice = synthesizer.CreateLegalNotice()
	class = uti.ReplaceAll(class, "legalNotice", legalNotice)

	// Create the access function.
	var accessFunction = synthesizer.CreateAccessFunction()
	class = uti.ReplaceAll(class, "accessFunction", accessFunction)

	// Create the constructor methods.
	var constructorMethods = synthesizer.CreateConstructorMethods()
	class = uti.ReplaceAll(class, "constructorMethods", constructorMethods)

	// Create the constant methods.
	var constantMethods = synthesizer.CreateConstantMethods()
	class = uti.ReplaceAll(class, "constantMethods", constantMethods)

	// Create the function methods.
	var functionMethods = synthesizer.CreateFunctionMethods()
	class = uti.ReplaceAll(class, "functionMethods", functionMethods)

	// Create the primary methods.
	var primaryMethods = synthesizer.CreatePrimaryMethods()
	class = uti.ReplaceAll(class, "primaryMethods", primaryMethods)

	// Create the attribute methods.
	var attributeMethods = synthesizer.CreateAttributeMethods()
	class = uti.ReplaceAll(class, "attributeMethods", attributeMethods)

	// Create the aspect methods.
	var aspectMethods = synthesizer.CreateAspectMethods()
	class = uti.ReplaceAll(class, "aspectMethods", aspectMethods)

	// Create the private methods.
	var privateMethods = synthesizer.CreatePrivateMethods()
	class = uti.ReplaceAll(class, "privateMethods", privateMethods)

	// Create the instance structure.
	var instanceStructure = synthesizer.CreateInstanceStructure()
	class = uti.ReplaceAll(class, "instanceStructure", instanceStructure)

	// Create the class structure.
	var classStructure = synthesizer.CreateClassStructure()
	class = uti.ReplaceAll(class, "classStructure", classStructure)

	// Create the class reference.
	var classReference = synthesizer.CreateClassReference()
	class = uti.ReplaceAll(class, "classReference", classReference)

	// Perform global updates (this must be done last).
	class = synthesizer.PerformGlobalUpdates(class)
	class = uti.ReplaceAll(class, "moduleName", moduleName)
	class = uti.ReplaceAll(class, "wikiPath", wikiPath)
	class = uti.ReplaceAll(class, "packageName", packageName)
	class = uti.ReplaceAll(class, "className", className)

	return class
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
	classTemplate_: `<LegalNotice>
package <PackageName><ModuleImports>

// CLASS INTERFACE
<AccessFunction>
<ConstructorMethods><ConstantMethods><FunctionMethods>

// INSTANCE INTERFACE
<PrimaryMethods><AttributeMethods><AspectMethods>

// PROTECTED INTERFACE
<PrivateMethods>
<InstanceStructure><ClassStructure><ClassReference>
`,
}
