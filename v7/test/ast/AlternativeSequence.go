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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func AlternativeSequenceClass() AlternativeSequenceClassLike {
	return alternativeSequenceClass()
}

// Constructor Methods

func (c *alternativeSequenceClass_) AlternativeSequence(
	delimiter string,
	sequence SequenceLike,
) AlternativeSequenceLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(sequence) {
		panic("The \"sequence\" attribute is required by this class.")
	}
	var instance = &alternativeSequence_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		sequence_:  sequence,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *alternativeSequence_) GetClass() AlternativeSequenceClassLike {
	return alternativeSequenceClass()
}

// Attribute Methods

func (v *alternativeSequence_) GetDelimiter() string {
	return v.delimiter_
}

func (v *alternativeSequence_) GetSequence() SequenceLike {
	return v.sequence_
}

// PROTECTED INTERFACE

// Instance Structure

type alternativeSequence_ struct {
	// Declare the instance attributes.
	delimiter_ string
	sequence_  SequenceLike
}

// Class Structure

type alternativeSequenceClass_ struct {
	// Declare the class constants.
}

// Class Reference

func alternativeSequenceClass() *alternativeSequenceClass_ {
	return alternativeSequenceClassReference_
}

var alternativeSequenceClassReference_ = &alternativeSequenceClass_{
	// Initialize the class constants.
}
