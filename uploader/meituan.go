package uploader

import (
	"bytes"
	"errors"
	"fmt"
	"go.uber.org/zap"
	http2 "net/http"
	"picbed-go/uploader/template"
	"picbed-go/util/http"
	"strconv"
	"strings"
	"time"
)

type MtUploader struct {
}

func (u *MtUploader) Upload(request *template.UploadRequest) *template.UploadResponse {
	// 获取visitId
	visitId := getVisitId()

	// 上传图片
	var mtRes mtResponse
	res, err := http.NewRestyClient().R().
		SetResult(&mtRes).
		SetHeaders(map[string]string{"CSC-VisitId": visitId}).
		SetFileReader("files", request.Filename, bytes.NewReader(request.FileBytes)).
		SetFormData(map[string]string{
			"fileName": request.Filename,
			"part":     "0",
			"partSize": "1",
			"fileID":   strconv.FormatInt(time.Now().UnixMilli(), 10),
		}).
		Execute("POST", "https://kf.dianping.com/api/file/burstUploadFile")
	if err != nil {
		panic(errors.New(err.Error()))
	}

	zap.L().Debug(fmt.Sprintf("mt upload raw res, status: %d, res: %#v", res.StatusCode(), string(res.Body())))

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
	if !mtRes.Success {
		panic(errors.New(mtRes.ErrMsg))
	}

	return template.NewUploadResponse(mtRes.Data.UploadPath, request.Filename, int64(len(request.FileBytes)))
}

func getVisitId() string {
	client := http.NoRedirectClient.R()
	// 获取visitId
	res, err := client.Execute("GET", "https://kf.dianping.com/csCenter/access/dealOrder_Help_DP_PC")
	if err != nil {
		errors.New("获取visitId异常")
	}

	location := res.Header().Get("Location")
	index := strings.Index(location, "?")
	param := location[index+1:]
	var visitId string
	for _, v := range strings.Split(param, "&") {
		split := strings.Split(v, "=")
		if split[0] == "visitId" {
			visitId = split[1]
			break
		}
	}
	return visitId
}

type mtResponse struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
	Data    struct {
		UploadPath string `json:"uploadPath"`
	} `json:"data"`
}
