/*
................................................................................
.                  Copyright (c) 2024.  All Rights Reserved.                   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package example

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Complex() ComplexClassLike {
	return complexReference()
}

// Constructor Methods

func (c *complexClass_) Make(
	realPart float64,
	imaginaryPart float64,
	form Form,
) ComplexLike {
	if uti.IsUndefined(realPart) {
		panic("The \"realPart\" attribute is required by this class.")
	}
	if uti.IsUndefined(imaginaryPart) {
		panic("The \"imaginaryPart\" attribute is required by this class.")
	}
	if uti.IsUndefined(form) {
		panic("The \"form\" attribute is required by this class.")
	}
	var instance = &complex_{
		// Initialize the instance attributes.
		realPart_:      realPart,
		imaginaryPart_: imaginaryPart,
		form_:          form,
	}
	return instance

}

func (c *complexClass_) MakeFromValue(
	value complex128,
) ComplexLike {
	var instance ComplexLike
	// TBD - Add the constructor implementation.
	return instance

}

// Constant Methods

func (c *complexClass_) Zero() ComplexLike {
	return c.zero_
}

func (c *complexClass_) Infinity() ComplexLike {
	return c.infinity_
}

// Function Methods

func (c *complexClass_) Inverse(
	value ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBD - Add the function implementation.
	return result_
}

func (c *complexClass_) Sum(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBD - Add the function implementation.
	return result_
}

func (c *complexClass_) Difference(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBD - Add the function implementation.
	return result_
}

func (c *complexClass_) Reciprocal(
	value ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBD - Add the function implementation.
	return result_
}

func (c *complexClass_) Product(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBD - Add the function implementation.
	return result_
}

func (c *complexClass_) Quotient(
	first ComplexLike,
	second ComplexLike,
) ComplexLike {
	var result_ ComplexLike
	// TBD - Add the function implementation.
	return result_
}

func (c *complexClass_) Norm(
	function NormFunction[ComplexLike],
	value ComplexLike,
) float64 {
	var result_ float64
	// TBD - Add the function implementation.
	return result_
}

// INSTANCE INTERFACE

// Primary Methods

func (v *complex_) GetClass() ComplexClassLike {
	return complexReference()
}

func (v *complex_) IsReal() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

func (v *complex_) IsImaginary() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

// Attribute Methods

func (v *complex_) GetRealPart() float64 {
	return v.realPart_
}

func (v *complex_) GetImaginaryPart() float64 {
	return v.imaginaryPart_
}

func (v *complex_) GetForm() Form {
	return v.form_
}

func (v *complex_) SetForm(
	form Form,
) {
	if uti.IsUndefined(form) {
		panic("The \"form\" attribute is required by this class.")
	}
	v.form_ = form
}

// Continuous Methods

func (v *complex_) IsZero() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

func (v *complex_) IsDiscrete() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

func (v *complex_) IsInfinity() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type complex_ struct {
	// Declare the instance attributes.
	realPart_      float64
	imaginaryPart_ float64
	form_          Form
}

// Class Structure

type complexClass_ struct {
	// Declare the class constants.
	zero_     ComplexLike
	infinity_ ComplexLike
}

// Class Reference

func complexReference() *complexClass_ {
	return complexReference_
}

var complexReference_ = &complexClass_{
	// Initialize the class constants.
	// zero_: constantValue,
	// infinity_: constantValue,
}
