package dto

type EmailModel struct {
	EmailDomain string `json:"emailDomain"`
}

type EmailVerifyReponseModel struct {
	HasMX        bool   `json:"hasMX"`
	HasSPF       bool   `json:"hasSPF"`
	SpfRecords   string `json:"spfRecords"`
	HasDMarc     bool   `json:"hasDMarc"`
	DMarcRecords string `json:"dMarcRecords"`
}
