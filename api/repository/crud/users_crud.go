// Package crud defines basic crud applications on the database by a user.
package crud

import (
	"github.com/Joe-Degs/abeg_mock_backend/api/helpers"
	"github.com/Joe-Degs/abeg_mock_backend/api/models"
	"gorm.io/gorm"
)

type users_crud struct {
	db *gorm.DB
}

func NewUsersCrud(db *gorm.DB) *users_crud {
	return &users_crud{db}
}

func (u *users_crud) SaveUser(user models.User) (models.User, error) {
	var err error
	ch := make(chan bool)

	go func(done chan<- bool) {
		defer close(done)
		err = u.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			done <- false
			return
		}
		done <- true
	}(ch)

	if helpers.OK(ch) {
		return user, nil
	}
	return models.User{}, err
}

func (u *users_crud) FindUser(phoneNumber string) (models.User, error) {
	var err error
	var user models.User
	ch := make(chan bool)

	go func(done chan<- bool) {
		defer close(done)
		err = u.db.Debug().Model(&models.User{}).Where("phone_number = ?", phoneNumber).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(ch)

	if helpers.OK(ch) {
		return user, nil
	}
	return models.User{}, err
}

//func (u *users_crud) UpdateUser(phoneNumber string, user models.User) (models.User, error) {
//	var rs *gorm.DB
//	ch := make(chan bool)
//
//	go func(done chan<- bool) {
//		defer close(done)
//		// not complete
//		//rs = u.db.Debug().Model(&models.User{}).Where("phonenumber = ?", phoneNumber).Take(&m)
//	}(ch)
//
//	if helpers.OK(ch) {
//		if rs.Error != nil {
//			return 0, rs.Error
//		}
//		return rs.RowAffected, nil
//	}
//	return 0, rs.Error
//}
//
//func (u *users_crud) DeleteUser(phoneNumber string) (models.User, error) {
//	var err error
//	ch := make(chan bool)
//
//	go func(done chan<- bool) {
//		defer close(done)
//	}(ch)
//}
