/*
Hava

Testing TeamsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_openapi_TeamsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test TeamsApiService TeamMembersAdd", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var accountId string
        var teamId string

        resp, httpRes, err := apiClient.TeamsApi.TeamMembersAdd(context.Background(), accountId, teamId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TeamsApiService TeamMembersInvite", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var accountId string
        var teamId string

        resp, httpRes, err := apiClient.TeamsApi.TeamMembersInvite(context.Background(), accountId, teamId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TeamsApiService TeamMembersRemove", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var accountId string
        var teamId string
        var userId string

        resp, httpRes, err := apiClient.TeamsApi.TeamMembersRemove(context.Background(), accountId, teamId, userId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TeamsApiService TeamsCreate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var accountId string

        resp, httpRes, err := apiClient.TeamsApi.TeamsCreate(context.Background(), accountId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TeamsApiService TeamsDestroy", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var accountId string
        var teamId string

        resp, httpRes, err := apiClient.TeamsApi.TeamsDestroy(context.Background(), accountId, teamId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TeamsApiService TeamsIndex", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var accountId string

        resp, httpRes, err := apiClient.TeamsApi.TeamsIndex(context.Background(), accountId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TeamsApiService TeamsShow", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var accountId string
        var teamId string

        resp, httpRes, err := apiClient.TeamsApi.TeamsShow(context.Background(), accountId, teamId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test TeamsApiService TeamsUpdate", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var accountId string
        var teamId string

        resp, httpRes, err := apiClient.TeamsApi.TeamsUpdate(context.Background(), accountId, teamId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
