package engine

type ParseFunc func(contents []byte, url string) ParseResult

type ParseResult struct {
	Requests []*Request
	Items    []*Item
}

type Request struct {
	Url           string
	ParseFunction ParseFunc //Url和contents解析为请求
	FetchFun      fetcher.FetchFun
}

func NewRequest(url string, parseFunction ParseFunc, fetchFun fetcher.FetchFun) *Request {
	return &Request{Url: url, ParseFunction: parseFunction, FetchFun: fetchFun}
}

type Item struct {
	Payload interface{}
}

func NewItem(payload interface{}) *Item {
	return &Item{Payload: payload}
}
