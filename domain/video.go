package domain

import (
    "github.com/asaskevich/govalidator"
    "time"
)

func init() {
    govalidator.SetFieldsRequiredByDefault(true)
}

type Video struct {
    ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"`
    ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:varchar(255)"`
    FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`
    CreatedAt  time.Time `json:"created_at" valid:"-"`
    UpdatedAt  time.Time `json:"updated_at" valid:"-"`
    Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
}

func NewVideo() *Video {
    return &Video{}
}

func (v *Video) Validate() error {
    _, err := govalidator.ValidateStruct(v)

    if err != nil {
        return err
    }

    return nil
}
