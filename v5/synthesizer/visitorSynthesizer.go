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

func VisitorSynthesizer() VisitorSynthesizerClassLike {
	return visitorSynthesizerReference()
}

// Constructor Methods

func (c *visitorSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) VisitorSynthesizerLike {
	var instance = &visitorSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzer().Make(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *visitorSynthesizer_) GetClass() VisitorSynthesizerClassLike {
	return visitorSynthesizerReference()
}

// TemplateDriven Methods

func (v *visitorSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *visitorSynthesizer_) CreateAccessFunction() string {
	var accessFunction = visitorSynthesizerReference().accessFunction_
	return accessFunction
}

func (v *visitorSynthesizer_) CreateConstructorMethods() string {
	var constructorMethods = visitorSynthesizerReference().constructorMethods_
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

func (v *visitorSynthesizer_) CreatePrimaryMethods() string {
	var primaryMethods = visitorSynthesizerReference().primaryMethods_
	return primaryMethods
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
	var privateMethods = visitorSynthesizerReference().privateMethods_
	privateMethods = uti.ReplaceAll(privateMethods, "visitMethods", visitMethods)
	return privateMethods
}

func (v *visitorSynthesizer_) CreateInstanceStructure() string {
	var instanceStructure = visitorSynthesizerReference().instanceStructure_
	return instanceStructure
}

func (v *visitorSynthesizer_) CreateClassStructure() string {
	var classStructure = visitorSynthesizerReference().classStructure_
	return classStructure
}

func (v *visitorSynthesizer_) CreateClassReference() string {
	var classReference = visitorSynthesizerReference().classReference_
	return classReference
}

func (v *visitorSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var syntaxName = v.analyzer_.GetSyntaxName()
	source = uti.ReplaceAll(source, "syntaxName", syntaxName)
	var classImports = visitorSynthesizerReference().classImports_
	source = uti.ReplaceAll(source, "classImports", classImports)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitorSynthesizer_) createInlineImplementation(
	ruleName string,
) (
	implementation string,
) {
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
) (
	implementation string,
) {
	var identifier = reference.GetIdentifier().GetAny().(string)
	switch {
	case not.MatchesType(identifier, not.LowercaseToken):
		implementation = v.createInlineToken(reference, variableName)
	case not.MatchesType(identifier, not.UppercaseToken):
		implementation = v.createInlineRule(reference, variableName)
	}
	return implementation
}

func (v *visitorSynthesizer_) createInlineRule(
	reference not.ReferenceLike,
	variableName string,
) (
	implementation string,
) {
	switch v.createPlurality(reference) {
	case "singular":
		implementation = visitorSynthesizerReference().singularRuleBlock_
	case "optional":
		implementation = visitorSynthesizerReference().optionalRuleBlock_
	case "repeated":
		implementation = visitorSynthesizerReference().repeatedRuleBlock_
	default:
		implementation = visitorSynthesizerReference().ruleBlock_
	}
	var ruleName = reference.GetIdentifier().GetAny().(string)
	implementation = uti.ReplaceAll(implementation, "ruleName", ruleName)
	implementation = uti.ReplaceAll(implementation, "variableName", variableName)
	return implementation
}

func (v *visitorSynthesizer_) createInlineSlot(
	ruleName string,
	slot uint,
) (
	implementation string,
) {
	implementation = visitorSynthesizerReference().slotBlock_
	implementation = uti.ReplaceAll(implementation, "slot", stc.Itoa(int(slot)))
	return implementation
}

func (v *visitorSynthesizer_) createInlineToken(
	reference not.ReferenceLike,
	variableName string,
) (
	implementation string,
) {
	switch v.createPlurality(reference) {
	case "singular":
		implementation = visitorSynthesizerReference().singularTokenBlock_
	case "optional":
		implementation = visitorSynthesizerReference().optionalTokenBlock_
	case "repeated":
		implementation = visitorSynthesizerReference().repeatedTokenBlock_
	default:
		implementation = visitorSynthesizerReference().tokenBlock_
	}
	var tokenName = reference.GetIdentifier().GetAny().(string)
	implementation = uti.ReplaceAll(implementation, "tokenName", tokenName)
	implementation = uti.ReplaceAll(implementation, "variableName", variableName)
	return implementation
}

func (v *visitorSynthesizer_) createMultilineImplementation(
	ruleName string,
) (
	implementation string,
) {
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
	implementation = visitorSynthesizerReference().multilineCases_
	implementation = uti.ReplaceAll(implementation, "ruleCases", ruleCases)
	implementation = uti.ReplaceAll(implementation, "tokenCases", tokenCases)
	return implementation
}

func (v *visitorSynthesizer_) createMultilineRule(
	ruleName string,
) (
	implementation string,
) {
	implementation = visitorSynthesizerReference().ruleCase_
	if v.analyzer_.IsPlural(ruleName) {
		implementation = visitorSynthesizerReference().singularRuleCase_
	}
	implementation = uti.ReplaceAll(
		implementation,
		"ruleName",
		ruleName,
	)
	return implementation
}

func (v *visitorSynthesizer_) createMultilineToken(
	tokenName string,
) (
	implementation string,
) {
	implementation = visitorSynthesizerReference().tokenCase_
	if v.analyzer_.IsPlural(tokenName) {
		implementation = visitorSynthesizerReference().singularTokenCase_
	}
	implementation = uti.ReplaceAll(
		implementation,
		"tokenName",
		tokenName,
	)
	return implementation
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
) (
	implementation string,
) {
	var methodImplementation string
	switch {
	case uti.IsDefined(v.analyzer_.GetIdentifiers(ruleName)):
		methodImplementation = v.createMultilineImplementation(ruleName)
	case uti.IsDefined(v.analyzer_.GetReferences(ruleName)):
		methodImplementation = v.createInlineImplementation(ruleName)
	}
	implementation = visitorSynthesizerReference().visitMethod_
	implementation = uti.ReplaceAll(
		implementation,
		"methodImplementation",
		methodImplementation,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"targetName",
		ruleName,
	)
	return implementation
}

// Instance Structure

type visitorSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ ana.SyntaxAnalyzerLike
}

// Class Structure

type visitorSynthesizerClass_ struct {
	// Declare the class constants.
	classImports_       string
	accessFunction_     string
	constructorMethods_ string
	primaryMethods_     string
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

func visitorSynthesizerReference() *visitorSynthesizerClass_ {
	return visitorSynthesizerReference_
}

var visitorSynthesizerReference_ = &visitorSynthesizerClass_{
	// Initialize the class constants.
	classImports_: `

import (
	fmt "fmt"
	ast "<ModuleName>/ast"
	uti "github.com/craterdog/go-missing-utilities/v2"
)`,

	accessFunction_: `
// Access Function

func Visitor() VisitorClassLike {
	return visitorReference()
}`,

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
}`,

	primaryMethods_: `
// Primary Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorReference()
}

func (v *visitor_) Visit<~SyntaxName>(
	<syntaxName_> ast.<~SyntaxName>Like,
) {
	v.processor_.Preprocess<~SyntaxName>(<syntaxName_>)
	v.visit<~SyntaxName>(<syntaxName_>)
	v.processor_.Postprocess<~SyntaxName>(<syntaxName_>)
}`,

	privateMethods_: `
// Private Methods<VisitMethods>`,

	visitMethod_: `

func (v *visitor_) visit<~TargetName>(<targetName_> ast.<~TargetName>Like) {<MethodImplementation>}`,

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
		case Scanner().MatchesType(actual, <~TokenName>Token):
			v.processor_.Process<~TokenName>(actual)`,

	singularTokenCase_: `
	case Scanner().MatchesType(actual, <~TokenName>Token):
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
}`,

	classStructure_: `
// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}`,

	classReference_: `
// Class Reference

func visitorReference() *visitorClass_ {
	return visitorReference_
}

var visitorReference_ = &visitorClass_{
	// Initialize the class constants.
}`,
}
