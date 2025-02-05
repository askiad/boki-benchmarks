package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	// "log"

	"cs.utexas.edu/zjia/faas/slib/statestore"
	"cs.utexas.edu/zjia/faas/types"
)

type PostInput struct {
	UserId string `json:"userId"`
	Body   string `json:"body"`
}

type PostOutput struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type postHandler struct {
	kind   string
	env    types.Environment
}

func NewSlibPostHandler(env types.Environment) types.FuncHandler {
	return &postHandler{
		kind: "slib",
		env:  env,
	}
}

func postSlib(ctx context.Context, env types.Environment, input *PostInput) (*PostOutput, error) {
	txn, err := statestore.CreateTxnEnv(ctx, env)
	if err != nil {
		return nil, err
	}

	userObj := txn.Object(fmt.Sprintf("userid:%s", input.UserId))

	prevCnt := "-1"
	if value, _ := userObj.Get("counter"); value.IsNull() {
		txn.TxnAbort()
		return &PostOutput{
			Success: false,
			Message: fmt.Sprintf("Cannot find str field with ID %s", input.UserId),
		}, nil
	} else {
		prevCnt = value.AsString()
	}

	cnt, _ := strconv.Atoi(prevCnt)
	cnt = (cnt + 1) % 1000
	newCnt := strconv.Itoa(cnt)
	userObj.SetString("counter", newCnt)

	// log.Printf("Request for counter with ID %s: counter was %s and setting to %s",
	// 	input.UserId, prevCnt, newCnt)

	if committed, err := txn.TxnCommit(); err != nil {
		return nil, err
	} else if !committed {
		return &PostOutput{
			Success: false,
			Message: "Failed to commit transaction due to conflicts",
		}, nil
	}

	return &PostOutput{Success: true}, nil
}

func (h *postHandler) onRequest(ctx context.Context, input *PostInput) (*PostOutput, error) {
	switch h.kind {
	case "slib":
		return postSlib(ctx, h.env, input)
	default:
		panic(fmt.Sprintf("Unknown kind: %s", h.kind))
	}
}

func (h *postHandler) Call(ctx context.Context, input []byte) ([]byte, error) {
	parsedInput := &PostInput{}
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
