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

package analyzer

import (
	mod "github.com/craterdog/go-class-model/v5"
	col "github.com/craterdog/go-collection-framework/v5"
	abs "github.com/craterdog/go-collection-framework/v5/collection"
)

// CLASS INTERFACE

// Access Function

func PackageAnalyzerClass() PackageAnalyzerClassLike {
	return packageAnalyzerClassReference()
}

// Constructor Methods

func (c *packageAnalyzerClass_) PackageAnalyzer(
	model mod.ModelLike,
) PackageAnalyzerLike {
	var instance = &packageAnalyzer_{
		// Initialize the instance attributes.
		importedPackages_:       col.Catalog[string, string](),
		typeDeclarations_:       col.List[mod.TypeDeclarationLike](),
		enumeratedValues_:       col.List[string](),
		functionalDeclarations_: col.List[mod.FunctionalDeclarationLike](),
		classDeclarations_:      col.List[mod.ClassDeclarationLike](),
		instanceDeclarations_:   col.List[mod.InstanceDeclarationLike](),
		aspectDeclarations_:     col.List[mod.AspectDeclarationLike](),
	}
	instance.visitor_ = mod.Visitor(instance)
	instance.visitor_.VisitModel(model)
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *packageAnalyzer_) GetClass() PackageAnalyzerClassLike {
	return packageAnalyzerClassReference()
}

func (v *packageAnalyzer_) GetLegalNotice() string {
	return v.legalNotice_
}

func (v *packageAnalyzer_) GetImportedPackages() abs.CatalogLike[string, string] {
	return v.importedPackages_
}

func (v *packageAnalyzer_) GetTypeDeclarations() abs.ListLike[mod.TypeDeclarationLike] {
	return v.typeDeclarations_
}

func (v *packageAnalyzer_) GetEnumeratedValues() abs.ListLike[string] {
	return v.enumeratedValues_
}

func (v *packageAnalyzer_) GetFunctionalDeclarations() abs.ListLike[mod.FunctionalDeclarationLike] {
	return v.functionalDeclarations_
}

func (v *packageAnalyzer_) GetClassDeclarations() abs.ListLike[mod.ClassDeclarationLike] {
	return v.classDeclarations_
}

func (v *packageAnalyzer_) GetInstanceDeclarations() abs.ListLike[mod.InstanceDeclarationLike] {
	return v.instanceDeclarations_
}

func (v *packageAnalyzer_) GetAspectDeclarations() abs.ListLike[mod.AspectDeclarationLike] {
	return v.aspectDeclarations_
}

// Methodical Methods

func (v *packageAnalyzer_) ProcessComment(
	comment string,
) {
}

func (v *packageAnalyzer_) ProcessName(
	name string,
) {
}

func (v *packageAnalyzer_) ProcessNewline(
	newline string,
) {
}

func (v *packageAnalyzer_) ProcessPath(
	path string,
) {
}

func (v *packageAnalyzer_) ProcessPrefix(
	prefix string,
) {
}

func (v *packageAnalyzer_) ProcessSpace(
	space string,
) {
}

func (v *packageAnalyzer_) PreprocessAbstraction(
	abstraction mod.AbstractionLike,
) {
}

func (v *packageAnalyzer_) ProcessAbstractionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessAbstraction(
	abstraction mod.AbstractionLike,
) {
}

