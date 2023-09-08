package routers

import (
	"pass-crypto/routers/private"
	"pass-crypto/routers/public"
)

type PublicRouterGroup struct {
	public.LoginRouter
	public.MyRouter
	public.RegisterRouter
}
type PrivateRouterGroup struct {
	private.PageRouter
}

// PublicRouterApp public route request
var PublicRouterApp = new(PublicRouterGroup)

// PrivateRouterApp private route request
var PrivateRouterApp = new(PrivateRouterGroup)
