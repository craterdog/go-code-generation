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

func Option() OptionClassLike {
	return optionReference()
}

// Constructor Methods

func (c *optionClass_) Make(
	repetitions abs.Sequential[RepetitionLike],
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *option_) GetClass() OptionClassLike {
	return optionReference()
}

// Attribute Methods

func (v *option_) GetRepetitions() abs.Sequential[RepetitionLike] {
	return v.repetitions_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type option_ struct {
	// Declare the instance attributes.
	repetitions_ abs.Sequential[RepetitionLike]
}

// Class Structure

type optionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func optionReference() *optionClass_ {
	return optionReference_
}

var optionReference_ = &optionClass_{
	// Initialize the class constants.
}
