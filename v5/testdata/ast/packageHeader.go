/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
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
│   Updates to any section other than the Private Methods may be overwritten.  │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func PackageHeader() PackageHeaderClassLike {
	return packageHeaderReference()
}

// Constructor Methods

func (c *packageHeaderClass_) Make(
	comment string,
	name string,
) PackageHeaderLike {
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	if uti.IsUndefined(name) {
		panic("The \"name\" attribute is required by this class.")
	}
	var instance = &packageHeader_{
		// Initialize the instance attributes.
		comment_: comment,
		name_:    name,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *packageHeader_) GetClass() PackageHeaderClassLike {
	return packageHeaderReference()
}

// Attribute Methods

func (v *packageHeader_) GetComment() string {
	return v.comment_
}

func (v *packageHeader_) GetName() string {
	return v.name_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type packageHeader_ struct {
	// Declare the instance attributes.
	comment_ string
	name_    string
}

// Class Structure

type packageHeaderClass_ struct {
	// Declare the class constants.
}

// Class Reference

func packageHeaderReference() *packageHeaderClass_ {
	return packageHeaderReference_
}

var packageHeaderReference_ = &packageHeaderClass_{
	// Initialize the class constants.
}
