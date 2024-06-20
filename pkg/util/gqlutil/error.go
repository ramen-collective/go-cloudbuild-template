package gqlutil

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ErrorCode is a string representing an error code.
type ErrorCode string

const (
	// InternalServerErrorCode should be returned on internal server error.
	// Equivalent to http 50x codes.
	InternalServerErrorCode ErrorCode = "INTERNAL_SERVER_ERROR"
	// UnauthenticatedCode should be returned when an unauthenticated user
	// is trying to access a auth only resource.
	// Equivalent to http 401 code.
	UnauthenticatedCode ErrorCode = "UNAUTHENTICATED"
	// ForbiddenCode should be returned when an authenticated user
	// does not have access to a resource.
	// Equivalent to http 403 code.
	ForbiddenCode ErrorCode = "FORBIDDEN"
	// BadUserInputCode should be returned on input data validation
	// failure.
	// Equivalent to 400 http code.
	BadUserInputCode ErrorCode = "BAD_USER_INPUT"
	// NotFoundCode should when a requested resource is not found.
	// Equivalent to 404 http code.
	NotFoundCode ErrorCode = "NOT_FOUND"
	// GiftCodeNotFound should be returned when the code doesn’t exist.
	// Equivalent to 404 http code.
	GiftCodeNotFound ErrorCode = "GIFT_CODE_NOT_FOUND"
	// GiftCodeAlreadyRedeemed should be returned when the code
	// has already been redeemed
	GiftCodeAlreadyRedeemed ErrorCode = "GIFT_CODE_ALREADY_REDEEMED"
	// GiftUserHasSubscription should be returned when
	// the user already has an active subscription
	GiftUserHasSubscription ErrorCode = "GIFT_USER_HAS_SUBSCRIPTION"
	// GiftCampaignExpired should be returned when the campaign
	// expiration date is in the past or if the campaign is canceled
	GiftCampaignExpired ErrorCode = "GIFT_CAMPAIGN_EXPIRED"

	// gqlgen error codes
	// https://github.com/99designs/gqlgen/blob/master/graphql/errcode/codes.go
	// No constructor for these errors as they should be only created
	// by gqlgen.

	// GQLValidationFailed should be returned when there's a query
	// validation error.
	GQLValidationFailed ErrorCode = errcode.ValidationFailed
	// GQLParseFailed should be returned when there's a query
	// parsing error.
	GQLParseFailed ErrorCode = errcode.ParseFailed
)

const (
	internalServerErrorString = "internal server error"
)

// ErrorPresenter ensure that an error is of type gqlerror.Error
// and has the gql path and error code filled before serialization.
// It also removes the error message for internal server errors
// to prevent technical insight leak.
func ErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
	error := graphql.ErrorOnPath(ctx, err).(*gqlerror.Error)
	if error.Extensions == nil {
		error.Extensions = map[string]interface{}{}
	}
	_, hasCode := error.Extensions["code"]
	if !hasCode {
		error.Extensions = map[string]interface{}{
			"code": InternalServerErrorCode,
		}
	}
	// Prevent low level error message from being
	// returned. You should log the original message instead.
	if error.Extensions["code"] == InternalServerErrorCode {
		error.Message = internalServerErrorString
	}
	return error
}

// GetErrorCode returns the error code. InternalServerErrorCode by default.
func GetErrorCode(err error) ErrorCode {
	var error *gqlerror.Error
	var code ErrorCode
	if errors.As(err, &error) {
		switch e := error.Extensions["code"].(type) {
		case ErrorCode:
			code = e
		case string:
			code = ErrorCode(e)
		}
	}
	if code == "" {
		code = InternalServerErrorCode
	}
	return code
}

// UnauthenticatedError returns a new unauthenticated gql error
// from a given error.
func UnauthenticatedError(err error) *gqlerror.Error {
	return Error(UnauthenticatedCode, err)
}

