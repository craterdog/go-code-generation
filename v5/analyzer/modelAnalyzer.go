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
	fmt "fmt"
	ast "github.com/craterdog/go-class-model/v5/ast"
	col "github.com/craterdog/go-collection-framework/v4"
	abs "github.com/craterdog/go-collection-framework/v4/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ModelAnalyzer() ModelAnalyzerClassLike {
	return modelAnalyzerReference()
}

// Constructor Methods

func (c *modelAnalyzerClass_) Make(
	model ast.ModelLike,
	className string,
) ModelAnalyzerLike {
	var instance = &modelAnalyzer_{
		// Initialize the instance attributes.
	}
	instance.analyzeClass(model, className)
	return instance

}

// INSTANCE INTERFACE

// Primary Methods

func (v *modelAnalyzer_) GetClass() ModelAnalyzerClassLike {
	return modelAnalyzerReference()
}

func (v *modelAnalyzer_) GetLegalNotice() string {
	return v.legalNotice_
}

func (v *modelAnalyzer_) GetPackageImports() ast.PackageImportsLike {
	return v.packageImports_
}

func (v *modelAnalyzer_) IsGeneric() bool {
	return v.isGeneric_
}

func (v *modelAnalyzer_) GetTypeConstraints() string {
	return v.typeConstraints_
}

func (v *modelAnalyzer_) GetTypeArguments() string {
	return v.typeArguments_
}

func (v *modelAnalyzer_) IsIntrinsic() bool {
	return v.isIntrinsic_
}

func (v *modelAnalyzer_) GetIntrinsicType() ast.AbstractionLike {
	return v.intrinsicType_
}

func (v *modelAnalyzer_) GetConstants() abs.CatalogLike[string, string] {
	return v.constants_
}

func (v *modelAnalyzer_) GetAttributes() abs.CatalogLike[string, string] {
	return v.attributes_
}

func (v *modelAnalyzer_) GetConstructorMethods() abs.ListLike[ast.ConstructorMethodLike] {
	return v.constructorMethods_
}

func (v *modelAnalyzer_) GetConstantMethods() abs.ListLike[ast.ConstantMethodLike] {
	return v.constantMethods_
}

func (v *modelAnalyzer_) GetFunctionMethods() abs.ListLike[ast.FunctionMethodLike] {
	return v.functionMethods_
}

func (v *modelAnalyzer_) GetPrimaryMethods() abs.ListLike[ast.PrimaryMethodLike] {
	return v.primaryMethods_
}

func (v *modelAnalyzer_) GetAttributeMethods() abs.ListLike[ast.AttributeMethodLike] {
	return v.attributeMethods_
}

func (v *modelAnalyzer_) GetAspectInterfaces() abs.ListLike[ast.AspectInterfaceLike] {
	return v.aspectInterfaces_
}

func (v *modelAnalyzer_) GetAspectDeclarations() abs.ListLike[ast.AspectDeclarationLike] {
	return v.aspectDeclarations_
}

// PROTECTED INTERFACE

// Private Methods

func (v *modelAnalyzer_) analyzeAspectDeclarations(
	interfaceDeclarations ast.InterfaceDeclarationsLike,
) {
	var aspectSection = interfaceDeclarations.GetAspectSection()
	var aspectDeclarations = aspectSection.GetAspectDeclarations()
	v.aspectDeclarations_ = col.List[ast.AspectDeclarationLike](aspectDeclarations)
}

func (v *modelAnalyzer_) analyzeAspectInterfaces(
	instanceDeclaration ast.InstanceDeclarationLike,
) {
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	var aspectSubsection = instanceMethods.GetOptionalAspectSubsection()
	if uti.IsDefined(aspectSubsection) {
		var aspectInterfaces = aspectSubsection.GetAspectInterfaces()
		v.aspectInterfaces_ = col.List[ast.AspectInterfaceLike](
			aspectInterfaces,
		)
	}
}

func (v *modelAnalyzer_) analyzeAttributeMethods(
	instanceDeclaration ast.InstanceDeclarationLike,
) {
	v.attributeMethods_ = col.List[ast.AttributeMethodLike]()
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

func (v *modelAnalyzer_) analyzeClass(
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
			v.analyzePrimaryMethods(instanceDeclaration)
			v.analyzeAttributeMethods(instanceDeclaration)
			v.analyzeAspectInterfaces(instanceDeclaration)
			break
		}
	}
}

