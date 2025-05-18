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
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func LiteralValueAlternativesClass() LiteralValueAlternativesClassLike {
	return literalValueAlternativesClass()
}

// Constructor Methods

func (c *literalValueAlternativesClass_) LiteralValueAlternatives(
	literalValues col.Sequential[LiteralValueLike],
) LiteralValueAlternativesLike {
	if uti.IsUndefined(literalValues) {
		panic("The \"literalValues\" attribute is required by this class.")
	}
	var instance = &literalValueAlternatives_{
		// Initialize the instance attributes.
		literalValues_: literalValues,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *literalValueAlternatives_) GetClass() LiteralValueAlternativesClassLike {
	return literalValueAlternativesClass()
}

// Attribute Methods

func (v *literalValueAlternatives_) GetLiteralValues() col.Sequential[LiteralValueLike] {
	return v.literalValues_
}

// PROTECTED INTERFACE

// Instance Structure

type literalValueAlternatives_ struct {
	// Declare the instance attributes.
	literalValues_ col.Sequential[LiteralValueLike]
}

// Class Structure

type literalValueAlternativesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func literalValueAlternativesClass() *literalValueAlternativesClass_ {
	return literalValueAlternativesClassReference_
}

var literalValueAlternativesClassReference_ = &literalValueAlternativesClass_{
	// Initialize the class constants.
}
