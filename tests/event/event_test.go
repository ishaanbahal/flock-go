package event

import (
	"fmt"
	"testing"

	"reflect"

	handler "../../src/event"
)

const (
	appInstallStr   = "{\"name\":\"app.install\",\"data\":\"message\"}"
	appUninstallStr = "{\"name\":\"app.uninstall\",\"data\":\"message\"}"
)

func TestHandler(t *testing.T) {
	eh := handler.EventHandler{}
	eh.SetOnAppInstallHandler(func(data map[string]interface{}) {
		fmt.Println("OnInstall ran")
		fmt.Println(data)
	})
	err := eh.HandleEvent([]byte(appInstallStr))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}

	eh.SetOnAppUninstallHandler(func(data map[string]interface{}) {
		fmt.Println("OnUnstall ran")
		fmt.Println(data)
	})

	err = eh.HandleEvent([]byte(appUninstallStr))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}

	err = eh.HandleEvent([]byte(appInstallStr))
	if err != nil {
		t.Errorf("Error occured: %s", err)
	}
}

func assertEquals(t *testing.T, actual, required map[string]interface{}) {
	if reflect.DeepEqual(actual, required) {
		return
	}
	t.Fatalf("Actual interface doesn't match required")
}
