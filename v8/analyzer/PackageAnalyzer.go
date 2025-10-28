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

package analyzer

import (
	mod "github.com/craterdog/go-class-model/v8"
	com "github.com/craterdog/go-essential-composites/v8"
)

// CLASS INTERFACE

// Access Function

func PackageAnalyzerClass() PackageAnalyzerClassLike {
	return packageAnalyzerClass()
}

// Constructor Methods

func (c *packageAnalyzerClass_) PackageAnalyzer(
	model mod.ModelLike,
) PackageAnalyzerLike {
	var instance = &packageAnalyzer_{
		// Initialize the instance attributes.
		importedPackages_:       com.Catalog[string, string](),
		typeDeclarations_:       com.List[mod.TypeDeclarationLike](),
		enumeratedValues_:       com.List[string](),
		functionalDeclarations_: com.List[mod.FunctionalDeclarationLike](),
		classDeclarations_:      com.List[mod.ClassDeclarationLike](),
		instanceDeclarations_:   com.List[mod.InstanceDeclarationLike](),
		aspectDeclarations_:     com.List[mod.AspectDeclarationLike](),

		// Initialize the inherited aspects.
		Methodical: mod.Processor(),
	}
	mod.Visitor(instance).VisitModel(model)
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *packageAnalyzer_) GetClass() PackageAnalyzerClassLike {
	return packageAnalyzerClass()
}

func (v *packageAnalyzer_) GetPackageName() string {
	return v.packageName_
}

func (v *packageAnalyzer_) GetLegalNotice() string {
	return v.legalNotice_
}

func (v *packageAnalyzer_) GetImportedPackages() com.CatalogLike[string, string] {
	return v.importedPackages_
}

func (v *packageAnalyzer_) GetTypeDeclarations() com.Sequential[mod.TypeDeclarationLike] {
	return v.typeDeclarations_
}

func (v *packageAnalyzer_) GetEnumeratedValues() com.Sequential[string] {
	return v.enumeratedValues_
}

func (v *packageAnalyzer_) GetFunctionalDeclarations() com.Sequential[mod.FunctionalDeclarationLike] {
	return v.functionalDeclarations_
}

func (v *packageAnalyzer_) GetClassDeclarations() com.Sequential[mod.ClassDeclarationLike] {
	return v.classDeclarations_
}

func (v *packageAnalyzer_) GetInstanceDeclarations() com.Sequential[mod.InstanceDeclarationLike] {
	return v.instanceDeclarations_
}

func (v *packageAnalyzer_) GetAspectDeclarations() com.Sequential[mod.AspectDeclarationLike] {
	return v.aspectDeclarations_
}

// Methodical Methods

func (v *packageAnalyzer_) PreprocessAdditionalValue(
	additionalValue mod.AdditionalValueLike,
	index uint,
	count uint,
) {
	var name = additionalValue.GetName()
	v.enumeratedValues_.AppendValue(name)
}

func (v *packageAnalyzer_) PreprocessAspectDeclaration(
	aspectDeclaration mod.AspectDeclarationLike,
	index uint,
	count uint,
) {
	v.aspectDeclarations_.AppendValue(aspectDeclaration)
}

func (v *packageAnalyzer_) PreprocessClassDeclaration(
	classDeclaration mod.ClassDeclarationLike,
	index uint,
	count uint,
) {
	v.classDeclarations_.AppendValue(classDeclaration)
}

func (v *packageAnalyzer_) PreprocessFunctionalDeclaration(
	functionalDeclaration mod.FunctionalDeclarationLike,
	index uint,
	count uint,
) {
	v.functionalDeclarations_.AppendValue(functionalDeclaration)
}

func (v *packageAnalyzer_) PreprocessImportedPackage(
	importedPackage mod.ImportedPackageLike,
	index uint,
	count uint,
) {
	var packageAcronym = importedPackage.GetName()
	var packagePath = importedPackage.GetPath()
	v.importedPackages_.SetValue(packageAcronym, packagePath)
}

func (v *packageAnalyzer_) PreprocessInstanceDeclaration(
	instanceDeclaration mod.InstanceDeclarationLike,
	index uint,
	count uint,
) {
	v.instanceDeclarations_.AppendValue(instanceDeclaration)
}

func (v *packageAnalyzer_) PreprocessPackageDeclaration(
	packageDeclaration mod.PackageDeclarationLike,
	index uint,
	count uint,
) {
	v.legalNotice_ = packageDeclaration.GetLegalNotice().GetComment()
}

func (v *packageAnalyzer_) PostprocessPackageDeclaration(
	packageDeclaration mod.PackageDeclarationLike,
	index uint,
	count uint,
) {
	v.importedPackages_.SetValue(
		"fmt",
		`"fmt"`,
	)
	v.importedPackages_.SetValue(
		"uti",
		`"github.com/craterdog/go-essential-utilities/v8"`,
	)
	v.importedPackages_.SetValue(
		"fra",
		`"github.com/craterdog/go-essential-composites/v8"`,
	)
	v.importedPackages_.SetValue(
		"syn",
		`"sync"`,
	)
}

func (v *packageAnalyzer_) PreprocessPackageHeader(
	packageHeader mod.PackageHeaderLike,
	index uint,
	count uint,
) {
	v.packageName_ = packageHeader.GetName()
}

func (v *packageAnalyzer_) PreprocessTypeDeclaration(
	typeDeclaration mod.TypeDeclarationLike,
	index uint,
	count uint,
) {
	v.typeDeclarations_.AppendValue(typeDeclaration)
}

func (v *packageAnalyzer_) PreprocessValue(
	value mod.ValueLike,
	index uint,
	count uint,
) {
	var name = value.GetName()
	v.enumeratedValues_.AppendValue(name)
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type packageAnalyzer_ struct {
	// Declare the instance attributes.
	legalNotice_            string
	packageName_            string
	importedPackages_       com.CatalogLike[string, string]
	typeDeclarations_       com.ListLike[mod.TypeDeclarationLike]
	enumeratedValues_       com.ListLike[string]
	functionalDeclarations_ com.ListLike[mod.FunctionalDeclarationLike]
	classDeclarations_      com.ListLike[mod.ClassDeclarationLike]
	instanceDeclarations_   com.ListLike[mod.InstanceDeclarationLike]
	aspectDeclarations_     com.ListLike[mod.AspectDeclarationLike]

	// Declare the inherited aspects.
	mod.Methodical
}

// Class Structure

type packageAnalyzerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func packageAnalyzerClass() *packageAnalyzerClass_ {
	return packageAnalyzerClassReference_
}

var packageAnalyzerClassReference_ = &packageAnalyzerClass_{
	// Initialize the class constants.
}
