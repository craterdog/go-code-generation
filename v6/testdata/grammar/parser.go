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
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	fra "github.com/craterdog/go-collection-framework/v5"
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	ast "github.com/craterdog/go-syntax-notation/v6/ast"
	mat "math"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ParserClass() ParserClassLike {
	return parserClass()
}

// Constructor Methods

func (c *parserClass_) Parser() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserClass()
}

func (v *parser_) ParseSource(
	source string,
) ast.SyntaxLike {
	v.source_ = sts.ReplaceAll(source, "\t", "    ")
	v.tokens_ = fra.Queue[TokenLike]()
	v.next_ = fra.Stack[TokenLike]()

	// The scanner runs in a separate Go routine.
	ScannerClass().Scanner(v.source_, v.tokens_)

	// Attempt to parse the syntax.
	var syntax, token, ok = v.parseSyntax()
	if !ok || !v.tokens_.IsEmpty() {
		var message = v.formatError("$Syntax", token)
		panic(message)
	}
	return syntax
}

// PROTECTED INTERFACE

// Private Methods

func (v *parser_) parseAlternative() (
	alternative ast.AlternativeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "|" delimiter.
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Alternative rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Alternative", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Option rule.
	var option ast.OptionLike
	option, token, ok = v.parseOption()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Option rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Alternative", token)
		panic(message)
	}

	// Found a single Alternative rule.
	ok = true
	v.remove(tokens)
	alternative = ast.AlternativeClass().Alternative(option)
	return
}

func (v *parser_) parseCardinality() (
	cardinality ast.CardinalityLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Constrained Cardinality.
	var constrained ast.ConstrainedLike
	constrained, token, ok = v.parseConstrained()
	if ok {
		// Found a single Constrained Cardinality.
		cardinality = ast.CardinalityClass().Cardinality(constrained)
		return
	}

	// Attempt to parse a single Quantified Cardinality.
	var quantified ast.QuantifiedLike
	quantified, token, ok = v.parseQuantified()
	if ok {
		// Found a single Quantified Cardinality.
		cardinality = ast.CardinalityClass().Cardinality(quantified)
		return
	}

	// This is not a single Cardinality rule.
	return
}

func (v *parser_) parseCharacter() (
	character ast.CharacterLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Implicit Character.
	var implicit ast.ImplicitLike
	implicit, token, ok = v.parseImplicit()
	if ok {
		// Found a single Implicit Character.
		character = ast.CharacterClass().Character(implicit)
		return
	}

	// Attempt to parse a single Explicit Character.
	var explicit ast.ExplicitLike
	explicit, token, ok = v.parseExplicit()
	if ok {
		// Found a single Explicit Character.
		character = ast.CharacterClass().Character(explicit)
		return
	}

	// This is not a single Character rule.
	return
}

func (v *parser_) parseConstrained() (
	constrained ast.ConstrainedLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single optional Constrained.
	var optional string
	optional, token, ok = v.parseToken(OptionalToken)
	if ok {
		// Found a single optional Constrained.
		constrained = ast.ConstrainedClass().Constrained(optional)
		return
	}

	// Attempt to parse a single repeated Constrained.
	var repeated string
	repeated, token, ok = v.parseToken(RepeatedToken)
	if ok {
		// Found a single repeated Constrained.
		constrained = ast.ConstrainedClass().Constrained(repeated)
		return
	}

	// This is not a single Constrained rule.
	return
}

func (v *parser_) parseDefinition() (
	definition ast.DefinitionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Multirule Definition.
	var multirule ast.MultiruleLike
	multirule, token, ok = v.parseMultirule()
	if ok {
		// Found a single Multirule Definition.
		definition = ast.DefinitionClass().Definition(multirule)
		return
	}

	// Attempt to parse a single Multiexpression Definition.
	var multiexpression ast.MultiexpressionLike
	multiexpression, token, ok = v.parseMultiexpression()
	if ok {
		// Found a single Multiexpression Definition.
		definition = ast.DefinitionClass().Definition(multiexpression)
		return
	}

	// Attempt to parse a single Inline Definition.
	var inline ast.InlineLike
	inline, token, ok = v.parseInline()
	if ok {
		// Found a single Inline Definition.
		definition = ast.DefinitionClass().Definition(inline)
		return
	}

	// This is not a single Definition rule.
	return
}

