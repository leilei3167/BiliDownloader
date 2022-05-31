package parser

import (
	"errors"
	"fmt"
	"github.com/leilei3167/BiliDownloader/internal/engine"
	"github.com/leilei3167/BiliDownloader/internal/fetcher"
	"math"

	"regexp"
	"strconv"
)

// ParseURL 解析url类型,构建请求
func ParseURL(url string) (*engine.Request, error) {
	var idType = ""
	var aid int64
	var bvid string
	var req *engine.Request
	//	var upid int64

	var params []string

	// bvid
	bvidRegexp := regexp.MustCompile(`/?(BV\w+)[/?]?`)
	params = bvidRegexp.FindStringSubmatch(url)
	if params != nil {
		idType = "bvid"
		bvid = params[1]
	}
	// aid
	aidRegexp := regexp.MustCompile(`/?(av\d+)/?`)
	params = aidRegexp.FindStringSubmatch(url)
	if params != nil {
		idType = "aid"
		aid, _ = strconv.ParseInt(params[1], 10, 64)
	}

	//根据不同的url创建请求Req
	switch idType {
	case "aid":
		req = getRequestByAid(aid)
		return req, nil
	case "bvid":
		tempbid := bv2av(bvid)
		req = getRequestByAid(tempbid)
		return req, nil
	default:

		return nil, errors.New("无法解析url")
	}

}

//根据av号创建req
func getRequestByAid(aid int64) *engine.Request {
	//组装api
	reqUrl := fmt.Sprintf(_getCidUrlTemp, aid)
	videoAid := NewVideoAidInfo(aid, fmt.Sprintf("%d", aid))
	reqParseFunction := GenGetAidChildrenParseFun(videoAid)
	req := engine.NewRequest(reqUrl, reqParseFunction, fetcher.DefaultFetcher)
	return req
}

//https://www.bilibili.com/video/BV1iT4y1B7um?spm_id_from=333.851.b_7265636f6d6d656e64.1

// bv2av bv转av
func bv2av(x string) int64 {
	tr = make(map[string]int)
	for i := 0; i < 58; i++ {
		tr[string(table[i])] = i
	}
	r := 0
	for i := 0; i < 6; i++ {
		r += tr[string(x[s[i]])] * int(math.Pow(float64(58), float64(i)))
	}
	return int64((r - add) ^ xor)
}
