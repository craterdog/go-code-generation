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
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This "Module.go" file was automatically generated.              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides a default constructor
for each commonly used class that is exported by the module.  Each constructor
delegates the actual construction process to its corresponding concrete class
declared in the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/craterdog/go-class-model/wiki
*/
package module

import (
	ast "github.com/craterdog/go-class-model/v5/ast"
	gra "github.com/craterdog/go-class-model/v5/grammar"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
)

// TYPE ALIASES

// Ast

type (
	AbstractionLike           = ast.AbstractionLike
	AdditionalArgumentLike    = ast.AdditionalArgumentLike
	AdditionalConstraintLike  = ast.AdditionalConstraintLike
	AdditionalValueLike       = ast.AdditionalValueLike
	ArgumentLike              = ast.ArgumentLike
	ArgumentsLike             = ast.ArgumentsLike
	ArrayLike                 = ast.ArrayLike
	AspectDeclarationLike     = ast.AspectDeclarationLike
	AspectInterfaceLike       = ast.AspectInterfaceLike
	AspectMethodLike          = ast.AspectMethodLike
	AspectSectionLike         = ast.AspectSectionLike
	AspectSubsectionLike      = ast.AspectSubsectionLike
	AttributeMethodLike       = ast.AttributeMethodLike
	AttributeSubsectionLike   = ast.AttributeSubsectionLike
	ChannelLike               = ast.ChannelLike
	ClassDeclarationLike      = ast.ClassDeclarationLike
	ClassMethodsLike          = ast.ClassMethodsLike
	ClassSectionLike          = ast.ClassSectionLike
	ConstantMethodLike        = ast.ConstantMethodLike
	ConstantSubsectionLike    = ast.ConstantSubsectionLike
	ConstraintLike            = ast.ConstraintLike
	ConstraintsLike           = ast.ConstraintsLike
	ConstructorMethodLike     = ast.ConstructorMethodLike
	ConstructorSubsectionLike = ast.ConstructorSubsectionLike
	DeclarationLike           = ast.DeclarationLike
	EnumerationLike           = ast.EnumerationLike
	FunctionMethodLike        = ast.FunctionMethodLike
	FunctionSubsectionLike    = ast.FunctionSubsectionLike
	FunctionalDeclarationLike = ast.FunctionalDeclarationLike
	FunctionalSectionLike     = ast.FunctionalSectionLike
	GetterMethodLike          = ast.GetterMethodLike
	ImportedPackageLike       = ast.ImportedPackageLike
	InstanceDeclarationLike   = ast.InstanceDeclarationLike
	InstanceMethodsLike       = ast.InstanceMethodsLike
	InstanceSectionLike       = ast.InstanceSectionLike
	InterfaceDeclarationsLike = ast.InterfaceDeclarationsLike
	LegalNoticeLike           = ast.LegalNoticeLike
	MapLike                   = ast.MapLike
	MethodLike                = ast.MethodLike
	ModelLike                 = ast.ModelLike
	MultivalueLike            = ast.MultivalueLike
	NoneLike                  = ast.NoneLike
	PackageDeclarationLike    = ast.PackageDeclarationLike
	PackageHeaderLike         = ast.PackageHeaderLike
	PackageImportsLike        = ast.PackageImportsLike
	ParameterLike             = ast.ParameterLike
	PrefixLike                = ast.PrefixLike
	PrimitiveDeclarationsLike = ast.PrimitiveDeclarationsLike
	PrincipalMethodLike       = ast.PrincipalMethodLike
	PrincipalSubsectionLike   = ast.PrincipalSubsectionLike
	ResultLike                = ast.ResultLike
	SetterMethodLike          = ast.SetterMethodLike
	SuffixLike                = ast.SuffixLike
	TypeDeclarationLike       = ast.TypeDeclarationLike
	TypeSectionLike           = ast.TypeSectionLike
	ValueLike                 = ast.ValueLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken     = gra.ErrorToken
	CommentToken   = gra.CommentToken
	DelimiterToken = gra.DelimiterToken
	NameToken      = gra.NameToken
	NewlineToken   = gra.NewlineToken
	PathToken      = gra.PathToken
	SpaceToken     = gra.SpaceToken
)

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

