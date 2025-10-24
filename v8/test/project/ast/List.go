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
	fra "github.com/craterdog/go-collection-framework/v8"
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func ListClass() ListClassLike {
	return listClass()
}

// Constructor Methods

func (c *listClass_) List(
	delimiter1 string,
	component ComponentLike,
	additionals fra.Sequential[AdditionalLike],
	delimiter2 string,
) ListLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionals) {
		panic("The \"additionals\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &list_{
		// Initialize the instance attributes.
		delimiter1_:  delimiter1,
		component_:   component,
		additionals_: additionals,
		delimiter2_:  delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *list_) GetClass() ListClassLike {
	return listClass()
}

// Attribute Methods

func (v *list_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *list_) GetComponent() ComponentLike {
	return v.component_
}

func (v *list_) GetAdditionals() fra.Sequential[AdditionalLike] {
	return v.additionals_
}

func (v *list_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type list_ struct {
	// Declare the instance attributes.
	delimiter1_  string
	component_   ComponentLike
	additionals_ fra.Sequential[AdditionalLike]
	delimiter2_  string
}

// Class Structure

type listClass_ struct {
	// Declare the class constants.
}

// Class Reference

func listClass() *listClass_ {
	return listClassReference_
}

var listClassReference_ = &listClass_{
	// Initialize the class constants.
}
