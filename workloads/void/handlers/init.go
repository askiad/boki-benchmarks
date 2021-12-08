package handlers

import (
	"context"
	"fmt"

	"cs.utexas.edu/zjia/faas/types"
)

type initHandler struct {
	kind   string
	env    types.Environment
}

func NewSlibInitHandler(env types.Environment) types.FuncHandler {
	return &initHandler{
		kind: "slib",
		env:  env,
	}
}

func initSlib(ctx context.Context, env types.Environment) error {
	return nil
}

func (h *initHandler) Call(ctx context.Context, input []byte) ([]byte, error) {
	var err error
	switch h.kind {
	case "slib":
		err = initSlib(ctx, h.env)
	default:
		panic(fmt.Sprintf("Unknown kind: %s", h.kind))
	}

	if err != nil {
		return nil, err
	} else {
		return []byte("Init done\n"), nil
	}
}
