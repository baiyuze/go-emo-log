package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/henrylee2cn/pholcus/common/pinyin"
)

type Location struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type CityLookupResponse struct {
	Code     string     `json:"code"`
	Location []Location `json:"location"`
}

var client = resty.New()

func GetCurrentWeather(location string) (string, error) {
	apiKey := os.Getenv("QWEATHER_API_KEY")

	url := fmt.Sprintf("https://nt5n8kqhju.re.qweatherapi.com/v7/weather/now?location=%s", location)
	fmt.Println(url)
	resp, err := client.R().
		SetHeader("X-QW-Api-Key", apiKey).
		Get(url)
	if err != nil {
		return "", err
	}

	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("unexpected status: %d, body: %s", resp.StatusCode(), resp.Body())
	}

	return string(resp.Body()), nil
}

func GetCityIDs(query string) ([]Location, error) {
	apiKey := os.Getenv("QWEATHER_API_KEY")

	resp, err := client.R().
		SetHeader("X-QW-Api-Key", apiKey).
		Get(fmt.Sprintf("https://nt5n8kqhju.re.qweatherapi.com/geo/v2/city/lookup?location=%s", query))

	if err != nil {
		return nil, fmt.Errorf("请求出错: %v", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status: %d, body: %s", resp.StatusCode(), resp.Body())
	}

	var cityResp CityLookupResponse
	if err := json.Unmarshal(resp.Body(), &cityResp); err != nil {
		return nil, fmt.Errorf("json解析失败: %v", err)
	}

	if cityResp.Code != "200" {
		return nil, fmt.Errorf("接口返回错误码: %s", cityResp.Code)
	}

	// var ids []string
	// for _, loc := range cityResp.Location {
	// 	ids = append(ids, loc.ID)
	// }
	fmt.Printf("%+v===>", cityResp.Location)
	return cityResp.Location, nil

}

func GetCityPinyin(hans string) string {
	vals := pinyin.Pinyin(hans, pinyin.Args{})
	var build strings.Builder
	for _, val := range vals {
		build.WriteString(val[0])
	}
	nameStr := build.String()
	fmt.Printf("%+v\n", vals)
	return nameStr
}
