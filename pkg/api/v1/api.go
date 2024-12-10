package v1

type User struct {
	Login    string `json:"login"`
	Password string `json:"pswd"`
}

type RespError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type RespLogin struct {
	Login string `json:"login"`
}

type RespToken struct {
	Token string `json:"token"`
}

type UploadReq struct {
	Meta
}

type Meta struct {
	Name   string   `json:"name"`
	Token  string   `json:"token"`
	Mime   string   `json:"mime"`
	Grant  []string `json:"grant"`
	File   bool     `json:"file"`
	Public bool     `json:"public"`
}

func (m *Meta) IsValid() bool {
	if m.Name == "" || m.Mime == "" || m.Token == "" {
		return false
	}

	return true
}

type UploadResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	File string `json:"file"`
}

type GetDocumentsResp struct {
	DataDocuments DataDocuments `json:"data"`
}

type Document struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Mime    string   `json:"mime"`
	File    bool     `json:"file"`
	Public  bool     `json:"public"`
	Created string   `json:"created"`
	Grant   []string `json:"grant"`
}

type DataDocuments struct {
	Docs []Document `json:"docs"`
}
