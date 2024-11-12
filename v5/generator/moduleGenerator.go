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

func ModuleGenerator() ModuleGeneratorClassLike {
	return moduleGeneratorReference()
}

// Constructor Methods

func (c *moduleGeneratorClass_) Make() ModuleGeneratorLike {
	var instance = &moduleGenerator_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *moduleGenerator_) GetClass() ModuleGeneratorClassLike {
	return moduleGeneratorReference()
}

func (v *moduleGenerator_) GenerateModule(
	wikiPath string,
	synthesizer ModuleTemplateDriven,
) string {
	// Begin with a module template.
	var source = moduleGeneratorReference().moduleTemplate_

	// Create the legal notice.
	var legalNotice = synthesizer.CreateLegalNotice()
	source = uti.ReplaceAll(source, "legalNotice", legalNotice)

	// Create the type aliases.
	var typeAliases = synthesizer.CreateTypeAliases()
	source = uti.ReplaceAll(source, "typeAliases", typeAliases)

	// Create the universal constructors.
	var universalConstructors = synthesizer.CreateUniversalConstructors()
	source = uti.ReplaceAll(source, "universalConstructors", universalConstructors)

	// Create the global functions.
	var globalFunctions = synthesizer.CreateGlobalFunctions()
	source = uti.ReplaceAll(source, "globalFunctions", globalFunctions)

	// Perform global updates (this must be done last).
	source = synthesizer.PerformGlobalUpdates(source)
	source = uti.ReplaceAll(source, "wikiPath", wikiPath)

	return source
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

func moduleGeneratorReference() *moduleGeneratorClass_ {
	return moduleGeneratorReference_
}

var moduleGeneratorReference_ = &moduleGeneratorClass_{
	// Initialize the class constants.
	moduleTemplate_: `<LegalNotice>
/*
Package "module" defines type aliases for the commonly used types defined in the
packages contained in this module.  It also provides a universal constructor for
each commonly used class that is exported by the module.  Each constructor
delegates the actual construction process to its corresponding concrete class
defined in the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - <WikiPath>
*/
package module<ModuleImports>

// TYPE ALIASES
<TypeAliases>

// UNIVERSAL CONSTRUCTORS
<UniversalConstructors>

// GLOBAL FUNCTIONS
<GlobalFunctions>`,
}
