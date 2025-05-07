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
	fmt "fmt"
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
	v.processor_.PreprocessSyntax(syntax)
	v.visitSyntax(syntax)
	v.processor_.PostprocessSyntax(syntax)
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAlternative(
	alternative ast.AlternativeLike,
) {
	// Visit a single option rule.
	var option = alternative.GetOption()
	v.processor_.PreprocessOption(option)
	v.visitOption(option)
	v.processor_.PostprocessOption(option)
}

func (v *visitor_) visitCardinality(
	cardinality ast.CardinalityLike,
) {
	// Visit the possible cardinality types.
	switch actual := cardinality.GetAny().(type) {
	case ast.ConstrainedLike:
		v.processor_.PreprocessConstrained(actual)
		v.visitConstrained(actual)
		v.processor_.PostprocessConstrained(actual)
	case ast.QuantifiedLike:
		v.processor_.PreprocessQuantified(actual)
		v.visitQuantified(actual)
		v.processor_.PostprocessQuantified(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitCharacter(
	character ast.CharacterLike,
) {
	// Visit the possible character types.
	switch actual := character.GetAny().(type) {
	case ast.ImplicitLike:
		v.processor_.PreprocessImplicit(actual)
		v.visitImplicit(actual)
		v.processor_.PostprocessImplicit(actual)
	case ast.ExplicitLike:
		v.processor_.PreprocessExplicit(actual)
		v.visitExplicit(actual)
		v.processor_.PostprocessExplicit(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitConstrained(
	constrained ast.ConstrainedLike,
) {
	// Visit the possible constrained types.
	switch actual := constrained.GetAny().(type) {
	case string:
		switch {
		case ScannerClass().MatchesType(actual, OptionalToken):
			v.processor_.ProcessOptional(actual)
		case ScannerClass().MatchesType(actual, RepeatedToken):
			v.processor_.ProcessRepeated(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitDefinition(
	definition ast.DefinitionLike,
) {
	// Visit the possible definition types.
	switch actual := definition.GetAny().(type) {
	case ast.MultiruleLike:
		v.processor_.PreprocessMultirule(actual)
		v.visitMultirule(actual)
		v.processor_.PostprocessMultirule(actual)
	case ast.MultiexpressionLike:
		v.processor_.PreprocessMultiexpression(actual)
		v.visitMultiexpression(actual)
		v.processor_.PostprocessMultiexpression(actual)
	case ast.MultiliteralLike:
		v.processor_.PreprocessMultiliteral(actual)
		v.visitMultiliteral(actual)
		v.processor_.PostprocessMultiliteral(actual)
	case ast.InlineLike:
		v.processor_.PreprocessInline(actual)
		v.visitInline(actual)
		v.processor_.PostprocessInline(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitElement(
	element ast.ElementLike,
) {
	// Visit the possible element types.
	switch actual := element.GetAny().(type) {
	case ast.GroupLike:
		v.processor_.PreprocessGroup(actual)
		v.visitGroup(actual)
		v.processor_.PostprocessGroup(actual)
	case ast.FilterLike:
		v.processor_.PreprocessFilter(actual)
		v.visitFilter(actual)
		v.processor_.PostprocessFilter(actual)
	case ast.TextLike:
		v.processor_.PreprocessText(actual)
		v.visitText(actual)
		v.processor_.PostprocessText(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitExplicit(
	explicit ast.ExplicitLike,
) {
	// Visit a single glyph token.
	var glyph = explicit.GetGlyph()
	v.processor_.ProcessGlyph(glyph)

	// Visit slot 1 between references.
	v.processor_.ProcessExplicitSlot(1)

	// Visit an optional extent rule.
	var optionalExtent = explicit.GetOptionalExtent()
	if uti.IsDefined(optionalExtent) {
		v.processor_.PreprocessExtent(optionalExtent)
		v.visitExtent(optionalExtent)
		v.processor_.PostprocessExtent(optionalExtent)
	}
}

func (v *visitor_) visitExpression(
	expression ast.ExpressionLike,
) {
	// Visit a single lowercase token.
	var lowercase = expression.GetLowercase()
	v.processor_.ProcessLowercase(lowercase)

	// Visit slot 1 between references.
	v.processor_.ProcessExpressionSlot(1)

	// Visit a single pattern rule.
	var pattern = expression.GetPattern()
	v.processor_.PreprocessPattern(pattern)
	v.visitPattern(pattern)
	v.processor_.PostprocessPattern(pattern)

	// Visit slot 2 between references.
	v.processor_.ProcessExpressionSlot(2)

	// Visit an optional note token.
	var optionalNote = expression.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitExpressionOption(
	expressionOption ast.ExpressionOptionLike,
) {
	// Visit a single newline token.
	var newline = expressionOption.GetNewline()
	v.processor_.ProcessNewline(newline)

	// Visit slot 1 between references.
	v.processor_.ProcessExpressionOptionSlot(1)

	// Visit a single lowercase token.
	var lowercase = expressionOption.GetLowercase()
	v.processor_.ProcessLowercase(lowercase)

	// Visit slot 2 between references.
	v.processor_.ProcessExpressionOptionSlot(2)

	// Visit an optional note token.
	var optionalNote = expressionOption.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitExtent(
	extent ast.ExtentLike,
) {
	// Visit a single glyph token.
	var glyph = extent.GetGlyph()
	v.processor_.ProcessGlyph(glyph)
}

func (v *visitor_) visitFilter(
	filter ast.FilterLike,
) {
	// Visit an optional excluded token.
	var optionalExcluded = filter.GetOptionalExcluded()
	if uti.IsDefined(optionalExcluded) {
		v.processor_.ProcessExcluded(optionalExcluded)
	}

	// Visit slot 1 between references.
	v.processor_.ProcessFilterSlot(1)

	// Visit each character rule.
	var characterIndex uint
	var characters = filter.GetCharacters().GetIterator()
	var charactersSize = uint(characters.GetSize())
	for characters.HasNext() {
		characterIndex++
		var character = characters.GetNext()
		v.processor_.PreprocessCharacter(
			character,
			characterIndex,
			charactersSize,
		)
		v.visitCharacter(character)
		v.processor_.PostprocessCharacter(
			character,
			characterIndex,
			charactersSize,
		)
	}
}

func (v *visitor_) visitGroup(
	group ast.GroupLike,
) {
	// Visit a single pattern rule.
	var pattern = group.GetPattern()
	v.processor_.PreprocessPattern(pattern)
	v.visitPattern(pattern)
	v.processor_.PostprocessPattern(pattern)
}

func (v *visitor_) visitIdentifier(
	identifier ast.IdentifierLike,
) {
	// Visit the possible identifier types.
	switch actual := identifier.GetAny().(type) {
	case string:
		switch {
		case ScannerClass().MatchesType(actual, LowercaseToken):
			v.processor_.ProcessLowercase(actual)
		case ScannerClass().MatchesType(actual, UppercaseToken):
			v.processor_.ProcessUppercase(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitImplicit(
	implicit ast.ImplicitLike,
) {
	// Visit a single intrinsic token.
	var intrinsic = implicit.GetIntrinsic()
	v.processor_.ProcessIntrinsic(intrinsic)
}

func (v *visitor_) visitInline(
	inline ast.InlineLike,
) {
	// Visit each term rule.
	var termIndex uint
	var terms = inline.GetTerms().GetIterator()
	var termsSize = uint(terms.GetSize())
	for terms.HasNext() {
		termIndex++
		var term = terms.GetNext()
		v.processor_.PreprocessTerm(
			term,
			termIndex,
			termsSize,
		)
		v.visitTerm(term)
		v.processor_.PostprocessTerm(
			term,
			termIndex,
			termsSize,
		)
	}

	// Visit slot 1 between references.
	v.processor_.ProcessInlineSlot(1)

	// Visit an optional note token.
	var optionalNote = inline.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitLimit(
	limit ast.LimitLike,
) {
	// Visit an optional number token.
	var optionalNumber = limit.GetOptionalNumber()
	if uti.IsDefined(optionalNumber) {
		v.processor_.ProcessNumber(optionalNumber)
	}
}

func (v *visitor_) visitLiteral(
	literal ast.LiteralLike,
) {
	// Visit a single quote token.
	var quote = literal.GetQuote()
	v.processor_.ProcessQuote(quote)
}

func (v *visitor_) visitLiteralOption(
	literalOption ast.LiteralOptionLike,
) {
	// Visit a single newline token.
	var newline = literalOption.GetNewline()
	v.processor_.ProcessNewline(newline)

	// Visit slot 1 between references.
	v.processor_.ProcessLiteralOptionSlot(1)

	// Visit a single quote token.
	var quote = literalOption.GetQuote()
	v.processor_.ProcessQuote(quote)

	// Visit slot 2 between references.
	v.processor_.ProcessLiteralOptionSlot(2)

	// Visit an optional note token.
	var optionalNote = literalOption.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitMultiexpression(
	multiexpression ast.MultiexpressionLike,
) {
	// Visit each expressionOption rule.
	var expressionOptionIndex uint
	var expressionOptions = multiexpression.GetExpressionOptions().GetIterator()
	var expressionOptionsSize = uint(expressionOptions.GetSize())
	for expressionOptions.HasNext() {
		expressionOptionIndex++
		var expressionOption = expressionOptions.GetNext()
		v.processor_.PreprocessExpressionOption(
			expressionOption,
			expressionOptionIndex,
			expressionOptionsSize,
		)
		v.visitExpressionOption(expressionOption)
		v.processor_.PostprocessExpressionOption(
			expressionOption,
			expressionOptionIndex,
			expressionOptionsSize,
		)
	}
}

func (v *visitor_) visitMultiliteral(
	multiliteral ast.MultiliteralLike,
) {
	// Visit each literalOption rule.
	var literalOptionIndex uint
	var literalOptions = multiliteral.GetLiteralOptions().GetIterator()
	var literalOptionsSize = uint(literalOptions.GetSize())
	for literalOptions.HasNext() {
		literalOptionIndex++
		var literalOption = literalOptions.GetNext()
		v.processor_.PreprocessLiteralOption(
			literalOption,
			literalOptionIndex,
			literalOptionsSize,
		)
		v.visitLiteralOption(literalOption)
		v.processor_.PostprocessLiteralOption(
			literalOption,
			literalOptionIndex,
			literalOptionsSize,
		)
	}
}

func (v *visitor_) visitMultirule(
	multirule ast.MultiruleLike,
) {
	// Visit each ruleOption rule.
	var ruleOptionIndex uint
	var ruleOptions = multirule.GetRuleOptions().GetIterator()
	var ruleOptionsSize = uint(ruleOptions.GetSize())
	for ruleOptions.HasNext() {
		ruleOptionIndex++
		var ruleOption = ruleOptions.GetNext()
		v.processor_.PreprocessRuleOption(
			ruleOption,
			ruleOptionIndex,
			ruleOptionsSize,
		)
		v.visitRuleOption(ruleOption)
		v.processor_.PostprocessRuleOption(
			ruleOption,
			ruleOptionIndex,
			ruleOptionsSize,
		)
	}
}

func (v *visitor_) visitNotice(
	notice ast.NoticeLike,
) {
	// Visit a single comment token.
	var comment = notice.GetComment()
	v.processor_.ProcessComment(comment)

	// Visit slot 1 between references.
	v.processor_.ProcessNoticeSlot(1)

	// Visit a single newline token.
	var newline = notice.GetNewline()
	v.processor_.ProcessNewline(newline)
}

func (v *visitor_) visitOption(
	option ast.OptionLike,
) {
	// Visit each repetition rule.
	var repetitionIndex uint
	var repetitions = option.GetRepetitions().GetIterator()
	var repetitionsSize = uint(repetitions.GetSize())
	for repetitions.HasNext() {
		repetitionIndex++
		var repetition = repetitions.GetNext()
		v.processor_.PreprocessRepetition(
			repetition,
			repetitionIndex,
			repetitionsSize,
		)
		v.visitRepetition(repetition)
		v.processor_.PostprocessRepetition(
			repetition,
			repetitionIndex,
			repetitionsSize,
		)
	}
}

func (v *visitor_) visitPattern(
	pattern ast.PatternLike,
) {
	// Visit a single option rule.
	var option = pattern.GetOption()
	v.processor_.PreprocessOption(option)
	v.visitOption(option)
	v.processor_.PostprocessOption(option)

	// Visit slot 1 between references.
	v.processor_.ProcessPatternSlot(1)

	// Visit each alternative rule.
	var alternativeIndex uint
	var alternatives = pattern.GetAlternatives().GetIterator()
	var alternativesSize = uint(alternatives.GetSize())
	for alternatives.HasNext() {
		alternativeIndex++
		var alternative = alternatives.GetNext()
		v.processor_.PreprocessAlternative(
			alternative,
			alternativeIndex,
			alternativesSize,
		)
		v.visitAlternative(alternative)
		v.processor_.PostprocessAlternative(
			alternative,
			alternativeIndex,
			alternativesSize,
		)
	}
}

func (v *visitor_) visitQuantified(
	quantified ast.QuantifiedLike,
) {
	// Visit a single number token.
	var number = quantified.GetNumber()
	v.processor_.ProcessNumber(number)

	// Visit slot 1 between references.
	v.processor_.ProcessQuantifiedSlot(1)

	// Visit an optional limit rule.
	var optionalLimit = quantified.GetOptionalLimit()
	if uti.IsDefined(optionalLimit) {
		v.processor_.PreprocessLimit(optionalLimit)
		v.visitLimit(optionalLimit)
		v.processor_.PostprocessLimit(optionalLimit)
	}
}

func (v *visitor_) visitReference(
	reference ast.ReferenceLike,
) {
	// Visit a single identifier rule.
	var identifier = reference.GetIdentifier()
	v.processor_.PreprocessIdentifier(identifier)
	v.visitIdentifier(identifier)
	v.processor_.PostprocessIdentifier(identifier)

	// Visit slot 1 between references.
	v.processor_.ProcessReferenceSlot(1)

	// Visit an optional cardinality rule.
	var optionalCardinality = reference.GetOptionalCardinality()
	if uti.IsDefined(optionalCardinality) {
		v.processor_.PreprocessCardinality(optionalCardinality)
		v.visitCardinality(optionalCardinality)
		v.processor_.PostprocessCardinality(optionalCardinality)
	}
}

func (v *visitor_) visitRepetition(
	repetition ast.RepetitionLike,
) {
	// Visit a single element rule.
	var element = repetition.GetElement()
	v.processor_.PreprocessElement(element)
	v.visitElement(element)
	v.processor_.PostprocessElement(element)

	// Visit slot 1 between references.
	v.processor_.ProcessRepetitionSlot(1)

	// Visit an optional cardinality rule.
	var optionalCardinality = repetition.GetOptionalCardinality()
	if uti.IsDefined(optionalCardinality) {
		v.processor_.PreprocessCardinality(optionalCardinality)
		v.visitCardinality(optionalCardinality)
		v.processor_.PostprocessCardinality(optionalCardinality)
	}
}

func (v *visitor_) visitRule(
	rule ast.RuleLike,
) {
	// Visit a single uppercase token.
	var uppercase = rule.GetUppercase()
	v.processor_.ProcessUppercase(uppercase)

	// Visit slot 1 between references.
	v.processor_.ProcessRuleSlot(1)

	// Visit a single definition rule.
	var definition = rule.GetDefinition()
	v.processor_.PreprocessDefinition(definition)
	v.visitDefinition(definition)
	v.processor_.PostprocessDefinition(definition)
}

func (v *visitor_) visitRuleOption(
	ruleOption ast.RuleOptionLike,
) {
	// Visit a single newline token.
	var newline = ruleOption.GetNewline()
	v.processor_.ProcessNewline(newline)

	// Visit slot 1 between references.
	v.processor_.ProcessRuleOptionSlot(1)

	// Visit a single uppercase token.
	var uppercase = ruleOption.GetUppercase()
	v.processor_.ProcessUppercase(uppercase)

	// Visit slot 2 between references.
	v.processor_.ProcessRuleOptionSlot(2)

	// Visit an optional note token.
	var optionalNote = ruleOption.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitSyntax(
	syntax ast.SyntaxLike,
) {
	// Visit a single notice rule.
	var notice = syntax.GetNotice()
	v.processor_.PreprocessNotice(notice)
	v.visitNotice(notice)
	v.processor_.PostprocessNotice(notice)

	// Visit slot 1 between references.
	v.processor_.ProcessSyntaxSlot(1)

	// Visit a single comment token.
	var comment1 = syntax.GetComment1()
	v.processor_.ProcessComment(comment1)

	// Visit slot 2 between references.
	v.processor_.ProcessSyntaxSlot(2)

	// Visit each rule rule.
	var ruleIndex uint
	var rules = syntax.GetRules().GetIterator()
	var rulesSize = uint(rules.GetSize())
	for rules.HasNext() {
		ruleIndex++
		var rule = rules.GetNext()
		v.processor_.PreprocessRule(
			rule,
			ruleIndex,
			rulesSize,
		)
		v.visitRule(rule)
		v.processor_.PostprocessRule(
			rule,
			ruleIndex,
			rulesSize,
		)
	}

	// Visit slot 3 between references.
	v.processor_.ProcessSyntaxSlot(3)

	// Visit a single comment token.
	var comment2 = syntax.GetComment2()
	v.processor_.ProcessComment(comment2)

	// Visit slot 4 between references.
	v.processor_.ProcessSyntaxSlot(4)

	// Visit each expression rule.
	var expressionIndex uint
	var expressions = syntax.GetExpressions().GetIterator()
	var expressionsSize = uint(expressions.GetSize())
	for expressions.HasNext() {
		expressionIndex++
		var expression = expressions.GetNext()
		v.processor_.PreprocessExpression(
			expression,
			expressionIndex,
			expressionsSize,
		)
		v.visitExpression(expression)
		v.processor_.PostprocessExpression(
			expression,
			expressionIndex,
			expressionsSize,
		)
	}
}

func (v *visitor_) visitTerm(
	term ast.TermLike,
) {
	// Visit the possible term types.
	switch actual := term.GetAny().(type) {
	case ast.LiteralLike:
		v.processor_.PreprocessLiteral(actual)
		v.visitLiteral(actual)
		v.processor_.PostprocessLiteral(actual)
	case ast.ReferenceLike:
		v.processor_.PreprocessReference(actual)
		v.visitReference(actual)
		v.processor_.PostprocessReference(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitText(
	text ast.TextLike,
) {
	// Visit the possible text types.
	switch actual := text.GetAny().(type) {
	case string:
		switch {
		case ScannerClass().MatchesType(actual, GlyphToken):
			v.processor_.ProcessGlyph(actual)
		case ScannerClass().MatchesType(actual, QuoteToken):
			v.processor_.ProcessQuote(actual)
		case ScannerClass().MatchesType(actual, LowercaseToken):
			v.processor_.ProcessLowercase(actual)
		case ScannerClass().MatchesType(actual, IntrinsicToken):
			v.processor_.ProcessIntrinsic(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
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
