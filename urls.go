// Package client provides methods, values and urls for interacting with Yandex Tracker
package client

var issuesBaseURL = "/issues/"
var issuesCreateURL = issuesBaseURL
var issuesCountURL = issuesBaseURL + "_count"
var issuesSearchURL = issuesBaseURL + "_search"
var issuesModifyURL = issuesBaseURL + "{issue_id}"
var issuesGetURL = issuesBaseURL + "{issue_id}"
var issueGetTransitionsURL = issuesBaseURL + "{issue_id}/transitions"
var issuesModifyStatusURL = issuesBaseURL + "{issue_id}/transitions/{transition_id}/_execute"
var issueAppendCommentURL = issuesBaseURL + "{issue_id}/comments"
var issueGetCommentsURL = issuesBaseURL + "{issue_id}/comments"
var issueGetCommentURL = issuesBaseURL + "{issue_id}/comments/{comment_id}"
var issueUpdateCommentURL = issuesBaseURL + "{issue_id}/comments/{comment_id}"
var issueDeleteCommentURL = issuesBaseURL + "{issue_id}/comments/{comment_id}"
var issueGetAttachmentsURL = issuesBaseURL + "{issue_id}/attachments"
var issueGetAttachmentURL = issuesBaseURL + "{issue_id}/attachments/{attachment_id}"
var issueAttachFileURL = issuesBaseURL + "{issue_id}/attachments"

var attachmentsBase = "/attachments/"
var attachmentUploadURL = attachmentsBase

var prioritiesBaseURL = "/priorities/"
var prioritiesGetURL = prioritiesBaseURL
var priorityGetURL = prioritiesBaseURL + "{priority_id}"

var myselfURL = "/myself"

var userBaseURL = "/users/"
var usersGetURL = userBaseURL
var userGetURL = userBaseURL + "{login_or_user_id}"

var componentBaseURL = "/components/"
var componentCreateURL = componentBaseURL
var componentsGetURL = componentBaseURL
var componentGetURL = componentBaseURL + "{component_id}"
var componentUpdateURL = componentBaseURL + "{component_id}"
