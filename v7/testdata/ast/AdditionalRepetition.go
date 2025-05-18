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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func AdditionalRepetitionClass() AdditionalRepetitionClassLike {
	return additionalRepetitionClass()
}

// Constructor Methods

func (c *additionalRepetitionClass_) AdditionalRepetition(
	repetition RepetitionLike,
) AdditionalRepetitionLike {
	if uti.IsUndefined(repetition) {
		panic("The \"repetition\" attribute is required by this class.")
	}
	var instance = &additionalRepetition_{
		// Initialize the instance attributes.
		repetition_: repetition,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalRepetition_) GetClass() AdditionalRepetitionClassLike {
	return additionalRepetitionClass()
}

// Attribute Methods

func (v *additionalRepetition_) GetRepetition() RepetitionLike {
	return v.repetition_
}

// PROTECTED INTERFACE

// Instance Structure

type additionalRepetition_ struct {
	// Declare the instance attributes.
	repetition_ RepetitionLike
}

// Class Structure

type additionalRepetitionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalRepetitionClass() *additionalRepetitionClass_ {
	return additionalRepetitionClassReference_
}

var additionalRepetitionClassReference_ = &additionalRepetitionClass_{
	// Initialize the class constants.
}
