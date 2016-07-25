package controllers

import (
	"err_code"
	"models"

	"github.com/bitly/go-simplejson"
	"github.com/chinarun/utils"
)

type PhotoController struct {
	BaseController
}

func (c *PhotoController) Prepare() {
	c.BaseController.Prepare()
}

func (c *PhotoController) Finish() {
	defer c.BaseController.Finish()
}

func (c *PhotoController) Post() {
	retJson := init_retJson()

	if retJson["result"] == err_code.OK {
		switch c.cmd {
		case "AddPhoto":
			c.Data["json"] = PhotoAdd(c.js)
		case "DelPhoto":
			c.Data["json"] = PhotoDel(c.js)

		default:
			retJson_edit(retJson, err_code.ErrCmd, "未知的cmd")
			c.Data["json"] = retJson
		}
	}
	c.ServeJson()
}

const (
	HP_REL_ID = "rel_id"
	HP_URL    = "url"
)

//photo 相关操作接口实现

func getPhotoParametersMap() map[string]utils.ParamInfo {
	return map[string]utils.ParamInfo{ //
		HP_REL_ID: {PARAM_REQUESTED, utils.DATA_TYPE_STRING},
		HP_TYPE:   {PARAM_REQUESTED, utils.DATA_TYPE_INT},
		HP_URL:    {PARAM_REQUESTED, utils.DATA_TYPE_ARRAY},
	}
}

func getPhotoValueFromParam(param_values map[string]interface{}) (photo []models.Photo, err error) {
	relId := uint32(utils.AddInt(param_values[HP_REL_ID]))
	photoType := uint8(utils.AddInt(param_values[HP_TYPE]))

	urls := (param_values[HP_URL]).([]interface{})

	for _, url := range urls {
		var tmpPhoto models.Photo
		tmpPhoto.RelId = relId
		tmpPhoto.Type = photoType
		tmpPhoto.Url = url.(string)
		photo = append(photo, tmpPhoto)
	}

	return
}

//插入图片
func insertPhoto(param_values map[string]interface{}, retJson map[string]interface{}) (num int64, err error) {
	photos, err := getPhotoValueFromParam(param_values)
	if len(photos) < 1 {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return
	}

	num, err = models.Orm.InsertMulti(1, &photos)
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "数据库写入Photo出错,"+err.Error())
		return
	}

	return
}

//删除图片
func deletePhoto(param_values map[string]interface{}, retJson map[string]interface{}) (num int64, err error) {
	relId := uint32(utils.AddInt(param_values[HP_REL_ID]))
	photoType := uint8(utils.AddInt(param_values[HP_TYPE]))

	urls := (param_values[HP_URL]).([]interface{})

	querySelector := models.Orm.QueryTable(models.DB_TABLE_PHOTO)
	querySelector = querySelector.Filter(HP_REL_ID, relId).Filter(HP_TYPE, photoType).Filter(HP_URL+"__in", urls)

	num, err = querySelector.Delete()
	if err != nil {
		retJson_edit(retJson, err_code.ServerErr, "数据库查询失败")
	}
	return
}

//api add photo
func PhotoAdd(js *simplejson.Json) map[string]interface{} {
	retJson := init_retJson()

	var photo_values map[string]interface{}
	var err error
	photo_values, err = utils.ParseParam(js, getPhotoParametersMap())
	utils.Logger.Debug("controllers/api_photo. AddPhoto : post data = %v,  params = %v, err = %v", js, photo_values, err)
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	num, err := insertPhoto(photo_values, retJson)
	if err != nil {
		return retJson
	}

	retJson["num"] = num

	return retJson
}

//api del photo
func PhotoDel(js *simplejson.Json) map[string]interface{} {
	retJson := init_retJson()

	var photo_values map[string]interface{}
	var err error
	photo_values, err = utils.ParseParam(js, getPhotoParametersMap())
	if err != nil {
		retJson_edit(retJson, err_code.InvalidData, err.Error())
		return retJson
	}

	num, err := deletePhoto(photo_values, retJson)

	retJson["num"] = num

	return retJson
}

// 根据图片类型和相关id查询所有图片
// 使用示例： 1.  GetPhotosByID("1", models.PHOTO_TYPE_PARTY, models.PHOTO_TYPE_PARTY_ROUTE)
//           2.
//               types:=[]int{models.PHOTO_TYPE_PARTY, models.PHOTO_TYPE_PARTY_ROUTE}
//               GetPhotosByID("1", types)
func GetPhotosByID(relId string, dbType ...interface{}) ([]*models.Photo, error) {
	query_result := models.Orm.QueryTable(models.DB_TABLE_PHOTO).Filter(models.DB_PHOTO_TYPE+"__in", dbType).Filter(models.DB_PHOTO_REL_ID, relId)

	var photos []*models.Photo
	num, err := query_result.All(&photos)
	if num < 1 {
		return nil, err
	}

	return photos, err
}
