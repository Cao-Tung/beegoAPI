package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // import your required driver

	"errors"
)


type Student struct {
	Id string   `orm:"pk"`
	Name string
}

func init() {
	orm.RegisterModel(new(Student))
	orm.RegisterDataBase("default", "mysql", "root:@/beego?charset=utf8", 30)
}

func Add(sv Student) (id string){
	isv := orm.NewOrm()
	_, err := isv.Insert(&sv)
	if err !=nil{
		return err.Error()
	}
	return sv.Id
}

func Dele(id string) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Student{Id: id}); err == nil {
		fmt.Println(num)
	}
}

func GetAl() (student[]Student){
	o := orm.NewOrm()
	var students []Student
	o.QueryTable("student").All(&students)
	return students
}

func GetOn(StudentId string) (studen Student) {
	o := orm.NewOrm()
	var student Student
	o.QueryTable("student").Filter("Id", StudentId).One(&student)
	return student
}

func Search(studentName string)(student[] Student){
	o := orm.NewOrm()
	qs := o.QueryTable("student")
	var students []Student
	qs.Filter("name__icontains", studentName).All(&students)
	return students
}

func Updat(studentId string, name string) (err error) {
	o := orm.NewOrm()
	student := Student{Id: studentId}
	if o.Read(&student) == nil {
		student.Name = name
		o.Update(&student)
		return nil
	}
	return errors.New("ObjectId Not Exist")
}


