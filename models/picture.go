package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Picture))
}

type Picture struct {
	Id   int64
	Name string
	Img  string `orm:"size(2048)"`
	Desc string
}

// 添加图片
func AddPicture(p Picture) (int64, error) {
	o := orm.NewOrm()
	picture := new(Picture)
	picture.Name = p.Name
	id, err := o.Insert(picture)
	return id, err
}

// 获取所有图片
func GetAllPictures() (pictures []orm.Params, count int64) {
	o := orm.NewOrm()
	pic := new(Picture)
	qs := o.QueryTable(pic)
	qs.Limit(100, 0).Values(&pictures)
	count, _ = qs.Count()
	return pictures, count
}

// 根据id获取一张图片
func GetPicture(id int64) (picture Picture, err error) {
	picture = Picture{Id: id}
	o := orm.NewOrm()
	err = o.Read(&picture, "Id")
	return picture, err
}
