package utils

import "github.com/astaxie/beego/orm"

func CheckDBIfStrExist(Orm orm.Ormer, table string, field_name string, str string) (bool, error) {
	count, err := Orm.QueryTable(table).Filter(field_name, str).Count()
	if err != nil {
		Logger.Error("failed in query table select count(*) from %s where %s == %s, err: %v", table, field_name, str, err)
		return false, err
	}

	return count > 0, nil
}
