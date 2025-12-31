package account

import (
	"kefu-server/internal/service"
)

type sAccount struct{}

func init() {
	service.RegisterAccount(&sAccount{})
}
