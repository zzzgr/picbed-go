package template

type IUploader interface {
	Upload(request *UploadRequest) *UploadResponse
}
