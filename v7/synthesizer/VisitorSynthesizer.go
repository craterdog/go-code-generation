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

package synthesizer

import (
	ana "github.com/craterdog/go-code-generation/v7/analyzer"
	uti "github.com/craterdog/go-missing-utilities/v7"
	not "github.com/craterdog/go-syntax-notation/v7"
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func VisitorSynthesizerClass() VisitorSynthesizerClassLike {
	return visitorSynthesizerClass()
}

// Constructor Methods

func (c *visitorSynthesizerClass_) VisitorSynthesizer(
	syntax not.SyntaxLike,
) VisitorSynthesizerLike {
	var instance = &visitorSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: ana.SyntaxAnalyzerClass().SyntaxAnalyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *visitorSynthesizer_) GetClass() VisitorSynthesizerClassLike {
	return visitorSynthesizerClass()
}

// TemplateDriven Methods

func (v *visitorSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetLegalNotice()
	return legalNotice
}

func (v *visitorSynthesizer_) CreateWarningMessage() string {
	var class = visitorSynthesizerClass()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *visitorSynthesizer_) CreateImportedPackages() string {
	var class = visitorSynthesizerClass()
	var importedPackages = class.importedPackages_
	return importedPackages
}

func (v *visitorSynthesizer_) CreateAccessFunction() string {
	var class = visitorSynthesizerClass()
	var accessFunction = class.accessFunction_
	return accessFunction
}

func (v *visitorSynthesizer_) CreateConstructorMethods() string {
	var class = visitorSynthesizerClass()
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
	var class = visitorSynthesizerClass()
	var principalMethods = class.principalMethods_
	return principalMethods
}

func (v *visitorSynthesizer_) CreateAttributeMethods() string {
	var attributeMethods string
	return attributeMethods
}

func (v *visitorSynthesizer_) CreateAspectInterfaces() string {
	var aspectInterfaces string
	return aspectInterfaces
}

func (v *visitorSynthesizer_) CreatePrivateMethods() string {
	var visitMethods = v.createVisitMethods()
	var class = visitorSynthesizerClass()
	var privateMethods = class.privateMethods_
	privateMethods = uti.ReplaceAll(
		privateMethods,
		"visitMethods",
		visitMethods,
	)
	return privateMethods
}

func (v *visitorSynthesizer_) CreateInstanceStructure() string {
	var class = visitorSynthesizerClass()
	var instanceStructure = class.instanceStructure_
	return instanceStructure
}

func (v *visitorSynthesizer_) CreateClassStructure() string {
	var class = visitorSynthesizerClass()
	var classStructure = class.classStructure_
	return classStructure
}

func (v *visitorSynthesizer_) CreateClass() string {
	var class = visitorSynthesizerClass()
	var classReference = class.classReference_
	return classReference
}

func (v *visitorSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	var syntaxName = v.analyzer_.GetSyntaxName()
	generated = uti.ReplaceAll(
		generated,
		"syntaxName",
		syntaxName,
	)
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitorSynthesizer_) createTermSequence(
	ruleName string,
) string {
	var implementation string
	var terms = v.analyzer_.GetRuleTerms(ruleName).GetIterator()
	var variables = v.analyzer_.GetVariables(ruleName).GetIterator()
	for terms.HasNext() && variables.HasNext() {
		var slot = uint(terms.GetSlot())
		if slot > 0 {
			implementation += v.createInlineSlot(ruleName, slot)
		}
		var term = terms.GetNext()
		var variableName = variables.GetNext()
		implementation += v.createInlineTerm(term, variableName)
	}
	return implementation
}

func (v *visitorSynthesizer_) createInlineTerm(
	term not.RuleTermLike,
	variableName string,
) string {
	var implementation string
	var cardinality = term.GetOptionalCardinality()
	var actual = term.GetComponent().GetAny().(string)
	switch {
	case not.MatchesType(actual, not.LiteralToken):
		implementation = v.createLiteralTerm(actual, cardinality)
	case not.MatchesType(actual, not.LowercaseToken):
		implementation = v.createExpressionTerm(actual, cardinality)
	case not.MatchesType(actual, not.UppercaseToken):
		implementation = v.createRuleName(actual, cardinality)
	}
	implementation = uti.ReplaceAll(
		implementation,
		"variableName",
		variableName,
	)
	return implementation
}

func (v *visitorSynthesizer_) createCardinality(
	cardinality not.CardinalityLike,
	optionalCardinality string,
	requiredCardinality string,
	repeatedCardinality string,
) string {
	var unlimited = "mat.MaxInt"
	var first string
	var last = unlimited
	var actualCardinality = requiredCardinality
	if uti.IsUndefined(cardinality) {
		return actualCardinality
	}
	switch actual := cardinality.GetAny().(type) {
	case not.ConstrainedLike:
		actualCardinality = repeatedCardinality
		switch actual.GetAny().(string) {
		case "?":
			// This is the "{0..1}" case.
			first = "0"
			last = "1"
			actualCardinality = optionalCardinality
		case "*":
			// This is the "{0..}" case.
			first = "0"
		case "+":
			// This is the "{1..}" case.
			first = "1"
		}
	case not.QuantifiedLike:
		first = actual.GetNumber()
		var limit = actual.GetOptionalLimit()
		if uti.IsUndefined(limit) {
			// This is the "{m}" case.
			last = first
		} else {
			last = limit.GetOptionalNumber()
			if uti.IsUndefined(last) {
				// This is the "{m..}" case.
				last = unlimited
			}
			// This is the "{m..n}" case.
		}
	}
	actualCardinality = uti.ReplaceAll(
		actualCardinality,
		"first",
		first,
	)
	actualCardinality = uti.ReplaceAll(
		actualCardinality,
		"last",
		last,
	)
	return actualCardinality
}

func (v *visitorSynthesizer_) createLiteralTerm(
	literal string,
	cardinality not.CardinalityLike,
) string {
	var class = visitorSynthesizerClass()
	var optionalTerm = class.optionalLiteral_
	var requiredTerm = class.requiredLiteral_
	var repeatedTerm = class.repeatedLiteral_
	var implementation = v.createCardinality(
		cardinality,
		optionalTerm,
		requiredTerm,
		repeatedTerm,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"literal",
		literal,
	)
	return implementation
}

func (v *visitorSynthesizer_) createExpressionTerm(
	lowercase string,
	cardinality not.CardinalityLike,
) string {
	var class = visitorSynthesizerClass()
	var optionalTerm = class.optionalExpression_
	var requiredTerm = class.requiredExpression_
	var repeatedTerm = class.repeatedExpression_
	var implementation = v.createCardinality(
		cardinality,
		optionalTerm,
		requiredTerm,
		repeatedTerm,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"lowercase",
		lowercase,
	)
	return implementation
}

func (v *visitorSynthesizer_) createRuleName(
	uppercase string,
	cardinality not.CardinalityLike,
) string {
	var class = visitorSynthesizerClass()
	var optionalTerm = class.optionalRule_
	var requiredTerm = class.requiredRule_
	var repeatedTerm = class.repeatedRule_
	var implementation = v.createCardinality(
		cardinality,
		optionalTerm,
		requiredTerm,
		repeatedTerm,
	)
	implementation = uti.ReplaceAll(
		implementation,
		"uppercase",
		uppercase,
	)
	return implementation
}

func (v *visitorSynthesizer_) createInlineSlot(
	ruleName string,
	slot uint,
) string {
	var class = visitorSynthesizerClass()
	var inlineSlot = class.slot_
	inlineSlot = uti.ReplaceAll(
		inlineSlot,
		"ruleName",
		ruleName,
	)
	inlineSlot = uti.ReplaceAll(
		inlineSlot,
		"slot",
		stc.Itoa(int(slot)),
	)
	return inlineSlot
}

func (v *visitorSynthesizer_) createLiteralAlternatives(
	ruleName string,
) string {
	var class = visitorSynthesizerClass()
	var literalValues string
	var iterator = v.analyzer_.GetLiteralValues(ruleName).GetIterator()
	for iterator.HasNext() {
		var literal = iterator.GetNext().GetLiteral()
		var literalValue = class.literalValue_
		literalValue = uti.ReplaceAll(
			literalValue,
			"literal",
			literal,
		)
		literalValues += literalValue
	}
	var implementation = class.literalAlternatives_
	implementation = uti.ReplaceAll(
		implementation,
		"literalValues",
		literalValues,
	)
	return implementation
}

func (v *visitorSynthesizer_) createExpressionAlternatives(
	ruleName string,
) string {
	var class = visitorSynthesizerClass()
	var implementation = class.expressionAlternatives_
	var expressionAlternatives string
	var expressionNames = v.analyzer_.GetExpressionNames(ruleName).GetIterator()
	for expressionNames.HasNext() {
		var expressionName = expressionNames.GetNext()
		var lowercase = expressionName.GetLowercase()
		var expression = class.expressionName_
		expression = uti.ReplaceAll(
			expression,
			"lowercase",
			lowercase,
		)
		expressionAlternatives += expression
	}
	implementation = uti.ReplaceAll(
		implementation,
		"expressionNames",
		expressionAlternatives,
	)
	return implementation
}

func (v *visitorSynthesizer_) createRuleAlternatives(
	ruleName string,
) string {
	var class = visitorSynthesizerClass()
	var implementation = class.ruleAlternatives_
	var ruleAlternatives string
	var ruleNames = v.analyzer_.GetRuleNames(ruleName).GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		var uppercase = ruleName.GetUppercase()
		var rule = class.ruleName_
		rule = uti.ReplaceAll(
			rule,
			"uppercase",
			uppercase,
		)
		ruleAlternatives += rule
	}
	implementation = uti.ReplaceAll(
		implementation,
		"ruleNames",
		ruleAlternatives,
	)
	return implementation
}

