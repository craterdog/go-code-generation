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
│         This "package_api.go" file was automatically generated using:        │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-syntax-notation/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AlternativeSequenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete alternative-sequence-like class.
*/
type AlternativeSequenceClassLike interface {
	// Constructor Methods
	AlternativeSequence(
		delimiter string,
		sequence SequenceLike,
	) AlternativeSequenceLike
}

/*
AlternativesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete alternatives-like class.
*/
type AlternativesClassLike interface {
	// Constructor Methods
	Alternatives(
		sequence SequenceLike,
		alternativeSequences fra.Sequential[AlternativeSequenceLike],
	) AlternativesLike
}

/*
CardinalityClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete cardinality-like class.
*/
type CardinalityClassLike interface {
	// Constructor Methods
	Cardinality(
		any_ any,
	) CardinalityLike
}

/*
CharacterClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete character-like class.
*/
type CharacterClassLike interface {
	// Constructor Methods
	Character(
		any_ any,
	) CharacterLike
}

/*
ComponentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete component-like class.
*/
type ComponentClassLike interface {
	// Constructor Methods
	Component(
		any_ any,
	) ComponentLike
}

/*
ConstrainedClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete constrained-like class.
*/
type ConstrainedClassLike interface {
	// Constructor Methods
	Constrained(
		any_ any,
	) ConstrainedLike
}

/*
DefinitionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete definition-like class.
*/
type DefinitionClassLike interface {
	// Constructor Methods
	Definition(
		any_ any,
	) DefinitionLike
}

/*
ElementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete element-like class.
*/
type ElementClassLike interface {
	// Constructor Methods
	Element(
		any_ any,
	) ElementLike
}

/*
ExplicitClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete explicit-like class.
*/
type ExplicitClassLike interface {
	// Constructor Methods
	Explicit(
		glyph string,
		optionalExtent ExtentLike,
	) ExplicitLike
}

/*
ExpressionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete expression-like class.
*/
type ExpressionClassLike interface {
	// Constructor Methods
	Expression(
		delimiter1 string,
		lowercase string,
		delimiter2 string,
		pattern PatternLike,
	) ExpressionLike
}

/*
ExpressionAlternativesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete expression-alternatives-like class.
*/
type ExpressionAlternativesClassLike interface {
	// Constructor Methods
	ExpressionAlternatives(
		expressionNames fra.Sequential[ExpressionNameLike],
	) ExpressionAlternativesLike
}

/*
ExpressionNameClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete expression-name-like class.
*/
type ExpressionNameClassLike interface {
	// Constructor Methods
	ExpressionName(
		newline string,
		lowercase string,
		optionalNote string,
	) ExpressionNameLike
}

/*
ExtentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete extent-like class.
*/
type ExtentClassLike interface {
	// Constructor Methods
	Extent(
		delimiter string,
		glyph string,
	) ExtentLike
}

/*
FilterClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete filter-like class.
*/
type FilterClassLike interface {
	// Constructor Methods
	Filter(
		optionalDelimiter string,
		delimiter1 string,
		characters fra.Sequential[CharacterLike],
		delimiter2 string,
	) FilterLike
}

/*
FragmentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete fragment-like class.
*/
type FragmentClassLike interface {
	// Constructor Methods
	Fragment(
		delimiter1 string,
		allcaps string,
		delimiter2 string,
		pattern PatternLike,
	) FragmentLike
}

/*
GroupClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete group-like class.
*/
type GroupClassLike interface {
	// Constructor Methods
	Group(
		delimiter1 string,
		alternatives AlternativesLike,
		delimiter2 string,
	) GroupLike
}

/*
ImplicitClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete implicit-like class.
*/
type ImplicitClassLike interface {
	// Constructor Methods
	Implicit(
		intrinsic string,
	) ImplicitLike
}

/*
LegalNoticeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete legal-notice-like class.
*/
type LegalNoticeClassLike interface {
	// Constructor Methods
	LegalNotice(
		comment string,
	) LegalNoticeLike
}

/*
LimitClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete limit-like class.
*/
type LimitClassLike interface {
	// Constructor Methods
	Limit(
		delimiter string,
		optionalNumber string,
	) LimitLike
}

/*
LiteralAlternativesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete literal-alternatives-like class.
*/
type LiteralAlternativesClassLike interface {
	// Constructor Methods
	LiteralAlternatives(
		literalValues fra.Sequential[LiteralValueLike],
	) LiteralAlternativesLike
}

/*
LiteralValueClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete literal-value-like class.
*/
type LiteralValueClassLike interface {
	// Constructor Methods
	LiteralValue(
		newline string,
		literal string,
		optionalNote string,
	) LiteralValueLike
}

/*
PatternClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete pattern-like class.
*/
type PatternClassLike interface {
	// Constructor Methods
	Pattern(
		alternatives AlternativesLike,
		optionalNote string,
	) PatternLike
}

/*
QuantifiedClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete quantified-like class.
*/
type QuantifiedClassLike interface {
	// Constructor Methods
	Quantified(
		delimiter1 string,
		number string,
		optionalLimit LimitLike,
		delimiter2 string,
	) QuantifiedLike
}

