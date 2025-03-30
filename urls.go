// Package client provides methods, values and urls for interacting with Yandex Tracker
package client

var issuesBaseURL = "/issues/"
var issuesCreateURL = issuesBaseURL
var issuesCountURL = issuesBaseURL + "_count"
var issuesSearchURL = issuesBaseURL + "_search"
var issuesModifyURL = issuesBaseURL + "{issue_id}"
var issuesGetURL = issuesBaseURL + "{issue_id}"
var issuesModifyStatusURL = issuesBaseURL + "{issue_id}/transitions/{transition_id}/_execute"
