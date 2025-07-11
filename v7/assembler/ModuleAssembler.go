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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ModuleAssemblerClass() ModuleAssemblerClassLike {
	return moduleAssemblerClass()
}

// Constructor Methods

func (c *moduleAssemblerClass_) ModuleAssembler() ModuleAssemblerLike {
	var instance = &moduleAssembler_{
		// Initialize the instance attributes.
	}
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *moduleAssembler_) GetClass() ModuleAssemblerClassLike {
	return moduleAssemblerClass()
}

func (v *moduleAssembler_) AssembleModule(
	moduleName string,
	wikiPath string,
	existing string,
	synthesizer ModuleTemplateDriven,
) string {
	// Begin with a module template.
	var generated = moduleAssemblerClass().moduleTemplate_

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

	// Create the class accessors.
	var classAccessors = synthesizer.CreateClassAccessors()
	generated = uti.ReplaceAll(
		generated,
		"classAccessors",
		classAccessors,
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
	var class = moduleAssemblerClass()
	generated = class.formatImportedPackages(
		existing,
		generated,
	)
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (c *moduleAssemblerClass_) createImportedPath(
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

func (c *moduleAssemblerClass_) extractImportedPackages(
	source string,
) (
	packages fra.CatalogLike[string, string],
) {
	packages = fra.Catalog[string, string]()
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

func (c *moduleAssemblerClass_) formatImportedPackages(
	existing string,
	generated string,
) string {
	// Start with the existing imported packages.
	var imports = c.extractImportedPackages(existing)

	// Add in the generated imported packages.
	imports = fra.CatalogClass[string, string]().Merge(
		imports,
		c.extractImportedPackages(generated),
	)

	// Sort the imported packages by path rather than name.
	imports.SortValuesWithRanker(
		func(first, second fra.AssociationLike[string, string]) fra.Rank {
			var firstValue = first.GetValue()
			var secondValue = second.GetValue()
			switch {
			case firstValue < secondValue:
				return fra.LesserRank
			case firstValue > secondValue:
				return fra.GreaterRank
			default:
				return fra.EqualRank
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

type moduleAssembler_ struct {
	// Declare the instance attributes.
}

// Class Structure

type moduleAssemblerClass_ struct {
	// Declare the class constants.
	moduleTemplate_ string
	importedPath_   string
}

// Class Reference

func moduleAssemblerClass() *moduleAssemblerClass_ {
	return moduleAssemblerClassReference_
}

var moduleAssemblerClassReference_ = &moduleAssemblerClass_{
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
// CLASS ACCESSORS
<ClassAccessors>
// GLOBAL FUNCTIONS
`,

	importedPath_: `
	<~packageAcronym> <packagePath>`,
}
