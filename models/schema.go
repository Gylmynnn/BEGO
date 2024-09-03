package models

import (
  "time"
)

type Character struct {
  Id int `json:"id" gorm:"primary_key"`
  Character string `json:"character"`
  Quote string `json:"quote"`
  CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
  UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
