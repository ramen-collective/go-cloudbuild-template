package gqlutil

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func TestErrorPresenter(t *testing.T) {
	require := require.New(t)
	testCases := []struct {
		name        string
		err         error
		expectedErr *gqlerror.Error
	}{
		{
			name: "Basic error",
			err:  errors.New("Basic error message"),
			expectedErr: &gqlerror.Error{
				Message: internalServerErrorString,
				Extensions: map[string]interface{}{
					"code": InternalServerErrorCode,
				},
			},
		},
		{
			name: "Gql error with no code",
			err: &gqlerror.Error{
				Message: "Basic error message",
			},
			expectedErr: &gqlerror.Error{
				Message: internalServerErrorString,
				Extensions: map[string]interface{}{
					"code": InternalServerErrorCode,
				},
			},
		},
		{
			name: "Gql error with internal error code",
			err: &gqlerror.Error{
				Message: "Basic error message",
				Extensions: map[string]interface{}{
					"code": InternalServerErrorCode,
				},
			},
			expectedErr: &gqlerror.Error{
				Message: internalServerErrorString,
				Extensions: map[string]interface{}{
					"code": InternalServerErrorCode,
				},
			},
		},
		{
			name: "Gql error with another error code",
			err: &gqlerror.Error{
				Message: "Basic error message",
				Extensions: map[string]interface{}{
					"code": UnauthenticatedCode,
				},
			},
			expectedErr: &gqlerror.Error{
				Message: "Basic error message",
				Extensions: map[string]interface{}{
					"code": UnauthenticatedCode,
				},
			},
		},
	}

	ctx := context.Background()
	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			err := ErrorPresenter(ctx, tc.err)
			require.Equal(tc.expectedErr.Message, err.Message)
			require.Equal(tc.expectedErr.Extensions["code"], err.Extensions["code"])
		})
	}
}

func TestGetErrorCode(t *testing.T) {
	require := require.New(t)
	testCases := []struct {
		name         string
		err          error
		expectedCode ErrorCode
	}{
		{name: "Basic error", err: errors.New("Basic error"), expectedCode: InternalServerErrorCode},
		{name: "Gql error with no code", err: &gqlerror.Error{Message: "Gql error"}, expectedCode: InternalServerErrorCode},
		{
			name: "Gql error with code",
			err: &gqlerror.Error{
				Message: "Gql error",
				Extensions: map[string]interface{}{
					"code": NotFoundCode,
				},
			},
			expectedCode: NotFoundCode},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			code := GetErrorCode(tc.err)
			require.Equal(tc.expectedCode, code)
		})
	}
}

func TestErrorCreation(t *testing.T) {
	require := require.New(t)
	tc := struct {
		err                error
		message            string
		arg                string
		internalErr        *gqlerror.Error
		unauthenticatedErr *gqlerror.Error
		forbiddenErr       *gqlerror.Error
		badUserInputErr    *gqlerror.Error
		notFoundErr        *gqlerror.Error
	}{
		err:     errors.New("error: test"),
		message: "error: %s",
		arg:     "test",
		internalErr: &gqlerror.Error{
			Message: "error: test",
			Extensions: map[string]interface{}{
				"code": InternalServerErrorCode,
			},
		},
		unauthenticatedErr: &gqlerror.Error{
			Message: "error: test",
			Extensions: map[string]interface{}{
				"code": UnauthenticatedCode,
			},
		},
		forbiddenErr: &gqlerror.Error{
			Message: "error: test",
			Extensions: map[string]interface{}{
				"code": ForbiddenCode,
			},
		},
		badUserInputErr: &gqlerror.Error{
			Message: "error: test",
			Extensions: map[string]interface{}{
				"code": BadUserInputCode,
			},
		},
		notFoundErr: &gqlerror.Error{
			Message: "error: test",
			Extensions: map[string]interface{}{
				"code": NotFoundCode,
			},
		},
	}

	error := InternalServerError(tc.err)
	require.Equal(tc.internalErr, error)
	error = InternalServerErrorf(tc.message, tc.arg)
	require.Equal(tc.internalErr, error)
	error = UnauthenticatedError(tc.err)
	require.Equal(tc.unauthenticatedErr, error)
	error = UnauthenticatedErrorf(tc.message, tc.arg)
	require.Equal(tc.unauthenticatedErr, error)
	error = ForbiddenError(tc.err)
	require.Equal(tc.forbiddenErr, error)
	error = ForbiddenErrorf(tc.message, tc.arg)
	require.Equal(tc.forbiddenErr, error)
	error = BadUserInputError(tc.err)
	require.Equal(tc.badUserInputErr, error)
	error = BadUserInputErrorf(tc.message, tc.arg)
	require.Equal(tc.badUserInputErr, error)
	error = NotFoundError(tc.err)
	require.Equal(tc.notFoundErr, error)
	error = NotFoundErrorf(tc.message, tc.arg)
	require.Equal(tc.notFoundErr, error)
}
