package util

import (
	"context"
	"errors"
	pd "github.com/PagerDuty/go-pagerduty"
	"veric-backend/logic/config"
)

func SendPagerNotify(incidentKey, action string, msg string, severity string) error {
	serviceKey := config.Get().PagerDuty.Key
	if serviceKey == "" {
		return errors.New("pagerduty service key not found. skipped")
	}

	_, err := pd.ManageEventWithContext(context.Background(), pd.V2Event{
		RoutingKey: serviceKey,
		Action:     action,
		DedupKey:   incidentKey,
		Payload: &pd.V2Payload{
			Summary:  incidentKey,
			Severity: severity,
			Source:   "as-gateway",
			Details:  msg,
		},
	})
	return err
}
