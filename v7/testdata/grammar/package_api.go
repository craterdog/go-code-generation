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
	col "github.com/craterdog/go-collection-framework/v7"
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
		tokens col.QueueLike[TokenLike],
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
	PreprocessAdditionalCharacter(
		additionalCharacter ast.AdditionalCharacterLike,
		index uint,
		count uint,
	)
	ProcessAdditionalCharacterSlot(
		slot uint,
	)
	PostprocessAdditionalCharacter(
		additionalCharacter ast.AdditionalCharacterLike,
		index uint,
		count uint,
	)
	PreprocessAdditionalRepetition(
		additionalRepetition ast.AdditionalRepetitionLike,
		index uint,
		count uint,
	)
	ProcessAdditionalRepetitionSlot(
		slot uint,
	)
	PostprocessAdditionalRepetition(
		additionalRepetition ast.AdditionalRepetitionLike,
		index uint,
		count uint,
	)
	PreprocessAllowedCharacters(
		allowedCharacters ast.AllowedCharactersLike,
		index uint,
		count uint,
	)
	ProcessAllowedCharactersSlot(
		slot uint,
	)
	PostprocessAllowedCharacters(
		allowedCharacters ast.AllowedCharactersLike,
		index uint,
		count uint,
	)
	PreprocessAlternativeSequence(
		alternativeSequence ast.AlternativeSequenceLike,
		index uint,
		count uint,
	)
	ProcessAlternativeSequenceSlot(
		slot uint,
	)
	PostprocessAlternativeSequence(
		alternativeSequence ast.AlternativeSequenceLike,
		index uint,
		count uint,
	)
	PreprocessAlternatives(
		alternatives ast.AlternativesLike,
		index uint,
		count uint,
	)
	ProcessAlternativesSlot(
		slot uint,
	)
	PostprocessAlternatives(
		alternatives ast.AlternativesLike,
		index uint,
		count uint,
	)
	PreprocessCardinality(
		cardinality ast.CardinalityLike,
		index uint,
		count uint,
	)
	ProcessCardinalitySlot(
		slot uint,
	)
	PostprocessCardinality(
		cardinality ast.CardinalityLike,
		index uint,
		count uint,
	)
	PreprocessCharacter(
		character ast.CharacterLike,
		index uint,
		count uint,
	)
	ProcessCharacterSlot(
		slot uint,
	)
	PostprocessCharacter(
		character ast.CharacterLike,
		index uint,
		count uint,
	)
	PreprocessComponent(
		component ast.ComponentLike,
		index uint,
		count uint,
	)
	ProcessComponentSlot(
		slot uint,
	)
	PostprocessComponent(
		component ast.ComponentLike,
		index uint,
		count uint,
	)
	PreprocessConstrained(
		constrained ast.ConstrainedLike,
		index uint,
		count uint,
	)
	ProcessConstrainedSlot(
		slot uint,
	)
	PostprocessConstrained(
		constrained ast.ConstrainedLike,
		index uint,
		count uint,
	)
	PreprocessDefinition(
		definition ast.DefinitionLike,
		index uint,
		count uint,
	)
	ProcessDefinitionSlot(
		slot uint,
	)
	PostprocessDefinition(
		definition ast.DefinitionLike,
		index uint,
		count uint,
	)
	PreprocessElement(
		element ast.ElementLike,
		index uint,
		count uint,
	)
	ProcessElementSlot(
		slot uint,
	)
	PostprocessElement(
		element ast.ElementLike,
		index uint,
		count uint,
	)
	PreprocessExplicit(
		explicit ast.ExplicitLike,
		index uint,
		count uint,
	)
	ProcessExplicitSlot(
		slot uint,
	)
	PostprocessExplicit(
		explicit ast.ExplicitLike,
		index uint,
		count uint,
	)
	PreprocessExpression(
		expression ast.ExpressionLike,
		index uint,
		count uint,
	)
	ProcessExpressionSlot(
		slot uint,
	)
	PostprocessExpression(
		expression ast.ExpressionLike,
		index uint,
		count uint,
	)
	PreprocessExtent(
		extent ast.ExtentLike,
		index uint,
		count uint,
	)
	ProcessExtentSlot(
		slot uint,
	)
	PostprocessExtent(
		extent ast.ExtentLike,
		index uint,
		count uint,
	)
	PreprocessFilter(
		filter ast.FilterLike,
		index uint,
		count uint,
	)
	ProcessFilterSlot(
		slot uint,
	)
	PostprocessFilter(
		filter ast.FilterLike,
		index uint,
		count uint,
	)
	PreprocessGroup(
		group ast.GroupLike,
		index uint,
		count uint,
	)
	ProcessGroupSlot(
		slot uint,
	)
	PostprocessGroup(
		group ast.GroupLike,
		index uint,
		count uint,
	)
	PreprocessImplicit(
		implicit ast.ImplicitLike,
		index uint,
		count uint,
	)
	ProcessImplicitSlot(
		slot uint,
	)
	PostprocessImplicit(
		implicit ast.ImplicitLike,
		index uint,
		count uint,
	)
	PreprocessLimit(
		limit ast.LimitLike,
		index uint,
		count uint,
	)
	ProcessLimitSlot(
		slot uint,
	)
	PostprocessLimit(
		limit ast.LimitLike,
		index uint,
		count uint,
	)
	PreprocessLiteralValue(
		literalValue ast.LiteralValueLike,
		index uint,
		count uint,
	)
	ProcessLiteralValueSlot(
		slot uint,
	)
	PostprocessLiteralValue(
		literalValue ast.LiteralValueLike,
		index uint,
		count uint,
	)
	PreprocessLiteralValueAlternatives(
		literalValueAlternatives ast.LiteralValueAlternativesLike,
		index uint,
		count uint,
	)
	ProcessLiteralValueAlternativesSlot(
		slot uint,
	)
	PostprocessLiteralValueAlternatives(
		literalValueAlternatives ast.LiteralValueAlternativesLike,
		index uint,
		count uint,
	)
	PreprocessNotice(
		notice ast.NoticeLike,
		index uint,
		count uint,
	)
	ProcessNoticeSlot(
		slot uint,
	)
	PostprocessNotice(
		notice ast.NoticeLike,
		index uint,
		count uint,
	)
	PreprocessPattern(
		pattern ast.PatternLike,
		index uint,
		count uint,
	)
	ProcessPatternSlot(
		slot uint,
	)
	PostprocessPattern(
		pattern ast.PatternLike,
		index uint,
		count uint,
	)
	PreprocessQuantified(
		quantified ast.QuantifiedLike,
		index uint,
		count uint,
	)
	ProcessQuantifiedSlot(
		slot uint,
	)
	PostprocessQuantified(
		quantified ast.QuantifiedLike,
		index uint,
		count uint,
	)
	PreprocessRepetition(
		repetition ast.RepetitionLike,
		index uint,
		count uint,
	)
	ProcessRepetitionSlot(
		slot uint,
	)
	PostprocessRepetition(
		repetition ast.RepetitionLike,
		index uint,
		count uint,
	)
	PreprocessRule(
		rule ast.RuleLike,
		index uint,
		count uint,
	)
	ProcessRuleSlot(
		slot uint,
	)
	PostprocessRule(
		rule ast.RuleLike,
		index uint,
		count uint,
	)
	PreprocessRuleName(
		ruleName ast.RuleNameLike,
		index uint,
		count uint,
	)
	ProcessRuleNameSlot(
		slot uint,
	)
	PostprocessRuleName(
		ruleName ast.RuleNameLike,
		index uint,
		count uint,
	)
	PreprocessRuleNameAlternatives(
		ruleNameAlternatives ast.RuleNameAlternativesLike,
		index uint,
		count uint,
	)
	ProcessRuleNameAlternativesSlot(
		slot uint,
	)
	PostprocessRuleNameAlternatives(
		ruleNameAlternatives ast.RuleNameAlternativesLike,
		index uint,
		count uint,
	)
	PreprocessRuleTerm(
		ruleTerm ast.RuleTermLike,
		index uint,
		count uint,
	)
	ProcessRuleTermSlot(
		slot uint,
	)
	PostprocessRuleTerm(
		ruleTerm ast.RuleTermLike,
		index uint,
		count uint,
	)
	PreprocessRuleTermSequence(
		ruleTermSequence ast.RuleTermSequenceLike,
		index uint,
		count uint,
	)
	ProcessRuleTermSequenceSlot(
		slot uint,
	)
	PostprocessRuleTermSequence(
		ruleTermSequence ast.RuleTermSequenceLike,
		index uint,
		count uint,
	)
	PreprocessSequence(
		sequence ast.SequenceLike,
		index uint,
		count uint,
	)
	ProcessSequenceSlot(
		slot uint,
	)
	PostprocessSequence(
		sequence ast.SequenceLike,
		index uint,
		count uint,
	)
	PreprocessSyntax(
		syntax ast.SyntaxLike,
		index uint,
		count uint,
	)
	ProcessSyntaxSlot(
		slot uint,
	)
	PostprocessSyntax(
		syntax ast.SyntaxLike,
		index uint,
		count uint,
	)
	PreprocessText(
		text ast.TextLike,
		index uint,
		count uint,
	)
	ProcessTextSlot(
		slot uint,
	)
	PostprocessText(
		text ast.TextLike,
		index uint,
		count uint,
	)
	PreprocessTokenName(
		tokenName ast.TokenNameLike,
		index uint,
		count uint,
	)
	ProcessTokenNameSlot(
		slot uint,
	)
	PostprocessTokenName(
		tokenName ast.TokenNameLike,
		index uint,
		count uint,
	)
	PreprocessTokenNameAlternatives(
		tokenNameAlternatives ast.TokenNameAlternativesLike,
		index uint,
		count uint,
	)
	ProcessTokenNameAlternativesSlot(
		slot uint,
	)
	PostprocessTokenNameAlternatives(
		tokenNameAlternatives ast.TokenNameAlternativesLike,
		index uint,
		count uint,
	)
}
