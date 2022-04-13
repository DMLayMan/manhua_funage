package service

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

const (
	token = "12345678901"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	signature := r.FormValue("signature")
	timestamp := r.FormValue("timestamp")
	nonce := r.FormValue("nonce")

	authSlice := []string{token, timestamp, nonce}
	sort.Slice(authSlice, func(i, j int) bool {
		return authSlice[i] < authSlice[j]
	})

	sortedAuthInfo := strings.Join(authSlice, "")

	encryptStr := fmt.Sprintf("%x", sha1.Sum([]byte(sortedAuthInfo)))
	var result = "false"
	if encryptStr == signature {
		result = "true"
	}
	_, _ = w.Write([]byte(result))
}
