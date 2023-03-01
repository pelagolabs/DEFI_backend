package db

type Manager struct {
	CommonModel

	Username    string       `gorm:"index:idx_manager_username;size:100;not null" json:"username"`
	Password    string       `gorm:"size:100;not null" json:"-"`
	Nickname    string       `gorm:"size:100;not null" json:"nickname"`
	Email       string       `gorm:"size:100;not null" json:"email"`
	Status      string       `gorm:"index:idx_manager_status;size:100;not null" json:"status"`
	Permissions []Permission `gorm:"many2many:manager_permissions;" json:"permissions"`
}

func FindManagerById(id uint) (manager *Manager, err error) {
	err = db.Where("id = ?", id).Find(&manager).Error
	return
}

func FindManagerByUserName(userName string) (manager *Manager, err error) {
	err = db.Where("username = ?", userName).Find(&manager).Error
	return
}

func FindManagerWithPermissionByUserName(userName string) (manager *Manager, err error) {
	err = db.Preload("Permissions").Where("username = ?", userName).Find(&manager).Error
	return
}

func GetManagerCountByStatus(status string) uint {
	var count int64

	query := db.Table("managers")

	if status != "all" {
		query.Where("status = ?", status)
	}

	query.Count(&count)

	return uint(count)
}

func FindPaginationManagerWithPermissionByStatus(status string, offset int, size int) (managers []Manager, err error) {
	query := db.Table("managers")

	if status != "all" {
		query.Where("status = ?", status)
	}

	err = query.Preload("Permissions").Limit(size).Offset(offset).Find(&managers).Error
	return
}

func ReplaceManagerPermission(manager *Manager, permissionIds []uint) (err error) {
	newPermissionSet := make([]Permission, 0)
	for _, permissionId := range permissionIds {
		newPermissionSet = append(newPermissionSet, Permission{
			CommonModel: CommonModel{
				ID: permissionId,
			},
		})
	}

	err = db.Model(manager).Association("Permissions").Replace(newPermissionSet)
	return
}

func DeleteManagerById(id uint) (err error) {
	return db.Delete(&Manager{}, id).Error
}

func SaveManager(manager *Manager) (err error) {
	return db.Save(manager).Error
}
