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
	fmt "fmt"
	ast "github.com/craterdog/go-class-model/v7/ast"
	com "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ClassAnalyzerClass() ClassAnalyzerClassLike {
	return classAnalyzerClass()
}

// Constructor Methods

func (c *classAnalyzerClass_) ClassAnalyzer(
	model ast.ModelLike,
	className string,
) ClassAnalyzerLike {
	var instance = &classAnalyzer_{
		// Initialize the instance attributes.
	}
	instance.analyzeClass(model, className)
	return instance

}

// INSTANCE INTERFACE

// Principal Methods

func (v *classAnalyzer_) GetClass() ClassAnalyzerClassLike {
	return classAnalyzerClass()
}

func (v *classAnalyzer_) GetLegalNotice() string {
	return v.legalNotice_
}

func (v *classAnalyzer_) GetImportedPackages() com.CatalogLike[string, string] {
	return v.importedPackages_
}

func (v *classAnalyzer_) IsGeneric() bool {
	return v.isGeneric_
}

func (v *classAnalyzer_) GetTypeConstraints() string {
	return v.typeConstraints_
}

func (v *classAnalyzer_) GetTypeArguments() string {
	return v.typeArguments_
}

func (v *classAnalyzer_) IsIntrinsic() bool {
	return v.isIntrinsic_
}

func (v *classAnalyzer_) GetIntrinsicType() ast.AbstractionLike {
	return v.intrinsicType_
}

func (v *classAnalyzer_) GetConstants() com.CatalogLike[string, string] {
	return v.constants_
}

func (v *classAnalyzer_) GetAttributes() com.CatalogLike[string, string] {
	return v.attributes_
}

func (v *classAnalyzer_) GetConstructorMethods() com.ListLike[ast.ConstructorMethodLike] {
	return v.constructorMethods_
}

func (v *classAnalyzer_) GetConstantMethods() com.ListLike[ast.ConstantMethodLike] {
	return v.constantMethods_
}

func (v *classAnalyzer_) GetFunctionMethods() com.ListLike[ast.FunctionMethodLike] {
	return v.functionMethods_
}

func (v *classAnalyzer_) GetPrincipalMethods() com.ListLike[ast.PrincipalMethodLike] {
	return v.principalMethods_
}

func (v *classAnalyzer_) GetAttributeMethods() com.ListLike[ast.AttributeMethodLike] {
	return v.attributeMethods_
}

func (v *classAnalyzer_) GetAspectInterfaces() com.ListLike[ast.AspectInterfaceLike] {
	return v.aspectInterfaces_
}

func (v *classAnalyzer_) GetAspectDeclarations() com.ListLike[ast.AspectDeclarationLike] {
	return v.aspectDeclarations_
}

// PROTECTED INTERFACE

// Private Methods

func (v *classAnalyzer_) analyzeAspectDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
) {
	var aspectSection = interfaceDeclarations.GetAspectSection()
	v.aspectDeclarations_ = aspectSection.GetAspectDeclarations()
}

func (v *classAnalyzer_) analyzeAspectInterfaces(
	instanceDeclaration ast.InstanceDeclarationLike,
) {
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	var aspectSubsection = instanceMethods.GetOptionalAspectSubsection()
	if uti.IsDefined(aspectSubsection) {
		v.aspectInterfaces_ = aspectSubsection.GetAspectInterfaces()
	}
}

func (v *classAnalyzer_) analyzeAttributeMethods(
	instanceDeclaration ast.InstanceDeclarationLike,
) {
	v.attributeMethods_ = com.List[ast.AttributeMethodLike]()
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	var attributeSubsection = instanceMethods.GetOptionalAttributeSubsection()
	if uti.IsDefined(attributeSubsection) {
		var attributeMethods = attributeSubsection.GetAttributeMethods().GetIterator()
		for attributeMethods.HasNext() {
			var attributeMethod = attributeMethods.GetNext()
			v.attributeMethods_.AppendValue(attributeMethod)
		}
	}
}

func (v *classAnalyzer_) analyzeClass(
	model ast.ModelLike,
	className string,
) {
	var packageDeclaration = model.GetPackageDeclaration()
	v.analyzePackageDeclaration(packageDeclaration)
	var interfaceDeclarations = model.GetInterfaceDeclarations()
	v.analyzeAspectDeclarations(interfaceDeclarations)
	var classSection = interfaceDeclarations.GetClassSection()
	var classDeclarations = classSection.GetClassDeclarations().GetIterator()
	var instanceSection = interfaceDeclarations.GetInstanceSection()
	var instanceDeclarations = instanceSection.GetInstanceDeclarations().GetIterator()
	for classDeclarations.HasNext() && instanceDeclarations.HasNext() {
		var classDeclaration = classDeclarations.GetNext()
		var instanceDeclaration = instanceDeclarations.GetNext()
		if className == v.extractClassName(classDeclaration) {
			v.analyzeClassGenerics(classDeclaration)
			v.analyzeClassStructure(instanceDeclaration)
			v.analyzeClassConstants(classDeclaration)
			v.analyzePublicAttributes(instanceDeclaration)
			v.analyzePrivateAttributes(classDeclaration)
			v.analyzeConstructorMethods(classDeclaration)
			v.analyzeConstantMethods(classDeclaration)
			v.analyzeFunctionMethods(classDeclaration)
			v.analyzePrincipalMethods(instanceDeclaration)
			v.analyzeAttributeMethods(instanceDeclaration)
			v.analyzeAspectInterfaces(instanceDeclaration)
			break
		}
	}
}

