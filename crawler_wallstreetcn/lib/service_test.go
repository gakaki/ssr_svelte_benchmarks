package lib

import "testing"

func TestGetInformationFlow(t *testing.T) {

	pageMain := PageMain{}

	pageMain.MarketReal, _ = GetMarketReal() //实时行情 首页

	pageMain.InformationFlow, _ = GetInformationFlow() //最新资讯

	pageMain.LiveFLow, _ = GetLiveFlow() //快讯 要闻

	pageMain.LessonMaster, _ = GetLessonMaster() //大师课 最近更新

	pageMain.Article, _ = GetArticlesHot() //最热 文章

	//全都写入Redis中 然后让ssr进行读取
	SetCache("main", &pageMain)
}
