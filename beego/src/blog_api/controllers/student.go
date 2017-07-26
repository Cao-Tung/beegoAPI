package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"blog_api/models"
)

// Operations about student
type StudentController struct {
	beego.Controller
}

// @Title Create
// @Description create Student
// @Param	body		body 	models.Student	true		"The student content"
// @Success 200 {string} models.Student.Id
// @Failure 403 body is empty
// @router / [post]
func (o *StudentController) Post() {
	var ob models.Student
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	studentId := models.Add(ob)
	o.Data["json"] = map[string]string{"StudentId": studentId}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the student
// @Param	studentId	path 	string	true		"The studentId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 studentId is empty
// @router /:studentId [delete]
func (o *StudentController) Delete() {
	studentId := o.Ctx.Input.Param(":studentId")
	models.Dele(studentId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

// @Title GetAll
// @Description get all students
// @Success 200 {object} models.Student
// @Failure 403 :studentId is empty
// @router / [get]
func (o *StudentController) GetAll() {
	obs := models.GetAl()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Get
// @Description find student by studentId
// @Param	studentId	path 	string	true		"the student you want to get"
// @Success 200 {student} models.Student
// @Failure 403 :studentId is empty
// @router /:studentId [get]
func (o *StudentController) Get() {
	studentId := o.Ctx.Input.Param(":studentId")
	if studentId != "" {
		ob := models.GetOn(studentId)

			o.Data["json"] = ob

	}
	o.ServeJSON()
}

// @Title Update
// @Description update the student
// @Param	studentId	path 	string	true		"The studentId you want to update"
// @Param	body		body 	models.Student	true		"The body"
// @Success 200 {object} models.Student
// @Failure 403 :studentId is empty
// @router /:studentId [put]
func (o *StudentController) Put() {
	studentId := o.Ctx.Input.Param(":studentId")
	var ob models.Student
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Updat(studentId, ob.Name)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Get
// @Description find students by name
// @Param	name	query 	string	true		"the students you want to get"
// @Success 200 {student} models.Student
// @Failure 403 :name is empty
// @router /searchByName [get]
func (o *StudentController) Search() {
	nameStudent := o.Ctx.Input.Query("name")
	if nameStudent != "" {
		obs := models.Search(nameStudent)

		o.Data["json"] = obs

	}
	o.ServeJSON()
}


