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
  - https://github.com/craterdog/go-syntax-notation/wiki
*/
package module

import (
	fra "github.com/craterdog/go-collection-framework/v8"
	ast "github.com/craterdog/go-syntax-notation/v8/ast"
	gra "github.com/craterdog/go-syntax-notation/v8/grammar"
)

// TYPE ALIASES

// Ast

type (
	AlternativeSequenceClassLike    = ast.AlternativeSequenceClassLike
	AlternativesClassLike           = ast.AlternativesClassLike
	CardinalityClassLike            = ast.CardinalityClassLike
	CharacterClassLike              = ast.CharacterClassLike
	ComponentClassLike              = ast.ComponentClassLike
	ConstrainedClassLike            = ast.ConstrainedClassLike
	DefinitionClassLike             = ast.DefinitionClassLike
	ElementClassLike                = ast.ElementClassLike
	ExplicitClassLike               = ast.ExplicitClassLike
	ExpressionClassLike             = ast.ExpressionClassLike
	ExpressionAlternativesClassLike = ast.ExpressionAlternativesClassLike
	ExpressionNameClassLike         = ast.ExpressionNameClassLike
	ExtentClassLike                 = ast.ExtentClassLike
	FilterClassLike                 = ast.FilterClassLike
	FragmentClassLike               = ast.FragmentClassLike
	GroupClassLike                  = ast.GroupClassLike
	ImplicitClassLike               = ast.ImplicitClassLike
	LegalNoticeClassLike            = ast.LegalNoticeClassLike
	LimitClassLike                  = ast.LimitClassLike
	LiteralAlternativesClassLike    = ast.LiteralAlternativesClassLike
	LiteralValueClassLike           = ast.LiteralValueClassLike
	PatternClassLike                = ast.PatternClassLike
	QuantifiedClassLike             = ast.QuantifiedClassLike
	RepetitionClassLike             = ast.RepetitionClassLike
	RuleClassLike                   = ast.RuleClassLike
	RuleAlternativesClassLike       = ast.RuleAlternativesClassLike
	RuleNameClassLike               = ast.RuleNameClassLike
	RuleTermClassLike               = ast.RuleTermClassLike
	SequenceClassLike               = ast.SequenceClassLike
	SyntaxClassLike                 = ast.SyntaxClassLike
	TermSequenceClassLike           = ast.TermSequenceClassLike
	TextClassLike                   = ast.TextClassLike
)

