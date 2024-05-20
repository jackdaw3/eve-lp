package net

import (
	"evelp/log"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type RetryStrategy struct {
	MaxRetries int
	Interval   time.Duration
	Factor     float64
}

type RetryableHttpClient struct {
	Client  *http.Client
	Retryer *RetryStrategy
}

var client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 3,
	},
	Timeout: 100 * time.Second,
}

var retryer = &RetryStrategy{
	MaxRetries: 5,
	Interval:   time.Second,
	Factor:     3,
}

var retryableClient = &RetryableHttpClient{
	Client:  client,
	Retryer: retryer,
}

func Get(req string) ([]byte, http.Header, error) {
	retries := 0
	interval := retryableClient.Retryer.Interval

	for {
		resp, err := retryableClient.Client.Get(req)
		if err != nil {
			return nil, nil, errors.WithStack(err)
		}
		defer resp.Body.Close()

		switch {
		case resp.StatusCode >= 200 && resp.StatusCode < 300:
			content, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, nil, errors.WithMessagef(err, "failed to read body for req %s", req)
			}
			return content, resp.Header, nil

		case resp.StatusCode >= 400 && resp.StatusCode < 500:
			content, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, nil, errors.WithMessagef(err, "failed to read error meesgae for req %s response code %d", req, resp.StatusCode)
			}

			return nil, nil, errors.Errorf("req %s client error: %s", req, content)

		case resp.StatusCode >= 500 && resp.StatusCode < 600:
			if retries >= retryableClient.Retryer.MaxRetries {
				return nil, nil, errors.Errorf("req %s max retries reached: %d", req, retries)
			}
			log.Debugf("req %s retrying request after %v, retry count: %d\n", req, interval, retries)
			time.Sleep(interval)
			retries++
			interval = time.Duration(float64(interval) * retryableClient.Retryer.Factor)
			continue

		default:
			return nil, nil, errors.Errorf("req %s unknow response code: %d", req, resp.StatusCode)
		}
	}
}
