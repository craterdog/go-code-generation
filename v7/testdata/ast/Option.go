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

func OptionClass() OptionClassLike {
	return optionClass()
}

// Constructor Methods

func (c *optionClass_) Option(
	repetitions col.Sequential[RepetitionLike],
) OptionLike {
	if uti.IsUndefined(repetitions) {
		panic("The \"repetitions\" attribute is required by this class.")
	}
	var instance = &option_{
		// Initialize the instance attributes.
		repetitions_: repetitions,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *option_) GetClass() OptionClassLike {
	return optionClass()
}

// Attribute Methods

func (v *option_) GetRepetitions() col.Sequential[RepetitionLike] {
	return v.repetitions_
}

// PROTECTED INTERFACE

// Instance Structure

type option_ struct {
	// Declare the instance attributes.
	repetitions_ col.Sequential[RepetitionLike]
}

// Class Structure

type optionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func optionClass() *optionClass_ {
	return optionClassReference_
}

var optionClassReference_ = &optionClass_{
	// Initialize the class constants.
}
