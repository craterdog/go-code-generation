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
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func LiteralValueClass() LiteralValueClassLike {
	return literalValueClass()
}

// Constructor Methods

func (c *literalValueClass_) LiteralValue(
	newline string,
	literal string,
	optionalNote string,
) LiteralValueLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	if uti.IsUndefined(literal) {
		panic("The \"literal\" attribute is required by this class.")
	}
	var instance = &literalValue_{
		// Initialize the instance attributes.
		newline_:      newline,
		literal_:      literal,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *literalValue_) GetClass() LiteralValueClassLike {
	return literalValueClass()
}

// Attribute Methods

func (v *literalValue_) GetNewline() string {
	return v.newline_
}

func (v *literalValue_) GetLiteral() string {
	return v.literal_
}

func (v *literalValue_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type literalValue_ struct {
	// Declare the instance attributes.
	newline_      string
	literal_      string
	optionalNote_ string
}

// Class Structure

type literalValueClass_ struct {
	// Declare the class constants.
}

// Class Reference

func literalValueClass() *literalValueClass_ {
	return literalValueClassReference_
}

var literalValueClassReference_ = &literalValueClass_{
	// Initialize the class constants.
}
