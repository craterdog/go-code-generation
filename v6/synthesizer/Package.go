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
Package "synthesizer" provides code synthesizers that extract attributes from a
model and integrate them with code templates to generate working Go code
fragments that are used by template driven code generators.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-code-generation/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package synthesizer

import (
	mod "github.com/craterdog/go-class-model/v5"
	ass "github.com/craterdog/go-code-generation/v6/assembler"
	col "github.com/craterdog/go-collection-framework/v5/collection"
	not "github.com/craterdog/go-syntax-notation/v6"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AstSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all ast-synthesizer-class-like classes.
*/
type AstSynthesizerClassLike interface {
	// Constructor Methods
	AstSynthesizer(
		syntax not.SyntaxLike,
	) AstSynthesizerLike
}

/*
ClassSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all class-synthesizer-class-like classes.
*/
type ClassSynthesizerClassLike interface {
	// Constructor Methods
	ClassSynthesizer(
		model mod.ModelLike,
		className string,
	) ClassSynthesizerLike
}

/*
FormatterSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all formatter-synthesizer-class-like classes.
*/
type FormatterSynthesizerClassLike interface {
	// Constructor Methods
	FormatterSynthesizer(
		syntax not.SyntaxLike,
	) FormatterSynthesizerLike
}

/*
GrammarSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all grammar-synthesizer-class-like classes.
*/
type GrammarSynthesizerClassLike interface {
	// Constructor Methods
	GrammarSynthesizer(
		syntax not.SyntaxLike,
	) GrammarSynthesizerLike
}

/*
ModuleSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all module-synthesizer-class-like classes.
*/
type ModuleSynthesizerClassLike interface {
	// Constructor Methods
	ModuleSynthesizer(
		moduleName string,
		models col.CatalogLike[string, mod.ModelLike],
	) ModuleSynthesizerLike
}

/*
NodeSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all node-synthesizer-class-like classes.
*/
type NodeSynthesizerClassLike interface {
	// Constructor Methods
	NodeSynthesizer(
		model mod.ModelLike,
		className string,
	) NodeSynthesizerLike
}

/*
PackageSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all package-synthesizer-class-like classes.
*/
type PackageSynthesizerClassLike interface {
	// Constructor Methods
	PackageSynthesizer() PackageSynthesizerLike
}

/*
ParserSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all parser-synthesizer-class-like classes.
*/
type ParserSynthesizerClassLike interface {
	// Constructor Methods
	ParserSynthesizer(
		syntax not.SyntaxLike,
	) ParserSynthesizerLike
}

/*
ProcessorSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all processor-synthesizer-class-like classes.
*/
type ProcessorSynthesizerClassLike interface {
	// Constructor Methods
	ProcessorSynthesizer(
		syntax not.SyntaxLike,
	) ProcessorSynthesizerLike
}

/*
ScannerSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all scanner-synthesizer-class-like classes.
*/
type ScannerSynthesizerClassLike interface {
	// Constructor Methods
	ScannerSynthesizer(
		syntax not.SyntaxLike,
	) ScannerSynthesizerLike
}

/*
TokenSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all token-synthesizer-class-like classes.
*/
type TokenSynthesizerClassLike interface {
	// Constructor Methods
	TokenSynthesizer(
		syntax not.SyntaxLike,
	) TokenSynthesizerLike
}

/*
ValidatorSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all validator-synthesizer-class-like classes.
*/
type ValidatorSynthesizerClassLike interface {
	// Constructor Methods
	ValidatorSynthesizer(
		syntax not.SyntaxLike,
	) ValidatorSynthesizerLike
}

/*
VisitorSynthesizerClassLike declares the set of class constants, constructors and
functions that must be supported by all visitor-synthesizer-class-like classes.
*/
type VisitorSynthesizerClassLike interface {
	// Constructor Methods
	VisitorSynthesizer(
		syntax not.SyntaxLike,
	) VisitorSynthesizerLike
}

// INSTANCE DECLARATIONS

/*
AstSynthesizerLike declares the set of aspects and methods that must be supported
by all ast-synthesizer-like instances.
*/
type AstSynthesizerLike interface {
	// Principal Methods
	GetClass() AstSynthesizerClassLike

	// Aspect Interfaces
	ass.PackageTemplateDriven
}

/*
ClassSynthesizerLike declares the set of aspects and methods that must
be supported by all class-synthesizer-like instances.
*/
type ClassSynthesizerLike interface {
	// Principal Methods
	GetClass() ClassSynthesizerClassLike

	// Aspect Interfaces
	ass.ClassTemplateDriven
}

/*
FormatterSynthesizerLike declares the set of aspects and methods that must
be supported by all formatter-synthesizer-like instances.
*/
type FormatterSynthesizerLike interface {
	// Principal Methods
	GetClass() FormatterSynthesizerClassLike

	// Aspect Interfaces
	ass.ClassTemplateDriven
}

/*
GrammarSynthesizerLike declares the set of aspects and methods that must
be supported by all grammar-synthesizer-like instances.
*/
type GrammarSynthesizerLike interface {
	// Principal Methods
	GetClass() GrammarSynthesizerClassLike

	// Aspect Interfaces
	ass.PackageTemplateDriven
}

/*
ModuleSynthesizerLike declares the set of aspects and methods that must
be supported by all module-synthesizer-like instances.
*/
type ModuleSynthesizerLike interface {
	// Principal Methods
	GetClass() ModuleSynthesizerClassLike

	// Aspect Interfaces
	ass.ModuleTemplateDriven
}

/*
NodeSynthesizerLike declares the set of aspects and methods that must
be supported by all node-synthesizer-like instances.
*/
type NodeSynthesizerLike interface {
	// Principal Methods
	GetClass() NodeSynthesizerClassLike

	// Aspect Interfaces
	ass.ClassTemplateDriven
}

/*
PackageSynthesizerLike declares the set of aspects and methods that must
be supported by all package-synthesizer-like instances.
*/
type PackageSynthesizerLike interface {
	// Principal Methods
	GetClass() PackageSynthesizerClassLike

	// Aspect Interfaces
	ass.PackageTemplateDriven
}

/*
ParserSynthesizerLike declares the set of aspects and methods that must
be supported by all parser-synthesizer-like instances.
*/
type ParserSynthesizerLike interface {
	// Principal Methods
	GetClass() ParserSynthesizerClassLike

	// Aspect Interfaces
	ass.ClassTemplateDriven
}

/*
ProcessorSynthesizerLike declares the set of aspects and methods that must
be supported by all processor-synthesizer-like instances.
*/
type ProcessorSynthesizerLike interface {
	// Principal Methods
	GetClass() ProcessorSynthesizerClassLike

	// Aspect Interfaces
	ass.ClassTemplateDriven
}

/*
ScannerSynthesizerLike declares the set of aspects and methods that must
be supported by all scanner-synthesizer-like instances.
*/
type ScannerSynthesizerLike interface {
	// Principal Methods
	GetClass() ScannerSynthesizerClassLike

	// Aspect Interfaces
	ass.ClassTemplateDriven
}

/*
TokenSynthesizerLike declares the set of aspects and methods that must
be supported by all token-synthesizer-like instances.
*/
type TokenSynthesizerLike interface {
	// Principal Methods
	GetClass() TokenSynthesizerClassLike

	// Aspect Interfaces
	ass.ClassTemplateDriven
}

/*
ValidatorSynthesizerLike declares the set of aspects and methods that must
be supported by all validator-synthesizer-like instances.
*/
type ValidatorSynthesizerLike interface {
	// Principal Methods
	GetClass() ValidatorSynthesizerClassLike

	// Aspect Interfaces
	ass.ClassTemplateDriven
}

/*
VisitorSynthesizerLike declares the set of aspects and methods that must
be supported by all visitor-synthesizer-like instances.
*/
type VisitorSynthesizerLike interface {
	// Principal Methods
	GetClass() VisitorSynthesizerClassLike

	// Aspect Interfaces
	ass.ClassTemplateDriven
}

// ASPECT DECLARATIONS
