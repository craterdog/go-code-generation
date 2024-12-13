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
│                 This class file was automatically generated.                 │
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	ast "github.com/craterdog/go-class-model/v5/ast"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func FormatterClass() FormatterClassLike {
	return formatterClassReference()
}

// Constructor Methods

func (c *formatterClass_) Formatter() FormatterLike {
	var instance = &formatter_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *formatter_) GetClass() FormatterClassLike {
	return formatterClassReference()
}

func (v *formatter_) FormatModel(model ast.ModelLike) string {
	v.visitor_.VisitModel(model)
	return v.getResult()
}

// Methodical Methods

func (v *formatter_) ProcessComment(
	comment string,
) {
	v.appendString(comment)
}

func (v *formatter_) ProcessName(
	name string,
) {
	v.appendString(name)
}

func (v *formatter_) ProcessNewline(
	newline string,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessPath(
	path string,
) {
	v.appendString(path)
}

func (v *formatter_) ProcessSpace(
	space string,
) {
	v.appendString(space)
}

func (v *formatter_) PreprocessAbstraction(
	abstraction ast.AbstractionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAbstractionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAbstraction(
	abstraction ast.AbstractionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAdditionalArgumentSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAdditionalConstraintSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAdditionalConstraint(
	additionalConstraint ast.AdditionalConstraintLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAdditionalValueSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessArgument(
	argument ast.ArgumentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessArgumentSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessArgument(
	argument ast.ArgumentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessArguments(
	arguments ast.ArgumentsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessArgumentsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessArguments(
	arguments ast.ArgumentsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessArray(
	array ast.ArrayLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessArraySlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessArray(
	array ast.ArrayLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAspectDeclarationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAspectDeclaration(
	aspectDeclaration ast.AspectDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAspectInterfaceSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAspectInterface(
	aspectInterface ast.AspectInterfaceLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAspectMethodSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAspectMethod(
	aspectMethod ast.AspectMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAspectSection(
	aspectSection ast.AspectSectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAspectSectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAspectSection(
	aspectSection ast.AspectSectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAspectSubsectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAspectSubsection(
	aspectSubsection ast.AspectSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAttributeMethodSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAttributeMethod(
	attributeMethod ast.AttributeMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAttributeSubsectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAttributeSubsection(
	attributeSubsection ast.AttributeSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessChannel(
	channel ast.ChannelLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessChannelSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessChannel(
	channel ast.ChannelLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessClassDeclarationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessClassDeclaration(
	classDeclaration ast.ClassDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessClassMethods(
	classMethods ast.ClassMethodsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessClassMethodsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessClassMethods(
	classMethods ast.ClassMethodsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessClassSection(
	classSection ast.ClassSectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessClassSectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessClassSection(
	classSection ast.ClassSectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessConstantMethodSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessConstantMethod(
	constantMethod ast.ConstantMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessConstantSubsectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessConstantSubsection(
	constantSubsection ast.ConstantSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessConstraint(
	constraint ast.ConstraintLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessConstraintSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessConstraint(
	constraint ast.ConstraintLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessConstraints(
	constraints ast.ConstraintsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessConstraintsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessConstraints(
	constraints ast.ConstraintsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessConstructorMethodSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessConstructorMethod(
	constructorMethod ast.ConstructorMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessConstructorSubsectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessConstructorSubsection(
	constructorSubsection ast.ConstructorSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessDeclaration(
	declaration ast.DeclarationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessDeclarationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessDeclaration(
	declaration ast.DeclarationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessEnumeration(
	enumeration ast.EnumerationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessEnumerationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessEnumeration(
	enumeration ast.EnumerationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessFunctionMethodSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessFunctionMethod(
	functionMethod ast.FunctionMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessFunctionSubsectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessFunctionSubsection(
	functionSubsection ast.FunctionSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessFunctionalDeclarationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessFunctionalDeclaration(
	functionalDeclaration ast.FunctionalDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessFunctionalSectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessFunctionalSection(
	functionalSection ast.FunctionalSectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessGetterMethodSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessGetterMethod(
	getterMethod ast.GetterMethodLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessImportedPackage(
	importedPackage ast.ImportedPackageLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessImportedPackageSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessImportedPackage(
	importedPackage ast.ImportedPackageLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInstanceDeclarationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInstanceDeclaration(
	instanceDeclaration ast.InstanceDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInstanceMethods(
	instanceMethods ast.InstanceMethodsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInstanceMethodsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInstanceMethods(
	instanceMethods ast.InstanceMethodsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInstanceSection(
	instanceSection ast.InstanceSectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInstanceSectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInstanceSection(
	instanceSection ast.InstanceSectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInterfaceDeclarationsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInterfaceDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessLegalNoticeSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMap(
	map_ ast.MapLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMapSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMap(
	map_ ast.MapLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMethod(
	method ast.MethodLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMethodSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMethod(
	method ast.MethodLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessModel(
	model ast.ModelLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessModelSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessModel(
	model ast.ModelLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMultivalue(
	multivalue ast.MultivalueLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMultivalueSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMultivalue(
	multivalue ast.MultivalueLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessNone(
	none ast.NoneLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessNoneSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessNone(
	none ast.NoneLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPackageDeclarationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessPackageHeader(
	packageHeader ast.PackageHeaderLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPackageHeaderSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPackageHeader(
	packageHeader ast.PackageHeaderLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessPackageImports(
	packageImports ast.PackageImportsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPackageImportsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPackageImports(
	packageImports ast.PackageImportsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessParameterSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessParameter(
	parameter ast.ParameterLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessPrefix(
	prefix ast.PrefixLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPrefixSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPrefix(
	prefix ast.PrefixLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPrimitiveDeclarationsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPrimitiveDeclarations(
	primitiveDeclarations ast.PrimitiveDeclarationsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessPrincipalMethod(
	principalMethod ast.PrincipalMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPrincipalMethodSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPrincipalMethod(
	principalMethod ast.PrincipalMethodLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPrincipalSubsectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPrincipalSubsection(
	principalSubsection ast.PrincipalSubsectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessResult(
	result ast.ResultLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessResultSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessResult(
	result ast.ResultLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessSetterMethodSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessSetterMethod(
	setterMethod ast.SetterMethodLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessSuffix(
	suffix ast.SuffixLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessSuffixSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessSuffix(
	suffix ast.SuffixLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessTypeDeclarationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessTypeDeclaration(
	typeDeclaration ast.TypeDeclarationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessTypeSection(
	typeSection ast.TypeSectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessTypeSectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessTypeSection(
	typeSection ast.TypeSectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessValue(
	value ast.ValueLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessValueSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessValue(
	value ast.ValueLike,
) {
	// TBD - Add formatting of the delimited rule.
}

// PROTECTED INTERFACE

// Private Methods

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "\t"
	var level uint
	for ; level < v.depth_; level++ {
		newline += indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(
	string_ string,
) {
	v.result_.WriteString(string_)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}

// Instance Structure

type formatter_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type formatterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func formatterClassReference() *formatterClass_ {
	return formatterClassReference_
}

var formatterClassReference_ = &formatterClass_{
	// Initialize the class constants.
}
