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
	com "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ExpressionAlternativesClass() ExpressionAlternativesClassLike {
	return expressionAlternativesClass()
}

// Constructor Methods

func (c *expressionAlternativesClass_) ExpressionAlternatives(
	expressionNames com.ListLike[ExpressionNameLike],
) ExpressionAlternativesLike {
	if uti.IsUndefined(expressionNames) {
		panic("The \"expressionNames\" attribute is required by this class.")
	}
	var instance = &expressionAlternatives_{
		// Initialize the instance attributes.
		expressionNames_: expressionNames,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *expressionAlternatives_) GetClass() ExpressionAlternativesClassLike {
	return expressionAlternativesClass()
}

// Attribute Methods

func (v *expressionAlternatives_) GetExpressionNames() com.ListLike[ExpressionNameLike] {
	return v.expressionNames_
}

// PROTECTED INTERFACE

// Instance Structure

type expressionAlternatives_ struct {
	// Declare the instance attributes.
	expressionNames_ com.ListLike[ExpressionNameLike]
}

// Class Structure

type expressionAlternativesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func expressionAlternativesClass() *expressionAlternativesClass_ {
	return expressionAlternativesClassReference_
}

var expressionAlternativesClassReference_ = &expressionAlternativesClass_{
	// Initialize the class constants.
}
