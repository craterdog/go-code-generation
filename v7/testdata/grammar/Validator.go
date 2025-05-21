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
	col "github.com/craterdog/go-collection-framework/v7"
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
		ruleNames_:   col.Set[string](),
		tokenNames_:  col.Set[string](),
		rules_:       col.Set[string](),
		expressions_: col.SetFromArray[string]([]string{"newline"}),

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
	var ruleNames = v.ruleNames_.GetIterator()
	for ruleNames.HasNext() {
		var ruleName = ruleNames.GetNext()
		if !v.rules_.ContainsValue(ruleName) {
			var message = fmt.Sprintf(
				"The following rule does not have an associated definition: %s",
				ruleName,
			)
			panic(message)
		}
	}
	var tokenNames = v.tokenNames_.GetIterator()
	for tokenNames.HasNext() {
		var tokenName = tokenNames.GetNext()
		if !v.expressions_.ContainsValue(tokenName) {
			var message = fmt.Sprintf(
				"The following expression does not have an associated definition: %s",
				tokenName,
			)
			panic(message)
		}
	}
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

func (v *validator_) PreprocessComponent(
	component ast.ComponentLike,
	index_ uint,
	count_ uint,
) {
	var scannerClass = ScannerClass()
	var actual = component.GetAny().(string)
	switch {
	case scannerClass.MatchesType(actual, UppercaseToken):
		v.ruleNames_.AddValue(actual)
	case scannerClass.MatchesType(actual, LowercaseToken):
		v.tokenNames_.AddValue(actual)
	}
}

func (v *validator_) PreprocessExpression(
	expression ast.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	var expressionName = expression.GetLowercase()
	v.expressions_.AddValue(expressionName)
}

func (v *validator_) PreprocessRule(
	rule ast.RuleLike,
	index_ uint,
	count_ uint,
) {
	var ruleName = rule.GetUppercase()
	v.rules_.AddValue(ruleName)
}

func (v *validator_) PreprocessRuleName(
	ruleName ast.RuleNameLike,
	index_ uint,
	count_ uint,
) {
	var uppercase = ruleName.GetUppercase()
	v.ruleNames_.AddValue(uppercase)
}

func (v *validator_) PreprocessTokenName(
	tokenName ast.TokenNameLike,
	index_ uint,
	count_ uint,
) {
	var lowercase = tokenName.GetLowercase()
	v.tokenNames_.AddValue(lowercase)
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
	visitor_     VisitorLike
	ruleNames_   col.SetLike[string]
	tokenNames_  col.SetLike[string]
	rules_       col.SetLike[string]
	expressions_ col.SetLike[string]

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
