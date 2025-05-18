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

func AlternativesClass() AlternativesClassLike {
	return alternativesClass()
}

// Constructor Methods

func (c *alternativesClass_) Alternatives(
	sequence SequenceLike,
	alternativeSequences col.Sequential[AlternativeSequenceLike],
) AlternativesLike {
	if uti.IsUndefined(sequence) {
		panic("The \"sequence\" attribute is required by this class.")
	}
	if uti.IsUndefined(alternativeSequences) {
		panic("The \"alternativeSequences\" attribute is required by this class.")
	}
	var instance = &alternatives_{
		// Initialize the instance attributes.
		sequence_:             sequence,
		alternativeSequences_: alternativeSequences,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *alternatives_) GetClass() AlternativesClassLike {
	return alternativesClass()
}

// Attribute Methods

func (v *alternatives_) GetSequence() SequenceLike {
	return v.sequence_
}

func (v *alternatives_) GetAlternativeSequences() col.Sequential[AlternativeSequenceLike] {
	return v.alternativeSequences_
}

// PROTECTED INTERFACE

// Instance Structure

type alternatives_ struct {
	// Declare the instance attributes.
	sequence_             SequenceLike
	alternativeSequences_ col.Sequential[AlternativeSequenceLike]
}

// Class Structure

type alternativesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func alternativesClass() *alternativesClass_ {
	return alternativesClassReference_
}

var alternativesClassReference_ = &alternativesClass_{
	// Initialize the class constants.
}
