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
	fra "github.com/craterdog/go-collection-framework/v8"
	uti "github.com/craterdog/go-missing-utilities/v8"
	not "github.com/craterdog/go-syntax-notation/v8"
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
	not.Visitor(instance).VisitSyntax(syntax)
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *syntaxAnalyzer_) GetClass() SyntaxAnalyzerClassLike {
	return syntaxAnalyzerClass()
}

func (v *syntaxAnalyzer_) GetLegalNotice() string {
	return v.legalNotice_
}

func (v *syntaxAnalyzer_) GetSyntaxName() string {
	return v.syntaxName_
}

func (v *syntaxAnalyzer_) GetRules() fra.SetLike[string] {
	return v.rules_
}

func (v *syntaxAnalyzer_) GetExpressions() fra.SetLike[string] {
	return v.expressions_
}

func (v *syntaxAnalyzer_) GetFragments() fra.SetLike[string] {
	return v.fragments_
}

func (v *syntaxAnalyzer_) GetLiteralValues(
	ruleName string,
) fra.Sequential[not.LiteralValueLike] {
	return v.literalValues_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetExpressionNames(
	ruleName string,
) fra.Sequential[not.ExpressionNameLike] {
	return v.expressionNames_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetRuleNames(
	ruleName string,
) fra.Sequential[not.RuleNameLike] {
	return v.ruleNames_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetRuleTerms(
	ruleName string,
) fra.Sequential[not.RuleTermLike] {
	return v.ruleTerms_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetVariables(
	ruleName string,
) fra.Sequential[string] {
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

func (v *syntaxAnalyzer_) GetPatterns() fra.CatalogLike[string, string] {
	return v.patterns_
}

func (v *syntaxAnalyzer_) GetDefinitions() fra.CatalogLike[string, not.DefinitionLike] {
	return v.definitions_
}

func (v *syntaxAnalyzer_) GetSyntaxMap() string {
	return v.syntaxMap_
}

// not.Methodical Methods

func (v *syntaxAnalyzer_) ProcessAllcaps(
	allcaps string,
) {
	switch {
	case v.inPattern_ > 0:
		// The fragment name is part of a fragment pattern.
		var lowercase = sts.ToLower(allcaps)
		v.pattern_ += `(?:" + ` + lowercase + `_ + ")`
	default:
		// The fragment name is part of a fragment definition.
		v.fragments_.AddValue(allcaps)
	}
}

func (v *syntaxAnalyzer_) ProcessDelimiter(
	delimiter string,
) {
	switch delimiter {
	case "~":
		v.pattern_ += "^"
	case "?", "*", "+":
		if v.inPattern_ > 0 {
			v.pattern_ += delimiter
		}
	}
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
		v.notGreedy_ = true // Turn off "greedy" for expressions containing ANY.
	}
	v.pattern_ += `" + ` + intrinsic + `_ + "`
}

func (v *syntaxAnalyzer_) ProcessLiteral(
	literal string,
) {
	// NOTE:
	// A scanned literal from the syntax notation corresponds to a delimiter in
	// the generated scanner and parser.
	var delimiter, err = stc.Unquote(literal) // Remove the double quotes.
	if err != nil {
		panic(err)
	}
	delimiter = v.escapeText(delimiter)
	switch {
	case v.inDefinition_:
		// The literal is part of a rule.
		v.delimiters_.AddValue(delimiter)
	case v.inPattern_ > 0:
		// The literal is part of an expression.
		v.pattern_ += delimiter
	}
}

func (v *syntaxAnalyzer_) ProcessLowercase(
	lowercase string,
) {
	switch {
	case v.inDefinition_:
		// The expression name is part of a rule definition.
		v.expressions_.AddValue(lowercase)
	case v.inPattern_ > 0:
		// The expression name is part of an expression pattern.
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

func (v *syntaxAnalyzer_) PreprocessAlternativeSequence(
	alternative not.AlternativeSequenceLike,
	index uint,
	count uint,
) {
	v.pattern_ += "|"
}

func (v *syntaxAnalyzer_) PostprocessConstrained(
	constrained not.ConstrainedLike,
	index uint,
	count uint,
) {
	if v.inPattern_ > 0 && v.notGreedy_ {
		v.pattern_ += "?"    // Specify non-greedy scanning.
		v.notGreedy_ = false // Reset scanning back to "greedy" (not notGreedy).
	}
}

func (v *syntaxAnalyzer_) PreprocessDefinition(
	definition not.DefinitionLike,
	index uint,
	count uint,
) {
	v.inDefinition_ = true
}

func (v *syntaxAnalyzer_) PostprocessDefinition(
	definition not.DefinitionLike,
	index uint,
	count uint,
) {
	v.inDefinition_ = false
}

func (v *syntaxAnalyzer_) PreprocessExpression(
	expression not.ExpressionLike,
	index uint,
	count uint,
) {
	v.pattern_ = `"(?:`
}

func (v *syntaxAnalyzer_) PostprocessExpression(
	expression not.ExpressionLike,
	index uint,
	count uint,
) {
	v.pattern_ += `)"`
	var expressionName = expression.GetLowercase()
	v.patterns_.SetValue(expressionName, v.pattern_)
}

func (v *syntaxAnalyzer_) PreprocessExtent(
	extent not.ExtentLike,
	index uint,
	count uint,
) {
	v.pattern_ += "-"
}

func (v *syntaxAnalyzer_) PreprocessFilter(
	filter not.FilterLike,
	index uint,
	count uint,
) {
	v.pattern_ += "["
}

func (v *syntaxAnalyzer_) PostprocessFilter(
	filter not.FilterLike,
	index uint,
	count uint,
) {
	v.pattern_ += "]"
}

func (v *syntaxAnalyzer_) PreprocessFragment(
	fragment not.FragmentLike,
	index uint,
	count uint,
) {
	v.pattern_ = `"(?:`
}

func (v *syntaxAnalyzer_) PostprocessFragment(
	fragment not.FragmentLike,
	index uint,
	count uint,
) {
	v.pattern_ += `)"`
	var allcaps = fragment.GetAllcaps()
	var lowercase = sts.ToLower(allcaps)
	v.patterns_.SetValue(lowercase, v.pattern_)
}

func (v *syntaxAnalyzer_) PreprocessGroup(
	group not.GroupLike,
	index uint,
	count uint,
) {
	v.pattern_ += "("
}

func (v *syntaxAnalyzer_) PostprocessGroup(
	group not.GroupLike,
	index uint,
	count uint,
) {
	v.pattern_ += ")"
}

func (v *syntaxAnalyzer_) PreprocessLimit(
	limit not.LimitLike,
	index uint,
	count uint,
) {
	if v.inPattern_ > 0 {
		v.pattern_ += ","
	}
}

func (v *syntaxAnalyzer_) PreprocessLiteralValue(
	literalValue not.LiteralValueLike,
	index uint,
	count uint,
) {
	var literalValues = v.literalValues_.GetValue(v.ruleName_)
	literalValues.AppendValue(literalValue)
	var literal = literalValue.GetLiteral()
	v.syntaxMap_ += "\n    " + literal
	var note = literalValue.GetOptionalNote()
	if uti.IsDefined(note) {
		v.syntaxMap_ += "  " + note
	}
}

func (v *syntaxAnalyzer_) PreprocessPattern(
	pattern not.PatternLike,
	index uint,
	count uint,
) {
	// Increase the pattern's nesting level.
	v.inPattern_++
}

func (v *syntaxAnalyzer_) PostprocessPattern(
	pattern not.PatternLike,
	index uint,
	count uint,
) {
	// Decrease the pattern's nesting level.
	v.inPattern_--
}

func (v *syntaxAnalyzer_) PreprocessQuantified(
	quantified not.QuantifiedLike,
	index uint,
	count uint,
) {
	if v.inPattern_ > 0 {
		v.pattern_ += "{"
	}
}

func (v *syntaxAnalyzer_) PostprocessQuantified(
	quantified not.QuantifiedLike,
	index uint,
	count uint,
) {
	if v.inPattern_ > 0 {
		v.pattern_ += "}"
		if v.notGreedy_ {
			v.pattern_ += "?"    // Specify non-greedy scanning.
			v.notGreedy_ = false // Reset scanning back to "greedy" (not notGreedy).
		}
	}
}

func (v *syntaxAnalyzer_) PreprocessRule(
	rule not.RuleLike,
	index uint,
	count uint,
) {
	var ruleName = rule.GetUppercase()
	v.ruleName_ = ruleName
	v.rules_.AddValue(ruleName)
	var definition = rule.GetDefinition()
	v.definitions_.SetValue(ruleName, definition)
	switch definition.GetAny().(type) {
	case not.LiteralAlternativesLike:
		var literalValues = fra.List[not.LiteralValueLike]()
		v.literalValues_.SetValue(ruleName, literalValues)
	case not.ExpressionAlternativesLike:
		var expressionNames = fra.List[not.ExpressionNameLike]()
		v.expressionNames_.SetValue(ruleName, expressionNames)
	case not.RuleAlternativesLike:
		var ruleNames = fra.List[not.RuleNameLike]()
		v.ruleNames_.SetValue(ruleName, ruleNames)
	case not.TermSequenceLike:
		var ruleTerms = fra.List[not.RuleTermLike]()
		v.ruleTerms_.SetValue(ruleName, ruleTerms)
	}
	v.syntaxMap_ += "\n\t\t\t\"$" + ruleName + "\": `"
}

func (v *syntaxAnalyzer_) PostprocessRule(
	rule not.RuleLike,
	index uint,
	count uint,
) {
	var ruleName = rule.GetUppercase()
	v.extractVariables(ruleName)
	v.syntaxMap_ += "`,"
}

func (v *syntaxAnalyzer_) PreprocessRuleName(
	ruleName not.RuleNameLike,
	index uint,
	count uint,
) {
	var ruleNames = v.ruleNames_.GetValue(v.ruleName_)
	ruleNames.AppendValue(ruleName)
	var uppercase = ruleName.GetUppercase()
	v.syntaxMap_ += "\n    " + uppercase
	var note = ruleName.GetOptionalNote()
	if uti.IsDefined(note) {
		v.syntaxMap_ += "  " + note
	}
}

func (v *syntaxAnalyzer_) PreprocessRuleTerm(
	ruleTerm not.RuleTermLike,
	index uint,
	count uint,
) {
	// Process the term.
	var ruleTerms = v.ruleTerms_.GetValue(v.ruleName_)
	ruleTerms.AppendValue(ruleTerm)

	// Process the component.
	var component = ruleTerm.GetComponent()
	var componentString = component.GetAny().(string)
	if index > 1 {
		v.syntaxMap_ += " "
	}
	v.syntaxMap_ += componentString

	// Process the optional cardinality.
	var cardinality = ruleTerm.GetOptionalCardinality()
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
					v.syntaxMap_ += last
				}
			}
			v.syntaxMap_ += "}"
		}
	}
}

func (v *syntaxAnalyzer_) PreprocessSyntax(
	syntax not.SyntaxLike,
	index uint,
	count uint,
) {
	v.notGreedy_ = false // The default is "greedy" (not notGreedy).
	v.syntaxName_ = v.extractSyntaxName(syntax)
	v.legalNotice_ = v.extractLegalNotice(syntax)
	v.rules_ = fra.Set[string]()
	v.expressions_ = fra.SetFromArray[string](
		[]string{"delimiter", "newline", "space"},
	)
	v.fragments_ = fra.SetFromArray[string](
		[]string{"ANY", "CONTROL", "DIGIT", "EOL", "LOWER", "UPPER"},
	)
	v.pluralNames_ = fra.Set[string]()
	v.delimiters_ = fra.Set[string]()
	v.patterns_ = fra.Catalog[string, string]()
	v.definitions_ = fra.Catalog[string, not.DefinitionLike]()
	v.ruleTerms_ = fra.Catalog[string, fra.ListLike[not.RuleTermLike]]()
	v.variables_ = fra.Catalog[string, fra.ListLike[string]]()
	v.ruleNames_ = fra.Catalog[string, fra.ListLike[not.RuleNameLike]]()
	v.expressionNames_ = fra.Catalog[string, fra.ListLike[not.ExpressionNameLike]]()
	v.literalValues_ = fra.Catalog[string, fra.ListLike[not.LiteralValueLike]]()
}

func (v *syntaxAnalyzer_) PostprocessSyntax(
	syntax not.SyntaxLike,
	index uint,
	count uint,
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
	v.patterns_.SetValue(
		"space",
		`"(?:[ \\t]+)"`,
	)
	v.patterns_.SetValue(
		"newline",
		`"(?:" + eol_ + ")"`,
	)
	v.patterns_.SetValue(
		"delimiter",
		delimiters,
	)
}

func (v *syntaxAnalyzer_) PostprocessTermSequence(
	termSequence not.TermSequenceLike,
	index uint,
	count uint,
) {
	var note = termSequence.GetOptionalNote()
	if uti.IsDefined(note) {
		v.syntaxMap_ += "  " + note
	}
}

func (v *syntaxAnalyzer_) PreprocessExpressionName(
	expressionName not.ExpressionNameLike,
	index uint,
	count uint,
) {
	var expressionNames = v.expressionNames_.GetValue(v.ruleName_)
	expressionNames.AppendValue(expressionName)
	v.syntaxMap_ += "\n    " + expressionName.GetLowercase()
	var note = expressionName.GetOptionalNote()
	if uti.IsDefined(note) {
		v.syntaxMap_ += "  " + note
	}
}

// PROTECTED INTERFACE

// Private Methods

func (v *syntaxAnalyzer_) checkPlurality(
	componentString string,
	cardinality not.CardinalityLike,
) {
	switch actual := cardinality.GetAny().(type) {
	case not.ConstrainedLike:
		switch actual.GetAny().(string) {
		case "*", "+":
			v.pluralNames_.AddValue(componentString)
		}
	case not.QuantifiedLike:
		v.pluralNames_.AddValue(componentString)
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
	var comment = syntax.GetLegalNotice().GetComment()

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
	ruleTerm not.RuleTermLike,
) string {
	var component = ruleTerm.GetComponent().GetAny().(string)
	if component[0] == '"' {
		// A literal string found in a rule definition represents a delimiter.
		component = "delimiter"
	}
	var variableName = uti.MakeLowerCase(component)
	var cardinality = ruleTerm.GetOptionalCardinality()
	if uti.IsDefined(cardinality) {
		switch actual := cardinality.GetAny().(type) {
		case not.ConstrainedLike:
			var constrained = actual.GetAny().(string)
			switch constrained {
			case "?":
				variableName = v.makeOptional(variableName)
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
	var ruleTerms fra.Sequential[not.RuleTermLike]
	var definition = v.definitions_.GetValue(ruleName)
	switch actual := definition.GetAny().(type) {
	case not.TermSequenceLike:
		ruleTerms = actual.GetRuleTerms()
	default:
		return
	}

	// Extract the variable names from the inline components.
	var variables = fra.List[string]()
	var iterator = ruleTerms.GetIterator()
	for iterator.HasNext() {
		var ruleTerm = iterator.GetNext()
		var variable = v.extractVariableName(ruleTerm)
		variables.AppendValue(variable)
	}

	// Make any duplicate variable names unique.
	var leftIndex int
	var rightIndex int
	var size = int(variables.GetSize())
	for leftIndex = 1; leftIndex <= size; leftIndex++ {
		var count = 1
		var leftName = variables.GetValue(leftIndex)
		for rightIndex = leftIndex + 1; rightIndex <= size; rightIndex++ {
			var rightName = variables.GetValue(rightIndex)
			if leftName == rightName {
				if count == 1 {
					// Add a count suffix of "1" to the first variable.
					var uniqueName = leftName + "1"
					variables.SetValue(leftIndex, uniqueName)
				}
				count++
				var uniqueName = rightName + stc.Itoa(count)
				variables.SetValue(rightIndex, uniqueName)
			}
		}
	}
	v.variables_.SetValue(ruleName, variables)
}

func (v *syntaxAnalyzer_) makeOptional(
	variableName string,
) string {
	var optional string
	if len(variableName) > 0 {
		optional = "optional" + uti.MakeUpperCase(variableName)
	}
	return optional
}

// Instance Structure

type syntaxAnalyzer_ struct {
	// Declare the instance attributes.
	notGreedy_       bool
	inDefinition_    bool
	inPattern_       uint8
	syntaxMap_       string
	syntaxName_      string
	legalNotice_     string
	ruleName_        string
	pattern_         string
	rules_           fra.SetLike[string]
	expressions_     fra.SetLike[string]
	fragments_       fra.SetLike[string]
	pluralNames_     fra.SetLike[string]
	delimiters_      fra.SetLike[string]
	patterns_        fra.CatalogLike[string, string]
	definitions_     fra.CatalogLike[string, not.DefinitionLike]
	ruleTerms_       fra.CatalogLike[string, fra.ListLike[not.RuleTermLike]]
	variables_       fra.CatalogLike[string, fra.ListLike[string]]
	ruleNames_       fra.CatalogLike[string, fra.ListLike[not.RuleNameLike]]
	expressionNames_ fra.CatalogLike[string, fra.ListLike[not.ExpressionNameLike]]
	literalValues_   fra.CatalogLike[string, fra.ListLike[not.LiteralValueLike]]

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
