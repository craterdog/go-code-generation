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
	col "github.com/craterdog/go-collection-framework/v7/collection"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func MultiexpressionClass() MultiexpressionClassLike {
	return multiexpressionClass()
}

// Constructor Methods

func (c *multiexpressionClass_) Multiexpression(
	expressionOptions col.Sequential[ExpressionOptionLike],
) MultiexpressionLike {
	if uti.IsUndefined(expressionOptions) {
		panic("The \"expressionOptions\" attribute is required by this class.")
	}
	var instance = &multiexpression_{
		// Initialize the instance attributes.
		expressionOptions_: expressionOptions,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multiexpression_) GetClass() MultiexpressionClassLike {
	return multiexpressionClass()
}

// Attribute Methods

func (v *multiexpression_) GetExpressionOptions() col.Sequential[ExpressionOptionLike] {
	return v.expressionOptions_
}

// PROTECTED INTERFACE

// Instance Structure

type multiexpression_ struct {
	// Declare the instance attributes.
	expressionOptions_ col.Sequential[ExpressionOptionLike]
}

// Class Structure

type multiexpressionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multiexpressionClass() *multiexpressionClass_ {
	return multiexpressionClassReference_
}

var multiexpressionClassReference_ = &multiexpressionClass_{
	// Initialize the class constants.
}