func (v *classAnalyzer_) analyzeClassConstants(
	classDeclaration ast.ClassDeclarationLike,
) {
	v.constants_ = com.Catalog[string, string]()
	var classMethods = classDeclaration.GetClassMethods()
	var constantSubsection = classMethods.GetOptionalConstantSubsection()
	if uti.IsDefined(constantSubsection) {
		var constantMethods = constantSubsection.GetConstantMethods().GetIterator()
		for constantMethods.HasNext() {
			var constantMethod = constantMethods.GetNext()
			var constantName = constantMethod.GetName()
			var constantType = v.extractType(constantMethod.GetAbstraction())
			v.constants_.SetValue(constantName, constantType)
		}
	}
}

func (v *classAnalyzer_) analyzeClassGenerics(
	classDeclaration ast.ClassDeclarationLike,
) {
	v.isGeneric_ = false
	var declaration = classDeclaration.GetDeclaration()
	var constraints = declaration.GetOptionalConstraints()
	if uti.IsDefined(constraints) {
		v.isGeneric_ = true
		v.analyzeTypeConstraints(classDeclaration)
		v.analyzeTypeArguments(classDeclaration)
	}
}

func (v *classAnalyzer_) analyzeClassStructure(
	instanceDeclaration ast.InstanceDeclarationLike,
) {
	v.isIntrinsic_ = false
	v.intrinsicType_ = nil
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	var principalSubsection = instanceMethods.GetPrincipalSubsection()
	var principalMethods = principalSubsection.GetPrincipalMethods().GetIterator()
	for principalMethods.HasNext() {
		var method = principalMethods.GetNext().GetMethod()
		var methodName = method.GetName()
		if methodName == "AsIntrinsic" {
			v.isIntrinsic_ = true
			var result = method.GetResult()
			v.intrinsicType_ = result.GetAny().(ast.AbstractionLike)
		}
	}
}

func (v *classAnalyzer_) analyzeConstantMethods(
	classDeclaration ast.ClassDeclarationLike,
) {
	v.constantMethods_ = com.List[ast.ConstantMethodLike]()
	var classMethods = classDeclaration.GetClassMethods()
	var constantSubsection = classMethods.GetOptionalConstantSubsection()
	if uti.IsDefined(constantSubsection) {
		var constantMethods = constantSubsection.GetConstantMethods().GetIterator()
		for constantMethods.HasNext() {
			var constantMethod = constantMethods.GetNext()
			v.constantMethods_.AppendValue(constantMethod)
		}
	}
}

func (v *classAnalyzer_) analyzeConstructorMethods(
	classDeclaration ast.ClassDeclarationLike,
) {
	v.constructorMethods_ = com.List[ast.ConstructorMethodLike]()
	var classMethods = classDeclaration.GetClassMethods()
	var constructorSubsection = classMethods.GetConstructorSubsection()
	var constructorMethods = constructorSubsection.GetConstructorMethods().GetIterator()
	for constructorMethods.HasNext() {
		var constructorMethod = constructorMethods.GetNext()
		v.constructorMethods_.AppendValue(constructorMethod)
	}
}

func (v *classAnalyzer_) analyzeFunctionMethods(
	classDeclaration ast.ClassDeclarationLike,
) {
	v.functionMethods_ = com.List[ast.FunctionMethodLike]()
	var classMethods = classDeclaration.GetClassMethods()
	var functionSubsection = classMethods.GetOptionalFunctionSubsection()
	if uti.IsDefined(functionSubsection) {
		var functionMethods = functionSubsection.GetFunctionMethods().GetIterator()
		for functionMethods.HasNext() {
			var functionMethod = functionMethods.GetNext()
			v.functionMethods_.AppendValue(functionMethod)
		}
	}
}

