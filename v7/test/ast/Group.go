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

func GroupClass() GroupClassLike {
	return groupClass()
}

// Constructor Methods

func (c *groupClass_) Group(
	delimiter1 string,
	alternatives AlternativesLike,
	delimiter2 string,
) GroupLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(alternatives) {
		panic("The \"alternatives\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &group_{
		// Initialize the instance attributes.
		delimiter1_:   delimiter1,
		alternatives_: alternatives,
		delimiter2_:   delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *group_) GetClass() GroupClassLike {
	return groupClass()
}

// Attribute Methods

func (v *group_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *group_) GetAlternatives() AlternativesLike {
	return v.alternatives_
}

func (v *group_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type group_ struct {
	// Declare the instance attributes.
	delimiter1_   string
	alternatives_ AlternativesLike
	delimiter2_   string
}

// Class Structure

type groupClass_ struct {
	// Declare the class constants.
}

// Class Reference

func groupClass() *groupClass_ {
	return groupClassReference_
}

var groupClassReference_ = &groupClass_{
	// Initialize the class constants.
}
