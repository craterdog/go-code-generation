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
│              This "Package.go" file was automatically generated.             │
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
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package grammar

import (
	col "github.com/craterdog/go-collection-framework/v5/collection"
	ast "github.com/craterdog/go-syntax-notation/v6/ast"
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
	ExcludedToken
	GlyphToken
	IntrinsicToken
	LowercaseToken
	NewlineToken
	NoteToken
	NumberToken
	OptionalToken
	QuoteToken
	RepeatedToken
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
	ProcessExcluded(
		excluded string,
	)
	ProcessGlyph(
		glyph string,
	)
	ProcessIntrinsic(
		intrinsic string,
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
	ProcessOptional(
		optional string,
	)
	ProcessQuote(
		quote string,
	)
	ProcessRepeated(
		repeated string,
	)
	ProcessSpace(
		space string,
	)
	ProcessUppercase(
		uppercase string,
	)
	PreprocessAlternative(
		alternative ast.AlternativeLike,
		index uint,
		size uint,
	)
	ProcessAlternativeSlot(
		slot uint,
	)
	PostprocessAlternative(
		alternative ast.AlternativeLike,
		index uint,
		size uint,
	)
	PreprocessCardinality(
		cardinality ast.CardinalityLike,
	)
	ProcessCardinalitySlot(
		slot uint,
	)
	PostprocessCardinality(
		cardinality ast.CardinalityLike,
	)
	PreprocessCharacter(
		character ast.CharacterLike,
		index uint,
		size uint,
	)
	ProcessCharacterSlot(
		slot uint,
	)
	PostprocessCharacter(
		character ast.CharacterLike,
		index uint,
		size uint,
	)
	PreprocessConstrained(
		constrained ast.ConstrainedLike,
	)
	ProcessConstrainedSlot(
		slot uint,
	)
	PostprocessConstrained(
		constrained ast.ConstrainedLike,
	)
	PreprocessDefinition(
		definition ast.DefinitionLike,
	)
	ProcessDefinitionSlot(
		slot uint,
	)
	PostprocessDefinition(
		definition ast.DefinitionLike,
	)
	PreprocessElement(
		element ast.ElementLike,
	)
	ProcessElementSlot(
		slot uint,
	)
	PostprocessElement(
		element ast.ElementLike,
	)
	PreprocessExplicit(
		explicit ast.ExplicitLike,
	)
	ProcessExplicitSlot(
		slot uint,
	)
	PostprocessExplicit(
		explicit ast.ExplicitLike,
	)
	PreprocessExpression(
		expression ast.ExpressionLike,
		index uint,
		size uint,
	)
	ProcessExpressionSlot(
		slot uint,
	)
	PostprocessExpression(
		expression ast.ExpressionLike,
		index uint,
		size uint,
	)
	PreprocessExpressionOption(
		expressionOption ast.ExpressionOptionLike,
		index uint,
		size uint,
	)
	ProcessExpressionOptionSlot(
		slot uint,
	)
	PostprocessExpressionOption(
		expressionOption ast.ExpressionOptionLike,
		index uint,
		size uint,
	)
	PreprocessExtent(
		extent ast.ExtentLike,
	)
	ProcessExtentSlot(
		slot uint,
	)
	PostprocessExtent(
		extent ast.ExtentLike,
	)
	PreprocessFilter(
		filter ast.FilterLike,
	)
	ProcessFilterSlot(
		slot uint,
	)
	PostprocessFilter(
		filter ast.FilterLike,
	)
	PreprocessGroup(
		group ast.GroupLike,
	)
	ProcessGroupSlot(
		slot uint,
	)
	PostprocessGroup(
		group ast.GroupLike,
	)
	PreprocessIdentifier(
		identifier ast.IdentifierLike,
	)
	ProcessIdentifierSlot(
		slot uint,
	)
	PostprocessIdentifier(
		identifier ast.IdentifierLike,
	)
	PreprocessImplicit(
		implicit ast.ImplicitLike,
	)
	ProcessImplicitSlot(
		slot uint,
	)
	PostprocessImplicit(
		implicit ast.ImplicitLike,
	)
	PreprocessInline(
		inline ast.InlineLike,
	)
	ProcessInlineSlot(
		slot uint,
	)
	PostprocessInline(
		inline ast.InlineLike,
	)
	PreprocessLimit(
		limit ast.LimitLike,
	)
	ProcessLimitSlot(
		slot uint,
	)
	PostprocessLimit(
		limit ast.LimitLike,
	)
	PreprocessLiteral(
		literal ast.LiteralLike,
	)
	ProcessLiteralSlot(
		slot uint,
	)
	PostprocessLiteral(
		literal ast.LiteralLike,
	)
	PreprocessMultiexpression(
		multiexpression ast.MultiexpressionLike,
	)
	ProcessMultiexpressionSlot(
		slot uint,
	)
	PostprocessMultiexpression(
		multiexpression ast.MultiexpressionLike,
	)
	PreprocessMultirule(
		multirule ast.MultiruleLike,
	)
	ProcessMultiruleSlot(
		slot uint,
	)
	PostprocessMultirule(
		multirule ast.MultiruleLike,
	)
	PreprocessNotice(
		notice ast.NoticeLike,
	)
	ProcessNoticeSlot(
		slot uint,
	)
	PostprocessNotice(
		notice ast.NoticeLike,
	)
	PreprocessOption(
		option ast.OptionLike,
	)
	ProcessOptionSlot(
		slot uint,
	)
	PostprocessOption(
		option ast.OptionLike,
	)
	PreprocessPattern(
		pattern ast.PatternLike,
	)
	ProcessPatternSlot(
		slot uint,
	)
	PostprocessPattern(
		pattern ast.PatternLike,
	)
	PreprocessQuantified(
		quantified ast.QuantifiedLike,
	)
	ProcessQuantifiedSlot(
		slot uint,
	)
	PostprocessQuantified(
		quantified ast.QuantifiedLike,
	)
	PreprocessReference(
		reference ast.ReferenceLike,
	)
	ProcessReferenceSlot(
		slot uint,
	)
	PostprocessReference(
		reference ast.ReferenceLike,
	)
	PreprocessRepetition(
		repetition ast.RepetitionLike,
		index uint,
		size uint,
	)
	ProcessRepetitionSlot(
		slot uint,
	)
	PostprocessRepetition(
		repetition ast.RepetitionLike,
		index uint,
		size uint,
	)
	PreprocessRule(
		rule ast.RuleLike,
		index uint,
		size uint,
	)
	ProcessRuleSlot(
		slot uint,
	)
	PostprocessRule(
		rule ast.RuleLike,
		index uint,
		size uint,
	)
	PreprocessRuleOption(
		ruleOption ast.RuleOptionLike,
		index uint,
		size uint,
	)
	ProcessRuleOptionSlot(
		slot uint,
	)
	PostprocessRuleOption(
		ruleOption ast.RuleOptionLike,
		index uint,
		size uint,
	)
	PreprocessSyntax(
		syntax ast.SyntaxLike,
	)
	ProcessSyntaxSlot(
		slot uint,
	)
	PostprocessSyntax(
		syntax ast.SyntaxLike,
	)
	PreprocessTerm(
		term ast.TermLike,
		index uint,
		size uint,
	)
	ProcessTermSlot(
		slot uint,
	)
	PostprocessTerm(
		term ast.TermLike,
		index uint,
		size uint,
	)
	PreprocessText(
		text ast.TextLike,
	)
	ProcessTextSlot(
		slot uint,
	)
	PostprocessText(
		text ast.TextLike,
	)
}
