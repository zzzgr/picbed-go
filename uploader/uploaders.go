package uploader

import "picbed-go/uploader/template"

var Mapping = map[string]template.IUploader{
	"bli": &BliUploader{},
	"mt":  &MtUploader{},
}
