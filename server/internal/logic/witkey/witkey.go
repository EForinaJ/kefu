package witkey

import "kefu-server/internal/service"

type sWitkey struct{}

func init() {
	service.RegisterWitkey(&sWitkey{})
}
