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
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	ast "github.com/craterdog/go-syntax-notation/v7/ast"
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
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *formatter_) GetClass() FormatterClassLike {
	return formatterClass()
}

func (v *formatter_) FormatSyntax(syntax ast.SyntaxLike) string {
	v.visitor_.VisitSyntax(syntax)
	return v.getResult()
}

// Methodical Methods

func (v *formatter_) ProcessComment(
	comment string,
) {
	v.appendString(comment)
}

func (v *formatter_) ProcessExcluded(
	excluded string,
) {
	v.appendString(excluded)
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

func (v *formatter_) PreprocessLiteralOption(
	literalOption ast.LiteralOptionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessLowercase(
	lowercase string,
) {
	v.appendString(lowercase)
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

func (v *formatter_) ProcessOptional(
	optional string,
) {
	v.appendString(optional)
}

func (v *formatter_) ProcessQuote(
	quote string,
) {
	v.appendString(quote)
}

func (v *formatter_) ProcessRepeated(
	repeated string,
) {
	v.appendString(repeated)
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

func (v *formatter_) PreprocessAlternative(
	alternative ast.AlternativeLike,
	index uint,
	size uint,
) {
	v.appendString(" |")
	if v.depth_ > 0 {
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessCharacter(
	character ast.CharacterLike,
	index uint,
	size uint,
) {
	if index > 1 {
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
	v.appendString("$")
}

func (v *formatter_) ProcessExpressionSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(":")
	}
}

func (v *formatter_) PostprocessExpression(
	expression ast.ExpressionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
	v.appendNewline()
}

func (v *formatter_) PreprocessExpressionOption(
	expressionOption ast.ExpressionOptionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessExtent(
	extent ast.ExtentLike,
) {
	v.appendString("..")
}

func (v *formatter_) ProcessFilterSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString("[")
	}
}

func (v *formatter_) PostprocessFilter(
	filter ast.FilterLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessGroup(
	group ast.GroupLike,
) {
	v.appendString("(")
	v.depth_++
}

func (v *formatter_) PostprocessGroup(
	group ast.GroupLike,
) {
	v.depth_--
	v.appendString(")")
}

func (v *formatter_) PostprocessInline(
	inline ast.InlineLike,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessLimit(
	limit ast.LimitLike,
) {
	v.appendString("..")
}

func (v *formatter_) PreprocessMultiexpression(
	multiexpression ast.MultiexpressionLike,
) {
	v.depth_++
}

func (v *formatter_) PostprocessMultiexpression(
	multiexpression ast.MultiexpressionLike,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessMultiliteral(
	multiliteral ast.MultiliteralLike,
) {
	v.depth_++
}

func (v *formatter_) PostprocessMultiliteral(
	multiliteral ast.MultiliteralLike,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PreprocessMultirule(
	multirule ast.MultiruleLike,
) {
	v.depth_++
}

func (v *formatter_) PostprocessMultirule(
	multirule ast.MultiruleLike,
) {
	v.depth_--
	v.appendNewline()
}

func (v *formatter_) PostprocessNotice(
	notice ast.NoticeLike,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessOption(
	option ast.OptionLike,
) {
	if v.depth_ == 0 {
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessQuantified(
	quantified ast.QuantifiedLike,
) {
	v.appendString("{")
}

func (v *formatter_) PostprocessQuantified(
	quantified ast.QuantifiedLike,
) {
	v.appendString("}")
}

func (v *formatter_) PreprocessRepetition(
	repetition ast.RepetitionLike,
	index uint,
	size uint,
) {
	if index > 1 {
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	v.appendString("$")
}

func (v *formatter_) ProcessRuleSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(":")
	}
}

func (v *formatter_) PostprocessRule(
	rule ast.RuleLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessRuleOption(
	ruleOption ast.RuleOptionLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessTerm(
	term ast.TermLike,
	index uint,
	size uint,
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
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

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
