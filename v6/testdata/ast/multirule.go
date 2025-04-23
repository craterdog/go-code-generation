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
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func MultiruleClass() MultiruleClassLike {
	return multiruleClass()
}

// Constructor Methods

func (c *multiruleClass_) Multirule(
	ruleOptions col.Sequential[RuleOptionLike],
) MultiruleLike {
	if uti.IsUndefined(ruleOptions) {
		panic("The \"ruleOptions\" attribute is required by this class.")
	}
	var instance = &multirule_{
		// Initialize the instance attributes.
		ruleOptions_: ruleOptions,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multirule_) GetClass() MultiruleClassLike {
	return multiruleClass()
}

// Attribute Methods

func (v *multirule_) GetRuleOptions() col.Sequential[RuleOptionLike] {
	return v.ruleOptions_
}

// PROTECTED INTERFACE

// Instance Structure

type multirule_ struct {
	// Declare the instance attributes.
	ruleOptions_ col.Sequential[RuleOptionLike]
}

// Class Structure

type multiruleClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multiruleClass() *multiruleClass_ {
	return multiruleClassReference_
}

var multiruleClassReference_ = &multiruleClass_{
	// Initialize the class constants.
}
