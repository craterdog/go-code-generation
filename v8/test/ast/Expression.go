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

func ExpressionClass() ExpressionClassLike {
	return expressionClass()
}

// Constructor Methods

func (c *expressionClass_) Expression(
	delimiter1 string,
	lowercase string,
	delimiter2 string,
	pattern PatternLike,
) ExpressionLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(lowercase) {
		panic("The \"lowercase\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(pattern) {
		panic("The \"pattern\" attribute is required by this class.")
	}
	var instance = &expression_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		lowercase_:  lowercase,
		delimiter2_: delimiter2,
		pattern_:    pattern,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *expression_) GetClass() ExpressionClassLike {
	return expressionClass()
}

// Attribute Methods

func (v *expression_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *expression_) GetLowercase() string {
	return v.lowercase_
}

func (v *expression_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *expression_) GetPattern() PatternLike {
	return v.pattern_
}

// PROTECTED INTERFACE

// Instance Structure

type expression_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	lowercase_  string
	delimiter2_ string
	pattern_    PatternLike
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
