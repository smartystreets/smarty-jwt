package jwt

import "errors"

type ClaimCallback func(claims map[string]interface{}, data interface{})

type headers struct {
	Algorithm string `json:"alg,omitempty"`
	KeyID     string `json:"kid,omitempty"`
	Type      string `json:"typ,omitempty"`
}

type Expiration interface {
	SetExpiration(int64)
}

type Issuer interface {
	SetIssuer(string)
}

type Audience interface {
	SetAudience(...string)
}

var (
	SegmentCountErr            = errors.New("a JWT must have three segments separated by period characters")
	MalformedHeaderErr         = errors.New("the header is malformed")
	MalformedHeaderContentErr  = errors.New("the header content is malformed")
	MalformedPayloadContentErr = errors.New("the payload content is malformed")
	MalformedSignatureErr      = errors.New("the signature is malformed")
	MissingKIDErr              = errors.New("kid is required")
)
