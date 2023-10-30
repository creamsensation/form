package form

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

const (
	requestTypeForm = iota
	requestTypeMultipartForm
)

func processRequest(req *http.Request, limit int) (url.Values, map[string][]*multipart.FileHeader, error) {
	requestType, err := parseForm(req, limit)
	if err != nil {
		return createEmptyProcessRequestResult(err)
	}
	if requestType == requestTypeMultipartForm {
		return req.MultipartForm.Value, req.MultipartForm.File, nil
	}
	return req.Form, make(map[string][]*multipart.FileHeader), nil
}

func processFormData(form *FormBuilder, data url.Values) {
	for i, field := range form.fields {
		for name, item := range data {
			if len(item) == 0 || name != field.name {
				continue
			}
			switch field.dataType {
			case fieldDataTypeString:
				if !field.multiple {
					form.fields[i].value = item[0]
				}
				if field.multiple {
					form.fields[i].value = item
				}
			case fieldDataTypeFloat:
				if !field.multiple {
					form.fields[i].value = convertToFloat(item[0])
				}
				if field.multiple {
					form.fields[i].value = convertSlice[string, float64](
						item, func(v string) float64 {
							return convertToFloat(v)
						},
					)
				}
			case fieldDataTypeInt:
				if !field.multiple {
					form.fields[i].value = convertToInt(item[0])
				}
				if field.multiple {
					form.fields[i].value = convertSlice[string, int](
						item, func(v string) int {
							return convertToInt(v)
						},
					)
				}
			case fieldDataTypeBool:
				if !field.multiple {
					form.fields[i].value = item[0] == "true"
				}
				if field.multiple {
					form.fields[i].value = convertSlice[string, bool](
						item, func(v string) bool {
							return v == "true"
						},
					)
				}
			}
		}
	}
}

func processFormFiles(form *FormBuilder, multipartFiles map[string][]*multipart.FileHeader) error {
	requestFiles := make([]Multipart, 0)
	for key, files := range multipartFiles {
		for _, file := range files {
			f, err := file.Open()
			if err != nil {
				return fmt.Errorf("error while opening multipart file: %w", err)
			}
			fileBytes, err := io.ReadAll(f)
			if err != nil {
				return fmt.Errorf("error while reading multipart file: %w", err)
			}
			requestFiles = append(
				requestFiles, Multipart{
					Key:    key,
					Name:   file.Filename,
					Type:   http.DetectContentType(fileBytes),
					Suffix: getFileSuffixFromName(file.Filename),
					Bytes:  fileBytes,
				},
			)
		}
	}
	for i, field := range form.fields {
		for _, file := range requestFiles {
			if file.Key != field.name {
				continue
			}
			if !field.multiple {
				form.fields[i].value = file
			}
			if field.multiple {
				form.fields[i].value = append(form.fields[i].value.([]Multipart), file)
			}
		}
	}
	return nil
}

func parseForm(req *http.Request, limit int) (int, error) {
	isForm := isRequestForm(req)
	isMultipartForm := isRequestMultipartForm(req)
	if !isForm && !isMultipartForm {
		return -1, nil
	}
	if isMultipartForm {
		if err := req.ParseMultipartForm(int64(limit) << 20); err != nil {
			return -1, err
		}
		return requestTypeMultipartForm, nil
	}
	if err := req.ParseForm(); err != nil {
		return -1, err
	}
	return requestTypeForm, nil
}

func createEmptyProcessRequestResult(err error) (url.Values, map[string][]*multipart.FileHeader, error) {
	return make(url.Values), make(map[string][]*multipart.FileHeader), err
}
