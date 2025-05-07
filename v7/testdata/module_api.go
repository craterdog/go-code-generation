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
│             This "module_api.go" file was automatically generated.           │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-syntax-notation/wiki
*/
package module

import (
	col "github.com/craterdog/go-collection-framework/v7"
	ast "github.com/craterdog/go-syntax-notation/v7/ast"
	gra "github.com/craterdog/go-syntax-notation/v7/grammar"
)

// TYPE ALIASES

// Ast

type (
	AlternativeClassLike      = ast.AlternativeClassLike
	CardinalityClassLike      = ast.CardinalityClassLike
	CharacterClassLike        = ast.CharacterClassLike
	ConstrainedClassLike      = ast.ConstrainedClassLike
	DefinitionClassLike       = ast.DefinitionClassLike
	ElementClassLike          = ast.ElementClassLike
	ExplicitClassLike         = ast.ExplicitClassLike
	ExpressionClassLike       = ast.ExpressionClassLike
	ExpressionOptionClassLike = ast.ExpressionOptionClassLike
	ExtentClassLike           = ast.ExtentClassLike
	FilterClassLike           = ast.FilterClassLike
	GroupClassLike            = ast.GroupClassLike
	IdentifierClassLike       = ast.IdentifierClassLike
	ImplicitClassLike         = ast.ImplicitClassLike
	InlineClassLike           = ast.InlineClassLike
	LimitClassLike            = ast.LimitClassLike
	LiteralClassLike          = ast.LiteralClassLike
	LiteralOptionClassLike    = ast.LiteralOptionClassLike
	MultiexpressionClassLike  = ast.MultiexpressionClassLike
	MultiliteralClassLike     = ast.MultiliteralClassLike
	MultiruleClassLike        = ast.MultiruleClassLike
	NoticeClassLike           = ast.NoticeClassLike
	OptionClassLike           = ast.OptionClassLike
	PatternClassLike          = ast.PatternClassLike
	QuantifiedClassLike       = ast.QuantifiedClassLike
	ReferenceClassLike        = ast.ReferenceClassLike
	RepetitionClassLike       = ast.RepetitionClassLike
	RuleClassLike             = ast.RuleClassLike
	RuleOptionClassLike       = ast.RuleOptionClassLike
	SyntaxClassLike           = ast.SyntaxClassLike
	TermClassLike             = ast.TermClassLike
	TextClassLike             = ast.TextClassLike
)

