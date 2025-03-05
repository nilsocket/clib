// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/

package chttp

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/nilsocket/clib/cconst"
)

type Response struct {
	Data  any    `json:"data"`
	Error string `json:"error,omitempty"`
}

// Send writes given `data` or `err` to `w` as `json`
func Send(w http.ResponseWriter, data any, err error) {

	resp := Response{
		Data:  data,
		Error: err.Error(),
	}

	if err != nil {
		resp.Data = nil
	}

	json.NewEncoder(w).Encode(resp)
}

// SendAsAttachment attaches given `data` as downloadable attachment with given `filename`
// it supports ranged requests
func SendAsAttachment(w http.ResponseWriter, r *http.Request, filename string, data io.ReadSeeker, err error) {

	if err != nil {
		Send(w, nil, err)
		return
	}

	w.Header().Set(cconst.HeaderContentDisposition, "attachment; filename="+filename+"\"")
	http.ServeContent(w, r, filename, time.Now(), data)
}
