// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/

package chttp

import "net/http"

func QueryParam(r *http.Request, param string) string {

	queryParams := r.URL.Query()[param]
	if len(queryParams) > 0 {
		return queryParams[0]
	}

	return ""
}
