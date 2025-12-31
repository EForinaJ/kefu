package site

import (
	"kefu-server/internal/service"
)

type sSite struct{}

func init() {
	service.RegisterSite(&sSite{})
}
