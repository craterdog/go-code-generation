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

func TermClass() TermClassLike {
	return termClass()
}

// Constructor Methods

func (c *termClass_) Term(
	any_ any,
) TermLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &term_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *term_) GetClass() TermClassLike {
	return termClass()
}

// Attribute Methods

func (v *term_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type term_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type termClass_ struct {
	// Declare the class constants.
}

// Class Reference

func termClass() *termClass_ {
	return termClassReference_
}

var termClassReference_ = &termClass_{
	// Initialize the class constants.
}