func (v *parser_) parseElement() (
	element ast.ElementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Group Element.
	var group ast.GroupLike
	group, token, ok = v.parseGroup()
	if ok {
		// Found a single Group Element.
		element = ast.ElementClass().Element(group)
		return
	}

	// Attempt to parse a single Filter Element.
	var filter ast.FilterLike
	filter, token, ok = v.parseFilter()
	if ok {
		// Found a single Filter Element.
		element = ast.ElementClass().Element(filter)
		return
	}

	// Attempt to parse a single Text Element.
	var text ast.TextLike
	text, token, ok = v.parseText()
	if ok {
		// Found a single Text Element.
		element = ast.ElementClass().Element(text)
		return
	}

	// This is not a single Element rule.
	return
}

func (v *parser_) parseExplicit() (
	explicit ast.ExplicitLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single glyph token.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single glyph token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Explicit", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Extent rule.
	var optionalExtent ast.ExtentLike
	optionalExtent, _, ok = v.parseExtent()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Explicit rule.
	ok = true
	v.remove(tokens)
	explicit = ast.ExplicitClass().Explicit(
		glyph,
		optionalExtent,
	)
	return
}

func (v *parser_) parseExpression() (
	expression ast.ExpressionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "$" delimiter.
	_, token, ok = v.parseDelimiter("$")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Expression rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Expression", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single lowercase token.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single lowercase token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Expression", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ":" delimiter.
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Expression rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Expression", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Pattern rule.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Pattern rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Expression", token)
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Expression rule.
	ok = true
	v.remove(tokens)
	expression = ast.ExpressionClass().Expression(
		lowercase,
		pattern,
		optionalNote,
	)
	return
}

func (v *parser_) parseExpressionOption() (
	expressionOption ast.ExpressionOptionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single newline token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExpressionOption", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single lowercase token.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single lowercase token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExpressionOption", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExpressionOption rule.
	ok = true
	v.remove(tokens)
	expressionOption = ast.ExpressionOptionClass().ExpressionOption(
		newline,
		lowercase,
		optionalNote,
	)
	return
}

func (v *parser_) parseExtent() (
	extent ast.ExtentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Extent rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Extent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single glyph token.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single glyph token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Extent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Extent rule.
	ok = true
	v.remove(tokens)
	extent = ast.ExtentClass().Extent(glyph)
	return
}

