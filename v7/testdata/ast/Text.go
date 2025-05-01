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

func TextClass() TextClassLike {
	return textClass()
}

// Constructor Methods

func (c *textClass_) Text(
	any_ any,
) TextLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &text_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *text_) GetClass() TextClassLike {
	return textClass()
}

// Attribute Methods

func (v *text_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type text_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type textClass_ struct {
	// Declare the class constants.
}

// Class Reference

func textClass() *textClass_ {
	return textClassReference_
}

var textClassReference_ = &textClass_{
	// Initialize the class constants.
}
