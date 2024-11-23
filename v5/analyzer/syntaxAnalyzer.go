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

package analyzer

import (
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	not "github.com/craterdog/go-syntax-notation/v5"
	gra "github.com/craterdog/go-syntax-notation/v5/grammar"
	stc "strconv"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func SyntaxAnalyzer() SyntaxAnalyzerClassLike {
	return syntaxAnalyzerReference()
}

// Constructor Methods

func (c *syntaxAnalyzerClass_) Make(
	syntax not.SyntaxLike,
) SyntaxAnalyzerLike {
	var instance = &syntaxAnalyzer_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: not.Processor(),
	}
	instance.visitor_ = gra.Visitor().Make(instance)
	instance.visitor_.VisitSyntax(syntax)
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *syntaxAnalyzer_) GetClass() SyntaxAnalyzerClassLike {
	return syntaxAnalyzerReference()
}

func (v *syntaxAnalyzer_) GetExpressions() abs.Sequential[abs.AssociationLike[string, string]] {
	return v.regexps_
}

func (v *syntaxAnalyzer_) GetIdentifiers(
	ruleName string,
) abs.Sequential[not.IdentifierLike] {
	return v.identifiers_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetLegalNotice() string {
	return v.legalNotice_
}

func (v *syntaxAnalyzer_) GetReferences(
	ruleName string,
) abs.Sequential[not.ReferenceLike] {
	return v.references_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetRuleNames() abs.Sequential[string] {
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
) abs.Sequential[not.TermLike] {
	return v.terms_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetTokenNames() abs.Sequential[string] {
	return v.tokenNames_
}

func (v *syntaxAnalyzer_) GetVariables(
	ruleName string,
) abs.Sequential[string] {
	return v.variables_.GetValue(ruleName)
}

func (v *syntaxAnalyzer_) GetVariableType(
	reference not.ReferenceLike,
) string {
	var variableType string
	var identifier = reference.GetIdentifier().GetAny().(string)
	var scannerClass = gra.Scanner()
	switch {
	case scannerClass.MatchesType(identifier, gra.LowercaseToken):
		variableType = "string"
	case scannerClass.MatchesType(identifier, gra.UppercaseToken):
		variableType = uti.MakeUpperCase(identifier) + "Like"
	}
	return variableType
}

func (v *syntaxAnalyzer_) HasPlurals() bool {
	return v.pluralNames_.GetSize() > 0
}

func (v *syntaxAnalyzer_) IsDelimited(
	ruleName string,
) bool {
	return v.delimited_.ContainsValue(ruleName)
}

func (v *syntaxAnalyzer_) IsPlural(
	name string,
) bool {
	return v.pluralNames_.ContainsValue(name)
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
	v.regexp_ += "^"
}

func (v *syntaxAnalyzer_) ProcessGlyph(
	glyph string,
) {
	var character = glyph[1:2] //Remove the single quotes.
	character = v.escapeText(character)
	v.regexp_ += character
}

func (v *syntaxAnalyzer_) ProcessIntrinsic(
	intrinsic string,
) {
	intrinsic = sts.ToLower(intrinsic)
	if intrinsic == "any" {
		v.isGreedy_ = false // Turn off "greedy" for expressions containing ANY.
	}
	v.regexp_ += `" + ` + intrinsic + `_ + "`
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
	if v.inDefinition_ {
		v.delimiters_.AddValue(delimiter)
	}
	v.regexp_ += delimiter
}

func (v *syntaxAnalyzer_) ProcessLowercase(
	lowercase string,
) {
	if v.inDefinition_ {
		v.tokenNames_.AddValue(lowercase)
	}
	if v.inPattern_ {
		v.regexp_ += `(?:" + ` + lowercase + `_ + ")`
	}
}

func (v *syntaxAnalyzer_) ProcessNumber(
	number string,
) {
	v.regexp_ += number
}

func (v *syntaxAnalyzer_) ProcessOptional(
	optional string,
) {
	v.regexp_ += optional
}

func (v *syntaxAnalyzer_) ProcessRepeated(
	repeated string,
) {
	v.regexp_ += repeated
}

func (v *syntaxAnalyzer_) PreprocessAlternative(
	alternative not.AlternativeLike,
	index uint,
	size uint,
) {
	v.regexp_ += "|"
}

func (v *syntaxAnalyzer_) PostprocessConstrained(
	constrained not.ConstrainedLike,
) {
	if !v.isGreedy_ {
		v.regexp_ += "?"
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
	v.regexp_ = `"(?:`
}

func (v *syntaxAnalyzer_) PostprocessExpression(
	expression not.ExpressionLike,
	index uint,
	size uint,
) {
	v.regexp_ += `)"`
	var name = expression.GetLowercase()
	v.regexps_.SetValue(name, v.regexp_)
}

func (v *syntaxAnalyzer_) PreprocessExtent(
	extent not.ExtentLike,
) {
	v.regexp_ += "-"
}

func (v *syntaxAnalyzer_) PreprocessFilter(
	filter not.FilterLike,
) {
	v.regexp_ += "["
}

func (v *syntaxAnalyzer_) PostprocessFilter(
	filter not.FilterLike,
) {
	v.regexp_ += "]"
}

func (v *syntaxAnalyzer_) PreprocessGroup(
	group not.GroupLike,
) {
	v.regexp_ += "("
}

func (v *syntaxAnalyzer_) PostprocessGroup(
	group not.GroupLike,
) {
	v.regexp_ += ")"
}

func (v *syntaxAnalyzer_) PreprocessIdentifier(
	identifier not.IdentifierLike,
) {
	var scannerClass = gra.Scanner()
	var name = identifier.GetAny().(string)
	if scannerClass.MatchesType(name, gra.LowercaseToken) {
		v.tokenNames_.AddValue(name)
	}
}

func (v *syntaxAnalyzer_) PostprocessInline(
	inline not.InlineLike,
) {
	var note = inline.GetOptionalNote()
	if uti.IsDefined(note) {
		v.syntaxMap_ += "  " + note
	}
}

func (v *syntaxAnalyzer_) PreprocessLimit(
	limit not.LimitLike,
) {
	v.regexp_ += ","
}

func (v *syntaxAnalyzer_) PreprocessLine(
	line not.LineLike,
	index uint,
	size uint,
) {
	var identifier = line.GetIdentifier()
	var identifiers = v.identifiers_.GetValue(v.ruleName_)
	identifiers.AppendValue(identifier)
	v.syntaxMap_ += "\n  - " + identifier.GetAny().(string)
	var note = line.GetOptionalNote()
	if uti.IsDefined(note) {
		v.syntaxMap_ += "  " + note
	}
}

func (v *syntaxAnalyzer_) PreprocessPattern(
	definition not.PatternLike,
) {
	v.inPattern_ = true
}

func (v *syntaxAnalyzer_) PostprocessPattern(
	definition not.PatternLike,
) {
	v.inPattern_ = false
}

func (v *syntaxAnalyzer_) PreprocessQuantified(
	quantified not.QuantifiedLike,
) {
	v.regexp_ += "{"
}

func (v *syntaxAnalyzer_) PostprocessQuantified(
	quantified not.QuantifiedLike,
) {
	v.regexp_ += "}"
	if !v.isGreedy_ {
		v.regexp_ += "?"
		v.isGreedy_ = true // Reset scanning back to "greedy".
	}
}

func (v *syntaxAnalyzer_) PreprocessReference(
	reference not.ReferenceLike,
) {
	var references = v.references_.GetValue(v.ruleName_)
	references.AppendValue(reference)

	// Process the identifier.
	var identifier = reference.GetIdentifier()
	v.syntaxMap_ += identifier.GetAny().(string)

	// Process the cardinality.
	var cardinality = reference.GetOptionalCardinality()
	if uti.IsDefined(cardinality) {
		var name = identifier.GetAny().(string)
		v.checkPlurality(name, cardinality)
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
	switch definition.GetAny().(type) {
	case not.InlineLike:
		var terms = col.List[not.TermLike]()
		v.terms_.SetValue(ruleName, terms)
		var references = col.List[not.ReferenceLike]()
		v.references_.SetValue(ruleName, references)
	case not.MultilineLike:
		var identifiers = col.List[not.IdentifierLike]()
		v.identifiers_.SetValue(ruleName, identifiers)
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
	if v.hasLiteral_ {
		v.delimited_.AddValue(ruleName)
	}
	v.syntaxMap_ += "`,"
}

func (v *syntaxAnalyzer_) PreprocessSyntax(
	syntax not.SyntaxLike,
) {
	v.isGreedy_ = true // The default is "greedy" scanning.
	v.syntaxName_ = v.extractSyntaxName(syntax)
	v.legalNotice_ = v.extractLegalNotice(syntax)
	v.ruleNames_ = col.Set[string]()
	v.tokenNames_ = col.Set[string]([]string{"delimiter", "newline", "space"})
	v.pluralNames_ = col.Set[string]()
	v.delimited_ = col.Set[string]()
	v.delimiters_ = col.Set[string]()
	var implicit = map[string]string{
		"newline": `"(?:" + eol_ + ")"`,
		"space":   `"(?:[ \\t]+)"`,
	}
	v.regexps_ = col.Catalog[string, string](implicit)
	v.terms_ = col.Catalog[string, abs.ListLike[not.TermLike]]()
	v.variables_ = col.Catalog[string, abs.ListLike[string]]()
	v.references_ = col.Catalog[string, abs.ListLike[not.ReferenceLike]]()
	v.identifiers_ = col.Catalog[string, abs.ListLike[not.IdentifierLike]]()
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
	v.regexps_.SetValue("delimiter", delimiters)
	v.regexps_.SortValues()
}

func (v *syntaxAnalyzer_) PreprocessTerm(
	term not.TermLike,
	index uint,
	size uint,
) {
	if index > 1 {
		v.syntaxMap_ += " "
	}
	switch actual := term.GetAny().(type) {
	case string:
		v.syntaxMap_ += actual
	}
	var terms = v.terms_.GetValue(v.ruleName_)
	terms.AppendValue(term)
}

// PROTECTED INTERFACE

// Private Methods

func (v *syntaxAnalyzer_) checkPlurality(
	name string,
	cardinality not.CardinalityLike,
) {
	switch actual := cardinality.GetAny().(type) {
	case not.ConstrainedLike:
		switch actual.GetAny().(string) {
		case "*", "+":
			v.pluralNames_.AddValue(name)
		}
	case not.QuantifiedLike:
		v.pluralNames_.AddValue(name)
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
	var name = rule.GetUppercase()
	return name
}

func (v *syntaxAnalyzer_) extractVariableName(
	reference not.ReferenceLike,
) string {
	var mixedCase = reference.GetIdentifier().GetAny().(string)
	var variableName = uti.MakeLowerCase(mixedCase)
	var cardinality = reference.GetOptionalCardinality()
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
	// Check for multiline rule.
	var references = v.references_.GetValue(ruleName)
	if uti.IsUndefined(references) {
		return
	}

	// Extract the variable names from the inline references.
	var variables = col.List[string]()
	var iterator = references.GetIterator()
	for iterator.HasNext() {
		var reference = iterator.GetNext()
		var variable = v.extractVariableName(reference)
		variables.AppendValue(variable)
	}

	// Make any duplicate variable names unique.
	var left = variables.GetIterator()
	var right = variables.GetIterator()
	for left.HasNext() {
		var leftName = left.GetNext()
		var slot = left.GetSlot()
		right.ToSlot(slot)
		for right.HasNext() {
			var count = 1
			var rightName = right.GetNext()
			if leftName == rightName {
				if count == 1 {
					var uniqueName = leftName + stc.Itoa(count)
					var index = left.GetSlot()
					variables.SetValue(index, uniqueName)
				}
				count++
				rightName = rightName + stc.Itoa(count)
				var index = right.GetSlot()
				variables.SetValue(index, rightName)
			}
		}
	}
	v.variables_.SetValue(ruleName, variables)
}

// Instance Structure

type syntaxAnalyzer_ struct {
	// Declare the instance attributes.
	visitor_      not.VisitorLike
	isGreedy_     bool
	inDefinition_ bool
	inPattern_    bool
	hasLiteral_   bool
	syntaxMap_    string
	syntaxName_   string
	legalNotice_  string
	ruleName_     string
	regexp_       string
	ruleNames_    abs.SetLike[string]
	tokenNames_   abs.SetLike[string]
	pluralNames_  abs.SetLike[string]
	delimited_    abs.SetLike[string]
	delimiters_   abs.SetLike[string]
	regexps_      abs.CatalogLike[string, string]
	terms_        abs.CatalogLike[string, abs.ListLike[not.TermLike]]
	variables_    abs.CatalogLike[string, abs.ListLike[string]]
	references_   abs.CatalogLike[string, abs.ListLike[not.ReferenceLike]]
	identifiers_  abs.CatalogLike[string, abs.ListLike[not.IdentifierLike]]

	// Declare the inherited aspects.
	not.Methodical
}

// Class Structure

type syntaxAnalyzerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func syntaxAnalyzerReference() *syntaxAnalyzerClass_ {
	return syntaxAnalyzerReference_
}

var syntaxAnalyzerReference_ = &syntaxAnalyzerClass_{
	// Initialize the class constants.
}
