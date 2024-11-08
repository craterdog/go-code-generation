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

package packages

import (
	not "github.com/craterdog/go-syntax-notation/v5"
)

// CLASS INTERFACE

// Access Function

func ExampleSynthesizer() ExampleSynthesizerClassLike {
	return exampleSynthesizerReference()
}

// Constructor Methods

func (c *exampleSynthesizerClass_) Make(
	syntax not.SyntaxLike,
) ExampleSynthesizerLike {
	var instance = &exampleSynthesizer_{
		// Initialize the instance attributes.
		analyzer_: not.Analyzer(syntax),
	}
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *exampleSynthesizer_) GetClass() ExampleSynthesizerClassLike {
	return exampleSynthesizerReference()
}

// TemplateDriven Methods

func (v *exampleSynthesizer_) CreateLegalNotice() string {
	var legalNotice = v.analyzer_.GetNotice()
	return legalNotice
}

func (v *exampleSynthesizer_) CreatePackageDescription() string {
	var packageDescription = exampleSynthesizerReference().packageDescription_
	return packageDescription
}

func (v *exampleSynthesizer_) CreateModuleImports() string {
	var moduleImports = exampleSynthesizerReference().moduleImports_
	return moduleImports
}

func (v *exampleSynthesizer_) CreateTypeDeclarations() string {
	var typeDeclarations = exampleSynthesizerReference().typeDeclarations_
	return typeDeclarations
}

func (v *exampleSynthesizer_) CreateFunctionalDeclarations() string {
	var functionalDeclarations = exampleSynthesizerReference().functionalDeclarations_
	return functionalDeclarations
}

func (v *exampleSynthesizer_) CreateClassDeclarations() string {
	var classDeclarations = exampleSynthesizerReference().classDeclarations_
	return classDeclarations
}

func (v *exampleSynthesizer_) CreateInstanceDeclarations() string {
	var instanceDeclarations = exampleSynthesizerReference().instanceDeclarations_
	return instanceDeclarations
}

func (v *exampleSynthesizer_) CreateAspectDeclarations() string {
	var aspectDeclarations = exampleSynthesizerReference().aspectDeclarations_
	return aspectDeclarations
}

// PROTECTED INTERFACE

// Instance Structure

type exampleSynthesizer_ struct {
	// Declare the instance attributes.
	analyzer_ not.AnalyzerLike
}

// Class Structure

type exampleSynthesizerClass_ struct {
	// Declare the class constants.
	packageDescription_     string
	moduleImports_          string
	typeDeclarations_       string
	functionalDeclarations_ string
	classDeclarations_      string
	instanceDeclarations_   string
	aspectDeclarations_     string
}

// Class Reference

func exampleSynthesizerReference() *exampleSynthesizerClass_ {
	return exampleSynthesizerReference_
}

