package threads

import (
	"encoding/json"
	"net/url"
	"regexp"
	"strconv"

	"github.com/pkg/errors"

	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	getUserDocID        = "23996318473300828"
	getUserThreadsDocID = "6232751443445612"
	getUserRepliesDocID = "6307072669391286"
	getPostDocID        = "5587632691339264"
	getPostLikersDocID  = "9360915773983802"
)

const (
	maxRedirectsNum = 20
	apiUrl          = "https://www.threads.net/api/graphql"
	tokenUrl        = "https://www.threads.net/@instagram"
)

// Threads implements Threads.net API wrapper.
type Threads struct {
	token          string
	defaultHeaders http.Header
}

// RequestData stores the request payload.
type RequestData struct {
	Lsd       string `json:"lsd"`
	Variables string `json:"variables"`
	DocID     string `json:"doc_id"`
}

// NewThreads constructs a Threads instance.
func NewThreads() (t *Threads, err error) {
	t = new(Threads)

	t.token, err = t.getToken()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token")
	}

	t.defaultHeaders = map[string][]string{
		"Authority":       {"www.threads.net"},
		"Accept":          {"*/*"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"Cache-Control":   {"no-cache"},
		"Content-Type":    {"application/x-www-form-urlencoded"},
		"Connection":      {"keep-alive"},
		"Origin":          {"https://www.threads.net"},
		"Pragma":          {"no-cache"},
		"Sec-Fetch-Site":  {"same-origin"},
		"User-Agent":      {"golang"},
		"X-ASBD-ID":       {"129477"},
		"X-IG-App-ID":     {"238260118697367"},
		"X-FB-LSD":        {t.token},
	}

	return t, nil
}

func (t *Threads) getToken() (string, error) {
	req, err := http.NewRequest(http.MethodGet, tokenUrl, nil)
	if err != nil {
		return "", errors.Wrap(err, "http.NewRequest")
	}

	req.Header = map[string][]string{
		"User-Agent": {"golang"},
	}

	client := &http.Client{
		CheckRedirect: func() func(req *http.Request, via []*http.Request) error {
			redirects := 0
			return func(req *http.Request, via []*http.Request) error {
				if redirects > maxRedirectsNum {
					return errors.New(fmt.Sprintf("stopped after %d redirects", maxRedirectsNum))
				}
				redirects++

				return nil
			}
		}(),
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "http Get request failed")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read response body")
	}

	bodyStr := string(body)

	tokenKeyPos := strings.Index(bodyStr, "\"token\"")
	token := bodyStr[tokenKeyPos+9 : tokenKeyPos+31]

	return token, nil
}

func (t *Threads) postRequest(variables map[string]int, docID string, headers http.Header) ([]byte, error) {
	variablesStr, err := json.Marshal(variables)
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("lsd", t.token)
	data.Set("variables", string(variablesStr))
	data.Set("doc_id", docID)

	req, err := http.NewRequest(http.MethodPost, apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header = headers

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetPost fetches a post.
func (t *Threads) GetPost(id int) ([]byte, error) {
	variables := map[string]int{"postID": id}

	headers := t.defaultHeaders.Clone()
	headers.Add("X-FB-Friendly-Name", "BarcelonaPostPageQuery")

	return t.postRequest(variables, getPostDocID, headers)
}

// GetPostLikers fetches all users who like the post.
func (t *Threads) GetPostLikers(id int) ([]byte, error) {
	variables := map[string]int{"postID": id}

	return t.postRequest(variables, getPostLikersDocID, t.defaultHeaders)
}

// GetUser fetches a user.
func (t *Threads) GetUser(id int) ([]byte, error) {
	variables := map[string]int{"userID": id}

	headers := t.defaultHeaders.Clone()
	headers.Add("X-FB-Friendly-Name", "BarcelonaProfileRootQuery")

	return t.postRequest(variables, getUserDocID, headers)
}

// GetUserThreads fetches a user's Threads.
func (t *Threads) GetUserThreads(id int) ([]byte, error) {
	variables := map[string]int{"userID": id}

	headers := t.defaultHeaders.Clone()
	headers.Add("X-FB-Friendly-Name", "BarcelonaProfileThreadsTabQuery")

	return t.postRequest(variables, getUserThreadsDocID, headers)
}

// GetUserReplies fetches a user's replies.
func (t *Threads) GetUserReplies(id int) ([]byte, error) {
	variables := map[string]int{"userID": id}

	headers := t.defaultHeaders.Clone()
	headers.Add("X-FB-Friendly-Name", "BarcelonaProfileRepliesTabQuery")

	return t.postRequest(variables, getUserRepliesDocID, headers)
}

// GetUserID fetches user's ID by username.
func (t *Threads) GetUserID(username string) (int, error) {
	baseURL := fmt.Sprintf("https://www.threads.net/@%s", username)

	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		return -1, err
	}

	req.Header = t.defaultHeaders.Clone()

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Referer", baseURL)
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")

	req.Header.Del("X-ASBD-ID")
	req.Header.Del("X-FB-LSD")
	req.Header.Del("X-IG-App-ID")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	userIdRegex := regexp.MustCompile(`"user_id":"(\d+)"`)
	userIdStr := userIdRegex.FindStringSubmatch(string(body))[1]

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return -1, err
	}

	return userId, nil
}
