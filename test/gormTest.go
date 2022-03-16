package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //代码部直接使用包 底层链接使用
)

type Student struct {
	Id   int
	Name string `gorm:"size:50;default:"xiaoming"`
	Age  int
	// Class int `gorm:"not null`
	// Join time.Time `gorm:"type:timestamp`
}

var GlobalDb *gorm.DB

func main() {

	//链接数据库
	db, err := gorm.Open("mysql", "root:zhang123@(127.0.0.1:3306)/test?parseTime=True&loc=Local")
	if err != nil {
		return
	}
	// defer db.Close()
	GlobalDb = db
	GlobalDb.DB().SetMaxIdleConns(10)
	GlobalDb.DB().SetMaxOpenConns(100)

	// // 创建数据库表--不能使用gorm创建数据库，提前使用sql语句，创建好想要的数据库。
	// fmt.Println(GlobalDb.AutoMigrate(new(Student2)).Error) //默认创建的表为复数类型，自动添加“s”
	// //创建之前添加db.SingularTable(true)可以创建非复数表名

	//插入数据
	// InsertData()
	//查询数据
	// Search()
	//更新数据
	// UpdateData()

	//删除数据
	DeleteData()

	// fmt.Printf("%#v\n", u)
	// var uu UserInfo
	// db.Find(&uu, "hobby=?", "足球")
	// fmt.Printf("%#v\n", uu)

	// // 删除
	// db.Delete(&u)
}

func InsertData() {
	var stu Student
	stu.Name = "zhangsan"
	stu.Age = 100
	// stu := Student{1, "zhangsan", 100}

	//插入创建数据
	fmt.Println(GlobalDb.Create(&stu).Error)
}

func Search() {
	// var stu Student
	// // GlobalDb.First(&stu)
	// GlobalDb.Select("name, age").First(&stu)

	var stu []Student
	// GlobalDb.Select("name, age").Find(&stu)
	// GlobalDb.Select("name, age").Where("name=?", "lisi").Find(&stu)
	// GlobalDb.Select("name, age").Where("name=?", "lisi").Where("age=?", 17).Find(&stu)
	GlobalDb.Select("name, age").Where("name=? and age=?", "lisi", 17).Find(&stu)

	//查询数据
	fmt.Println(stu)
}

func UpdateData() {
	//更新数据
	// fmt.Println(GlobalDb.Model(new(Student)).Where("name=?", "lisi").Update("name", "zhaoliu").Error)
	fmt.Println(GlobalDb.Model(new(Student)).Where("name=?", "zhaoliu").
		Updates(map[string]interface{}{"name": "liuqi", "age": 77}).Error)
}

type Student2 struct {
	gorm.Model //go语言中，匿名成员---继承
	Name       string
	Age        int
}

func DeleteData() {
	//删除数据
	fmt.Println(GlobalDb.Where("name=?", "zhangsan").Delete(new(Student2)).Error)
	// fmt.Println(GlobalDb.Unscoped().Where("name=?", "zhangsan").Delete(new(Student2)).Error)

	// var stu Student2
	// stu.Name = "zhangsan"
	// stu.Age = 100
	// fmt.Println(GlobalDb.Create(&stu).Error)

	// var stu []Student2
	// GlobalDb.Unscoped().Find(&stu)
	// fmt.Println(stu)
}
