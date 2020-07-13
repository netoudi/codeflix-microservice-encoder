package domain_test

import (
    "encoder/domain"
    uuid "github.com/satori/go.uuid"
    "github.com/stretchr/testify/require"
    "testing"
    "time"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
    video := domain.NewVideo()

    err := video.Validate()

    require.Error(t, err)
}

func TestVideoIdIsNotAUuid(t *testing.T) {
    video := domain.NewVideo()

    video.ID = uuid.NewV4().String()
    video.ResourceID = "a"
    video.FilePath = "path"
    video.CreatedAt = time.Now()
    video.UpdatedAt = time.Now()

    err := video.Validate()

    require.Nil(t, err)
}
