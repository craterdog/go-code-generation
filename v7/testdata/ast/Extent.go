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

func ExtentClass() ExtentClassLike {
	return extentClass()
}

// Constructor Methods

func (c *extentClass_) Extent(
	glyph string,
) ExtentLike {
	if uti.IsUndefined(glyph) {
		panic("The \"glyph\" attribute is required by this class.")
	}
	var instance = &extent_{
		// Initialize the instance attributes.
		glyph_: glyph,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *extent_) GetClass() ExtentClassLike {
	return extentClass()
}

// Attribute Methods

func (v *extent_) GetGlyph() string {
	return v.glyph_
}

// PROTECTED INTERFACE

// Instance Structure

type extent_ struct {
	// Declare the instance attributes.
	glyph_ string
}

// Class Structure

type extentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func extentClass() *extentClass_ {
	return extentClassReference_
}

var extentClassReference_ = &extentClass_{
	// Initialize the class constants.
}
