/*
................................................................................
.                  Copyright (c) 2025.  All Rights Reserved.                   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package example

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v8"
	syn "sync"
)

// CLASS INTERFACE

// Access Function

func ArrayClass[V any]() ArrayClassLike[V] {
	return arrayClass[V]()
}

// Constructor Methods

func (c *arrayClass_[V]) Array(
	array []V,
) ArrayLike[V] {
	return array_[V](array)
}

func (c *arrayClass_[V]) ArrayWithSize(
	size Cardinal,
) ArrayLike[V] {
	if uti.IsUndefined(size) {
		panic("The \"size\" attribute is required by this class.")
	}
	var instance = &array_[V]{
		// Initialize the instance attributes.
	}
	return instance
}

func (c *arrayClass_[V]) ArrayFromSequence(
	values Sequential[V],
) ArrayLike[V] {
	var instance ArrayLike[V]
	// TBD - Add the constructor implementation.
	return instance
}

// Constant Methods

// Function Methods

func (c *arrayClass_[V]) Merge(
	arrays ...ArrayLike[V],
) ArrayLike[V] {
	var result_ ArrayLike[V]
	// TBD - Add the function implementation.
	return result_
}

// INSTANCE INTERFACE

// Principal Methods

func (v array_[V]) GetClass() ArrayClassLike[V] {
	return arrayClass[V]()
}

func (v array_[V]) AsIntrinsic() []V {
	return []V(v)
}

func (v array_[V]) SortValuesWithRanker(
	ranker RankingFunction[V],
) {
	// TBD - Add the method implementation.
}

// Attribute Methods

// Accessible[V] Methods

func (v array_[V]) GetValue(
	index Ordinal,
) V {
	var result_ V
	// TBD - Add the method implementation.
	return result_
}

func (v array_[V]) GetValues(
	first Ordinal,
	last Ordinal,
) Sequential[V] {
	var result_ Sequential[V]
	// TBD - Add the method implementation.
	return result_
}

// Sequential[V] Methods

func (v array_[V]) IsEmpty() bool {
	var result_ bool
	// TBD - Add the method implementation.
	return result_
}

func (v array_[V]) GetSize() Cardinal {
	var result_ Cardinal
	// TBD - Add the method implementation.
	return result_
}

func (v array_[V]) AsArray() []V {
	var result_ []V
	// TBD - Add the method implementation.
	return result_
}

func (v array_[V]) GetIterator() IteratorLike[V] {
	var result_ IteratorLike[V]
	// TBD - Add the method implementation.
	return result_
}

// Updatable[V] Methods

func (v array_[V]) SetValue(
	index Ordinal,
	value V,
) {
	// TBD - Add the method implementation.
}

func (v array_[V]) SetValues(
	index Ordinal,
	values Sequential[V],
) {
	// TBD - Add the method implementation.
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type array_[V any] []V

// Class Structure

type arrayClass_[V any] struct {
	// Declare the class constants.
}

// Class Reference

var arrayMap_ = map[string]any{}
var arrayMutex_ syn.Mutex

func arrayClass[V any]() *arrayClass_[V] {
	// Generate the name of the bound class type.
	var class *arrayClass_[V]
	var name = fmt.Sprintf("%T", class)

	// Check for an existing bound class type.
	arrayMutex_.Lock()
	var value = arrayMap_[name]
	switch actual := value.(type) {
	case *arrayClass_[V]:
		// This bound class type already exists.
		class = actual
	default:
		// Add a new bound class type.
		class = &arrayClass_[V]{
			// Initialize the class constants.
		}
		arrayMap_[name] = class
	}
	arrayMutex_.Unlock()

	// Return a reference to the bound class type.
	return class
}
