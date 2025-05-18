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

func TermSequenceClass() TermSequenceClassLike {
	return termSequenceClass()
}

// Constructor Methods

func (c *termSequenceClass_) TermSequence(
	ruleTerms col.Sequential[RuleTermLike],
	optionalNote string,
) TermSequenceLike {
	if uti.IsUndefined(ruleTerms) {
		panic("The \"ruleTerms\" attribute is required by this class.")
	}
	var instance = &termSequence_{
		// Initialize the instance attributes.
		ruleTerms_:    ruleTerms,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *termSequence_) GetClass() TermSequenceClassLike {
	return termSequenceClass()
}

// Attribute Methods

func (v *termSequence_) GetRuleTerms() col.Sequential[RuleTermLike] {
	return v.ruleTerms_
}

func (v *termSequence_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type termSequence_ struct {
	// Declare the instance attributes.
	ruleTerms_    col.Sequential[RuleTermLike]
	optionalNote_ string
}

// Class Structure

type termSequenceClass_ struct {
	// Declare the class constants.
}

// Class Reference

func termSequenceClass() *termSequenceClass_ {
	return termSequenceClassReference_
}

var termSequenceClassReference_ = &termSequenceClass_{
	// Initialize the class constants.
}
