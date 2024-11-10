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

package grammar

import (
	fmt "fmt"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	ast "github.com/craterdog/go-syntax-notation/v5/ast"
	mat "math"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func Parser() ParserClassLike {
	return parserReference()
}

// Constructor Methods

func (c *parserClass_) Make() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserReference()
}

func (v *parser_) ParseSource(
	source string,
) ast.SyntaxLike {
	v.source_ = source
	v.tokens_ = col.Queue[TokenLike](parserReference().queueSize_)
	v.next_ = col.Stack[TokenLike](parserReference().stackSize_)

	// The scanner runs in a separate Go routine.
	Scanner().Make(v.source_, v.tokens_)

	// Attempt to parse the syntax.
	var syntax, token, ok = v.parseSyntax()
	if !ok {
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
	var tokens = col.List[TokenLike]()

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
		// This is not a single Alternative rule.
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
	alternative = ast.Alternative().Make(option)
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
		cardinality = ast.Cardinality().Make(constrained)
		return
	}

	// Attempt to parse a single Quantified Cardinality.
	var quantified ast.QuantifiedLike
	quantified, token, ok = v.parseQuantified()
	if ok {
		// Found a single Quantified Cardinality.
		cardinality = ast.Cardinality().Make(quantified)
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
	// Attempt to parse a single Explicit Character.
	var explicit ast.ExplicitLike
	explicit, token, ok = v.parseExplicit()
	if ok {
		// Found a single Explicit Character.
		character = ast.Character().Make(explicit)
		return
	}

	// Attempt to parse a single intrinsic Character.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken)
	if ok {
		// Found a single intrinsic Character.
		character = ast.Character().Make(intrinsic)
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
		constrained = ast.Constrained().Make(optional)
		return
	}

	// Attempt to parse a single repeated Constrained.
	var repeated string
	repeated, token, ok = v.parseToken(RepeatedToken)
	if ok {
		// Found a single repeated Constrained.
		constrained = ast.Constrained().Make(repeated)
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
	// Attempt to parse a single Multiline Definition.
	var multiline ast.MultilineLike
	multiline, token, ok = v.parseMultiline()
	if ok {
		// Found a single Multiline Definition.
		definition = ast.Definition().Make(multiline)
		return
	}

	// Attempt to parse a single Inline Definition.
	var inline ast.InlineLike
	inline, token, ok = v.parseInline()
	if ok {
		// Found a single Inline Definition.
		definition = ast.Definition().Make(inline)
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
		element = ast.Element().Make(group)
		return
	}

	// Attempt to parse a single Filter Element.
	var filter ast.FilterLike
	filter, token, ok = v.parseFilter()
	if ok {
		// Found a single Filter Element.
		element = ast.Element().Make(filter)
		return
	}

	// Attempt to parse a single Text Element.
	var text ast.TextLike
	text, token, ok = v.parseText()
	if ok {
		// Found a single Text Element.
		element = ast.Element().Make(text)
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
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single glyph token.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Explicit rule.
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
	explicit = ast.Explicit().Make(
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
	var tokens = col.List[TokenLike]()

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
		// This is not a single Expression rule.
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
	expression = ast.Expression().Make(
		lowercase,
		pattern,
		optionalNote,
	)
	return
}

func (v *parser_) parseExtent() (
	extent ast.ExtentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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

	// Found a single Extent rule.
	ok = true
	v.remove(tokens)
	extent = ast.Extent().Make(glyph)
	return
}

func (v *parser_) parseFilter() (
	filter ast.FilterLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

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
	var characters = col.List[ast.CharacterLike]()
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
	filter = ast.Filter().Make(
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
	var tokens = col.List[TokenLike]()

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
		// This is not a single Group rule.
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
	group = ast.Group().Make(pattern)
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
		identifier = ast.Identifier().Make(lowercase)
		return
	}

	// Attempt to parse a single uppercase Identifier.
	var uppercase string
	uppercase, token, ok = v.parseToken(UppercaseToken)
	if ok {
		// Found a single uppercase Identifier.
		identifier = ast.Identifier().Make(uppercase)
		return
	}

	// This is not a single Identifier rule.
	return
}

func (v *parser_) parseInline() (
	inline ast.InlineLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse multiple Term rules.
	var terms = col.List[ast.TermLike]()
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
	inline = ast.Inline().Make(
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
	var tokens = col.List[TokenLike]()

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
	limit = ast.Limit().Make(optionalNumber)
	return
}

func (v *parser_) parseLine() (
	line ast.LineLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "-" delimiter.
	_, token, ok = v.parseDelimiter("-")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Line rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Line", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Identifier rule.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Line rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Line", token)
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Line rule.
	ok = true
	v.remove(tokens)
	line = ast.Line().Make(
		identifier,
		optionalNote,
	)
	return
}

func (v *parser_) parseMultiline() (
	multiline ast.MultilineLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse multiple Line rules.
	var lines = col.List[ast.LineLike]()
linesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var line ast.LineLike
		line, token, ok = v.parseLine()
		if !ok {
			switch {
			case count >= 1:
				break linesLoop
			case uti.IsDefined(tokens):
				// This is not multiple Line rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Multiline", token)
				message += "1 or more Line rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		lines.AppendValue(line)
	}

	// Found a single Multiline rule.
	ok = true
	v.remove(tokens)
	multiline = ast.Multiline().Make(lines)
	return
}

func (v *parser_) parseNotice() (
	notice ast.NoticeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Notice rule.
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
	notice = ast.Notice().Make(comment)
	return
}

func (v *parser_) parseOption() (
	option ast.OptionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse multiple Repetition rules.
	var repetitions = col.List[ast.RepetitionLike]()
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
	option = ast.Option().Make(repetitions)
	return
}

func (v *parser_) parsePattern() (
	pattern ast.PatternLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Option rule.
	var option ast.OptionLike
	option, token, ok = v.parseOption()
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
		var message = v.formatError("$Pattern", token)
		panic(message)
	}

	// Attempt to parse multiple Alternative rules.
	var alternatives = col.List[ast.AlternativeLike]()
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
	pattern = ast.Pattern().Make(
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
	var tokens = col.List[TokenLike]()

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
	quantified = ast.Quantified().Make(
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
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Identifier rule.
	var identifier ast.IdentifierLike
	identifier, token, ok = v.parseIdentifier()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Reference rule.
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
	reference = ast.Reference().Make(
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
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Element rule.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Repetition rule.
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
	repetition = ast.Repetition().Make(
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
	var tokens = col.List[TokenLike]()

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
		// This is not a single Rule rule.
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
	rule = ast.Rule().Make(
		uppercase,
		definition,
	)
	return
}

func (v *parser_) parseSyntax() (
	syntax ast.SyntaxLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Notice rule.
	var notice ast.NoticeLike
	notice, token, ok = v.parseNotice()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Syntax rule.
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
			// This is not a single Syntax rule.
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
	var rules = col.List[ast.RuleLike]()
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
			// This is not a single Syntax rule.
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
	var expressions = col.List[ast.ExpressionLike]()
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
	syntax = ast.Syntax().Make(
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
	// Attempt to parse a single Reference Term.
	var reference ast.ReferenceLike
	reference, token, ok = v.parseReference()
	if ok {
		// Found a single Reference Term.
		term = ast.Term().Make(reference)
		return
	}

	// Attempt to parse a single literal Term.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken)
	if ok {
		// Found a single literal Term.
		term = ast.Term().Make(literal)
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
	// Attempt to parse a single intrinsic Text.
	var intrinsic string
	intrinsic, token, ok = v.parseToken(IntrinsicToken)
	if ok {
		// Found a single intrinsic Text.
		text = ast.Text().Make(intrinsic)
		return
	}

	// Attempt to parse a single glyph Text.
	var glyph string
	glyph, token, ok = v.parseToken(GlyphToken)
	if ok {
		// Found a single glyph Text.
		text = ast.Text().Make(glyph)
		return
	}

	// Attempt to parse a single literal Text.
	var literal string
	literal, token, ok = v.parseToken(LiteralToken)
	if ok {
		// Found a single literal Text.
		text = ast.Text().Make(literal)
		return
	}

	// Attempt to parse a single lowercase Text.
	var lowercase string
	lowercase, token, ok = v.parseToken(LowercaseToken)
	if ok {
		// Found a single lowercase Text.
		text = ast.Text().Make(lowercase)
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

func (v *parser_) parseToken(tokenType TokenType) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	var tokens = col.List[TokenLike]()
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
		Scanner().FormatToken(token),
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
	return parserReference().syntax_.GetValue(ruleName)
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveTop()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveHead() // This will wait for a token.
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
	tokens abs.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

func (v *parser_) remove(
	tokens abs.Sequential[TokenLike],
) {
}

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ abs.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   abs.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	queueSize_ uint
	stackSize_ uint
	syntax_    abs.CatalogLike[string, string]
}

// Class Reference

func parserReference() *parserClass_ {
	return parserReference_
}

var parserReference_ = &parserClass_{
	// Initialize the class constants.
	queueSize_: 16,
	stackSize_: 16,
	syntax_: col.Catalog[string, string](
		map[string]string{
			"$Syntax": `Notice comment Rule+ comment Expression+`,
			"$Notice": `comment`,
			"$Rule":   `"$" uppercase ":" Definition`,
			"$Definition": `
  - Multiline
  - Inline`,
			"$Multiline": `Line+`,
			"$Line":      `"-" Identifier note?`,
			"$Identifier": `
  - lowercase
  - uppercase`,
			"$Inline": `Term+ note?`,
			"$Term": `
  - Reference
  - literal`,
			"$Reference": `Identifier Cardinality?  ! The default cardinality is one.`,
			"$Cardinality": `
  - Constrained
  - Quantified`,
			"$Constrained": `
  - optional
  - repeated`,
			"$Quantified":  `"{" number Limit? "}"`,
			"$Limit":       `".." number?  ! The limit of a range of numbers is inclusive.`,
			"$Expression":  `"$" lowercase ":" Pattern note?`,
			"$Pattern":     `Option Alternative*`,
			"$Alternative": `"|" Option`,
			"$Option":      `Repetition+`,
			"$Repetition":  `Element Cardinality?  ! The default cardinality is one.`,
			"$Element": `
  - Group
  - Filter
  - Text`,
			"$Group":  `"(" Pattern ")"`,
			"$Filter": `excluded? "[" Character+ "]"`,
			"$Character": `
  - Explicit
  - intrinsic`,
			"$Explicit": `glyph Extent?`,
			"$Extent":   `".." glyph  ! The extent of a range of glyphs is inclusive.`,
			"$Text": `
  - intrinsic
  - glyph
  - literal
  - lowercase`,
		},
	),
}
