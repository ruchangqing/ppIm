package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"ppIm/api"
	"ppIm/global"
	"ppIm/model"
	"strconv"
)

// 附近的用户列表
func GeoUsers(ctx *gin.Context) {
	// 经纬度校验
	longitude, err1 := strconv.ParseFloat(ctx.PostForm("longitude"), 64)
	latitude, err2 := strconv.ParseFloat(ctx.PostForm("latitude"), 64)
	if err1 != nil || err2 != nil {
		fmt.Println(err1)
		fmt.Println(err2)
		api.R(ctx, 500, "数据非法", nil)
		return
	}
	// 距离范围，默认100
	distance := ctx.PostForm("distance")
	if distance == "" {
		distance = "100"
	}

	query := elastic.NewGeoDistanceQuery("location").Distance(distance + "km").Lat(latitude).Lon(longitude)
	sort := elastic.NewGeoDistanceSort("location").Point(latitude, longitude).Asc().DistanceType("arc").Unit("km")
	res, err3 := global.Elasticsearch.Search().Index("user_location").Query(query).SortBy(sort).Do(context.Background())
	if err3 != nil {
		fmt.Println(err3)
		api.R(ctx, 500, "数据非法", nil)
		return
	}

	// 解析es数据数组
	type Data map[string]interface{}
	// es数组变量
	var data []Data

	// 循环es结果
	for _, hit := range res.Hits.Hits {
		var userLocation model.UserLocation
		err := json.Unmarshal(hit.Source, &userLocation) // json解析结果
		if err != nil {
			fmt.Println(err)
		}
		temp := make(Data)
		temp["uid"] = userLocation.Uid
		temp["location"] = userLocation.Location
		temp["distance"] = hit.Sort
		data = append(data, temp)
	}

	api.Rt(ctx, 200, "ok", gin.H{
		"users": data,
	})
}
