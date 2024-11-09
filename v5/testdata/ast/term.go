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

func Term() TermClassLike {
	return termReference()
}

// Constructor Methods

func (c *termClass_) Make(
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *term_) GetClass() TermClassLike {
	return termReference()
}

// Attribute Methods

func (v *term_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Private Methods

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

func termReference() *termClass_ {
	return termReference_
}

var termReference_ = &termClass_{
	// Initialize the class constants.
}
