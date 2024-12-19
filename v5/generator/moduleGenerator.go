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

func ModuleGeneratorClass() ModuleGeneratorClassLike {
	return moduleGeneratorClassReference()
}

// Constructor Methods

func (c *moduleGeneratorClass_) ModuleGenerator() ModuleGeneratorLike {
	var instance = &moduleGenerator_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *moduleGenerator_) GetClass() ModuleGeneratorClassLike {
	return moduleGeneratorClassReference()
}

func (v *moduleGenerator_) GenerateModule(
	moduleName string,
	wikiPath string,
	existing string,
	synthesizer ModuleTemplateDriven,
) string {
	// Begin with a module template.
	var generated = moduleGeneratorClassReference().moduleTemplate_

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

	// Create the type aliases.
	var typeAliases = synthesizer.CreateTypeAliases()
	generated = uti.ReplaceAll(
		generated,
		"typeAliases",
		typeAliases,
	)

	// Create the class constructors.
	var classConstructors = synthesizer.CreateClassConstructors()
	generated = uti.ReplaceAll(
		generated,
		"classConstructors",
		classConstructors,
	)

	// Perform global updates (this must be done last).
	generated = synthesizer.PerformGlobalUpdates(existing, generated)
	generated = uti.ReplaceAll(generated, "moduleName", moduleName)
	generated = uti.ReplaceAll(generated, "wikiPath", wikiPath)

	return generated
}

// PROTECTED INTERFACE

// Instance Structure

type moduleGenerator_ struct {
	// Declare the instance attributes.
}

// Class Structure

type moduleGeneratorClass_ struct {
	// Declare the class constants.
	moduleTemplate_ string
}

// Class Reference

func moduleGeneratorClassReference() *moduleGeneratorClass_ {
	return moduleGeneratorClassReference_
}

var moduleGeneratorClassReference_ = &moduleGeneratorClass_{
	// Initialize the class constants.
	moduleTemplate_: `<LegalNotice>
/*<WarningMessage>
Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - <WikiPath>
*/
package module

import (<ImportedPackages>)

// TYPE ALIASES
<TypeAliases>
// CLASS CONSTRUCTORS
<ClassConstructors>
// GLOBAL FUNCTIONS
`,
}
