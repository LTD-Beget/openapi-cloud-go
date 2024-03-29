/*
API Управляемых сервисов

Testing PostgresqlStatisticServiceApiService

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

func Test_begetOpenapiCloud_PostgresqlStatisticServiceApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test PostgresqlStatisticServiceApiService PostgresqlStatisticServiceGetCpu", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetCpu(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PostgresqlStatisticServiceApiService PostgresqlStatisticServiceGetCpuDetails", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetCpuDetails(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PostgresqlStatisticServiceApiService PostgresqlStatisticServiceGetDisk", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetDisk(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PostgresqlStatisticServiceApiService PostgresqlStatisticServiceGetDiskUsage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetDiskUsage(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PostgresqlStatisticServiceApiService PostgresqlStatisticServiceGetLoadAverage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetLoadAverage(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PostgresqlStatisticServiceApiService PostgresqlStatisticServiceGetMemory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetMemory(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PostgresqlStatisticServiceApiService PostgresqlStatisticServiceGetNetwork", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetNetwork(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
