package models

import (
	"time"

	"cmdb/utils"

	"github.com/astaxie/beego/orm"
)

// User 用户对象
type User struct {
	ID         int        `orm:"column(id)"`
	StaffID    string     `orm:"column(staff_id);size(32)"`
	Name       string     `orm:"size(64)"`
	Nickname   string     `orm:"size(64)"`
	Password   string     `orm:"size(1024)"`
	Gender     int        `orm:""`
	Tel        string     `orm:"size(32)"`
	Addr       string     `orm:"size(128)"`
	Email      string     `orm:"size(64)"`
	Department string     `orm:"size(128)"`
	Status     int        `orm:""`
	CreatedAt  *time.Time `orm:"auto_now_add"`
	UpdatedAt  *time.Time `orm:"auto_now"`
	DeletedAt  *time.Time `orm:"null"`
}

// ValidPassword 验证用户密码是否正确
func (u *User) ValidPassword(password string) bool {
	return utils.CheckPassword(password, u.Password)
}

// GenderText 性别显示
func (u *User) GenderText() string {
	if u.Gender == 1 {
		return "女"
	}
	return "男"
}

// StatusText 状态显示
func (u *User) StatusText() string {
	switch u.Status {
	case 0:
		return "正常"
	case 1:
		return "锁定"
	case 2:
		return "离职"
	}
	return "未知"
}

// // GetUserByName 通过用户名获取用户
// func GetUserByName(name string) *User {
// 	user := &User{Name: name}
// 	ormer := orm.NewOrm()
// 	if err := ormer.Read(user, "Name"); err == nil {
// 		return user
// 	}
// 	return nil
// }

// // QueryUser 查询用户
// func QueryUser(q string) []*User {
// 	var users []*User
// 	queryset := orm.NewOrm().QueryTable(&User{})
// 	if q != "" {
// 		cond := orm.NewCondition()
// 		cond = cond.Or("name__icontains", q)
// 		cond = cond.Or("nickname__icontains", q)
// 		cond = cond.Or("tel__icontains", q)
// 		cond = cond.Or("addr__icontains", q)
// 		cond = cond.Or("email__icontains", q)
// 		cond = cond.Or("department__icontains", q)
// 		queryset = queryset.SetCond(cond)
// 	}
// 	queryset.All(&users)
// 	return users
// }

func init() {
	orm.RegisterModel(new(User))
}
