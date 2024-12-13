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
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v5"
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func VisitorSynthesizerClass() VisitorSynthesizerClassLike {
	return visitorSynthesizerClassReference()
}

// Constructor Methods

func (c *visitorSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) VisitorSynthesizerLike {
	var instance = &visitorSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *visitorSynthesizer_) GetClass() VisitorSynthesizerClassLike {
	return visitorSynthesizerClassReference()
}

// TemplateDriven Methods

func (v *visitorSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *visitorSynthesizer_) CreateWarningMessage() string {
	var class = visitorSynthesizerClassReference()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *visitorSynthesizer_) CreateAccessFunction() string {
	var class = visitorSynthesizerClassReference()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *visitorSynthesizer_) CreateConstructorMethods() string {
	var class = visitorSynthesizerClassReference()
	var constructorMethods = class.constructorMethods_
	return constructorMethods
}

func (v *visitorSynthesizer_) CreateConstantMethods() string {
	var constantMethods string
	return constantMethods
}

func (v *visitorSynthesizer_) CreateFunctionMethods() string {
	var functionMethods string
	return functionMethods
}

func (v *visitorSynthesizer_) CreatePrincipalMethods() string {
	var class = visitorSynthesizerClassReference()
	var principalMethods = class.principalMethods_
	return principalMethods
}

func (v *visitorSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *visitorSynthesizer_) CreateAspectMethods() string {
	var aspectMethods string
	return aspectMethods
}

func (v *visitorSynthesizer_) CreatePrivateMethods() string {
	var visitMethods = v.createVisitMethods()
	var class = visitorSynthesizerClassReference()
	var privateMethods = class.privateMethods_
	privateMethods = uti.ReplaceAll(
		privateMethods,
		"visitMethods",
		visitMethods,
	)
	return privateMethods
}

func (v *visitorSynthesizer_) CreateInstanceStructure() string {
	var class = visitorSynthesizerClassReference()
	var instanceStructure = class.instanceStructure_
	return instanceStructure
}

func (v *visitorSynthesizer_) CreateClassStructure() string {
	var class = visitorSynthesizerClassReference()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *visitorSynthesizer_) CreateClassReference() string {
	var class = visitorSynthesizerClassReference()
	var classReference = class.classReference_
	return classReference
}

