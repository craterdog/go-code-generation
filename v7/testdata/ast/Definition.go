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

func DefinitionClass() DefinitionClassLike {
	return definitionClass()
}

// Constructor Methods

func (c *definitionClass_) Definition(
	any_ any,
) DefinitionLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &definition_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *definition_) GetClass() DefinitionClassLike {
	return definitionClass()
}

// Attribute Methods

func (v *definition_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type definition_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type definitionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func definitionClass() *definitionClass_ {
	return definitionClassReference_
}

var definitionClassReference_ = &definitionClass_{
	// Initialize the class constants.
}
