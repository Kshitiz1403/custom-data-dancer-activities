package common

import (
	"fmt"

	"github.com/finbox-in/data-dancer/v2/activities"
)

type SampleActivities struct {
	activities.BaseActivity
}

func (a *SampleActivities) SampleActivity(data string, args map[string]any) (interface{}, error) {
	fmt.Println("SampleActivity", data)
	fmt.Println("Activity name", a.GetActivityName())
	return nil, nil
}
