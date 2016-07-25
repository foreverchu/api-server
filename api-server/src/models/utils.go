package models

import "github.com/chinarun/utils"

// IsValueExists 表示检查一个db里
func IsValueExists(table string, column string, value interface{}) bool {
	utils.Logger.Debug("models.IsValueExists : table: %s, columns: %s, value: %v", table, column, value)
	return Orm.QueryTable(table).Filter(column, value).Exist()
}
