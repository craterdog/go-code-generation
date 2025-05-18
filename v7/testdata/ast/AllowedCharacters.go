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
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func AllowedCharactersClass() AllowedCharactersClassLike {
	return allowedCharactersClass()
}

// Constructor Methods

func (c *allowedCharactersClass_) AllowedCharacters(
	character CharacterLike,
	additionalCharacters col.Sequential[AdditionalCharacterLike],
) AllowedCharactersLike {
	if uti.IsUndefined(character) {
		panic("The \"character\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalCharacters) {
		panic("The \"additionalCharacters\" attribute is required by this class.")
	}
	var instance = &allowedCharacters_{
		// Initialize the instance attributes.
		character_:            character,
		additionalCharacters_: additionalCharacters,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *allowedCharacters_) GetClass() AllowedCharactersClassLike {
	return allowedCharactersClass()
}

// Attribute Methods

func (v *allowedCharacters_) GetCharacter() CharacterLike {
	return v.character_
}

func (v *allowedCharacters_) GetAdditionalCharacters() col.Sequential[AdditionalCharacterLike] {
	return v.additionalCharacters_
}

// PROTECTED INTERFACE

// Instance Structure

type allowedCharacters_ struct {
	// Declare the instance attributes.
	character_            CharacterLike
	additionalCharacters_ col.Sequential[AdditionalCharacterLike]
}

// Class Structure

type allowedCharactersClass_ struct {
	// Declare the class constants.
}

// Class Reference

func allowedCharactersClass() *allowedCharactersClass_ {
	return allowedCharactersClassReference_
}

var allowedCharactersClassReference_ = &allowedCharactersClass_{
	// Initialize the class constants.
}
