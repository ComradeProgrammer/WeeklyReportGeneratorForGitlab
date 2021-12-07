package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

//自动周报生成器
func main() {
	var sinceDateString, endDateString, endpoint, personalAccessToken string
	var repoID int
	flag.StringVar(&sinceDateString, "since", "", "pull commits since this date,YYYY-MM-DD")
	flag.StringVar(&endDateString, "until", "", "pull commits before this date,YYYY-MM-DD")
	flag.StringVar(&endpoint, "endpoint", "http://10.134.180.201/", "endpoint url of gitlab")
	flag.IntVar(&repoID, "id", 11, "repository id")
	flag.StringVar(&personalAccessToken, "token", "", "personal access token of your gitlab")
	flag.Parse()

	if sinceDateString=="" && endDateString==""{
		sinceDateString,endDateString=GetLastWeekISO8601()
	}

	commits, err := RequestCommitInfo(personalAccessToken, endpoint, repoID, sinceDateString, endDateString)
	if err != nil {
		panic(err)
	}

	var buffer strings.Builder
	for _, commit := range commits {
		buffer.WriteString(commit.generateBadgeItem())
		buffer.WriteString("\n")
	}
	var durationString = fmt.Sprintf("(%s-%s)", sinceDateString, endDateString)
	res := ReplaceInTemplate(map[string]string{
		"duration": durationString,
		"commits":  buffer.String(),
	})

	ioutil.WriteFile("weekly_report"+durationString+".md", []byte(res), 0666)
	fmt.Println(GetLastWeekISO8601())
}
