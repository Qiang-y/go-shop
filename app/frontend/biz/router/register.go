// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	auth "biz-demo/gomall/app/frontend/biz/router/auth"
	cart "biz-demo/gomall/app/frontend/biz/router/cart"
	category "biz-demo/gomall/app/frontend/biz/router/category"
	checkout "biz-demo/gomall/app/frontend/biz/router/checkout"
	home "biz-demo/gomall/app/frontend/biz/router/home"
	order "biz-demo/gomall/app/frontend/biz/router/order"
	product "biz-demo/gomall/app/frontend/biz/router/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	order.Register(r)

	checkout.Register(r)

	cart.Register(r)

	product.Register(r)

	auth.Register(r)

	category.Register(r)

	home.Register(r)

}
