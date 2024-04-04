/*
accent-auth

Testing PoliciesAPIService

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

func Test_auth_PoliciesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PoliciesAPIService AddGroupPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var policyUuid string

		httpRes, err := apiClient.PoliciesAPI.AddGroupPolicy(context.Background(), groupUuid, policyUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService AddPolicyAccess", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyUuid string
		var access string

		httpRes, err := apiClient.PoliciesAPI.AddPolicyAccess(context.Background(), policyUuid, access).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService AddUserPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyUuid string
		var userUuid string

		httpRes, err := apiClient.PoliciesAPI.AddUserPolicy(context.Background(), policyUuid, userUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService CreatePolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PoliciesAPI.CreatePolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService DeleteGroupPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var policyUuid string

		httpRes, err := apiClient.PoliciesAPI.DeleteGroupPolicy(context.Background(), groupUuid, policyUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService DeletePolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyUuid string

		httpRes, err := apiClient.PoliciesAPI.DeletePolicy(context.Background(), policyUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService DeletePolicyAccess", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyUuid string
		var access string

		httpRes, err := apiClient.PoliciesAPI.DeletePolicyAccess(context.Background(), policyUuid, access).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService DeleteUserPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyUuid string
		var userUuid string

		httpRes, err := apiClient.PoliciesAPI.DeleteUserPolicy(context.Background(), policyUuid, userUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService EditPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyUuid string

		resp, httpRes, err := apiClient.PoliciesAPI.EditPolicy(context.Background(), policyUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService GetPolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PoliciesAPI.GetPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService GetPolicy", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var policyUuid string

		resp, httpRes, err := apiClient.PoliciesAPI.GetPolicy(context.Background(), policyUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService GetUserPolicies", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string

		resp, httpRes, err := apiClient.PoliciesAPI.GetUserPolicies(context.Background(), userUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
