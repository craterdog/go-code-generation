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

package grammar

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v2"
	ast "github.com/craterdog/go-syntax-notation/v5/ast"
)

// CLASS INTERFACE

// Access Function

func Visitor() VisitorClassLike {
	return visitorReference()
}

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

// INSTANCE INTERFACE

// Primary Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorReference()
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

func (v *visitor_) visitAlternative(alternative ast.AlternativeLike) {
	// Visit a single option rule.
	var option = alternative.GetOption()
	v.processor_.PreprocessOption(option)
	v.visitOption(option)
	v.processor_.PostprocessOption(option)
}

func (v *visitor_) visitCardinality(cardinality ast.CardinalityLike) {
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

func (v *visitor_) visitCharacter(character ast.CharacterLike) {
	// Visit the possible character types.
	switch actual := character.GetAny().(type) {
	case ast.ExplicitLike:
		v.processor_.PreprocessExplicit(actual)
		v.visitExplicit(actual)
		v.processor_.PostprocessExplicit(actual)
	case string:
		switch {
		case Scanner().MatchesType(actual, IntrinsicToken):
			v.processor_.ProcessIntrinsic(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitConstrained(constrained ast.ConstrainedLike) {
	// Visit the possible constrained types.
	switch actual := constrained.GetAny().(type) {
	case string:
		switch {
		case Scanner().MatchesType(actual, OptionalToken):
			v.processor_.ProcessOptional(actual)
		case Scanner().MatchesType(actual, RepeatedToken):
			v.processor_.ProcessRepeated(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitDefinition(definition ast.DefinitionLike) {
	// Visit the possible definition types.
	switch actual := definition.GetAny().(type) {
	case ast.MultilineLike:
		v.processor_.PreprocessMultiline(actual)
		v.visitMultiline(actual)
		v.processor_.PostprocessMultiline(actual)
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

func (v *visitor_) visitElement(element ast.ElementLike) {
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

func (v *visitor_) visitExplicit(explicit ast.ExplicitLike) {
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

func (v *visitor_) visitExpression(expression ast.ExpressionLike) {
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

func (v *visitor_) visitExtent(extent ast.ExtentLike) {
	// Visit a single glyph token.
	var glyph = extent.GetGlyph()
	v.processor_.ProcessGlyph(glyph)
}

func (v *visitor_) visitFilter(filter ast.FilterLike) {
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

func (v *visitor_) visitGroup(group ast.GroupLike) {
	// Visit a single pattern rule.
	var pattern = group.GetPattern()
	v.processor_.PreprocessPattern(pattern)
	v.visitPattern(pattern)
	v.processor_.PostprocessPattern(pattern)
}

func (v *visitor_) visitIdentifier(identifier ast.IdentifierLike) {
	// Visit the possible identifier types.
	switch actual := identifier.GetAny().(type) {
	case string:
		switch {
		case Scanner().MatchesType(actual, LowercaseToken):
			v.processor_.ProcessLowercase(actual)
		case Scanner().MatchesType(actual, UppercaseToken):
			v.processor_.ProcessUppercase(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitInline(inline ast.InlineLike) {
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

func (v *visitor_) visitLimit(limit ast.LimitLike) {
	// Visit an optional number token.
	var optionalNumber = limit.GetOptionalNumber()
	if uti.IsDefined(optionalNumber) {
		v.processor_.ProcessNumber(optionalNumber)
	}
}

func (v *visitor_) visitLine(line ast.LineLike) {
	// Visit a single identifier rule.
	var identifier = line.GetIdentifier()
	v.processor_.PreprocessIdentifier(identifier)
	v.visitIdentifier(identifier)
	v.processor_.PostprocessIdentifier(identifier)

	// Visit slot 1 between references.
	v.processor_.ProcessLineSlot(1)

	// Visit an optional note token.
	var optionalNote = line.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitMultiline(multiline ast.MultilineLike) {
	// Visit each line rule.
	var lineIndex uint
	var lines = multiline.GetLines().GetIterator()
	var linesSize = uint(lines.GetSize())
	for lines.HasNext() {
		lineIndex++
		var line = lines.GetNext()
		v.processor_.PreprocessLine(
			line,
			lineIndex,
			linesSize,
		)
		v.visitLine(line)
		v.processor_.PostprocessLine(
			line,
			lineIndex,
			linesSize,
		)
	}
}

func (v *visitor_) visitNotice(notice ast.NoticeLike) {
	// Visit a single comment token.
	var comment = notice.GetComment()
	v.processor_.ProcessComment(comment)
}

func (v *visitor_) visitOption(option ast.OptionLike) {
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

func (v *visitor_) visitPattern(pattern ast.PatternLike) {
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

func (v *visitor_) visitQuantified(quantified ast.QuantifiedLike) {
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

func (v *visitor_) visitReference(reference ast.ReferenceLike) {
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

func (v *visitor_) visitRepetition(repetition ast.RepetitionLike) {
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

func (v *visitor_) visitRule(rule ast.RuleLike) {
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

func (v *visitor_) visitSyntax(syntax ast.SyntaxLike) {
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

func (v *visitor_) visitTerm(term ast.TermLike) {
	// Visit the possible term types.
	switch actual := term.GetAny().(type) {
	case ast.ReferenceLike:
		v.processor_.PreprocessReference(actual)
		v.visitReference(actual)
		v.processor_.PostprocessReference(actual)
	case string:
		switch {
		case Scanner().MatchesType(actual, LiteralToken):
			v.processor_.ProcessLiteral(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitText(text ast.TextLike) {
	// Visit the possible text types.
	switch actual := text.GetAny().(type) {
	case string:
		switch {
		case Scanner().MatchesType(actual, IntrinsicToken):
			v.processor_.ProcessIntrinsic(actual)
		case Scanner().MatchesType(actual, GlyphToken):
			v.processor_.ProcessGlyph(actual)
		case Scanner().MatchesType(actual, LiteralToken):
			v.processor_.ProcessLiteral(actual)
		case Scanner().MatchesType(actual, LowercaseToken):
			v.processor_.ProcessLowercase(actual)
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

func visitorReference() *visitorClass_ {
	return visitorReference_
}

var visitorReference_ = &visitorClass_{
	// Initialize the class constants.
}
