package tests

import (
	"fmt"
	"testing"

	handler ".."
)

func TestHandler(t *testing.T) {
	eh := handler.EventHandler{}
	eh.SetOnAppInstallHandler(func(data map[string]interface{}) {
		fmt.Println("OnInstall ran")
		fmt.Println(data)
	})
	err := eh.HandleEvent([]byte("{\"name\":\"app.install\",\"data\":\"message\"}"))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}

	eh.SetOnAppUninstallHandler(func(data map[string]interface{}) {
		fmt.Println("OnUnstall ran")
		fmt.Println(data)
	})

	err = eh.HandleEvent([]byte("{\"name\":\"app.uninstall\",\"data\":\"message\"}"))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}

	err = eh.HandleEvent([]byte("{\"name\":\"app.install\",\"data\":\"message\"}"))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
}