func (v *visitorSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var syntaxName = v.analyzer_.GetSyntaxName()
	source = uti.ReplaceAll(
		source,
		"syntaxName",
		syntaxName,
	)
	var class = visitorSynthesizerClassReference()
	var importedPackages = class.importedPackages_
	source = uti.ReplaceAll(
		source,
		"importedPackages",
		importedPackages,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitorSynthesizer_) createInlineImplementation(
	ruleName string,
) string {
	var implementation string
	var references = v.analyzer_.GetReferences(ruleName).GetIterator()
	var variables = v.analyzer_.GetVariables(ruleName).GetIterator()
	for references.HasNext() && variables.HasNext() {
		var slot uint = uint(references.GetSlot())
		if slot > 0 {
			implementation += v.createInlineSlot(ruleName, slot)
		}
		var reference = references.GetNext()
		var variableName = variables.GetNext()
		implementation += v.createInlineReference(reference, variableName)
	}
	return implementation
}

func (v *visitorSynthesizer_) createInlineReference(
	reference not.ReferenceLike,
	variableName string,
) string {
	var inlineReference string
	var identifier = reference.GetIdentifier().GetAny().(string)
	switch {
	case not.MatchesType(identifier, not.LowercaseToken):
		inlineReference = v.createInlineToken(reference, variableName)
	case not.MatchesType(identifier, not.UppercaseToken):
		inlineReference = v.createInlineRule(reference, variableName)
	}
	return inlineReference
}

func (v *visitorSynthesizer_) createInlineRule(
	reference not.ReferenceLike,
	variableName string,
) string {
	var inlineRule string
	var class = visitorSynthesizerClassReference()
	switch v.createPlurality(reference) {
	case "singular":
		inlineRule = class.singularRuleBlock_
	case "optional":
		inlineRule = class.optionalRuleBlock_
	case "repeated":
		inlineRule = class.repeatedRuleBlock_
	default:
		inlineRule = class.ruleBlock_
	}
	var ruleName = reference.GetIdentifier().GetAny().(string)
	inlineRule = uti.ReplaceAll(
		inlineRule,
		"ruleName",
		ruleName,
	)
	inlineRule = uti.ReplaceAll(
		inlineRule,
		"variableName",
		variableName,
	)
	return inlineRule
}

func (v *visitorSynthesizer_) createInlineSlot(
	ruleName string,
	slot uint,
) string {
	var class = visitorSynthesizerClassReference()
	var inlineSlot = class.slotBlock_
	inlineSlot = uti.ReplaceAll(
		inlineSlot,
		"slot",
		stc.Itoa(int(slot)),
	)
	return inlineSlot
}

func (v *visitorSynthesizer_) createInlineToken(
	reference not.ReferenceLike,
	variableName string,
) string {
	var inlineToken string
	var class = visitorSynthesizerClassReference()
	switch v.createPlurality(reference) {
	case "singular":
		inlineToken = class.singularTokenBlock_
	case "optional":
		inlineToken = class.optionalTokenBlock_
	case "repeated":
		inlineToken = class.repeatedTokenBlock_
	default:
		inlineToken = class.tokenBlock_
	}
	var tokenName = reference.GetIdentifier().GetAny().(string)
	inlineToken = uti.ReplaceAll(
		inlineToken,
		"tokenName",
		tokenName,
	)
	inlineToken = uti.ReplaceAll(
		inlineToken,
		"variableName",
		variableName,
	)
	return inlineToken
}

func (v *visitorSynthesizer_) createMultilineImplementation(
	ruleName string,
) string {
	var implementation string
	var tokenCases, ruleCases string
	var identifiers = v.analyzer_.GetIdentifiers(ruleName).GetIterator()
	for identifiers.HasNext() {
		var identifier = identifiers.GetNext().GetAny().(string)
		switch {
		case not.MatchesType(identifier, not.LowercaseToken):
			tokenCases += v.createMultilineToken(identifier)
		case not.MatchesType(identifier, not.UppercaseToken):
			ruleCases += v.createMultilineRule(identifier)
		}
	}
	var class = visitorSynthesizerClassReference()
	implementation = class.multilineCases_
	implementation = uti.ReplaceAll(
		implementation,
		"ruleCases",
		ruleCases,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"tokenCases",
		tokenCases,
	)
	return implementation
}

func (v *visitorSynthesizer_) createMultilineRule(
	ruleName string,
) string {
	var class = visitorSynthesizerClassReference()
	var multilineRule = class.ruleCase_
	if v.analyzer_.IsPlural(ruleName) {
		multilineRule = class.singularRuleCase_
	}
	multilineRule = uti.ReplaceAll(
		multilineRule,
		"ruleName",
		ruleName,
	)
	return multilineRule
}

func (v *visitorSynthesizer_) createMultilineToken(
	tokenName string,
) string {
	var class = visitorSynthesizerClassReference()
	var multilineToken = class.tokenCase_
	if v.analyzer_.IsPlural(tokenName) {
		multilineToken = class.singularTokenCase_
	}
	multilineToken = uti.ReplaceAll(
		multilineToken,
		"tokenName",
		tokenName,
	)
	return multilineToken
}

func (v *visitorSynthesizer_) createPlurality(
	reference not.ReferenceLike,
) (
	plurality string,
) {
	var name = reference.GetIdentifier().GetAny().(string)
	var cardinality = reference.GetOptionalCardinality()
	if uti.IsUndefined(cardinality) {
		if v.analyzer_.IsPlural(name) {
			plurality = "singular"
		}
		return plurality
	}
	switch actual := cardinality.GetAny().(type) {
	case not.ConstrainedLike:
		var token = actual.GetAny().(string)
		switch {
		case not.MatchesType(token, not.OptionalToken):
			plurality = "optional"
			if v.analyzer_.IsPlural(name) {
				plurality = "singular"
			}
		case not.MatchesType(token, not.RepeatedToken):
			plurality = "repeated"
		}
	case not.QuantifiedLike:
		plurality = "repeated"
	}
	return plurality
}

func (v *visitorSynthesizer_) createVisitMethods() string {
	var visitMethods string
	var ruleNames = v.analyzer_.GetRuleNames().GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		var visitMethod = v.createVisitMethod(ruleName)
		visitMethods += visitMethod
	}
	return visitMethods
}

