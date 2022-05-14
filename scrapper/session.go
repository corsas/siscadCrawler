package scrapper

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func isWebSiteWorking(c *http.Client, reqUrl string) bool {
	resp, err := c.Get(reqUrl)
	if err != nil || resp.StatusCode >= 400 {
		return false
	}
	resp.Body.Close()
	return true
}

func loginAttempt(c *http.Client, loginUrl, authUrl, successfulPattern string) bool {
	respAuthentication, err := c.PostForm(loginUrl, getCredentials())
	if err != nil || respAuthentication.StatusCode >= 400 {
		return false
	}

	respAuthentication.Body.Close()

	respAuthorization, errAuth := c.Get(authUrl)
	if errAuth != nil || respAuthentication.StatusCode >= 400 {
		return false
	}

	defer respAuthorization.Body.Close()
	body, _ := ioutil.ReadAll(respAuthorization.Body)
	bodyString := string(body)

	if strings.Contains(bodyString, successfulPattern) {
		return true
	}
	return false
}

func getCredentials() map[string][]string {
	return url.Values{
		"login":    {os.Getenv("TEST_USERNAME")},
		"password": {os.Getenv("TEST_PASSWORD")},
	}
}
