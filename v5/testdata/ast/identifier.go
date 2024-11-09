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

func Identifier() IdentifierClassLike {
	return identifierReference()
}

// Constructor Methods

func (c *identifierClass_) Make(
	any_ any,
) IdentifierLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &identifier_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *identifier_) GetClass() IdentifierClassLike {
	return identifierReference()
}

// Attribute Methods

func (v *identifier_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type identifier_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type identifierClass_ struct {
	// Declare the class constants.
}

// Class Reference

func identifierReference() *identifierClass_ {
	return identifierReference_
}

var identifierReference_ = &identifierClass_{
	// Initialize the class constants.
}
