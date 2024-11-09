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
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Inline() InlineClassLike {
	return inlineReference()
}

// Constructor Methods

func (c *inlineClass_) Make(
	terms abs.Sequential[TermLike],
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *inline_) GetClass() InlineClassLike {
	return inlineReference()
}

// Attribute Methods

func (v *inline_) GetTerms() abs.Sequential[TermLike] {
	return v.terms_
}

func (v *inline_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type inline_ struct {
	// Declare the instance attributes.
	terms_        abs.Sequential[TermLike]
	optionalNote_ string
}

// Class Structure

type inlineClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inlineReference() *inlineClass_ {
	return inlineReference_
}

var inlineReference_ = &inlineClass_{
	// Initialize the class constants.
}
