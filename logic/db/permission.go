package db

type Permission struct {
	CommonModel

	Name     string `gorm:"size:100;not null" json:"name"`
	Describe string `gorm:"size:200" json:"describe"`
	GroupId  uint   `gorm:"index:idx_permission_group_id;not null" json:"group_id"`
	Identity string `gorm:"size:100;not null" json:"identity"`
}

func GetAllPermissionCountByGroup(group uint) uint {
	var count int64

	query := db.Table("permissions")

	query.Where("group_id = ?", group).Count(&count)

	return uint(count)
}

func FindAllPermissionByGroup(group uint) (permissions []Permission, err error) {
	query := db.Table("permissions")

	err = query.Where("group_id = ?", group).Order("id desc").Find(&permissions).Error
	return
}

func FindAvailablePermissionByIdGroup(ids []uint, group uint) (permissions []Permission, err error) {
	err = db.Where("id IN ? AND group_id = ?", ids, group).Find(&permissions).Error
	return
}
