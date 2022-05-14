package scrapper

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"siscadCrawler/util"
	"time"
)

func Run() {

	c := clientConfig()

	// Session
	if !isWebSiteWorking(&c, util.BaseURL) {
		return
	}
	if !loginAttempt(&c, util.BaseURL, util.GradesURL, util.AuthenticatedPattern) {
		return
	}

	// Crawler
	var results []string

	gradeList := getGradesURL(&c, util.GradesURL, util.GradesSelector)

	for _, v := range gradeList {
		getGrades(&c, util.Host+v, util.IndividualGradeSelector, util.SubjectNameSelector)
		results = append(results, getGrades(&c, util.Host+v, util.IndividualGradeSelector, util.SubjectNameSelector))
	}

	fmt.Println(results)
}

func clientConfig() http.Client {

	jar, _ := cookiejar.New(nil)
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	return http.Client{
		Transport:     t,
		CheckRedirect: nil,
		Jar:           jar,
		Timeout:       time.Second * 4,
	}

}
