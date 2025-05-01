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

func ReferenceClass() ReferenceClassLike {
	return referenceClass()
}

// Constructor Methods

func (c *referenceClass_) Reference(
	identifier IdentifierLike,
	optionalCardinality CardinalityLike,
) ReferenceLike {
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
	}
	var instance = &reference_{
		// Initialize the instance attributes.
		identifier_:          identifier,
		optionalCardinality_: optionalCardinality,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *reference_) GetClass() ReferenceClassLike {
	return referenceClass()
}

// Attribute Methods

func (v *reference_) GetIdentifier() IdentifierLike {
	return v.identifier_
}

func (v *reference_) GetOptionalCardinality() CardinalityLike {
	return v.optionalCardinality_
}

// PROTECTED INTERFACE

// Instance Structure

type reference_ struct {
	// Declare the instance attributes.
	identifier_          IdentifierLike
	optionalCardinality_ CardinalityLike
}

// Class Structure

type referenceClass_ struct {
	// Declare the class constants.
}

// Class Reference

func referenceClass() *referenceClass_ {
	return referenceClassReference_
}

var referenceClassReference_ = &referenceClass_{
	// Initialize the class constants.
}
