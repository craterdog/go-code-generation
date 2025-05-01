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

func ConstrainedClass() ConstrainedClassLike {
	return constrainedClass()
}

// Constructor Methods

func (c *constrainedClass_) Constrained(
	any_ any,
) ConstrainedLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &constrained_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *constrained_) GetClass() ConstrainedClassLike {
	return constrainedClass()
}

// Attribute Methods

func (v *constrained_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type constrained_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type constrainedClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constrainedClass() *constrainedClass_ {
	return constrainedClassReference_
}

var constrainedClassReference_ = &constrainedClass_{
	// Initialize the class constants.
}
