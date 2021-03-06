package habitica

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Tag struct {
	ID   string `json: id`
	Name string `json: name`
}

type ReorderTag struct {
	TagID string `json: tagId`
	To    int    `json: to`
}

type TagResponse struct {
	Success bool `json:"success"`
	Data    *Tag `json:"data,omitempty"`
}

type TagsResponse struct {
	Success bool  `json:"success"`
	Data    []Tag `json:"data,omitempty"`
}

type TagService struct {
	client *HabiticaClient
}

func newTagService(h *HabiticaClient) *TagService {
	return &TagService{
		client: h,
	}
}

func (s *TagService) Create(ctx context.Context, tag *Tag) (*TagResponse, error) {
	req, err := s.client.NewRequest(http.MethodPost, "tags", tag)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return s.getTagResponse(ctx, req)
}

func (s *TagService) Delete(ctx context.Context, id string) (*TagResponse, error) {
	req, err := s.client.NewRequest(http.MethodDelete, fmt.Sprintf("tags/%s", id), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return s.getTagResponse(ctx, req)
}

func (s *TagService) Get(ctx context.Context, id string) (*TagResponse, error) {
	req, err := s.client.NewRequest(http.MethodGet, fmt.Sprintf("tags/%s", id), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return s.getTagResponse(ctx, req)
}

func (s *TagService) List(ctx context.Context) (*TagsResponse, error) {
	req, err := s.client.NewRequest(http.MethodGet, "tags", nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()

	var tagsResp TagsResponse
	err = json.NewDecoder(resp.Body).Decode(&tagsResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response body: %s", err)
	}
	return &tagsResp, err
}

func (s *TagService) Reorder(ctx context.Context, t *ReorderTag) (*TagResponse, error) {
	req, err := s.client.NewRequest(http.MethodPost, "reorder-tags", t)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return s.getTagResponse(ctx, req)
}

func (s *TagService) Update(ctx context.Context, id string, t *Tag) (*TagResponse, error) {
	req, err := s.client.NewRequest(http.MethodPut, fmt.Sprintf("tags/%s", id), t)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %s", err)
	}

	return s.getTagResponse(ctx, req)
}

func (s *TagService) getTagResponse(ctx context.Context, req *http.Request) (*TagResponse, error) {
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}
	defer resp.Body.Close()

	var tagResp TagResponse
	err = json.NewDecoder(resp.Body).Decode(&tagResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response body: %s", err)
	}
	return &tagResp, err
}
