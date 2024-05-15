package engine

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"context"
	"errors"

	"explorer_webarchiv/internal/config"
	"explorer_webarchiv/internal/types"
	"explorer_webarchiv/internal/utils"
)

const API_URL = "https://web.archive.org/cdx/search/cdx?url=*."

const MAX_CONNECTION_ATTEMPTS = 2

func GetWebHistory(ctx context.Context, targetDomain string, timeStamp string) ([]string, error) {
	query_url := API_URL + targetDomain

	cfg, ok := ctx.Value(types.ConfigObject).(*config.Config)
	if !ok {
		return nil, fmt.Errorf("'%s' not found in context", types.ConfigObject)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", query_url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	content := string(body)
	lines := strings.Split(string(content), "\n")

	results := make([]string, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		data := strings.Split(line, " ")

		if len(data) != 7 || data[4] != "200" || utils.IsValueExists(data[3], cfg.Rules.BlackListHttpTypes) {
			continue
		}

		if strings.Contains(data[0], ")/") == true {
			savedpage := strings.Split(data[0], ")/")[1]
			url := targetDomain + "/" + savedpage
			timestamp := string(data[1])

			if strings.HasPrefix(timestamp, timeStamp) {
				wayback_url := timestamp + "/" + url
				results = append(results, wayback_url)
			}
		}
	}

	return results, nil
}

func GetContent(ctx context.Context, url string) (string, error) {
	for i := 0; i <= MAX_CONNECTION_ATTEMPTS; i++ {
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return "", fmt.Errorf("error creating request: %w", err)
		}

		response, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", fmt.Errorf("error making request: %w", err)
		}

		if response.StatusCode != http.StatusOK {
			continue
		}

		data, err := io.ReadAll(response.Body)
		if err != nil {
			return "", fmt.Errorf("%w", err)
		}

		result := string(data)
		response.Body.Close()

		return result, nil
	}

	return "", errors.New("without any response data")
}
