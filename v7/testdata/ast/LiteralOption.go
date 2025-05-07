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

func LiteralOptionClass() LiteralOptionClassLike {
	return literalOptionClass()
}

// Constructor Methods

func (c *literalOptionClass_) LiteralOption(
	newline string,
	quote string,
	optionalNote string,
) LiteralOptionLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	if uti.IsUndefined(quote) {
		panic("The \"quote\" attribute is required by this class.")
	}
	var instance = &literalOption_{
		// Initialize the instance attributes.
		newline_:      newline,
		quote_:        quote,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *literalOption_) GetClass() LiteralOptionClassLike {
	return literalOptionClass()
}

// Attribute Methods

func (v *literalOption_) GetNewline() string {
	return v.newline_
}

func (v *literalOption_) GetQuote() string {
	return v.quote_
}

func (v *literalOption_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type literalOption_ struct {
	// Declare the instance attributes.
	newline_      string
	quote_        string
	optionalNote_ string
}

// Class Structure

type literalOptionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func literalOptionClass() *literalOptionClass_ {
	return literalOptionClassReference_
}

var literalOptionClassReference_ = &literalOptionClass_{
	// Initialize the class constants.
}