type (
	AlternativeSequenceLike    = ast.AlternativeSequenceLike
	AlternativesLike           = ast.AlternativesLike
	CardinalityLike            = ast.CardinalityLike
	CharacterLike              = ast.CharacterLike
	ComponentLike              = ast.ComponentLike
	ConstrainedLike            = ast.ConstrainedLike
	DefinitionLike             = ast.DefinitionLike
	ElementLike                = ast.ElementLike
	ExplicitLike               = ast.ExplicitLike
	ExpressionLike             = ast.ExpressionLike
	ExpressionAlternativesLike = ast.ExpressionAlternativesLike
	ExpressionNameLike         = ast.ExpressionNameLike
	ExtentLike                 = ast.ExtentLike
	FilterLike                 = ast.FilterLike
	FragmentLike               = ast.FragmentLike
	GroupLike                  = ast.GroupLike
	ImplicitLike               = ast.ImplicitLike
	LegalNoticeLike            = ast.LegalNoticeLike
	LimitLike                  = ast.LimitLike
	LiteralAlternativesLike    = ast.LiteralAlternativesLike
	LiteralValueLike           = ast.LiteralValueLike
	PatternLike                = ast.PatternLike
	QuantifiedLike             = ast.QuantifiedLike
	RepetitionLike             = ast.RepetitionLike
	RuleLike                   = ast.RuleLike
	RuleAlternativesLike       = ast.RuleAlternativesLike
	RuleNameLike               = ast.RuleNameLike
	RuleTermLike               = ast.RuleTermLike
	SequenceLike               = ast.SequenceLike
	SyntaxLike                 = ast.SyntaxLike
	TermSequenceLike           = ast.TermSequenceLike
	TextLike                   = ast.TextLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken     = gra.ErrorToken
	AllcapsToken   = gra.AllcapsToken
	CommentToken   = gra.CommentToken
	DelimiterToken = gra.DelimiterToken
	GlyphToken     = gra.GlyphToken
	IntrinsicToken = gra.IntrinsicToken
	LiteralToken   = gra.LiteralToken
	LowercaseToken = gra.LowercaseToken
	NewlineToken   = gra.NewlineToken
	NoteToken      = gra.NoteToken
	NumberToken    = gra.NumberToken
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

func AlternativeSequenceClass() AlternativeSequenceClassLike {
	return ast.AlternativeSequenceClass()
}

func AlternativeSequence(
	delimiter string,
	sequence ast.SequenceLike,
) AlternativeSequenceLike {
	return AlternativeSequenceClass().AlternativeSequence(
		delimiter,
		sequence,
	)
}

func AlternativesClass() AlternativesClassLike {
	return ast.AlternativesClass()
}

func Alternatives(
	sequence ast.SequenceLike,
	alternativeSequences fra.Sequential[ast.AlternativeSequenceLike],
) AlternativesLike {
	return AlternativesClass().Alternatives(
		sequence,
		alternativeSequences,
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
	delimiter1 string,
	lowercase string,
	delimiter2 string,
	pattern ast.PatternLike,
) ExpressionLike {
	return ExpressionClass().Expression(
		delimiter1,
		lowercase,
		delimiter2,
		pattern,
	)
}

func ExpressionAlternativesClass() ExpressionAlternativesClassLike {
	return ast.ExpressionAlternativesClass()
}

func ExpressionAlternatives(
	expressionNames fra.Sequential[ast.ExpressionNameLike],
) ExpressionAlternativesLike {
	return ExpressionAlternativesClass().ExpressionAlternatives(
		expressionNames,
	)
}

func ExpressionNameClass() ExpressionNameClassLike {
	return ast.ExpressionNameClass()
}

func ExpressionName(
	newline string,
	lowercase string,
	optionalNote string,
) ExpressionNameLike {
	return ExpressionNameClass().ExpressionName(
		newline,
		lowercase,
		optionalNote,
	)
}

func ExtentClass() ExtentClassLike {
	return ast.ExtentClass()
}

func Extent(
	delimiter string,
	glyph string,
) ExtentLike {
	return ExtentClass().Extent(
		delimiter,
		glyph,
	)
}

func FilterClass() FilterClassLike {
	return ast.FilterClass()
}

func Filter(
	optionalDelimiter string,
	delimiter1 string,
	characters fra.Sequential[ast.CharacterLike],
	delimiter2 string,
) FilterLike {
	return FilterClass().Filter(
		optionalDelimiter,
		delimiter1,
		characters,
		delimiter2,
	)
}

func FragmentClass() FragmentClassLike {
	return ast.FragmentClass()
}

func Fragment(
	delimiter1 string,
	allcaps string,
	delimiter2 string,
	pattern ast.PatternLike,
) FragmentLike {
	return FragmentClass().Fragment(
		delimiter1,
		allcaps,
		delimiter2,
		pattern,
	)
}

func GroupClass() GroupClassLike {
	return ast.GroupClass()
}

func Group(
	delimiter1 string,
	alternatives ast.AlternativesLike,
	delimiter2 string,
) GroupLike {
	return GroupClass().Group(
		delimiter1,
		alternatives,
		delimiter2,
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

func LegalNoticeClass() LegalNoticeClassLike {
	return ast.LegalNoticeClass()
}

func LegalNotice(
	comment string,
) LegalNoticeLike {
	return LegalNoticeClass().LegalNotice(
		comment,
	)
}

func LimitClass() LimitClassLike {
	return ast.LimitClass()
}

func Limit(
	delimiter string,
	optionalNumber string,
) LimitLike {
	return LimitClass().Limit(
		delimiter,
		optionalNumber,
	)
}

func LiteralAlternativesClass() LiteralAlternativesClassLike {
	return ast.LiteralAlternativesClass()
}

func LiteralAlternatives(
	literalValues fra.Sequential[ast.LiteralValueLike],
) LiteralAlternativesLike {
	return LiteralAlternativesClass().LiteralAlternatives(
		literalValues,
	)
}

func LiteralValueClass() LiteralValueClassLike {
	return ast.LiteralValueClass()
}

func LiteralValue(
	newline string,
	literal string,
	optionalNote string,
) LiteralValueLike {
	return LiteralValueClass().LiteralValue(
		newline,
		literal,
		optionalNote,
	)
}

func PatternClass() PatternClassLike {
	return ast.PatternClass()
}

func Pattern(
	alternatives ast.AlternativesLike,
	optionalNote string,
) PatternLike {
	return PatternClass().Pattern(
		alternatives,
		optionalNote,
	)
}

func QuantifiedClass() QuantifiedClassLike {
	return ast.QuantifiedClass()
}

func Quantified(
	delimiter1 string,
	number string,
	optionalLimit ast.LimitLike,
	delimiter2 string,
) QuantifiedLike {
	return QuantifiedClass().Quantified(
		delimiter1,
		number,
		optionalLimit,
		delimiter2,
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
	delimiter1 string,
	uppercase string,
	delimiter2 string,
	definition ast.DefinitionLike,
) RuleLike {
	return RuleClass().Rule(
		delimiter1,
		uppercase,
		delimiter2,
		definition,
	)
}

func RuleAlternativesClass() RuleAlternativesClassLike {
	return ast.RuleAlternativesClass()
}

func RuleAlternatives(
	ruleNames fra.Sequential[ast.RuleNameLike],
) RuleAlternativesLike {
	return RuleAlternativesClass().RuleAlternatives(
		ruleNames,
	)
}

func RuleNameClass() RuleNameClassLike {
	return ast.RuleNameClass()
}

func RuleName(
	newline string,
	uppercase string,
	optionalNote string,
) RuleNameLike {
	return RuleNameClass().RuleName(
		newline,
		uppercase,
		optionalNote,
	)
}

func RuleTermClass() RuleTermClassLike {
	return ast.RuleTermClass()
}

func RuleTerm(
	component ast.ComponentLike,
	optionalCardinality ast.CardinalityLike,
) RuleTermLike {
	return RuleTermClass().RuleTerm(
		component,
		optionalCardinality,
	)
}

func SequenceClass() SequenceClassLike {
	return ast.SequenceClass()
}

func Sequence(
	repetitions fra.Sequential[ast.RepetitionLike],
) SequenceLike {
	return SequenceClass().Sequence(
		repetitions,
	)
}

func SyntaxClass() SyntaxClassLike {
	return ast.SyntaxClass()
}

func Syntax(
	legalNotice ast.LegalNoticeLike,
	comment1 string,
	rules fra.Sequential[ast.RuleLike],
	comment2 string,
	expressions fra.Sequential[ast.ExpressionLike],
	comment3 string,
	fragments fra.Sequential[ast.FragmentLike],
) SyntaxLike {
	return SyntaxClass().Syntax(
		legalNotice,
		comment1,
		rules,
		comment2,
		expressions,
		comment3,
		fragments,
	)
}

func TermSequenceClass() TermSequenceClassLike {
	return ast.TermSequenceClass()
}

func TermSequence(
	ruleTerms fra.Sequential[ast.RuleTermLike],
	optionalNote string,
) TermSequenceLike {
	return TermSequenceClass().TermSequence(
		ruleTerms,
		optionalNote,
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
