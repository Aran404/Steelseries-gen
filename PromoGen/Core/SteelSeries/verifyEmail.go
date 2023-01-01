package steelseries

import (
	"path/filepath"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func (in *Instance) VerifyEmail() error {
	extPath, _ := filepath.Abs("./ex.crx")

	u := launcher.New().
		// Must use abs path for an extension
		Set("load-extension", extPath).
		// Headless mode doesn't support extension yet.
		// Reason: https://bugs.chromium.org/p/chromium/issues/detail?id=706008#c5
		// You can use XVFB to get rid of it: https://github.com/go-rod/rod/blob/master/lib/examples/launch-managed/main.go
		Headless(false).
		MustLaunch()

	page := rod.New().ControlURL(u).MustConnect().MustPage(in.VerificationLink)

	time.Sleep(time.Second * 40)
	page.MustScreenshot("")
	page.MustClose()

	return nil
}