type (
	Methodical = gra.Methodical
)

// DEFAULT CONSTRUCTORS

// Ast

func Abstraction(
	optionalPrefix PrefixLike,
	name string,
	optionalSuffix SuffixLike,
	optionalArguments ArgumentsLike,
) AbstractionLike {
	var instance_ AbstractionLike
	instance_ = ast.AbstractionClass().Make(
		optionalPrefix,
		name,
		optionalSuffix,
		optionalArguments,
	)
	return instance_
}

func AdditionalArgument(
	argument ArgumentLike,
) AdditionalArgumentLike {
	var instance_ AdditionalArgumentLike
	instance_ = ast.AdditionalArgumentClass().Make(
		argument,
	)
	return instance_
}

func AdditionalConstraint(
	constraint ConstraintLike,
) AdditionalConstraintLike {
	var instance_ AdditionalConstraintLike
	instance_ = ast.AdditionalConstraintClass().Make(
		constraint,
	)
	return instance_
}

func AdditionalValue(
	name string,
) AdditionalValueLike {
	var instance_ AdditionalValueLike
	instance_ = ast.AdditionalValueClass().Make(
		name,
	)
	return instance_
}

func Argument(
	abstraction AbstractionLike,
) ArgumentLike {
	var instance_ ArgumentLike
	instance_ = ast.ArgumentClass().Make(
		abstraction,
	)
	return instance_
}

func Arguments(
	argument ArgumentLike,
	additionalArguments abs.Sequential[AdditionalArgumentLike],
) ArgumentsLike {
	var instance_ ArgumentsLike
	instance_ = ast.ArgumentsClass().Make(
		argument,
		additionalArguments,
	)
	return instance_
}

func Array() ArrayLike {
	var instance_ ArrayLike
	instance_ = ast.ArrayClass().Make()
	return instance_
}

func AspectDeclaration(
	declaration DeclarationLike,
	aspectMethods abs.Sequential[AspectMethodLike],
) AspectDeclarationLike {
	var instance_ AspectDeclarationLike
	instance_ = ast.AspectDeclarationClass().Make(
		declaration,
		aspectMethods,
	)
	return instance_
}

func AspectInterface(
	abstraction AbstractionLike,
) AspectInterfaceLike {
	var instance_ AspectInterfaceLike
	instance_ = ast.AspectInterfaceClass().Make(
		abstraction,
	)
	return instance_
}

func AspectMethod(
	method MethodLike,
) AspectMethodLike {
	var instance_ AspectMethodLike
	instance_ = ast.AspectMethodClass().Make(
		method,
	)
	return instance_
}

func AspectSection(
	aspectDeclarations abs.Sequential[AspectDeclarationLike],
) AspectSectionLike {
	var instance_ AspectSectionLike
	instance_ = ast.AspectSectionClass().Make(
		aspectDeclarations,
	)
	return instance_
}

func AspectSubsection(
	aspectInterfaces abs.Sequential[AspectInterfaceLike],
) AspectSubsectionLike {
	var instance_ AspectSubsectionLike
	instance_ = ast.AspectSubsectionClass().Make(
		aspectInterfaces,
	)
	return instance_
}

func AttributeMethod(
	any_ any,
) AttributeMethodLike {
	var instance_ AttributeMethodLike
	instance_ = ast.AttributeMethodClass().Make(
		any_,
	)
	return instance_
}

func AttributeSubsection(
	attributeMethods abs.Sequential[AttributeMethodLike],
) AttributeSubsectionLike {
	var instance_ AttributeSubsectionLike
	instance_ = ast.AttributeSubsectionClass().Make(
		attributeMethods,
	)
	return instance_
}

func Channel() ChannelLike {
	var instance_ ChannelLike
	instance_ = ast.ChannelClass().Make()
	return instance_
}

func ClassDeclaration(
	declaration DeclarationLike,
	classMethods ClassMethodsLike,
) ClassDeclarationLike {
	var instance_ ClassDeclarationLike
	instance_ = ast.ClassDeclarationClass().Make(
		declaration,
		classMethods,
	)
	return instance_
}

