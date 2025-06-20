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
│            This "package_api.go" file was automatically generated.           │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.

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
package grammar

import (
	fra "github.com/craterdog/go-component-framework/v7"
	ast "github.com/craterdog/go-syntax-notation/v7/ast"
)

// TYPE DECLARATIONS

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	AllcapsToken
	CommentToken
	DelimiterToken
	GlyphToken
	IntrinsicToken
	LiteralToken
	LowercaseToken
	NewlineToken
	NoteToken
	NumberToken
	SpaceToken
	UppercaseToken
)

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
FormatterClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructor Methods
	Formatter() FormatterLike
}

/*
ParserClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor Methods
	Parser() ParserLike
}

/*
ProcessorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor Methods
	Processor() ProcessorLike
}

/*
ScannerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructor Methods
	Scanner(
		source string,
		tokens fra.QueueLike[TokenLike],
	) ScannerLike

	// Function Methods
	FormatToken(
		token TokenLike,
	) string
	FormatType(
		tokenType TokenType,
	) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type TokenClassLike interface {
	// Constructor Methods
	Token(
		line uint,
		position uint,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructor Methods
	Validator() ValidatorLike
}

/*
VisitorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor Methods
	Visitor(
		processor Methodical,
	) VisitorLike
}

// INSTANCE DECLARATIONS

/*
FormatterLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Principal Methods
	GetClass() FormatterClassLike
	FormatSyntax(
		syntax ast.SyntaxLike,
	) string

	// Aspect Interfaces
	Methodical
}

/*
ParserLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Principal Methods
	GetClass() ParserClassLike
	ParseSource(
		source string,
	) ast.SyntaxLike
}

/*
ProcessorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Principal Methods
	GetClass() ProcessorClassLike

	// Aspect Interfaces
	Methodical
}

/*
ScannerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Principal Methods
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Principal Methods
	GetClass() TokenClassLike

	// Attribute Methods
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Principal Methods
	GetClass() ValidatorClassLike
	ValidateSyntax(
		syntax ast.SyntaxLike,
	)

	// Aspect Interfaces
	Methodical
}

/*
VisitorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Principal Methods
	GetClass() VisitorClassLike
	VisitSyntax(
		syntax ast.SyntaxLike,
	)
}

// ASPECT DECLARATIONS

/*
Methodical declares the set of method signatures that must be supported by
all methodical processors.
*/
type Methodical interface {
	ProcessAllcaps(
		allcaps string,
	)
	ProcessComment(
		comment string,
	)
	ProcessDelimiter(
		delimiter string,
	)
	ProcessGlyph(
		glyph string,
	)
	ProcessIntrinsic(
		intrinsic string,
	)
	ProcessLiteral(
		literal string,
	)
	ProcessLowercase(
		lowercase string,
	)
	ProcessNewline(
		newline string,
	)
	ProcessNote(
		note string,
	)
	ProcessNumber(
		number string,
	)
	ProcessSpace(
		space string,
	)
	ProcessUppercase(
		uppercase string,
	)
	PreprocessAlternativeSequence(
		alternativeSequence ast.AlternativeSequenceLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAlternativeSequence(
		alternativeSequence ast.AlternativeSequenceLike,
		index_ uint,
		count_ uint,
	)
	ProcessAlternativeSequenceSlot(
		alternativeSequence ast.AlternativeSequenceLike,
		slot_ uint,
	)
	PreprocessAlternatives(
		alternatives ast.AlternativesLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAlternatives(
		alternatives ast.AlternativesLike,
		index_ uint,
		count_ uint,
	)
	ProcessAlternativesSlot(
		alternatives ast.AlternativesLike,
		slot_ uint,
	)
	PreprocessCardinality(
		cardinality ast.CardinalityLike,
		index_ uint,
		count_ uint,
	)
	PostprocessCardinality(
		cardinality ast.CardinalityLike,
		index_ uint,
		count_ uint,
	)
	ProcessCardinalitySlot(
		cardinality ast.CardinalityLike,
		slot_ uint,
	)
	PreprocessCharacter(
		character ast.CharacterLike,
		index_ uint,
		count_ uint,
	)
	PostprocessCharacter(
		character ast.CharacterLike,
		index_ uint,
		count_ uint,
	)
	ProcessCharacterSlot(
		character ast.CharacterLike,
		slot_ uint,
	)
	PreprocessComponent(
		component ast.ComponentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessComponent(
		component ast.ComponentLike,
		index_ uint,
		count_ uint,
	)
	ProcessComponentSlot(
		component ast.ComponentLike,
		slot_ uint,
	)
	PreprocessConstrained(
		constrained ast.ConstrainedLike,
		index_ uint,
		count_ uint,
	)
	PostprocessConstrained(
		constrained ast.ConstrainedLike,
		index_ uint,
		count_ uint,
	)
	ProcessConstrainedSlot(
		constrained ast.ConstrainedLike,
		slot_ uint,
	)
	PreprocessDefinition(
		definition ast.DefinitionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessDefinition(
		definition ast.DefinitionLike,
		index_ uint,
		count_ uint,
	)
	ProcessDefinitionSlot(
		definition ast.DefinitionLike,
		slot_ uint,
	)
	PreprocessElement(
		element ast.ElementLike,
		index_ uint,
		count_ uint,
	)
	PostprocessElement(
		element ast.ElementLike,
		index_ uint,
		count_ uint,
	)
	ProcessElementSlot(
		element ast.ElementLike,
		slot_ uint,
	)
	PreprocessExplicit(
		explicit ast.ExplicitLike,
		index_ uint,
		count_ uint,
	)
	PostprocessExplicit(
		explicit ast.ExplicitLike,
		index_ uint,
		count_ uint,
	)
	ProcessExplicitSlot(
		explicit ast.ExplicitLike,
		slot_ uint,
	)
	PreprocessExpression(
		expression ast.ExpressionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessExpression(
		expression ast.ExpressionLike,
		index_ uint,
		count_ uint,
	)
	ProcessExpressionSlot(
		expression ast.ExpressionLike,
		slot_ uint,
	)
	PreprocessExpressionAlternatives(
		expressionAlternatives ast.ExpressionAlternativesLike,
		index_ uint,
		count_ uint,
	)
	PostprocessExpressionAlternatives(
		expressionAlternatives ast.ExpressionAlternativesLike,
		index_ uint,
		count_ uint,
	)
	ProcessExpressionAlternativesSlot(
		expressionAlternatives ast.ExpressionAlternativesLike,
		slot_ uint,
	)
	PreprocessExpressionName(
		expressionName ast.ExpressionNameLike,
		index_ uint,
		count_ uint,
	)
	PostprocessExpressionName(
		expressionName ast.ExpressionNameLike,
		index_ uint,
		count_ uint,
	)
	ProcessExpressionNameSlot(
		expressionName ast.ExpressionNameLike,
		slot_ uint,
	)
	PreprocessExtent(
		extent ast.ExtentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessExtent(
		extent ast.ExtentLike,
		index_ uint,
		count_ uint,
	)
	ProcessExtentSlot(
		extent ast.ExtentLike,
		slot_ uint,
	)
	PreprocessFilter(
		filter ast.FilterLike,
		index_ uint,
		count_ uint,
	)
	PostprocessFilter(
		filter ast.FilterLike,
		index_ uint,
		count_ uint,
	)
	ProcessFilterSlot(
		filter ast.FilterLike,
		slot_ uint,
	)
	PreprocessFragment(
		fragment ast.FragmentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessFragment(
		fragment ast.FragmentLike,
		index_ uint,
		count_ uint,
	)
	ProcessFragmentSlot(
		fragment ast.FragmentLike,
		slot_ uint,
	)
	PreprocessGroup(
		group ast.GroupLike,
		index_ uint,
		count_ uint,
	)
	PostprocessGroup(
		group ast.GroupLike,
		index_ uint,
		count_ uint,
	)
	ProcessGroupSlot(
		group ast.GroupLike,
		slot_ uint,
	)
	PreprocessImplicit(
		implicit ast.ImplicitLike,
		index_ uint,
		count_ uint,
	)
	PostprocessImplicit(
		implicit ast.ImplicitLike,
		index_ uint,
		count_ uint,
	)
	ProcessImplicitSlot(
		implicit ast.ImplicitLike,
		slot_ uint,
	)
	PreprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
		index_ uint,
		count_ uint,
	)
	ProcessLegalNoticeSlot(
		legalNotice ast.LegalNoticeLike,
		slot_ uint,
	)
	PreprocessLimit(
		limit ast.LimitLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLimit(
		limit ast.LimitLike,
		index_ uint,
		count_ uint,
	)
	ProcessLimitSlot(
		limit ast.LimitLike,
		slot_ uint,
	)
	PreprocessLiteralAlternatives(
		literalAlternatives ast.LiteralAlternativesLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLiteralAlternatives(
		literalAlternatives ast.LiteralAlternativesLike,
		index_ uint,
		count_ uint,
	)
	ProcessLiteralAlternativesSlot(
		literalAlternatives ast.LiteralAlternativesLike,
		slot_ uint,
	)
	PreprocessLiteralValue(
		literalValue ast.LiteralValueLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLiteralValue(
		literalValue ast.LiteralValueLike,
		index_ uint,
		count_ uint,
	)
	ProcessLiteralValueSlot(
		literalValue ast.LiteralValueLike,
		slot_ uint,
	)
	PreprocessPattern(
		pattern ast.PatternLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPattern(
		pattern ast.PatternLike,
		index_ uint,
		count_ uint,
	)
	ProcessPatternSlot(
		pattern ast.PatternLike,
		slot_ uint,
	)
	PreprocessQuantified(
		quantified ast.QuantifiedLike,
		index_ uint,
		count_ uint,
	)
	PostprocessQuantified(
		quantified ast.QuantifiedLike,
		index_ uint,
		count_ uint,
	)
	ProcessQuantifiedSlot(
		quantified ast.QuantifiedLike,
		slot_ uint,
	)
	PreprocessRepetition(
		repetition ast.RepetitionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessRepetition(
		repetition ast.RepetitionLike,
		index_ uint,
		count_ uint,
	)
	ProcessRepetitionSlot(
		repetition ast.RepetitionLike,
		slot_ uint,
	)
	PreprocessRule(
		rule ast.RuleLike,
		index_ uint,
		count_ uint,
	)
	PostprocessRule(
		rule ast.RuleLike,
		index_ uint,
		count_ uint,
	)
	ProcessRuleSlot(
		rule ast.RuleLike,
		slot_ uint,
	)
	PreprocessRuleAlternatives(
		ruleAlternatives ast.RuleAlternativesLike,
		index_ uint,
		count_ uint,
	)
	PostprocessRuleAlternatives(
		ruleAlternatives ast.RuleAlternativesLike,
		index_ uint,
		count_ uint,
	)
	ProcessRuleAlternativesSlot(
		ruleAlternatives ast.RuleAlternativesLike,
		slot_ uint,
	)
	PreprocessRuleName(
		ruleName ast.RuleNameLike,
		index_ uint,
		count_ uint,
	)
	PostprocessRuleName(
		ruleName ast.RuleNameLike,
		index_ uint,
		count_ uint,
	)
	ProcessRuleNameSlot(
		ruleName ast.RuleNameLike,
		slot_ uint,
	)
	PreprocessRuleTerm(
		ruleTerm ast.RuleTermLike,
		index_ uint,
		count_ uint,
	)
	PostprocessRuleTerm(
		ruleTerm ast.RuleTermLike,
		index_ uint,
		count_ uint,
	)
	ProcessRuleTermSlot(
		ruleTerm ast.RuleTermLike,
		slot_ uint,
	)
	PreprocessSequence(
		sequence ast.SequenceLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSequence(
		sequence ast.SequenceLike,
		index_ uint,
		count_ uint,
	)
	ProcessSequenceSlot(
		sequence ast.SequenceLike,
		slot_ uint,
	)
	PreprocessSyntax(
		syntax ast.SyntaxLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSyntax(
		syntax ast.SyntaxLike,
		index_ uint,
		count_ uint,
	)
	ProcessSyntaxSlot(
		syntax ast.SyntaxLike,
		slot_ uint,
	)
	PreprocessTermSequence(
		termSequence ast.TermSequenceLike,
		index_ uint,
		count_ uint,
	)
	PostprocessTermSequence(
		termSequence ast.TermSequenceLike,
		index_ uint,
		count_ uint,
	)
	ProcessTermSequenceSlot(
		termSequence ast.TermSequenceLike,
		slot_ uint,
	)
	PreprocessText(
		text ast.TextLike,
		index_ uint,
		count_ uint,
	)
	PostprocessText(
		text ast.TextLike,
		index_ uint,
		count_ uint,
	)
	ProcessTextSlot(
		text ast.TextLike,
		slot_ uint,
	)
}
