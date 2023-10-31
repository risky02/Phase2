package controller

import (
	"net/http"
	"ugdc11/dto"
	"ugdc11/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) UserDB {
	return UserDB{DB: db}
}

func (user UserDB) Register(c echo.Context) error {
	reqBody := dto.UserRegister{}
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Request Body")
	}

	var userName entity.User
	result := user.DB.Where("username = ?", reqBody.Username).First(&userName)
	if result.RowsAffected > 0 {
		return c.JSON(http.StatusBadRequest, "Username sudah terpakai")
	}

	newUser := entity.User{
		Username: reqBody.Username,
		Password: reqBody.Password,
		Deposit_Amount: reqBody.Deposit_Amount,
	}

	err := user.DB.Create(&newUser).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal menyimpan ke database")
	} else {
		return c.JSON(http.StatusCreated, "Registrasi Sukses")
	}
}