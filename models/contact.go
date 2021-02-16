package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	MobileNumber string `gorm:"size:255;null;unique" json:"mobileNumber,omitempty"`
	Country      string `gorm:"size:255;null;" json:"country,omitempty"`
	User         *User  `gorm:"foreignkey:UserID" json:"user"`
	UserID       int64  `gorm:"not null" json:"userId"`
}

func (c *Contact) Prepare() {
	c.Country = html.EscapeString(strings.TrimSpace(c.Country))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Contact) Save(db *gorm.DB) (*Contact, error) {
	err := db.Debug().Model(&Contact{}).Where("user_id=?", c.UserID).FirstOrCreate(&c).Error
	if err != nil {
		return &Contact{}, err
	}
	return c, nil
}

func (c *Contact) Find(db *gorm.DB, id int64) (*Contact, error) {
	err := db.Debug().Model(&Contact{}).Where("id=?", id).Find(&c).Error
	if err != nil {
		return &Contact{}, err
	}
	return c, nil
}

func (c *Contact) FindAll(db *gorm.DB) ([]Contact, error) {
	contacts := []Contact{}
	err := db.Debug().Model(&Contact{}).Find(&contacts).Error
	if err != nil {
		return []Contact{}, err
	}
	return contacts, nil
}
