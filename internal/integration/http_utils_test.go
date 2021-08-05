package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const LCAPIHTTPTimeout = 60

const (
	lcAPIHTTPAddress = "http://localhost:54012"

	v0Path     = "v0"
	chartsPath = "charts"

	userAgentHeader = "user-agent"
	userAgent       = "lc-acceptance-tests"

	contentTypeHeader = "content-type"
	contentTypeJSON   = "application/json"
)

func createChartAndParseGoodReplyHTTP(ctx context.Context, requestPath string) (*ChartReply, error) {
	reqBody, err := ioutil.ReadFile(requestPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read request fixture %s: %w", requestPath, err)
	}

	resp, err := createChartHTTP(ctx, string(reqBody))
	if err != nil {
		return nil, fmt.Errorf("unable to create chart via HTTP: %w", err)
	}

	rep, err := parseChartReplyBody(resp)
	if err != nil {
		return nil, fmt.Errorf("unable to parse chart reply body: %w", err)
	}

	return rep, nil
}

func createChartHTTP(ctx context.Context, body string) (*http.Response, error) {
	b := bytes.NewReader([]byte(body))

	url := strings.Join([]string{lcAPIHTTPAddress, v0Path, chartsPath}, "/")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, b)
	if err != nil {
		return nil, fmt.Errorf("unable to prepare HTTP request: %w", err)
	}

	req.Header.Set(userAgentHeader, userAgent)
	req.Header.Set(contentTypeHeader, contentTypeJSON)

	c := &http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to do HTTP request: %w", err)
	}

	return resp, nil
}

func parseChartReplyBody(resp *http.Response) (*ChartReply, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %w", err)
	}
	defer resp.Body.Close()

	chartRep := struct {
		ChartReply `json:"chart"`
	}{}

	if err := json.Unmarshal(body, &chartRep); err != nil {
		return nil, fmt.Errorf("unable to unmarshal chart response body: %w", err)
	}

	return &chartRep.ChartReply, nil
}
