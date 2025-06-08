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
	com "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func SequenceClass() SequenceClassLike {
	return sequenceClass()
}

// Constructor Methods

func (c *sequenceClass_) Sequence(
	repetitions com.ListLike[RepetitionLike],
) SequenceLike {
	if uti.IsUndefined(repetitions) {
		panic("The \"repetitions\" attribute is required by this class.")
	}
	var instance = &sequence_{
		// Initialize the instance attributes.
		repetitions_: repetitions,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *sequence_) GetClass() SequenceClassLike {
	return sequenceClass()
}

// Attribute Methods

func (v *sequence_) GetRepetitions() com.ListLike[RepetitionLike] {
	return v.repetitions_
}

// PROTECTED INTERFACE

// Instance Structure

type sequence_ struct {
	// Declare the instance attributes.
	repetitions_ com.ListLike[RepetitionLike]
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
