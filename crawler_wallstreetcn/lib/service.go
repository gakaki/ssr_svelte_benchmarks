package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var BaseUrl = "https://api-one-wscn.awtmt.com"
var BaseUrlAPI = BaseUrl + "/apiv1"
var BaseUrlContent = BaseUrlAPI + "/content"
var UrlInformationFlow = BaseUrlContent + "/information-flow"
var UrlLiveFlow = BaseUrlContent + "/lives"
var UrlLessonMaster = BaseUrlAPI + "/vipv3/club_articles"
var UrlMarketReal = "https://api-ddc-wscn.awtmt.com/market/real"
var UrlArticlesHot = BaseUrlContent + "/articles/hot"

func HelperRequest[T any](url string, queryParams map[string]string) (*T, error) {
	toStruct, err := RequestToStruct[BodyData[T]](url, queryParams)
	if err != nil {
		return nil, err
	}
	return &toStruct.Data, nil
}

// 资讯
func GetInformationFlow() (*InformationFLow, error) {
	queryParams := map[string]string{
		"channel": "global",
		"accept":  "article",
		"cursor":  "", // Empty string for the cursor parameter as it's not provided
		"limit":   "20",
		"action":  "upglide",
	}
	return HelperRequest[InformationFLow](UrlInformationFlow, queryParams)
}

// 快讯
func GetLiveFlow() (*LiveFLow, error) {
	queryParams := map[string]string{
		"channel":    "global-channel",
		"client":     "pc",
		"limit":      "20",
		"first_page": "true",
		"accept":     "live,vip-live",
	}
	return HelperRequest[LiveFLow](UrlLiveFlow, queryParams)
}

// 大师课-最近更新 可能不太对凑活用吧
func GetLessonMaster() (*LessonMaster, error) {
	queryParams := map[string]string{
		"vip_type": "master_course",
		"limit":    "20",
		"cursor":   "",
	}
	return HelperRequest[LessonMaster](UrlLessonMaster, queryParams)
}

// 各种指数
func GetMarketReal() (*MarketReal, error) {
	queryParams := map[string]string{
		"prod_code": "000001.SS,DXY.OTC,US10YR.OTC,USDCNH.OTC,399001.SZ,399006.SZ",
		"fields":    "prod_name,preclose_px,last_px,px_change,px_change_rate,price_precision",
	}
	return HelperRequest[MarketReal](UrlMarketReal, queryParams)
}

// 最新文章
func GetArticlesHot() (*Article, error) {
	queryParams := map[string]string{
		"period": "all",
	}
	return HelperRequest[Article](UrlArticlesHot, queryParams)
}

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func SetCache[T any](key string, t *T) {

	v, _ := json.Marshal(t)
	err := rdb.Set(ctx, key, v, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}
