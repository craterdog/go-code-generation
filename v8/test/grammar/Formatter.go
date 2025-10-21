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
	ast "github.com/craterdog/go-syntax-notation/v8/ast"
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

func (v *formatter_) FormatSyntax(syntax ast.SyntaxLike) string {
	VisitorClass().Visitor(v).VisitSyntax(syntax)
	return v.getResult()
}

// Methodical Methods

func (v *formatter_) ProcessAllcaps(
	allcaps string,
) {
	v.appendString(allcaps)
}

func (v *formatter_) ProcessComment(
	comment string,
) {
	v.appendString(comment)
}

func (v *formatter_) ProcessDelimiter(
	delimiter string,
) {
	v.appendString(delimiter)
}

func (v *formatter_) ProcessGlyph(
	glyph string,
) {
	v.appendString(glyph)
}

func (v *formatter_) ProcessIntrinsic(
	intrinsic string,
) {
	v.appendString(intrinsic)
}

func (v *formatter_) ProcessLiteral(
	literal string,
) {
	v.appendString(literal)
}

func (v *formatter_) ProcessLowercase(
	lowercase string,
) {
	v.appendString(lowercase)
}

func (v *formatter_) ProcessNewline(
	newline string,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessNote(
	note string,
) {
	v.appendString("  ")
	v.appendString(note)
}

func (v *formatter_) ProcessNumber(
	number string,
) {
	v.appendString(number)
}

func (v *formatter_) ProcessSpace(
	space string,
) {
	v.appendString(space)
}

func (v *formatter_) ProcessUppercase(
	uppercase string,
) {
	v.appendString(uppercase)
}

func (v *formatter_) PreprocessCharacter(
	character ast.CharacterLike,
	index_ uint,
	count_ uint,
) {
	if index_ > 1 {
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessRepetition(
	repetition ast.RepetitionLike,
	index_ uint,
	count_ uint,
) {
	if index_ > 1 {
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessAlternativeSequence(
	alternativeSequence ast.AlternativeSequenceLike,
	index_ uint,
	count_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessAlternativeSequenceSlot(
	alternativeSequence ast.AlternativeSequenceLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessDefinition(
	definition ast.DefinitionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_++
}

func (v *formatter_) PostprocessDefinition(
	definition ast.DefinitionLike,
	index_ uint,
	count_ uint,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PostprocessExpression(
	expression ast.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PostprocessFragment(
	fragment ast.FragmentLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PostprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessPattern(
	pattern ast.PatternLike,
	index_ uint,
	count_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PostprocessPattern(
	pattern ast.PatternLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PostprocessRule(
	rule ast.RuleLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessRuleTerm(
	ruleTerm ast.RuleTermLike,
	index_ uint,
	count_ uint,
) {
	v.appendString(" ")
}

const _indentation = "    "

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
