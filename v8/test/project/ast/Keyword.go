/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func KeywordClass() KeywordClassLike {
	return keywordClass()
}

// Constructor Methods

func (c *keywordClass_) Keyword(
	any_ any,
) KeywordLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &keyword_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *keyword_) GetClass() KeywordClassLike {
	return keywordClass()
}

// Attribute Methods

func (v *keyword_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type keyword_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type keywordClass_ struct {
	// Declare the class constants.
}

// Class Reference

func keywordClass() *keywordClass_ {
	return keywordClassReference_
}

var keywordClassReference_ = &keywordClass_{
	// Initialize the class constants.
}
