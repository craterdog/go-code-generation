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

func Element() ElementClassLike {
	return elementReference()
}

// Constructor Methods

func (c *elementClass_) Make(
	any_ any,
) ElementLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &element_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *element_) GetClass() ElementClassLike {
	return elementReference()
}

// Attribute Methods

func (v *element_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type element_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type elementClass_ struct {
	// Declare the class constants.
}

// Class Reference

func elementReference() *elementClass_ {
	return elementReference_
}

var elementReference_ = &elementClass_{
	// Initialize the class constants.
}
