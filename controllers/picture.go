package controllers

import (
	"encoding/json"
	"strconv"
	"yxdzb/models"

	"github.com/astaxie/beego"
)

// PictureController Operations about Picture
type PictureController struct {
	beego.Controller
}

// Post comment
// @Title CreatePicture
// @Description create Picture
// @Param	body		body 	models.Picture	true		"body for Picture content"
// @Success 200 {int} models.Picture.Id
// @Failure 403 body is empty
// @router / [post]
func (p *PictureController) Post() {
	var picture models.Picture
	json.Unmarshal(p.Ctx.Input.RequestBody, &picture)
	picID, _ := models.AddPicture(picture)
	p.Data["json"] = map[string]int64{"picID": picID}
	p.ServeJSON()
}

// @Title GetAll
// @Description get all Pictures
// @Success 200 {object} models.Picture
// @router / [get]
func (p *PictureController) GetAll() {
	response := make(map[string]interface{})
	pictures, count := models.GetAllPictures()
	response["data"] = pictures
	response["count"] = count
	p.Data["json"] = response
	p.ServeJSON()
}

// @Title Get
// @Description get picture by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Picture
// @Failure 403 :id is empty
// @router /:id [get]
func (p *PictureController) Get() {
	pid := p.GetString(":id")
	intPid, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		p.Data["json"] = err.Error()
	}
	if pid != "" {
		picture, err := models.GetPicture(intPid)
		if err != nil {
			p.Data["json"] = err.Error()
		} else {
			p.Data["json"] = picture
		}
	}
	p.ServeJSON()
}
