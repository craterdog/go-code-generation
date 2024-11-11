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
	ast "github.com/craterdog/go-syntax-notation/v5/ast"
)

// CLASS INTERFACE

// Access Function

func Validator() ValidatorClassLike {
	return validatorReference()
}

// Constructor Methods

func (c *validatorClass_) Make() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: Processor().Make(),
	}
	instance.visitor_ = Visitor().Make(instance)
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorReference()
}

func (v *validator_) ValidateSyntax(
	syntax ast.SyntaxLike,
) {
	v.visitor_.VisitSyntax(syntax)
}

// Methodical Methods

func (v *validator_) ProcessComment(
	comment string,
) {
	v.validateToken(comment, CommentToken)
}

func (v *validator_) ProcessExcluded(
	excluded string,
) {
	v.validateToken(excluded, ExcludedToken)
}

func (v *validator_) ProcessGlyph(
	glyph string,
) {
	v.validateToken(glyph, GlyphToken)
}

func (v *validator_) ProcessIntrinsic(
	intrinsic string,
) {
	v.validateToken(intrinsic, IntrinsicToken)
}

func (v *validator_) ProcessLiteral(
	literal string,
) {
	v.validateToken(literal, LiteralToken)
}

func (v *validator_) ProcessLowercase(
	lowercase string,
) {
	v.validateToken(lowercase, LowercaseToken)
}

func (v *validator_) ProcessNewline(
	newline string,
) {
	v.validateToken(newline, NewlineToken)
}

func (v *validator_) ProcessNote(
	note string,
) {
	v.validateToken(note, NoteToken)
}

func (v *validator_) ProcessNumber(
	number string,
) {
	v.validateToken(number, NumberToken)
}

func (v *validator_) ProcessOptional(
	optional string,
) {
	v.validateToken(optional, OptionalToken)
}

func (v *validator_) ProcessRepeated(
	repeated string,
) {
	v.validateToken(repeated, RepeatedToken)
}

func (v *validator_) ProcessSpace(
	space string,
) {
	v.validateToken(space, SpaceToken)
}

func (v *validator_) ProcessUppercase(
	uppercase string,
) {
	v.validateToken(uppercase, UppercaseToken)
}

func (v *validator_) PreprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAlternativeSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCardinality(
	cardinality ast.CardinalityLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessCardinalitySlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessCardinality(
	cardinality ast.CardinalityLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCharacter(
	character ast.CharacterLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessCharacterSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessCharacter(
	character ast.CharacterLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstrained(
	constrained ast.ConstrainedLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessConstrainedSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstrained(
	constrained ast.ConstrainedLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDefinition(
	definition ast.DefinitionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessDefinitionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessDefinition(
	definition ast.DefinitionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessElement(
	element ast.ElementLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessElementSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessElement(
	element ast.ElementLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExplicit(
	explicit ast.ExplicitLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessExplicitSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessExplicit(
	explicit ast.ExplicitLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessExpressionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExtent(
	extent ast.ExtentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessExtentSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessExtent(
	extent ast.ExtentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFilter(
	filter ast.FilterLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessFilterSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFilter(
	filter ast.FilterLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessGroup(
	group ast.GroupLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessGroupSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessGroup(
	group ast.GroupLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIdentifier(
	identifier ast.IdentifierLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessIdentifierSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessIdentifier(
	identifier ast.IdentifierLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInline(
	inline ast.InlineLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInlineSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInline(
	inline ast.InlineLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLimit(
	limit ast.LimitLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLimitSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLimit(
	limit ast.LimitLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLineSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLine(
	line ast.LineLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultiline(
	multiline ast.MultilineLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMultilineSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMultiline(
	multiline ast.MultilineLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNotice(
	notice ast.NoticeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessNoticeSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNotice(
	notice ast.NoticeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessOption(
	option ast.OptionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessOptionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessOption(
	option ast.OptionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPattern(
	pattern ast.PatternLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPatternSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPattern(
	pattern ast.PatternLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessQuantified(
	quantified ast.QuantifiedLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessQuantifiedSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessQuantified(
	quantified ast.QuantifiedLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessReference(
	reference ast.ReferenceLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessReferenceSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessReference(
	reference ast.ReferenceLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRepetition(
	repetition ast.RepetitionLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessRepetitionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRepetition(
	repetition ast.RepetitionLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessRuleSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSyntax(
	syntax ast.SyntaxLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessSyntaxSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSyntax(
	syntax ast.SyntaxLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTerm(
	term ast.TermLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessTermSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessTerm(
	term ast.TermLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessText(
	text ast.TextLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessTextSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessText(
	text ast.TextLike,
) {
	// TBD - Add any validation checks.
}

// PROTECTED INTERFACE

// Private Methods

func (v *validator_) validateToken(
	tokenValue string,
	tokenType TokenType,
) {
	if !Scanner().MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			Scanner().FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}

// Instance Structure

type validator_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type validatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func validatorReference() *validatorClass_ {
	return validatorReference_
}

var validatorReference_ = &validatorClass_{
	// Initialize the class constants.
}
