package product

import (
	"kefu-server/internal/service"
)

type sProduct struct{}

func init() {
	service.RegisterProduct(&sProduct{})
}
