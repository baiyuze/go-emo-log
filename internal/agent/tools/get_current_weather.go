package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/mozillazg/go-pinyin"
)

func GetCurrentWeather(location string) (string, error) {
	pinyinName := getCityPinyin(location)
	// location, idErr :=
	ids, err := getCityIDs(pinyinName)

	if err != nil {
		return "", err
	}

	fmt.Println(pinyinName, location, "---pinyinName----")
	apiKey := os.Getenv("QWEATHER_API_KEY")

	client := resty.New()
	url := fmt.Sprintf("https://nt5n8kqhju.re.qweatherapi.com/v7/weather/now?location=%s", ids[0])
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

type CityLookupResponse struct {
	Code     string `json:"code"`
	Location []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"location"`
}

func getCityIDs(query string) ([]string, error) {
	client := resty.New()
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

	var ids []string
	for _, loc := range cityResp.Location {
		ids = append(ids, loc.ID)
	}
	fmt.Println(ids, "--城市ID---")
	return ids, nil

}

func getCityPinyin(hans string) string {
	p := pinyin.NewArgs()
	vals := pinyin.Pinyin(hans, p)
	var build strings.Builder
	for _, val := range vals {
		build.WriteString(val[0])
	}
	nameStr := build.String()
	return nameStr
}
