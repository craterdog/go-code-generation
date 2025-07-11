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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func RuleAlternativesClass() RuleAlternativesClassLike {
	return ruleAlternativesClass()
}

// Constructor Methods

func (c *ruleAlternativesClass_) RuleAlternatives(
	ruleNames fra.ListLike[RuleNameLike],
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

func (v *ruleAlternatives_) GetRuleNames() fra.ListLike[RuleNameLike] {
	return v.ruleNames_
}

// PROTECTED INTERFACE

// Instance Structure

type ruleAlternatives_ struct {
	// Declare the instance attributes.
	ruleNames_ fra.ListLike[RuleNameLike]
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
