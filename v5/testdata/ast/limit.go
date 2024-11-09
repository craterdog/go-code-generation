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

// CLASS INTERFACE

// Access Function

func Limit() LimitClassLike {
	return limitReference()
}

// Constructor Methods

func (c *limitClass_) Make(
	optionalNumber string,
) LimitLike {
	var instance = &limit_{
		// Initialize the instance attributes.
		optionalNumber_: optionalNumber,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *limit_) GetClass() LimitClassLike {
	return limitReference()
}

// Attribute Methods

func (v *limit_) GetOptionalNumber() string {
	return v.optionalNumber_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type limit_ struct {
	// Declare the instance attributes.
	optionalNumber_ string
}

// Class Structure

type limitClass_ struct {
	// Declare the class constants.
}

// Class Reference

func limitReference() *limitClass_ {
	return limitReference_
}

var limitReference_ = &limitClass_{
	// Initialize the class constants.
}
