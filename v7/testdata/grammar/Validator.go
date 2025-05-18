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

func (v *validator_) PreprocessAdditionalCharacter(
	additionalCharacter ast.AdditionalCharacterLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalCharacter(
	additionalCharacter ast.AdditionalCharacterLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalRepetition(
	additionalRepetition ast.AdditionalRepetitionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalRepetition(
	additionalRepetition ast.AdditionalRepetitionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAllowedCharacters(
	allowedCharacters ast.AllowedCharactersLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAllowedCharacters(
	allowedCharacters ast.AllowedCharactersLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAlternativeSequence(
	alternativeSequence ast.AlternativeSequenceLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAlternativeSequence(
	alternativeSequence ast.AlternativeSequenceLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAlternatives(
	alternatives ast.AlternativesLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAlternatives(
	alternatives ast.AlternativesLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCardinality(
	cardinality ast.CardinalityLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessCardinality(
	cardinality ast.CardinalityLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCharacter(
	character ast.CharacterLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessCharacter(
	character ast.CharacterLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessComponent(
	component ast.ComponentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessComponent(
	component ast.ComponentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConstrained(
	constrained ast.ConstrainedLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConstrained(
	constrained ast.ConstrainedLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDefinition(
	definition ast.DefinitionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessDefinition(
	definition ast.DefinitionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessElement(
	element ast.ElementLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessElement(
	element ast.ElementLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExplicit(
	explicit ast.ExplicitLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessExplicit(
	explicit ast.ExplicitLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExtent(
	extent ast.ExtentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessExtent(
	extent ast.ExtentLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFilter(
	filter ast.FilterLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFilter(
	filter ast.FilterLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessGroup(
	group ast.GroupLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessGroup(
	group ast.GroupLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessImplicit(
	implicit ast.ImplicitLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessImplicit(
	implicit ast.ImplicitLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLimit(
	limit ast.LimitLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLimit(
	limit ast.LimitLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLiteralValue(
	literalValue ast.LiteralValueLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLiteralValue(
	literalValue ast.LiteralValueLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLiteralValueAlternatives(
	literalValueAlternatives ast.LiteralValueAlternativesLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLiteralValueAlternatives(
	literalValueAlternatives ast.LiteralValueAlternativesLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNotice(
	notice ast.NoticeLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNotice(
	notice ast.NoticeLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPattern(
	pattern ast.PatternLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPattern(
	pattern ast.PatternLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessQuantified(
	quantified ast.QuantifiedLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessQuantified(
	quantified ast.QuantifiedLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRepetition(
	repetition ast.RepetitionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRepetition(
	repetition ast.RepetitionLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRule(
	rule ast.RuleLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRuleName(
	ruleName ast.RuleNameLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRuleName(
	ruleName ast.RuleNameLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRuleNameAlternatives(
	ruleNameAlternatives ast.RuleNameAlternativesLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRuleNameAlternatives(
	ruleNameAlternatives ast.RuleNameAlternativesLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRuleTerm(
	ruleTerm ast.RuleTermLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRuleTerm(
	ruleTerm ast.RuleTermLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRuleTermSequence(
	ruleTermSequence ast.RuleTermSequenceLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRuleTermSequence(
	ruleTermSequence ast.RuleTermSequenceLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSequence(
	sequence ast.SequenceLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSequence(
	sequence ast.SequenceLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSyntax(
	syntax ast.SyntaxLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSyntax(
	syntax ast.SyntaxLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessText(
	text ast.TextLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessText(
	text ast.TextLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTokenName(
	tokenName ast.TokenNameLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessTokenName(
	tokenName ast.TokenNameLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTokenNameAlternatives(
	tokenNameAlternatives ast.TokenNameAlternativesLike,
	index uint,
	count uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessTokenNameAlternatives(
	tokenNameAlternatives ast.TokenNameAlternativesLike,
	index uint,
	count uint,
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