func ClassMethods(
	constructorSubsection ConstructorSubsectionLike,
	optionalConstantSubsection ConstantSubsectionLike,
	optionalFunctionSubsection FunctionSubsectionLike,
) ClassMethodsLike {
	var instance_ ClassMethodsLike
	instance_ = ast.ClassMethodsClass().Make(
		constructorSubsection,
		optionalConstantSubsection,
		optionalFunctionSubsection,
	)
	return instance_
}

func ClassSection(
	classDeclarations abs.Sequential[ClassDeclarationLike],
) ClassSectionLike {
	var instance_ ClassSectionLike
	instance_ = ast.ClassSectionClass().Make(
		classDeclarations,
	)
	return instance_
}

func ConstantMethod(
	name string,
	abstraction AbstractionLike,
) ConstantMethodLike {
	var instance_ ConstantMethodLike
	instance_ = ast.ConstantMethodClass().Make(
		name,
		abstraction,
	)
	return instance_
}

func ConstantSubsection(
	constantMethods abs.Sequential[ConstantMethodLike],
) ConstantSubsectionLike {
	var instance_ ConstantSubsectionLike
	instance_ = ast.ConstantSubsectionClass().Make(
		constantMethods,
	)
	return instance_
}

func Constraint(
	name string,
	abstraction AbstractionLike,
) ConstraintLike {
	var instance_ ConstraintLike
	instance_ = ast.ConstraintClass().Make(
		name,
		abstraction,
	)
	return instance_
}

func Constraints(
	constraint ConstraintLike,
	additionalConstraints abs.Sequential[AdditionalConstraintLike],
) ConstraintsLike {
	var instance_ ConstraintsLike
	instance_ = ast.ConstraintsClass().Make(
		constraint,
		additionalConstraints,
	)
	return instance_
}

func ConstructorMethod(
	name string,
	parameters abs.Sequential[ParameterLike],
	abstraction AbstractionLike,
) ConstructorMethodLike {
	var instance_ ConstructorMethodLike
	instance_ = ast.ConstructorMethodClass().Make(
		name,
		parameters,
		abstraction,
	)
	return instance_
}

func ConstructorSubsection(
	constructorMethods abs.Sequential[ConstructorMethodLike],
) ConstructorSubsectionLike {
	var instance_ ConstructorSubsectionLike
	instance_ = ast.ConstructorSubsectionClass().Make(
		constructorMethods,
	)
	return instance_
}

func Declaration(
	comment string,
	name string,
	optionalConstraints ConstraintsLike,
) DeclarationLike {
	var instance_ DeclarationLike
	instance_ = ast.DeclarationClass().Make(
		comment,
		name,
		optionalConstraints,
	)
	return instance_
}

func Enumeration(
	value ValueLike,
	additionalValues abs.Sequential[AdditionalValueLike],
) EnumerationLike {
	var instance_ EnumerationLike
	instance_ = ast.EnumerationClass().Make(
		value,
		additionalValues,
	)
	return instance_
}

func FunctionMethod(
	name string,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) FunctionMethodLike {
	var instance_ FunctionMethodLike
	instance_ = ast.FunctionMethodClass().Make(
		name,
		parameters,
		result,
	)
	return instance_
}

func FunctionSubsection(
	functionMethods abs.Sequential[FunctionMethodLike],
) FunctionSubsectionLike {
	var instance_ FunctionSubsectionLike
	instance_ = ast.FunctionSubsectionClass().Make(
		functionMethods,
	)
	return instance_
}

func FunctionalDeclaration(
	declaration DeclarationLike,
	parameters abs.Sequential[ParameterLike],
	result ResultLike,
) FunctionalDeclarationLike {
	var instance_ FunctionalDeclarationLike
	instance_ = ast.FunctionalDeclarationClass().Make(
		declaration,
		parameters,
		result,
	)
	return instance_
}

func FunctionalSection(
	functionalDeclarations abs.Sequential[FunctionalDeclarationLike],
) FunctionalSectionLike {
	var instance_ FunctionalSectionLike
	instance_ = ast.FunctionalSectionClass().Make(
		functionalDeclarations,
	)
	return instance_
}

