package models

import (
	_ "errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "html"
	_ "strings"
	_ "time"
)

type Resource struct {
	ImageId        string    `gorm:"ize:255;not null" json:"image_id"`
	InstanceId  string    `gorm:"size:255;not null;unique" json:"instance_id"`
	InstanceType     string    `gorm:"size:255;not null;" json:"instance_type"`
	RootDeviceType  string    `gorm:"size:255;not null;" json:"root_device_type"`
	RootDeviceName  string    `gorm:"size:255;not null;" json:"root_device_name"`
	PrivateDnsName  string    `gorm:"size:255;not null;" json:"private_dns_name"`
	UserID        uint32    `gorm:"primary_key;auto_increment" json:"user_id"`
}

func (r *Resource) SaveResource(db *gorm.DB) (*Resource, error) {
	var err error
	err = db.Debug().Create(&r).Error
	if err != nil {
		return &Resource{}, err
	}
	return r, nil
}


func (p *Resource) FindAllResources(db *gorm.DB, userId uint32) (*[]Resource, error) {
	var err error
	var resources []Resource
	err = db.Debug().Model(&Resource{}).Where("user_id = ?", userId).Limit(100).Find(&resources).Error
	if err != nil {
		return &[]Resource{}, err
	}
	if len(resources) > 0 {
		for i, _ := range resources {
			err := db.Debug().Model(&User{}).Where("instanceId = ?", resources[i].InstanceId).Error
			if err != nil {
				return &[]Resource{}, err
			}
		}
	}
	return &resources, nil
}