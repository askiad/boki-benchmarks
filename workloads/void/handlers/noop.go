package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"cs.utexas.edu/zjia/faas/types"
)

type NoopInput struct {
	Input string `json:"input"`
}

type NoopOutput struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type noopHandler struct {
	kind   string
	env    types.Environment
}

func NewSlibNoopHandler(env types.Environment) types.FuncHandler {
	return &noopHandler{
		kind: "slib",
		env:  env,
	}
}

func noopSlib(ctx context.Context, env types.Environment, input *NoopInput) (*NoopOutput, error) {
	output := &NoopOutput{
		Success: true,
	}
	return output, nil
}

func (h *noopHandler) onRequest(ctx context.Context, input *NoopInput) (*NoopOutput, error) {
	switch h.kind {
	case "slib":
		return noopSlib(ctx, h.env, input)
	default:
		panic(fmt.Sprintf("Unknown kind: %s", h.kind))
	}
}

func (h *noopHandler) Call(ctx context.Context, input []byte) ([]byte, error) {
	parsedInput := &NoopInput{}
	err := json.Unmarshal(input, parsedInput)
	if err != nil {
		return nil, err
	}
	output, err := h.onRequest(ctx, parsedInput)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}
