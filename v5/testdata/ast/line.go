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

func Line() LineClassLike {
	return lineReference()
}

// Constructor Methods

func (c *lineClass_) Make(
	identifier IdentifierLike,
	optionalNote string,
) LineLike {
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
	}
	var instance = &line_{
		// Initialize the instance attributes.
		identifier_:   identifier,
		optionalNote_: optionalNote,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *line_) GetClass() LineClassLike {
	return lineReference()
}

// Attribute Methods

func (v *line_) GetIdentifier() IdentifierLike {
	return v.identifier_
}

func (v *line_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type line_ struct {
	// Declare the instance attributes.
	identifier_   IdentifierLike
	optionalNote_ string
}

// Class Structure

type lineClass_ struct {
	// Declare the class constants.
}

// Class Reference

func lineReference() *lineClass_ {
	return lineReference_
}

var lineReference_ = &lineClass_{
	// Initialize the class constants.
}