func (v *modelAnalyzer_) analyzeClassConstants(
	classDeclaration ast.ClassDeclarationLike,
) {
	v.constants_ = col.Catalog[string, string]()
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

func (v *modelAnalyzer_) analyzeClassGenerics(
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

func (v *modelAnalyzer_) analyzeClassStructure(
	instanceDeclaration ast.InstanceDeclarationLike,
) {
	v.isIntrinsic_ = false
	v.intrinsicType_ = nil
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	var primarySubsection = instanceMethods.GetPrimarySubsection()
	var primaryMethods = primarySubsection.GetPrimaryMethods().GetIterator()
	for primaryMethods.HasNext() {
		var method = primaryMethods.GetNext().GetMethod()
		var methodName = method.GetName()
		if methodName == "GetIntrinsic" {
			v.isIntrinsic_ = true
			var result = method.GetOptionalResult()
			v.intrinsicType_ = result.GetAny().(ast.AbstractionLike)
		}
	}
}

func (v *modelAnalyzer_) analyzeConstantMethods(
	classDeclaration ast.ClassDeclarationLike,
) {
	v.constantMethods_ = col.List[ast.ConstantMethodLike]()
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

func (v *modelAnalyzer_) analyzeConstructorMethods(
	classDeclaration ast.ClassDeclarationLike,
) {
	v.constructorMethods_ = col.List[ast.ConstructorMethodLike]()
	var classMethods = classDeclaration.GetClassMethods()
	var constructorSubsection = classMethods.GetConstructorSubsection()
	var constructorMethods = constructorSubsection.GetConstructorMethods().GetIterator()
	for constructorMethods.HasNext() {
		var constructorMethod = constructorMethods.GetNext()
		v.constructorMethods_.AppendValue(constructorMethod)
	}
}

func (v *modelAnalyzer_) analyzeFunctionMethods(
	classDeclaration ast.ClassDeclarationLike,
) {
	v.functionMethods_ = col.List[ast.FunctionMethodLike]()
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

func (v *modelAnalyzer_) analyzePackageDeclaration(
	packageDeclaration ast.PackageDeclarationLike,
) {
	v.legalNotice_ = packageDeclaration.GetLegalNotice().GetComment()
}

func (v *modelAnalyzer_) analyzePrimaryMethods(
	instanceDeclaration ast.InstanceDeclarationLike,
) {
	v.primaryMethods_ = col.List[ast.PrimaryMethodLike]()
	var instanceMethods = instanceDeclaration.GetInstanceMethods()
	var primarySubsection = instanceMethods.GetPrimarySubsection()
	var primaryMethods = primarySubsection.GetPrimaryMethods().GetIterator()
	for primaryMethods.HasNext() {
		var primaryMethod = primaryMethods.GetNext()
		v.primaryMethods_.AppendValue(primaryMethod)
	}
}

func (v *modelAnalyzer_) analyzePrivateAttributes(
	classDeclaration ast.ClassDeclarationLike,
) {
	var classMethods = classDeclaration.GetClassMethods()
	var constructorSubsection = classMethods.GetConstructorSubsection()
	var constructorMethods = constructorSubsection.GetConstructorMethods().GetIterator()
	for constructorMethods.HasNext() {
		var constructorMethod = constructorMethods.GetNext()
		var name = constructorMethod.GetName()
		// Focus only on constructors that are passed attributes as arguments.
		if !v.isIntrinsic_ &&
			(name == "Make" || sts.HasPrefix(name, "MakeWith")) {
			var parameters = constructorMethod.GetParameters().GetIterator()
			for parameters.HasNext() {
				var parameter = parameters.GetNext()
				var attributeName = sts.TrimSuffix(parameter.GetName(), "_")
				var attributeType = v.extractType(parameter.GetAbstraction())
				v.attributes_.SetValue(attributeName, attributeType)
			}
		}
	}
}

func (v *modelAnalyzer_) analyzePublicAttributes(
	instanceDeclaration ast.InstanceDeclarationLike,
) {
	v.attributes_ = col.Catalog[string, string]()
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

func (v *modelAnalyzer_) analyzeTypeArguments(
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

func (v *modelAnalyzer_) analyzeTypeConstraints(
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

func (v *modelAnalyzer_) extractAttributeName(
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

func (v *modelAnalyzer_) extractClassName(
	classDeclaration ast.ClassDeclarationLike,
) string {
	var className = classDeclaration.GetDeclaration().GetName()
	className = sts.TrimSuffix(className, "ClassLike")
	className = uti.MakeLowerCase(className)
	return className
}

func (v *modelAnalyzer_) extractType(
	abstraction ast.AbstractionLike,
) string {
	var abstractType string
	var prefix = abstraction.GetOptionalPrefix()
	if uti.IsDefined(prefix) {
		switch actual := prefix.GetAny().(type) {
		case ast.ArrayLike:
			abstractType = "[]"
		case ast.MapLike:
			abstractType = "map[" + actual.GetName() + "]"
		case ast.ChannelLike:
			abstractType = "chan "
		}
	}
	var name = abstraction.GetName()
	abstractType += name
	var suffix = abstraction.GetOptionalSuffix()
	if uti.IsDefined(suffix) {
		abstractType += "." + suffix.GetName()
	}
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

type modelAnalyzer_ struct {
	// Declare the instance attributes.
	legalNotice_        string
	isGeneric_          bool
	typeConstraints_    string
	typeArguments_      string
	isIntrinsic_        bool
	intrinsicType_      ast.AbstractionLike
	constants_          abs.CatalogLike[string, string]
	attributes_         abs.CatalogLike[string, string]
	constructorMethods_ abs.ListLike[ast.ConstructorMethodLike]
	constantMethods_    abs.ListLike[ast.ConstantMethodLike]
	functionMethods_    abs.ListLike[ast.FunctionMethodLike]
	primaryMethods_     abs.ListLike[ast.PrimaryMethodLike]
	attributeMethods_   abs.ListLike[ast.AttributeMethodLike]
	aspectInterfaces_   abs.ListLike[ast.AspectInterfaceLike]
	aspectDeclarations_ abs.ListLike[ast.AspectDeclarationLike]
	packageImports_     ast.PackageImportsLike
}

// Class Structure

type modelAnalyzerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func modelAnalyzerReference() *modelAnalyzerClass_ {
	return modelAnalyzerReference_
}

var modelAnalyzerReference_ = &modelAnalyzerClass_{
	// Initialize the class constants.
}
