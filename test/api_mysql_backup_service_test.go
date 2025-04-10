/*
API Управляемых сервисов

Testing MysqlBackupServiceApiService

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

func Test_begetOpenapiCloud_MysqlBackupServiceApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MysqlBackupServiceApiService MysqlBackupServiceGetList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MysqlBackupServiceApi.MysqlBackupServiceGetList(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MysqlBackupServiceApiService MysqlBackupServiceGetOrders", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MysqlBackupServiceApi.MysqlBackupServiceGetOrders(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MysqlBackupServiceApiService MysqlBackupServiceRestore", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var copyId string

        resp, httpRes, err := apiClient.MysqlBackupServiceApi.MysqlBackupServiceRestore(context.Background(), copyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
