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

func RuleClass() RuleClassLike {
	return ruleClass()
}

// Constructor Methods

func (c *ruleClass_) Rule(
	delimiter1 string,
	uppercase string,
	delimiter2 string,
	definition DefinitionLike,
) RuleLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(uppercase) {
		panic("The \"uppercase\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(definition) {
		panic("The \"definition\" attribute is required by this class.")
	}
	var instance = &rule_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		uppercase_:  uppercase,
		delimiter2_: delimiter2,
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

func (v *rule_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *rule_) GetUppercase() string {
	return v.uppercase_
}

func (v *rule_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *rule_) GetDefinition() DefinitionLike {
	return v.definition_
}

// PROTECTED INTERFACE

// Instance Structure

type rule_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	uppercase_  string
	delimiter2_ string
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