var exampleSynthesizerReference_ = &exampleSynthesizerClass_{
	// Initialize the class constants.
	packageDescription_: `
Package "example" provides...`,

	moduleImports_: `

import (
// abs "github.com/craterdog/go-collection-framework/v4/collection"
)`,

	typeDeclarations_: `

/*
Form is a constrained type representing the possible notational forms for the
complex number.
*/
type Form uint8

const (
	Polar Form = iota
	Rectangular
)

/*
Rank is a constrained type representing the possible rankings for two values.
*/
type Rank uint8

const (
	LesserRank Rank = iota
	EqualRank
	GreaterRank
)

/*
Units is a constrained type representing the possible units for an angle.
*/
type Units uint8

const (
	Degrees Units = iota
	Radians
	Gradians
)`,

	functionalDeclarations_: `

/*
NormFunction[V any] is a functional type that defines the signature for any
mathematical norm function.
*/
type NormFunction[V any] func(
	value V,
) float64

/*
RankingFunction[V any] is a functional type that defines the signature for any
function that can determine the relative ranking of two values.
*/
type RankingFunction[V any] func(
	first V,
	second V,
) Rank

/*
TrigonometricFunction is a functional type that defines the signature for any
trigonometric function.
*/
type TrigonometricFunction func(
	angle AngleLike,
) float64`,

	classDeclarations_: `

/*
AngleClassLike is a class interface that defines the set of class constants,
constructors and functions that must be supported by each angle-like concrete
class.
*/
type AngleClassLike interface {
	// Constructor Methods
	Make(
		intrinsic float64,
	) AngleLike
	MakeFromString(
		value string,
	) AngleLike

	// Constant Methods
	Pi() AngleLike
	Tau() AngleLike

	// Function Methods
	Apply(
		function TrigonometricFunction,
		angle AngleLike,
	) float64
	Sine(
		angle AngleLike,
	) float64
	Cosine(
		angle AngleLike,
	) float64
	Tangent(
		angle AngleLike,
	) float64
}

/*
ArrayClassLike[V any] is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete array-like class.
*/
type ArrayClassLike[V any] interface {
	// Constructor Methods
	Make(
		intrinsic []V,
	) ArrayLike[V]
	MakeFromSize(
		size uint,
	) ArrayLike[V]
	MakeFromSequence(
		values Sequential[V],
	) ArrayLike[V]

	// Constant Methods
	DefaultRanker() RankingFunction[V]
}

/*
AssociationClassLike[K comparable, V any] is a class interface that defines
the complete set of class constants, constructors and functions that must be
supported by each concrete association-like class.
*/
type AssociationClassLike[K comparable, V any] interface {
	// Constructor Methods
	Make(
		key K,
		value V,
	) AssociationLike[K, V]
}

/*
CatalogClassLike[K comparable, V any] is a class interface that defines the
complete set of class constants, constructors and functions that must be
supported by each concrete catalog-like class.

The following functions are supported:

Extract() returns a new catalog containing only the associations that are in
the specified catalog that have the specified keys.  The associations in the
resulting catalog will be in the same order as the specified keys.

Merge() returns a new catalog containing all of the associations that are in
the specified catalogs in the order that they appear in each catalog.  If a
key is present in both catalogs, the value of the key from the second
catalog takes precedence.
*/
type CatalogClassLike[K comparable, V any] interface {
	// Constructor Methods
	Make() CatalogLike[K, V]
	MakeFromArray(
		associations []AssociationLike[K, V],
	) CatalogLike[K, V]
	MakeFromMap(
		associations map[K]V,
	) CatalogLike[K, V]
	MakeFromSequence(
		associations Sequential[AssociationLike[K, V]],
	) CatalogLike[K, V]

	// Constant Methods
	DefaultRanker() RankingFunction[AssociationLike[K, V]]

	// Function Methods
	Extract(
		catalog CatalogLike[K, V],
		keys Sequential[K],
	) CatalogLike[K, V]
	Merge(
		first CatalogLike[K, V],
		second CatalogLike[K, V],
	) CatalogLike[K, V]
}

/*
ComplexClassLike is a class interface that defines the set of class constants,
constructors and functions that must be supported by each complex-like concrete
class.
*/
type ComplexClassLike interface {
	// Constructor Methods
	Make(
		realPart float64,
		imaginaryPart float64,
		form Form,
	) ComplexLike
	MakeFromValue(
		value complex128,
	) ComplexLike

	// Constant Methods
	Zero() ComplexLike
	Infinity() ComplexLike

	// Function Methods
	Inverse(
		value ComplexLike,
	) ComplexLike
	Sum(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Difference(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Reciprocal(
		value ComplexLike,
	) ComplexLike
	Product(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Quotient(
		first ComplexLike,
		second ComplexLike,
	) ComplexLike
	Norm(
		function NormFunction[ComplexLike],
		value ComplexLike,
	) float64
}`,

	instanceDeclarations_: `

/*
AngleLike is an instance interface that defines the complete set of attributes,
abstractions and methods that must be supported by each instance of a concrete
angle-like class.
*/
type AngleLike interface {
	// Primary Methods
	GetClass() AngleClassLike
	GetIntrinsic() float64
	IsZero() bool

	// Aspect Interfaces
	Angular
}

/*
ArrayLike[V any] is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete array-like class.

An array-like class maintains a fixed length indexed sequence of values.  Each
value is associated with an implicit positive integer index. An array-like class
uses ORDINAL based indexing rather than the more common—and nonsensical—ZERO
based indexing scheme.
*/
type ArrayLike[V any] interface {
	// Primary Methods
	GetClass() ArrayClassLike[V]
	GetIntrinsic() []V
	SortValues()
	SortValuesWithRanker(
		ranker RankingFunction[V],
	)

	// Aspect Interfaces
	Accessible[V]
	Sequential[V]
	Updatable[V]
}

/*
AssociationLike[K comparable, V any] is an instance interface that defines the
complete set of instance attributes, abstractions and methods that must be
supported by each instance of a concrete association-like class.
*/
type AssociationLike[K comparable, V any] interface {
	// Primary Methods
	GetClass() AssociationClassLike[K, V]

	// Attribute Methods
	GetKey() K
	GetValue() V
	SetValue(
		value V,
	)
}

/*
CatalogLike[K comparable, V any] is an instance interface that defines the
complete set of instance attributes, abstractions and methods that must be
supported by each instance of a concrete catalog-like class.
*/
type CatalogLike[K comparable, V any] interface {
	// Primary Methods
	GetClass() CatalogClassLike[K, V]
	SortValues()
	SortValuesWithRanker(
		ranker RankingFunction[AssociationLike[K, V]],
	)

	// Aspect Interfaces
	Associative[K, V]
	Sequential[AssociationLike[K, V]]
}

/*
ComplexLike is an instance interface that defines the complete set of attributes,
abstractions and methods that must be supported by each instance of a concrete
complex-like class.
*/
type ComplexLike interface {
	// Primary Methods
	GetClass() ComplexClassLike
	IsReal() bool
	IsImaginary() bool

	// Attribute Methods
	GetRealPart() float64
	GetImaginaryPart() float64
	GetForm() Form
	SetForm(
		form Form,
	)

	// Aspect Interfaces
	Continuous
}`,

	aspectDeclarations_: `

/*
Accessible[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of an accessible concrete class.  The
values in an accessible class are accessed using indices. The indices of an
accessible class are ORDINAL rather than ZERO based—which never really made
sense except for pointer offsets.

This approach allows for positive indices starting at the beginning of the
sequence, and negative indices starting at the end of the sequence as follows:

|      1           2           3             N      |
|  [value 1] . [value 2] . [value 3] ... [value N]  |
|     -N        -(N-1)      -(N-2)          -1      |

Notice that because the indices are ordinal based, the positive and negative
indices are symmetrical.
*/
type Accessible[V any] interface {
	GetValue(
		index int,
	) V
	GetValues(
		first int,
		last int,
	) Sequential[V]
}

/*
Angular is an aspect interface that defines a set of method signatures that
must be supported by each instance of an angular concrete class.
*/
type Angular interface {
	AsNormalized() AngleLike
	InUnits(
		units Units,
	) float64
}

/*
Associative[K comparable, V any] defines the set of method signatures that
must be supported by all sequences of key-value associations.
*/
type Associative[K comparable, V any] interface {
	GetKeys() Sequential[K]
	GetValue(
		key K,
	) V
	RemoveValue(
		key K,
	) V
	SetValue(
		key K,
		value V,
	)
}

/*
Continuous is an aspect interface that defines a set of method signatures
that must be supported by each instance of a continuous concrete class.
*/
type Continuous interface {
	IsZero() bool
	IsDiscrete() bool
	IsInfinity() bool
}

/*
Sequential[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of a sequential concrete class.

NOTE: that sizes should be of type "uint" but the Go language does not allow
arithmetic and comparison operations between "int" and "uint" so we us "int" for
the return type to make it easier to use.
*/
type Sequential[V any] interface {
	IsEmpty() bool
	GetSize() int
	AsArray() []V
}

/*
Updatable[V any] is an aspect interface that defines a set of method signatures
that must be supported by each instance of an updatable concrete class.
*/
type Updatable[V any] interface {
	SetValue(
		index int,
		value V,
	)
	SetValues(
		index int,
		values Sequential[V],
	)
}`,
}
