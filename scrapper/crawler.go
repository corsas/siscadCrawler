package scrapper

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func getGradesURL(c *http.Client, reqURL, selectorPath string) []string {

	resp, err := c.Get(reqURL)
	if err != nil || resp.StatusCode >= 400 {
		return nil
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	nodes := doc.Find(selectorPath)
	var URLs []string

	for _, v := range nodes.Nodes[1:] {

		if v.Attr[0].Val == "listHeader" {
			break
		}

		URLs = append(URLs, v.FirstChild.NextSibling.FirstChild.FirstChild.FirstChild.Attr[0].Val)
	}

	return URLs
}

func getGrades(c *http.Client, gradeURL, gradeSelector, SubjectNameSelector string) string {
	resp, err := c.Get(gradeURL)
	if err != nil || resp.StatusCode >= 400 {
		return ""
	}
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	name := doc.Find(SubjectNameSelector)
	grades := doc.Find(gradeSelector)

	var output []string

	if name.Nodes != nil {
		output = append(output, name.Nodes[0].FirstChild.Data)
	}
	if grades.Nodes != nil {

		for _, v := range grades.Nodes {
			output = append(output, v.FirstChild.Data)
		}

	}

	return strings.Join(output, " - ")
}
