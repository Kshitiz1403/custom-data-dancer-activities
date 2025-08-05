package octopus

import (
	"context"
	"fmt"

	"github.com/finbox-in/data-dancer/v2/activities"
)

// CryptoActivities contains all crypto-related workflow activities
type CryptoActivities struct {
	activities.BaseActivity
	// You can inject dependencies here
}

func (a *CryptoActivities) Encrypt(ctx context.Context, args map[string]any) (interface{}, error) {
	activityInfo := activities.GetInfo(ctx)
	fmt.Println("Encrypting data", args)
	fmt.Println("Activity name", activityInfo.ActivityName)
	fmt.Println("Activity name", a.GetActivityName())
	return nil, nil
}

func (a *CryptoActivities) Decrypt(ctx context.Context, args map[string]any) (interface{}, error) {
	activityInfo := activities.GetInfo(ctx)
	fmt.Println("Decrypting data", args)
	fmt.Println("Activity name", activityInfo.ActivityName)
	fmt.Println("Activity name", a.GetActivityName())
	return nil, nil
}
