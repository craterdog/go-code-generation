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

func AdditionalCharacterClass() AdditionalCharacterClassLike {
	return additionalCharacterClass()
}

// Constructor Methods

func (c *additionalCharacterClass_) AdditionalCharacter(
	character CharacterLike,
) AdditionalCharacterLike {
	if uti.IsUndefined(character) {
		panic("The \"character\" attribute is required by this class.")
	}
	var instance = &additionalCharacter_{
		// Initialize the instance attributes.
		character_: character,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalCharacter_) GetClass() AdditionalCharacterClassLike {
	return additionalCharacterClass()
}

// Attribute Methods

func (v *additionalCharacter_) GetCharacter() CharacterLike {
	return v.character_
}

// PROTECTED INTERFACE

// Instance Structure

type additionalCharacter_ struct {
	// Declare the instance attributes.
	character_ CharacterLike
}

// Class Structure

type additionalCharacterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalCharacterClass() *additionalCharacterClass_ {
	return additionalCharacterClassReference_
}

var additionalCharacterClassReference_ = &additionalCharacterClass_{
	// Initialize the class constants.
}
