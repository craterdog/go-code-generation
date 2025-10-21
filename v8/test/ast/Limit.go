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

func LimitClass() LimitClassLike {
	return limitClass()
}

// Constructor Methods

func (c *limitClass_) Limit(
	delimiter string,
	optionalNumber string,
) LimitLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	var instance = &limit_{
		// Initialize the instance attributes.
		delimiter_:      delimiter,
		optionalNumber_: optionalNumber,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *limit_) GetClass() LimitClassLike {
	return limitClass()
}

// Attribute Methods

func (v *limit_) GetDelimiter() string {
	return v.delimiter_
}

func (v *limit_) GetOptionalNumber() string {
	return v.optionalNumber_
}

// PROTECTED INTERFACE

// Instance Structure

type limit_ struct {
	// Declare the instance attributes.
	delimiter_      string
	optionalNumber_ string
}

// Class Structure

type limitClass_ struct {
	// Declare the class constants.
}

// Class Reference

func limitClass() *limitClass_ {
	return limitClassReference_
}

var limitClassReference_ = &limitClass_{
	// Initialize the class constants.
}
