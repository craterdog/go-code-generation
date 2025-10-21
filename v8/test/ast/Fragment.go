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

func FragmentClass() FragmentClassLike {
	return fragmentClass()
}

// Constructor Methods

func (c *fragmentClass_) Fragment(
	delimiter1 string,
	allcaps string,
	delimiter2 string,
	pattern PatternLike,
) FragmentLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(allcaps) {
		panic("The \"allcaps\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(pattern) {
		panic("The \"pattern\" attribute is required by this class.")
	}
	var instance = &fragment_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		allcaps_:    allcaps,
		delimiter2_: delimiter2,
		pattern_:    pattern,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *fragment_) GetClass() FragmentClassLike {
	return fragmentClass()
}

// Attribute Methods

func (v *fragment_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *fragment_) GetAllcaps() string {
	return v.allcaps_
}

func (v *fragment_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *fragment_) GetPattern() PatternLike {
	return v.pattern_
}

// PROTECTED INTERFACE

// Instance Structure

type fragment_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	allcaps_    string
	delimiter2_ string
	pattern_    PatternLike
}

// Class Structure

type fragmentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func fragmentClass() *fragmentClass_ {
	return fragmentClassReference_
}

var fragmentClassReference_ = &fragmentClass_{
	// Initialize the class constants.
}
