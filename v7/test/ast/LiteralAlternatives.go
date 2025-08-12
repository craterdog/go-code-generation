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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func LiteralAlternativesClass() LiteralAlternativesClassLike {
	return literalAlternativesClass()
}

// Constructor Methods

func (c *literalAlternativesClass_) LiteralAlternatives(
	literalValues fra.Sequential[LiteralValueLike],
) LiteralAlternativesLike {
	if uti.IsUndefined(literalValues) {
		panic("The \"literalValues\" attribute is required by this class.")
	}
	var instance = &literalAlternatives_{
		// Initialize the instance attributes.
		literalValues_: literalValues,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *literalAlternatives_) GetClass() LiteralAlternativesClassLike {
	return literalAlternativesClass()
}

// Attribute Methods

func (v *literalAlternatives_) GetLiteralValues() fra.Sequential[LiteralValueLike] {
	return v.literalValues_
}

// PROTECTED INTERFACE

// Instance Structure

type literalAlternatives_ struct {
	// Declare the instance attributes.
	literalValues_ fra.Sequential[LiteralValueLike]
}

// Class Structure

type literalAlternativesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func literalAlternativesClass() *literalAlternativesClass_ {
	return literalAlternativesClassReference_
}

var literalAlternativesClassReference_ = &literalAlternativesClass_{
	// Initialize the class constants.
}
