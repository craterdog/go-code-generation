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
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	fra "github.com/craterdog/go-collection-framework/v8"
	uti "github.com/craterdog/go-missing-utilities/v8"
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
	rules fra.Sequential[RuleLike],
	comment2 string,
	expressions fra.Sequential[ExpressionLike],
	comment3 string,
	fragments fra.Sequential[FragmentLike],
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
	if uti.IsUndefined(comment3) {
		panic("The \"comment3\" attribute is required by this class.")
	}
	if uti.IsUndefined(fragments) {
		panic("The \"fragments\" attribute is required by this class.")
	}
	var instance = &syntax_{
		// Initialize the instance attributes.
		legalNotice_: legalNotice,
		comment1_:    comment1,
		rules_:       rules,
		comment2_:    comment2,
		expressions_: expressions,
		comment3_:    comment3,
		fragments_:   fragments,
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

func (v *syntax_) GetRules() fra.Sequential[RuleLike] {
	return v.rules_
}

func (v *syntax_) GetComment2() string {
	return v.comment2_
}

func (v *syntax_) GetExpressions() fra.Sequential[ExpressionLike] {
	return v.expressions_
}

func (v *syntax_) GetComment3() string {
	return v.comment3_
}

func (v *syntax_) GetFragments() fra.Sequential[FragmentLike] {
	return v.fragments_
}

// PROTECTED INTERFACE

// Instance Structure

type syntax_ struct {
	// Declare the instance attributes.
	legalNotice_ LegalNoticeLike
	comment1_    string
	rules_       fra.Sequential[RuleLike]
	comment2_    string
	expressions_ fra.Sequential[ExpressionLike]
	comment3_    string
	fragments_   fra.Sequential[FragmentLike]
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
