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

func Filter() FilterClassLike {
	return filterReference()
}

// Constructor Methods

func (c *filterClass_) Make(
	optionalExcluded string,
	characters abs.Sequential[CharacterLike],
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *filter_) GetClass() FilterClassLike {
	return filterReference()
}

// Attribute Methods

func (v *filter_) GetOptionalExcluded() string {
	return v.optionalExcluded_
}

func (v *filter_) GetCharacters() abs.Sequential[CharacterLike] {
	return v.characters_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type filter_ struct {
	// Declare the instance attributes.
	optionalExcluded_ string
	characters_       abs.Sequential[CharacterLike]
}

// Class Structure

type filterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func filterReference() *filterClass_ {
	return filterReference_
}

var filterReference_ = &filterClass_{
	// Initialize the class constants.
}
