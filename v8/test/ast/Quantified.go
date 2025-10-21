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

func QuantifiedClass() QuantifiedClassLike {
	return quantifiedClass()
}

// Constructor Methods

func (c *quantifiedClass_) Quantified(
	delimiter1 string,
	number string,
	optionalLimit LimitLike,
	delimiter2 string,
) QuantifiedLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(number) {
		panic("The \"number\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &quantified_{
		// Initialize the instance attributes.
		delimiter1_:    delimiter1,
		number_:        number,
		optionalLimit_: optionalLimit,
		delimiter2_:    delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *quantified_) GetClass() QuantifiedClassLike {
	return quantifiedClass()
}

// Attribute Methods

func (v *quantified_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *quantified_) GetNumber() string {
	return v.number_
}

func (v *quantified_) GetOptionalLimit() LimitLike {
	return v.optionalLimit_
}

func (v *quantified_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type quantified_ struct {
	// Declare the instance attributes.
	delimiter1_    string
	number_        string
	optionalLimit_ LimitLike
	delimiter2_    string
}

// Class Structure

type quantifiedClass_ struct {
	// Declare the class constants.
}

// Class Reference

func quantifiedClass() *quantifiedClass_ {
	return quantifiedClassReference_
}

var quantifiedClassReference_ = &quantifiedClass_{
	// Initialize the class constants.
}
