package main

import (
	"fmt"
	"github.com/sunfmin/checkdailyplan/config"
	"github.com/sunfmin/qortexapi"
	"github.com/sunfmin/qortexapiclient/client"
	"log"
	"strings"
	"time"
)

func main() {

	var err error

	client.ApiDomain = qortexApiDomain

	var session string
	session, err = getQortexSession()

	var authService qortexapi.AuthUserService
	authService, err = client.DefaultPublicService.GetAuthUserService(session, config.OrgId)
	if err != nil {
		return
	}

	var users []*qortexapi.EmbedUser

	users, _, err = authService.GetAllUsers()

	if err != nil {
		log.Println(err)
	}
	// for _, u := range users {
	// 	log.Println(u.Name)
	// }

	// log.Println(users)

	var entries []*qortexapi.Entry
	entries, err = authService.GetGroupEntries(config.CheckGroupId, "", "", int(config.MaxCheckEntries), true)

	if err != nil {
		log.Println(err)
		return
	}

	// log.Printf("Got %d Entries\n", len(entries))

	var writtenUsernames []string
	zone, _ := time.LoadLocation("Local")
	yesterday := time.Now().In(zone).Add(-24 * time.Hour).Format("2006-01-02")
	log.Println("Yesterday: ", yesterday)
	for _, entry := range entries {
		// log.Println(entry.Author.Name, entry.BumpedUpAt.In(zone), entry.UpdatedAt.In(zone))
		if entry.BumpedUpAt.In(zone).Format("2006-01-02") >= yesterday ||
			entry.UpdatedAt.In(zone).Format("2006-01-02") >= yesterday {
			writtenUsernames = append(writtenUsernames, entry.Author.Name)
		}
	}

	leftUsers := removeUsersByNames(users, writtenUsernames)

	for _, lu := range leftUsers {
		fmt.Printf("@%s\n", lu.Name)
	}
	// log.Println(writtenUsernames)

}

var qortexSession string
var qortexApiDomain string = "https://qortex.com/api"

func removeUsersByNames(users []*qortexapi.EmbedUser, names []string) (leftUsers []*qortexapi.EmbedUser) {
	for _, user := range users {
		if strings.Contains(config.IgnoreUserNameList, user.Name) {
			continue
		}
		if !exist(names, user) {
			leftUsers = append(leftUsers, user)
		}
	}
	return
}

func exist(names []string, user *qortexapi.EmbedUser) bool {
	for _, name := range names {
		if user.Name == name {
			return true
		}
	}
	return false
}

func getQortexSession() (session string, err error) {
	if qortexSession != "" {
		session = qortexSession
		return
	}

	client.ApiDomain = qortexApiDomain
	session, err = client.DefaultPublicService.GetSession(config.LoginEmail, config.LoginPassword, "en")
	return
}
