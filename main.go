package main

import (
	"fmt"

	"net/http"

	"github.com/go-playground/webhooks/v6/github"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := github.New(github.Options.Secret("tabienpatestear"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent,
			github.PullRequestEvent,
			github.PushEvent,
			github.CommitCommentEvent,
			github.IssueCommentEvent,
			github.IssuesEvent,
			github.PullRequestReviewEvent,
			github.PullRequestReviewCommentEvent,
			github.CreateEvent,
			github.DeleteEvent,
			github.ForkEvent,
			github.GollumEvent,
			github.WatchEvent,
			github.DeploymentEvent,
			github.DeploymentStatusEvent,
			github.GollumEvent,
			github.MemberEvent,
			github.PublicEvent,
			github.TeamAddEvent,
			github.StatusEvent,
			github.PageBuildEvent,
			github.ProjectCardEvent,
			github.ProjectColumnEvent,
			github.ProjectEvent,
		)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
				fmt.Println("not a github event we asked to parse")
			}
		}
		switch payload.(type) {
		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Println(release.Release.Name)
			fmt.Printf("realese: %+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			fmt.Println(pullRequest.PullRequest.Title)
			// Do whatever you want from here...
			fmt.Println(pullRequest.Repository.Name)

		case github.CommitCommentPayload:
			commitComment := payload.(github.CommitCommentPayload)
			fmt.Println(commitComment.Comment.Body)
			// Do whatever you want from here...
			fmt.Printf("commit comment: %+v", commitComment)

		case github.PushPayload:
			push := payload.(github.PushPayload)
			fmt.Println(push.Ref)
			fmt.Println("===============================================")
			fmt.Println("")
			fmt.Println("push.HeadCommit.Message", push.HeadCommit.Message)
			fmt.Println("push.HeadCommit.Author.Name", push.HeadCommit.Author.Name)
			fmt.Println("push.HeadCommit.Author.Email", push.HeadCommit.Author.Email)
			fmt.Println("push.HeadCommit.Author.Username", push.HeadCommit.Author.Username)
			fmt.Println("push.HeadCommit.Committer.Name", push.HeadCommit.Committer.Name)
			fmt.Println("push.HeadCommit.Committer.Email", push.HeadCommit.Committer.Email)
			fmt.Println("push.HeadCommit.Committer.Username", push.HeadCommit.Committer.Username)
			fmt.Println("push.HeadCommit.Timestamp", push.HeadCommit.Timestamp)
			fmt.Println("push.HeadCommit.URL", push.HeadCommit.URL)
			fmt.Println("push.HeadCommit.ID", push.HeadCommit.ID)
			fmt.Println("push.HeadCommit.TreeID", push.HeadCommit.TreeID)
			fmt.Println("push.HeadCommit.Added", push.HeadCommit.Added)
			fmt.Println("push.HeadCommit.Removed", push.HeadCommit.Removed)
			fmt.Println("push.HeadCommit.Modified", push.HeadCommit.Modified)
			fmt.Println("push.Repository.Name", push.Repository.Name)
			fmt.Println("")
			fmt.Println("===============================================")
			// Do whatever you want from here...

		case github.IssueCommentPayload:
			issueComment := payload.(github.IssueCommentPayload)
			fmt.Println(issueComment.Comment.Body)
			// Do whatever you want from here...
			fmt.Printf("issue comment: %+v", issueComment)

		case github.IssuesPayload:
			issue := payload.(github.IssuesPayload)
			fmt.Println(issue.Issue.Title)
			// Do whatever you want from here...
			fmt.Printf("issue: %+v", issue)

		case github.PullRequestReviewPayload:
			pullRequestReview := payload.(github.PullRequestReviewPayload)
			fmt.Println(pullRequestReview.Review.Body)
			// Do whatever you want from here...
			fmt.Printf("pull request review: %+v", pullRequestReview)

		case github.PullRequestReviewCommentPayload:
			pullRequestReviewComment := payload.(github.PullRequestReviewCommentPayload)
			fmt.Println(pullRequestReviewComment.Comment.Body)
			// Do whatever you want from here...
			fmt.Printf("pull request review comment: %+v", pullRequestReviewComment)

		case github.CreatePayload:
			create := payload.(github.CreatePayload)
			fmt.Println(create.Ref)
			// Do whatever you want from here...
			fmt.Printf("create: %+v", create)

		case github.DeletePayload:
			delete := payload.(github.DeletePayload)
			fmt.Println(delete.Ref)
			// Do whatever you want from here...
			fmt.Printf("delete: %+v", delete)

		case github.ForkPayload:
			fork := payload.(github.ForkPayload)
			fmt.Println(fork.Forkee.FullName)
			// Do whatever you want from here...
			fmt.Printf("fork: %+v", fork)

		case github.GollumPayload:
			gollum := payload.(github.GollumPayload)
			fmt.Println(gollum.Pages[0].Title)
			// Do whatever you want from here...
			fmt.Printf("gollum: %+v", gollum)

		case github.WatchPayload:
			watch := payload.(github.WatchPayload)
			fmt.Println(watch.Action)
			// Do whatever you want from here...
			fmt.Printf("watch: %+v", watch)

		case github.DeploymentPayload:
			deployment := payload.(github.DeploymentPayload)
			fmt.Println(deployment.Deployment.URL)
			// Do whatever you want from here...
			fmt.Printf("deployment: %+v", deployment)

		case github.DeploymentStatusPayload:
			deploymentStatus := payload.(github.DeploymentStatusPayload)
			fmt.Println(deploymentStatus.DeploymentStatus.State)
			// Do whatever you want from here...
			fmt.Printf("deployment status: %+v", deploymentStatus)
		}

	})
	http.ListenAndServe(":3000", nil)
}
