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

package synthesizer

import (
	ana "github.com/craterdog/go-code-generation/v5/analyzer"
	col "github.com/craterdog/go-collection-framework/v5"
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v5"
	reg "regexp"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ValidatorSynthesizerClass() ValidatorSynthesizerClassLike {
	return validatorSynthesizerClassReference()
}

// Constructor Methods

func (c *validatorSynthesizerClass_) ValidatorSynthesizer(
	syntax not.SyntaxLike,
) ValidatorSynthesizerLike {
	var instance = &validatorSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().SyntaxAnalyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *validatorSynthesizer_) GetClass() ValidatorSynthesizerClassLike {
	return validatorSynthesizerClassReference()
}

// TemplateDriven Methods

func (v *validatorSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *validatorSynthesizer_) CreateWarningMessage() string {
	var class = validatorSynthesizerClassReference()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *validatorSynthesizer_) CreateAccessFunction() string {
	var class = validatorSynthesizerClassReference()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *validatorSynthesizer_) CreateConstructorMethods() string {
	var class = validatorSynthesizerClassReference()
	var constructorMethods = class.constructorMethods_
	return constructorMethods
}

func (v *validatorSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *validatorSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *validatorSynthesizer_) CreatePrincipalMethods() string {
	var class = validatorSynthesizerClassReference()
	var principalMethods = class.principalMethods_
	var syntaxMap = v.analyzer_.GetSyntaxMap()
	principalMethods = uti.ReplaceAll(
		principalMethods,
		"syntaxMap",
		syntaxMap,
	)
	var syntaxName = v.analyzer_.GetSyntaxName()
	principalMethods = uti.ReplaceAll(
		principalMethods,
		"syntaxName",
		syntaxName,
	)
	return principalMethods
}

func (v *validatorSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *validatorSynthesizer_) CreateAspectMethods() string {
	var class = validatorSynthesizerClassReference()
	var aspectMethods = class.aspectMethods_
	var processTokens = v.createProcessTokens()
	aspectMethods = uti.ReplaceAll(
		aspectMethods,
		"processTokens",
		processTokens,
	)
	var processRules = v.createProcessRules()
	aspectMethods = uti.ReplaceAll(
		aspectMethods,
		"processRules",
		processRules,
	)
	return aspectMethods
}

func (v *validatorSynthesizer_) CreatePrivateMethods() string {
	var class = validatorSynthesizerClassReference()
	var privateMethods = class.privateMethods_
	return privateMethods
}

func (v *validatorSynthesizer_) CreateInstanceStructure() string {
	var class = validatorSynthesizerClassReference()
	var instanceStructure = class.instanceStructure_
	return instanceStructure
}

func (v *validatorSynthesizer_) CreateClassStructure() string {
	var class = validatorSynthesizerClassReference()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *validatorSynthesizer_) CreateClassReference() string {
	var class = validatorSynthesizerClassReference()
	var classReference = class.classReference_
	return classReference
}

func (v *validatorSynthesizer_) PerformGlobalUpdates(
	moduleName string,
	packageName string,
	className string,
	existing string,
	generated string,
) string {
	generated = v.preserveExistingCode(existing, generated)
	generated = v.updateImportedPackages(moduleName, existing, generated)
	return generated
}

func (v *validatorSynthesizer_) preserveExistingCode(
	existing string,
	generated string,
) string {
	// Preserve the methodical method implementations.
	var pattern = `// Methodical Methods(.|\r?\n)+// PROTECTED INTERFACE`
	generated = v.replacePattern(pattern, existing, generated)
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (v *validatorSynthesizer_) createImportedPath(
	packageAcronym string,
	packagePath string,
) string {
	var class = validatorSynthesizerClassReference()
	var importedPath = class.importedPath_
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

func (v *validatorSynthesizer_) createProcessRule(
	ruleName string,
) string {
	var class = validatorSynthesizerClassReference()
	var processRule = class.processRule_
	if v.analyzer_.IsPlural(ruleName) {
		processRule = class.processIndexedRule_
	}
	processRule = uti.ReplaceAll(
		processRule,
		"ruleName",
		ruleName,
	)
	return processRule
}

func (v *validatorSynthesizer_) createProcessRules() string {
	var processRules string
	var ruleNames = v.analyzer_.GetRuleNames().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		var processRule = v.createProcessRule(ruleName)
		processRules += processRule
	}
	return processRules
}

func (v *validatorSynthesizer_) createProcessToken(
	tokenName string,
) string {
	var processToken string
	if tokenName == "delimiter" {
		return processToken
	}
	var class = validatorSynthesizerClassReference()
	processToken = class.processToken_
	if v.analyzer_.IsPlural(tokenName) {
		processToken = class.processIndexedToken_
	}
	processToken = uti.ReplaceAll(
		processToken,
		"tokenName",
		tokenName,
	)
	return processToken
}

func (v *validatorSynthesizer_) createProcessTokens() string {
	var processTokens string
	var tokenNames = v.analyzer_.GetTokenNames().GetIterator()
	for tokenNames.HasNext() {
		var tokenName = tokenNames.GetNext()
		var processToken = v.createProcessToken(tokenName)
		processTokens += processToken
	}
	return processTokens
}

func (v *validatorSynthesizer_) extractImportedPackages(
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

func (v *validatorSynthesizer_) replacePattern(
	pattern string,
	existing string,
	generated string,
) string {
	var matcher = reg.MustCompile(pattern)
	var existingPattern = matcher.FindString(existing)
	var generatedPattern = matcher.FindString(generated)
	generated = sts.ReplaceAll(
		generated,
		generatedPattern,
		existingPattern,
	)
	return generated
}

func (v *validatorSynthesizer_) updateImportedPackages(
	moduleName string,
	existing string,
	generated string,
) string {
	// Seed the imported packages with the most common ones.
	var imports = abs.CatalogClass[string, string]().CatalogFromMap(
		map[string]string{
			"fmt": `"fmt"`,
			"ast": `"<ModuleName>/ast"`,
			"col": `"github.com/craterdog/go-collection-framework/v5"`,
			"abs": `"github.com/craterdog/go-collection-framework/v5/collection"`,
			"uti": `"github.com/craterdog/go-missing-utilities/v2"`,
			"ref": `"reflect"`,
			"sts": `"strings"`,
		},
	)

	// Add in the imported packages from the existing code.
	imports = abs.CatalogClass[string, string]().Merge(
		imports, v.extractImportedPackages(existing),
	)
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

	// Create imported package statements for each imported package.
	var importedPackages string
	var packages = imports.GetIterator()
	for packages.HasNext() {
		var association = packages.GetNext()
		var packageAcronym = association.GetKey()
		var packagePath = association.GetValue()
		if sts.Contains(generated, packageAcronym+".") {
			// Only import packages that are actually used in the generated code.
			importedPackages += v.createImportedPath(
				packageAcronym,
				packagePath,
			)
		}
	}
	if uti.IsDefined(importedPackages) {
		importedPackages += "\n"
	}

	// Insert the imported packages into the generated code.
	generated = uti.ReplaceAll(
		generated,
		"importedPackages",
		importedPackages,
	)
	return generated
}

// Instance Structure

type validatorSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type validatorSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_      string
	importedPath_        string
	accessFunction_      string
	constructorMethods_  string
	principalMethods_    string
	aspectMethods_       string
	processToken_        string
	processIndexedToken_ string
	processRule_         string
	processIndexedRule_  string
	privateMethods_      string
	instanceStructure_   string
	classStructure_      string
	classReference_      string
}

// Class Reference

func validatorSynthesizerClassReference() *validatorSynthesizerClass_ {
	return validatorSynthesizerClassReference_
}

var validatorSynthesizerClassReference_ = &validatorSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPath_: `
	<~packageAcronym> <packagePath>`,

	accessFunction_: `
// Access Function

func ValidatorClass() ValidatorClassLike {
	return validatorClassReference()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *validatorClass_) Validator() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}
`,

	principalMethods_: `
// Principal Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorClassReference()
}

func (v *validator_) Validate<~SyntaxName>(
	<syntaxName_> ast.<~SyntaxName>Like,
) {
	v.visitor_.Visit<~SyntaxName>(<syntaxName_>)
}
`,

	aspectMethods_: `
// Methodical Methods
<ProcessTokens><ProcessRules>`,

	processToken_: `
func (v *validator_) Process<~TokenName>(
	<tokenName_> string,
) {
	v.validateToken(<tokenName_>, <~TokenName>Token)
}
`,

	processIndexedToken_: `
func (v *validator_) Process<~TokenName>(
	<tokenName_> string,
	index uint,
	size uint,
) {
	v.validateToken(<tokenName_>, <~TokenName>Token)
}
`,

	processRule_: `
func (v *validator_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {
	// TBD - Add any validation checks.
}
`,

	processIndexedRule_: `
func (v *validator_) Preprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Process<~RuleName>Slot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) Postprocess<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}
`,

	privateMethods_: `
// Private Methods

func (v *validator_) validateToken(
	tokenValue string,
	tokenType TokenType,
) {
	var scannerClass = ScannerClass()
	if !scannerClass.MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			scannerClass.FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}
`,

	instanceStructure_: `
// Instance Structure

type validator_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike

	// Declare the inherited aspects.
	Methodical
}
`,

	classStructure_: `
// Class Structure

type validatorClass_ struct {
	// Declare the class constants.
}
`,

	classReference_: `
// Class Reference

func validatorClassReference() *validatorClass_ {
	return validatorClassReference_
}

var validatorClassReference_ = &validatorClass_{
	// Initialize the class constants.
}
`,
}
