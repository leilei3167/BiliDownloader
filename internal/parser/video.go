package parser

import (
	"github.com/leilei3167/BiliDownloader/internal/engine"
	"github.com/leilei3167/BiliDownloader/internal/fetcher"
	"github.com/tidwall/gjson"
)

// 大视频所以分成了不同部分提交，但是最终显示的只有一个视频文件
func GenVideoDownloadParseFun(videoCid *VideoCid) engine.ParseFunc {
	return func(contents []byte, url string) engine.ParseResult {
		retParseResult := engine.ParseResult{}

		durlSlice := gjson.GetBytes(contents, "durl").Array()
		videoCid.AllOrder = int64(len(durlSlice))
		item := engine.NewItem(videoCid)
		retParseResult.Items = append(retParseResult.Items, item)

		for _, i := range durlSlice {
			video := &Video{Order: i.Get("order").Int(), ParCid: videoCid}
			videoUrl := i.Get("url").String()
			req := engine.NewRequest(videoUrl, recordCidParseFun(video), fetcher.GenVideoFetcher(video))
			retParseResult.Requests = append(retParseResult.Requests, req)
		}
		return retParseResult
	}
}

func recordCidParseFun(Video *Video) engine.ParseFunc {
	return func(contents []byte, url string) engine.ParseResult {
		var retResult engine.ParseResult
		item := engine.NewItem(Video)
		retResult.Items = append(retResult.Items, item)
		return retResult
	}
}
