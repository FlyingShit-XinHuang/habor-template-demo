package v1

import (
	"bytes"
	"fmt"
)

type ListStatus struct {
	Status string `json:"status"`
}

type Link struct {
	Uri    string `json:"uri"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type ListPageLinks struct {
	Links []Link `json:"link"`
}

func (l ListPageLinks) String() string {
	buf := bytes.NewBuffer([]byte{})
	i := 0
	for _, link := range l.Links {
		if i > 0 {
			fmt.Fprintln(buf)
		}
		fmt.Fprintf(buf, "%s page: %s", link.Rel, link.Uri)
		i++
	}
	return buf.String()
}

type PaginationOptions struct {
	Offset int
	Limit  int
}

type DLR struct {
	Period       string `json:"period"`
	Rule         string `json:"rule"`
	Type         string `json:"type"`
	PublishToWeb bool   `json:"publishToWeb"`
	ExpiryDay    int    `json:"expiryDay"`
	ExpiryHour   int    `json:"expiryHour"`
	ExpiryMin    int    `json:"expiryMin"`
	FeedIds      string `json:"feedIds"`
	Bool         bool   `json:"bool"`
}

type EmailMsg struct {
	MsgCommon
	Footer string `json:"footer"`
}

func (e *EmailMsg) String() string {
	if nil == e {
		return "    null"
	}
	return fmt.Sprintf("    body: %s\n    footer: %s\n", e.Body, e.Footer)
}

type WebMsg struct {
	MsgCommon
}

func (w *WebMsg) String() string {
	if nil == w {
		return "    null"
	}
	return fmt.Sprintf("    body: %s\n", w.Body)
}

type VoiceMsg struct {
	MsgCommon
	Header string `json:"header"`
}

func (v *VoiceMsg) String() string {
	if nil == v {
		return "    null"
	}
	return fmt.Sprintf("    header: %s\n    body: %s\n", v.Header, v.Body)
}

type SocialMsg struct {
	List []SocialMsgDetail `json:"social"`
}

type SocialMsgDetail struct {
	Id   string `json:"id"`
	Body string `json:"body"`
}

type MsgCommon struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

type OperationLinks []Link

func (l OperationLinks) String() string {
	buf := bytes.NewBuffer([]byte{})
	i := 0
	for _, link := range l {
		if "self" == link.Rel {
			continue
		}
		fmt.Fprintf(buf, "    %s: (%s)%s\n", link.Rel, link.Method, link.Uri)
		i++
	}
	return buf.String()
}
