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
	ast "github.com/craterdog/go-example-syntax/ast"
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
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorClass()
}

func (v *validator_) ValidateDocument(
	document ast.DocumentLike,
) {
	VisitorClass().Visitor(v).VisitDocument(document)
}

// Methodical Methods

func (v *validator_) ProcessNewline(
	newline string,
) {
	v.validateExpression(newline, NewlineExpression)
}

func (v *validator_) ProcessNumber(
	number string,
) {
	v.validateExpression(number, NumberExpression)
}

func (v *validator_) ProcessRune(
	rune_ string,
) {
	v.validateExpression(rune_, RuneExpression)
}

func (v *validator_) ProcessSpace(
	space string,
) {
	v.validateExpression(space, SpaceExpression)
}

func (v *validator_) ProcessText(
	text string,
) {
	v.validateExpression(text, TextExpression)
}

func (v *validator_) PreprocessAdditional(
	additional ast.AdditionalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessComponent(
	component ast.ComponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDocument(
	document ast.DocumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIntrinsic(
	intrinsic ast.IntrinsicLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessKeyword(
	keyword ast.KeywordLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessList(
	list ast.ListLike,
	index_ uint,
	count_ uint,
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
