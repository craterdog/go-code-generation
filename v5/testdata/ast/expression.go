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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Expression() ExpressionClassLike {
	return expressionReference()
}

// Constructor Methods

func (c *expressionClass_) Make(
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *expression_) GetClass() ExpressionClassLike {
	return expressionReference()
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

// Private Methods

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

func expressionReference() *expressionClass_ {
	return expressionReference_
}

var expressionReference_ = &expressionClass_{
	// Initialize the class constants.
}
