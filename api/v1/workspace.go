package v1

import (
	"bytes"
	"fmt"
)

type WorkspaceList struct {
	ListStatus
	ListPageLinks
	Workspaces []Workspace `json:"workspaces"`
}

type Workspace struct {
	Name              string `json:"projectName"`
	Number            string `json:"projectNumber"`
	Status            string `json:"status"`
	BillingCostCentre string `json:"billingcostcentre"`
	Id                string `json:"id"`
	Links             []Link `json:"link"`
}

func (l WorkspaceList) String() string {
	buf := bytes.Buffer{}
	buf.Write([]byte("workspaces: ["))
	for _, s := range l.Workspaces {
		fmt.Fprintf(&buf, "%s", s)
	}
	fmt.Fprintf(&buf, "\n]\n%s", l.ListPageLinks)
	return buf.String()
}

func (w Workspace) String() string {
	var link string
	if len(w.Links) > 0 {
		link = w.Links[0].Uri
	}
	return fmt.Sprintf("\n"+
		"  %v: {\n"+
		"    id: %s\n"+
		"    status: %s\n"+
		"    link: %s\n"+
		"  }", w.Name, w.Id, w.Status, link)
}
