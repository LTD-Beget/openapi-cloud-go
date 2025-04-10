/*
API Управляемых сервисов

Testing S3ServiceApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package begetOpenapiCloud

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_begetOpenapiCloud_S3ServiceApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test S3ServiceApiService S3ServiceChangeAccessKey", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.S3ServiceApi.S3ServiceChangeAccessKey(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test S3ServiceApiService S3ServiceChangeCors", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.S3ServiceApi.S3ServiceChangeCors(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test S3ServiceApiService S3ServiceChangeDomain", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.S3ServiceApi.S3ServiceChangeDomain(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test S3ServiceApiService S3ServiceChangePublic", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.S3ServiceApi.S3ServiceChangePublic(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test S3ServiceApiService S3ServiceEnableFtp", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.S3ServiceApi.S3ServiceEnableFtp(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test S3ServiceApiService S3ServiceGetPrefix", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.S3ServiceApi.S3ServiceGetPrefix(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test S3ServiceApiService S3ServiceGetPrice", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.S3ServiceApi.S3ServiceGetPrice(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test S3ServiceApiService S3ServiceGetQuota", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.S3ServiceApi.S3ServiceGetQuota(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
