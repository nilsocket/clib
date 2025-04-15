package chttp

import "net/http"

func QueryParam(r *http.Request, param string) string {

	queryParams := r.URL.Query()[param]
	if len(queryParams) > 0 {
		return queryParams[0]
	}

	return ""
}
