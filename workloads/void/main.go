package main

import (
	"fmt"

	"cs.utexas.edu/zjia/faas-void/handlers"

	"cs.utexas.edu/zjia/faas"
	"cs.utexas.edu/zjia/faas/types"
)

type funcHandlerFactory struct {
}

func (f *funcHandlerFactory) New(env types.Environment, funcName string) (types.FuncHandler, error) {
	switch funcName {
	case "VoidInit":
		return handlers.NewSlibInitHandler(env), nil
	case "VoidNoop":
		return handlers.NewSlibNoopHandler(env), nil
	default:
		return nil, fmt.Errorf("Unknown function name: %s", funcName)
	}
}

func (f *funcHandlerFactory) GrpcNew(env types.Environment, service string) (types.GrpcFuncHandler, error) {
	return nil, fmt.Errorf("Not implemented")
}

func main() {
	faas.Serve(&funcHandlerFactory{})
}
