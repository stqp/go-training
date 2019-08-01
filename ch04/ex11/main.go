package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type response struct {
	*http.Response
	NextPage  int
	PrevPage  int
	FirstPage int
	LastPage  int
}

const (
	baseURL = "https://api.github.com/"
	repo    = "go-training"
	owner   = "stqp"
)

var (
	gitUser = "stqp"
	gitPass = os.Getenv("git_pass")
	client  = &http.Client{}
)

// Issue is Issue
type Issue struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Labels    []string `json:"labels,omitempty"`
	Assignee  string   `json:"assignee,omitempty"`
	State     string   `json:"state,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
}

func l(data interface{}) {
	fmt.Println(data)
}

func create(issue Issue) {

	url := baseURL + fmt.Sprintf("repos/%v/%v/issues", owner, repo)
	data, err := json.Marshal(issue)
	if err != nil {
		panic(err)
	}

	l(bytes.NewBuffer(data))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	resp, err := do(req)
	l(resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Content-type", "application/json")
	req.SetBasicAuth(gitUser, gitPass)
	return client.Do(req)
}

func edit(issue Issue, id int) {
	url := baseURL + fmt.Sprintf("repos/%v/%v/issues/%v", owner, repo, id)
	data, err := json.Marshal(issue)
	if err != nil {
		panic(err)
	}
	l(bytes.NewBuffer(data))
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	resp, err := do(req)
	l(resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func get(id int) {
	url := baseURL + fmt.Sprintf("repos/%v/%v/issues/%v", owner, repo, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	res, err := do(req)
	if err != nil {
		panic(err)
	}
	l(res.Status)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var issue Issue
	json.Unmarshal(body, &issue)

	l(issue)
}

func close(id int) {
	edit(Issue{State: "close"}, id)
}

func usageAndExit() {
	l("invalid command call.")
	os.Exit(1)
}

var (
	title   = flag.String("title", "", "issue title")
	body    = flag.String("body", "", "issue body")
	issueID = flag.Int("id", -1, "issue id")
	command = flag.String("command", "get", "command")
)

func main() {

	if len(os.Args) < 2 {
		usageAndExit()
	}

	flag.Parse()

	issue := Issue{
		Title:     *title,
		Body:      *body,
		Labels:    nil,
		Assignee:  "",
		State:     "",
		Milestone: 0,
		Assignees: nil,
	}

	switch *command {
	case "create":
		create(issue)
	case "edit":
		edit(issue, *issueID)
	case "get":
		get(*issueID)
	case "close":
		close(*issueID)
	default:
		l("You must specify command.")
	}

}
