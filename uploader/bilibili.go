package uploader

import (
	"bytes"
	"errors"
	"fmt"
	"go.uber.org/zap"
	http2 "net/http"
	"picbed-go/setting"
	"picbed-go/uploader/template"
	"picbed-go/util/http"
	strUtil "picbed-go/util/string"
)

type BliUploader struct {
}

func (u *BliUploader) Upload(request *template.UploadRequest) *template.UploadResponse {
	var bliRes BliResponse
	res, err := http.NewRestyClient().R().
		SetResult(&bliRes).
		SetFileReader("binary", request.Filename, bytes.NewReader(request.FileBytes)).
		SetFormData(map[string]string{
			"csrf":     u.getCsrf(),
			"biz":      "new_dyn",
			"category": "daily",
		}).
		SetHeader("Cookie", setting.App.Config.BliCookie).
		Execute("POST", "https://api.bilibili.com/x/article/creative/article/upcover")
	if err != nil {
		panic(errors.New(err.Error()))
	}

	zap.L().Debug(fmt.Sprintf("bilibili upload raw res, status: %d, res: %#v", res.StatusCode(), string(res.Body())))

	// http状态码
	statusCode := res.StatusCode()
	if statusCode != http2.StatusOK {
		if statusCode == http2.StatusRequestEntityTooLarge {
			panic(errors.New("文件过大"))
		} else {
			panic(errors.New("第三方接口异常"))
		}
	}

	// 业务码
	if bliRes.Code != 0 {
		panic(errors.New(bliRes.Message))
	}
	return template.NewUploadResponse(bliRes.Data.Url, request.Filename, int64(len(request.FileBytes)))
}

func (u *BliUploader) getCsrf() string {
	cookieMap := strUtil.SplitCookie(setting.App.Config.BliCookie)
	return cookieMap["bili_jct"]
}

type BliResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Url string `json:"url"`
	} `json:"data"`
}
