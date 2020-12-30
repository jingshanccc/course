package util

import "strings"

//GeneratePageSql: 构建分页模糊搜索 sql 语句
func GeneratePageSql(createTime []string, blurry string, sort []string, blurryProperties []string, beforeOrder string) (string, string) {
	// blurry 模糊搜索字段 sort 排序数组 ["id,desc","name,asc"]
	forCount := ""
	and := true
	if createTime != nil {
		and = false
		forCount += "where x.create_time between '" + createTime[0] + "' and '" + createTime[1] + "'"
	}
	if blurry != "" {
		if and {
			forCount += "where "
		} else {
			forCount += " and "
		}
		for i, property := range blurryProperties {
			forCount += "x." + property + " like '%" + blurry + "%'"
			if i != len(blurryProperties)-1 {
				forCount += " or "
			}
		}
	}
	forPage := forCount + beforeOrder
	if sort != nil {
		forPage += " order by "
		for i, s := range sort {
			sorts := strings.Split(s, ",")
			forPage += "x." + sorts[0] + " " + sorts[1]
			if i != len(sort)-1 {
				forPage += ","
			}
		}
	}
	forPage += " limit ?,?"
	return forCount, forPage
}