/*
RepetitionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete repetition-like class.
*/
type RepetitionClassLike interface {
	// Constructor Methods
	Repetition(
		element ElementLike,
		optionalCardinality CardinalityLike,
	) RepetitionLike
}

/*
RuleClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete rule-like class.
*/
type RuleClassLike interface {
	// Constructor Methods
	Rule(
		delimiter1 string,
		uppercase string,
		delimiter2 string,
		definition DefinitionLike,
	) RuleLike
}

/*
RuleAlternativesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete rule-alternatives-like class.
*/
type RuleAlternativesClassLike interface {
	// Constructor Methods
	RuleAlternatives(
		ruleNames fra.Sequential[RuleNameLike],
	) RuleAlternativesLike
}

/*
RuleNameClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete rule-name-like class.
*/
type RuleNameClassLike interface {
	// Constructor Methods
	RuleName(
		newline string,
		uppercase string,
		optionalNote string,
	) RuleNameLike
}

/*
RuleTermClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete rule-term-like class.
*/
type RuleTermClassLike interface {
	// Constructor Methods
	RuleTerm(
		component ComponentLike,
		optionalCardinality CardinalityLike,
	) RuleTermLike
}

/*
SequenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete sequence-like class.
*/
type SequenceClassLike interface {
	// Constructor Methods
	Sequence(
		repetitions fra.Sequential[RepetitionLike],
	) SequenceLike
}

/*
SyntaxClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete syntax-like class.
*/
type SyntaxClassLike interface {
	// Constructor Methods
	Syntax(
		legalNotice LegalNoticeLike,
		comment1 string,
		rules fra.Sequential[RuleLike],
		comment2 string,
		expressions fra.Sequential[ExpressionLike],
		comment3 string,
		fragments fra.Sequential[FragmentLike],
	) SyntaxLike
}

/*
TermSequenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete term-sequence-like class.
*/
type TermSequenceClassLike interface {
	// Constructor Methods
	TermSequence(
		ruleTerms fra.Sequential[RuleTermLike],
		optionalNote string,
	) TermSequenceLike
}

/*
TextClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete text-like class.
*/
type TextClassLike interface {
	// Constructor Methods
	Text(
		any_ any,
	) TextLike
}

// INSTANCE DECLARATIONS

/*
AlternativeSequenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete alternative-sequence-like class.
*/
type AlternativeSequenceLike interface {
	// Principal Methods
	GetClass() AlternativeSequenceClassLike

	// Attribute Methods
	GetDelimiter() string
	GetSequence() SequenceLike
}

/*
AlternativesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete alternatives-like class.
*/
type AlternativesLike interface {
	// Principal Methods
	GetClass() AlternativesClassLike

	// Attribute Methods
	GetSequence() SequenceLike
	GetAlternativeSequences() fra.Sequential[AlternativeSequenceLike]
}

/*
CardinalityLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete cardinality-like class.
*/
type CardinalityLike interface {
	// Principal Methods
	GetClass() CardinalityClassLike

	// Attribute Methods
	GetAny() any
}

/*
CharacterLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete character-like class.
*/
type CharacterLike interface {
	// Principal Methods
	GetClass() CharacterClassLike

	// Attribute Methods
	GetAny() any
}

/*
ComponentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete component-like class.
*/
type ComponentLike interface {
	// Principal Methods
	GetClass() ComponentClassLike

	// Attribute Methods
	GetAny() any
}

/*
ConstrainedLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete constrained-like class.
*/
type ConstrainedLike interface {
	// Principal Methods
	GetClass() ConstrainedClassLike

	// Attribute Methods
	GetAny() any
}

/*
DefinitionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete definition-like class.
*/
type DefinitionLike interface {
	// Principal Methods
	GetClass() DefinitionClassLike

	// Attribute Methods
	GetAny() any
}

/*
ElementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete element-like class.
*/
type ElementLike interface {
	// Principal Methods
	GetClass() ElementClassLike

	// Attribute Methods
	GetAny() any
}

/*
ExplicitLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete explicit-like class.
*/
type ExplicitLike interface {
	// Principal Methods
	GetClass() ExplicitClassLike

	// Attribute Methods
	GetGlyph() string
	GetOptionalExtent() ExtentLike
}

/*
ExpressionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete expression-like class.
*/
type ExpressionLike interface {
	// Principal Methods
	GetClass() ExpressionClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetLowercase() string
	GetDelimiter2() string
	GetPattern() PatternLike
}

/*
ExpressionAlternativesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete expression-alternatives-like class.
*/
type ExpressionAlternativesLike interface {
	// Principal Methods
	GetClass() ExpressionAlternativesClassLike

	// Attribute Methods
	GetExpressionNames() fra.Sequential[ExpressionNameLike]
}

/*
ExpressionNameLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete expression-name-like class.
*/
type ExpressionNameLike interface {
	// Principal Methods
	GetClass() ExpressionNameClassLike

	// Attribute Methods
	GetNewline() string
	GetLowercase() string
	GetOptionalNote() string
}

