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

func RuleNameClass() RuleNameClassLike {
	return ruleNameClass()
}

// Constructor Methods

func (c *ruleNameClass_) RuleName(
	newline string,
	uppercase string,
	optionalNote string,
) RuleNameLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	if uti.IsUndefined(uppercase) {
		panic("The \"uppercase\" attribute is required by this class.")
	}
	var instance = &ruleName_{
		// Initialize the instance attributes.
		newline_:      newline,
		uppercase_:    uppercase,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *ruleName_) GetClass() RuleNameClassLike {
	return ruleNameClass()
}

// Attribute Methods

func (v *ruleName_) GetNewline() string {
	return v.newline_
}

func (v *ruleName_) GetUppercase() string {
	return v.uppercase_
}

func (v *ruleName_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type ruleName_ struct {
	// Declare the instance attributes.
	newline_      string
	uppercase_    string
	optionalNote_ string
}

// Class Structure

type ruleNameClass_ struct {
	// Declare the class constants.
}

// Class Reference

func ruleNameClass() *ruleNameClass_ {
	return ruleNameClassReference_
}

var ruleNameClassReference_ = &ruleNameClass_{
	// Initialize the class constants.
}
