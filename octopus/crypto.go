package octopus

import (
	"fmt"

	"github.com/finbox-in/data-dancer/v2/activities"
)

// CryptoActivities contains all crypto-related workflow activities
type CryptoActivities struct {
	activities.BaseActivity
	// You can inject dependencies here
}

func (a *CryptoActivities) Encrypt(data string, args map[string]any) (interface{}, error) {
	fmt.Println("Encrypting data", data)
	fmt.Println("Activity name", a.GetActivityName())
	return nil, nil
}

func (a *CryptoActivities) Decrypt(data string, args map[string]any) (interface{}, error) {
	fmt.Println("Decrypting data", data)
	fmt.Println("Activity name", a.GetActivityName())
	return nil, nil
}
