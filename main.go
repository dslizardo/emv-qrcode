package main

import (
	"fmt"
	"log"

	"github.com/dongri/emv-qrcode/emv/mpm"
	"github.com/skip2/go-qrcode"
)

func main() {

	// conflicting payment type instapay qr value on there docs
	// 99960300, 99964403
	// 99964403 is being used by Paymaya and UBP
	const (
		ubpCode = "UBPHPHMMXXX"
		paymayaCode = "PAPHPHM1XXX"
		cebCode = "CELRPHM1XXX"
		bpiCode = "BOPIPHMMXXX"

		phQrCode = "com.p2pqrpay"

		qrP2pInstapay = "99964403"

		// for p2p, most likely it will be 6016 (Personal Information)
		// copied this value from UB and Paymaya
		p2pMerchantCategoryCode = "6016"

		phCurrency = "608"
	)

	receiver := "639175606349"


	// MPM Encode
	emvqr := &mpm.EMVQR{}
	emvqr.SetPayloadFormatIndicator("01")
	emvqr.SetPointOfInitiationMethod("11") // 11 is static qrcode
	merchantAccountInfoPH := &mpm.MerchantAccountInformation{}
	merchantAccountInfoPH.SetGloballyUniqueIdentifier(phQrCode)
	merchantAccountInfoPH.AddPaymentNetworkSpecific("01", paymayaCode)
	//merchantAccountInfoPH.AddPaymentNetworkSpecific("02", qrP2pInstapay)
	merchantAccountInfoPH.SetPaymentType("02", qrP2pInstapay)
	// 04 is receiver/merchant account
	merchantAccountInfoPH.AddPaymentNetworkSpecific("04", receiver)
	// O5 is optional, this is mobile number to notify the account
	//merchantAccountInfoPH.AddPaymentNetworkSpecific("05", "09175606349")
	emvqr.AddMerchantAccountInformation("27", merchantAccountInfoPH)

	emvqr.SetMerchantCategoryCode(p2pMerchantCategoryCode)
	emvqr.SetTransactionCurrency(phCurrency)
	//emvqr.SetTransactionAmount("999.123")
	emvqr.SetCountryCode("PH")
	emvqr.SetMerchantName("danielpaymaya")
	emvqr.SetMerchantCity("BocaueBulacan")
	emvqr.SetCRC("AC1D")

	code, err := mpm.Encode(emvqr)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(code)
	err = qrcode.WriteFile(code, qrcode.High, 512, "recreatedanielpaymayaqr.png")
	if err != nil {
		log.Println(err)
	}
}
