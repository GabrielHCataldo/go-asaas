package asaas

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type request[T any] struct {
	ctx         context.Context
	env         Env
	accessToken string
}

type Request[T any] interface {
	make(method string, path string, payload any) (*T, Error)
	makeMultipartForm(method string, path string, payload any) (*T, Error)
}

func NewRequest[T any](ctx context.Context, env Env, accessToken string) Request[T] {
	return request[T]{
		ctx:         ctx,
		env:         env,
		accessToken: accessToken,
	}
}

func (r request[T]) make(method string, path string, payload any) (*T, Error) {
	req, err := r.createHttpRequest(r.ctx, method, path, payload)
	if err != nil {
		return nil, NewByError(err)
	}
	logInfoSkipCaller(6, r.env, "request url:", req.URL.String())
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, NewByError(err)
	}
	defer r.closeBody(res.Body)
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusBadRequest {
		var result T
		err = r.readResponse(res, &result)
		if err != nil {
			return nil, NewByError(err)
		}
		return &result, nil
	}
	return nil, NewError(ERROR_UNEXPECTED, "response status code not expected:", res.StatusCode)
}

func (r request[T]) makeMultipartForm(method string, path string, payload any) (*T, Error) {
	multipartPayload, err := r.prepareMultipartPayload(payload)
	if err != nil {
		return nil, NewByError(err)
	}
	req, err := r.createHttpRequestMultipartForm(r.ctx, method, path, multipartPayload)
	if err != nil {
		return nil, NewByError(err)
	}
	logInfoSkipCaller(6, r.env, "request url:", req.URL.String())
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, NewByError(err)
	}
	defer r.closeBody(res.Body)
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusBadRequest {
		var result T
		err = r.readResponse(res, &result)
		if err != nil {
			return nil, NewByError(err)
		}
		return &result, nil
	}
	return nil, NewError(ERROR_UNEXPECTED, "response status code not expected:", res.StatusCode)
}

func (r request[T]) createHttpRequest(ctx context.Context, method string, path string, payload any) (
	*http.Request, error) {
	var payloadToSend io.Reader
	if payload != nil {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		logInfoSkipCaller(6, r.env, "sending request body:", string(payloadBytes))
		payloadToSend = bytes.NewReader(payloadBytes)
	}
	req, err := http.NewRequestWithContext(ctx, method, r.env.BaseURL()+path, payloadToSend)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", r.accessToken)
	return req, nil
}

func (r request[T]) createHttpRequestMultipartForm(
	ctx context.Context,
	method string,
	path string,
	values map[string]io.Reader,
) (req *http.Request, err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for k, v := range values {
		err = r.prepareMultipartWriter(w, k, v)
		if err != nil {
			return nil, err
		}
	}
	defer r.closeWriter(w)
	logInfoSkipCaller(6, r.env, "request body:", strings.ReplaceAll(b.String(), "\n", ""))
	req, err = http.NewRequestWithContext(ctx, method, r.env.BaseURL()+path, &b)
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
	logInfoSkipCaller(6, r.env, "response body:", string(respBody))
	return json.Unmarshal(respBody, result)
}

func (r request[T]) prepareMultipartPayload(payload any) (map[string]io.Reader, error) {
	multipartPayload := map[string]io.Reader{}
	for _, field := range structs.Fields(payload) {
		k := strings.Replace(field.Tag("json"), ",omitempty", "", 1)
		v := field.Value()
		if b, ok := v.(bool); ok {
			multipartPayload[k] = strings.NewReader(strconv.FormatBool(b))
		} else if s, ok := v.(string); ok {
			multipartPayload[k] = strings.NewReader(s)
		} else if f, ok := v.(*os.File); ok {
			multipartPayload[k] = f
		} else if fs, ok := v.([]*os.File); ok {
			for _, file := range fs {
				multipartPayload[file.Name()] = file
			}
		} else {
			multipartPayload[k] = strings.NewReader(fmt.Sprintf(`%s`, v))
		}
	}
	return multipartPayload, nil
}

func (r request[T]) prepareMultipartWriter(form *multipart.Writer, k string, reader io.Reader) (err error) {
	var fw io.Writer
	if c, ok := reader.(io.Closer); ok {
		defer r.closeCloser(c)
	}
	if file, ok := reader.(*os.File); ok {
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
