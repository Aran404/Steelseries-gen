// package main

// import (
// 	client "PromoGen/Core/Client"
// 	tof "PromoGen/Core/SteelSeries"
// 	utils "PromoGen/Core/Utils"
// 	"fmt"
// )

// var (
// 	Proxies *utils.Iterator
// 	Config  tof.Config
// )

// func init() {
// 	Config = utils.LoadConfig("config.json")

// 	if !Config.Proxyless {
// 		Proxies = utils.NewFromFile("Data/proxies.txt")
// 	}

// 	utils.Clear()
// }

// func createThread(proxy, userAgent, version string, timeout int) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			utils.LogError("Goroutine Panicked, Error: %v", r)
// 		}
// 	}()

// 	tlsClient, err := client.NewTlsClient(proxy, "107", true, timeout)

// 	if err != nil {
// 		utils.LogError("Tls Client: %v", err)
// 		return
// 	}

// 	Instance := tof.Instance{
// 		Client:    tlsClient,
// 		UserAgent: userAgent,
// 	}

// 	err = Instance.GetPage()

// 	if err != nil {
// 		utils.LogError("An error occured: %v", err)
// 		return
// 	}

// 	_, err = Instance.GetNewEmail()

// 	if err != nil {
// 		utils.LogError("An error occured: %v", err)
// 		return
// 	}

// 	err = Instance.SubmitEmail()

// 	if err != nil {
// 		utils.LogError("An error occured: %v", err)
// 		return
// 	}

// 	inbox, err := Instance.WaitForMail(60, true)

// 	if err != nil {
// 		utils.LogError("An error occured: %v", err)
// 		return
// 	}

// 	err = Instance.SubmitCode(fmt.Sprintf("%v", inbox))

// 	if err != nil {
// 		utils.LogError("An error occured: %v", err)
// 		return
// 	}

// 	utils.LogSuccess(Instance.Email + ":" + Instance.Password)
// 	utils.AppendLine("Data/Success.txt", Instance.Email+":"+Instance.Password)
// }

// func main() {
// 	guard := make(chan struct{}, Config.Threads)

// 	for i := 0; i < Config.Iterations*Config.Threads; i++ {
// 		var Proxy string

// 		if Config.Proxyless {
// 			Proxy = ""
// 		} else {
// 			Proxy = Proxies.Next()
// 		}

// 		guard <- struct{}{}
// 		go func() {
// 			createThread(Proxy, Config.Fingerprinting.UserAgent, Config.Fingerprinting.ChromeVersion, Config.Fingerprinting.ClientTimeout)
// 			<-guard
// 		}()
// 	}
// }

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