/*
ExtentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete extent-like class.
*/
type ExtentLike interface {
	// Principal Methods
	GetClass() ExtentClassLike

	// Attribute Methods
	GetDelimiter() string
	GetGlyph() string
}

/*
FilterLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete filter-like class.
*/
type FilterLike interface {
	// Principal Methods
	GetClass() FilterClassLike

	// Attribute Methods
	GetOptionalDelimiter() string
	GetDelimiter1() string
	GetCharacters() fra.Sequential[CharacterLike]
	GetDelimiter2() string
}

/*
FragmentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete fragment-like class.
*/
type FragmentLike interface {
	// Principal Methods
	GetClass() FragmentClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetAllcaps() string
	GetDelimiter2() string
	GetPattern() PatternLike
}

/*
GroupLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete group-like class.
*/
type GroupLike interface {
	// Principal Methods
	GetClass() GroupClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetAlternatives() AlternativesLike
	GetDelimiter2() string
}

/*
ImplicitLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete implicit-like class.
*/
type ImplicitLike interface {
	// Principal Methods
	GetClass() ImplicitClassLike

	// Attribute Methods
	GetIntrinsic() string
}

/*
LegalNoticeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete legal-notice-like class.
*/
type LegalNoticeLike interface {
	// Principal Methods
	GetClass() LegalNoticeClassLike

	// Attribute Methods
	GetComment() string
}

/*
LimitLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete limit-like class.
*/
type LimitLike interface {
	// Principal Methods
	GetClass() LimitClassLike

	// Attribute Methods
	GetDelimiter() string
	GetOptionalNumber() string
}

/*
LiteralAlternativesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete literal-alternatives-like class.
*/
type LiteralAlternativesLike interface {
	// Principal Methods
	GetClass() LiteralAlternativesClassLike

	// Attribute Methods
	GetLiteralValues() fra.Sequential[LiteralValueLike]
}

/*
LiteralValueLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete literal-value-like class.
*/
type LiteralValueLike interface {
	// Principal Methods
	GetClass() LiteralValueClassLike

	// Attribute Methods
	GetNewline() string
	GetLiteral() string
	GetOptionalNote() string
}

/*
PatternLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete pattern-like class.
*/
type PatternLike interface {
	// Principal Methods
	GetClass() PatternClassLike

	// Attribute Methods
	GetAlternatives() AlternativesLike
	GetOptionalNote() string
}

/*
QuantifiedLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete quantified-like class.
*/
type QuantifiedLike interface {
	// Principal Methods
	GetClass() QuantifiedClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetNumber() string
	GetOptionalLimit() LimitLike
	GetDelimiter2() string
}

/*
RepetitionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete repetition-like class.
*/
type RepetitionLike interface {
	// Principal Methods
	GetClass() RepetitionClassLike

	// Attribute Methods
	GetElement() ElementLike
	GetOptionalCardinality() CardinalityLike
}

/*
RuleLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete rule-like class.
*/
type RuleLike interface {
	// Principal Methods
	GetClass() RuleClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetUppercase() string
	GetDelimiter2() string
	GetDefinition() DefinitionLike
}

/*
RuleAlternativesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete rule-alternatives-like class.
*/
type RuleAlternativesLike interface {
	// Principal Methods
	GetClass() RuleAlternativesClassLike

	// Attribute Methods
	GetRuleNames() fra.Sequential[RuleNameLike]
}

/*
RuleNameLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete rule-name-like class.
*/
type RuleNameLike interface {
	// Principal Methods
	GetClass() RuleNameClassLike

	// Attribute Methods
	GetNewline() string
	GetUppercase() string
	GetOptionalNote() string
}

/*
RuleTermLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete rule-term-like class.
*/
type RuleTermLike interface {
	// Principal Methods
	GetClass() RuleTermClassLike

	// Attribute Methods
	GetComponent() ComponentLike
	GetOptionalCardinality() CardinalityLike
}

/*
SequenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete sequence-like class.
*/
type SequenceLike interface {
	// Principal Methods
	GetClass() SequenceClassLike

	// Attribute Methods
	GetRepetitions() fra.Sequential[RepetitionLike]
}

/*
SyntaxLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete syntax-like class.
*/
type SyntaxLike interface {
	// Principal Methods
	GetClass() SyntaxClassLike

	// Attribute Methods
	GetLegalNotice() LegalNoticeLike
	GetComment1() string
	GetRules() fra.Sequential[RuleLike]
	GetComment2() string
	GetExpressions() fra.Sequential[ExpressionLike]
	GetComment3() string
	GetFragments() fra.Sequential[FragmentLike]
}

/*
TermSequenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete term-sequence-like class.
*/
type TermSequenceLike interface {
	// Principal Methods
	GetClass() TermSequenceClassLike

	// Attribute Methods
	GetRuleTerms() fra.Sequential[RuleTermLike]
	GetOptionalNote() string
}

/*
TextLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete text-like class.
*/
type TextLike interface {
	// Principal Methods
	GetClass() TextClassLike

	// Attribute Methods
	GetAny() any
}

// ASPECT DECLARATIONS
