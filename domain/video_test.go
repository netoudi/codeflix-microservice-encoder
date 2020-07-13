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
    video.CreateAt = time.Now()
    video.UpdateAt = time.Now()

    err := video.Validate()

    require.Nil(t, err)
}