func (v *visitorSynthesizer_) createVisitMethod(
	ruleName string,
) string {
	var methodImplementation string
	switch {
	case uti.IsDefined(v.analyzer_.GetIdentifiers(ruleName)):
		methodImplementation = v.createMultilineImplementation(ruleName)
	case uti.IsDefined(v.analyzer_.GetReferences(ruleName)):
		methodImplementation = v.createInlineImplementation(ruleName)
	}
	var class = visitorSynthesizerClassReference()
	var visitMethod = class.visitMethod_
	visitMethod = uti.ReplaceAll(
		visitMethod,
		"methodImplementation",
		methodImplementation,
	)
	visitMethod = uti.ReplaceAll(
		visitMethod,
		"targetName",
		ruleName,
	)
	return visitMethod
}

// Instance Structure

type visitorSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type visitorSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_     string
	importedPackages_   string
	accessFunction_     string
	constructorMethods_ string
	principalMethods_   string
	privateMethods_     string
	visitMethod_        string
	multilineCases_     string
	ruleCase_           string
	singularRuleCase_   string
	tokenCase_          string
	singularTokenCase_  string
	ruleBlock_          string
	singularRuleBlock_  string
	optionalRuleBlock_  string
	repeatedRuleBlock_  string
	tokenBlock_         string
	singularTokenBlock_ string
	optionalTokenBlock_ string
	repeatedTokenBlock_ string
	slotBlock_          string
	instanceStructure_  string
	classStructure_     string
	classReference_     string
}

// Class Reference

func visitorSynthesizerClassReference() *visitorSynthesizerClass_ {
	return visitorSynthesizerClassReference_
}

