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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ExpressionOptionClass() ExpressionOptionClassLike {
	return expressionOptionClass()
}

// Constructor Methods

func (c *expressionOptionClass_) ExpressionOption(
	newline string,
	lowercase string,
	optionalNote string,
) ExpressionOptionLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	if uti.IsUndefined(lowercase) {
		panic("The \"lowercase\" attribute is required by this class.")
	}
	var instance = &expressionOption_{
		// Initialize the instance attributes.
		newline_:      newline,
		lowercase_:    lowercase,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *expressionOption_) GetClass() ExpressionOptionClassLike {
	return expressionOptionClass()
}

// Attribute Methods

func (v *expressionOption_) GetNewline() string {
	return v.newline_
}

func (v *expressionOption_) GetLowercase() string {
	return v.lowercase_
}

func (v *expressionOption_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type expressionOption_ struct {
	// Declare the instance attributes.
	newline_      string
	lowercase_    string
	optionalNote_ string
}

// Class Structure

type expressionOptionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func expressionOptionClass() *expressionOptionClass_ {
	return expressionOptionClassReference_
}

var expressionOptionClassReference_ = &expressionOptionClass_{
	// Initialize the class constants.
}
