package v1

import (
	"bytes"
	"fmt"
)

type TmplList struct {
	ListStatus
	ListPageLinks
	Tmpls []MsgTmpl `json:"messagetemplates"`
}

type MsgTmpl struct {
	Name        string `json:"messageTemplateName"`
	Description string `json:"messageTemplateDescription"`
	Id          string `json:"id"`
	Links       []Link `json:"link"`
}

func (l TmplList) String() string {
	buf := bytes.Buffer{}
	buf.Write([]byte("templates: ["))
	for _, tmpl := range l.Tmpls {
		fmt.Fprintf(&buf, "%s", tmpl)
	}
	fmt.Fprintf(&buf, "\n]\n%s", l.ListPageLinks)
	return buf.String()
}

func (t MsgTmpl) String() string {
	var link string
	if len(t.Links) > 0 {
		link = t.Links[0].Uri
	}
	return fmt.Sprintf("\n"+
		"  %v: {\n"+
		"    id: %s\n"+
		"    description: %s\n"+
		"    link: %s\n"+
		"  }", t.Name, t.Id, t.Description, link)
}

type MsgTmplDetail struct {
	Name        string `json:"messageTemplateName"`
	Description string `json:"messageTemplateDescription"`
	Subject     string `json:"subject"`
	Tags        string `json:"tags"`
	*DLR        `json:"dlr"`
	Notes       string         `json:"notes"`
	Body        string         `json:"body"`
	Type        string         `json:"type"`
	RespTmplId  string         `json:"responseTemplateId"`
	Links       OperationLinks `json:"link"`
	Email       *EmailMsg      `json:"email"`
	Web         *WebMsg        `json:"web"`
	Voice       *VoiceMsg      `json:"voice"`
	Social      *SocialMsg     `json:"social"`
}

func (d *MsgTmplDetail) String() string {
	return fmt.Sprintf("template '%s': {\n"+
		"  description: %s\n"+
		"  subject: %s\n"+
		"  SMS: %s\n"+
		"  ------------------- email message ------------------- \n"+
		"%v"+
		"  ------------------- email message ------------------- \n"+
		"  ------------------- web message ------------------- \n"+
		"%v"+
		"  ------------------- web message -------------------\n"+
		"  ------------------- voice message -------------------\n"+
		"%v"+
		"  ------------------- voice message -------------------\n"+
		"  link: [\n"+
		"%v"+
		"  ]\n"+
		"}",
		d.Name, d.Description, d.Subject, d.Body, d.Email, d.Web, d.Voice, d.Links)
}
