// Code generated by goctl. DO NOT EDIT!
// Source: transform.proto

package transformer

import "errors"

var errJsonConvert = errors.New("json convert error")

type (
	ExpandReq struct {
		Shorten string `json:"shorten,omitempty"`
	}

	ExpandResp struct {
		Url string `json:"url,omitempty"`
	}

	ShortenReq struct {
		Url string `json:"url,omitempty"`
	}

	ShortenResp struct {
		Shorten string `json:"shorten,omitempty"`
	}
)
