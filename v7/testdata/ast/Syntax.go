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

func SyntaxClass() SyntaxClassLike {
	return syntaxClass()
}

// Constructor Methods

func (c *syntaxClass_) Syntax(
	legalNotice LegalNoticeLike,
	comment1 string,
	rules col.ListLike[RuleLike],
	comment2 string,
	expressions col.ListLike[ExpressionLike],
) SyntaxLike {
	if uti.IsUndefined(legalNotice) {
		panic("The \"legalNotice\" attribute is required by this class.")
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
		legalNotice_: legalNotice,
		comment1_:    comment1,
		rules_:       rules,
		comment2_:    comment2,
		expressions_: expressions,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *syntax_) GetClass() SyntaxClassLike {
	return syntaxClass()
}

// Attribute Methods

func (v *syntax_) GetLegalNotice() LegalNoticeLike {
	return v.legalNotice_
}

func (v *syntax_) GetComment1() string {
	return v.comment1_
}

func (v *syntax_) GetRules() col.ListLike[RuleLike] {
	return v.rules_
}

func (v *syntax_) GetComment2() string {
	return v.comment2_
}

func (v *syntax_) GetExpressions() col.ListLike[ExpressionLike] {
	return v.expressions_
}

// PROTECTED INTERFACE

// Instance Structure

type syntax_ struct {
	// Declare the instance attributes.
	legalNotice_ LegalNoticeLike
	comment1_    string
	rules_       col.ListLike[RuleLike]
	comment2_    string
	expressions_ col.ListLike[ExpressionLike]
}

// Class Structure

type syntaxClass_ struct {
	// Declare the class constants.
}

// Class Reference

func syntaxClass() *syntaxClass_ {
	return syntaxClassReference_
}

var syntaxClassReference_ = &syntaxClass_{
	// Initialize the class constants.
}
