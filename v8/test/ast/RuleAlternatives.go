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
	fra "github.com/craterdog/go-collection-framework/v8"
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func RuleAlternativesClass() RuleAlternativesClassLike {
	return ruleAlternativesClass()
}

// Constructor Methods

func (c *ruleAlternativesClass_) RuleAlternatives(
	ruleNames fra.Sequential[RuleNameLike],
) RuleAlternativesLike {
	if uti.IsUndefined(ruleNames) {
		panic("The \"ruleNames\" attribute is required by this class.")
	}
	var instance = &ruleAlternatives_{
		// Initialize the instance attributes.
		ruleNames_: ruleNames,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *ruleAlternatives_) GetClass() RuleAlternativesClassLike {
	return ruleAlternativesClass()
}

// Attribute Methods

func (v *ruleAlternatives_) GetRuleNames() fra.Sequential[RuleNameLike] {
	return v.ruleNames_
}

// PROTECTED INTERFACE

// Instance Structure

type ruleAlternatives_ struct {
	// Declare the instance attributes.
	ruleNames_ fra.Sequential[RuleNameLike]
}

// Class Structure

type ruleAlternativesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func ruleAlternativesClass() *ruleAlternativesClass_ {
	return ruleAlternativesClassReference_
}

var ruleAlternativesClassReference_ = &ruleAlternativesClass_{
	// Initialize the class constants.
}
