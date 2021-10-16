package xueqiu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
)

var (
	// 股票行情
	API_PATH_QUOTE = "https://stock.xueqiu.com/v5/stock/quote.json?symbol=%s&extend=detail"
)

type XQApiClient struct {
	httpClient *httpclient.Client
}

func (xq XQApiClient) GetQuote(code string) (*QuoteResp, error) {
	url := fmt.Sprintf(API_PATH_QUOTE, code)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	// xq_a_token=cddd72fc3c69466c73a5a4d4ed47425fd46595e7;
	request.AddCookie(&http.Cookie{
		Name:  "xq_a_token",
		Value: "cddd72fc3c69466c73a5a4d4ed47425fd46595e7",
	})

	response, err := xq.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	data := new(QuoteResp)
	if err = xq.parseResponse(response, data); err != nil {
		return nil, err
	}

	return data, nil
}

func (xq XQApiClient) parseResponse(resp *http.Response, data interface{}) error {

	t := reflect.TypeOf(data)
	if t.Elem() == nil {
		return fmt.Errorf("数据类型不符,请传递指针数据类型")
	}

	if resp.Body == nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("状态码错误,状态码:%d", resp.StatusCode)
	}

	var respInfo xqResponse
	if err := json.Unmarshal(body, &respInfo); err != nil {
		return err
	}

	if respInfo.ErrorCode != 0 {
		return err
	}

	if err = json.Unmarshal(respInfo.Data, data); err != nil {
		return err
	}

	return nil
}

type xqResponse struct {
	ErrorCode        int             `json:"error_code"`
	ErrorDescription string          `json:"error_description"`
	Data             json.RawMessage `json:"data"`
}

func NewXQApiClient() *XQApiClient {
	httpclient := httpclient.NewClient(
		httpclient.WithHTTPTimeout(30*time.Second),
		httpclient.WithRetryCount(1),
	)

	return &XQApiClient{
		httpClient: httpclient,
	}
}
