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

/*
Package "module" defines type aliases for the commonly used types defined in the
packages contained in this module.  It also provides a universal constructor for
each commonly used class that is exported by the module.  Each constructor
delegates the actual construction process to its corresponding concrete class
defined in the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - github.com/craterdog/go-syntax-notation/wiki
*/
package module

// TYPE ALIASES

// Ast

type (
	AlternativeLike = ast.AlternativeLike
	CardinalityLike = ast.CardinalityLike
	CharacterLike   = ast.CharacterLike
	ConstrainedLike = ast.ConstrainedLike
	DefinitionLike  = ast.DefinitionLike
	ElementLike     = ast.ElementLike
	ExplicitLike    = ast.ExplicitLike
	ExpressionLike  = ast.ExpressionLike
	ExtentLike      = ast.ExtentLike
	FilterLike      = ast.FilterLike
	GroupLike       = ast.GroupLike
	IdentifierLike  = ast.IdentifierLike
	InlineLike      = ast.InlineLike
	LimitLike       = ast.LimitLike
	LineLike        = ast.LineLike
	MultilineLike   = ast.MultilineLike
	NoticeLike      = ast.NoticeLike
	OptionLike      = ast.OptionLike
	PatternLike     = ast.PatternLike
	QuantifiedLike  = ast.QuantifiedLike
	ReferenceLike   = ast.ReferenceLike
	RepetitionLike  = ast.RepetitionLike
	RuleLike        = ast.RuleLike
	SyntaxLike      = ast.SyntaxLike
	TermLike        = ast.TermLike
	TextLike        = ast.TextLike
)

// Grammar

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

// UNIVERSAL CONSTRUCTORS

// GLOBAL FUNCTIONS
