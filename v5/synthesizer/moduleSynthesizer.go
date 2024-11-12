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

package synthesizer

import (
	mod "github.com/craterdog/go-class-model/v5/ast"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ModuleSynthesizer() ModuleSynthesizerClassLike {
	return moduleSynthesizerReference()
}

// Constructor Methods

func (c *moduleSynthesizerClass_) Make(
	models abs.CatalogLike[string, mod.ModelLike],
) ModuleSynthesizerLike {
	var instance = &moduleSynthesizer_{
		// Initialize the instance attributes.
		models_: models,
	}
	return instance
}

// INSTANCE INTERFACE

// Primary Methods

func (v *moduleSynthesizer_) GetClass() ModuleSynthesizerClassLike {
	return moduleSynthesizerReference()
}

// TemplateDriven Methods

func (v *moduleSynthesizer_) CreateLegalNotice() string {
	var packageNames = v.models_.GetKeys().AsArray()
	var firstPackage = packageNames[0]
	var firstModel = v.models_.GetValue(firstPackage)
	var moduleDeclaration = firstModel.GetModuleDeclaration()
	var legalNotice = moduleDeclaration.GetLegalNotice().GetComment()
	return legalNotice
}

func (v *moduleSynthesizer_) CreateTypeAliases() string {
	var typeAliases string
	var models = v.models_.GetIterator()
	for models.HasNext() {
		var association = models.GetNext()
		var packageName = association.GetKey()
		var model = association.GetValue()
		var packageAliases = v.createPackageAliases(packageName, model)
		var packageAcronym = packageName[0:3]
		packageAliases = uti.ReplaceAll(
			packageAliases,
			"packageAcronym",
			packageAcronym,
		)
		typeAliases += packageAliases
	}
	return typeAliases
}

func (v *moduleSynthesizer_) CreateUniversalConstructors() string {
	return ""
}

func (v *moduleSynthesizer_) CreateGlobalFunctions() string {
	return ""
}

func (v *moduleSynthesizer_) PerformGlobalUpdates(
	source string,
) string {
	var moduleImports string
	var importedPackages = v.createImportedPackages(source)
	if uti.IsDefined(importedPackages) {
		moduleImports = moduleSynthesizerReference().moduleImports_
		moduleImports = uti.ReplaceAll(
			moduleImports,
			"importedPackages",
			importedPackages,
		)
	}
	source = uti.ReplaceAll(
		source,
		"moduleImports",
		moduleImports,
	)
	return source
}

// PROTECTED INTERFACE

// Private Methods

func (v *moduleSynthesizer_) createImportedPackages(
	source string,
) string {
	var importedPackages string
	return importedPackages
}

func (v *moduleSynthesizer_) createPackageAliases(
	packageName string,
	model mod.ModelLike,
) string {
	var typeAliases string
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classes = classSection.GetClassDeclarations().GetIterator()
	for classes.HasNext() {
		var class = classes.GetNext()
		var className = sts.TrimSuffix(class.GetDeclaration().GetName(), "ClassLike")
		var typeAlias = moduleSynthesizerReference().typeAlias_
		typeAlias = uti.ReplaceAll(
			typeAlias,
			"className",
			className,
		)
		typeAliases += typeAlias
	}
	var packageAliases = moduleSynthesizerReference().packageAliases_
	packageAliases = uti.ReplaceAll(
		packageAliases,
		"packageName",
		packageName,
	)
	packageAliases = uti.ReplaceAll(
		packageAliases,
		"typeAliases",
		typeAliases,
	)

	return packageAliases
}

// Instance Structure

type moduleSynthesizer_ struct {
	// Declare the instance attributes.
	models_ abs.CatalogLike[string, mod.ModelLike]
}

// Class Structure

type moduleSynthesizerClass_ struct {
	// Declare the class constants.
	moduleImports_  string
	packageAlias_   string
	packageAliases_ string
	typeAlias_      string
}

// Class Reference

func moduleSynthesizerReference() *moduleSynthesizerClass_ {
	return moduleSynthesizerReference_
}

var moduleSynthesizerReference_ = &moduleSynthesizerClass_{
	// Initialize the class constants.
	moduleImports_: `

import (<ImportedPackages>)`,

	packageAlias_: `
	<~packageAcronym> <packagePath>`,

	packageAliases_: `

// <~PackageName>

type (<TypeAliases>
)`,

	typeAlias_: `
	<~ClassName>Like = <~packageAcronym>.<~ClassName>Like`,
}