func (v *classAnalyzer_) analyzePackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
) {
	v.legalNotice_ = packageDeclaration.GetLegalNotice().GetComment()
	v.importedPackages_ = com.Catalog[string, string]()
	var packageImports = packageDeclaration.GetPackageImports()
	var importList = packageImports.GetOptionalImportList()
	if uti.IsUndefined(importList) {
		// This package has no imports.
		return
	}
	var packages = importList.GetImportedPackages().GetIterator()
	for packages.HasNext() {
		var importedPackage = packages.GetNext()
		var packagePath = importedPackage.GetPath()
		var packageAcronym = importedPackage.GetName()
		v.importedPackages_.SetValue(packageAcronym, packagePath)
	}
	v.importedPackages_.SetValue(
		"fmt",
		`"fmt"`,
	)
	v.importedPackages_.SetValue(
		"uti",
		`"github.com/craterdog/go-missing-utilities/v7"`,
	)
	v.importedPackages_.SetValue(
		"com",
		`"github.com/craterdog/go-component-framework/v7"`,
	)
	v.importedPackages_.SetValue(
		"syn",
		`"sync"`,
	)
}

func (v *classAnalyzer_) analyzePrincipalMethods(
	instanceDeclaration ast.InstanceDeclarationLike,
) {
	v.principalMethods_ = com.List[ast.PrincipalMethodLike]()
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	var principalSubsection = instanceMethods.GetPrincipalSubsection()
	var principalMethods = principalSubsection.GetPrincipalMethods().GetIterator()
	for principalMethods.HasNext() {
		var principalMethod = principalMethods.GetNext()
		v.principalMethods_.AppendValue(principalMethod)
	}
}

func (v *classAnalyzer_) analyzePrivateAttributes(
	classDeclaration ast.ClassDeclarationLike,
) {
	var classMethods = classDeclaration.GetClassMethods()
	var className = v.extractClassName(classDeclaration)
	var constructorSubsection = classMethods.GetConstructorSubsection()
	var constructorMethods = constructorSubsection.GetConstructorMethods().GetIterator()
	for constructorMethods.HasNext() {
		var constructorMethod = constructorMethods.GetNext()
		var name = constructorMethod.GetName()
		// Focus only on constructors that are passed attributes as arguments.
		if !v.isIntrinsic_ &&
			(name == className || sts.HasPrefix(name, className+"With")) {
			var parameterList = constructorMethod.GetOptionalParameterList()
			if uti.IsUndefined(parameterList) {
				// This constructor method takes no parameters.
				continue
			}
			var parameters = parameterList.GetParameters().GetIterator()
			for parameters.HasNext() {
				var parameter = parameters.GetNext()
				var attributeName = sts.TrimSuffix(parameter.GetName(), "_")
				var attributeType = v.extractType(parameter.GetAbstraction())
				v.attributes_.SetValue(attributeName, attributeType)
			}
		}
	}
}

func (v *classAnalyzer_) analyzePublicAttributes(
	instanceDeclaration ast.InstanceDeclarationLike,
) {
	v.attributes_ = com.Catalog[string, string]()
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	var attributeSubsection = instanceMethods.GetOptionalAttributeSubsection()
	if uti.IsDefined(attributeSubsection) {
		var attributeMethods = attributeSubsection.GetAttributeMethods().GetIterator()
		for attributeMethods.HasNext() {
			var attributeMethod = attributeMethods.GetNext()
			var attributeName string
			var abstraction ast.AbstractionLike
			switch actual := attributeMethod.GetAny().(type) {
			case ast.GetterMethodLike:
				attributeName = v.extractAttributeName(actual.GetName())
				abstraction = actual.GetAbstraction()
			case ast.SetterMethodLike:
				attributeName = v.extractAttributeName(actual.GetName())
				abstraction = actual.GetParameter().GetAbstraction()
			}
			var attributeType = v.extractType(abstraction)
			v.attributes_.SetValue(attributeName, attributeType)
		}
	}
}

func (v *classAnalyzer_) analyzeTypeArguments(
	classDeclaration ast.ClassDeclarationLike,
) {
	if v.isGeneric_ {
		v.typeArguments_ = "["
		var classDeclaration = classDeclaration.GetDeclaration()
		var optionalConstraints = classDeclaration.GetOptionalConstraints()
		var constraint = optionalConstraints.GetConstraint()
		var argument = constraint.GetName()
		v.typeArguments_ += argument
		var additionalConstraints = optionalConstraints.GetAdditionalConstraints().GetIterator()
		for additionalConstraints.HasNext() {
			constraint = additionalConstraints.GetNext().GetConstraint()
			argument = constraint.GetName()
			v.typeArguments_ += ", " + argument
		}
		v.typeArguments_ += "]"
	}
}

