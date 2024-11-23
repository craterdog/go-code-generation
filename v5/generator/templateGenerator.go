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

func TemplateGenerator() TemplateGeneratorClassLike {
	return templateGeneratorReference()
}

// Constructor Methods

func (c *templateGeneratorClass_) Make() TemplateGeneratorLike {
	var instance = &templateGenerator_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *templateGenerator_) GetClass() TemplateGeneratorClassLike {
	return templateGeneratorReference()
}

func (v *templateGenerator_) GenerateClass(
	moduleName string,
	wikiPath string,
	packageName string,
	className string,
	synthesizer ClassTemplateDriven,
) string {
	// Begin with a class template.
	var source = templateGeneratorReference().classTemplate_

	// Create the legal notice.
	var legalNotice = synthesizer.CreateLegalNotice()
	source = uti.ReplaceAll(source, "legalNotice", legalNotice)

	// Create the access function.
	var accessFunction = synthesizer.CreateAccessFunction()
	source = uti.ReplaceAll(source, "accessFunction", accessFunction)

	// Create the constructor methods.
	var constructorMethods = synthesizer.CreateConstructorMethods()
	source = uti.ReplaceAll(source, "constructorMethods", constructorMethods)

	// Create the constant methods.
	var constantMethods = synthesizer.CreateConstantMethods()
	source = uti.ReplaceAll(source, "constantMethods", constantMethods)

	// Create the function methods.
	var functionMethods = synthesizer.CreateFunctionMethods()
	source = uti.ReplaceAll(source, "functionMethods", functionMethods)

	// Create the principal methods.
	var principalMethods = synthesizer.CreatePrincipalMethods()
	source = uti.ReplaceAll(source, "principalMethods", principalMethods)

	// Create the attribute methods.
	var attributeMethods = synthesizer.CreateAttributeMethods()
	source = uti.ReplaceAll(source, "attributeMethods", attributeMethods)

	// Create the aspect methods.
	var aspectMethods = synthesizer.CreateAspectMethods()
	source = uti.ReplaceAll(source, "aspectMethods", aspectMethods)

	// Create the private methods.
	var privateMethods = synthesizer.CreatePrivateMethods()
	source = uti.ReplaceAll(source, "privateMethods", privateMethods)

	// Create the instance structure.
	var instanceStructure = synthesizer.CreateInstanceStructure()
	source = uti.ReplaceAll(source, "instanceStructure", instanceStructure)

	// Create the class structure.
	var classStructure = synthesizer.CreateClassStructure()
	source = uti.ReplaceAll(source, "classStructure", classStructure)

	// Create the class reference.
	var classReference = synthesizer.CreateClassReference()
	source = uti.ReplaceAll(source, "classReference", classReference)

	// Perform global updates (this must be done last).
	source = synthesizer.PerformGlobalUpdates(source)
	source = uti.ReplaceAll(source, "moduleName", moduleName)
	source = uti.ReplaceAll(source, "wikiPath", wikiPath)
	source = uti.ReplaceAll(source, "packageName", packageName)
	source = uti.ReplaceAll(source, "className", className)

	return source
}

// PROTECTED INTERFACE

// Instance Structure

type templateGenerator_ struct {
	// Declare the instance attributes.
}

// Class Structure

type templateGeneratorClass_ struct {
	// Declare the class constants.
	classTemplate_ string
}

// Class Reference

func templateGeneratorReference() *templateGeneratorClass_ {
	return templateGeneratorReference_
}

var templateGeneratorReference_ = &templateGeneratorClass_{
	// Initialize the class constants.
	classTemplate_: `<LegalNotice>
package <PackageName>
<ClassImports>
// CLASS INTERFACE
<AccessFunction><ConstructorMethods><ConstantMethods><FunctionMethods>
// INSTANCE INTERFACE
<PrincipalMethods><AttributeMethods><AspectMethods>
// PROTECTED INTERFACE
<PrivateMethods><InstanceStructure><ClassStructure><ClassReference>`,
}
