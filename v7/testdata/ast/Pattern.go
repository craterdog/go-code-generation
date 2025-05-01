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
	col "github.com/craterdog/go-collection-framework/v7/collection"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func PatternClass() PatternClassLike {
	return patternClass()
}

// Constructor Methods

func (c *patternClass_) Pattern(
	option OptionLike,
	alternatives col.Sequential[AlternativeLike],
) PatternLike {
	if uti.IsUndefined(option) {
		panic("The \"option\" attribute is required by this class.")
	}
	if uti.IsUndefined(alternatives) {
		panic("The \"alternatives\" attribute is required by this class.")
	}
	var instance = &pattern_{
		// Initialize the instance attributes.
		option_:       option,
		alternatives_: alternatives,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *pattern_) GetClass() PatternClassLike {
	return patternClass()
}

// Attribute Methods

func (v *pattern_) GetOption() OptionLike {
	return v.option_
}

func (v *pattern_) GetAlternatives() col.Sequential[AlternativeLike] {
	return v.alternatives_
}

// PROTECTED INTERFACE

// Instance Structure

type pattern_ struct {
	// Declare the instance attributes.
	option_       OptionLike
	alternatives_ col.Sequential[AlternativeLike]
}

// Class Structure

type patternClass_ struct {
	// Declare the class constants.
}

// Class Reference

func patternClass() *patternClass_ {
	return patternClassReference_
}

var patternClassReference_ = &patternClass_{
	// Initialize the class constants.
}
