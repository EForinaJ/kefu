package dashboard

import (
	"server/internal/service"
)

type sDashboard struct{}

func init() {
	service.RegisterDashboard(&sDashboard{})
}
