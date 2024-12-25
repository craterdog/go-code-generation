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

		// Initialize the inherited aspects.
		Methodical: mod.Processor(),
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

func (v *packageAnalyzer_) GetPackageName() string {
	return v.packageName_
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

func (v *packageAnalyzer_) PreprocessAdditionalValue(
	additionalValue mod.AdditionalValueLike,
	index uint,
	size uint,
) {
	var name = additionalValue.GetName()
	v.enumeratedValues_.AppendValue(name)
}

func (v *packageAnalyzer_) PreprocessAspectDeclaration(
	aspectDeclaration mod.AspectDeclarationLike,
	index uint,
	size uint,
) {
	v.aspectDeclarations_.AppendValue(aspectDeclaration)
}

func (v *packageAnalyzer_) PreprocessClassDeclaration(
	classDeclaration mod.ClassDeclarationLike,
	index uint,
	size uint,
) {
	v.classDeclarations_.AppendValue(classDeclaration)
}

func (v *packageAnalyzer_) PreprocessFunctionalDeclaration(
	functionalDeclaration mod.FunctionalDeclarationLike,
	index uint,
	size uint,
) {
	v.functionalDeclarations_.AppendValue(functionalDeclaration)
}

func (v *packageAnalyzer_) PreprocessImportedPackage(
	importedPackage mod.ImportedPackageLike,
	index uint,
	size uint,
) {
	var packageAcronym = importedPackage.GetName()
	var packagePath = importedPackage.GetPath()
	v.importedPackages_.SetValue(packageAcronym, packagePath)
}

func (v *packageAnalyzer_) PreprocessInstanceDeclaration(
	instanceDeclaration mod.InstanceDeclarationLike,
	index uint,
	size uint,
) {
	v.instanceDeclarations_.AppendValue(instanceDeclaration)
}

func (v *packageAnalyzer_) PreprocessPackageDeclaration(
	packageDeclaration mod.PackageDeclarationLike,
) {
	v.legalNotice_ = packageDeclaration.GetLegalNotice().GetComment()
}

func (v *packageAnalyzer_) PostprocessPackageDeclaration(
	packageDeclaration mod.PackageDeclarationLike,
) {
	v.importedPackages_.SetValue(
		"fmt",
		`"fmt"`,
	)
	v.importedPackages_.SetValue(
		"uti",
		`"github.com/craterdog/go-missing-utilities/v2"`,
	)
	v.importedPackages_.SetValue(
		"col",
		`"github.com/craterdog/go-collection-framework/v5"`,
	)
	v.importedPackages_.SetValue(
		"abs",
		`"github.com/craterdog/go-collection-framework/v5/collection"`,
	)
	v.importedPackages_.SetValue(
		"syn",
		`"sync"`,
	)
}

func (v *packageAnalyzer_) PreprocessPackageHeader(
	packageHeader mod.PackageHeaderLike,
) {
	v.packageName_ = packageHeader.GetName()
}

func (v *packageAnalyzer_) PreprocessTypeDeclaration(
	typeDeclaration mod.TypeDeclarationLike,
	index uint,
	size uint,
) {
	v.typeDeclarations_.AppendValue(typeDeclaration)
}

func (v *packageAnalyzer_) PreprocessValue(
	value mod.ValueLike,
) {
	var name = value.GetName()
	v.enumeratedValues_.AppendValue(name)
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type packageAnalyzer_ struct {
	// Declare the instance attributes.
	visitor_                mod.VisitorLike
	legalNotice_            string
	packageName_            string
	importedPackages_       abs.CatalogLike[string, string]
	typeDeclarations_       abs.ListLike[mod.TypeDeclarationLike]
	enumeratedValues_       abs.ListLike[string]
	functionalDeclarations_ abs.ListLike[mod.FunctionalDeclarationLike]
	classDeclarations_      abs.ListLike[mod.ClassDeclarationLike]
	instanceDeclarations_   abs.ListLike[mod.InstanceDeclarationLike]
	aspectDeclarations_     abs.ListLike[mod.AspectDeclarationLike]

	// Declare the inherited aspects.
	mod.Methodical
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
