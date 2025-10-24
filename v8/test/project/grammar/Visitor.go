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

package grammar

import (
	ast "github.com/craterdog/go-example-syntax/ast"
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func VisitorClass() VisitorClassLike {
	return visitorClass()
}

// Constructor Methods

func (c *visitorClass_) Visitor(
	processor Methodical,
) VisitorLike {
	if uti.IsUndefined(processor) {
		panic("The \"processor\" attribute is required by this class.")
	}
	var instance = &visitor_{
		// Initialize the instance attributes.
		processor_: processor,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorClass()
}

func (v *visitor_) VisitDocument(
	document ast.DocumentLike,
) {
	v.processor_.PreprocessDocument(
		document,
		0,
		0,
	)
	v.visitDocument(document)
	v.processor_.PostprocessDocument(
		document,
		0,
		0,
	)
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAdditional(
	additional ast.AdditionalLike,
) {
	var delimiter = additional.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAdditionalSlot(
		additional,
		1,
	)

	var component = additional.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		0,
		0,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		0,
		0,
	)
}

func (v *visitor_) visitComponent(
	component ast.ComponentLike,
) {
	// Visit the possible component rule types.
	switch actual := component.GetAny().(type) {
	case ast.IntrinsicLike:
		v.processor_.PreprocessIntrinsic(
			actual,
			0,
			0,
		)
		v.visitIntrinsic(actual)
		v.processor_.PostprocessIntrinsic(
			actual,
			0,
			0,
		)
	case ast.ListLike:
		v.processor_.PreprocessList(
			actual,
			0,
			0,
		)
		v.visitList(actual)
		v.processor_.PostprocessList(
			actual,
			0,
			0,
		)
	case ast.KeywordLike:
		v.processor_.PreprocessKeyword(
			actual,
			0,
			0,
		)
		v.visitKeyword(actual)
		v.processor_.PostprocessKeyword(
			actual,
			0,
			0,
		)
	}
}

func (v *visitor_) visitDocument(
	document ast.DocumentLike,
) {
	var components = document.GetComponents()
	v.processor_.PreprocessComponent(
		components,
		0,
		0,
	)
	v.visitComponent(components)
	v.processor_.PostprocessComponent(
		components,
		0,
		0,
	)
}

func (v *visitor_) visitIntrinsic(
	intrinsic ast.IntrinsicLike,
) {
	// Visit the possible intrinsic expression types.
	var actual = intrinsic.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, NumberToken):
		v.processor_.ProcessNumber(actual)
	case ScannerClass().MatchesType(actual, RuneToken):
		v.processor_.ProcessRune(actual)
	case ScannerClass().MatchesType(actual, TextToken):
		v.processor_.ProcessText(actual)
	}
}

func (v *visitor_) visitKeyword(
	keyword ast.KeywordLike,
) {
	// Visit the possible keyword literal values.
	var actual = keyword.GetAny().(string)
	switch actual {
	case "if":
		v.processor_.ProcessDelimiter("if")
	case "then":
		v.processor_.ProcessDelimiter("then")
	case "else":
		v.processor_.ProcessDelimiter("else")
	}
}

func (v *visitor_) visitList(
	list ast.ListLike,
) {
	var delimiter1 = list.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessListSlot(
		list,
		1,
	)

	var component = list.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		0,
		0,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		0,
		0,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessListSlot(
		list,
		2,
	)

	var additionalsIndex uint
	var additionals = list.GetAdditionals().GetIterator()
	var additionalsCount = uint(additionals.GetSize())
	for additionals.HasNext() {
		additionalsIndex++
		var rule = additionals.GetNext()
		v.processor_.PreprocessAdditional(
			rule,
			additionalsIndex,
			additionalsCount,
		)
		v.visitAdditional(rule)
		v.processor_.PostprocessAdditional(
			rule,
			additionalsIndex,
			additionalsCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessListSlot(
		list,
		3,
	)

	var delimiter2 = list.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

// Instance Structure

type visitor_ struct {
	// Declare the instance attributes.
	processor_ Methodical
}

// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func visitorClass() *visitorClass_ {
	return visitorClassReference_
}

var visitorClassReference_ = &visitorClass_{
	// Initialize the class constants.
}
