package main

import (
	"fmt"

	"cs.utexas.edu/zjia/faas-retwis/handlers"

	"cs.utexas.edu/zjia/faas"
	"cs.utexas.edu/zjia/faas/types"
)

type funcHandlerFactory struct {
}

func (f *funcHandlerFactory) New(env types.Environment, funcName string) (types.FuncHandler, error) {
	switch funcName {
	case "RetwisInit":
		return handlers.NewSlibInitHandler(env), nil
	case "RetwisRegister":
		return handlers.NewSlibRegisterHandler(env), nil
	case "RetwisPost":
		return handlers.NewSlibPostHandler(env), nil
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
