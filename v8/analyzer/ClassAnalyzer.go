/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	mod "github.com/craterdog/go-class-model/v8"
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ClassAnalyzerClass() ClassAnalyzerClassLike {
	return classAnalyzerClass()
}

// Constructor Methods

func (c *classAnalyzerClass_) ClassAnalyzer(
	model mod.ModelLike,
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

func (v *classAnalyzer_) GetIntrinsicType() mod.AbstractionLike {
	return v.intrinsicType_
}

func (v *classAnalyzer_) GetConstants() com.CatalogLike[string, string] {
	return v.constants_
}

func (v *classAnalyzer_) GetAttributes() com.CatalogLike[string, string] {
	return v.attributes_
}

func (v *classAnalyzer_) GetConstructorMethods() com.Sequential[mod.ConstructorMethodLike] {
	return v.constructorMethods_
}

func (v *classAnalyzer_) GetConstantMethods() com.Sequential[mod.ConstantMethodLike] {
	return v.constantMethods_
}

func (v *classAnalyzer_) GetFunctionMethods() com.Sequential[mod.FunctionMethodLike] {
	return v.functionMethods_
}

func (v *classAnalyzer_) GetPrincipalMethods() com.Sequential[mod.PrincipalMethodLike] {
	return v.principalMethods_
}

func (v *classAnalyzer_) GetAttributeMethods() com.Sequential[mod.AttributeMethodLike] {
	return v.attributeMethods_
}

func (v *classAnalyzer_) GetAspectInterfaces() com.Sequential[mod.AspectInterfaceLike] {
	return v.aspectInterfaces_
}

func (v *classAnalyzer_) GetAspectDeclarations() com.Sequential[mod.AspectDeclarationLike] {
	return v.aspectDeclarations_
}

// PROTECTED INTERFACE

// Private Methods

func (v *classAnalyzer_) analyzeAspectDeclarations(
	interfaceDeclarations mod.InterfaceDeclarationsLike,
) {
	var aspectSection = interfaceDeclarations.GetAspectSection()
	v.aspectDeclarations_ = aspectSection.GetAspectDeclarations()
}

func (v *classAnalyzer_) analyzeAspectInterfaces(
	instanceDeclaration mod.InstanceDeclarationLike,
) {
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	var aspectSubsection = instanceMethods.GetOptionalAspectSubsection()
	if uti.IsDefined(aspectSubsection) {
		v.aspectInterfaces_ = aspectSubsection.GetAspectInterfaces()
	}
}

func (v *classAnalyzer_) analyzeAttributeMethods(
	instanceDeclaration mod.InstanceDeclarationLike,
) {
	v.attributeMethods_ = com.List[mod.AttributeMethodLike]()
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
	model mod.ModelLike,
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
	classDeclaration mod.ClassDeclarationLike,
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
	classDeclaration mod.ClassDeclarationLike,
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
	instanceDeclaration mod.InstanceDeclarationLike,
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
			v.intrinsicType_ = result.GetAny().(mod.AbstractionLike)
		}
	}
}

func (v *classAnalyzer_) analyzeConstantMethods(
	classDeclaration mod.ClassDeclarationLike,
) {
	v.constantMethods_ = com.List[mod.ConstantMethodLike]()
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
	classDeclaration mod.ClassDeclarationLike,
) {
	v.constructorMethods_ = com.List[mod.ConstructorMethodLike]()
	var classMethods = classDeclaration.GetClassMethods()
	var constructorSubsection = classMethods.GetConstructorSubsection()
	var constructorMethods = constructorSubsection.GetConstructorMethods().GetIterator()
	for constructorMethods.HasNext() {
		var constructorMethod = constructorMethods.GetNext()
		v.constructorMethods_.AppendValue(constructorMethod)
	}
}

