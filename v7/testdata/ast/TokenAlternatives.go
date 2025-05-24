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

func TokenAlternativesClass() TokenAlternativesClassLike {
	return tokenAlternativesClass()
}

// Constructor Methods

func (c *tokenAlternativesClass_) TokenAlternatives(
	tokenNames col.ListLike[TokenNameLike],
) TokenAlternativesLike {
	if uti.IsUndefined(tokenNames) {
		panic("The \"tokenNames\" attribute is required by this class.")
	}
	var instance = &tokenAlternatives_{
		// Initialize the instance attributes.
		tokenNames_: tokenNames,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *tokenAlternatives_) GetClass() TokenAlternativesClassLike {
	return tokenAlternativesClass()
}

// Attribute Methods

func (v *tokenAlternatives_) GetTokenNames() col.ListLike[TokenNameLike] {
	return v.tokenNames_
}

// PROTECTED INTERFACE

// Instance Structure

type tokenAlternatives_ struct {
	// Declare the instance attributes.
	tokenNames_ col.ListLike[TokenNameLike]
}

// Class Structure

type tokenAlternativesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func tokenAlternativesClass() *tokenAlternativesClass_ {
	return tokenAlternativesClassReference_
}

var tokenAlternativesClassReference_ = &tokenAlternativesClass_{
	// Initialize the class constants.
}
