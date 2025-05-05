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

func InlineClass() InlineClassLike {
	return inlineClass()
}

// Constructor Methods

func (c *inlineClass_) Inline(
	terms col.Sequential[TermLike],
	optionalNote string,
) InlineLike {
	if uti.IsUndefined(terms) {
		panic("The \"terms\" attribute is required by this class.")
	}
	var instance = &inline_{
		// Initialize the instance attributes.
		terms_:        terms,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inline_) GetClass() InlineClassLike {
	return inlineClass()
}

// Attribute Methods

func (v *inline_) GetTerms() col.Sequential[TermLike] {
	return v.terms_
}

func (v *inline_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type inline_ struct {
	// Declare the instance attributes.
	terms_        col.Sequential[TermLike]
	optionalNote_ string
}

// Class Structure

type inlineClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inlineClass() *inlineClass_ {
	return inlineClassReference_
}

var inlineClassReference_ = &inlineClass_{
	// Initialize the class constants.
}