func (v *classAnalyzer_) analyzeTypeConstraints(
	classDeclaration ast.ClassDeclarationLike,
) {
	if v.isGeneric_ {
		v.typeConstraints_ = "["
		var classDeclaration = classDeclaration.GetDeclaration()
		var optionalConstraints = classDeclaration.GetOptionalConstraints()
		var constraint = optionalConstraints.GetConstraint()
		var constraintName = constraint.GetName()
		var constraintType = v.extractType(constraint.GetAbstraction())
		v.typeConstraints_ += constraintName + " " + constraintType
		var additionalConstraints = optionalConstraints.GetAdditionalConstraints().GetIterator()
		for additionalConstraints.HasNext() {
			constraint = additionalConstraints.GetNext().GetConstraint()
			constraintName = constraint.GetName()
			constraintType = v.extractType(constraint.GetAbstraction())
			v.typeConstraints_ += ", " + constraintName + " " + constraintType
		}
		v.typeConstraints_ += "]"
	}
}

func (v *classAnalyzer_) extractAttributeName(
	accessorName string,
) string {
	var attributeName string
	switch {
	case sts.HasPrefix(accessorName, "Get"):
		attributeName = sts.TrimPrefix(accessorName, "Get")
	case sts.HasPrefix(accessorName, "Set"):
		attributeName = sts.TrimPrefix(accessorName, "Set")
	case sts.HasPrefix(accessorName, "Is"):
		attributeName = sts.TrimPrefix(accessorName, "Is")
	case sts.HasPrefix(accessorName, "Was"):
		attributeName = sts.TrimPrefix(accessorName, "Was")
	case sts.HasPrefix(accessorName, "Are"):
		attributeName = sts.TrimPrefix(accessorName, "Are")
	case sts.HasPrefix(accessorName, "Were"):
		attributeName = sts.TrimPrefix(accessorName, "Were")
	case sts.HasPrefix(accessorName, "Has"):
		attributeName = sts.TrimPrefix(accessorName, "Has")
	case sts.HasPrefix(accessorName, "Had"):
		attributeName = sts.TrimPrefix(accessorName, "Had")
	case sts.HasPrefix(accessorName, "Have"):
		attributeName = sts.TrimPrefix(accessorName, "Have")
	default:
		var message = fmt.Sprintf(
			"An unknown accessor name was found: %q",
			accessorName,
		)
		panic(message)
	}
	attributeName = uti.MakeLowerCase(attributeName)
	return attributeName
}

func (v *classAnalyzer_) extractClassName(
	classDeclaration ast.ClassDeclarationLike,
) string {
	var className = classDeclaration.GetDeclaration().GetName()
	className = sts.TrimSuffix(className, "ClassLike")
	className = uti.MakeLowerCase(className)
	return className
}

func (v *classAnalyzer_) extractType(
	abstraction ast.AbstractionLike,
) string {
	var abstractType string
	var wrapper = abstraction.GetOptionalWrapper()
	if uti.IsDefined(wrapper) {
		switch actual := wrapper.GetAny().(type) {
		case ast.StarLike:
			abstractType = "*"
		case ast.ArrayLike:
			abstractType = "[]"
		case ast.MapLike:
			abstractType = "map[" + actual.GetName() + "]"
		case ast.ChannelLike:
			abstractType = "chan "
		}
	}
	var prefix = abstraction.GetOptionalPrefix()
	if uti.IsDefined(prefix) {
		abstractType += prefix
	}
	var name = abstraction.GetName()
	abstractType += name
	var arguments = abstraction.GetOptionalArguments()
	if uti.IsDefined(arguments) {
		var argument = v.extractType(arguments.GetArgument().GetAbstraction())
		abstractType += "[" + argument
		var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
		for additionalArguments.HasNext() {
			var additionalArgument = additionalArguments.GetNext().GetArgument()
			argument = v.extractType(additionalArgument.GetAbstraction())
			abstractType += ", " + argument
		}
		abstractType += "]"
	}
	return abstractType
}

// Instance Structure

type classAnalyzer_ struct {
	// Declare the instance attributes.
	legalNotice_        string
	isGeneric_          bool
	importedPackages_   com.CatalogLike[string, string]
	typeConstraints_    string
	typeArguments_      string
	isIntrinsic_        bool
	intrinsicType_      ast.AbstractionLike
	constants_          com.CatalogLike[string, string]
	attributes_         com.CatalogLike[string, string]
	constructorMethods_ com.ListLike[ast.ConstructorMethodLike]
	constantMethods_    com.ListLike[ast.ConstantMethodLike]
	functionMethods_    com.ListLike[ast.FunctionMethodLike]
	principalMethods_   com.ListLike[ast.PrincipalMethodLike]
	attributeMethods_   com.ListLike[ast.AttributeMethodLike]
	aspectInterfaces_   com.ListLike[ast.AspectInterfaceLike]
	aspectDeclarations_ com.ListLike[ast.AspectDeclarationLike]
}

// Class Structure

type classAnalyzerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func classAnalyzerClass() *classAnalyzerClass_ {
	return classAnalyzerClassReference_
}

var classAnalyzerClassReference_ = &classAnalyzerClass_{
	// Initialize the class constants.
}
