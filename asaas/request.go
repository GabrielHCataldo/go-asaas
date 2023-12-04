package asaas

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GabrielHCataldo/go-asaas/internal/util"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type request[T any] struct {
	ctx         context.Context
	env         Env
	accessToken string
}

type Request[T any] interface {
	make(method string, path string, payload any) (*T, error)
	makeMultipartForm(method string, path string, payload any) (*T, error)
}

func NewRequest[T any](ctx context.Context, env Env, accessToken string) Request[T] {
	return request[T]{
		ctx:         ctx,
		env:         env,
		accessToken: accessToken,
	}
}

func (r request[T]) make(method string, path string, payload any) (*T, error) {
	req, err := r.createHttpRequest(r.ctx, method, path, payload)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.closeBody(res.Body)
	var respBody T
	err = r.readResponse(res, &respBody)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK ||
		res.StatusCode == http.StatusBadRequest ||
		(res.StatusCode == http.StatusNotFound && (method == http.MethodGet || method == http.MethodPut)) {
		return &respBody, nil
	}
	return r.prepareResponseUnexpected(res)
}

func (r request[T]) makeMultipartForm(method string, path string, payload any) (*T, error) {
	multipartPayload, err := r.prepareMultipartPayload(payload)
	if err != nil {
		return nil, err
	}
	req, err := r.createHttpRequestMultipartForm(r.ctx, method, path, multipartPayload)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.closeBody(res.Body)
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusBadRequest {
		var result T
		err = r.readResponse(res, &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
	return r.prepareResponseUnexpected(res)
}

func (r request[T]) createHttpRequest(ctx context.Context, method string, path string, payload any) (
	*http.Request, error) {
	rUrl := r.env.BaseUrl() + path
	var payloadToSend io.Reader
	var payloadBytes []byte
	var err error
	if payload != nil {
		switch method {
		case http.MethodGet, http.MethodDelete:
			rPayload := reflect.ValueOf(payload)
			rtPayload := reflect.TypeOf(payload)
			params := url.Values{}
			for i := 0; i < rPayload.NumField(); i++ {
				f := rPayload.Field(i)
				ft := rtPayload.Field(i)
				if f.IsNil() || f.IsZero() || !f.IsValid() {
					continue
				}
				k := strings.Split(ft.Tag.Get("json"), ",")
				if len(k) > 0 {
					v := f.Elem().Interface()
					params.Add(k[0], fmt.Sprintf(`%s`, v))
				}
			}
			encode := params.Encode()
			if util.IsNotBlank(&encode) {
				rUrl += "?" + encode
			}
			break
		default:
			payloadBytes, err = json.Marshal(payload)
			if err != nil {
				return nil, err
			}
			payloadToSend = bytes.NewReader(payloadBytes)
		}
	}
	logInfoSkipCaller(5, r.env, "request url:", rUrl, "method:", method)
	if len(payloadBytes) > 0 {
		logInfoSkipCaller(5, r.env, "request body:", string(payloadBytes))
	}
	req, err := http.NewRequestWithContext(ctx, method, rUrl, payloadToSend)
	if err != nil {
		return nil, err
	}
	var t T
	accept := HttpContentTypeJSON
	switch any(t).(type) {
	case string, FileTextPlainResponse:
		accept = HttpContentTypeText
	}
	req.Header.Add("Accept", accept)
	if method != http.MethodGet && method != http.MethodDelete && accept == HttpContentTypeJSON {
		req.Header.Add("Content-Type", HttpContentTypeJSON)
	}
	req.Header.Add("access_token", r.accessToken)
	return req, nil
}

func (r request[T]) createHttpRequestMultipartForm(
	ctx context.Context,
	method string,
	path string,
	values map[string][]io.Reader,
) (req *http.Request, err error) {
	rUrl := r.env.BaseUrl() + path
	logInfoSkipCaller(5, r.env, "request url:", rUrl, "method:", method)
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for k, v := range values {
		for _, reader := range v {
			err = r.prepareMultipartWriter(w, k, reader)
			if err != nil {
				return nil, err
			}
		}
	}
	logInfoSkipCaller(5, r.env, "request body:", util.ReplaceAllSpacesRepeat(b.String()))
	defer r.closeWriter(w)
	req, err = http.NewRequestWithContext(ctx, method, rUrl, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Add("access_token", r.accessToken)
	req.ContentLength = 0
	return req, nil
}

func (r request[T]) closeBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		logError("error close read body:", err)
	}
}

func (r request[T]) closeWriter(writer *multipart.Writer) {
	err := writer.Close()
	if err != nil {
		logError("error close writer:", err)
	}
}

func (r request[T]) closeCloser(c io.Closer) {
	err := c.Close()
	if err != nil {
		logError("error close reader closer:", err)
	}
}

func (r request[T]) readResponse(res *http.Response, result *T) error {
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	logInfoSkipCaller(6, r.env, "response status:", res.StatusCode, "body:", string(respBody))
	if len(respBody) == 0 {
		return nil
	} else if x, ok := any(*result).(FileTextPlainResponse); ok {
		x.Data = string(respBody)
		*result = any(x).(T)
		return nil
	}
	return json.Unmarshal(respBody, result)
}

func (r request[T]) prepareMultipartPayload(payload any) (map[string][]io.Reader, error) {
	rPayload := reflect.ValueOf(payload)
	rtPayload := reflect.TypeOf(payload)

	multipartPayload := map[string][]io.Reader{}
	for i := 0; i < rPayload.NumField(); i++ {
		field := rPayload.Field(i)
		fieldType := rtPayload.Field(i)
		if field.IsNil() || field.IsZero() || !field.IsValid() {
			continue
		}
		sk := strings.Split(fieldType.Tag.Get("json"), ",")
		if len(sk) == 0 {
			continue
		}
		k := sk[0]
		vf := field.Elem().Interface()
		var b bool
		var s string
		var f *os.File
		var fs []*os.File
		var ok bool
		if b, ok = vf.(bool); ok {
			multipartPayload[k] = []io.Reader{strings.NewReader(strconv.FormatBool(b))}
		} else if s, ok = vf.(string); ok {
			multipartPayload[k] = []io.Reader{strings.NewReader(s)}
		} else if f, ok = vf.(*os.File); ok && f != nil {
			multipartPayload[k] = []io.Reader{f}
		} else if fs, ok = vf.([]*os.File); ok && fs != nil {
			var files []io.Reader
			for _, file := range fs {
				if file == nil {
					continue
				}
				files = append(files, file)
			}
			multipartPayload[k] = files
		} else {
			multipartPayload[k] = []io.Reader{strings.NewReader(fmt.Sprintf(`%s`, vf))}
		}
	}
	return multipartPayload, nil
}

func (r request[T]) prepareMultipartWriter(form *multipart.Writer, k string, reader io.Reader) (err error) {
	var fw io.Writer
	if c, ok := reader.(io.Closer); ok {
		defer r.closeCloser(c)
	}
	if file, ok := reader.(*os.File); ok && file != nil {
		if fw, err = form.CreateFormFile(k, file.Name()); err != nil {
			return err
		}
	} else if fw, err = form.CreateFormField(k); err != nil {
		return err
	}
	if _, err = io.Copy(fw, reader); err != nil {
		return err
	}
	return
}

func (r request[T]) prepareResponseUnexpected(res *http.Response) (*T, error) {
	var respBody T
	rv := reflect.ValueOf(&respBody)
	fv := rv.Elem().FieldByName("Errors")
	if rv.Kind() == reflect.Struct && fv.Kind() != reflect.Slice {
		return nil, errors.New("poorly formatted response structure, does not contain the errors field")
	} else if fv.Kind() == reflect.Slice {
		fv.Set(reflect.MakeSlice(fv.Type(), 1, 1))
		fv.Index(0).Set(reflect.ValueOf(ErrorResponse{
			Code:        res.Status,
			Description: "response status code not expected",
		}))
	}
	return &respBody, nil
}
