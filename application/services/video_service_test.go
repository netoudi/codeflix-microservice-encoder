package services_test

import (
    "encoder/application/repositories"
    "encoder/application/services"
    "encoder/domain"
    "encoder/framework/database"
    "github.com/joho/godotenv"
    uuid "github.com/satori/go.uuid"
    "github.com/stretchr/testify/require"
    "log"
    "testing"
    "time"
)

func init() {
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func prepare() (*domain.Video, repositories.VideoRepository) {
    db := database.NewDbTest()
    defer db.Close()

    video := domain.NewVideo()
    video.ID = uuid.NewV4().String()
    video.FilePath = "video.mp4"
    video.CreatedAt = time.Now()
    video.UpdatedAt = time.Now()

    repo := repositories.VideoRepositoryDb{Db: db}

    return video, repo
}

func TestVideoServiceDownload(t *testing.T) {
    video, repo := prepare()

    videoService := services.NewVideoService()
    videoService.Video = video
    videoService.VideoRepository = repo

    err := videoService.Download("codeflix-microservice-videos")
    require.Nil(t, err)

    err = videoService.Fragment()
    require.Nil(t, err)
}
