package main

import (
	client "PromoGen/Core/Client"
	ss "PromoGen/Core/SteelSeries"
	utils "PromoGen/Core/Utils"
	"fmt"
)

func main() {
	tlsClient, err := client.NewTlsClient("", "103", true, 60)

	if err != nil {
		utils.LogError("%s", err.Error())
		return
	}

	in := ss.Instance{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		Password:  "Helloyoumonkey@1",
		Client:    tlsClient,
	}

	err = in.GetPage()

	if err != nil {
		utils.LogError("%s", err.Error())
		return
	}

	_, err = in.GetNewEmail()

	if err != nil {
		utils.LogError("%s", err.Error())
		return
	}

	err = in.Register()

	if err != nil {
		utils.LogError("%s", err.Error())
		return
	}

	err = in.Login()

	if err != nil {
		utils.LogError("%s", err.Error())
		return
	}

	token, err := in.WaitForMail(60, true)

	if err != nil {
		utils.LogError("%s", err.Error())
		return
	}

	in.VerificationLink = fmt.Sprintf("%v", token)

	fmt.Println(in.VerificationLink)
	fmt.Println(in.Email, in.Password)
	err = in.VerifyEmail()

	if err != nil {
		utils.LogError("%s", err.Error())
		return
	}

	err = in.GetCode()

	if err != nil {
		utils.LogError("%s", err.Error())
		return
	}

}
