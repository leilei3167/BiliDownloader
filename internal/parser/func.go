package parser

import (
	"crypto/md5"
	"fmt"
	"github.com/leilei3167/BiliDownloader/internal/engine"
	"github.com/leilei3167/BiliDownloader/internal/fetcher"
	"github.com/leilei3167/BiliDownloader/pkg/tool"
	"github.com/tidwall/gjson"
	"strconv"
)

func GenGetAidChildrenParseFun(videoAid *VideoAid) engine.ParseFunc {
	return func(contents []byte, url string) engine.ParseResult {
		var retParseResult engine.ParseResult

		if videoAid.Title == strconv.FormatInt(videoAid.Aid, 10) {
			//gjson从一个json字节流中快速取出数据
			title := gjson.GetBytes(contents, "data.title").String()
			title = tool.TitleEdit(title) // remove special characters
			videoAid.Title = title
		} //转换为十进制表示字符串
		data := gjson.GetBytes(contents, "data.pages").Array()
		fmt.Println("即将开始下载：", videoAid.Title)
		appKey, sec := tool.GetAppKey(_entropy)

		var videoTotalPage int64
		for _, i := range data {
			cid := i.Get("cid").Int()
			page := i.Get("page").Int()
			part := i.Get("part").String()
			part = tool.TitleEdit(part) //remove special characters
			videoCid := NewVideoCidInfo(cid, videoAid, page, part)
			videoTotalPage += 1
			cidStr := strconv.FormatInt(videoCid.Cid, 10)

			params := fmt.Sprintf(_paramsTemp, appKey, cidStr, _quality, _quality)
			chksum := fmt.Sprintf("%x", md5.Sum([]byte(params+sec)))

			urlApi := fmt.Sprintf(_playApiTemp, params, chksum)

			req := engine.NewRequest(urlApi, GenVideoDownloadParseFun(videoCid), fetcher.DefaultFetcher)
			retParseResult.Requests = append(retParseResult.Requests, req)
		}

		videoAid.SetPage(videoTotalPage)
		item := engine.NewItem(videoAid)
		retParseResult.Items = append(retParseResult.Items, item)

		return retParseResult

	}

}
func GetRequestByAid(aid int64) *engine.Request {
	reqUrl := fmt.Sprintf(_getCidUrlTemp, aid)
	videoAid := NewVideoAidInfo(aid, fmt.Sprintf("%d", aid))
	reqParseFunction := GenGetAidChildrenParseFun(videoAid)
	req := engine.NewRequest(reqUrl, reqParseFunction, fetcher.DefaultFetcher)
	return req
}
