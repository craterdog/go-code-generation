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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func ClassAssemblerClass() ClassAssemblerClassLike {
	return classAssemblerClass()
}

// Constructor Methods

func (c *classAssemblerClass_) ClassAssembler() ClassAssemblerLike {
	var instance = &classAssembler_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *classAssembler_) GetClass() ClassAssemblerClassLike {
	return classAssemblerClass()
}

func (v *classAssembler_) AssembleClass(
	moduleName string,
	packageName string,
	className string,
	existing string,
	synthesizer ClassTemplateDriven,
) string {
	// Begin with a class template.
	var generated = classAssemblerClass().classTemplate_

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

	// Create the imported packages.
	var importedPackages = synthesizer.CreateImportedPackages()
	generated = uti.ReplaceAll(
		generated,
		"importedPackages",
		importedPackages,
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

	// Create the aspect interface methods.
	var aspectInterfaces = synthesizer.CreateAspectInterfaces()
	generated = uti.ReplaceAll(
		generated,
		"aspectInterfaces",
		aspectInterfaces,
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
	var classReference = synthesizer.CreateClass()
	generated = uti.ReplaceAll(
		generated,
		"classReference",
		classReference,
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
		"packageName",
		packageName,
	)
	generated = uti.ReplaceAll(
		generated,
		"className",
		className,
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

type classAssembler_ struct {
	// Declare the instance attributes.
}

// Class Structure

type classAssemblerClass_ struct {
	// Declare the class constants.
	classTemplate_ string
	importedPath_  string
}

// Class Reference

func classAssemblerClass() *classAssemblerClass_ {
	return classAssemblerClassReference_
}

var classAssemblerClassReference_ = &classAssemblerClass_{
	// Initialize the class constants.
	classTemplate_: `<LegalNotice><WarningMessage>
package <PackageName>

import (<ImportedPackages>)

// CLASS INTERFACE
<AccessFunction><ConstructorMethods><ConstantMethods><FunctionMethods>
// INSTANCE INTERFACE
<PrincipalMethods><AttributeMethods><AspectInterfaces>
// PROTECTED INTERFACE
<PrivateMethods><InstanceStructure><ClassStructure><ClassReference>`,

	importedPath_: `
	<~packageAcronym> <packagePath>`,
}
