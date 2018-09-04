package gapi

import (
	"bytes"
	"encoding/json"

	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)


type FolderCreateResponse struct {
	Id        int    `json:"id"`
	uid       string `json:"uid"`
	title     string `json:"title"`
	url       string `json:"url"`
	hasAcl    bool   `json:"hasAcl"`
	canSave   bool   `json:"canSave"`
	canEdit   bool   `json:"canEdit"`
	canAdmin  bool   `json:"canAdmin"`
	createdBy string `json:"createdBy"`
	created   string `json:"created"`
	updatedBy string `json:"updatedBy"`
	updated	  string `json:"updated"`
	version   int    `json:"version"`
}

type Folder struct {
	Title string `json:"title"`
	Id    int    `json:"id"`
}

func (c *Client) CreateFolder(model map[string]interface{}) (*FolderCreateResponse, error) {
	wrapper := map[string]interface{}{
		"title": model["title"],
		"uid": 	 model["uid"],
	}
	data, err := json.Marshal(wrapper)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to marshall folder JSON")
	}
	req, err := c.newRequest("POST", "/api/folders", nil, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to perform HTTP request")
	}

	if resp.StatusCode != 200 {
		var gmsg GrafanaErrorMessage
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&gmsg)
		return nil, fmt.Errorf("Request to Grafana returned %+v status code with the following message: %+v", resp.StatusCode, gmsg.Message)
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &FolderCreateResponse{}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) GetFolderByUID(slug string) (*Folder, error) {
	path := fmt.Sprintf("/api/folders/%s", slug)
	req, err := c.newRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var gmsg GrafanaErrorMessage
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&gmsg)
		errMsg := fmt.Sprintf("Request to Grafana returned %+v status code with the following message: %+v", resp.StatusCode, gmsg.Message)

		// TODO: Give support to other custom errors
		switch resp.StatusCode {
		case 404:
			return nil, NewErrNotFound(errMsg)
		default:
			return nil, fmt.Errorf(errMsg)
		}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &Folder{}
	err = json.Unmarshal(data, &result)
	return result, err
}

