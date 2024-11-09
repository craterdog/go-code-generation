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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func Rule() RuleClassLike {
	return ruleReference()
}

// Constructor Methods

func (c *ruleClass_) Make(
	uppercase string,
	definition DefinitionLike,
) RuleLike {
	if uti.IsUndefined(uppercase) {
		panic("The \"uppercase\" attribute is required by this class.")
	}
	if uti.IsUndefined(definition) {
		panic("The \"definition\" attribute is required by this class.")
	}
	var instance = &rule_{
		// Initialize the instance attributes.
		uppercase_:  uppercase,
		definition_: definition,
	}
	return instance

}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Primary Methods

func (v *rule_) GetClass() RuleClassLike {
	return ruleReference()
}

// Attribute Methods

func (v *rule_) GetUppercase() string {
	return v.uppercase_
}

func (v *rule_) GetDefinition() DefinitionLike {
	return v.definition_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type rule_ struct {
	// Declare the instance attributes.
	uppercase_  string
	definition_ DefinitionLike
}

// Class Structure

type ruleClass_ struct {
	// Declare the class constants.
}

// Class Reference

func ruleReference() *ruleClass_ {
	return ruleReference_
}

var ruleReference_ = &ruleClass_{
	// Initialize the class constants.
}
