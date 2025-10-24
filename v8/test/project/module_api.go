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
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-example-notation/wiki
*/
package module

import (
	fra "github.com/craterdog/go-collection-framework/v8"
	ast "github.com/craterdog/go-example-syntax/ast"
	gra "github.com/craterdog/go-example-syntax/grammar"
)

// TYPE ALIASES

// Ast

type (
	AdditionalClassLike = ast.AdditionalClassLike
	ComponentClassLike  = ast.ComponentClassLike
	DocumentClassLike   = ast.DocumentClassLike
	IntrinsicClassLike  = ast.IntrinsicClassLike
	KeywordClassLike    = ast.KeywordClassLike
	ListClassLike       = ast.ListClassLike
)

type (
	AdditionalLike = ast.AdditionalLike
	ComponentLike  = ast.ComponentLike
	DocumentLike   = ast.DocumentLike
	IntrinsicLike  = ast.IntrinsicLike
	KeywordLike    = ast.KeywordLike
	ListLike       = ast.ListLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken     = gra.ErrorToken
	DelimiterToken = gra.DelimiterToken
	NewlineToken   = gra.NewlineToken
	NumberToken    = gra.NumberToken
	RuneToken      = gra.RuneToken
	SpaceToken     = gra.SpaceToken
	TextToken      = gra.TextToken
)

type (
	FormatterClassLike = gra.FormatterClassLike
	ParserClassLike    = gra.ParserClassLike
	ProcessorClassLike = gra.ProcessorClassLike
	ScannerClassLike   = gra.ScannerClassLike
	TokenClassLike     = gra.TokenClassLike
	ValidatorClassLike = gra.ValidatorClassLike
	VisitorClassLike   = gra.VisitorClassLike
)

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

type (
	Methodical = gra.Methodical
)

// CLASS ACCESSORS

// Ast

func AdditionalClass() AdditionalClassLike {
	return ast.AdditionalClass()
}

func Additional(
	delimiter string,
	component ast.ComponentLike,
) AdditionalLike {
	return AdditionalClass().Additional(
		delimiter,
		component,
	)
}

func ComponentClass() ComponentClassLike {
	return ast.ComponentClass()
}

func Component(
	any_ any,
) ComponentLike {
	return ComponentClass().Component(
		any_,
	)
}

func DocumentClass() DocumentClassLike {
	return ast.DocumentClass()
}

func Document(
	components fra.Sequential[ast.ComponentLike],
) DocumentLike {
	return DocumentClass().Document(
		components,
	)
}

func IntrinsicClass() IntrinsicClassLike {
	return ast.IntrinsicClass()
}

func Intrinsic(
	any_ any,
) IntrinsicLike {
	return IntrinsicClass().Intrinsic(
		any_,
	)
}

func KeywordClass() KeywordClassLike {
	return ast.KeywordClass()
}

func Keyword(
	any_ any,
) KeywordLike {
	return KeywordClass().Keyword(
		any_,
	)
}

func ListClass() ListClassLike {
	return ast.ListClass()
}

func List(
	delimiter1 string,
	component ast.ComponentLike,
	additionals fra.Sequential[ast.AdditionalLike],
	delimiter2 string,
) ListLike {
	return ListClass().List(
		delimiter1,
		component,
		additionals,
		delimiter2,
	)
}

// Grammar

func FormatterClass() FormatterClassLike {
	return gra.FormatterClass()
}

func Formatter() FormatterLike {
	return FormatterClass().Formatter()
}

func ParserClass() ParserClassLike {
	return gra.ParserClass()
}

func Parser() ParserLike {
	return ParserClass().Parser()
}

func ProcessorClass() ProcessorClassLike {
	return gra.ProcessorClass()
}

func Processor() ProcessorLike {
	return ProcessorClass().Processor()
}

func ScannerClass() ScannerClassLike {
	return gra.ScannerClass()
}

func Scanner(
	source string,
	tokens fra.QueueLike[gra.TokenLike],
) ScannerLike {
	return ScannerClass().Scanner(
		source,
		tokens,
	)
}

func TokenClass() TokenClassLike {
	return gra.TokenClass()
}

func Token(
	line uint,
	position uint,
	type_ gra.TokenType,
	value string,
) TokenLike {
	return TokenClass().Token(
		line,
		position,
		type_,
		value,
	)
}

func ValidatorClass() ValidatorClassLike {
	return gra.ValidatorClass()
}

func Validator() ValidatorLike {
	return ValidatorClass().Validator()
}

func VisitorClass() VisitorClassLike {
	return gra.VisitorClass()
}

func Visitor(
	processor gra.Methodical,
) VisitorLike {
	return VisitorClass().Visitor(
		processor,
	)
}

// GLOBAL FUNCTIONS
