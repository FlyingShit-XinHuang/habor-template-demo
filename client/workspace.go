package client

import "habor-template-demo/api/v1"

type WorkspaceClient interface {
	List(opts *v1.PaginationOptions) (*v1.WorkspaceList, error)
}

type workspaceClient struct {
	client *Client
}

func NewWorkspaceClient(apiKey, user, password string) WorkspaceClient {
	return &workspaceClient{
		client: NewClient(apiKey, user, password),
	}
}

func (c *workspaceClient) List(opts *v1.PaginationOptions) (*v1.WorkspaceList, error) {
	list := &v1.WorkspaceList{}
	err := c.client.Get().
		Resource(ResourceWorkspace).
		Pagination(opts).
		SetAccept().
		Do().
		Into(list)
	return list, err
}
