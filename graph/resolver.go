package graph

import "sms-campaign-manager/config"

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *config.DB
}
