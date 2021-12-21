package controllers

import "gorm.io/gorm"

type controller struct {
    DB *gorm.DB
}

func New(db *gorm.DB) controller {
    return controller{db}
}
