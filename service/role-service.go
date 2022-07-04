package service

import (
	"github.com/Lasang3012/facilityManage/entity"
	"github.com/Lasang3012/facilityManage/repository"
)

type RoleService interface {
	All() []entity.Role
}

type roleService struct {
	roleRepository repository.RoleRepository
}

//NewBookService .....
func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{
		roleRepository: roleRepo,
	}
}

func (service *roleService) All() []entity.Role {
	return service.roleRepository.AllRole()
}
