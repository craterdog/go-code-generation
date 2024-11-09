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

func Pattern() PatternClassLike {
	return patternReference()
}

// Constructor Methods

func (c *patternClass_) Make(
	option OptionLike,
	alternatives abs.Sequential[AlternativeLike],
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *pattern_) GetClass() PatternClassLike {
	return patternReference()
}

// Attribute Methods

func (v *pattern_) GetOption() OptionLike {
	return v.option_
}

func (v *pattern_) GetAlternatives() abs.Sequential[AlternativeLike] {
	return v.alternatives_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type pattern_ struct {
	// Declare the instance attributes.
	option_       OptionLike
	alternatives_ abs.Sequential[AlternativeLike]
}

// Class Structure

type patternClass_ struct {
	// Declare the class constants.
}

// Class Reference

func patternReference() *patternClass_ {
	return patternReference_
}

var patternReference_ = &patternClass_{
	// Initialize the class constants.
}
