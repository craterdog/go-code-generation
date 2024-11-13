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
	CharacterLike = ast.CharacterLike
	ConstrainedLike = ast.ConstrainedLike
	DefinitionLike = ast.DefinitionLike
	ElementLike = ast.ElementLike
	ExplicitLike = ast.ExplicitLike
	ExpressionLike = ast.ExpressionLike
	ExtentLike = ast.ExtentLike
	FilterLike = ast.FilterLike
	GroupLike = ast.GroupLike
	IdentifierLike = ast.IdentifierLike
	InlineLike = ast.InlineLike
	LimitLike = ast.LimitLike
	LineLike = ast.LineLike
	MultilineLike = ast.MultilineLike
	NoticeLike = ast.NoticeLike
	OptionLike = ast.OptionLike
	PatternLike = ast.PatternLike
	QuantifiedLike = ast.QuantifiedLike
	ReferenceLike = ast.ReferenceLike
	RepetitionLike = ast.RepetitionLike
	RuleLike = ast.RuleLike
	SyntaxLike = ast.SyntaxLike
	TermLike = ast.TermLike
	TextLike = ast.TextLike
)

// Grammar

type (
	FormatterLike = gra.FormatterLike
	ParserLike = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike = gra.ScannerLike
	TokenLike = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike = gra.VisitorLike
)

// UNIVERSAL CONSTRUCTORS

// Ast

func Alternative(arguments ...any) AlternativeLike {
	// Initialize the possible arguments.
	var option OptionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case OptionLike:
			option = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Alternative constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var alternative AlternativeLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Alternative constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return alternative
}

func Cardinality(arguments ...any) CardinalityLike {
	// Initialize the possible arguments.
	var any_ any

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			any_ = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Cardinality constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var cardinality CardinalityLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Cardinality constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return cardinality
}

func Character(arguments ...any) CharacterLike {
	// Initialize the possible arguments.
	var any_ any

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			any_ = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Character constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var character CharacterLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Character constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return character
}

func Constrained(arguments ...any) ConstrainedLike {
	// Initialize the possible arguments.
	var any_ any

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			any_ = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Constrained constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var constrained ConstrainedLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Constrained constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return constrained
}

func Definition(arguments ...any) DefinitionLike {
	// Initialize the possible arguments.
	var any_ any

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			any_ = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Definition constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var definition DefinitionLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Definition constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return definition
}

func Element(arguments ...any) ElementLike {
	// Initialize the possible arguments.
	var any_ any

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			any_ = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Element constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var element ElementLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Element constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return element
}

func Explicit(arguments ...any) ExplicitLike {
	// Initialize the possible arguments.
	var glyph string
	var optionalExtent ExtentLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			glyph = actual
		case ExtentLike:
			optionalExtent = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Explicit constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var explicit ExplicitLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Explicit constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return explicit
}

func Expression(arguments ...any) ExpressionLike {
	// Initialize the possible arguments.
	var lowercase string
	var pattern PatternLike
	var optionalNote string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			lowercase = actual
		case PatternLike:
			pattern = actual
		case string:
			optionalNote = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Expression constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var expression ExpressionLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Expression constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return expression
}

func Extent(arguments ...any) ExtentLike {
	// Initialize the possible arguments.
	var glyph string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			glyph = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Extent constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var extent ExtentLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Extent constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return extent
}

