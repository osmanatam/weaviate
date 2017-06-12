package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// WeaviateLocationsPatchURL generates an URL for the weaviate locations patch operation
type WeaviateLocationsPatchURL struct {
	LocationID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *WeaviateLocationsPatchURL) WithBasePath(bp string) *WeaviateLocationsPatchURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *WeaviateLocationsPatchURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *WeaviateLocationsPatchURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/locations/{locationId}"

	locationID := o.LocationID
	if locationID != "" {
		_path = strings.Replace(_path, "{locationId}", locationID, -1)
	} else {
		return nil, errors.New("LocationID is required on WeaviateLocationsPatchURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/weaviate/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *WeaviateLocationsPatchURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *WeaviateLocationsPatchURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *WeaviateLocationsPatchURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on WeaviateLocationsPatchURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on WeaviateLocationsPatchURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *WeaviateLocationsPatchURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}