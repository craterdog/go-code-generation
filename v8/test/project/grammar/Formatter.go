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
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	ast "github.com/craterdog/go-example-syntax/ast"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func FormatterClass() FormatterClassLike {
	return formatterClass()
}

// Constructor Methods

func (c *formatterClass_) Formatter() FormatterLike {
	var instance = &formatter_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *formatter_) GetClass() FormatterClassLike {
	return formatterClass()
}

func (v *formatter_) FormatDocument(document ast.DocumentLike) string {
	VisitorClass().Visitor(v).VisitDocument(document)
	return v.getResult()
}

// Methodical Methods

func (v *formatter_) ProcessDelimiter(
	delimiter string,
) {
	v.appendString(delimiter)
}

func (v *formatter_) ProcessNewline(
	newline string,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessNumber(
	number string,
) {
	v.appendString(number)
}

func (v *formatter_) ProcessRune(
	rune_ string,
) {
	v.appendString(rune_)
}

func (v *formatter_) ProcessSpace(
	space string,
) {
	v.appendString(space)
}

func (v *formatter_) ProcessText(
	text string,
) {
	v.appendString(text)
}

func (v *formatter_) PreprocessAdditional(
	additional ast.AdditionalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessAdditionalSlot(
	additional ast.AdditionalLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessComponent(
	component ast.ComponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessComponentSlot(
	component ast.ComponentLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessDocument(
	document ast.DocumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessDocumentSlot(
	document ast.DocumentLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessIntrinsic(
	intrinsic ast.IntrinsicLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessIntrinsicSlot(
	intrinsic ast.IntrinsicLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessKeyword(
	keyword ast.KeywordLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessKeywordSlot(
	keyword ast.KeywordLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessList(
	list ast.ListLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessListSlot(
	list ast.ListLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendString(" ")
	}
}

const _indentation = "\t"

// PROTECTED INTERFACE

// Private Methods

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var level uint
	for ; level < v.depth_; level++ {
		newline += _indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(
	string_ string,
) {
	v.result_.WriteString(string_)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}

// Instance Structure

type formatter_ struct {
	// Declare the instance attributes.
	depth_  uint
	result_ sts.Builder

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type formatterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func formatterClass() *formatterClass_ {
	return formatterClassReference_
}

var formatterClassReference_ = &formatterClass_{
	// Initialize the class constants.
}