func (v *parser_) parseFilter() (
	filter ast.FilterLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse an optional excluded token.
	var optionalExcluded string
	optionalExcluded, token, ok = v.parseToken(ExcludedToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Filter rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Filter", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Character rules.
	var characters = fra.List[ast.CharacterLike]()
charactersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var character ast.CharacterLike
		character, token, ok = v.parseCharacter()
		if !ok {
			switch {
			case count >= 1:
				break charactersLoop
			case uti.IsDefined(tokens):
				// This is not multiple Character rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Filter", token)
				message += "1 or more Character rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		characters.AppendValue(character)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Filter rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Filter", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Filter rule.
	ok = true
	v.remove(tokens)
	filter = ast.FilterClass().Filter(
		optionalExcluded,
		characters,
	)
	return
}

func (v *parser_) parseGroup() (
	group ast.GroupLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Group rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Group", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Pattern rule.
	var pattern ast.PatternLike
	pattern, token, ok = v.parsePattern()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Pattern rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Group", token)
		panic(message)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Group rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Group", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Group rule.
	ok = true
	v.remove(tokens)
	group = ast.GroupClass().Group(pattern)
	return
}

func (v *parser_) parseIdentifier() (
	identifier ast.IdentifierLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single lowercase Identifier.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if ok {
		// Found a single lowercase Identifier.
		identifier = ast.IdentifierClass().Identifier(lowercase)
		return
	}

	// Attempt to parse a single uppercase Identifier.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken)
	if ok {
		// Found a single uppercase Identifier.
		identifier = ast.IdentifierClass().Identifier(uppercase)
		return
	}

	// This is not a single Identifier rule.
	return
}

func (v *parser_) parseImplicit() (
	implicit ast.ImplicitLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single intrinsic token.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single intrinsic token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Implicit", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Implicit rule.
	ok = true
	v.remove(tokens)
	implicit = ast.ImplicitClass().Implicit(intrinsic)
	return
}

func (v *parser_) parseInline() (
	inline ast.InlineLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple Term rules.
	var terms = fra.List[ast.TermLike]()
termsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var term ast.TermLike
		term, token, ok = v.parseTerm()
		if !ok {
			switch {
			case count >= 1:
				break termsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Term rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Inline", token)
				message += "1 or more Term rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		terms.AppendValue(term)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Inline rule.
	ok = true
	v.remove(tokens)
	inline = ast.InlineClass().Inline(
		terms,
		optionalNote,
	)
	return
}

func (v *parser_) parseLimit() (
	limit ast.LimitLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Limit rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Limit", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional number token.
	var optionalNumber string
	optionalNumber, token, ok = v.parseToken(NumberToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Limit rule.
	ok = true
	v.remove(tokens)
	limit = ast.LimitClass().Limit(optionalNumber)
	return
}

func (v *parser_) parseLiteral() (
	literal ast.LiteralLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single quote token.
	var quote string
	quote, token, ok = v.parseToken(QuoteToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single quote token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Literal", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Literal rule.
	ok = true
	v.remove(tokens)
	literal = ast.LiteralClass().Literal(quote)
	return
}

func (v *parser_) parseMultiexpression() (
	multiexpression ast.MultiexpressionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple ExpressionOption rules.
	var expressionOptions = fra.List[ast.ExpressionOptionLike]()
expressionOptionsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var expressionOption ast.ExpressionOptionLike
		expressionOption, token, ok = v.parseExpressionOption()
		if !ok {
			switch {
			case count >= 1:
				break expressionOptionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple ExpressionOption rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Multiexpression", token)
				message += "1 or more ExpressionOption rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		expressionOptions.AppendValue(expressionOption)
	}

	// Found a single Multiexpression rule.
	ok = true
	v.remove(tokens)
	multiexpression = ast.MultiexpressionClass().Multiexpression(expressionOptions)
	return
}

func (v *parser_) parseMultirule() (
	multirule ast.MultiruleLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple RuleOption rules.
	var ruleOptions = fra.List[ast.RuleOptionLike]()
ruleOptionsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var ruleOption ast.RuleOptionLike
		ruleOption, token, ok = v.parseRuleOption()
		if !ok {
			switch {
			case count >= 1:
				break ruleOptionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple RuleOption rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Multirule", token)
				message += "1 or more RuleOption rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		ruleOptions.AppendValue(ruleOption)
	}

	// Found a single Multirule rule.
	ok = true
	v.remove(tokens)
	multirule = ast.MultiruleClass().Multirule(ruleOptions)
	return
}

func (v *parser_) parseNotice() (
	notice ast.NoticeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single comment token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Notice", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single newline token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Notice", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Notice rule.
	ok = true
	v.remove(tokens)
	notice = ast.NoticeClass().Notice(
		comment,
		newline,
	)
	return
}

func (v *parser_) parseOption() (
	option ast.OptionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple Repetition rules.
	var repetitions = fra.List[ast.RepetitionLike]()
repetitionsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var repetition ast.RepetitionLike
		repetition, token, ok = v.parseRepetition()
		if !ok {
			switch {
			case count >= 1:
				break repetitionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Repetition rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Option", token)
				message += "1 or more Repetition rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		repetitions.AppendValue(repetition)
	}

	// Found a single Option rule.
	ok = true
	v.remove(tokens)
	option = ast.OptionClass().Option(repetitions)
	return
}

func (v *parser_) parsePattern() (
	pattern ast.PatternLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Option rule.
	var option ast.OptionLike
	option, token, ok = v.parseOption()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Option rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Pattern", token)
		panic(message)
	}

	// Attempt to parse multiple Alternative rules.
	var alternatives = fra.List[ast.AlternativeLike]()
alternativesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var alternative ast.AlternativeLike
		alternative, token, ok = v.parseAlternative()
		if !ok {
			switch {
			case count >= 0:
				break alternativesLoop
			case uti.IsDefined(tokens):
				// This is not multiple Alternative rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Pattern", token)
				message += "0 or more Alternative rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		alternatives.AppendValue(alternative)
	}

	// Found a single Pattern rule.
	ok = true
	v.remove(tokens)
	pattern = ast.PatternClass().Pattern(
		option,
		alternatives,
	)
	return
}

func (v *parser_) parseQuantified() (
	quantified ast.QuantifiedLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Quantified rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Quantified", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single number token.
	var number string
	number, token, ok = v.parseToken(NumberToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single number token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Quantified", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Limit rule.
	var optionalLimit ast.LimitLike
	optionalLimit, _, ok = v.parseLimit()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Quantified rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Quantified", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Quantified rule.
	ok = true
	v.remove(tokens)
	quantified = ast.QuantifiedClass().Quantified(
		number,
		optionalLimit,
	)
	return
}

func (v *parser_) parseReference() (
	reference ast.ReferenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Identifier rule.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Identifier rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Reference", token)
		panic(message)
	}

	// Attempt to parse an optional Cardinality rule.
	var optionalCardinality ast.CardinalityLike
	optionalCardinality, _, ok = v.parseCardinality()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Reference rule.
	ok = true
	v.remove(tokens)
	reference = ast.ReferenceClass().Reference(
		identifier,
		optionalCardinality,
	)
	return
}

func (v *parser_) parseRepetition() (
	repetition ast.RepetitionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Element rule.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Element rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Repetition", token)
		panic(message)
	}

	// Attempt to parse an optional Cardinality rule.
	var optionalCardinality ast.CardinalityLike
	optionalCardinality, _, ok = v.parseCardinality()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Repetition rule.
	ok = true
	v.remove(tokens)
	repetition = ast.RepetitionClass().Repetition(
		element,
		optionalCardinality,
	)
	return
}

func (v *parser_) parseRule() (
	rule ast.RuleLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "$" delimiter.
	_, token, ok = v.parseDelimiter("$")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Rule rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Rule", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single uppercase token.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single uppercase token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Rule", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ":" delimiter.
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Rule rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Rule", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Definition rule.
	var definition ast.DefinitionLike
	definition, token, ok = v.parseDefinition()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Definition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Rule", token)
		panic(message)
	}

	// Found a single Rule rule.
	ok = true
	v.remove(tokens)
	rule = ast.RuleClass().Rule(
		uppercase,
		definition,
	)
	return
}

func (v *parser_) parseRuleOption() (
	ruleOption ast.RuleOptionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single newline token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RuleOption", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single uppercase token.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single uppercase token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RuleOption", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single RuleOption rule.
	ok = true
	v.remove(tokens)
	ruleOption = ast.RuleOptionClass().RuleOption(
		newline,
		uppercase,
		optionalNote,
	)
	return
}

func (v *parser_) parseSyntax() (
	syntax ast.SyntaxLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Notice rule.
	var notice ast.NoticeLike
	notice, token, ok = v.parseNotice()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Notice rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Syntax", token)
		panic(message)
	}

	// Attempt to parse a single comment token.
	var comment1 string
	comment1, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single comment token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Syntax", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Rule rules.
	var rules = fra.List[ast.RuleLike]()
rulesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var rule ast.RuleLike
		rule, token, ok = v.parseRule()
		if !ok {
			switch {
			case count >= 1:
				break rulesLoop
			case uti.IsDefined(tokens):
				// This is not multiple Rule rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Syntax", token)
				message += "1 or more Rule rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		rules.AppendValue(rule)
	}

	// Attempt to parse a single comment token.
	var comment2 string
	comment2, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single comment token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Syntax", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple Expression rules.
	var expressions = fra.List[ast.ExpressionLike]()
expressionsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var expression ast.ExpressionLike
		expression, token, ok = v.parseExpression()
		if !ok {
			switch {
			case count >= 1:
				break expressionsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Expression rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Syntax", token)
				message += "1 or more Expression rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		expressions.AppendValue(expression)
	}

	// Found a single Syntax rule.
	ok = true
	v.remove(tokens)
	syntax = ast.SyntaxClass().Syntax(
		notice,
		comment1,
		rules,
		comment2,
		expressions,
	)
	return
}

func (v *parser_) parseTerm() (
	term ast.TermLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Literal Term.
	var literal ast.LiteralLike
	literal, token, ok = v.parseLiteral()
	if ok {
		// Found a single Literal Term.
		term = ast.TermClass().Term(literal)
		return
	}

	// Attempt to parse a single Reference Term.
	var reference ast.ReferenceLike
	reference, token, ok = v.parseReference()
	if ok {
		// Found a single Reference Term.
		term = ast.TermClass().Term(reference)
		return
	}

	// This is not a single Term rule.
	return
}

func (v *parser_) parseText() (
	text ast.TextLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single glyph Text.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if ok {
		// Found a single glyph Text.
		text = ast.TextClass().Text(glyph)
		return
	}

	// Attempt to parse a single quote Text.
	var quote string
	quote, token, ok = v.parseToken(QuoteToken)
	if ok {
		// Found a single quote Text.
		text = ast.TextClass().Text(quote)
		return
	}

	// Attempt to parse a single lowercase Text.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if ok {
		// Found a single lowercase Text.
		text = ast.TextClass().Text(lowercase)
		return
	}

	// Attempt to parse a single intrinsic Text.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken)
	if ok {
		// Found a single intrinsic Text.
		text = ast.TextClass().Text(intrinsic)
		return
	}

	// This is not a single Text rule.
	return
}

func (v *parser_) parseDelimiter(
	expectedValue string,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == expectedValue {
			// Found the desired delimiter.
			return
		}
		v.next_.AddValue(token)
		ok = false
	}

	// This is not the desired delimiter.
	return
}

func (v *parser_) parseToken(
	tokenType TokenType,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	var tokens = fra.List[TokenLike]()
	token = v.getNextToken()
	for token != nil {
		tokens.AppendValue(token)
		switch token.GetType() {
		case tokenType:
			// Found the desired token type.
			value = token.GetValue()
			ok = true
			return
		case SpaceToken, NewlineToken:
			// Ignore any unrequested whitespace.
			token = v.getNextToken()
		default:
			// This is not the desired token type.
			v.putBack(tokens)
			return
		}
	}

	// We are at the end-of-file marker.
	return
}

func (v *parser_) formatError(
	ruleName string,
	token TokenLike,
) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		ScannerClass().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source lines with the error in it.
	message += "\033[36m"
	for index := line - 3; index < line; index++ {
		if index > 1 {
			message += fmt.Sprintf("%04d: ", index) + string(lines[index-1]) + "\n"
		}
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uint(len(lines)) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"
	if uti.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			v.getDefinition(ruleName),
		)
	}
	return message
}

func (v *parser_) getDefinition(
	ruleName string,
) string {
	var syntax = parserClass().syntax_
	var definition = syntax.GetValue(ruleName)
	return definition
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveLast()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveFirst() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError("", token)
		panic(message)
	}

	return token
}

func (v *parser_) putBack(
	tokens col.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

func (v *parser_) remove(
	tokens col.Sequential[TokenLike],
) {
	// NOTE: This method does nothing but must exist to satisfy the lint
	// check on the generated parser code.
}

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   col.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	syntax_ col.CatalogLike[string, string]
}

// Class Reference

func parserClass() *parserClass_ {
	return parserClassReference_
}

var parserClassReference_ = &parserClass_{
	// Initialize the class constants.
	syntax_: fra.CatalogFromMap[string, string](
		map[string]string{
			"$Syntax": `Notice comment Rule+ comment Expression+`,
			"$Notice": `comment newline`,
			"$Rule":   `"$" uppercase ":" Definition`,
			"$Definition": `
    Multirule
    Multiexpression
    Inline  ! This must be the last option since it skips newlines.`,
			"$Multirule":        `RuleOption+`,
			"$RuleOption":       `newline uppercase note?`,
			"$Multiexpression":  `ExpressionOption+`,
			"$ExpressionOption": `newline lowercase note?`,
			"$Inline":           `Term+ note?`,
			"$Term": `
    Literal
    Reference`,
			"$Literal":   `quote`,
			"$Reference": `Identifier Cardinality?  ! The default cardinality is one.`,
			"$Identifier": `
    lowercase
    uppercase`,
			"$Cardinality": `
    Constrained
    Quantified`,
			"$Constrained": `
    optional
    repeated`,
			"$Quantified":  `"{" number Limit? "}"`,
			"$Limit":       `".." number?  ! The limit of a range of numbers is inclusive.`,
			"$Expression":  `"$" lowercase ":" Pattern note?`,
			"$Pattern":     `Option Alternative*`,
			"$Alternative": `"|" Option`,
			"$Option":      `Repetition+`,
			"$Repetition":  `Element Cardinality?  ! The default cardinality is one.`,
			"$Element": `
    Group
    Filter
    Text`,
			"$Group":  `"(" Pattern ")"`,
			"$Filter": `excluded? "[" Character+ "]"`,
			"$Character": `
    Implicit
    Explicit`,
			"$Implicit": `intrinsic`,
			"$Explicit": `glyph Extent?`,
			"$Extent":   `".." glyph  ! The extent of a range of glyphs is inclusive.`,
			"$Text": `
    glyph
    quote
    lowercase
    intrinsic`,
		},
	),
}
