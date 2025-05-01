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
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	ast "github.com/craterdog/go-syntax-notation/v7/ast"
)

// CLASS INTERFACE

// Access Function

func ValidatorClass() ValidatorClassLike {
	return validatorClass()
}

// Constructor Methods

func (c *validatorClass_) Validator() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorClass()
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

func (v *validator_) ProcessQuote(
	quote string,
) {
	v.validateToken(quote, QuoteToken)
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

func (v *validator_) PreprocessExpressionOption(
	expressionOption ast.ExpressionOptionLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessExpressionOptionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessExpressionOption(
	expressionOption ast.ExpressionOptionLike,
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

func (v *validator_) PreprocessImplicit(
	implicit ast.ImplicitLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessImplicitSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessImplicit(
	implicit ast.ImplicitLike,
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

func (v *validator_) PreprocessLiteral(
	literal ast.LiteralLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLiteralSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLiteral(
	literal ast.LiteralLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultiexpression(
	multiexpression ast.MultiexpressionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMultiexpressionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMultiexpression(
	multiexpression ast.MultiexpressionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultirule(
	multirule ast.MultiruleLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMultiruleSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMultirule(
	multirule ast.MultiruleLike,
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

func (v *validator_) PreprocessRuleOption(
	ruleOption ast.RuleOptionLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessRuleOptionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRuleOption(
	ruleOption ast.RuleOptionLike,
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
	var scannerClass = ScannerClass()
	if !scannerClass.MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			scannerClass.FormatType(tokenType),
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

func validatorClass() *validatorClass_ {
	return validatorClassReference_
}

var validatorClassReference_ = &validatorClass_{
	// Initialize the class constants.
}
