package repository

import (
	"github.com/Lasang3012/facilityManage/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	AllRole() []entity.Role
}

type roleConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates an instance BookRepository
func NewRoleRepository(dbConn *gorm.DB) RoleRepository {
	return &roleConnection{
		connection: dbConn,
	}
}

func (db *roleConnection) AllRole() []entity.Role {
	var roles []entity.Role
	db.connection.Preload("Role").Find(&roles)
	return roles
}
