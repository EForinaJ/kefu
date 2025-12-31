package order

import "kefu-server/internal/service"

type sOrder struct{}

func init() {
	service.RegisterOrder(&sOrder{})
}
