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

func Explicit() ExplicitClassLike {
	return explicitReference()
}

// Constructor Methods

func (c *explicitClass_) Make(
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *explicit_) GetClass() ExplicitClassLike {
	return explicitReference()
}

// Attribute Methods

func (v *explicit_) GetGlyph() string {
	return v.glyph_
}

func (v *explicit_) GetOptionalExtent() ExtentLike {
	return v.optionalExtent_
}

// PROTECTED INTERFACE

// Private Methods

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

func explicitReference() *explicitClass_ {
	return explicitReference_
}

var explicitReference_ = &explicitClass_{
	// Initialize the class constants.
}
