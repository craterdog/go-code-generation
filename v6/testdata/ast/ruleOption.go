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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func RuleOptionClass() RuleOptionClassLike {
	return ruleOptionClass()
}

// Constructor Methods

func (c *ruleOptionClass_) RuleOption(
	newline string,
	uppercase string,
	optionalNote string,
) RuleOptionLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	if uti.IsUndefined(uppercase) {
		panic("The \"uppercase\" attribute is required by this class.")
	}
	var instance = &ruleOption_{
		// Initialize the instance attributes.
		newline_:      newline,
		uppercase_:    uppercase,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *ruleOption_) GetClass() RuleOptionClassLike {
	return ruleOptionClass()
}

// Attribute Methods

func (v *ruleOption_) GetNewline() string {
	return v.newline_
}

func (v *ruleOption_) GetUppercase() string {
	return v.uppercase_
}

func (v *ruleOption_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type ruleOption_ struct {
	// Declare the instance attributes.
	newline_      string
	uppercase_    string
	optionalNote_ string
}

// Class Structure

type ruleOptionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func ruleOptionClass() *ruleOptionClass_ {
	return ruleOptionClassReference_
}

var ruleOptionClassReference_ = &ruleOptionClass_{
	// Initialize the class constants.
}
