package scw

import (
	. "errors"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
)

// IsScwError returns true if the given error is a Scaleway SDK error.
func IsScwError(err error) bool {
	var e *errors.Error
	return As(err, &e)
}

//
// Scaleway API errors.
//

// ResponseError returns true if the given error is a response error.
func IsResponseError(err error) bool {
	var e *errors.ResponseError
	return As(err, &e)
}

// IsInvalidArgumentsError returns true if the given error reports that a request contains an invalid argument.
func IsInvalidArgumentsError(err error) bool {
	var e *errors.InvalidArgumentsError
	return As(err, &e)
}

// IsQuotasExceededError returns true if the given error reports that you exceeded a resource quotas.
func IsQuotasExceededError(err error) bool {
	var e *errors.QuotasExceededError
	return As(err, &e)
}

// IsPermissionsDeniedError returns true if the given error reports that you try to make a forbidden action.
func IsPermissionsDeniedError(err error) bool {
	var e *errors.PermissionsDeniedError
	return As(err, &e)
}

// IsResourceNotFoundError returns true if the given error is reports that the targeted resource is not found.
func IsResourceNotFoundError(err error) bool {
	var e *errors.ResourceNotFoundError
	return As(err, &e)
}

// IsOutOfStockError returns true if the given error reports that the target product is out of stock.
func IsOutOfStockError(err error) bool {
	var e *errors.OutOfStockError
	return As(err, &e)
}

//
// Client Validation errors.
//

// IsNoClientCredentialProvidedError returns true if the given error reports that no credential has been provided
// for a client creation.
func IsNoClientCredentialProvidedError(err error) bool {
	var e *errors.NoClientCredentialProvidedError
	return As(err, &e)
}
