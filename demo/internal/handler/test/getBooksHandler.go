package test

import (
	"net/http"

	"github.com/shyandsy/shygoctl/demo/internal/logic/test"
	"github.com/shyandsy/shygoctl/demo/internal/svc"
	"github.com/shyandsy/shygoctl/demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// get books
func GetBooksHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBookReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := test.NewGetBooksLogic(r.Context(), svcCtx)
		resp, err := l.GetBooks(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
