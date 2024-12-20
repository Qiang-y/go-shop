package category

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"biz-demo/gomall/app/frontend/biz/service"
	"biz-demo/gomall/app/frontend/biz/utils"
	category "biz-demo/gomall/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Category .
// @router /category/:category [GET]
func Category(ctx context.Context, c *app.RequestContext) {
	var err error
	var req category.CategoryReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//resp := &common.Empty{}
	resp, err := service.NewCategoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	hlog.Infof("resp products: %v", resp)
	c.HTML(consts.StatusOK, "category", utils.WarpResponse(ctx, c, resp))
}