var visitorSynthesizerClassReference_ = &visitorSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPackages_: `
	fmt "fmt"
	ast "<ModuleName>/ast"
	uti "github.com/craterdog/go-missing-utilities/v2"
`,

	accessFunction_: `
// Access Function

func VisitorClass() VisitorClassLike {
	return visitorClassReference()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *visitorClass_) Make(
	processor Methodical,
) VisitorLike {
	if uti.IsUndefined(processor) {
		panic("The \"processor\" attribute is required by this class.")
	}
	var instance = &visitor_{
		// Initialize the instance attributes.
		processor_: processor,
	}
	return instance
}
`,

	principalMethods_: `
// Principal Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorClassReference()
}

func (v *visitor_) Visit<~SyntaxName>(
	<syntaxName_> ast.<~SyntaxName>Like,
) {
	v.processor_.Preprocess<~SyntaxName>(<syntaxName_>)
	v.visit<~SyntaxName>(<syntaxName_>)
	v.processor_.Postprocess<~SyntaxName>(<syntaxName_>)
}
`,

	privateMethods_: `
// Private Methods
<VisitMethods>`,

	visitMethod_: `
func (v *visitor_) visit<~TargetName>(
	<targetName_> ast.<~TargetName>Like,
) {<MethodImplementation>}
`,

	multilineCases_: `
	// Visit the possible <~targetName> types.
	switch actual := <targetName_>.GetAny().(type) {<RuleCases>
	case string:
		switch {<TokenCases>
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
`,

	ruleCase_: `
	case ast.<~RuleName>Like:
		v.processor_.Preprocess<~RuleName>(actual)
		v.visit<~RuleName>(actual)
		v.processor_.Postprocess<~RuleName>(actual)`,

	singularRuleCase_: `
	case ast.<~RuleName>Like:
		v.processor_.Preprocess<~RuleName>(actual, 1, 1)
		v.visit<~RuleName>(actual)
		v.processor_.Postprocess<~RuleName>(actual, 1, 1)`,

	tokenCase_: `
		case ScannerClass().MatchesType(actual, <~TokenName>Token):
			v.processor_.Process<~TokenName>(actual)`,

	singularTokenCase_: `
	case ClassScanner().MatchesType(actual, <~TokenName>Token):
			v.processor_.Process<~TokenName>(actual, 1, 1)`,

	ruleBlock_: `
	// Visit a single <~ruleName> rule.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	v.processor_.Preprocess<~RuleName>(<variableName_>)
	v.visit<~RuleName>(<variableName_>)
	v.processor_.Postprocess<~RuleName>(<variableName_>)
`,

	singularRuleBlock_: `
	// Visit a single <~ruleName> rule.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	if uti.IsDefined(<variableName_>) {
		v.processor_.Preprocess<~RuleName>(<variableName_>, 1, 1)
		v.visit<~RuleName>(<variableName_>)
		v.processor_.Postprocess<~RuleName>(<variableName_>, 1, 1)
	}
`,

	optionalRuleBlock_: `
	// Visit an optional <~ruleName> rule.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	if uti.IsDefined(<variableName_>) {
		v.processor_.Preprocess<~RuleName>(<variableName_>)
		v.visit<~RuleName>(<variableName_>)
		v.processor_.Postprocess<~RuleName>(<variableName_>)
	}
`,

	repeatedRuleBlock_: `
	// Visit each <~ruleName> rule.
	var <~ruleName>Index uint
	var <variableName_> = <targetName_>.Get<~VariableName>().GetIterator()
	var <variableName>Size = uint(<variableName_>.GetSize())
	for <variableName_>.HasNext() {
		<~ruleName>Index++
		var <ruleName_> = <variableName_>.GetNext()
		v.processor_.Preprocess<~RuleName>(
			<ruleName_>,
			<~ruleName>Index,
			<variableName>Size,
		)
		v.visit<~RuleName>(<ruleName_>)
		v.processor_.Postprocess<~RuleName>(
			<ruleName_>,
			<~ruleName>Index,
			<variableName>Size,
		)
	}
`,

	tokenBlock_: `
	// Visit a single <~tokenName> token.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	v.processor_.Process<~TokenName>(<variableName_>)
`,

	singularTokenBlock_: `
	// Visit a single <~tokenName> token.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	if uti.IsDefined(<variableName_>) {
		v.processor_.Process<~TokenName>(<variableName_>, 1, 1)
	}
`,

	optionalTokenBlock_: `
	// Visit an optional <~tokenName> token.
	var <variableName_> = <targetName_>.Get<~VariableName>()
	if uti.IsDefined(<variableName_>) {
		v.processor_.Process<~TokenName>(<variableName_>)
	}
`,

	repeatedTokenBlock_: `
	// Visit each <~tokenName> token.
	var <~tokenName>Index uint
	var <variableName_> = <targetName_>.Get<~VariableName>().GetIterator()
	var <variableName>Size = uint(<variableName_>.GetSize())
	for <variableName_>.HasNext() {
		<~tokenName>Index++
		var <tokenName_> = <variableName_>.GetNext()
		v.processor_.Process<~TokenName>(
			<tokenName_>,
			<~tokenName>Index,
			<variableName>Size,
		)
	}
`,

	slotBlock_: `
	// Visit slot <Slot> between references.
	v.processor_.Process<~TargetName>Slot(<Slot>)
`,

	instanceStructure_: `
// Instance Structure

type visitor_ struct {
	// Declare the instance attributes.
	processor_ Methodical
}
`,

	classStructure_: `
// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}
`,

	classReference_: `
// Class Reference

func visitorClassReference() *visitorClass_ {
	return visitorClassReference_
}

var visitorClassReference_ = &visitorClass_{
	// Initialize the class constants.
}
`,
}
