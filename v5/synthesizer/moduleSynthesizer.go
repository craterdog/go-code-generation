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
	mod "github.com/craterdog/go-class-model/v5"
	ana "github.com/craterdog/go-code-generation/v5/analyzer"
	col "github.com/craterdog/go-collection-framework/v5"
	abs "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	reg "regexp"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ModuleSynthesizerClass() ModuleSynthesizerClassLike {
	return moduleSynthesizerClassReference()
}

// Constructor Methods

func (c *moduleSynthesizerClass_) ModuleSynthesizer(
	models abs.CatalogLike[string, mod.ModelLike],
) ModuleSynthesizerLike {
	var instance = &moduleSynthesizer_{
		// Initialize the instance attributes.
		models_: models,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *moduleSynthesizer_) GetClass() ModuleSynthesizerClassLike {
	return moduleSynthesizerClassReference()
}

// TemplateDriven Methods

func (v *moduleSynthesizer_) CreateLegalNotice() string {
	var models = v.models_.GetIterator()
	var model = models.GetNext().GetValue()
	var packageDeclaration = model.GetPackageDeclaration()
	var legalNotice = packageDeclaration.GetLegalNotice().GetComment()
	return legalNotice
}

func (v *moduleSynthesizer_) CreateWarningMessage() string {
	var class = moduleSynthesizerClassReference()
	var warningMessage = class.warningMessage_
	return warningMessage
}

func (v *moduleSynthesizer_) CreateTypeAliases() string {
	var typeAliases string
	var models = v.models_.GetIterator()
	for models.HasNext() {
		var association = models.GetNext()
		var packageName = association.GetKey()
		var model = association.GetValue()
		var packageAliases = v.createPackageAliases(
			packageName,
			model,
		)
		typeAliases += packageAliases
	}
	return typeAliases
}

func (v *moduleSynthesizer_) CreateClassConstructors() string {
	var classConstructors string
	var models = v.models_.GetIterator()
	for models.HasNext() {
		var association = models.GetNext()
		var packageName = association.GetKey()
		var model = association.GetValue()
		var constructorFunctions = v.createConstructorFunctions(
			packageName,
			model,
		)
		classConstructors += constructorFunctions
	}
	return classConstructors
}

func (v *moduleSynthesizer_) PerformGlobalUpdates(
	existing string,
	generated string,
) string {
	generated = v.preserveExistingCode(existing, generated)
	generated = v.updateImportedPackages(existing, generated)
	return generated
}

// PROTECTED INTERFACE

// Private Methods

func (v *moduleSynthesizer_) createAspectAliases(
	aspectDeclarations abs.ListLike[mod.AspectDeclarationLike],
) (
	aspectAliases string,
) {
	var nameAliases string
	var aspects = aspectDeclarations.GetIterator()
	for aspects.HasNext() {
		var declaration = aspects.GetNext().GetDeclaration()
		var nameAlias = v.createNameAlias(declaration)
		nameAliases += nameAlias
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClassReference()
		aspectAliases = class.typeAliases_
		aspectAliases = uti.ReplaceAll(
			aspectAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createClassAliases(
	classDeclarations abs.ListLike[mod.ClassDeclarationLike],
) (
	classAliases string,
) {
	var nameAliases string
	var classes = classDeclarations.GetIterator()
	for classes.HasNext() {
		var declaration = classes.GetNext().GetDeclaration()
		nameAliases += v.createNameAlias(declaration)
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClassReference()
		classAliases = class.typeAliases_
		classAliases = uti.ReplaceAll(
			classAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createInstanceAliases(
	instanceDeclarations abs.ListLike[mod.InstanceDeclarationLike],
) (
	instanceAliases string,
) {
	var nameAliases string
	var instances = instanceDeclarations.GetIterator()
	for instances.HasNext() {
		var declaration = instances.GetNext().GetDeclaration()
		nameAliases += v.createNameAlias(declaration)
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClassReference()
		instanceAliases = class.typeAliases_
		instanceAliases = uti.ReplaceAll(
			instanceAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createConstructorFunction(
	constructorMethod mod.ConstructorMethodLike,
	className string,
	model mod.ModelLike,
) string {
	var class = moduleSynthesizerClassReference()
	var constructorFunction = class.constructorFunction_
	var methodName = constructorMethod.GetName()
	var parameters = v.extractParameters(constructorMethod, model)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"methodName",
		methodName,
	)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"parameters",
		parameters,
	)
	var parameterNames = v.extractParameterNames(constructorMethod)
	constructorFunction = uti.ReplaceAll(
		constructorFunction,
		"parameterNames",
		parameterNames,
	)
	return constructorFunction
}

func (v *moduleSynthesizer_) createClassConstructors(
	model mod.ModelLike,
	classDeclaration mod.ClassDeclarationLike,
) string {
	var constructors string
	var className = sts.TrimSuffix(
		classDeclaration.GetDeclaration().GetName(),
		"ClassLike",
	)
	className = uti.MakeLowerCase(className)
	var analyzerClass = ana.ClassAnalyzerClass()
	var analyzer = analyzerClass.ClassAnalyzer(model, className)
	var constructorMethods = analyzer.GetConstructorMethods().GetIterator()
	for constructorMethods.HasNext() {
		var constructorMethod = constructorMethods.GetNext()
		var constructorFunction = v.createConstructorFunction(
			constructorMethod,
			className,
			model,
		)
		constructors += constructorFunction
	}
	var class = moduleSynthesizerClassReference()
	var classConstructors = class.classConstructors_
	classConstructors = uti.ReplaceAll(
		classConstructors,
		"constructors",
		constructors,
	)
	classConstructors = uti.ReplaceAll(
		classConstructors,
		"className",
		className,
	)
	var constraints = analyzer.GetTypeConstraints()
	var arguments = analyzer.GetTypeArguments()
	classConstructors = uti.ReplaceAll(
		classConstructors,
		"constraints",
		constraints,
	)
	classConstructors = uti.ReplaceAll(
		classConstructors,
		"arguments",
		arguments,
	)

	return classConstructors
}

func (v *moduleSynthesizer_) createConstructorFunctions(
	packageName string,
	model mod.ModelLike,
) string {
	var constructorFunctions string
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	var classSection = interfaceDeclarations.GetClassSection()
	var classes = classSection.GetClassDeclarations().GetIterator()
	for classes.HasNext() {
		var class = classes.GetNext()
		var constructors = v.createClassConstructors(model, class)
		constructorFunctions += constructors
	}
	constructorFunctions = v.replacePackageAttributes(
		constructorFunctions,
		packageName,
	)
	return constructorFunctions
}

func (v *moduleSynthesizer_) createEnumeratedAliases(
	enumeratedValues abs.ListLike[string],
) (
	enumeratedAliases string,
) {
	var nameAliases string
	var class = moduleSynthesizerClassReference()
	var names = enumeratedValues.GetIterator()
	for names.HasNext() {
		var name = names.GetNext()
		var nameAlias = class.nameAlias_
		nameAlias = uti.ReplaceAll(
			nameAlias,
			"name",
			name,
		)
		nameAliases += nameAlias
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		enumeratedAliases = class.enumeratedAliases_
		enumeratedAliases = uti.ReplaceAll(
			enumeratedAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createFunctionalAliases(
	functionalDeclarations abs.ListLike[mod.FunctionalDeclarationLike],
) (
	functionalAliases string,
) {
	var nameAliases string
	var functionals = functionalDeclarations.GetIterator()
	for functionals.HasNext() {
		var declaration = functionals.GetNext().GetDeclaration()
		var nameAlias = v.createNameAlias(declaration)
		nameAliases += nameAlias
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClassReference()
		functionalAliases = class.typeAliases_
		functionalAliases = uti.ReplaceAll(
			functionalAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) createImportedPath(
	packageAcronym string,
	packagePath string,
) string {
	var class = moduleSynthesizerClassReference()
	var importedPath = class.importedPath_
	importedPath = uti.ReplaceAll(
		importedPath,
		"packageAcronym",
		packageAcronym,
	)
	importedPath = uti.ReplaceAll(
		importedPath,
		"packagePath",
		packagePath,
	)
	return importedPath
}

func (v *moduleSynthesizer_) createNameAlias(
	declaration mod.DeclarationLike,
) (
	nameAlias string,
) {
	var name = declaration.GetName()
	if uti.IsDefined(declaration.GetOptionalConstraints()) {
		// Type aliases are not supported for generic types in Go.
		return
	}
	var class = moduleSynthesizerClassReference()
	nameAlias = class.nameAlias_
	nameAlias = uti.ReplaceAll(
		nameAlias,
		"name",
		name,
	)
	return
}

func (v *moduleSynthesizer_) createPackageAliases(
	packageName string,
	model mod.ModelLike,
) string {
	var analyzer = ana.PackageAnalyzerClass().PackageAnalyzer(model)
	var typeAliases = v.createTypeAliases(analyzer.GetTypeDeclarations())
	typeAliases += v.createEnumeratedAliases(analyzer.GetEnumeratedValues())
	typeAliases += v.createFunctionalAliases(analyzer.GetFunctionalDeclarations())
	typeAliases += v.createClassAliases(analyzer.GetClassDeclarations())
	typeAliases += v.createInstanceAliases(analyzer.GetInstanceDeclarations())
	typeAliases += v.createAspectAliases(analyzer.GetAspectDeclarations())
	var class = moduleSynthesizerClassReference()
	var packageAliases = class.packageAliases_
	packageAliases = uti.ReplaceAll(
		packageAliases,
		"typeAliases",
		typeAliases,
	)
	packageAliases = v.replacePackageAttributes(
		packageAliases,
		packageName,
	)
	return packageAliases
}

func (v *moduleSynthesizer_) createTypeAliases(
	typeDeclarations abs.ListLike[mod.TypeDeclarationLike],
) (
	typeAliases string,
) {
	var nameAliases string
	var types = typeDeclarations.GetIterator()
	for types.HasNext() {
		var typeDeclaration = types.GetNext().GetDeclaration()
		nameAliases += v.createNameAlias(typeDeclaration)
	}
	if uti.IsDefined(nameAliases) {
		nameAliases += "\n"
		var class = moduleSynthesizerClassReference()
		typeAliases = class.typeAliases_
		typeAliases = uti.ReplaceAll(
			typeAliases,
			"nameAliases",
			nameAliases,
		)
	}
	return
}

func (v *moduleSynthesizer_) extractImportedPackages(
	source string,
) (
	packages abs.CatalogLike[string, string],
) {
	packages = col.Catalog[string, string]()
	var lower_ = `\p{Ll}`
	var digit_ = `\p{Nd}`
	var acronym_ = `(` + lower_ + `(?:` + lower_ + `|` + digit_ + `){2})`
	var white_ = `[ \t\r\n]*`
	var path_ = `("[^"]+")`
	var pattern = `import \((?:.|\r?\n)*?\)`
	var matcher = reg.MustCompile(pattern)
	var imports = matcher.FindString(source)
	var lines = sts.Split(imports, "\n")
	var count = len(lines)
	if count > 2 {
		pattern = acronym_ + white_ + path_
		matcher = reg.MustCompile(pattern)
		lines = lines[1 : count-1]
		for _, line := range lines {
			var matches = matcher.FindStringSubmatch(line)
			var packageAcronym = matches[1]
			var packagePath = matches[2]
			packages.SetValue(packageAcronym, packagePath)
		}
	}
	return
}

func (v *moduleSynthesizer_) extractParameterNames(
	constructorMethod mod.ConstructorMethodLike,
) string {
	var parameterNames string
	var parameters = constructorMethod.GetParameters().GetIterator()
	if parameters.IsEmpty() {
		return parameterNames
	}
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var parameterName = parameter.GetName()
		parameterNames += "\n\t\t" + parameterName + ","
	}
	parameterNames += "\n\t"
	return parameterNames
}

func (v *moduleSynthesizer_) extractParameters(
	constructorMethod mod.ConstructorMethodLike,
	model mod.ModelLike,
) string {
	var methodParameters string
	var parameters = constructorMethod.GetParameters().GetIterator()
	for parameters.HasNext() {
		var parameter = parameters.GetNext()
		var parameterName = parameter.GetName()
		var parameterType = v.extractType(
			parameter.GetAbstraction(),
			model,
		)
		var class = moduleSynthesizerClassReference()
		var methodParameter = class.methodParameter_
		methodParameter = uti.ReplaceAll(
			methodParameter,
			"parameterName",
			parameterName,
		)
		methodParameter = uti.ReplaceAll(
			methodParameter,
			"parameterType",
			parameterType,
		)
		methodParameters += methodParameter
	}
	if uti.IsDefined(methodParameters) {
		methodParameters += "\n"
	}
	return methodParameters
}

func (v *moduleSynthesizer_) extractType(
	abstraction mod.AbstractionLike,
	model mod.ModelLike,
) string {
	var abstractType string
	var wrapper = abstraction.GetOptionalWrapper()
	if uti.IsDefined(wrapper) {
		switch actual := wrapper.GetAny().(type) {
		case mod.ArrayLike:
			abstractType = "[]"
		case mod.MapLike:
			abstractType = "map[" + actual.GetName() + "]"
		case mod.ChannelLike:
			abstractType = "chan "
		}
	}
	var analyzer = ana.PackageAnalyzerClass().PackageAnalyzer(model)
	var packageAcronym = analyzer.GetPackageName()[0:3] + "."
	var name = abstraction.GetName()
	var prefix = abstraction.GetOptionalPrefix()
	if uti.IsUndefined(prefix) {
		var typeDeclarations = analyzer.GetTypeDeclarations()
		var types = typeDeclarations.GetIterator()
		for types.HasNext() {
			var typeName = types.GetNext().GetDeclaration().GetName()
			if typeName == name {
				prefix = packageAcronym
				break
			}
		}
		var enumeratedValues = analyzer.GetEnumeratedValues()
		var enumerations = enumeratedValues.GetIterator()
		for enumerations.HasNext() {
			var enumerationName = enumerations.GetNext()
			if enumerationName == name {
				prefix = packageAcronym
				break
			}
		}
		var functionalDeclarations = analyzer.GetFunctionalDeclarations()
		var functionals = functionalDeclarations.GetIterator()
		for functionals.HasNext() {
			var functionalName = functionals.GetNext().GetDeclaration().GetName()
			if functionalName == name {
				prefix = packageAcronym
				break
			}
		}
		var classDeclarations = analyzer.GetClassDeclarations()
		var classes = classDeclarations.GetIterator()
		for classes.HasNext() {
			var className = classes.GetNext().GetDeclaration().GetName()
			if className == name {
				prefix = packageAcronym
				break
			}
		}
		var instanceDeclarations = analyzer.GetInstanceDeclarations()
		var instances = instanceDeclarations.GetIterator()
		for instances.HasNext() {
			var instanceName = instances.GetNext().GetDeclaration().GetName()
			if instanceName == name {
				prefix = packageAcronym
				break
			}
		}
		var aspectDeclarations = analyzer.GetAspectDeclarations()
		var aspects = aspectDeclarations.GetIterator()
		for aspects.HasNext() {
			var aspectName = aspects.GetNext().GetDeclaration().GetName()
			if aspectName == name {
				prefix = packageAcronym
				break
			}
		}
	}
	abstractType += prefix
	abstractType += name
	var arguments = abstraction.GetOptionalArguments()
	if uti.IsDefined(arguments) {
		var argument = v.extractType(
			arguments.GetArgument().GetAbstraction(),
			model,
		)
		abstractType += "[" + argument
		var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
		for additionalArguments.HasNext() {
			var additionalArgument = additionalArguments.GetNext().GetArgument()
			argument = v.extractType(
				additionalArgument.GetAbstraction(),
				model,
			)
			abstractType += ", " + argument
		}
		abstractType += "]"
	}
	return abstractType
}

func (v *moduleSynthesizer_) preserveExistingCode(
	existing string,
	generated string,
) string {
	// Preserve the module description.
	var pattern = `└──────────────────────────────────────────────────────────────────────────────┘(.|\r?\n)+package module`
	generated = v.replacePattern(pattern, existing, generated)

	// Preserve the global functions.
	pattern = `// GLOBAL FUNCTIONS(.|\r?\n)*`
	generated = v.replacePattern(pattern, existing, generated)
	return generated
}

func (v *moduleSynthesizer_) replacePackageAttributes(
	generated string,
	packageName string,
) string {
	generated = uti.ReplaceAll(
		generated,
		"packageName",
		packageName,
	)
	var packageAcronym = packageName[0:3]
	generated = uti.ReplaceAll(
		generated,
		"packageAcronym",
		packageAcronym,
	)
	return generated
}

func (v *moduleSynthesizer_) replacePattern(
	pattern string,
	existing string,
	generated string,
) string {
	var matcher = reg.MustCompile(pattern)
	var existingPattern = matcher.FindString(existing)
	var generatedPattern = matcher.FindString(generated)
	generated = sts.ReplaceAll(
		generated,
		generatedPattern,
		existingPattern,
	)
	return generated
}

func (v *moduleSynthesizer_) updateImportedPackages(
	existing string,
	generated string,
) string {
	// Seed the imported packages with the most common ones.
	var imports = abs.CatalogClass[string, string]().CatalogFromMap(
		map[string]string{
			"fmt": `"fmt"`,
			"col": `"github.com/craterdog/go-collection-framework/v5"`,
			"abs": `"github.com/craterdog/go-collection-framework/v5/collection"`,
			"uti": `"github.com/craterdog/go-missing-utilities/v2"`,
			"ref": `"reflect"`,
			"sts": `"strings"`,
		},
	)

	// Add in the imported packages from the existing code.
	imports = abs.CatalogClass[string, string]().Merge(
		imports, v.extractImportedPackages(existing),
	)

	// Add in the imported packages from each class-model.
	var models = v.models_.GetIterator()
	for models.HasNext() {
		var association = models.GetNext()
		var packageName = association.GetKey()
		var packageAcronym = packageName[0:3]
		var packagePath = `"<moduleName>/` + packageName + `"`
		imports.SetValue(packageAcronym, packagePath)
		var model = association.GetValue()
		var analyzer = ana.PackageAnalyzerClass().PackageAnalyzer(model)
		imports = abs.CatalogClass[string, string]().Merge(
			imports, analyzer.GetImportedPackages(),
		)
	}
	imports.SortValuesWithRanker(
		func(first, second abs.AssociationLike[string, string]) col.Rank {
			var firstValue = first.GetValue()
			var secondValue = second.GetValue()
			switch {
			case firstValue < secondValue:
				return col.LesserRank
			case firstValue > secondValue:
				return col.GreaterRank
			default:
				return col.EqualRank
			}
		},
	)

	// Create imported package statements for each imported package.
	var importedPackages string
	var packages = imports.GetIterator()
	for packages.HasNext() {
		var association = packages.GetNext()
		var packageAcronym = association.GetKey()
		var packagePath = association.GetValue()
		if sts.Contains(generated, packageAcronym+".") {
			// Only import packages that are actually used in the generated code.
			importedPackages += v.createImportedPath(
				packageAcronym,
				packagePath,
			)
		}
	}
	if uti.IsDefined(importedPackages) {
		importedPackages += "\n"
	}

	// Insert the imported packages into the generated code.
	generated = uti.ReplaceAll(
		generated,
		"importedPackages",
		importedPackages,
	)
	return generated
}

// Instance Structure

type moduleSynthesizer_ struct {
	// Declare the instance attributes.
	models_ abs.CatalogLike[string, mod.ModelLike]
}

// Class Structure

type moduleSynthesizerClass_ struct {
	// Declare the class constants.
	warningMessage_      string
	importedPath_        string
	packageAliases_      string
	typeAliases_         string
	enumeratedAliases_   string
	nameAlias_           string
	classConstructors_   string
	constructorFunction_ string
	methodParameter_     string
}

// Class Reference

func moduleSynthesizerClassReference() *moduleSynthesizerClass_ {
	return moduleSynthesizerClassReference_
}

var moduleSynthesizerClassReference_ = &moduleSynthesizerClass_{
	// Initialize the class constants.
	warningMessage_: `
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│              This "Module.go" file was automatically generated.              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘
`,

	importedPath_: `
	<~packageAcronym> <packagePath>`,

	packageAliases_: `
// <~PackageName><TypeAliases>
`,

	typeAliases_: `

type (<NameAliases>)`,

	enumeratedAliases_: `

const (<NameAliases>)`,

	nameAlias_: `
	<Name> = <~packageAcronym>.<Name>`,

	classConstructors_: `
// <~PackageName>/<~ClassName><Constructors>
`,

	constructorFunction_: `

func <MethodName><Constraints>(<Parameters>) <~packageAcronym>.<~ClassName>Like<Arguments> {
	return <~packageAcronym>.<~ClassName>Class<Arguments>().<MethodName>(<ParameterNames>)
}`,

	methodParameter_: `
	<parameterName_> <ParameterType>,`,
}
