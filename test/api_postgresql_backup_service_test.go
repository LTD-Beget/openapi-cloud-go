/*
API Управляемых сервисов

Testing PostgresqlBackupServiceApiService

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

func Test_begetOpenapiCloud_PostgresqlBackupServiceApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test PostgresqlBackupServiceApiService PostgresqlBackupServiceGetList", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PostgresqlBackupServiceApi.PostgresqlBackupServiceGetList(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PostgresqlBackupServiceApiService PostgresqlBackupServiceGetOrders", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PostgresqlBackupServiceApi.PostgresqlBackupServiceGetOrders(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PostgresqlBackupServiceApiService PostgresqlBackupServiceRestore", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var copyId string

        resp, httpRes, err := apiClient.PostgresqlBackupServiceApi.PostgresqlBackupServiceRestore(context.Background(), copyId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}