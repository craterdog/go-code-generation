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

func ImplicitClass() ImplicitClassLike {
	return implicitClass()
}

// Constructor Methods

func (c *implicitClass_) Implicit(
	intrinsic string,
) ImplicitLike {
	if uti.IsUndefined(intrinsic) {
		panic("The \"intrinsic\" attribute is required by this class.")
	}
	var instance = &implicit_{
		// Initialize the instance attributes.
		intrinsic_: intrinsic,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *implicit_) GetClass() ImplicitClassLike {
	return implicitClass()
}

// Attribute Methods

func (v *implicit_) GetIntrinsic() string {
	return v.intrinsic_
}

// PROTECTED INTERFACE

// Instance Structure

type implicit_ struct {
	// Declare the instance attributes.
	intrinsic_ string
}

// Class Structure

type implicitClass_ struct {
	// Declare the class constants.
}

// Class Reference

func implicitClass() *implicitClass_ {
	return implicitClassReference_
}

var implicitClassReference_ = &implicitClass_{
	// Initialize the class constants.
}
