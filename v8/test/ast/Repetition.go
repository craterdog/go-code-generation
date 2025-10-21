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
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func RepetitionClass() RepetitionClassLike {
	return repetitionClass()
}

// Constructor Methods

func (c *repetitionClass_) Repetition(
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

// INSTANCE INTERFACE

// Principal Methods

func (v *repetition_) GetClass() RepetitionClassLike {
	return repetitionClass()
}

// Attribute Methods

func (v *repetition_) GetElement() ElementLike {
	return v.element_
}

func (v *repetition_) GetOptionalCardinality() CardinalityLike {
	return v.optionalCardinality_
}

// PROTECTED INTERFACE

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

func repetitionClass() *repetitionClass_ {
	return repetitionClassReference_
}

var repetitionClassReference_ = &repetitionClass_{
	// Initialize the class constants.
}
