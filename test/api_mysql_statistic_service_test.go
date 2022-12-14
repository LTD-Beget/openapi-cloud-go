/*
API Управляемых сервисов

Testing MysqlStatisticServiceApiService

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

func Test_begetOpenapiCloud_MysqlStatisticServiceApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MysqlStatisticServiceApiService MysqlStatisticServiceGetCpu", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetCpu(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MysqlStatisticServiceApiService MysqlStatisticServiceGetCpuDetails", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetCpuDetails(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MysqlStatisticServiceApiService MysqlStatisticServiceGetDisk", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetDisk(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MysqlStatisticServiceApiService MysqlStatisticServiceGetDiskUsage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetDiskUsage(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MysqlStatisticServiceApiService MysqlStatisticServiceGetLoadAverage", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetLoadAverage(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MysqlStatisticServiceApiService MysqlStatisticServiceGetMemory", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetMemory(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MysqlStatisticServiceApiService MysqlStatisticServiceGetNetwork", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var serviceId string

        resp, httpRes, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetNetwork(context.Background(), serviceId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
