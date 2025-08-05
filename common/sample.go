package common

import (
	"context"
	"fmt"

	"github.com/finbox-in/data-dancer/v2/activities"
)

type SampleActivities struct {
	activities.BaseActivity
}

func (a *SampleActivities) SampleActivity(ctx context.Context, args map[string]any) (interface{}, error) {
	fmt.Println("SampleActivity", args)
	activityInfo := activities.GetInfo(ctx)
	fmt.Println("Activity name", activityInfo.ActivityName)
	fmt.Println("Activity name", a.GetActivityName())
	return nil, nil
}
