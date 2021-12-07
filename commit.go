package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

type Commit struct {
	ID         string `json:"id"`
	ShortID    string `json:"short_id"`
	CreatedAt  string `json:"created_at"`
	Title      string `json:"title"`
	Message    string `json:"message"`
	CommitDate string `json:"committed_date"`
	WebURL     string `json:"web_url"`
}

func (c *Commit) generateBadgeItem() string {
	return fmt.Sprintf("![img](https://img.shields.io/static/v1?label=commit&message=%s&color=brightgreen)[%s](%s)<br>", c.ShortID, c.Title, c.WebURL)
}

func RequestCommitInfo(personalAccessToken string, endpoint string, repoID int, sinceDateString, endDateString string) ([]Commit, error) {
	//1. 确定上周起止日期
	if ok, err := regexp.Match(`^\d\d\d\d-\d\d-\d\d$`, []byte(sinceDateString)); err != nil || !ok {
		return nil, fmt.Errorf("invalid since date %s", sinceDateString)
	}
	if ok, err := regexp.Match(`^\d\d\d\d-\d\d-\d\d$`, []byte(endDateString)); err != nil || !ok {
		return nil, fmt.Errorf("invalid since date %s", endDateString)
	}
	//2. 请求实验室的Gitlab API
	gitlabURL, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}
	gitlabURL.Path = fmt.Sprintf("/api/v4/projects/%d/repository/commits", repoID)
	var param = gitlabURL.Query()
	param.Add("since", fmt.Sprintf("%sT00:00:00Z", sinceDateString))
	param.Add("end", fmt.Sprintf("%sT23:59:59Z", endDateString))
	gitlabURL.RawQuery = param.Encode()

	fmt.Println(gitlabURL.String())
	req, err := http.NewRequest("GET", gitlabURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("PRIVATE-TOKEN", personalAccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var res []Commit
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil

}
