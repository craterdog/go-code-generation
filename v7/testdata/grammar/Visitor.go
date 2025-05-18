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

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
	ast "github.com/craterdog/go-syntax-notation/v7/ast"
)

// CLASS INTERFACE

// Access Function

func VisitorClass() VisitorClassLike {
	return visitorClass()
}

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

// INSTANCE INTERFACE

// Principal Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorClass()
}

func (v *visitor_) VisitSyntax(
	syntax ast.SyntaxLike,
) {
	v.processor_.PreprocessSyntax(
		syntax,
		1,
		1,
	)
	v.visitSyntax(syntax)
	v.processor_.PostprocessSyntax(
		syntax,
		1,
		1,
	)
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAlternativeSequence(
	alternativeSequence ast.AlternativeSequenceLike,
) {
	var delimiter = alternativeSequence.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAlternativeSequenceSlot(1)

	var sequence = alternativeSequence.GetSequence()
	v.processor_.PreprocessSequence(
		sequence,
		1,
		1,
	)
	v.visitSequence(sequence)
	v.processor_.PostprocessSequence(
		sequence,
		1,
		1,
	)
}

func (v *visitor_) visitAlternatives(
	alternatives ast.AlternativesLike,
) {
	var sequence = alternatives.GetSequence()
	v.processor_.PreprocessSequence(
		sequence,
		1,
		1,
	)
	v.visitSequence(sequence)
	v.processor_.PostprocessSequence(
		sequence,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessAlternativesSlot(1)

	var alternativeSequencesIndex uint
	var alternativeSequences = alternatives.GetAlternativeSequences().GetIterator()
	var alternativeSequencesCount = uint(alternativeSequences.GetSize())
	for alternativeSequences.HasNext() {
		alternativeSequencesIndex++
		var rule = alternativeSequences.GetNext()
		v.processor_.PreprocessAlternativeSequence(
			rule,
			alternativeSequencesIndex,
			alternativeSequencesCount,
		)
		v.visitAlternativeSequence(rule)
		v.processor_.PostprocessAlternativeSequence(
			rule,
			alternativeSequencesIndex,
			alternativeSequencesCount,
		)
	}
}

func (v *visitor_) visitCardinality(
	cardinality ast.CardinalityLike,
) {
	// Visit the possible cardinality rule types.
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		v.processor_.PreprocessConstrained(
			actual,
			1,
			1,
		)
		v.visitConstrained(actual)
		v.processor_.PostprocessConstrained(
			actual,
			1,
			1,
		)
	case ast.QuantifiedLike:
		v.processor_.PreprocessQuantified(
			actual,
			1,
			1,
		)
		v.visitQuantified(actual)
		v.processor_.PostprocessQuantified(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitCharacter(
	character ast.CharacterLike,
) {
	// Visit the possible character rule types.
	switch actual := character.GetAny().(type) {
	case ast.ImplicitLike:
		v.processor_.PreprocessImplicit(
			actual,
			1,
			1,
		)
		v.visitImplicit(actual)
		v.processor_.PostprocessImplicit(
			actual,
			1,
			1,
		)
	case ast.ExplicitLike:
		v.processor_.PreprocessExplicit(
			actual,
			1,
			1,
		)
		v.visitExplicit(actual)
		v.processor_.PostprocessExplicit(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitComponent(
	component ast.ComponentLike,
) {
	// Visit the possible component expression types.
	var actual = component.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, LiteralToken):
		v.processor_.ProcessLiteral(actual)
	case ScannerClass().MatchesType(actual, LowercaseToken):
		v.processor_.ProcessLowercase(actual)
	case ScannerClass().MatchesType(actual, UppercaseToken):
		v.processor_.ProcessUppercase(actual)
	}
}

func (v *visitor_) visitConstrained(
	constrained ast.ConstrainedLike,
) {
	// Visit the possible constrained literal values.
	var actual = constrained.GetAny().(string)
	switch actual {
	case "?":
		v.processor_.ProcessDelimiter("?")
	case "*":
		v.processor_.ProcessDelimiter("*")
	case "+":
		v.processor_.ProcessDelimiter("+")
	}
}

func (v *visitor_) visitDefinition(
	definition ast.DefinitionLike,
) {
	// Visit the possible definition rule types.
	switch actual := definition.GetAny().(type) {
	case ast.LiteralAlternativesLike:
		v.processor_.PreprocessLiteralAlternatives(
			actual,
			1,
			1,
		)
		v.visitLiteralAlternatives(actual)
		v.processor_.PostprocessLiteralAlternatives(
			actual,
			1,
			1,
		)
	case ast.TokenAlternativesLike:
		v.processor_.PreprocessTokenAlternatives(
			actual,
			1,
			1,
		)
		v.visitTokenAlternatives(actual)
		v.processor_.PostprocessTokenAlternatives(
			actual,
			1,
			1,
		)
	case ast.RuleAlternativesLike:
		v.processor_.PreprocessRuleAlternatives(
			actual,
			1,
			1,
		)
		v.visitRuleAlternatives(actual)
		v.processor_.PostprocessRuleAlternatives(
			actual,
			1,
			1,
		)
	case ast.TermSequenceLike:
		v.processor_.PreprocessTermSequence(
			actual,
			1,
			1,
		)
		v.visitTermSequence(actual)
		v.processor_.PostprocessTermSequence(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitElement(
	element ast.ElementLike,
) {
	// Visit the possible element rule types.
	switch actual := element.GetAny().(type) {
	case ast.GroupLike:
		v.processor_.PreprocessGroup(
			actual,
			1,
			1,
		)
		v.visitGroup(actual)
		v.processor_.PostprocessGroup(
			actual,
			1,
			1,
		)
	case ast.FilterLike:
		v.processor_.PreprocessFilter(
			actual,
			1,
			1,
		)
		v.visitFilter(actual)
		v.processor_.PostprocessFilter(
			actual,
			1,
			1,
		)
	case ast.TextLike:
		v.processor_.PreprocessText(
			actual,
			1,
			1,
		)
		v.visitText(actual)
		v.processor_.PostprocessText(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitExplicit(
	explicit ast.ExplicitLike,
) {
	var glyph = explicit.GetGlyph()
	v.processor_.ProcessGlyph(glyph)
	// Visit slot 1 between terms.
	v.processor_.ProcessExplicitSlot(1)

	var optionalExtent = explicit.GetOptionalExtent()
	if uti.IsDefined(optionalExtent) {
		v.processor_.PreprocessExtent(
			optionalExtent,
			1,
			1,
		)
		v.visitExtent(optionalExtent)
		v.processor_.PostprocessExtent(
			optionalExtent,
			1,
			1,
		)
	}
}

func (v *visitor_) visitExpression(
	expression ast.ExpressionLike,
) {
	var delimiter1 = expression.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessExpressionSlot(1)

	var lowercase = expression.GetLowercase()
	v.processor_.ProcessLowercase(lowercase)
	// Visit slot 2 between terms.
	v.processor_.ProcessExpressionSlot(2)

	var delimiter2 = expression.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessExpressionSlot(3)

	var pattern = expression.GetPattern()
	v.processor_.PreprocessPattern(
		pattern,
		1,
		1,
	)
	v.visitPattern(pattern)
	v.processor_.PostprocessPattern(
		pattern,
		1,
		1,
	)
}

func (v *visitor_) visitExtent(
	extent ast.ExtentLike,
) {
	var delimiter = extent.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessExtentSlot(1)

	var glyph = extent.GetGlyph()
	v.processor_.ProcessGlyph(glyph)
}

func (v *visitor_) visitFilter(
	filter ast.FilterLike,
) {
	var optionalDelimiter = filter.GetOptionalDelimiter()
	if uti.IsDefined(optionalDelimiter) {
		v.processor_.ProcessDelimiter(optionalDelimiter)
	}
	// Visit slot 1 between terms.
	v.processor_.ProcessFilterSlot(1)

	var delimiter1 = filter.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessFilterSlot(2)

	var charactersIndex uint
	var characters = filter.GetCharacters().GetIterator()
	var charactersCount = uint(characters.GetSize())
	for characters.HasNext() {
		charactersIndex++
		var rule = characters.GetNext()
		v.processor_.PreprocessCharacter(
			rule,
			charactersIndex,
			charactersCount,
		)
		v.visitCharacter(rule)
		v.processor_.PostprocessCharacter(
			rule,
			charactersIndex,
			charactersCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessFilterSlot(3)

	var delimiter2 = filter.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitGroup(
	group ast.GroupLike,
) {
	var delimiter1 = group.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessGroupSlot(1)

	var alternatives = group.GetAlternatives()
	v.processor_.PreprocessAlternatives(
		alternatives,
		1,
		1,
	)
	v.visitAlternatives(alternatives)
	v.processor_.PostprocessAlternatives(
		alternatives,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessGroupSlot(2)

	var delimiter2 = group.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitImplicit(
	implicit ast.ImplicitLike,
) {
	var intrinsic = implicit.GetIntrinsic()
	v.processor_.ProcessIntrinsic(intrinsic)
}

func (v *visitor_) visitLimit(
	limit ast.LimitLike,
) {
	var delimiter = limit.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessLimitSlot(1)

	var optionalNumber = limit.GetOptionalNumber()
	if uti.IsDefined(optionalNumber) {
		v.processor_.ProcessNumber(optionalNumber)
	}
}

func (v *visitor_) visitLiteralAlternatives(
	literalAlternatives ast.LiteralAlternativesLike,
) {
	var literalValuesIndex uint
	var literalValues = literalAlternatives.GetLiteralValues().GetIterator()
	var literalValuesCount = uint(literalValues.GetSize())
	for literalValues.HasNext() {
		literalValuesIndex++
		var rule = literalValues.GetNext()
		v.processor_.PreprocessLiteralValue(
			rule,
			literalValuesIndex,
			literalValuesCount,
		)
		v.visitLiteralValue(rule)
		v.processor_.PostprocessLiteralValue(
			rule,
			literalValuesIndex,
			literalValuesCount,
		)
	}
}

func (v *visitor_) visitLiteralValue(
	literalValue ast.LiteralValueLike,
) {
	var newline = literalValue.GetNewline()
	v.processor_.ProcessNewline(newline)
	// Visit slot 1 between terms.
	v.processor_.ProcessLiteralValueSlot(1)

	var literal = literalValue.GetLiteral()
	v.processor_.ProcessLiteral(literal)
	// Visit slot 2 between terms.
	v.processor_.ProcessLiteralValueSlot(2)

	var optionalNote = literalValue.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitNotice(
	notice ast.NoticeLike,
) {
	var comment = notice.GetComment()
	v.processor_.ProcessComment(comment)
	// Visit slot 1 between terms.
	v.processor_.ProcessNoticeSlot(1)

	var newline = notice.GetNewline()
	v.processor_.ProcessNewline(newline)
}

func (v *visitor_) visitPattern(
	pattern ast.PatternLike,
) {
	var alternatives = pattern.GetAlternatives()
	v.processor_.PreprocessAlternatives(
		alternatives,
		1,
		1,
	)
	v.visitAlternatives(alternatives)
	v.processor_.PostprocessAlternatives(
		alternatives,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessPatternSlot(1)

	var optionalNote = pattern.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitQuantified(
	quantified ast.QuantifiedLike,
) {
	var delimiter1 = quantified.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessQuantifiedSlot(1)

	var number = quantified.GetNumber()
	v.processor_.ProcessNumber(number)
	// Visit slot 2 between terms.
	v.processor_.ProcessQuantifiedSlot(2)

	var optionalLimit = quantified.GetOptionalLimit()
	if uti.IsDefined(optionalLimit) {
		v.processor_.PreprocessLimit(
			optionalLimit,
			1,
			1,
		)
		v.visitLimit(optionalLimit)
		v.processor_.PostprocessLimit(
			optionalLimit,
			1,
			1,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessQuantifiedSlot(3)

	var delimiter2 = quantified.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitRepetition(
	repetition ast.RepetitionLike,
) {
	var element = repetition.GetElement()
	v.processor_.PreprocessElement(
		element,
		1,
		1,
	)
	v.visitElement(element)
	v.processor_.PostprocessElement(
		element,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessRepetitionSlot(1)

	var optionalCardinality = repetition.GetOptionalCardinality()
	if uti.IsDefined(optionalCardinality) {
		v.processor_.PreprocessCardinality(
			optionalCardinality,
			1,
			1,
		)
		v.visitCardinality(optionalCardinality)
		v.processor_.PostprocessCardinality(
			optionalCardinality,
			1,
			1,
		)
	}
}

func (v *visitor_) visitRule(
	rule ast.RuleLike,
) {
	var delimiter1 = rule.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessRuleSlot(1)

	var uppercase = rule.GetUppercase()
	v.processor_.ProcessUppercase(uppercase)
	// Visit slot 2 between terms.
	v.processor_.ProcessRuleSlot(2)

	var delimiter2 = rule.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessRuleSlot(3)

	var definition = rule.GetDefinition()
	v.processor_.PreprocessDefinition(
		definition,
		1,
		1,
	)
	v.visitDefinition(definition)
	v.processor_.PostprocessDefinition(
		definition,
		1,
		1,
	)
}

func (v *visitor_) visitRuleAlternatives(
	ruleAlternatives ast.RuleAlternativesLike,
) {
	var ruleNamesIndex uint
	var ruleNames = ruleAlternatives.GetRuleNames().GetIterator()
	var ruleNamesCount = uint(ruleNames.GetSize())
	for ruleNames.HasNext() {
		ruleNamesIndex++
		var rule = ruleNames.GetNext()
		v.processor_.PreprocessRuleName(
			rule,
			ruleNamesIndex,
			ruleNamesCount,
		)
		v.visitRuleName(rule)
		v.processor_.PostprocessRuleName(
			rule,
			ruleNamesIndex,
			ruleNamesCount,
		)
	}
}

func (v *visitor_) visitRuleName(
	ruleName ast.RuleNameLike,
) {
	var newline = ruleName.GetNewline()
	v.processor_.ProcessNewline(newline)
	// Visit slot 1 between terms.
	v.processor_.ProcessRuleNameSlot(1)

	var uppercase = ruleName.GetUppercase()
	v.processor_.ProcessUppercase(uppercase)
	// Visit slot 2 between terms.
	v.processor_.ProcessRuleNameSlot(2)

	var optionalNote = ruleName.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitRuleTerm(
	ruleTerm ast.RuleTermLike,
) {
	var component = ruleTerm.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		1,
		1,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessRuleTermSlot(1)

	var optionalCardinality = ruleTerm.GetOptionalCardinality()
	if uti.IsDefined(optionalCardinality) {
		v.processor_.PreprocessCardinality(
			optionalCardinality,
			1,
			1,
		)
		v.visitCardinality(optionalCardinality)
		v.processor_.PostprocessCardinality(
			optionalCardinality,
			1,
			1,
		)
	}
}

func (v *visitor_) visitSequence(
	sequence ast.SequenceLike,
) {
	var repetitionsIndex uint
	var repetitions = sequence.GetRepetitions().GetIterator()
	var repetitionsCount = uint(repetitions.GetSize())
	for repetitions.HasNext() {
		repetitionsIndex++
		var rule = repetitions.GetNext()
		v.processor_.PreprocessRepetition(
			rule,
			repetitionsIndex,
			repetitionsCount,
		)
		v.visitRepetition(rule)
		v.processor_.PostprocessRepetition(
			rule,
			repetitionsIndex,
			repetitionsCount,
		)
	}
}

func (v *visitor_) visitSyntax(
	syntax ast.SyntaxLike,
) {
	var notice = syntax.GetNotice()
	v.processor_.PreprocessNotice(
		notice,
		1,
		1,
	)
	v.visitNotice(notice)
	v.processor_.PostprocessNotice(
		notice,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessSyntaxSlot(1)

	var comment1 = syntax.GetComment1()
	v.processor_.ProcessComment(comment1)
	// Visit slot 2 between terms.
	v.processor_.ProcessSyntaxSlot(2)

	var rulesIndex uint
	var rules = syntax.GetRules().GetIterator()
	var rulesCount = uint(rules.GetSize())
	for rules.HasNext() {
		rulesIndex++
		var rule = rules.GetNext()
		v.processor_.PreprocessRule(
			rule,
			rulesIndex,
			rulesCount,
		)
		v.visitRule(rule)
		v.processor_.PostprocessRule(
			rule,
			rulesIndex,
			rulesCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessSyntaxSlot(3)

	var comment2 = syntax.GetComment2()
	v.processor_.ProcessComment(comment2)
	// Visit slot 4 between terms.
	v.processor_.ProcessSyntaxSlot(4)

	var expressionsIndex uint
	var expressions = syntax.GetExpressions().GetIterator()
	var expressionsCount = uint(expressions.GetSize())
	for expressions.HasNext() {
		expressionsIndex++
		var rule = expressions.GetNext()
		v.processor_.PreprocessExpression(
			rule,
			expressionsIndex,
			expressionsCount,
		)
		v.visitExpression(rule)
		v.processor_.PostprocessExpression(
			rule,
			expressionsIndex,
			expressionsCount,
		)
	}
}

func (v *visitor_) visitTermSequence(
	termSequence ast.TermSequenceLike,
) {
	var ruleTermsIndex uint
	var ruleTerms = termSequence.GetRuleTerms().GetIterator()
	var ruleTermsCount = uint(ruleTerms.GetSize())
	for ruleTerms.HasNext() {
		ruleTermsIndex++
		var rule = ruleTerms.GetNext()
		v.processor_.PreprocessRuleTerm(
			rule,
			ruleTermsIndex,
			ruleTermsCount,
		)
		v.visitRuleTerm(rule)
		v.processor_.PostprocessRuleTerm(
			rule,
			ruleTermsIndex,
			ruleTermsCount,
		)
	}
	// Visit slot 1 between terms.
	v.processor_.ProcessTermSequenceSlot(1)

	var optionalNote = termSequence.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitText(
	text ast.TextLike,
) {
	// Visit the possible text expression types.
	var actual = text.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, GlyphToken):
		v.processor_.ProcessGlyph(actual)
	case ScannerClass().MatchesType(actual, LiteralToken):
		v.processor_.ProcessLiteral(actual)
	case ScannerClass().MatchesType(actual, LowercaseToken):
		v.processor_.ProcessLowercase(actual)
	case ScannerClass().MatchesType(actual, IntrinsicToken):
		v.processor_.ProcessIntrinsic(actual)
	}
}

func (v *visitor_) visitTokenAlternatives(
	tokenAlternatives ast.TokenAlternativesLike,
) {
	var tokenNamesIndex uint
	var tokenNames = tokenAlternatives.GetTokenNames().GetIterator()
	var tokenNamesCount = uint(tokenNames.GetSize())
	for tokenNames.HasNext() {
		tokenNamesIndex++
		var rule = tokenNames.GetNext()
		v.processor_.PreprocessTokenName(
			rule,
			tokenNamesIndex,
			tokenNamesCount,
		)
		v.visitTokenName(rule)
		v.processor_.PostprocessTokenName(
			rule,
			tokenNamesIndex,
			tokenNamesCount,
		)
	}
}

func (v *visitor_) visitTokenName(
	tokenName ast.TokenNameLike,
) {
	var newline = tokenName.GetNewline()
	v.processor_.ProcessNewline(newline)
	// Visit slot 1 between terms.
	v.processor_.ProcessTokenNameSlot(1)

	var lowercase = tokenName.GetLowercase()
	v.processor_.ProcessLowercase(lowercase)
	// Visit slot 2 between terms.
	v.processor_.ProcessTokenNameSlot(2)

	var optionalNote = tokenName.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

// Instance Structure

type visitor_ struct {
	// Declare the instance attributes.
	processor_ Methodical
}

// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func visitorClass() *visitorClass_ {
	return visitorClassReference_
}

var visitorClassReference_ = &visitorClass_{
	// Initialize the class constants.
}
