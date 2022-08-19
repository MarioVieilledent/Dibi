package httpService

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

func ServeWebApp() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	fmt.Println("App served on port", port)

	// openBrowser("http://" + Url + ":" + Port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error serving static files")
	}
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		fmt.Println("Error opening browser :", err)
	}
}
