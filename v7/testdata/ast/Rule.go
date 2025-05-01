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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func RuleClass() RuleClassLike {
	return ruleClass()
}

// Constructor Methods

func (c *ruleClass_) Rule(
	uppercase string,
	definition DefinitionLike,
) RuleLike {
	if uti.IsUndefined(uppercase) {
		panic("The \"uppercase\" attribute is required by this class.")
	}
	if uti.IsUndefined(definition) {
		panic("The \"definition\" attribute is required by this class.")
	}
	var instance = &rule_{
		// Initialize the instance attributes.
		uppercase_:  uppercase,
		definition_: definition,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *rule_) GetClass() RuleClassLike {
	return ruleClass()
}

// Attribute Methods

func (v *rule_) GetUppercase() string {
	return v.uppercase_
}

func (v *rule_) GetDefinition() DefinitionLike {
	return v.definition_
}

// PROTECTED INTERFACE

// Instance Structure

type rule_ struct {
	// Declare the instance attributes.
	uppercase_  string
	definition_ DefinitionLike
}

// Class Structure

type ruleClass_ struct {
	// Declare the class constants.
}

// Class Reference

func ruleClass() *ruleClass_ {
	return ruleClassReference_
}

var ruleClassReference_ = &ruleClass_{
	// Initialize the class constants.
}
