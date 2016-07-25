package models

import (
	"fmt"
)

type Tag struct {
	Id   uint32
	Name string
}

func GetTagIdByTagName(tagname string) (tagid uint32, err error) {
	query := Orm.QueryTable("tag").Filter("name", tagname)
	count, err := query.Count()
	if err != nil {
		return tagid, err
	}
	if count != 1 {
		return tagid, fmt.Errorf("找不到类型”%s“下的比赛 ", tagname)
	}

	var tag Tag
	err = query.One(&tag)
	if err != nil {
		return tagid, err
	}

	return tag.Id, err
}
