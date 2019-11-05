package scw

import (
	. "errors"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
)

// AsScwError returns true if the given error is a Scaleway SDK error.
func AsScwError(err error) (bool, error) {
	var e *errors.Error
	return As(err, &e), e
}

//
// Scaleway API errors.
//

// AsResponseError returns true if the given error is a response error.
func AsResponseError(err error) (bool, error) {
	var e *errors.ResponseError
	return As(err, &e), e
}

// AsInvalidArgumentsError returns true if the given error reports that a request contains an invalid argument.
func AsInvalidArgumentsError(err error) (bool, error) {
	var e *errors.InvalidArgumentsError
	return As(err, &e), e
}

// AsQuotasExceededError returns true if the given error reports that you exceeded a resource quotas.
func AsQuotasExceededError(err error) (bool, error) {
	var e *errors.QuotasExceededError
	return As(err, &e), e
}

// AsPermissionsDeniedError returns true if the given error reports that you try to make a forbidden action.
func AsPermissionsDeniedError(err error) (bool, error) {
	var e *errors.PermissionsDeniedError
	return As(err, &e), e
}

// AsResourceNotFoundError returns true if the given error is reports that the targeted resource is not found.
func AsResourceNotFoundError(err error) (bool, error) {
	var e *errors.ResourceNotFoundError
	return As(err, &e), e
}

// AsOutOfStockError returns true if the given error reports that the target product is out of stock.
func AsOutOfStockError(err error) (bool, error) {
	var e *errors.OutOfStockError
	return As(err, &e), e
}

//
// Client Validation errors.
//

// AsNoClientCredentialProvidedError returns true if the given error reports that no credential has been provided
// for a client creation.
func AsNoClientCredentialProvidedError(err error) (bool, error) {
	var e *errors.NoClientCredentialProvidedError
	return As(err, &e), e
}
