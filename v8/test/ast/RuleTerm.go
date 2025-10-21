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
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func RuleTermClass() RuleTermClassLike {
	return ruleTermClass()
}

// Constructor Methods

func (c *ruleTermClass_) RuleTerm(
	component ComponentLike,
	optionalCardinality CardinalityLike,
) RuleTermLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &ruleTerm_{
		// Initialize the instance attributes.
		component_:           component,
		optionalCardinality_: optionalCardinality,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *ruleTerm_) GetClass() RuleTermClassLike {
	return ruleTermClass()
}

// Attribute Methods

func (v *ruleTerm_) GetComponent() ComponentLike {
	return v.component_
}

func (v *ruleTerm_) GetOptionalCardinality() CardinalityLike {
	return v.optionalCardinality_
}

// PROTECTED INTERFACE

// Instance Structure

type ruleTerm_ struct {
	// Declare the instance attributes.
	component_           ComponentLike
	optionalCardinality_ CardinalityLike
}

// Class Structure

type ruleTermClass_ struct {
	// Declare the class constants.
}

// Class Reference

func ruleTermClass() *ruleTermClass_ {
	return ruleTermClassReference_
}

var ruleTermClassReference_ = &ruleTermClass_{
	// Initialize the class constants.
}
