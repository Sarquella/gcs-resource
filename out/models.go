package out

import (
	"github.com/frodenas/gcs-resource"
)

type OutRequest struct {
	Source gcsresource.Source `json:"source"`
	Params Params             `json:"params"`
}

type Params struct {
	File             string `json:"file"`
	PredefinedACL    string `json:"predefined_acl"`
	ContentType      string `json:"content_type"`
	CacheControl     string `json:"cache_control"`
	DynamicPathValue string `json:"dynamic_path_value"`
	DynamicPathFile  string `json:"dynamic_path_file"`
}

func (params Params) IsValid() (bool, string) {
	if params.File == "" {
		return false, "please specify the file"
	}

	return true, ""
}

type OutResponse struct {
	Version  gcsresource.Version        `json:"version"`
	Metadata []gcsresource.MetadataPair `json:"metadata"`
}
