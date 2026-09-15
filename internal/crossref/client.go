package crossref

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	userAgent  string
}

func NewClient(userMail string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		userAgent: "ABNTConverter/1.0 (mailto:" + userMail + ")",
		baseURL:   "https://api.crossref.org/works/",
	}
}

func extractDoi(raw string) (string, error) {
	start := strings.Index(raw, "10.")

	if start < 0 {
		return "", fmt.Errorf("nenhum padrão do DOI foi encontrado: %s", raw)
	}
	raw = raw[start:]
	//O motivo dessa substituição é que na api do crossref a busca é feita substituindo o / por %2F
	raw = strings.Replace(raw, "/", "%2F", 1)
	return raw, nil

}

func (c *Client) GetMetadata(doi string) (*ArticleMetadata, error) {

	cleanDoi, err := extractDoi(doi)

	if err != nil {
		return nil, err
	}

	searchUrl := c.baseURL + cleanDoi
	req, err := http.NewRequest(http.MethodGet, searchUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar a requisição http: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("falha ao fazer a requisição a API do Crossref: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ocorreu algum erro no servidor do crossref: %s", resp.Status)
	}

	var apiResponse response

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("falha ao decodificar a resposta do Crossref: %w", err)
	}

	return &apiResponse.Message, nil

}
