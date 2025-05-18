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
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func SequenceClass() SequenceClassLike {
	return sequenceClass()
}

// Constructor Methods

func (c *sequenceClass_) Sequence(
	repetition RepetitionLike,
	additionalRepetitions col.Sequential[AdditionalRepetitionLike],
) SequenceLike {
	if uti.IsUndefined(repetition) {
		panic("The \"repetition\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalRepetitions) {
		panic("The \"additionalRepetitions\" attribute is required by this class.")
	}
	var instance = &sequence_{
		// Initialize the instance attributes.
		repetition_:            repetition,
		additionalRepetitions_: additionalRepetitions,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *sequence_) GetClass() SequenceClassLike {
	return sequenceClass()
}

// Attribute Methods

func (v *sequence_) GetRepetition() RepetitionLike {
	return v.repetition_
}

func (v *sequence_) GetAdditionalRepetitions() col.Sequential[AdditionalRepetitionLike] {
	return v.additionalRepetitions_
}

// PROTECTED INTERFACE

// Instance Structure

type sequence_ struct {
	// Declare the instance attributes.
	repetition_            RepetitionLike
	additionalRepetitions_ col.Sequential[AdditionalRepetitionLike]
}

// Class Structure

type sequenceClass_ struct {
	// Declare the class constants.
}

// Class Reference

func sequenceClass() *sequenceClass_ {
	return sequenceClassReference_
}

var sequenceClassReference_ = &sequenceClass_{
	// Initialize the class constants.
}
