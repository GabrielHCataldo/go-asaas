package asaas

import (
	"encoding/base64"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"net/url"
	"strings"
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

func (f *FileRequest) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == `""` {
		return nil
	}
	split := strings.Split(s, ";")
	if len(split) != 3 {
		return nil
	}
	mimeType := FileMimeType(strings.Replace(split[0], "data:", "", 1))
	name := strings.Replace(split[1], "name=", "", 1)
	data, err := base64.StdEncoding.DecodeString(strings.Replace(split[2], "base64,", "", 1))
	if err != nil {
		return err
	}
	*f = FileRequest{
		Name: name,
		Mime: mimeType,
		Data: data,
	}
	return err
}

func (f FileRequest) MarshalJSON() ([]byte, error) {
	result := `null`
	if util.IsNotBlank(&f.Name) && f.Mime.IsEnumValid() && len(f.Data) > 0 {
		name := url.PathEscape(f.Name)
		b64 := base64.StdEncoding.EncodeToString(f.Data)
		result = fmt.Sprintf(`"data:%s;name=%s;base64,%s"`, f.Mime.String(), name, b64)
	}
	return []byte(result), nil
}
