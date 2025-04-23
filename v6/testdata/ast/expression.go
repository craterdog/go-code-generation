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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ExpressionClass() ExpressionClassLike {
	return expressionClass()
}

// Constructor Methods

func (c *expressionClass_) Expression(
	lowercase string,
	pattern PatternLike,
	optionalNote string,
) ExpressionLike {
	if uti.IsUndefined(lowercase) {
		panic("The \"lowercase\" attribute is required by this class.")
	}
	if uti.IsUndefined(pattern) {
		panic("The \"pattern\" attribute is required by this class.")
	}
	var instance = &expression_{
		// Initialize the instance attributes.
		lowercase_:    lowercase,
		pattern_:      pattern,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *expression_) GetClass() ExpressionClassLike {
	return expressionClass()
}

// Attribute Methods

func (v *expression_) GetLowercase() string {
	return v.lowercase_
}

func (v *expression_) GetPattern() PatternLike {
	return v.pattern_
}

func (v *expression_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type expression_ struct {
	// Declare the instance attributes.
	lowercase_    string
	pattern_      PatternLike
	optionalNote_ string
}

// Class Structure

type expressionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func expressionClass() *expressionClass_ {
	return expressionClassReference_
}

var expressionClassReference_ = &expressionClass_{
	// Initialize the class constants.
}
