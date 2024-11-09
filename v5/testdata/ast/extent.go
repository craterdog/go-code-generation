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

func Extent() ExtentClassLike {
	return extentReference()
}

// Constructor Methods

func (c *extentClass_) Make(
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *extent_) GetClass() ExtentClassLike {
	return extentReference()
}

// Attribute Methods

func (v *extent_) GetGlyph() string {
	return v.glyph_
}

// PROTECTED INTERFACE

// Private Methods

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

func extentReference() *extentClass_ {
	return extentReference_
}

var extentReference_ = &extentClass_{
	// Initialize the class constants.
}
