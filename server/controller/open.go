package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"picbed-go/model/response"
	"picbed-go/setting"
	"picbed-go/uploader"
	"picbed-go/uploader/template"
)

// InfoHandler 网站信息
func InfoHandler(c *gin.Context) {

	response.Ok(c, "", map[string]interface{}{
		"name":   setting.App.Name,
		"notice": setting.App.Notice,
		"uploaders": []map[string]string{
			{"id": "bli", "name": "哔哩哔哩"},
			{"id": "mt", "name": "美团"},
		},
	})
}

// UploadHandler 上传文件
func UploadHandler(c *gin.Context) {
	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		response.Fail(c, http.StatusOK, response.FileNotFound)
		return
	}

	// 获取uploader
	id := c.Param("id")
	uploader, ok := uploader.Mapping[id]
	if !ok {
		zap.L().Debug("uploader不存在")
		response.Fail(c, http.StatusOK, response.UploaderNotFound)
		return
	}

	// 构造UploadRequest
	open, err := file.Open()
	if err != nil {
		return
	}
	defer open.Close()
	fileBytes := make([]byte, file.Size)
	open.Read(fileBytes)
	uploadReq := &template.UploadRequest{
		FileBytes: fileBytes,
		Filename:  file.Filename,
	}

	// 上传
	uploadResponse := uploader.Upload(uploadReq)
	response.Ok(c, "上传成功", uploadResponse)
}