type (
	AlternativeLike      = ast.AlternativeLike
	CardinalityLike      = ast.CardinalityLike
	CharacterLike        = ast.CharacterLike
	ConstrainedLike      = ast.ConstrainedLike
	DefinitionLike       = ast.DefinitionLike
	ElementLike          = ast.ElementLike
	ExplicitLike         = ast.ExplicitLike
	ExpressionLike       = ast.ExpressionLike
	ExpressionOptionLike = ast.ExpressionOptionLike
	ExtentLike           = ast.ExtentLike
	FilterLike           = ast.FilterLike
	GroupLike            = ast.GroupLike
	IdentifierLike       = ast.IdentifierLike
	ImplicitLike         = ast.ImplicitLike
	InlineLike           = ast.InlineLike
	LimitLike            = ast.LimitLike
	LiteralLike          = ast.LiteralLike
	LiteralOptionLike    = ast.LiteralOptionLike
	MultiexpressionLike  = ast.MultiexpressionLike
	MultiliteralLike     = ast.MultiliteralLike
	MultiruleLike        = ast.MultiruleLike
	NoticeLike           = ast.NoticeLike
	OptionLike           = ast.OptionLike
	PatternLike          = ast.PatternLike
	QuantifiedLike       = ast.QuantifiedLike
	ReferenceLike        = ast.ReferenceLike
	RepetitionLike       = ast.RepetitionLike
	RuleLike             = ast.RuleLike
	RuleOptionLike       = ast.RuleOptionLike
	SyntaxLike           = ast.SyntaxLike
	TermLike             = ast.TermLike
	TextLike             = ast.TextLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken     = gra.ErrorToken
	CommentToken   = gra.CommentToken
	DelimiterToken = gra.DelimiterToken
	ExcludedToken  = gra.ExcludedToken
	GlyphToken     = gra.GlyphToken
	IntrinsicToken = gra.IntrinsicToken
	LowercaseToken = gra.LowercaseToken
	NewlineToken   = gra.NewlineToken
	NoteToken      = gra.NoteToken
	NumberToken    = gra.NumberToken
	OptionalToken  = gra.OptionalToken
	QuoteToken     = gra.QuoteToken
	RepeatedToken  = gra.RepeatedToken
	SpaceToken     = gra.SpaceToken
	UppercaseToken = gra.UppercaseToken
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

func AlternativeClass() AlternativeClassLike {
	return ast.AlternativeClass()
}

func Alternative(
	option ast.OptionLike,
) AlternativeLike {
	return AlternativeClass().Alternative(
		option,
	)
}

func CardinalityClass() CardinalityClassLike {
	return ast.CardinalityClass()
}

func Cardinality(
	any_ any,
) CardinalityLike {
	return CardinalityClass().Cardinality(
		any_,
	)
}

func CharacterClass() CharacterClassLike {
	return ast.CharacterClass()
}

func Character(
	any_ any,
) CharacterLike {
	return CharacterClass().Character(
		any_,
	)
}

func ConstrainedClass() ConstrainedClassLike {
	return ast.ConstrainedClass()
}

func Constrained(
	any_ any,
) ConstrainedLike {
	return ConstrainedClass().Constrained(
		any_,
	)
}

func DefinitionClass() DefinitionClassLike {
	return ast.DefinitionClass()
}

func Definition(
	any_ any,
) DefinitionLike {
	return DefinitionClass().Definition(
		any_,
	)
}

func ElementClass() ElementClassLike {
	return ast.ElementClass()
}

func Element(
	any_ any,
) ElementLike {
	return ElementClass().Element(
		any_,
	)
}

func ExplicitClass() ExplicitClassLike {
	return ast.ExplicitClass()
}

func Explicit(
	glyph string,
	optionalExtent ast.ExtentLike,
) ExplicitLike {
	return ExplicitClass().Explicit(
		glyph,
		optionalExtent,
	)
}

func ExpressionClass() ExpressionClassLike {
	return ast.ExpressionClass()
}

func Expression(
	lowercase string,
	pattern ast.PatternLike,
	optionalNote string,
) ExpressionLike {
	return ExpressionClass().Expression(
		lowercase,
		pattern,
		optionalNote,
	)
}

func ExpressionOptionClass() ExpressionOptionClassLike {
	return ast.ExpressionOptionClass()
}

func ExpressionOption(
	newline string,
	lowercase string,
	optionalNote string,
) ExpressionOptionLike {
	return ExpressionOptionClass().ExpressionOption(
		newline,
		lowercase,
		optionalNote,
	)
}

func ExtentClass() ExtentClassLike {
	return ast.ExtentClass()
}

func Extent(
	glyph string,
) ExtentLike {
	return ExtentClass().Extent(
		glyph,
	)
}

func FilterClass() FilterClassLike {
	return ast.FilterClass()
}

func Filter(
	optionalExcluded string,
	characters col.Sequential[ast.CharacterLike],
) FilterLike {
	return FilterClass().Filter(
		optionalExcluded,
		characters,
	)
}

func GroupClass() GroupClassLike {
	return ast.GroupClass()
}

func Group(
	pattern ast.PatternLike,
) GroupLike {
	return GroupClass().Group(
		pattern,
	)
}

func IdentifierClass() IdentifierClassLike {
	return ast.IdentifierClass()
}

func Identifier(
	any_ any,
) IdentifierLike {
	return IdentifierClass().Identifier(
		any_,
	)
}

func ImplicitClass() ImplicitClassLike {
	return ast.ImplicitClass()
}

func Implicit(
	intrinsic string,
) ImplicitLike {
	return ImplicitClass().Implicit(
		intrinsic,
	)
}

func InlineClass() InlineClassLike {
	return ast.InlineClass()
}

func Inline(
	terms col.Sequential[ast.TermLike],
	optionalNote string,
) InlineLike {
	return InlineClass().Inline(
		terms,
		optionalNote,
	)
}

func LimitClass() LimitClassLike {
	return ast.LimitClass()
}

func Limit(
	optionalNumber string,
) LimitLike {
	return LimitClass().Limit(
		optionalNumber,
	)
}

func LiteralClass() LiteralClassLike {
	return ast.LiteralClass()
}

func Literal(
	quote string,
) LiteralLike {
	return LiteralClass().Literal(
		quote,
	)
}

func LiteralOptionClass() LiteralOptionClassLike {
	return ast.LiteralOptionClass()
}

func LiteralOption(
	newline string,
	quote string,
	optionalNote string,
) LiteralOptionLike {
	return LiteralOptionClass().LiteralOption(
		newline,
		quote,
		optionalNote,
	)
}

func MultiexpressionClass() MultiexpressionClassLike {
	return ast.MultiexpressionClass()
}

func Multiexpression(
	expressionOptions col.Sequential[ast.ExpressionOptionLike],
) MultiexpressionLike {
	return MultiexpressionClass().Multiexpression(
		expressionOptions,
	)
}

func MultiliteralClass() MultiliteralClassLike {
	return ast.MultiliteralClass()
}

func Multiliteral(
	literalOptions col.Sequential[ast.LiteralOptionLike],
) MultiliteralLike {
	return MultiliteralClass().Multiliteral(
		literalOptions,
	)
}

func MultiruleClass() MultiruleClassLike {
	return ast.MultiruleClass()
}

func Multirule(
	ruleOptions col.Sequential[ast.RuleOptionLike],
) MultiruleLike {
	return MultiruleClass().Multirule(
		ruleOptions,
	)
}

func NoticeClass() NoticeClassLike {
	return ast.NoticeClass()
}

func Notice(
	comment string,
	newline string,
) NoticeLike {
	return NoticeClass().Notice(
		comment,
		newline,
	)
}

func OptionClass() OptionClassLike {
	return ast.OptionClass()
}

func Option(
	repetitions col.Sequential[ast.RepetitionLike],
) OptionLike {
	return OptionClass().Option(
		repetitions,
	)
}

func PatternClass() PatternClassLike {
	return ast.PatternClass()
}

func Pattern(
	option ast.OptionLike,
	alternatives col.Sequential[ast.AlternativeLike],
) PatternLike {
	return PatternClass().Pattern(
		option,
		alternatives,
	)
}

func QuantifiedClass() QuantifiedClassLike {
	return ast.QuantifiedClass()
}

func Quantified(
	number string,
	optionalLimit ast.LimitLike,
) QuantifiedLike {
	return QuantifiedClass().Quantified(
		number,
		optionalLimit,
	)
}

func ReferenceClass() ReferenceClassLike {
	return ast.ReferenceClass()
}

func Reference(
	identifier ast.IdentifierLike,
	optionalCardinality ast.CardinalityLike,
) ReferenceLike {
	return ReferenceClass().Reference(
		identifier,
		optionalCardinality,
	)
}

func RepetitionClass() RepetitionClassLike {
	return ast.RepetitionClass()
}

func Repetition(
	element ast.ElementLike,
	optionalCardinality ast.CardinalityLike,
) RepetitionLike {
	return RepetitionClass().Repetition(
		element,
		optionalCardinality,
	)
}

func RuleClass() RuleClassLike {
	return ast.RuleClass()
}

func Rule(
	uppercase string,
	definition ast.DefinitionLike,
) RuleLike {
	return RuleClass().Rule(
		uppercase,
		definition,
	)
}

func RuleOptionClass() RuleOptionClassLike {
	return ast.RuleOptionClass()
}

func RuleOption(
	newline string,
	uppercase string,
	optionalNote string,
) RuleOptionLike {
	return RuleOptionClass().RuleOption(
		newline,
		uppercase,
		optionalNote,
	)
}

func SyntaxClass() SyntaxClassLike {
	return ast.SyntaxClass()
}

func Syntax(
	notice ast.NoticeLike,
	comment1 string,
	rules col.Sequential[ast.RuleLike],
	comment2 string,
	expressions col.Sequential[ast.ExpressionLike],
) SyntaxLike {
	return SyntaxClass().Syntax(
		notice,
		comment1,
		rules,
		comment2,
		expressions,
	)
}

func TermClass() TermClassLike {
	return ast.TermClass()
}

func Term(
	any_ any,
) TermLike {
	return TermClass().Term(
		any_,
	)
}

func TextClass() TextClassLike {
	return ast.TextClass()
}

func Text(
	any_ any,
) TextLike {
	return TextClass().Text(
		any_,
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
	tokens col.QueueLike[gra.TokenLike],
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

func FormatSyntax(
	syntax SyntaxLike,
) string {
	var formatter = Formatter()
	return formatter.FormatSyntax(syntax)
}

func MatchesType(
	tokenValue string,
	tokenType TokenType,
) bool {
	var scannerClass = gra.ScannerClass()
	return scannerClass.MatchesType(tokenValue, tokenType)
}

func ParseSource(
	source string,
) SyntaxLike {
	var parser = Parser()
	return parser.ParseSource(source)
}

func ValidateSyntax(
	syntax SyntaxLike,
) {
	var validator = Validator()
	validator.ValidateSyntax(syntax)
}