func (v *visitorSynthesizer_) createVisitMethods() string {
	var visitMethods string
	var ruleNames = v.analyzer_.GetRules().GetIterator()
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
	var definition = v.analyzer_.GetDefinitions().GetValue(ruleName)
	switch definition.GetAny().(type) {
	case not.LiteralAlternativesLike:
		methodImplementation = v.createLiteralAlternatives(ruleName)
	case not.ExpressionAlternativesLike:
		methodImplementation = v.createExpressionAlternatives(ruleName)
	case not.RuleAlternativesLike:
		methodImplementation = v.createRuleAlternatives(ruleName)
	case not.TermSequenceLike:
		methodImplementation = v.createTermSequence(ruleName)
	}
	var class = visitorSynthesizerClass()
	var visitMethod = class.visitMethod_
	visitMethod = uti.ReplaceAll(
		visitMethod,
		"methodImplementation",
		methodImplementation,
	)
	visitMethod = uti.ReplaceAll(
		visitMethod,
		"ruleName",
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
	warningMessage_         string
	importedPackages_       string
	accessFunction_         string
	constructorMethods_     string
	principalMethods_       string
	privateMethods_         string
	visitMethod_            string
	literalAlternatives_    string
	literalValue_           string
	expressionAlternatives_ string
	expressionName_         string
	ruleAlternatives_       string
	ruleName_               string
	optionalLiteral_        string
	requiredLiteral_        string
	repeatedLiteral_        string
	optionalExpression_     string
	requiredExpression_     string
	repeatedExpression_     string
	optionalRule_           string
	requiredRule_           string
	repeatedRule_           string
	slot_                   string
	instanceStructure_      string
	classStructure_         string
	classReference_         string
}

// Class Reference

func visitorSynthesizerClass() *visitorSynthesizerClass_ {
	return visitorSynthesizerClassReference_
}

var visitorSynthesizerClassReference_ = &visitorSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPackages_: `
	fmt "fmt"
	ast "<ModuleName>/ast"
	uti "github.com/craterdog/go-missing-utilities/v7"
`,

	accessFunction_: `
// Access Function

func VisitorClass() VisitorClassLike {
	return visitorClass()
}
`,

	constructorMethods_: `
// Constructor Methods

func (c *visitorClass_) Visitor(
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
	return visitorClass()
}

func (v *visitor_) Visit<~SyntaxName>(
	<syntaxName_> ast.<~SyntaxName>Like,
) {
	v.processor_.Preprocess<~SyntaxName>(
		<syntaxName_>,
		1,
		1,
	)
	v.visit<~SyntaxName>(<syntaxName_>)
	v.processor_.Postprocess<~SyntaxName>(
		<syntaxName_>,
		1,
		1,
	)
}
`,

	privateMethods_: `
// Private Methods
<VisitMethods>`,

	visitMethod_: `
func (v *visitor_) visit<~RuleName>(
	<ruleName_> ast.<~RuleName>Like,
) {<MethodImplementation>}
`,

	literalAlternatives_: `
	// Visit the possible <~ruleName> literal values.
	var actual = <ruleName_>.GetAny().(string)
	switch actual {<LiteralValues>}
`,

	literalValue_: `
	case <literal>:
		v.processor_.ProcessDelimiter(<literal>)`,

	expressionAlternatives_: `
	// Visit the possible <~ruleName> expression types.
	var actual = <ruleName_>.GetAny().(string)
	switch {<ExpressionNames>}
`,

	expressionName_: `
	case ScannerClass().MatchesType(actual, <~Lowercase>Token):
		v.processor_.Process<~Lowercase>(actual)`,

	ruleAlternatives_: `
	// Visit the possible <~ruleName> rule types.
	switch actual := <ruleName_>.GetAny().(type) {<RuleNames>}
`,

	ruleName_: `
	case ast.<~Uppercase>Like:
		v.processor_.Preprocess<~Uppercase>(
			actual,
			1,
			1,
		)
		v.visit<~Uppercase>(actual)
		v.processor_.Postprocess<~Uppercase>(
			actual,
			1,
			1,
		)`,

	optionalLiteral_: `
	var <variableName_> = <ruleName_>.Get<~VariableName>()
	if uti.IsDefined(<variableName_>) {
		v.processor_.ProcessDelimiter(<variableName_>)
	}`,

	requiredLiteral_: `
	var <variableName_> = <ruleName_>.Get<~VariableName>()
	v.processor_.ProcessDelimiter(<variableName_>)`,

	repeatedLiteral_: `
	var <variableName_> = <ruleName_>.Get<~VariableName>().GetIterator()
	for <variableName_>.HasNext() {
		var literal = <variableName_>.GetNext()
		v.processor_.ProcessDelimiter(literal)
	}`,

	optionalExpression_: `
	var <variableName_> = <ruleName_>.Get<~VariableName>()
	if uti.IsDefined(<variableName_>) {
		v.processor_.Process<~Lowercase>(<variableName_>)
	}`,

	requiredExpression_: `
	var <variableName_> = <ruleName_>.Get<~VariableName>()
	v.processor_.Process<~Lowercase>(<variableName_>)`,

	repeatedExpression_: `
	var <variableName_> = <ruleName_>.Get<~VariableName>().GetIterator()
	for <variableName_>.HasNext() {
		var expression = <variableName_>.GetNext()
		v.processor_.Process<~Lowercase>(expression)
	}`,

	optionalRule_: `
	var <variableName_> = <ruleName_>.Get<~VariableName>()
	if uti.IsDefined(<variableName_>) {
		v.processor_.Preprocess<~Uppercase>(
			<variableName_>,
			1,
			1,
		)
		v.visit<~Uppercase>(<variableName_>)
		v.processor_.Postprocess<~Uppercase>(
			<variableName_>,
			1,
			1,
		)
	}`,

	requiredRule_: `
	var <variableName_> = <ruleName_>.Get<~VariableName>()
	v.processor_.Preprocess<~Uppercase>(
		<variableName_>,
		1,
		1,
	)
	v.visit<~Uppercase>(<variableName_>)
	v.processor_.Postprocess<~Uppercase>(
		<variableName_>,
		1,
		1,
	)`,

	repeatedRule_: `
	var <~variableName>Index uint
	var <variableName_> = <ruleName_>.Get<~VariableName>().GetIterator()
	var <~variableName>Count = uint(<variableName_>.GetSize())
	for <variableName_>.HasNext() {
		<~variableName>Index++
		var rule = <variableName_>.GetNext()
		v.processor_.Preprocess<~Uppercase>(
			rule,
			<~variableName>Index,
			<~variableName>Count,
		)
		v.visit<~Uppercase>(rule)
		v.processor_.Postprocess<~Uppercase>(
			rule,
			<~variableName>Index,
			<~variableName>Count,
		)
	}`,

	slot_: `
	// Visit slot <slot> between terms.
	v.processor_.Process<~RuleName>Slot(
		<ruleName_>,
		<slot>,
	)
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

func visitorClass() *visitorClass_ {
	return visitorClassReference_
}

var visitorClassReference_ = &visitorClass_{
	// Initialize the class constants.
}
`,
}
