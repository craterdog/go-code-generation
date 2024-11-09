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

func Character() CharacterClassLike {
	return characterReference()
}

// Constructor Methods

func (c *characterClass_) Make(
	any_ any,
) CharacterLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &character_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *character_) GetClass() CharacterClassLike {
	return characterReference()
}

// Attribute Methods

func (v *character_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type character_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type characterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func characterReference() *characterClass_ {
	return characterReference_
}

var characterReference_ = &characterClass_{
	// Initialize the class constants.
}
