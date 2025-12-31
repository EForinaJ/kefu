package auth

import (
	"kefu-server/internal/service"
)

type sAuth struct{}

func init() {
	service.RegisterAuth(&sAuth{})
}
