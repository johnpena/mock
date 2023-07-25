package core

//go:generate mockgen -package core -self_package github.com/johnpena/mock/mockgen/internal/tests/self_package -destination mock.go github.com/johnpena/mock/mockgen/internal/tests/self_package Methods

type Info struct{}

type Methods interface {
	getInfo() Info
}
