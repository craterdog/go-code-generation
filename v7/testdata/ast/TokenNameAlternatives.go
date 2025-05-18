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

func TokenNameAlternativesClass() TokenNameAlternativesClassLike {
	return tokenNameAlternativesClass()
}

// Constructor Methods

func (c *tokenNameAlternativesClass_) TokenNameAlternatives(
	tokenNames col.Sequential[TokenNameLike],
) TokenNameAlternativesLike {
	if uti.IsUndefined(tokenNames) {
		panic("The \"tokenNames\" attribute is required by this class.")
	}
	var instance = &tokenNameAlternatives_{
		// Initialize the instance attributes.
		tokenNames_: tokenNames,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *tokenNameAlternatives_) GetClass() TokenNameAlternativesClassLike {
	return tokenNameAlternativesClass()
}

// Attribute Methods

func (v *tokenNameAlternatives_) GetTokenNames() col.Sequential[TokenNameLike] {
	return v.tokenNames_
}

// PROTECTED INTERFACE

// Instance Structure

type tokenNameAlternatives_ struct {
	// Declare the instance attributes.
	tokenNames_ col.Sequential[TokenNameLike]
}

// Class Structure

type tokenNameAlternativesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func tokenNameAlternativesClass() *tokenNameAlternativesClass_ {
	return tokenNameAlternativesClassReference_
}

var tokenNameAlternativesClassReference_ = &tokenNameAlternativesClass_{
	// Initialize the class constants.
}