func GetterMethod(
	name string,
	abstraction AbstractionLike,
) GetterMethodLike {
	var instance_ GetterMethodLike
	instance_ = ast.GetterMethodClass().Make(
		name,
		abstraction,
	)
	return instance_
}

func ImportedPackage(
	name string,
	path string,
) ImportedPackageLike {
	var instance_ ImportedPackageLike
	instance_ = ast.ImportedPackageClass().Make(
		name,
		path,
	)
	return instance_
}

func InstanceDeclaration(
	declaration DeclarationLike,
	instanceMethods InstanceMethodsLike,
) InstanceDeclarationLike {
	var instance_ InstanceDeclarationLike
	instance_ = ast.InstanceDeclarationClass().Make(
		declaration,
		instanceMethods,
	)
	return instance_
}

func InstanceMethods(
	principalSubsection PrincipalSubsectionLike,
	optionalAttributeSubsection AttributeSubsectionLike,
	optionalAspectSubsection AspectSubsectionLike,
) InstanceMethodsLike {
	var instance_ InstanceMethodsLike
	instance_ = ast.InstanceMethodsClass().Make(
		principalSubsection,
		optionalAttributeSubsection,
		optionalAspectSubsection,
	)
	return instance_
}

func InstanceSection(
	instanceDeclarations abs.Sequential[InstanceDeclarationLike],
) InstanceSectionLike {
	var instance_ InstanceSectionLike
	instance_ = ast.InstanceSectionClass().Make(
		instanceDeclarations,
	)
	return instance_
}

func InterfaceDeclarations(
	classSection ClassSectionLike,
	instanceSection InstanceSectionLike,
	aspectSection AspectSectionLike,
) InterfaceDeclarationsLike {
	var instance_ InterfaceDeclarationsLike
	instance_ = ast.InterfaceDeclarationsClass().Make(
		classSection,
		instanceSection,
		aspectSection,
	)
	return instance_
}

func LegalNotice(
	comment string,
) LegalNoticeLike {
	var instance_ LegalNoticeLike
	instance_ = ast.LegalNoticeClass().Make(
		comment,
	)
	return instance_
}

func Map(
	name string,
) MapLike {
	var instance_ MapLike
	instance_ = ast.MapClass().Make(
		name,
	)
	return instance_
}

func Method(
	name string,
	parameters abs.Sequential[ParameterLike],
	optionalResult ResultLike,
) MethodLike {
	var instance_ MethodLike
	instance_ = ast.MethodClass().Make(
		name,
		parameters,
		optionalResult,
	)
	return instance_
}

func Model(
	packageDeclaration PackageDeclarationLike,
	primitiveDeclarations PrimitiveDeclarationsLike,
	interfaceDeclarations InterfaceDeclarationsLike,
) ModelLike {
	var instance_ ModelLike
	instance_ = ast.ModelClass().Make(
		packageDeclaration,
		primitiveDeclarations,
		interfaceDeclarations,
	)
	return instance_
}

func Multivalue(
	parameters abs.Sequential[ParameterLike],
) MultivalueLike {
	var instance_ MultivalueLike
	instance_ = ast.MultivalueClass().Make(
		parameters,
	)
	return instance_
}

func None(
	newline string,
) NoneLike {
	var instance_ NoneLike
	instance_ = ast.NoneClass().Make(
		newline,
	)
	return instance_
}

func PackageDeclaration(
	legalNotice LegalNoticeLike,
	packageHeader PackageHeaderLike,
	packageImports PackageImportsLike,
) PackageDeclarationLike {
	var instance_ PackageDeclarationLike
	instance_ = ast.PackageDeclarationClass().Make(
		legalNotice,
		packageHeader,
		packageImports,
	)
	return instance_
}

func PackageHeader(
	comment string,
	name string,
) PackageHeaderLike {
	var instance_ PackageHeaderLike
	instance_ = ast.PackageHeaderClass().Make(
		comment,
		name,
	)
	return instance_
}