func Filter(arguments ...any) FilterLike {
	// Initialize the possible arguments.
	var optionalExcluded string
	var characters abs.Sequential[CharacterLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			optionalExcluded = actual
		case abs.Sequential[CharacterLike]:
			characters = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Filter constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var filter FilterLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Filter constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return filter
}

func Group(arguments ...any) GroupLike {
	// Initialize the possible arguments.
	var pattern PatternLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case PatternLike:
			pattern = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Group constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var group GroupLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Group constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return group
}

func Identifier(arguments ...any) IdentifierLike {
	// Initialize the possible arguments.
	var any_ any

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			any_ = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Identifier constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var identifier IdentifierLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Identifier constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return identifier
}

func Inline(arguments ...any) InlineLike {
	// Initialize the possible arguments.
	var terms abs.Sequential[TermLike]
	var optionalNote string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[TermLike]:
			terms = actual
		case string:
			optionalNote = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Inline constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var inline InlineLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Inline constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return inline
}

func Limit(arguments ...any) LimitLike {
	// Initialize the possible arguments.
	var optionalNumber string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			optionalNumber = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Limit constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var limit LimitLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Limit constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return limit
}

func Line(arguments ...any) LineLike {
	// Initialize the possible arguments.
	var identifier IdentifierLike
	var optionalNote string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case IdentifierLike:
			identifier = actual
		case string:
			optionalNote = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Line constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var line LineLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Line constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return line
}

func Multiline(arguments ...any) MultilineLike {
	// Initialize the possible arguments.
	var lines abs.Sequential[LineLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[LineLike]:
			lines = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Multiline constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var multiline MultilineLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Multiline constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return multiline
}

func Notice(arguments ...any) NoticeLike {
	// Initialize the possible arguments.
	var comment string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			comment = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Notice constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var notice NoticeLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Notice constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return notice
}

func Option(arguments ...any) OptionLike {
	// Initialize the possible arguments.
	var repetitions abs.Sequential[RepetitionLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case abs.Sequential[RepetitionLike]:
			repetitions = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Option constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var option OptionLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Option constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return option
}

func Pattern(arguments ...any) PatternLike {
	// Initialize the possible arguments.
	var option OptionLike
	var alternatives abs.Sequential[AlternativeLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case OptionLike:
			option = actual
		case abs.Sequential[AlternativeLike]:
			alternatives = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Pattern constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var pattern PatternLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Pattern constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return pattern
}

func Quantified(arguments ...any) QuantifiedLike {
	// Initialize the possible arguments.
	var number string
	var optionalLimit LimitLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			number = actual
		case LimitLike:
			optionalLimit = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Quantified constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var quantified QuantifiedLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Quantified constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return quantified
}

func Reference(arguments ...any) ReferenceLike {
	// Initialize the possible arguments.
	var identifier IdentifierLike
	var optionalCardinality CardinalityLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case IdentifierLike:
			identifier = actual
		case CardinalityLike:
			optionalCardinality = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Reference constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var reference ReferenceLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Reference constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return reference
}

func Repetition(arguments ...any) RepetitionLike {
	// Initialize the possible arguments.
	var element ElementLike
	var optionalCardinality CardinalityLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case ElementLike:
			element = actual
		case CardinalityLike:
			optionalCardinality = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Repetition constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var repetition RepetitionLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Repetition constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return repetition
}

func Rule(arguments ...any) RuleLike {
	// Initialize the possible arguments.
	var uppercase string
	var definition DefinitionLike

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			uppercase = actual
		case DefinitionLike:
			definition = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Rule constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var rule RuleLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Rule constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return rule
}

func Syntax(arguments ...any) SyntaxLike {
	// Initialize the possible arguments.
	var notice NoticeLike
	var comment1 string
	var rules abs.Sequential[RuleLike]
	var comment2 string
	var expressions abs.Sequential[ExpressionLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case NoticeLike:
			notice = actual
		case string:
			comment1 = actual
		case abs.Sequential[RuleLike]:
			rules = actual
		case string:
			comment2 = actual
		case abs.Sequential[ExpressionLike]:
			expressions = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Syntax constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var syntax SyntaxLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Syntax constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return syntax
}

func Term(arguments ...any) TermLike {
	// Initialize the possible arguments.
	var any_ any

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			any_ = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Term constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var term TermLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Term constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return term
}

func Text(arguments ...any) TextLike {
	// Initialize the possible arguments.
	var any_ any

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case any:
			any_ = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Text constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var text TextLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Text constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return text
}

// Grammar

func Formatter(arguments ...any) FormatterLike {
	// Initialize the possible arguments.

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Formatter constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var formatter FormatterLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Formatter constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return formatter
}

func Parser(arguments ...any) ParserLike {
	// Initialize the possible arguments.

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Parser constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var parser ParserLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Parser constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return parser
}

func Processor(arguments ...any) ProcessorLike {
	// Initialize the possible arguments.

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Processor constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var processor ProcessorLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Processor constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return processor
}

func Scanner(arguments ...any) ScannerLike {
	// Initialize the possible arguments.
	var source string
	var tokens abs.QueueLike[TokenLike]

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case string:
			source = actual
		case abs.QueueLike[TokenLike]:
			tokens = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Scanner constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var scanner ScannerLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Scanner constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return scanner
}

func Token(arguments ...any) TokenLike {
	// Initialize the possible arguments.
	var line uint
	var position uint
	var type_ TokenType
	var value string

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case uint:
			line = actual
		case uint:
			position = actual
		case TokenType:
			type_ = actual
		case string:
			value = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Token constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var token TokenLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Token constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return token
}

func Validator(arguments ...any) ValidatorLike {
	// Initialize the possible arguments.

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Validator constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var validator ValidatorLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Validator constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return validator
}

func Visitor(arguments ...any) VisitorLike {
	// Initialize the possible arguments.
	var processor Methodical

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {
		case Methodical:
			processor = actual
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Visitor constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the appropriate constructor.
	var visitor VisitorLike
	switch {
		default:
			var message = fmt.Sprintf(
				"A Visitor constructor matching the arguments was not found: $v\n",
				arguments,
			)
			panic(message)
	}
	return visitor
}