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

func ExplicitClass() ExplicitClassLike {
	return explicitClass()
}

// Constructor Methods

func (c *explicitClass_) Explicit(
	glyph string,
	optionalExtent ExtentLike,
) ExplicitLike {
	if uti.IsUndefined(glyph) {
		panic("The \"glyph\" attribute is required by this class.")
	}
	var instance = &explicit_{
		// Initialize the instance attributes.
		glyph_:          glyph,
		optionalExtent_: optionalExtent,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *explicit_) GetClass() ExplicitClassLike {
	return explicitClass()
}

// Attribute Methods

func (v *explicit_) GetGlyph() string {
	return v.glyph_
}

func (v *explicit_) GetOptionalExtent() ExtentLike {
	return v.optionalExtent_
}

// PROTECTED INTERFACE

// Instance Structure

type explicit_ struct {
	// Declare the instance attributes.
	glyph_          string
	optionalExtent_ ExtentLike
}

// Class Structure

type explicitClass_ struct {
	// Declare the class constants.
}

// Class Reference

func explicitClass() *explicitClass_ {
	return explicitClassReference_
}

var explicitClassReference_ = &explicitClass_{
	// Initialize the class constants.
}
