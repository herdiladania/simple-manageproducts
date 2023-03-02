package entity

import "github.com/labstack/echo/v4"

type Core struct {
	ID         uint
	FullName   string
	FirstOrder string
	CreatedAt  string
	UpdateAt   string
	DeleteAt   string
}

type UserHandler interface {
	Create() echo.HandlerFunc
	List() echo.HandlerFunc
	Detail() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type UserService interface {
	Create(newUser Core) (Core, error)
	List(page int) (map[string]interface{}, []Core, error)
	Detail(userID uint) (Core, error)
	Update(userID uint) (Core, error)
	Delete(userID uint) error
}

type UserData interface {
	Create(newUser Core) (Core, error)
	List(page int) (map[string]interface{}, []Core, error)
	Detail(userID uint) (Core, error)
	Update(userID uint) (Core, error)
	Delete(userID uint) error
}
