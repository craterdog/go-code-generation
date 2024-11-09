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

func Reference() ReferenceClassLike {
	return referenceReference()
}

// Constructor Methods

func (c *referenceClass_) Make(
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *reference_) GetClass() ReferenceClassLike {
	return referenceReference()
}

// Attribute Methods

func (v *reference_) GetIdentifier() IdentifierLike {
	return v.identifier_
}

func (v *reference_) GetOptionalCardinality() CardinalityLike {
	return v.optionalCardinality_
}

// PROTECTED INTERFACE

// Private Methods

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

func referenceReference() *referenceClass_ {
	return referenceReference_
}

var referenceReference_ = &referenceClass_{
	// Initialize the class constants.
}
