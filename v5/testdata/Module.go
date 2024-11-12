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
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Alternative constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var alternative = ast.Alternative().Make(<Arguments>)
	return alternative
}

func Cardinality(arguments ...any) CardinalityLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Cardinality constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var cardinality = ast.Cardinality().Make(<Arguments>)
	return cardinality
}

func Character(arguments ...any) CharacterLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Character constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var character = ast.Character().Make(<Arguments>)
	return character
}

func Constrained(arguments ...any) ConstrainedLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Constrained constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var constrained = ast.Constrained().Make(<Arguments>)
	return constrained
}

func Definition(arguments ...any) DefinitionLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Definition constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var definition = ast.Definition().Make(<Arguments>)
	return definition
}

func Element(arguments ...any) ElementLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Element constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var element = ast.Element().Make(<Arguments>)
	return element
}

func Explicit(arguments ...any) ExplicitLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Explicit constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var explicit = ast.Explicit().Make(<Arguments>)
	return explicit
}

func Expression(arguments ...any) ExpressionLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Expression constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var expression = ast.Expression().Make(<Arguments>)
	return expression
}

func Extent(arguments ...any) ExtentLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Extent constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var extent = ast.Extent().Make(<Arguments>)
	return extent
}

func Filter(arguments ...any) FilterLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Filter constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var filter = ast.Filter().Make(<Arguments>)
	return filter
}

func Group(arguments ...any) GroupLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Group constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var group = ast.Group().Make(<Arguments>)
	return group
}

func Identifier(arguments ...any) IdentifierLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Identifier constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var identifier = ast.Identifier().Make(<Arguments>)
	return identifier
}

func Inline(arguments ...any) InlineLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Inline constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var inline = ast.Inline().Make(<Arguments>)
	return inline
}

func Limit(arguments ...any) LimitLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Limit constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var limit = ast.Limit().Make(<Arguments>)
	return limit
}

func Line(arguments ...any) LineLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Line constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var line = ast.Line().Make(<Arguments>)
	return line
}

func Multiline(arguments ...any) MultilineLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Multiline constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var multiline = ast.Multiline().Make(<Arguments>)
	return multiline
}

func Notice(arguments ...any) NoticeLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Notice constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var notice = ast.Notice().Make(<Arguments>)
	return notice
}

func Option(arguments ...any) OptionLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Option constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var option = ast.Option().Make(<Arguments>)
	return option
}

func Pattern(arguments ...any) PatternLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Pattern constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var pattern = ast.Pattern().Make(<Arguments>)
	return pattern
}

func Quantified(arguments ...any) QuantifiedLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Quantified constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var quantified = ast.Quantified().Make(<Arguments>)
	return quantified
}

func Reference(arguments ...any) ReferenceLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Reference constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var reference = ast.Reference().Make(<Arguments>)
	return reference
}

func Repetition(arguments ...any) RepetitionLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Repetition constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var repetition = ast.Repetition().Make(<Arguments>)
	return repetition
}

func Rule(arguments ...any) RuleLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Rule constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var rule = ast.Rule().Make(<Arguments>)
	return rule
}

func Syntax(arguments ...any) SyntaxLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Syntax constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var syntax = ast.Syntax().Make(<Arguments>)
	return syntax
}

func Term(arguments ...any) TermLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Term constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var term = ast.Term().Make(<Arguments>)
	return term
}

func Text(arguments ...any) TextLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Text constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var text = ast.Text().Make(<Arguments>)
	return text
}

// Grammar

func Formatter(arguments ...any) FormatterLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Formatter constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var formatter = gra.Formatter().Make(<Arguments>)
	return formatter
}

func Parser(arguments ...any) ParserLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Parser constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var parser = gra.Parser().Make(<Arguments>)
	return parser
}

func Processor(arguments ...any) ProcessorLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Processor constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var processor = gra.Processor().Make(<Arguments>)
	return processor
}

func Scanner(arguments ...any) ScannerLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Scanner constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var scanner = gra.Scanner().Make(<Arguments>)
	return scanner
}

func Token(arguments ...any) TokenLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Token constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var token = gra.Token().Make(<Arguments>)
	return token
}

func Validator(arguments ...any) ValidatorLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Validator constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var validator = gra.Validator().Make(<Arguments>)
	return validator
}

func Visitor(arguments ...any) VisitorLike {
	// Initialize the possible arguments.<ArgumentInitialization>

	// Process the actual arguments.
	for _, argument := range arguments {
		switch actual := argument.(type) {<ValidationCases>
		default:
			var message = fmt.Sprintf(
				"An unknown argument type was passed into the Visitor constructor: %T\n",
				actual,
			)
			panic(message)
		}
	}

	// Call the constructor.
	var visitor = gra.Visitor().Make(<Arguments>)
	return visitor
}