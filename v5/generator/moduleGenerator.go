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
	col "github.com/craterdog/go-collection-framework/v5"
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	reg "regexp"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ModuleGeneratorClass() ModuleGeneratorClassLike {
	return moduleGeneratorClass()
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
	return moduleGeneratorClass()
}

func (v *moduleGenerator_) GenerateModule(
	moduleName string,
	wikiPath string,
	existing string,
	synthesizer ModuleTemplateDriven,
) string {
	// Begin with a module template.
	var generated = moduleGeneratorClass().moduleTemplate_

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

	// Create the imported packages.
	var importedPackages = synthesizer.CreateImportedPackages()
	generated = uti.ReplaceAll(
		generated,
		"importedPackages",
		importedPackages,
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
		"wikiPath",
		wikiPath,
	)

	// Clean up and format the imported packages (must be done last).
	var class = moduleGeneratorClass()
	generated = class.formatImportedPackages(
		existing,
		generated,
	)
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (c *moduleGeneratorClass_) createImportedPath(
	packageAcronym string,
	packagePath string,
) string {
	var importedPath = c.importedPath_
	importedPath = uti.ReplaceAll(
		importedPath,
		"packageAcronym",
		packageAcronym,
	)
	importedPath = uti.ReplaceAll(
		importedPath,
		"packagePath",
		packagePath,
	)
	return importedPath
}

func (c *moduleGeneratorClass_) extractImportedPackages(
	source string,
) (
	packages abs.CatalogLike[string, string],
) {
	packages = col.Catalog[string, string]()
	var lower_ = `\p{Ll}`
	var digit_ = `\p{Nd}`
	var acronym_ = `(` + lower_ + `(?:` + lower_ + `|` + digit_ + `){2})`
	var white_ = `[ \t\r\n]*`
	var path_ = `("[^"]+")`
	var pattern = `import \((?:.|\r?\n)*?\)`
	var matcher = reg.MustCompile(pattern)
	var imports = matcher.FindString(source)
	var lines = sts.Split(imports, "\n")
	var count = len(lines)
	if count > 2 {
		pattern = acronym_ + white_ + path_
		matcher = reg.MustCompile(pattern)
		lines = lines[1 : count-1]
		for _, line := range lines {
			var matches = matcher.FindStringSubmatch(line)
			var packageAcronym = matches[1]
			var packagePath = matches[2]
			packages.SetValue(packageAcronym, packagePath)
		}
	}
	return
}

func (c *moduleGeneratorClass_) formatImportedPackages(
	existing string,
	generated string,
) string {
	// Start with the existing imported packages.
	var imports = c.extractImportedPackages(existing)

	// Add in the generated imported packages.
	imports = abs.CatalogClass[string, string]().Merge(
		imports, c.extractImportedPackages(generated),
	)

	// Sort the imported packages by path rather than name.
	imports.SortValuesWithRanker(
		func(first, second abs.AssociationLike[string, string]) col.Rank {
			var firstValue = first.GetValue()
			var secondValue = second.GetValue()
			switch {
			case firstValue < secondValue:
				return col.LesserRank
			case firstValue > secondValue:
				return col.GreaterRank
			default:
				return col.EqualRank
			}
		},
	)

	// Create revised imported package statements for each imported package.
	var importedPackages string
	var packages = imports.GetIterator()
	for packages.HasNext() {
		var association = packages.GetNext()
		var packageName = association.GetKey()
		var packagePath = association.GetValue()
		if sts.Contains(generated, packageName+".") {
			// Only import packages that are actually used in the generated code.
			importedPackages += c.createImportedPath(
				packageName,
				packagePath,
			)
		}
	}
	if uti.IsDefined(importedPackages) {
		importedPackages += "\n"
	}

	// Replace the generated imported packages with the revised ones.
	var pattern = `import \((?:.|\r?\n)*?\)`
	var matcher = reg.MustCompile(pattern)
	var generatedImports = matcher.FindString(generated)
	var revisedImports = "import (" + importedPackages + ")"
	generated = sts.ReplaceAll(
		generated,
		generatedImports,
		revisedImports,
	)
	return generated
}

// Instance Structure

type moduleGenerator_ struct {
	// Declare the instance attributes.
}

// Class Structure

type moduleGeneratorClass_ struct {
	// Declare the class constants.
	moduleTemplate_ string
	importedPath_   string
}

// Class Reference

func moduleGeneratorClass() *moduleGeneratorClass_ {
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

	importedPath_: `
	<~packageAcronym> <packagePath>`,
}
