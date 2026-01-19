package workorder

import (
	"server/internal/service"
)

type sWorkOrder struct {
}

func init() {
	service.RegisterWorkOrder(&sWorkOrder{})
}
