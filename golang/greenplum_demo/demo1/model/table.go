package model

import (
	"fmt"
	"gorm.io/gorm"
	"greenplum_demo/conn"
)

type User struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
	Age  uint   `gorm:"type:integer" json:"age"`
}

//This migrate all tables
func Automigrate() {
	err := conn.DB.AutoMigrate(&User{})
	if err != nil {
		panic(fmt.Errorf("Automigrate table error, err: %v", err))
	}
}

func (u *User) Add (name string, age uint) {
	u.Name = name
	u.Age = age
	err := conn.DB.Model(&User{}).Create(u).Error
	if err != nil {
		panic(fmt.Errorf("create user error, err: %v", err))
	}
	fmt.Println("create user success, user: ", u)
}

func (u *User) Get () {
	err := conn.DB.Model(&User{}).First(u).Error
	if err != nil {
		panic(fmt.Errorf("get user error, err: %v", err))
	}
	fmt.Println("get user success, user: ", u)
}

func (u *User) UpdateNameByAge(name string, age uint) {
	us := []User{}
	err := conn.DB.Model(&User{}).Where("age = ?", age).Update("name", name).Find(&us).Error
	if err != nil {
		panic(fmt.Errorf("UpdateNameByAge error, err: %v", err))
	}
	fmt.Println("UpdateNameByAge success, user: ", us)
}

func (u *User) DeleteUser(age uint)  {
	err := conn.DB.Where("age=?", age).Delete(&User{}).Error
	if err != nil {
		panic(fmt.Errorf("DeleteUser error, err: %v", err))
	}
	fmt.Println("DeleteUser success")
}

func (u *User) TransactionUser()  {
	err := conn.DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&User{Name: "GiraffeTTTTT", Age: 222}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		//return errors.New("我就是故意停止")
		if err := tx.Create(&User{Name: "LionMMM", Age: 2222}).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		fmt.Errorf("TransactionUser error, err: %v", err)
	}
	fmt.Println("TransactionUser success")
}



