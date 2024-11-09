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

func Syntax() SyntaxClassLike {
	return syntaxReference()
}

// Constructor Methods

func (c *syntaxClass_) Make(
	notice NoticeLike,
	comment1 string,
	rules abs.Sequential[RuleLike],
	comment2 string,
	expressions abs.Sequential[ExpressionLike],
) SyntaxLike {
	if uti.IsUndefined(notice) {
		panic("The \"notice\" attribute is required by this class.")
	}
	if uti.IsUndefined(comment1) {
		panic("The \"comment1\" attribute is required by this class.")
	}
	if uti.IsUndefined(rules) {
		panic("The \"rules\" attribute is required by this class.")
	}
	if uti.IsUndefined(comment2) {
		panic("The \"comment2\" attribute is required by this class.")
	}
	if uti.IsUndefined(expressions) {
		panic("The \"expressions\" attribute is required by this class.")
	}
	var instance = &syntax_{
		// Initialize the instance attributes.
		notice_:      notice,
		comment1_:    comment1,
		rules_:       rules,
		comment2_:    comment2,
		expressions_: expressions,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *syntax_) GetClass() SyntaxClassLike {
	return syntaxReference()
}

// Attribute Methods

func (v *syntax_) GetNotice() NoticeLike {
	return v.notice_
}

func (v *syntax_) GetComment1() string {
	return v.comment1_
}

func (v *syntax_) GetRules() abs.Sequential[RuleLike] {
	return v.rules_
}

func (v *syntax_) GetComment2() string {
	return v.comment2_
}

func (v *syntax_) GetExpressions() abs.Sequential[ExpressionLike] {
	return v.expressions_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type syntax_ struct {
	// Declare the instance attributes.
	notice_      NoticeLike
	comment1_    string
	rules_       abs.Sequential[RuleLike]
	comment2_    string
	expressions_ abs.Sequential[ExpressionLike]
}

// Class Structure

type syntaxClass_ struct {
	// Declare the class constants.
}

// Class Reference

func syntaxReference() *syntaxClass_ {
	return syntaxReference_
}

var syntaxReference_ = &syntaxClass_{
	// Initialize the class constants.
}
