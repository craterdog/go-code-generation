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

package modules

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
	wiki string,
	synthesizer TemplateDriven,
) GeneratorLike {
	var instance = &generator_{
		// Initialize the instance attributes.
		wiki_:        wiki,
		synthesizer_: synthesizer,
	}
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *generator_) GetClass() GeneratorClassLike {
	return generatorReference()
}

func (v *generator_) GenerateModule() string {
	// Begin with a module template.
	var result = generatorReference().moduleTemplate_

	// Insert the legal notice.
	var legalNotice = v.synthesizer_.GenerateLegalNotice()
	result = uti.ReplaceAll(result, "legalNotice", legalNotice)

	// Insert the type aliases.
	var typeAliases = v.synthesizer_.GenerateTypeAliases()
	result = uti.ReplaceAll(result, "typeAliases", typeAliases)

	// Insert the universal constructors.
	var universalConstructors = v.synthesizer_.GenerateUniversalConstructors()
	result = uti.ReplaceAll(result, "universalConstructors", universalConstructors)

	// Insert the global functions.
	var globalFunctions = v.synthesizer_.GenerateGlobalFunctions()
	result = uti.ReplaceAll(result, "globalFunctions", globalFunctions)

	// Insert the module imports (this must be done last).
	var moduleImports = v.synthesizer_.GenerateModuleImports()
	result = uti.ReplaceAll(result, "moduleImports", moduleImports)

	return result
}

// PROTECTED INTERFACE

// Instance Structure

type generator_ struct {
	// Declare the instance attributes.
	wiki_        string
	synthesizer_ TemplateDriven
}

// Class Structure

type generatorClass_ struct {
	// Declare the class constants.
	moduleTemplate_ string
}

// Class Reference

func generatorReference() *generatorClass_ {
	return generatorReference_
}

var generatorReference_ = &generatorClass_{
	// Initialize the class constants.
	moduleTemplate_: `<Notice>
/*
Package "module" defines type aliases for the commonly used types defined in the
packages contained in this module.  It also provides a universal constructor for
each commonly used class that is exported by this module.  Each constructor
delegates the actual construction process to its corresponding concrete class
defined in a package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - <Wiki>
*/
package module<ModuleImports>

// Type Aliases<TypeAliases>

// Universal Constructors<UniversalConstructors>

// Global Functions<GlobalFunctions>
`,
}
