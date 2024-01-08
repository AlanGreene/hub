// Code generated by goa v3.14.4, DO NOT EDIT.
//
// catalog HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	"context"
	"errors"
	"net/http"
	"strings"

	catalog "github.com/tektoncd/hub/api/gen/catalog"
	catalogviews "github.com/tektoncd/hub/api/gen/catalog/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeRefreshResponse returns an encoder for responses returned by the
// catalog Refresh endpoint.
func EncodeRefreshResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(*catalogviews.Job)
		enc := encoder(ctx, w)
		body := NewRefreshResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRefreshRequest returns a decoder for requests sent to the catalog
// Refresh endpoint.
func DecodeRefreshRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			catalogName string
			token       string
			err         error

			params = mux.Vars(r)
		)
		catalogName = params["catalogName"]
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewRefreshPayload(catalogName, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeRefreshError returns an encoder for errors returned by the Refresh
// catalog endpoint.
func EncodeRefreshError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "not-found":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewRefreshNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "internal-error":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewRefreshInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRefreshAllResponse returns an encoder for responses returned by the
// catalog RefreshAll endpoint.
func EncodeRefreshAllResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.([]*catalog.Job)
		enc := encoder(ctx, w)
		body := NewRefreshAllResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRefreshAllRequest returns a decoder for requests sent to the catalog
// RefreshAll endpoint.
func DecodeRefreshAllRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			token string
			err   error
		)
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewRefreshAllPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeRefreshAllError returns an encoder for errors returned by the
// RefreshAll catalog endpoint.
func EncodeRefreshAllError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "internal-error":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewRefreshAllInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCatalogErrorResponse returns an encoder for responses returned by the
// catalog CatalogError endpoint.
func EncodeCatalogErrorResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*catalog.CatalogErrorResult)
		enc := encoder(ctx, w)
		body := NewCatalogErrorResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCatalogErrorRequest returns a decoder for requests sent to the catalog
// CatalogError endpoint.
func DecodeCatalogErrorRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			catalogName string
			token       string
			err         error

			params = mux.Vars(r)
		)
		catalogName = params["catalogName"]
		token = r.Header.Get("Authorization")
		if token == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("token", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewCatalogErrorPayload(catalogName, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}

		return payload, nil
	}
}

// EncodeCatalogErrorError returns an encoder for errors returned by the
// CatalogError catalog endpoint.
func EncodeCatalogErrorError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "internal-error":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCatalogErrorInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalCatalogJobToJobResponse builds a value of type *JobResponse from a
// value of type *catalog.Job.
func marshalCatalogJobToJobResponse(v *catalog.Job) *JobResponse {
	res := &JobResponse{
		ID:          v.ID,
		CatalogName: v.CatalogName,
		Status:      v.Status,
	}

	return res
}

// marshalCatalogCatalogErrorsToCatalogErrorsResponseBody builds a value of
// type *CatalogErrorsResponseBody from a value of type *catalog.CatalogErrors.
func marshalCatalogCatalogErrorsToCatalogErrorsResponseBody(v *catalog.CatalogErrors) *CatalogErrorsResponseBody {
	res := &CatalogErrorsResponseBody{
		Type: v.Type,
	}
	if v.Errors != nil {
		res.Errors = make([]string, len(v.Errors))
		for i, val := range v.Errors {
			res.Errors[i] = val
		}
	} else {
		res.Errors = []string{}
	}

	return res
}
