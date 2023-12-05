package asaas

import (
	"bytes"
	"context"
	"encoding/json"
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
	r.readResponse(res, &respBody)
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
		r.readResponse(res, &result)
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
				k := util.GetJsonFieldNameByReflect(ft)
				v := util.GetValueByReflect(f)
				if util.IsBlank(&k) || v == nil {
					continue
				}
				params.Add(k, fmt.Sprintf(`%s`, v))
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
	req, _ := http.NewRequestWithContext(ctx, method, rUrl, payloadToSend)
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
	_ = Body.Close()
}

func (r request[T]) closeWriter(writer *multipart.Writer) {
	_ = writer.Close()
}

func (r request[T]) closeCloser(c io.Closer) {
	_ = c.Close()
}

func (r request[T]) readResponse(res *http.Response, result *T) {
	respBody, _ := io.ReadAll(res.Body)
	logInfoSkipCaller(6, r.env, "response status:", res.StatusCode, "body:", string(respBody))
	if x, ok := any(*result).(FileTextPlainResponse); ok && res.StatusCode == http.StatusOK {
		x.Data = string(respBody)
		*result = any(x).(T)
	} else if util.IsJson(respBody) {
		_ = json.Unmarshal(respBody, result)
	}
}

func (r request[T]) prepareMultipartPayload(payload any) (map[string][]io.Reader, error) {
	rPayload := reflect.ValueOf(payload)
	rtPayload := reflect.TypeOf(payload)
	multipartPayload := map[string][]io.Reader{}
	for i := 0; i < rPayload.NumField(); i++ {
		fd := rPayload.Field(i)
		ft := rtPayload.Field(i)
		if fd.IsZero() || !fd.IsValid() {
			continue
		}
		k := util.GetJsonFieldNameByReflect(ft)
		vf := util.GetValueByReflect(fd)
		var b bool
		var s string
		var in int
		var f *os.File
		var fs []*os.File
		var ok bool
		if b, ok = vf.(bool); ok {
			multipartPayload[k] = []io.Reader{strings.NewReader(strconv.FormatBool(b))}
		} else if s, ok = vf.(string); ok {
			multipartPayload[k] = []io.Reader{strings.NewReader(s)}
		} else if in, ok = vf.(int); ok {
			multipartPayload[k] = []io.Reader{strings.NewReader(strconv.Itoa(in))}
		} else if f, ok = vf.(*os.File); ok && f != nil {
			multipartPayload[k] = []io.Reader{f}
		} else if fs, ok = vf.([]*os.File); ok && fs != nil {
			var files []io.Reader
			for _, file := range fs {
				if file != nil {
					files = append(files, file)
				}
			}
			multipartPayload[k] = files
		} else if vf != nil {
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
	if fv.Kind() == reflect.Slice {
		fv.Set(reflect.MakeSlice(fv.Type(), 1, 1))
		fv.Index(0).Set(reflect.ValueOf(ErrorResponse{
			Code:        res.Status,
			Description: "response status code not expected",
		}))
	}
	return &respBody, nil
}
