package config

import (
	// "log"
	"os"
	"strconv"
)

var (
	LoginEmail         string
	LoginPassword      string
	CheckGroupId       string
	IgnoreUserNameList string
	MaxCheckEntries    int64
	OrgId              string
)

func init() {
	LoginEmail = envOrPanic("CDP_LOGIN_EMAIL")
	LoginPassword = envOrPanic("CDP_LOGIN_PASSWORD")
	CheckGroupId = envOrPanic("CDP_CHECK_GROUP_ID")
	IgnoreUserNameList = envOrPanic("CDP_IGNORE_USERNAME_LIST")
	MaxCheckEntries, _ = strconv.ParseInt(envOrPanic("CDP_MAX_CHECK_ENTRIES"), 10, 32)
	OrgId = envOrPanic("CDP_ORG_ID")
}

func envOrPanic(key string) (r string) {
	r = os.Getenv(key)
	if r == "" {
		panic("env " + key + " is not set")
	}
	// log.Printf("Configure: %s = %s\n", key, r)
	return
}
