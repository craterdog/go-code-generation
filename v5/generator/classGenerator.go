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

func ClassGeneratorClass() ClassGeneratorClassLike {
	return classGeneratorClassReference()
}

// Constructor Methods

func (c *classGeneratorClass_) ClassGenerator() ClassGeneratorLike {
	var instance = &classGenerator_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *classGenerator_) GetClass() ClassGeneratorClassLike {
	return classGeneratorClassReference()
}

func (v *classGenerator_) GenerateClass(
	moduleName string,
	packageName string,
	className string,
	existing string,
	synthesizer ClassTemplateDriven,
) string {
	// Begin with a class template.
	var generated = classGeneratorClassReference().classTemplate_

	// Create the legal notice.
	var legalNotice = synthesizer.CreateLegalNotice()
	generated = uti.ReplaceAll(
		generated,
		"legalNotice",
		legalNotice,
	)

	// Create the warning message.
	var warningMessage = synthesizer.CreateWarningMessage()
	if uti.IsDefined(warningMessage) {
		warningMessage = "\n/*" + warningMessage + "*/\n"
	}
	generated = uti.ReplaceAll(
		generated,
		"warningMessage",
		warningMessage,
	)

	// Create the access function.
	var accessFunction = synthesizer.CreateAccessFunction()
	generated = uti.ReplaceAll(
		generated,
		"accessFunction",
		accessFunction,
	)

	// Create the constructor methods.
	var constructorMethods = synthesizer.CreateConstructorMethods()
	generated = uti.ReplaceAll(
		generated,
		"constructorMethods",
		constructorMethods,
	)

	// Create the constant methods.
	var constantMethods = synthesizer.CreateConstantMethods()
	generated = uti.ReplaceAll(
		generated,
		"constantMethods",
		constantMethods,
	)

	// Create the function methods.
	var functionMethods = synthesizer.CreateFunctionMethods()
	generated = uti.ReplaceAll(
		generated,
		"functionMethods",
		functionMethods,
	)

	// Create the principal methods.
	var principalMethods = synthesizer.CreatePrincipalMethods()
	generated = uti.ReplaceAll(
		generated,
		"principalMethods",
		principalMethods,
	)

	// Create the attribute methods.
	var attributeMethods = synthesizer.CreateAttributeMethods()
	generated = uti.ReplaceAll(
		generated,
		"attributeMethods",
		attributeMethods,
	)

	// Create the aspect methods.
	var aspectMethods = synthesizer.CreateAspectMethods()
	generated = uti.ReplaceAll(
		generated,
		"aspectMethods",
		aspectMethods,
	)

	// Create the private methods.
	var privateMethods = synthesizer.CreatePrivateMethods()
	generated = uti.ReplaceAll(
		generated,
		"privateMethods",
		privateMethods,
	)

	// Create the instance structure.
	var instanceStructure = synthesizer.CreateInstanceStructure()
	generated = uti.ReplaceAll(
		generated,
		"instanceStructure",
		instanceStructure,
	)

	// Create the class structure.
	var classStructure = synthesizer.CreateClassStructure()
	generated = uti.ReplaceAll(
		generated,
		"classStructure",
		classStructure,
	)

	// Create the class reference.
	var classReference = synthesizer.CreateClassReference()
	generated = uti.ReplaceAll(
		generated,
		"classReference",
		classReference,
	)

	// Perform global updates (this must be done last).
	generated = synthesizer.PerformGlobalUpdates(existing, generated)
	generated = uti.ReplaceAll(generated, "moduleName", moduleName)
	generated = uti.ReplaceAll(generated, "packageName", packageName)
	generated = uti.ReplaceAll(generated, "className", className)

	return generated
}

// PROTECTED INTERFACE

// Instance Structure

type classGenerator_ struct {
	// Declare the instance attributes.
}

// Class Structure

type classGeneratorClass_ struct {
	// Declare the class constants.
	classTemplate_ string
}

// Class Reference

func classGeneratorClassReference() *classGeneratorClass_ {
	return classGeneratorClassReference_
}

var classGeneratorClassReference_ = &classGeneratorClass_{
	// Initialize the class constants.
	classTemplate_: `<LegalNotice><WarningMessage>
package <PackageName>

import (<ImportedPackages>)

// CLASS INTERFACE
<AccessFunction><ConstructorMethods><ConstantMethods><FunctionMethods>
// INSTANCE INTERFACE
<PrincipalMethods><AttributeMethods><AspectMethods>
// PROTECTED INTERFACE
<PrivateMethods><InstanceStructure><ClassStructure><ClassReference>`,
}
