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

func FilterClass() FilterClassLike {
	return filterClass()
}

// Constructor Methods

func (c *filterClass_) Filter(
	optionalExcluded string,
	characters col.Sequential[CharacterLike],
) FilterLike {
	if uti.IsUndefined(characters) {
		panic("The \"characters\" attribute is required by this class.")
	}
	var instance = &filter_{
		// Initialize the instance attributes.
		optionalExcluded_: optionalExcluded,
		characters_:       characters,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *filter_) GetClass() FilterClassLike {
	return filterClass()
}

// Attribute Methods

func (v *filter_) GetOptionalExcluded() string {
	return v.optionalExcluded_
}

func (v *filter_) GetCharacters() col.Sequential[CharacterLike] {
	return v.characters_
}

// PROTECTED INTERFACE

// Instance Structure

type filter_ struct {
	// Declare the instance attributes.
	optionalExcluded_ string
	characters_       col.Sequential[CharacterLike]
}

// Class Structure

type filterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func filterClass() *filterClass_ {
	return filterClassReference_
}

var filterClassReference_ = &filterClass_{
	// Initialize the class constants.
}
