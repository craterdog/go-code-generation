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
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	fra "github.com/craterdog/go-collection-framework/v8"
	ast "github.com/craterdog/go-syntax-notation/v8/ast"
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
		ruleNames_:       fra.Set[string](),
		expressionNames_: fra.Set[string](),
		fragmentNames_:   fra.Set[string](),
		rules_:           fra.Set[string](),
		expressions_: fra.SetFromArray[string](
			[]string{"newline"},
		),
		fragments_: fra.SetFromArray[string](
			[]string{"ANY", "CONTROL", "DIGIT", "EOL", "LOWER", "UPPER"},
		),

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
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
	VisitorClass().Visitor(v).VisitSyntax(syntax)
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
	var expressionNames = v.expressionNames_.GetIterator()
	for expressionNames.HasNext() {
		var expressionName = expressionNames.GetNext()
		if !v.expressions_.ContainsValue(expressionName) {
			var message = fmt.Sprintf(
				"The following expression does not have an associated definition: %s",
				expressionName,
			)
			panic(message)
		}
	}
	var fragmentNames = v.fragmentNames_.GetIterator()
	for fragmentNames.HasNext() {
		var fragmentName = fragmentNames.GetNext()
		if !v.fragments_.ContainsValue(fragmentName) {
			var message = fmt.Sprintf(
				"The following fragment does not have an associated definition: %s",
				fragmentName,
			)
			panic(message)
		}
	}
	// We can "cheat" here because we know the code generator uses lists.
	var fragments = syntax.GetFragments().(fra.ListLike[ast.FragmentLike])
	fragments.SortValuesWithRanker(
		func(
			first ast.FragmentLike,
			second ast.FragmentLike,
		) fra.Rank {
			var firstName = first.GetAllcaps()
			var secondName = second.GetAllcaps()
			switch {
			case firstName < secondName:
				return fra.LesserRank
			case firstName > secondName:
				return fra.GreaterRank
			default:
				return fra.EqualRank
			}
		},
	)
}

// Methodical Methods

func (v *validator_) ProcessAllcaps(
	allcaps string,
) {
	v.validateToken(allcaps, AllcapsToken)
}

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
		v.expressionNames_.AddValue(actual)
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

func (v *validator_) PreprocessExpressionName(
	expressionName ast.ExpressionNameLike,
	index_ uint,
	count_ uint,
) {
	var lowercase = expressionName.GetLowercase()
	v.expressionNames_.AddValue(lowercase)
}

func (v *validator_) PreprocessFragment(
	fragment ast.FragmentLike,
	index_ uint,
	count_ uint,
) {
	var allcaps = fragment.GetAllcaps()
	v.fragments_.AddValue(allcaps)
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

func (v *validator_) PreprocessText(
	text ast.TextLike,
	index_ uint,
	count_ uint,
) {
	var scannerClass = ScannerClass()
	var actual = text.GetAny().(string)
	switch {
	case scannerClass.MatchesType(actual, LowercaseToken):
		v.expressionNames_.AddValue(actual)
	case scannerClass.MatchesType(actual, AllcapsToken):
		v.fragmentNames_.AddValue(actual)
	}
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
	ruleNames_       fra.SetLike[string]
	expressionNames_ fra.SetLike[string]
	fragmentNames_   fra.SetLike[string]
	rules_           fra.SetLike[string]
	expressions_     fra.SetLike[string]
	fragments_       fra.SetLike[string]

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
