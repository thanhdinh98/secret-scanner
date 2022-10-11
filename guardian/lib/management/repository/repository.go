package repository

import (
	"context"

	"guardian/guardian/lib/management/repository/state"
)

func AddNewForUser(ctx context.Context, state *state.State, email string) error {
	return nil
}

func TriggerScanByUser(
	ctx context.Context, state *state.State,
	repoTitle, email string,
) error {
	return nil
}

func ReturnScanResultForUser(
	ctx context.Context, state *state.State,
	email string,
) error {
	return nil
}
