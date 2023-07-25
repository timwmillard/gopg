package main

import "context"

func QueryExec[ParamsT, ReturnT any](ctx context.Context, args ParamsT) (ReturnT, error) {
	var rt ReturnT
	return rt, nil
}
