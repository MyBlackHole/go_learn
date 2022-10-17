package main

import (
	"encoding/json"
	"fmt"
)

type NegativeKeywordPost struct {
	KeywordID   *int64  `json:"keywordId,omitempty"`
	CampaignID  *int64  `json:"campaignId,omitempty"`
	AdGroupID   *int64  `json:"adGroupId,omitempty"`
	State       *string `json:"state,omitempty"`
	KeywordText *string `json:"keywordText,omitempty"`
	MatchType   *string `json:"matchType,omitempty"`
}

type DATA interface {
}

func (this *NegativeKeywordPost) String() string {
	b, _ := json.Marshal(this)
	return string(b)
}

func main() {
	n := new(NegativeKeywordPost)
	i := int64(1234)
	n.AdGroupID = &i
	fmt.Printf("%+v", n)
}
