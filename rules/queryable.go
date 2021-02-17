// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package rules

type Queryable interface {
	Cast() Queryable
	GetCollectionName() string
}
