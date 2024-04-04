/*
accent-auth

Testing TokenAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package auth

import (
	"context"
	openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_auth_TokenAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TokenAPIService CheckTokenContext", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		httpRes, err := apiClient.TokenAPI.CheckTokenContext(context.Background(), token).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TokenAPIService CheckTokenScopes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		resp, httpRes, err := apiClient.TokenAPI.CheckTokenScopes(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TokenAPIService CreateToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TokenAPI.CreateToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TokenAPIService DeleteRefreshTokens", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuidOrMe string
		var clientId string

		httpRes, err := apiClient.TokenAPI.DeleteRefreshTokens(context.Background(), userUuidOrMe, clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TokenAPIService GetToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		resp, httpRes, err := apiClient.TokenAPI.GetToken(context.Background(), token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TokenAPIService GetTokens", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TokenAPI.GetTokens(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TokenAPIService GetUserTokens", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuidOrMe string

		resp, httpRes, err := apiClient.TokenAPI.GetUserTokens(context.Background(), userUuidOrMe).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TokenAPIService RevokeToken", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var token string

		httpRes, err := apiClient.TokenAPI.RevokeToken(context.Background(), token).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}