package domain

import (
	"email-verify-go/dto"
	"fmt"
	"net"
	"strings"
)

type IVerifyEmail interface {
	IsEmailValid(string) dto.EmailVerifyReponseModel
}

type VerifyEmailImpl struct {
}

func (vei VerifyEmailImpl) IsEmailValid(email string) dto.EmailVerifyReponseModel {
	hasMX := vei.lookupMXRecords(email)
	hasSPF, SPFRecords := vei.lookupSPFRecords(email)
	hasDMarc, dMarcRecords := vei.lookupDMarcRecordss(email)

	return dto.EmailVerifyReponseModel{
		HasMX:        hasMX,
		HasSPF:       hasSPF,
		SpfRecords:   getDummy(SPFRecords),
		HasDMarc:     hasDMarc,
		DMarcRecords: getDummy(dMarcRecords),
	}

}

func getDummy(val *string) string {
	if val == nil {
		return ""
	} else {
		return *val
	}
}

func (vei VerifyEmailImpl) lookupMXRecords(email string) bool {
	mxRecord, err := net.LookupMX(email)
	if err != nil {
		fmt.Printf("Error in LookupMX %v\n", err)
		return false
	}

	if len(mxRecord) > 0 {
		return true
	}

	return false
}

func (vei VerifyEmailImpl) lookupSPFRecords(email string) (bool, *string) {
	txtRecords, err := net.LookupTXT(email)
	if err != nil {
		fmt.Printf("Error in LookupSPF %v\n", err)
		return false, nil
	}

	fmt.Printf("SPF: %v\n", txtRecords)

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			return true, &record
		}
	}

	return false, nil
}

func (vei VerifyEmailImpl) lookupDMarcRecordss(email string) (bool, *string) {
	txtRecords, err := net.LookupTXT("_dmarc." + email)
	if err != nil {
		fmt.Printf("Error in LookupDMarc %v\n", err)
		return false, nil
	}

	fmt.Printf("DMarc: %v\n", txtRecords)

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			return true, &record
		}
	}

	return false, nil
}

func NewVerifyEmailImpl() IVerifyEmail {
	return VerifyEmailImpl{}
}
