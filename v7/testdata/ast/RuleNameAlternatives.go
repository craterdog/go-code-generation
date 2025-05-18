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

func RuleNameAlternativesClass() RuleNameAlternativesClassLike {
	return ruleNameAlternativesClass()
}

// Constructor Methods

func (c *ruleNameAlternativesClass_) RuleNameAlternatives(
	ruleNames col.Sequential[RuleNameLike],
) RuleNameAlternativesLike {
	if uti.IsUndefined(ruleNames) {
		panic("The \"ruleNames\" attribute is required by this class.")
	}
	var instance = &ruleNameAlternatives_{
		// Initialize the instance attributes.
		ruleNames_: ruleNames,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *ruleNameAlternatives_) GetClass() RuleNameAlternativesClassLike {
	return ruleNameAlternativesClass()
}

// Attribute Methods

func (v *ruleNameAlternatives_) GetRuleNames() col.Sequential[RuleNameLike] {
	return v.ruleNames_
}

// PROTECTED INTERFACE

// Instance Structure

type ruleNameAlternatives_ struct {
	// Declare the instance attributes.
	ruleNames_ col.Sequential[RuleNameLike]
}

// Class Structure

type ruleNameAlternativesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func ruleNameAlternativesClass() *ruleNameAlternativesClass_ {
	return ruleNameAlternativesClassReference_
}

var ruleNameAlternativesClassReference_ = &ruleNameAlternativesClass_{
	// Initialize the class constants.
}
