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

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package grammar

import (
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	ast "github.com/craterdog/go-syntax-notation/v5/ast"
)

// Type Declarations

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
	LiteralToken
	LowercaseToken
	NewlineToken
	NoteToken
	NumberToken
	OptionalToken
	RepeatedToken
	SpaceToken
	UppercaseToken
)

// Class Declarations

/*
AnalyzerClassLike defines the set of class constants, constructors and
functions that must be supported by all analyzer-class-like classes.
*/
type AnalyzerClassLike interface {
	// Constructor Methods
	Make(
		syntax ast.SyntaxLike,
	) AnalyzerLike
}

/*
FormatterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructor Methods
	Make() FormatterLike
}

/*
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor Methods
	Make() ParserLike
}

/*
ProcessorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor Methods
	Make() ProcessorLike
}

/*
ScannerClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructor Methods
	Make(
		source string,
		tokens abs.QueueLike[TokenLike],
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
TokenClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete token-like class.
*/
type TokenClassLike interface {
	// Constructor Methods
	Make(
		line uint,
		position uint,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructor Methods
	Make() ValidatorLike
}

/*
VisitorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor Methods
	Make(
		processor Methodical,
	) VisitorLike
}

// Instance Declarations

/*
AnalyzerLike defines the set of aspects and methods that must be supported by
all analyzer-like instances.
*/
type AnalyzerLike interface {
	// Primary Methods
	GetClass() AnalyzerClassLike
	GetExpressions() abs.Sequential[abs.AssociationLike[string, string]]
	GetIdentifiers(
		ruleName string,
	) abs.Sequential[ast.IdentifierLike]
	GetNotice() string
	GetReferences(
		ruleName string,
	) abs.Sequential[ast.ReferenceLike]
	GetRuleNames() abs.Sequential[string]
	GetSyntaxMap() string
	GetSyntaxName() string
	GetTerms(
		ruleName string,
	) abs.Sequential[ast.TermLike]
	GetTokenNames() abs.Sequential[string]
	GetVariableType(
		reference ast.ReferenceLike,
	) string
	GetVariables(
		ruleName string,
	) abs.Sequential[string]
	HasPlurals() bool
	IsDelimited(
		ruleName string,
	) bool
	IsPlural(
		name string,
	) bool

	// Aspect Interfaces
	Methodical
}

/*
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Primary Methods
	GetClass() FormatterClassLike
	FormatSyntax(
		syntax ast.SyntaxLike,
	) string

	// Aspect Interfaces
	Methodical
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Primary Methods
	GetClass() ParserClassLike
	ParseSource(
		source string,
	) ast.SyntaxLike
}

/*
ProcessorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Primary Methods
	GetClass() ProcessorClassLike

	// Aspect Interfaces
	Methodical
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Primary Methods
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Primary Methods
	GetClass() TokenClassLike

	// Attribute Methods
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Primary Methods
	GetClass() ValidatorClassLike
	ValidateSyntax(
		syntax ast.SyntaxLike,
	)

	// Aspect Interfaces
	Methodical
}

/*
VisitorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Primary Methods
	GetClass() VisitorClassLike
	VisitSyntax(
		syntax ast.SyntaxLike,
	)
}

// Aspect Declarations

/*
Methodical defines the set of method signatures that must be supported
by all methodical processors.
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
	ProcessOptional(
		optional string,
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
	PreprocessLine(
		line ast.LineLike,
		index uint,
		size uint,
	)
	ProcessLineSlot(
		slot uint,
	)
	PostprocessLine(
		line ast.LineLike,
		index uint,
		size uint,
	)
	PreprocessMultiline(
		multiline ast.MultilineLike,
	)
	ProcessMultilineSlot(
		slot uint,
	)
	PostprocessMultiline(
		multiline ast.MultilineLike,
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
