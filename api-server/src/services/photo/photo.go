package photoSrv

import (
	"models"
)

type SrvPhoto struct {
	modelsPhotoList []models.Photo
}

func (p *SrvPhoto) GetModelsPhotoList() []models.Photo {
	return p.modelsPhotoList
}

func (p *SrvPhoto) GetSrvPhotoByRelId(rel_id uint32) error {
	qs := models.Orm.QueryTable("photo").Filter("rel_id", rel_id)
	count, err := qs.Count()
	if err != nil {
		return err
	}
	if count <= 0 {
		return nil
	}

	_, err = qs.All(&p.modelsPhotoList)
	return err
}
