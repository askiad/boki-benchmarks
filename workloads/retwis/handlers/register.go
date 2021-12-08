package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"cs.utexas.edu/zjia/faas/slib/statestore"
	"cs.utexas.edu/zjia/faas/types"
)

type RegisterInput struct {
	UserName string `json:"username"`
}

type RegisterOutput struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	UserId  string `json:"userId"`
}

type registerHandler struct {
	kind   string
	env    types.Environment
}

func NewSlibRegisterHandler(env types.Environment) types.FuncHandler {
	return &registerHandler{
		kind: "slib",
		env:  env,
	}
}

func registerSlib(ctx context.Context, env types.Environment, input *RegisterInput) (*RegisterOutput, error) {
	store := statestore.CreateEnv(ctx, env)
	nextUserIdObj := store.Object("next_user_id")
	result := nextUserIdObj.NumberFetchAdd("value", 1)
	if result.Err != nil {
		return nil, result.Err
	}
	userIdValue := uint32(result.Value.AsNumber())

	txn, err := statestore.CreateTxnEnv(ctx, env)
	if err != nil {
		return nil, err
	}

	userNameObj := txn.Object(fmt.Sprintf("username:%s", input.UserName))
	if value, _ := userNameObj.Get("id"); !value.IsNull() {
		txn.TxnAbort()
		return &RegisterOutput{
			Success: false,
			Message: fmt.Sprintf("User name \"%s\" already exists", input.UserName),
		}, nil
	}

	userId := fmt.Sprintf("%08x", userIdValue)
	userNameObj.SetString("id", userId)

	userObj := txn.Object(fmt.Sprintf("userid:%s", userId))
	userObj.SetString("username", input.UserName)
	userObj.SetString("counter", "0")

	if committed, err := txn.TxnCommit(); err != nil {
		return nil, err
	} else if committed {
		log.Printf("[counter-log] Registered counter with ID %s", userId)
		return &RegisterOutput{
			Success: true,
			UserId:  userId,
		}, nil
	} else {
		return &RegisterOutput{
			Success: false,
			Message: "Failed to commit transaction due to conflicts",
		}, nil
	}
}

func (h *registerHandler) onRequest(ctx context.Context, input *RegisterInput) (*RegisterOutput, error) {
	switch h.kind {
	case "slib":
		return registerSlib(ctx, h.env, input)
	default:
		panic(fmt.Sprintf("Unknown kind: %s", h.kind))
	}
}

func (h *registerHandler) Call(ctx context.Context, input []byte) ([]byte, error) {
	parsedInput := &RegisterInput{}
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
