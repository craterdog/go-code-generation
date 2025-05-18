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

func TokenNameClass() TokenNameClassLike {
	return tokenNameClass()
}

// Constructor Methods

func (c *tokenNameClass_) TokenName(
	newline string,
	lowercase string,
	optionalNote string,
) TokenNameLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	if uti.IsUndefined(lowercase) {
		panic("The \"lowercase\" attribute is required by this class.")
	}
	var instance = &tokenName_{
		// Initialize the instance attributes.
		newline_:      newline,
		lowercase_:    lowercase,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *tokenName_) GetClass() TokenNameClassLike {
	return tokenNameClass()
}

// Attribute Methods

func (v *tokenName_) GetNewline() string {
	return v.newline_
}

func (v *tokenName_) GetLowercase() string {
	return v.lowercase_
}

func (v *tokenName_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type tokenName_ struct {
	// Declare the instance attributes.
	newline_      string
	lowercase_    string
	optionalNote_ string
}

// Class Structure

type tokenNameClass_ struct {
	// Declare the class constants.
}

// Class Reference

func tokenNameClass() *tokenNameClass_ {
	return tokenNameClassReference_
}

var tokenNameClassReference_ = &tokenNameClass_{
	// Initialize the class constants.
}
