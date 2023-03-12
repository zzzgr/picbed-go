package template

import "time"

type UploadRequest struct {
	FileBytes []byte
	Filename  string
}

type UploadResponse struct {
	Url      string `json:"url"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	Time     int64  `json:"time"`
}

func NewUploadResponse(url string, filename string, size int64) *UploadResponse {
	return &UploadResponse{
		Url:      url,
		Filename: filename,
		Size:     size,
		Time:     time.Now().Unix(),
	}
}
