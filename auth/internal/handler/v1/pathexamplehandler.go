package v1

import (
	"net/http"

	"auth/internal/logic/v1"
	"auth/internal/svc"
	"auth/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PathExampleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathExampleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := v1.NewPathExampleLogic(r.Context(), svcCtx)
		resp, err := l.PathExample(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
