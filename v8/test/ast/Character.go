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
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func CharacterClass() CharacterClassLike {
	return characterClass()
}

// Constructor Methods

func (c *characterClass_) Character(
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

// INSTANCE INTERFACE

// Principal Methods

func (v *character_) GetClass() CharacterClassLike {
	return characterClass()
}

// Attribute Methods

func (v *character_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

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

func characterClass() *characterClass_ {
	return characterClassReference_
}

var characterClassReference_ = &characterClass_{
	// Initialize the class constants.
}
