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

package analyzer

import (
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	not "github.com/craterdog/go-syntax-notation/v7"
	stc "strconv"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func SyntaxAnalyzerClass() SyntaxAnalyzerClassLike {
	return syntaxAnalyzerClass()
}

// Constructor Methods

func (c *syntaxAnalyzerClass_) SyntaxAnalyzer(
	syntax not.SyntaxLike,
) SyntaxAnalyzerLike {
	var instance = &syntaxAnalyzer_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: not.Processor(),
	}
	instance.visitor_ = not.Visitor(instance)
	instance.visitor_.VisitSyntax(syntax)
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *syntaxAnalyzer_) GetClass() SyntaxAnalyzerClassLike {
	return syntaxAnalyzerClass()
}

func (v *syntaxAnalyzer_) GetPatterns() col.CatalogLike[string, string] {
	return v.patterns_
}

func (v *syntaxAnalyzer_) GetDefinitions() col.CatalogLike[string, not.DefinitionLike] {
	return v.definitions_
}

func (v *syntaxAnalyzer_) GetRuleOptions(
	ruleName string,
) col.ListLike[not.RuleOptionLike] {
	return v.ruleOptions_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetExpressionOptions(
	ruleName string,
) col.ListLike[not.ExpressionOptionLike] {
	return v.tokenOptions_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetLiteralOptions(
	ruleName string,
) col.ListLike[not.LiteralOptionLike] {
	return v.literalOptions_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetLegalNotice() string {
	return v.legalNotice_
}

func (v *syntaxAnalyzer_) GetRuleNames() col.SetLike[string] {
	return v.ruleNames_
}

func (v *syntaxAnalyzer_) GetSyntaxMap() string {
	return v.syntaxMap_
}

func (v *syntaxAnalyzer_) GetSyntaxName() string {
	return v.syntaxName_
}

func (v *syntaxAnalyzer_) GetTerms(
	ruleName string,
) col.ListLike[not.TermLike] {
	return v.terms_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetTokenNames() col.SetLike[string] {
	return v.tokenNames_
}

func (v *syntaxAnalyzer_) GetVariables(
	ruleName string,
) col.ListLike[string] {
	return v.variables_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetVariableType(
	component not.ComponentLike,
) string {
	var variableType string
	var componentString = component.GetAny().(string)
	switch {
	case not.MatchesType(componentString, not.LiteralToken):
		variableType = "string"
	case not.MatchesType(componentString, not.LowercaseToken):
		variableType = "string"
	case not.MatchesType(componentString, not.UppercaseToken):
		variableType = componentString + "Like"
	}
	return variableType
}

func (v *syntaxAnalyzer_) MakeOptional(
	mixedCase string,
) string {
	var optional string
	if len(mixedCase) > 0 {
		optional = "optional" + uti.MakeUpperCase(mixedCase)
	}
	return optional
}

// not.Methodical Methods

func (v *syntaxAnalyzer_) ProcessExcluded(
	excluded string,
) {
	v.pattern_ += "^"
}

func (v *syntaxAnalyzer_) ProcessGlyph(
	glyph string,
) {
	var character = glyph[1 : len(glyph)-1] // Remove the single quotes.
	character = v.escapeText(character)
	v.pattern_ += character
}

func (v *syntaxAnalyzer_) ProcessIntrinsic(
	intrinsic string,
) {
	intrinsic = sts.ToLower(intrinsic)
	if intrinsic == "any" {
		v.isGreedy_ = false // Turn off "greedy" for expressions containing ANY.
	}
	v.pattern_ += `" + ` + intrinsic + `_ + "`
}

func (v *syntaxAnalyzer_) ProcessLiteral(
	literal string,
) {
	v.hasLiteral_ = true
	var delimiter, err = stc.Unquote(literal) // Remove the double quotes.
	if err != nil {
		panic(err)
	}
	delimiter = v.escapeText(delimiter)
	switch {
	case v.inDefinition_:
		v.delimiters_.AddValue(delimiter)
	case v.inPattern_ > 0:
		v.pattern_ += delimiter
	}
}

func (v *syntaxAnalyzer_) ProcessLowercase(
	lowercase string,
) {
	switch {
	case v.inDefinition_:
		v.tokenNames_.AddValue(lowercase)
	case v.inPattern_ > 0:
		v.pattern_ += `(?:" + ` + lowercase + `_ + ")`
	}
}

func (v *syntaxAnalyzer_) ProcessNumber(
	number string,
) {
	if v.inPattern_ > 0 {
		v.pattern_ += number
	}
}

func (v *syntaxAnalyzer_) ProcessOptional(
	optional string,
) {
	if v.inPattern_ > 0 {
		v.pattern_ += optional
	}
}

func (v *syntaxAnalyzer_) ProcessRepeated(
	repeated string,
) {
	if v.inPattern_ > 0 {
		v.pattern_ += repeated
	}
}

func (v *syntaxAnalyzer_) PreprocessAlternative(
	alternative not.AlternativeLike,
	index uint,
	size uint,
) {
	v.pattern_ += "|"
}

func (v *syntaxAnalyzer_) PostprocessConstrained(
	constrained not.ConstrainedLike,
) {
	if v.inPattern_ > 0 && !v.isGreedy_ {
		v.pattern_ += "?"
		v.isGreedy_ = true // Reset scanning back to "greedy".
	}
}

func (v *syntaxAnalyzer_) PreprocessDefinition(
	definition not.DefinitionLike,
) {
	v.inDefinition_ = true
}

func (v *syntaxAnalyzer_) PostprocessDefinition(
	definition not.DefinitionLike,
) {
	v.inDefinition_ = false
}

func (v *syntaxAnalyzer_) PreprocessExpression(
	expression not.ExpressionLike,
	index uint,
	size uint,
) {
	v.pattern_ = `"(?:`
}

func (v *syntaxAnalyzer_) PostprocessExpression(
	expression not.ExpressionLike,
	index uint,
	size uint,
) {
	v.pattern_ += `)"`
	var expressionName = expression.GetLowercase()
	v.patterns_.SetValue(expressionName, v.pattern_)
}

func (v *syntaxAnalyzer_) PreprocessExpressionOption(
	tokenOption not.ExpressionOptionLike,
	index uint,
	size uint,
) {
	var tokenOptions = v.tokenOptions_.GetValue(v.ruleName_)
	tokenOptions.AppendValue(tokenOption)
	v.syntaxMap_ += "\n    " + tokenOption.GetLowercase()
	var note = tokenOption.GetOptionalNote()
	if uti.IsDefined(note) {
		v.syntaxMap_ += "  " + note
	}
}

func (v *syntaxAnalyzer_) PreprocessExtent(
	extent not.ExtentLike,
) {
	v.pattern_ += "-"
}

func (v *syntaxAnalyzer_) PreprocessFilter(
	filter not.FilterLike,
) {
	v.pattern_ += "["
}

func (v *syntaxAnalyzer_) PostprocessFilter(
	filter not.FilterLike,
) {
	v.pattern_ += "]"
}

func (v *syntaxAnalyzer_) PreprocessGroup(
	group not.GroupLike,
) {
	v.pattern_ += "("
}

func (v *syntaxAnalyzer_) PostprocessGroup(
	group not.GroupLike,
) {
	v.pattern_ += ")"
}

func (v *syntaxAnalyzer_) PostprocessInlineRule(
	inline not.InlineRuleLike,
) {
	var note = inline.GetOptionalNote()
	if uti.IsDefined(note) {
		v.syntaxMap_ += "  " + note
	}
}

func (v *syntaxAnalyzer_) PreprocessLimit(
	limit not.LimitLike,
) {
	if v.inPattern_ > 0 {
		v.pattern_ += ","
	}
}

func (v *syntaxAnalyzer_) PreprocessLiteralOption(
	literalOption not.LiteralOptionLike,
	index uint,
	size uint,
) {
	var literalOptions = v.literalOptions_.GetValue(v.ruleName_)
	literalOptions.AppendValue(literalOption)
	v.syntaxMap_ += "\n    " + literalOption.GetLiteral()
	var note = literalOption.GetOptionalNote()
	if uti.IsDefined(note) {
		v.syntaxMap_ += "  " + note
	}
}

func (v *syntaxAnalyzer_) PreprocessPattern(
	definition not.PatternLike,
) {
	v.inPattern_++
}

func (v *syntaxAnalyzer_) PostprocessPattern(
	definition not.PatternLike,
) {
	v.inPattern_--
}

func (v *syntaxAnalyzer_) PreprocessQuantified(
	quantified not.QuantifiedLike,
) {
	if v.inPattern_ > 0 {
		v.pattern_ += "{"
	}
}

func (v *syntaxAnalyzer_) PostprocessQuantified(
	quantified not.QuantifiedLike,
) {
	if v.inPattern_ > 0 {
		v.pattern_ += "}"
		if !v.isGreedy_ {
			v.pattern_ += "?"
			v.isGreedy_ = true // Reset scanning back to "greedy".
		}
	}
}

func (v *syntaxAnalyzer_) PreprocessRule(
	rule not.RuleLike,
	index uint,
	size uint,
) {
	v.hasLiteral_ = false
	var ruleName = rule.GetUppercase()
	v.ruleName_ = ruleName
	v.ruleNames_.AddValue(ruleName)
	var definition = rule.GetDefinition()
	v.definitions_.SetValue(ruleName, definition)
	switch definition.GetAny().(type) {
	case not.MultiLiteralLike:
		var literalOptions = col.List[not.LiteralOptionLike]()
		v.literalOptions_.SetValue(ruleName, literalOptions)
	case not.MultiExpressionLike:
		var tokenOptions = col.List[not.ExpressionOptionLike]()
		v.tokenOptions_.SetValue(ruleName, tokenOptions)
	case not.MultiRuleLike:
		var ruleOptions = col.List[not.RuleOptionLike]()
		v.ruleOptions_.SetValue(ruleName, ruleOptions)
	case not.InlineRuleLike:
		var terms = col.List[not.TermLike]()
		v.terms_.SetValue(ruleName, terms)
	}
	v.syntaxMap_ += "\n\t\t\t\"$" + ruleName + "\": `"
}

func (v *syntaxAnalyzer_) PostprocessRule(
	rule not.RuleLike,
	index uint,
	size uint,
) {
	var ruleName = rule.GetUppercase()
	v.extractVariables(ruleName)
	v.syntaxMap_ += "`,"
}

func (v *syntaxAnalyzer_) PreprocessRuleOption(
	ruleOption not.RuleOptionLike,
	index uint,
	size uint,
) {
	var ruleOptions = v.ruleOptions_.GetValue(v.ruleName_)
	ruleOptions.AppendValue(ruleOption)
	v.syntaxMap_ += "\n    " + ruleOption.GetUppercase()
	var note = ruleOption.GetOptionalNote()
	if uti.IsDefined(note) {
		v.syntaxMap_ += "  " + note
	}
}

func (v *syntaxAnalyzer_) PreprocessSyntax(
	syntax not.SyntaxLike,
) {
	v.isGreedy_ = true // The default is "greedy" scanning.
	v.syntaxName_ = v.extractSyntaxName(syntax)
	v.legalNotice_ = v.extractLegalNotice(syntax)
	v.ruleNames_ = col.Set[string]()
	v.tokenNames_ = col.SetFromArray[string](
		[]string{"delimiter", "newline", "space"},
	)
	v.pluralNames_ = col.Set[string]()
	v.delimiters_ = col.Set[string]()
	v.patterns_ = col.Catalog[string, string]()
	v.patterns_.SetValue(
		"delimiter",
		`"(?:)"`, // The delimiters will be filled in later but need to be first.
	)
	v.patterns_.SetValue(
		"newline",
		`"(?:" + eol_ + ")"`,
	)
	v.patterns_.SetValue(
		"space",
		`"(?:[ \\t]+)"`,
	)
	v.definitions_ = col.Catalog[string, not.DefinitionLike]()
	v.terms_ = col.Catalog[string, col.ListLike[not.TermLike]]()
	v.variables_ = col.Catalog[string, col.ListLike[string]]()
	v.ruleOptions_ = col.Catalog[string, col.ListLike[not.RuleOptionLike]]()
	v.tokenOptions_ = col.Catalog[string, col.ListLike[not.ExpressionOptionLike]]()
	v.literalOptions_ = col.Catalog[string, col.ListLike[not.LiteralOptionLike]]()
}

func (v *syntaxAnalyzer_) PostprocessSyntax(
	syntax not.SyntaxLike,
) {
	var delimiters = `"(?:`
	if !v.delimiters_.IsEmpty() {
		var iterator = v.delimiters_.GetIterator()
		iterator.ToEnd() // These must be assembled in reverse alphabetical order.
		delimiters += iterator.GetPrevious()
		for iterator.HasPrevious() {
			delimiters += "|" + iterator.GetPrevious()
		}
	}
	delimiters += `)"`
	v.patterns_.SetValue("delimiter", delimiters)
}

func (v *syntaxAnalyzer_) PreprocessTerm(
	term not.TermLike,
	index uint,
	size uint,
) {
	// Process the term.
	var terms = v.terms_.GetValue(v.ruleName_)
	terms.AppendValue(term)

	// Process the component.
	var component = term.GetComponent()
	var componentString = component.GetAny().(string)
	if index > 1 {
		v.syntaxMap_ += " "
	}
	v.syntaxMap_ += componentString

	// Process the optional cardinality.
	var cardinality = term.GetOptionalCardinality()
	if uti.IsDefined(cardinality) {
		v.checkPlurality(componentString, cardinality)
		switch actual := cardinality.GetAny().(type) {
		case not.ConstrainedLike:
			v.syntaxMap_ += actual.GetAny().(string)
		case not.QuantifiedLike:
			var first = actual.GetNumber()
			v.syntaxMap_ += "{" + first
			var limit = actual.GetOptionalLimit()
			if uti.IsDefined(limit) {
				v.syntaxMap_ += ".."
				var last = limit.GetOptionalNumber()
				if uti.IsDefined(last) {
					v.syntaxMap_ += last + "}"
				}
			}
		}
	}
}

// PROTECTED INTERFACE

// Private Methods

func (v *syntaxAnalyzer_) checkPlurality(
	identifierName string,
	cardinality not.CardinalityLike,
) {
	switch actual := cardinality.GetAny().(type) {
	case not.ConstrainedLike:
		switch actual.GetAny().(string) {
		case "*", "+":
			v.pluralNames_.AddValue(identifierName)
		}
	case not.QuantifiedLike:
		v.pluralNames_.AddValue(identifierName)
	}
}

func (v *syntaxAnalyzer_) escapeText(
	text string,
) string {
	var escaped string
	for _, character := range text {
		switch character {
		case '"':
			escaped += `\`
		case '.', '|', '^', '$', '+', '*', '?', '(', ')', '[', ']', '{', '}':
			escaped += `\\`
		case '\\':
			escaped += `\\\`
		}
		escaped += string(character)
	}
	return escaped
}

func (v *syntaxAnalyzer_) extractLegalNotice(
	syntax not.SyntaxLike,
) string {
	var comment = syntax.GetNotice().GetComment()

	// Strip off the syntax style comment delimiters.
	comment = comment[2 : len(comment)-3]

	// Add in the go style comment delimiters.
	var notice = "/*" + comment + "*/\n"

	return notice
}

func (v *syntaxAnalyzer_) extractSyntaxName(
	syntax not.SyntaxLike,
) string {
	var rules = syntax.GetRules().GetIterator()
	// The first rule name is the name of the syntax.
	var rule = rules.GetNext()
	var syntaxName = rule.GetUppercase()
	return syntaxName
}

func (v *syntaxAnalyzer_) extractVariableName(
	term not.TermLike,
) string {
	var component = term.GetComponent().GetAny().(string)
	if component[0] == '"' {
		// A literal string found in a rule definition represents a delimiter.
		component = "delimiter"
	}
	var variableName = uti.MakeLowerCase(component)
	var cardinality = term.GetOptionalCardinality()
	if uti.IsDefined(cardinality) {
		switch actual := cardinality.GetAny().(type) {
		case not.ConstrainedLike:
			var constrained = actual.GetAny().(string)
			switch constrained {
			case "?":
				variableName = v.MakeOptional(variableName)
			case "*", "+":
				variableName = uti.MakePlural(variableName)
			}
		case not.QuantifiedLike:
			variableName = uti.MakePlural(variableName)
		}
	}
	return variableName
}

func (v *syntaxAnalyzer_) extractVariables(
	ruleName string,
) {
	// Only inline rules have variables.
	var terms col.Sequential[not.TermLike]
	var definition = v.definitions_.GetValue(ruleName)
	switch actual := definition.GetAny().(type) {
	case not.InlineRuleLike:
		terms = actual.GetTerms()
	default:
		return
	}

	// Extract the variable names from the inline components.
	var variables = col.List[string]()
	var iterator = terms.GetIterator()
	for iterator.HasNext() {
		var term = iterator.GetNext()
		var variable = v.extractVariableName(term)
		variables.AppendValue(variable)
	}

	// Make any duplicate variable names unique.
	var left = variables.GetIterator()
	var right = variables.GetIterator()
	for left.HasNext() {
		var leftName = left.GetNext()
		var slot = left.GetSlot()
		right.SetSlot(slot)
		for right.HasNext() {
			var count = 1
			var rightName = right.GetNext()
			if leftName == rightName {
				if count == 1 {
					var uniqueName = leftName + stc.Itoa(count)
					var index = col.Index(left.GetSlot())
					variables.SetValue(index, uniqueName)
				}
				count++
				rightName = rightName + stc.Itoa(count)
				var index = col.Index(right.GetSlot())
				variables.SetValue(index, rightName)
			}
		}
	}
	v.variables_.SetValue(ruleName, variables)
}

// Instance Structure

type syntaxAnalyzer_ struct {
	// Declare the instance attributes.
	visitor_        not.VisitorLike
	isGreedy_       bool
	inDefinition_   bool
	inPattern_      uint8
	hasLiteral_     bool
	syntaxMap_      string
	syntaxName_     string
	legalNotice_    string
	ruleName_       string
	pattern_        string
	ruleNames_      col.SetLike[string]
	tokenNames_     col.SetLike[string]
	pluralNames_    col.SetLike[string]
	delimiters_     col.SetLike[string]
	patterns_       col.CatalogLike[string, string]
	definitions_    col.CatalogLike[string, not.DefinitionLike]
	terms_          col.CatalogLike[string, col.ListLike[not.TermLike]]
	variables_      col.CatalogLike[string, col.ListLike[string]]
	ruleOptions_    col.CatalogLike[string, col.ListLike[not.RuleOptionLike]]
	tokenOptions_   col.CatalogLike[string, col.ListLike[not.ExpressionOptionLike]]
	literalOptions_ col.CatalogLike[string, col.ListLike[not.LiteralOptionLike]]

	// Declare the inherited aspects.
	not.Methodical
}

// Class Structure

type syntaxAnalyzerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func syntaxAnalyzerClass() *syntaxAnalyzerClass_ {
	return syntaxAnalyzerClassReference_
}

var syntaxAnalyzerClassReference_ = &syntaxAnalyzerClass_{
	// Initialize the class constants.
}
