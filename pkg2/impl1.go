package pkg2

import "github.com/korjavin/goimplements/pkg1"

type B struct {
}

func (b B) Add() {
}

var _ pkg1.A = B{}
