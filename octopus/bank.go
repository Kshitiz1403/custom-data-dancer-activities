package octopus

import (
	"context"

	"github.com/finbox-in/data-dancer/v2/activities"
)

// BankActivities contains all bank-related workflow activities
type BankActivities struct {
	activities.BaseActivity
	// You can inject dependencies here
}

func (a *BankActivities) ValidateAccount(ctx context.Context, args map[string]any) (interface{}, error) {
	return nil, nil
}

func (a *BankActivities) GetBalance(ctx context.Context, args map[string]any) (interface{}, error) {
	return nil, nil
}
