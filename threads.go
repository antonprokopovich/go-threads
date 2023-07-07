package threads

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

const (
	apiUrl   = "https://www.threads.net/api/graphql"
	tokenUrl = "https://www.threads.net/@instagram"
)

// Threads implements Threads.net API wrapper.
type Threads struct {
	token   string
	headers http.Header
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

	t.token, err = t.GetToken()
	if err != nil {
		return nil, err
	}

	t.headers = map[string][]string{
		"Content-Type":   {"application/x-www-form-urlencoded"},
		"X-IG-App-ID":    {"238260118697367"},
		"X-FB-LSD":       {t.token},
		"Sec-Fetch-Site": {"same-origin"},
	}

	return t, nil
}

// GetToken fetches a token.
func (t *Threads) GetToken() (string, error) {
	resp, err := http.Get(tokenUrl)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	bodyStr := string(body)

	tokenKeyPos := strings.Index(bodyStr, "\"token\"")
	token := bodyStr[tokenKeyPos+9 : tokenKeyPos+31]

	return token, nil
}

func (t *Threads) postRequest(variables map[string]int, docID string) ([]byte, error) {
	variablesStr, err := json.Marshal(variables)
	if err != nil {
		return nil, err
	}

	data := RequestData{
		Lsd:       t.token,
		Variables: string(variablesStr),
		DocID:     docID,
	}

	dataStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(dataStr))
	if err != nil {
		return nil, err
	}

	req.Header = t.headers

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
	return t.postRequest(variables, "5587632691339264")
}

// GetUser fetches a user.
func (t *Threads) GetUser(id int) ([]byte, error) {
	variables := map[string]int{"userID": id}
	return t.postRequest(variables, "23996318473300828")
}

// GetUserThreads fetches a user's Threads.
func (t *Threads) GetUserThreads(id int) ([]byte, error) {
	variables := map[string]int{"userID": id}
	return t.postRequest(variables, "6232751443445612")
}

// GetUserReplies fetches a user's replies.
func (t *Threads) GetUserReplies(id int) ([]byte, error) {
	variables := map[string]int{"userID": id}
	return t.postRequest(variables, "6307072669391286")
}
