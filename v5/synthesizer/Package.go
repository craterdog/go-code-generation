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
Package "synthesizer" provides code synthesizers that extract attributes from a
model and integrate them with code templates to generate working Go code
fragments that are used by template driven code generators.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-code-generation/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package synthesizer

import (
	mod "github.com/craterdog/go-class-model/v5"
	gen "github.com/craterdog/go-code-generation/v5/generator"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	not "github.com/craterdog/go-syntax-notation/v5"
)

// Class Declarations

/*
AstSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all ast-synthesizer-class-like classes.
*/
type AstSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) AstSynthesizerLike
}

/*
ClassSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all class-synthesizer-class-like classes.
*/
type ClassSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		model mod.ModelLike,
		className string,
	) ClassSynthesizerLike
}

/*
ExampleSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all example-synthesizer-class-like classes.
*/
type ExampleSynthesizerClassLike interface {
	// Constructor Methods
	Make() ExampleSynthesizerLike
}

/*
FormatterSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all formatter-synthesizer-class-like classes.
*/
type FormatterSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) FormatterSynthesizerLike
}

/*
GrammarSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all grammar-synthesizer-class-like classes.
*/
type GrammarSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) GrammarSynthesizerLike
}

/*
ModuleSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all module-synthesizer-class-like classes.
*/
type ModuleSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		models abs.Sequential[mod.ModelLike],
	) ModuleSynthesizerLike
}

/*
ParserSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all parser-synthesizer-class-like classes.
*/
type ParserSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) ParserSynthesizerLike
}

/*
ProcessorSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all processor-synthesizer-class-like classes.
*/
type ProcessorSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) ProcessorSynthesizerLike
}

/*
ScannerSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all scanner-synthesizer-class-like classes.
*/
type ScannerSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) ScannerSynthesizerLike
}

/*
TokenSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all token-synthesizer-class-like classes.
*/
type TokenSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) TokenSynthesizerLike
}

/*
ValidatorSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all validator-synthesizer-class-like classes.
*/
type ValidatorSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) ValidatorSynthesizerLike
}

/*
VisitorSynthesizerClassLike defines the set of class constants, constructors and
functions that must be supported by all visitor-synthesizer-class-like classes.
*/
type VisitorSynthesizerClassLike interface {
	// Constructor Methods
	Make(
		syntax not.SyntaxLike,
	) VisitorSynthesizerLike
}

// Instance Declarations

/*
AstSynthesizerLike defines the set of aspects and methods that must be supported
by all ast-synthesizer-like instances.
*/
type AstSynthesizerLike interface {
	// Primary Methods
	GetClass() AstSynthesizerClassLike

	// Aspect Interfaces
	gen.PackageTemplateDriven
}

/*
ClassSynthesizerLike defines the set of aspects and methods that must be
supported by all class-synthesizer-like instances.
*/
type ClassSynthesizerLike interface {
	// Primary Methods
	GetClass() ClassSynthesizerClassLike

	// Aspect Interfaces
	gen.ClassTemplateDriven
}

/*
ExampleSynthesizerLike defines the set of aspects and methods that must be
supported by all example-synthesizer-like instances.
*/
type ExampleSynthesizerLike interface {
	// Primary Methods
	GetClass() ExampleSynthesizerClassLike

	// Aspect Interfaces
	gen.PackageTemplateDriven
}

/*
FormatterSynthesizerLike defines the set of aspects and methods that must be
supported by all formatter-synthesizer-like instances.
*/
type FormatterSynthesizerLike interface {
	// Primary Methods
	GetClass() FormatterSynthesizerClassLike

	// Aspect Interfaces
	gen.ClassTemplateDriven
}

/*
GrammarSynthesizerLike defines the set of aspects and methods that must be
supported by all grammar-synthesizer-like instances.
*/
type GrammarSynthesizerLike interface {
	// Primary Methods
	GetClass() GrammarSynthesizerClassLike

	// Aspect Interfaces
	gen.PackageTemplateDriven
}

/*
ModuleSynthesizerLike defines the set of aspects and methods that must be
supported by all module-synthesizer-like instances.
*/
type ModuleSynthesizerLike interface {
	// Primary Methods
	GetClass() ModuleSynthesizerClassLike

	// Aspect Interfaces
	gen.ModuleTemplateDriven
}

/*
ParserSynthesizerLike defines the set of aspects and methods that must be
supported by all parser-synthesizer-like instances.
*/
type ParserSynthesizerLike interface {
	// Primary Methods
	GetClass() ParserSynthesizerClassLike

	// Aspect Interfaces
	gen.ClassTemplateDriven
}

/*
ProcessorSynthesizerLike defines the set of aspects and methods that must be
supported by all processor-synthesizer-like instances.
*/
type ProcessorSynthesizerLike interface {
	// Primary Methods
	GetClass() ProcessorSynthesizerClassLike

	// Aspect Interfaces
	gen.ClassTemplateDriven
}

/*
ScannerSynthesizerLike defines the set of aspects and methods that must be
supported by all scanner-synthesizer-like instances.
*/
type ScannerSynthesizerLike interface {
	// Primary Methods
	GetClass() ScannerSynthesizerClassLike

	// Aspect Interfaces
	gen.ClassTemplateDriven
}

/*
TokenSynthesizerLike defines the set of aspects and methods that must be
supported by all token-synthesizer-like instances.
*/
type TokenSynthesizerLike interface {
	// Primary Methods
	GetClass() TokenSynthesizerClassLike

	// Aspect Interfaces
	gen.ClassTemplateDriven
}

/*
ValidatorSynthesizerLike defines the set of aspects and methods that must be
supported by all validator-synthesizer-like instances.
*/
type ValidatorSynthesizerLike interface {
	// Primary Methods
	GetClass() ValidatorSynthesizerClassLike

	// Aspect Interfaces
	gen.ClassTemplateDriven
}

/*
VisitorSynthesizerLike defines the set of aspects and methods that must be
supported by all visitor-synthesizer-like instances.
*/
type VisitorSynthesizerLike interface {
	// Primary Methods
	GetClass() VisitorSynthesizerClassLike

	// Aspect Interfaces
	gen.ClassTemplateDriven
}
