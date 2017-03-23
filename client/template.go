package client

import (
	"habor-template-demo/api/v1"
	"fmt"
)

type TemplateClient interface {
	List(opts *v1.PaginationOptions) (*v1.TmplList, error)
	Get(id string) (*v1.MsgTmplDetail, error)
	CreateFromFile(file string) error
	UpdateFromFile(id, file string) error
	Delete(id string) error
	SendMsg(id, to string) (string, error)
}

type tmplClient struct {
	client    *Client
	workspace string
}

func NewTemplateClient(apiKey, user, password, workspace string) TemplateClient {
	return &tmplClient{
		client:    NewClient(apiKey, user, password),
		workspace: workspace,
	}
}

func (c *tmplClient) List(opts *v1.PaginationOptions) (*v1.TmplList, error) {
	list := &v1.TmplList{}
	err := c.client.Get().
		Workspace(c.workspace).
		Resource(ResourceTemplate).
		Pagination(opts).
		SetAccept().
		Do().
		Into(list)
	return list, err
}

func (c *tmplClient) Get(id string) (*v1.MsgTmplDetail, error) {
	tmpl := &v1.MsgTmplDetail{}
	err := c.client.Get().
		Workspace(c.workspace).
		Resource(ResourceTemplate).
		ID(id).
		SetAccept().
		Do().
		Into(tmpl)
	return tmpl, err
}

func (c *tmplClient) CreateFromFile(file string) error {
	return c.client.Post().
		Workspace(c.workspace).
		Resource(ResourceTemplate).
		BodyFromFile(file).
		SetContentType().
		Do().
		Err()
}

func (c *tmplClient) UpdateFromFile(id, file string) error {
	return c.client.Put().
		Workspace(c.workspace).
		Resource(ResourceTemplate).
		ID(id).
		BodyFromFile(file).
		SetContentType().
		Do().
		Err()
}

func (c *tmplClient) Delete(id string) error {
	return c.client.Delete().
		Workspace(c.workspace).
		Resource(ResourceTemplate).
		ID(id).
		Do().
		Err()
}

func (c *tmplClient) SendMsg(id, to string) (string, error) {
	location := ""
	body := fmt.Sprintf(`{"to":"%s", "messageTemplateId":"%s"}`, to, id)
	err := c.client.Post().
		Workspace(c.workspace).
		Resource(ResourceMessage).
		Body([]byte(body)).
		SetContentType().
		Do().
		ExtractLocation(&location).
		Err()
	return location, err
}