func (v *classAnalyzer_) analyzeFunctionMethods(
	classDeclaration mod.ClassDeclarationLike,
) {
	v.functionMethods_ = com.List[mod.FunctionMethodLike]()
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
	packageDeclaration mod.PackageDeclarationLike,
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

func (v *classAnalyzer_) analyzePrincipalMethods(
	instanceDeclaration mod.InstanceDeclarationLike,
) {
	v.principalMethods_ = com.List[mod.PrincipalMethodLike]()
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	var principalSubsection = instanceMethods.GetPrincipalSubsection()
	var principalMethods = principalSubsection.GetPrincipalMethods().GetIterator()
	for principalMethods.HasNext() {
		var principalMethod = principalMethods.GetNext()
		v.principalMethods_.AppendValue(principalMethod)
	}
}

func (v *classAnalyzer_) analyzePrivateAttributes(
	classDeclaration mod.ClassDeclarationLike,
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
	instanceDeclaration mod.InstanceDeclarationLike,
) {
	v.attributes_ = com.Catalog[string, string]()
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	var attributeSubsection = instanceMethods.GetOptionalAttributeSubsection()
	if uti.IsDefined(attributeSubsection) {
		var attributeMethods = attributeSubsection.GetAttributeMethods().GetIterator()
		for attributeMethods.HasNext() {
			var attributeMethod = attributeMethods.GetNext()
			var attributeName string
			var abstraction mod.AbstractionLike
			switch actual := attributeMethod.GetAny().(type) {
			case mod.GetterMethodLike:
				attributeName = v.extractAttributeName(actual.GetName())
				abstraction = actual.GetAbstraction()
			case mod.SetterMethodLike:
				attributeName = v.extractAttributeName(actual.GetName())
				abstraction = actual.GetParameter().GetAbstraction()
			}
			var attributeType = v.extractType(abstraction)
			v.attributes_.SetValue(attributeName, attributeType)
		}
	}
}

func (v *classAnalyzer_) analyzeTypeArguments(
	classDeclaration mod.ClassDeclarationLike,
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
	classDeclaration mod.ClassDeclarationLike,
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
	classDeclaration mod.ClassDeclarationLike,
) string {
	var className = classDeclaration.GetDeclaration().GetName()
	className = sts.TrimSuffix(className, "ClassLike")
	className = uti.MakeLowerCase(className)
	return className
}

func (v *classAnalyzer_) extractType(
	abstraction mod.AbstractionLike,
) string {
	var abstractType string
	var wrapper = abstraction.GetOptionalWrapper()
	if uti.IsDefined(wrapper) {
		switch actual := wrapper.GetAny().(type) {
		case mod.DotsLike:
			abstractType = "..."
		case mod.StarLike:
			abstractType = "*"
		case mod.ArrayLike:
			abstractType = "[]"
		case mod.ChannelLike:
			abstractType = "chan "
		case mod.MapLike:
			abstractType = "map[" + actual.GetName() + "]"
		}
	}
	switch actual := abstraction.GetType().GetAny().(type) {
	case mod.FunctionalLike:
		abstractType += "func("
		var parameterList = actual.GetOptionalParameterList()
		if uti.IsDefined(parameterList) {
			var parameters = parameterList.GetParameters().GetIterator()
			for parameters.HasNext() {
				var parameter = parameters.GetNext()
				abstractType += parameter.GetName() + " "
				abstractType += v.extractType(parameter.GetAbstraction())
			}
		}
		abstractType += ")"
		var result = actual.GetOptionalResult()
		if uti.IsDefined(result) {
			switch actual := result.GetAny().(type) {
			case mod.NoneLike:
			case mod.AbstractionLike:
				abstractType += v.extractType(actual)
			case mod.MultivalueLike:
				abstractType += "("
				var parameterList = actual.GetParameterList()
				var parameters = parameterList.GetParameters().GetIterator()
				for parameters.HasNext() {
					var parameter = parameters.GetNext()
					abstractType += parameter.GetName() + " "
					abstractType += v.extractType(parameter.GetAbstraction())
				}
				abstractType += ")"
			}
		}
	case mod.NamedLike:
		var prefix = actual.GetOptionalPrefix()
		if uti.IsDefined(prefix) {
			abstractType += prefix
		}
		var name = actual.GetName()
		abstractType += name
		var arguments = actual.GetOptionalArguments()
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
	intrinsicType_      mod.AbstractionLike
	constants_          com.CatalogLike[string, string]
	attributes_         com.CatalogLike[string, string]
	constructorMethods_ com.ListLike[mod.ConstructorMethodLike]
	constantMethods_    com.ListLike[mod.ConstantMethodLike]
	functionMethods_    com.ListLike[mod.FunctionMethodLike]
	principalMethods_   com.ListLike[mod.PrincipalMethodLike]
	attributeMethods_   com.ListLike[mod.AttributeMethodLike]
	aspectInterfaces_   com.Sequential[mod.AspectInterfaceLike]
	aspectDeclarations_ com.Sequential[mod.AspectDeclarationLike]
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
