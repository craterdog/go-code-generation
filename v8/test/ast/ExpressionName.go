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

func ExpressionNameClass() ExpressionNameClassLike {
	return expressionNameClass()
}

// Constructor Methods

func (c *expressionNameClass_) ExpressionName(
	newline string,
	lowercase string,
	optionalNote string,
) ExpressionNameLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	if uti.IsUndefined(lowercase) {
		panic("The \"lowercase\" attribute is required by this class.")
	}
	var instance = &expressionName_{
		// Initialize the instance attributes.
		newline_:      newline,
		lowercase_:    lowercase,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *expressionName_) GetClass() ExpressionNameClassLike {
	return expressionNameClass()
}

// Attribute Methods

func (v *expressionName_) GetNewline() string {
	return v.newline_
}

func (v *expressionName_) GetLowercase() string {
	return v.lowercase_
}

func (v *expressionName_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type expressionName_ struct {
	// Declare the instance attributes.
	newline_      string
	lowercase_    string
	optionalNote_ string
}

// Class Structure

type expressionNameClass_ struct {
	// Declare the class constants.
}

// Class Reference

func expressionNameClass() *expressionNameClass_ {
	return expressionNameClassReference_
}

var expressionNameClassReference_ = &expressionNameClass_{
	// Initialize the class constants.
}
