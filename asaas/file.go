package asaas

import (
	"encoding/base64"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"net/url"
)

type FileRequest struct {
	Name string       `json:"name"`
	Mime FileMimeType `json:"mime"`
	Data []byte       `json:"data"`
}

type FileTextPlainResponse struct {
	Data   string          `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

func (f FileRequest) MarshalJSON() ([]byte, error) {
	result := `null`
	if util.IsNotBlank(&f.Name) && f.Mime != "" && len(f.Data) > 0 {
		name := url.PathEscape(f.Name)
		b64 := base64.StdEncoding.EncodeToString(f.Data)
		result = fmt.Sprintf(`"data:%s;name=%s;base64,%s"`, f.Mime.String(), name, b64)
	}
	return []byte(result), nil
}
