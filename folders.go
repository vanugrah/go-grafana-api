package gapi

import (
	"encoding/json"
	"fmt"
	"time"

	"io/ioutil"

	"github.com/pkg/errors"
)

type Folder struct {
	Id        int       `json:"id"`
	Uid       string    `json:"uid"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	HasAcl    bool      `json:"hasAcl"`
	CanSave   bool      `json:"canSave"`
	CanEdit   bool      `json:"canEdit"`
	CanAdmin  bool      `json:"canAdmin"`
	CreatedBy string    `json:"createdBy"`
	Created   time.Time `json:"created"`
	UpdatedBy string    `json:"updatedBy"`
	Updated   time.Time `json:"updated"`
	Version   int       `json:"version"`
}

type FolderCreateOpts struct {
	Title string `json:"title"`
	Id    int    `json:"id"`
}

func (c *Client) GetAllFolders() ([]Folder, error) {
	folders := make([]Folder, 0)
	req, err := c.newRequest("GET", "/api/folders/", nil, nil)
	if err != nil {
		return folders, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return folders, errors.Wrap(err, "Unable to perform HTTP request")
	}

	if resp.StatusCode != 200 {
		var gmsg GrafanaErrorMessage
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&gmsg)
		return folders, &GrafanaError{resp.StatusCode, gmsg.Message}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return folders, err
	}

	err = json.Unmarshal(data, &folders)
	if err != nil {
		return folders, err
	}
	return folders, err
}

// func (c *Client) CreateFolder(model map[string]interface{}) (*FolderCreateResponse, error) {
// 	wrapper := map[string]interface{}{
// 		"title": model["title"],
// 		"uid":   model["uid"],
// 	}
// 	data, err := json.Marshal(wrapper)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Failed to marshall folder JSON")
// 	}
// 	req, err := c.newRequest("POST", "/api/folders", nil, bytes.NewBuffer(data))
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp, err := c.Do(req)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Unable to perform HTTP request")
// 	}

// 	if resp.StatusCode != 200 {
// 		var gmsg GrafanaErrorMessage
// 		dec := json.NewDecoder(resp.Body)
// 		dec.Decode(&gmsg)
// 		return nil, fmt.Errorf("Request to Grafana returned %+v status code with the following message: %+v", resp.StatusCode, gmsg.Message)
// 	}

// 	data, err = ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result := &FolderCreateResponse{}
// 	err = json.Unmarshal(data, &result)
// 	return result, err
// }

func (c *Client) GetFolderByUID(uid string) (*Folder, error) {
	path := fmt.Sprintf("/api/folders/%s", uid)
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
		return nil, &GrafanaError{resp.StatusCode, gmsg.Message}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &Folder{}
	err = json.Unmarshal(data, &result)
	return result, err
}
