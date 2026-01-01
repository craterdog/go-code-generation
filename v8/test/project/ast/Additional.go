/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func AdditionalClass() AdditionalClassLike {
	return additionalClass()
}

// Constructor Methods

func (c *additionalClass_) Additional(
	delimiter string,
	component ComponentLike,
) AdditionalLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &additional_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		component_: component,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *additional_) GetClass() AdditionalClassLike {
	return additionalClass()
}

// Attribute Methods

func (v *additional_) GetDelimiter() string {
	return v.delimiter_
}

func (v *additional_) GetComponent() ComponentLike {
	return v.component_
}

// PROTECTED INTERFACE

// Instance Structure

type additional_ struct {
	// Declare the instance attributes.
	delimiter_ string
	component_ ComponentLike
}

// Class Structure

type additionalClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalClass() *additionalClass_ {
	return additionalClassReference_
}

var additionalClassReference_ = &additionalClass_{
	// Initialize the class constants.
}
