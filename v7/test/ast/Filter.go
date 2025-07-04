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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func FilterClass() FilterClassLike {
	return filterClass()
}

// Constructor Methods

func (c *filterClass_) Filter(
	optionalDelimiter string,
	delimiter1 string,
	characters fra.ListLike[CharacterLike],
	delimiter2 string,
) FilterLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(characters) {
		panic("The \"characters\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &filter_{
		// Initialize the instance attributes.
		optionalDelimiter_: optionalDelimiter,
		delimiter1_:        delimiter1,
		characters_:        characters,
		delimiter2_:        delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *filter_) GetClass() FilterClassLike {
	return filterClass()
}

// Attribute Methods

func (v *filter_) GetOptionalDelimiter() string {
	return v.optionalDelimiter_
}

func (v *filter_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *filter_) GetCharacters() fra.ListLike[CharacterLike] {
	return v.characters_
}

func (v *filter_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type filter_ struct {
	// Declare the instance attributes.
	optionalDelimiter_ string
	delimiter1_        string
	characters_        fra.ListLike[CharacterLike]
	delimiter2_        string
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
