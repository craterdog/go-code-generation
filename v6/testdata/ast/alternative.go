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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func AlternativeClass() AlternativeClassLike {
	return alternativeClass()
}

// Constructor Methods

func (c *alternativeClass_) Alternative(
	option OptionLike,
) AlternativeLike {
	if uti.IsUndefined(option) {
		panic("The \"option\" attribute is required by this class.")
	}
	var instance = &alternative_{
		// Initialize the instance attributes.
		option_: option,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *alternative_) GetClass() AlternativeClassLike {
	return alternativeClass()
}

// Attribute Methods

func (v *alternative_) GetOption() OptionLike {
	return v.option_
}

// PROTECTED INTERFACE

// Instance Structure

type alternative_ struct {
	// Declare the instance attributes.
	option_ OptionLike
}

// Class Structure

type alternativeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func alternativeClass() *alternativeClass_ {
	return alternativeClassReference_
}

var alternativeClassReference_ = &alternativeClass_{
	// Initialize the class constants.
}