// UnauthenticatedErrorf returns a new unauthenticated gql error
// from a given message.
func UnauthenticatedErrorf(message string, args ...interface{}) *gqlerror.Error {
	return Errorf(UnauthenticatedCode, message, args...)
}

// ForbiddenError returns a new forbidden gql error
// from a given error.
func ForbiddenError(err error) *gqlerror.Error {
	return Error(ForbiddenCode, err)
}

// ForbiddenErrorf returns a new forbidden gql error
// from a given message.
func ForbiddenErrorf(message string, args ...interface{}) *gqlerror.Error {
	return Errorf(ForbiddenCode, message, args...)
}

// BadUserInputError returns a new bad user input gql error
// from a given error.
func BadUserInputError(err error) *gqlerror.Error {
	return Error(BadUserInputCode, err)
}

// BadUserInputErrorf returns a new bad user input gql error
// from a given message.
func BadUserInputErrorf(message string, args ...interface{}) *gqlerror.Error {
	return Errorf(BadUserInputCode, message, args...)
}

// NotFoundError returns a new not found gql error
// from a given error.
func NotFoundError(err error) *gqlerror.Error {
	return Error(NotFoundCode, err)
}

// NotFoundErrorf returns a new not found gql error
// from a given message.
func NotFoundErrorf(message string, args ...interface{}) *gqlerror.Error {
	return Errorf(NotFoundCode, message, args...)
}

// GiftCodeNotFoundError returns a new not found gql error
// from a given error.
func GiftCodeNotFoundError(err error) *gqlerror.Error {
	return Error(GiftCodeNotFound, err)
}

// GiftCodeNotFoundErrorf returns a new not found gql error
// from a given message.
func GiftCodeNotFoundErrorf(message string, args ...interface{}) *gqlerror.Error {
	return Errorf(GiftCodeNotFound, message, args...)
}

// GiftCodeAlreadyRedeemedError returns a new not found gql error
// from a given error.
func GiftCodeAlreadyRedeemedError(err error) *gqlerror.Error {
	return Error(GiftCodeAlreadyRedeemed, err)
}

// GiftCodeAlreadyRedeemedErrorf returns a new not found gql error
// from a given message.
func GiftCodeAlreadyRedeemedErrorf(message string, args ...interface{}) *gqlerror.Error {
	return Errorf(GiftCodeAlreadyRedeemed, message, args...)
}

// GiftUserHasSubscriptionError returns a new not found gql error
// from a given error.
func GiftUserHasSubscriptionError(err error) *gqlerror.Error {
	return Error(GiftUserHasSubscription, err)
}

// GiftUserHasSubscriptionErrorf returns a new not found gql error
// from a given message.
func GiftUserHasSubscriptionErrorf(message string, args ...interface{}) *gqlerror.Error {
	return Errorf(GiftUserHasSubscription, message, args...)
}

// GiftCampaignExpiredError returns a new not found gql error
// from a given error.
func GiftCampaignExpiredError(err error) *gqlerror.Error {
	return Error(GiftCampaignExpired, err)
}

// GiftCampaignExpiredErrorf returns a new not found gql error
// from a given message.
func GiftCampaignExpiredErrorf(message string, args ...interface{}) *gqlerror.Error {
	return Errorf(GiftCampaignExpired, message, args...)
}

// InternalServerError returns a new internal server gql error
// from a given error.
func InternalServerError(err error) *gqlerror.Error {
	return Error(InternalServerErrorCode, err)
}

// InternalServerErrorf returns a new internal server gql error
// from a given error.
func InternalServerErrorf(message string, args ...interface{}) *gqlerror.Error {
	return Errorf(InternalServerErrorCode, message, args...)
}

// Errorf returns a new gql error from a given code and message.
func Errorf(code ErrorCode, message string, args ...interface{}) *gqlerror.Error {
	return &gqlerror.Error{
		Message: fmt.Sprintf(message, args...),
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}

// Error returns a new gql error from a given code and error.
func Error(code ErrorCode, err error) *gqlerror.Error {
	return &gqlerror.Error{
		Message: err.Error(),
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}
