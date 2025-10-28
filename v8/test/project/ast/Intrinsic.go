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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func IntrinsicClass() IntrinsicClassLike {
	return intrinsicClass()
}

// Constructor Methods

func (c *intrinsicClass_) Intrinsic(
	any_ any,
) IntrinsicLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &intrinsic_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *intrinsic_) GetClass() IntrinsicClassLike {
	return intrinsicClass()
}

// Attribute Methods

func (v *intrinsic_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type intrinsic_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type intrinsicClass_ struct {
	// Declare the class constants.
}

// Class Reference

func intrinsicClass() *intrinsicClass_ {
	return intrinsicClassReference_
}

var intrinsicClassReference_ = &intrinsicClass_{
	// Initialize the class constants.
}
