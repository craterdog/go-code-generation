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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Group() GroupClassLike {
	return groupReference()
}

// Constructor Methods

func (c *groupClass_) Make(
	pattern PatternLike,
) GroupLike {
	if uti.IsUndefined(pattern) {
		panic("The \"pattern\" attribute is required by this class.")
	}
	var instance = &group_{
		// Initialize the instance attributes.
		pattern_: pattern,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *group_) GetClass() GroupClassLike {
	return groupReference()
}

// Attribute Methods

func (v *group_) GetPattern() PatternLike {
	return v.pattern_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type group_ struct {
	// Declare the instance attributes.
	pattern_ PatternLike
}

// Class Structure

type groupClass_ struct {
	// Declare the class constants.
}

// Class Reference

func groupReference() *groupClass_ {
	return groupReference_
}

var groupReference_ = &groupClass_{
	// Initialize the class constants.
}
