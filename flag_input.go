package main

type Flag[T comparable] struct {
	limits     []T
	defaultVal T
}

func (f *Flag[T]) Select(flag T) T {
	for _, i := range f.limits {
		if i == flag {
			return flag
		}
	}
	return f.defaultVal
}

func NewFlag[T comparable](defaultVal T, limits ...T) *Flag[T] {
	return &Flag[T]{
		limits:     limits,
		defaultVal: defaultVal,
	}
}
