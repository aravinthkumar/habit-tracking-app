package controller

import (
	"net/http"
	"regexp"
)

type habitController struct {
	habitIDPattern *regexp.Regexpß
}

func (hc habitController) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path = "/habits" {
		switch req.Method {
		case http.MethodGet:
			uc.get
			
		}
	}

}
