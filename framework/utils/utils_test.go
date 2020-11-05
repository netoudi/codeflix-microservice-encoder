package utils_test

import (
    "encoder/framework/utils"
    "github.com/stretchr/testify/require"
    "testing"
)

func TestIsJson(t *testing.T) {
    json := `{
        "id": "814578f6-8213-4eae-806c-701734cc5427",
        "file_path": "video.mp4",
        "status": "pending"
    }`

    err := utils.IsJson(json)
    require.Nil(t, err)

    json = `json invalid`

    err = utils.IsJson(json)
    require.Error(t, err)
}
