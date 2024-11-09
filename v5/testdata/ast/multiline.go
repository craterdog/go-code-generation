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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Multiline() MultilineClassLike {
	return multilineReference()
}

// Constructor Methods

func (c *multilineClass_) Make(
	lines abs.Sequential[LineLike],
) MultilineLike {
	if uti.IsUndefined(lines) {
		panic("The \"lines\" attribute is required by this class.")
	}
	var instance = &multiline_{
		// Initialize the instance attributes.
		lines_: lines,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *multiline_) GetClass() MultilineClassLike {
	return multilineReference()
}

// Attribute Methods

func (v *multiline_) GetLines() abs.Sequential[LineLike] {
	return v.lines_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type multiline_ struct {
	// Declare the instance attributes.
	lines_ abs.Sequential[LineLike]
}

// Class Structure

type multilineClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multilineReference() *multilineClass_ {
	return multilineReference_
}

var multilineReference_ = &multilineClass_{
	// Initialize the class constants.
}
