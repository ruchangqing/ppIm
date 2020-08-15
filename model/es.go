package model

// 用户位置index，用来创建index
var CreateUserLocationIndex = `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"properties":{
			"uid":{
				"type":"keyword"
			},
			"location":{
				"type":"geo_point"
			}
		}
	}
}`

// 用户位置 index 数据格式，用来json解析es查询结果
type UserLocation struct {
	Uid      string
	Location string
}