func (v *packageAnalyzer_) PreprocessAdditionalArgument(
	additionalArgument mod.AdditionalArgumentLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) ProcessAdditionalArgumentSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessAdditionalArgument(
	additionalArgument mod.AdditionalArgumentLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessAdditionalConstraint(
	additionalConstraint mod.AdditionalConstraintLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) ProcessAdditionalConstraintSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessAdditionalConstraint(
	additionalConstraint mod.AdditionalConstraintLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessAdditionalValue(
	additionalValue mod.AdditionalValueLike,
	index uint,
	size uint,
) {
	var name = additionalValue.GetName()
	v.enumeratedValues_.AppendValue(name)
}

func (v *packageAnalyzer_) ProcessAdditionalValueSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessAdditionalValue(
	additionalValue mod.AdditionalValueLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessArgument(
	argument mod.ArgumentLike,
) {
}

func (v *packageAnalyzer_) ProcessArgumentSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessArgument(
	argument mod.ArgumentLike,
) {
}

func (v *packageAnalyzer_) PreprocessArguments(
	arguments mod.ArgumentsLike,
) {
}

func (v *packageAnalyzer_) ProcessArgumentsSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessArguments(
	arguments mod.ArgumentsLike,
) {
}

func (v *packageAnalyzer_) PreprocessArray(
	array mod.ArrayLike,
) {
}

func (v *packageAnalyzer_) ProcessArraySlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessArray(
	array mod.ArrayLike,
) {
}

func (v *packageAnalyzer_) PreprocessAspectDeclaration(
	aspectDeclaration mod.AspectDeclarationLike,
	index uint,
	size uint,
) {
	v.aspectDeclarations_.AppendValue(aspectDeclaration)
}

func (v *packageAnalyzer_) ProcessAspectDeclarationSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessAspectDeclaration(
	aspectDeclaration mod.AspectDeclarationLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessAspectInterface(
	aspectInterface mod.AspectInterfaceLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) ProcessAspectInterfaceSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessAspectInterface(
	aspectInterface mod.AspectInterfaceLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessAspectMethod(
	aspectMethod mod.AspectMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) ProcessAspectMethodSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessAspectMethod(
	aspectMethod mod.AspectMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessAspectSection(
	aspectSection mod.AspectSectionLike,
) {
}

func (v *packageAnalyzer_) ProcessAspectSectionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessAspectSection(
	aspectSection mod.AspectSectionLike,
) {
}

func (v *packageAnalyzer_) PreprocessAspectSubsection(
	aspectSubsection mod.AspectSubsectionLike,
) {
}

func (v *packageAnalyzer_) ProcessAspectSubsectionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessAspectSubsection(
	aspectSubsection mod.AspectSubsectionLike,
) {
}

func (v *packageAnalyzer_) PreprocessAttributeMethod(
	attributeMethod mod.AttributeMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) ProcessAttributeMethodSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessAttributeMethod(
	attributeMethod mod.AttributeMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessAttributeSubsection(
	attributeSubsection mod.AttributeSubsectionLike,
) {
}

func (v *packageAnalyzer_) ProcessAttributeSubsectionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessAttributeSubsection(
	attributeSubsection mod.AttributeSubsectionLike,
) {
}

func (v *packageAnalyzer_) PreprocessChannel(
	channel mod.ChannelLike,
) {
}

func (v *packageAnalyzer_) ProcessChannelSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessChannel(
	channel mod.ChannelLike,
) {
}

func (v *packageAnalyzer_) PreprocessClassDeclaration(
	classDeclaration mod.ClassDeclarationLike,
	index uint,
	size uint,
) {
	v.classDeclarations_.AppendValue(classDeclaration)
}

func (v *packageAnalyzer_) ProcessClassDeclarationSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessClassDeclaration(
	classDeclaration mod.ClassDeclarationLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessClassMethods(
	classMethods mod.ClassMethodsLike,
) {
}

func (v *packageAnalyzer_) ProcessClassMethodsSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessClassMethods(
	classMethods mod.ClassMethodsLike,
) {
}

func (v *packageAnalyzer_) PreprocessClassSection(
	classSection mod.ClassSectionLike,
) {
}

func (v *packageAnalyzer_) ProcessClassSectionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessClassSection(
	classSection mod.ClassSectionLike,
) {
}

func (v *packageAnalyzer_) PreprocessConstantMethod(
	constantMethod mod.ConstantMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) ProcessConstantMethodSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessConstantMethod(
	constantMethod mod.ConstantMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessConstantSubsection(
	constantSubsection mod.ConstantSubsectionLike,
) {
}

func (v *packageAnalyzer_) ProcessConstantSubsectionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessConstantSubsection(
	constantSubsection mod.ConstantSubsectionLike,
) {
}

func (v *packageAnalyzer_) PreprocessConstraint(
	constraint mod.ConstraintLike,
) {
}

func (v *packageAnalyzer_) ProcessConstraintSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessConstraint(
	constraint mod.ConstraintLike,
) {
}

func (v *packageAnalyzer_) PreprocessConstraints(
	constraints mod.ConstraintsLike,
) {
}

func (v *packageAnalyzer_) ProcessConstraintsSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessConstraints(
	constraints mod.ConstraintsLike,
) {
}

func (v *packageAnalyzer_) PreprocessConstructorMethod(
	constructorMethod mod.ConstructorMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) ProcessConstructorMethodSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessConstructorMethod(
	constructorMethod mod.ConstructorMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessConstructorSubsection(
	constructorSubsection mod.ConstructorSubsectionLike,
) {
}

func (v *packageAnalyzer_) ProcessConstructorSubsectionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessConstructorSubsection(
	constructorSubsection mod.ConstructorSubsectionLike,
) {
}

func (v *packageAnalyzer_) PreprocessDeclaration(
	declaration mod.DeclarationLike,
) {
}

func (v *packageAnalyzer_) ProcessDeclarationSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessDeclaration(
	declaration mod.DeclarationLike,
) {
}

func (v *packageAnalyzer_) PreprocessEnumeration(
	enumeration mod.EnumerationLike,
) {
}

func (v *packageAnalyzer_) ProcessEnumerationSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessEnumeration(
	enumeration mod.EnumerationLike,
) {
}

func (v *packageAnalyzer_) PreprocessFunctionMethod(
	functionMethod mod.FunctionMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) ProcessFunctionMethodSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessFunctionMethod(
	functionMethod mod.FunctionMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessFunctionSubsection(
	functionSubsection mod.FunctionSubsectionLike,
) {
}

func (v *packageAnalyzer_) ProcessFunctionSubsectionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessFunctionSubsection(
	functionSubsection mod.FunctionSubsectionLike,
) {
}

func (v *packageAnalyzer_) PreprocessFunctionalDeclaration(
	functionalDeclaration mod.FunctionalDeclarationLike,
	index uint,
	size uint,
) {
	v.functionalDeclarations_.AppendValue(functionalDeclaration)
}

func (v *packageAnalyzer_) ProcessFunctionalDeclarationSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessFunctionalDeclaration(
	functionalDeclaration mod.FunctionalDeclarationLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessFunctionalSection(
	functionalSection mod.FunctionalSectionLike,
) {
}

func (v *packageAnalyzer_) ProcessFunctionalSectionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessFunctionalSection(
	functionalSection mod.FunctionalSectionLike,
) {
}

func (v *packageAnalyzer_) PreprocessGetterMethod(
	getterMethod mod.GetterMethodLike,
) {
}

func (v *packageAnalyzer_) ProcessGetterMethodSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessGetterMethod(
	getterMethod mod.GetterMethodLike,
) {
}

func (v *packageAnalyzer_) PreprocessImportedPackage(
	importedPackage mod.ImportedPackageLike,
	index uint,
	size uint,
) {
	var packagePath = importedPackage.GetPath()
	packagePath = packagePath[1 : len(packagePath)-1]
	var packageAcronym = importedPackage.GetName()
	v.importedPackages_.SetValue(packagePath, packageAcronym)
}

func (v *packageAnalyzer_) ProcessImportedPackageSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessImportedPackage(
	importedPackage mod.ImportedPackageLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessInstanceDeclaration(
	instanceDeclaration mod.InstanceDeclarationLike,
	index uint,
	size uint,
) {
	v.instanceDeclarations_.AppendValue(instanceDeclaration)
}

func (v *packageAnalyzer_) ProcessInstanceDeclarationSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessInstanceDeclaration(
	instanceDeclaration mod.InstanceDeclarationLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessInstanceMethods(
	instanceMethods mod.InstanceMethodsLike,
) {
}

func (v *packageAnalyzer_) ProcessInstanceMethodsSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessInstanceMethods(
	instanceMethods mod.InstanceMethodsLike,
) {
}

func (v *packageAnalyzer_) PreprocessInstanceSection(
	instanceSection mod.InstanceSectionLike,
) {
}

func (v *packageAnalyzer_) ProcessInstanceSectionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessInstanceSection(
	instanceSection mod.InstanceSectionLike,
) {
}

func (v *packageAnalyzer_) PreprocessInterfaceDeclarations(
	interfaceDeclarations mod.InterfaceDeclarationsLike,
) {
}

func (v *packageAnalyzer_) ProcessInterfaceDeclarationsSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessInterfaceDeclarations(
	interfaceDeclarations mod.InterfaceDeclarationsLike,
) {
}

func (v *packageAnalyzer_) PreprocessLegalNotice(
	legalNotice mod.LegalNoticeLike,
) {
}

func (v *packageAnalyzer_) ProcessLegalNoticeSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessLegalNotice(
	legalNotice mod.LegalNoticeLike,
) {
}

func (v *packageAnalyzer_) PreprocessMap(
	map_ mod.MapLike,
) {
}

func (v *packageAnalyzer_) ProcessMapSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessMap(
	map_ mod.MapLike,
) {
}

func (v *packageAnalyzer_) PreprocessMethod(
	method mod.MethodLike,
) {
}

func (v *packageAnalyzer_) ProcessMethodSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessMethod(
	method mod.MethodLike,
) {
}

func (v *packageAnalyzer_) PreprocessModel(
	model mod.ModelLike,
) {
}

func (v *packageAnalyzer_) ProcessModelSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessModel(
	model mod.ModelLike,
) {
}

func (v *packageAnalyzer_) PreprocessMultivalue(
	multivalue mod.MultivalueLike,
) {
}

func (v *packageAnalyzer_) ProcessMultivalueSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessMultivalue(
	multivalue mod.MultivalueLike,
) {
}

func (v *packageAnalyzer_) PreprocessNone(
	none mod.NoneLike,
) {
}

func (v *packageAnalyzer_) ProcessNoneSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessNone(
	none mod.NoneLike,
) {
}

func (v *packageAnalyzer_) PreprocessPackageDeclaration(
	packageDeclaration mod.PackageDeclarationLike,
) {
	v.legalNotice_ = packageDeclaration.GetLegalNotice().GetComment()
}

func (v *packageAnalyzer_) ProcessPackageDeclarationSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessPackageDeclaration(
	packageDeclaration mod.PackageDeclarationLike,
) {
	v.importedPackages_.SetValue(
		"fmt",
		"fmt",
	)
	v.importedPackages_.SetValue(
		"github.com/craterdog/go-missing-utilities/v2",
		"uti",
	)
	v.importedPackages_.SetValue(
		"github.com/craterdog/go-collection-framework/v5",
		"col",
	)
	v.importedPackages_.SetValue(
		"github.com/craterdog/go-collection-framework/v5/collection",
		"abs",
	)
	v.importedPackages_.SetValue(
		"sync",
		"syn",
	)
}

func (v *packageAnalyzer_) PreprocessPackageHeader(
	packageHeader mod.PackageHeaderLike,
) {
}

func (v *packageAnalyzer_) ProcessPackageHeaderSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessPackageHeader(
	packageHeader mod.PackageHeaderLike,
) {
}

func (v *packageAnalyzer_) PreprocessPackageImports(
	packageImports mod.PackageImportsLike,
) {
}

func (v *packageAnalyzer_) ProcessPackageImportsSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessPackageImports(
	packageImports mod.PackageImportsLike,
) {
}

func (v *packageAnalyzer_) PreprocessParameter(
	parameter mod.ParameterLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) ProcessParameterSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessParameter(
	parameter mod.ParameterLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessPrimitiveDeclarations(
	primitiveDeclarations mod.PrimitiveDeclarationsLike,
) {
}

func (v *packageAnalyzer_) ProcessPrimitiveDeclarationsSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessPrimitiveDeclarations(
	primitiveDeclarations mod.PrimitiveDeclarationsLike,
) {
}

func (v *packageAnalyzer_) PreprocessPrincipalMethod(
	principalMethod mod.PrincipalMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) ProcessPrincipalMethodSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessPrincipalMethod(
	principalMethod mod.PrincipalMethodLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessPrincipalSubsection(
	principalSubsection mod.PrincipalSubsectionLike,
) {
}

func (v *packageAnalyzer_) ProcessPrincipalSubsectionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessPrincipalSubsection(
	principalSubsection mod.PrincipalSubsectionLike,
) {
}

func (v *packageAnalyzer_) PreprocessResult(
	result mod.ResultLike,
) {
}

func (v *packageAnalyzer_) ProcessResultSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessResult(
	result mod.ResultLike,
) {
}

func (v *packageAnalyzer_) PreprocessSetterMethod(
	setterMethod mod.SetterMethodLike,
) {
}

func (v *packageAnalyzer_) ProcessSetterMethodSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessSetterMethod(
	setterMethod mod.SetterMethodLike,
) {
}

func (v *packageAnalyzer_) PreprocessTypeDeclaration(
	typeDeclaration mod.TypeDeclarationLike,
	index uint,
	size uint,
) {
	v.typeDeclarations_.AppendValue(typeDeclaration)
}

func (v *packageAnalyzer_) ProcessTypeDeclarationSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessTypeDeclaration(
	typeDeclaration mod.TypeDeclarationLike,
	index uint,
	size uint,
) {
}

func (v *packageAnalyzer_) PreprocessTypeSection(
	typeSection mod.TypeSectionLike,
) {
}

func (v *packageAnalyzer_) ProcessTypeSectionSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessTypeSection(
	typeSection mod.TypeSectionLike,
) {
}

func (v *packageAnalyzer_) PreprocessValue(
	value mod.ValueLike,
) {
	var name = value.GetName()
	v.enumeratedValues_.AppendValue(name)
}

func (v *packageAnalyzer_) ProcessValueSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessValue(
	value mod.ValueLike,
) {
}

func (v *packageAnalyzer_) PreprocessWrapper(
	wrapper mod.WrapperLike,
) {
}

func (v *packageAnalyzer_) ProcessWrapperSlot(
	slot uint,
) {
}

func (v *packageAnalyzer_) PostprocessWrapper(
	wrapper mod.WrapperLike,
) {
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type packageAnalyzer_ struct {
	// Declare the instance attributes.
	visitor_                mod.VisitorLike
	legalNotice_            string
	importedPackages_       abs.CatalogLike[string, string]
	typeDeclarations_       abs.ListLike[mod.TypeDeclarationLike]
	enumeratedValues_       abs.ListLike[string]
	functionalDeclarations_ abs.ListLike[mod.FunctionalDeclarationLike]
	classDeclarations_      abs.ListLike[mod.ClassDeclarationLike]
	instanceDeclarations_   abs.ListLike[mod.InstanceDeclarationLike]
	aspectDeclarations_     abs.ListLike[mod.AspectDeclarationLike]
}

// Class Structure

type packageAnalyzerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func packageAnalyzerClassReference() *packageAnalyzerClass_ {
	return packageAnalyzerClassReference_
}

var packageAnalyzerClassReference_ = &packageAnalyzerClass_{
	// Initialize the class constants.
}
