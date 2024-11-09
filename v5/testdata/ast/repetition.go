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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Repetition() RepetitionClassLike {
	return repetitionReference()
}

// Constructor Methods

func (c *repetitionClass_) Make(
	element ElementLike,
	optionalCardinality CardinalityLike,
) RepetitionLike {
	if uti.IsUndefined(element) {
		panic("The \"element\" attribute is required by this class.")
	}
	var instance = &repetition_{
		// Initialize the instance attributes.
		element_:             element,
		optionalCardinality_: optionalCardinality,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *repetition_) GetClass() RepetitionClassLike {
	return repetitionReference()
}

// Attribute Methods

func (v *repetition_) GetElement() ElementLike {
	return v.element_
}

func (v *repetition_) GetOptionalCardinality() CardinalityLike {
	return v.optionalCardinality_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type repetition_ struct {
	// Declare the instance attributes.
	element_             ElementLike
	optionalCardinality_ CardinalityLike
}

// Class Structure

type repetitionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func repetitionReference() *repetitionClass_ {
	return repetitionReference_
}

var repetitionReference_ = &repetitionClass_{
	// Initialize the class constants.
}