func PackageImports(
	importedPackages abs.Sequential[ImportedPackageLike],
) PackageImportsLike {
	var instance_ PackageImportsLike
	instance_ = ast.PackageImportsClass().Make(
		importedPackages,
	)
	return instance_
}

func Parameter(
	name string,
	abstraction AbstractionLike,
) ParameterLike {
	var instance_ ParameterLike
	instance_ = ast.ParameterClass().Make(
		name,
		abstraction,
	)
	return instance_
}

func Prefix(
	any_ any,
) PrefixLike {
	var instance_ PrefixLike
	instance_ = ast.PrefixClass().Make(
		any_,
	)
	return instance_
}

func PrimitiveDeclarations(
	typeSection TypeSectionLike,
	functionalSection FunctionalSectionLike,
) PrimitiveDeclarationsLike {
	var instance_ PrimitiveDeclarationsLike
	instance_ = ast.PrimitiveDeclarationsClass().Make(
		typeSection,
		functionalSection,
	)
	return instance_
}

func PrincipalMethod(
	method MethodLike,
) PrincipalMethodLike {
	var instance_ PrincipalMethodLike
	instance_ = ast.PrincipalMethodClass().Make(
		method,
	)
	return instance_
}

func PrincipalSubsection(
	principalMethods abs.Sequential[PrincipalMethodLike],
) PrincipalSubsectionLike {
	var instance_ PrincipalSubsectionLike
	instance_ = ast.PrincipalSubsectionClass().Make(
		principalMethods,
	)
	return instance_
}

func Result(
	any_ any,
) ResultLike {
	var instance_ ResultLike
	instance_ = ast.ResultClass().Make(
		any_,
	)
	return instance_
}

func SetterMethod(
	name string,
	parameter ParameterLike,
) SetterMethodLike {
	var instance_ SetterMethodLike
	instance_ = ast.SetterMethodClass().Make(
		name,
		parameter,
	)
	return instance_
}

func Suffix(
	name string,
) SuffixLike {
	var instance_ SuffixLike
	instance_ = ast.SuffixClass().Make(
		name,
	)
	return instance_
}

func TypeDeclaration(
	declaration DeclarationLike,
	abstraction AbstractionLike,
	optionalEnumeration EnumerationLike,
) TypeDeclarationLike {
	var instance_ TypeDeclarationLike
	instance_ = ast.TypeDeclarationClass().Make(
		declaration,
		abstraction,
		optionalEnumeration,
	)
	return instance_
}

func TypeSection(
	typeDeclarations abs.Sequential[TypeDeclarationLike],
) TypeSectionLike {
	var instance_ TypeSectionLike
	instance_ = ast.TypeSectionClass().Make(
		typeDeclarations,
	)
	return instance_
}

func Value(
	name string,
	abstraction AbstractionLike,
) ValueLike {
	var instance_ ValueLike
	instance_ = ast.ValueClass().Make(
		name,
		abstraction,
	)
	return instance_
}

// Grammar

func Formatter() FormatterLike {
	var instance_ FormatterLike
	instance_ = gra.FormatterClass().Make()
	return instance_
}

func Parser() ParserLike {
	var instance_ ParserLike
	instance_ = gra.ParserClass().Make()
	return instance_
}

func Processor() ProcessorLike {
	var instance_ ProcessorLike
	instance_ = gra.ProcessorClass().Make()
	return instance_
}

func Scanner(
	source string,
	tokens abs.QueueLike[TokenLike],
) ScannerLike {
	var instance_ ScannerLike
	instance_ = gra.ScannerClass().Make(
		source,
		tokens,
	)
	return instance_
}

func Token(
	line uint,
	position uint,
	type_ TokenType,
	value string,
) TokenLike {
	var instance_ TokenLike
	instance_ = gra.TokenClass().Make(
		line,
		position,
		type_,
		value,
	)
	return instance_
}

func Validator() ValidatorLike {
	var instance_ ValidatorLike
	instance_ = gra.ValidatorClass().Make()
	return instance_
}

func Visitor(
	processor Methodical,
) VisitorLike {
	var instance_ VisitorLike
	instance_ = gra.VisitorClass().Make(
		processor,
	)
	return instance_
}

// GLOBAL FUNCTIONS
