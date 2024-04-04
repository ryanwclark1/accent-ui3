/*
accent-auth

Testing TenantsAPIService

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

func Test_auth_TenantsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TenantsAPIService CreateTenant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TenantsAPI.CreateTenant(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TenantsAPIService DeleteTenant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantUuid string

		httpRes, err := apiClient.TenantsAPI.DeleteTenant(context.Background(), tenantUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TenantsAPIService GetTenant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantUuid string

		resp, httpRes, err := apiClient.TenantsAPI.GetTenant(context.Background(), tenantUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TenantsAPIService GetTenants", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TenantsAPI.GetTenants(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TenantsAPIService UpdateTenant", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tenantUuid string

		resp, httpRes, err := apiClient.TenantsAPI.UpdateTenant(context.Background(), tenantUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}