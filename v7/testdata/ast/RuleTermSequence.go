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

func RuleTermSequenceClass() RuleTermSequenceClassLike {
	return ruleTermSequenceClass()
}

// Constructor Methods

func (c *ruleTermSequenceClass_) RuleTermSequence(
	ruleTerms col.Sequential[RuleTermLike],
	optionalNote string,
) RuleTermSequenceLike {
	if uti.IsUndefined(ruleTerms) {
		panic("The \"ruleTerms\" attribute is required by this class.")
	}
	var instance = &ruleTermSequence_{
		// Initialize the instance attributes.
		ruleTerms_:    ruleTerms,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *ruleTermSequence_) GetClass() RuleTermSequenceClassLike {
	return ruleTermSequenceClass()
}

// Attribute Methods

func (v *ruleTermSequence_) GetRuleTerms() col.Sequential[RuleTermLike] {
	return v.ruleTerms_
}

func (v *ruleTermSequence_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type ruleTermSequence_ struct {
	// Declare the instance attributes.
	ruleTerms_    col.Sequential[RuleTermLike]
	optionalNote_ string
}

// Class Structure

type ruleTermSequenceClass_ struct {
	// Declare the class constants.
}

// Class Reference

func ruleTermSequenceClass() *ruleTermSequenceClass_ {
	return ruleTermSequenceClassReference_
}

var ruleTermSequenceClassReference_ = &ruleTermSequenceClass_{
	// Initialize the class constants.
}
