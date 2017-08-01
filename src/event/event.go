package event

import (
	"encoding/json"
)

const (
	APP_INSTALL                   = "app.install"
	APP_UNINSTALL                 = "app.uninstall"
	CHAT_GENERATE_URL_PREVIEW     = "chat.generateUrlPreview"
	CHAT_RECEIVE_MESSAGE          = "chat.receiveMessage"
	CLIENT_FLOCKML_ACTION         = "client.flockmlAction"
	CLIENT_MESSAGE_ACTION         = "client.messageAction"
	CLIENT_OPEN_ATTACHMENT_WIDGET = "client.openAttachmentWidget"
	CLIENT_PRESS_BUTTON           = "client.pressButton"
	CLIENT_SLASH_COMMAND          = "client.slashCommand"
	CLIENT_WIDGET_ACTION          = "client.widgetAction"
)

type EventHandler struct {
	_eventMap map[string]func(map[string]interface{})
}

func (eh *EventHandler) HandleEvent(event []byte) error {
	var data map[string]interface{}
	err := json.Unmarshal(event, &data)
	if err != nil {
		return err
	}
	eh._handle(data)
	return nil
}

func (eh *EventHandler) _handle(data map[string]interface{}) {
	methodName := data["name"].(string)
	if val, ok := eh._eventMap[methodName]; ok {
		if val != nil {
			val(data)
		}
	}
}

func (eh *EventHandler) _setEventHandler(event string, callback func(map[string]interface{})) {
	if eh._eventMap == nil {
		eh._eventMap = make(map[string]func(map[string]interface{}))
	}
	eh._eventMap[event] = callback
}

func (eh *EventHandler) SetOnAppInstallHandler(callback func(map[string]interface{})) {
	eh._setEventHandler(APP_INSTALL, callback)
}

func (eh *EventHandler) SetOnAppUninstallHandler(callback func(map[string]interface{})) {
	eh._setEventHandler(APP_UNINSTALL, callback)
}

func (eh *EventHandler) SetOnChatGenerateURLPreviewHandler(callback func(map[string]interface{})) {
	eh._setEventHandler(CHAT_GENERATE_URL_PREVIEW, callback)
}

func (eh *EventHandler) SetOnChatReceiveMessageHandler(callback func(map[string]interface{})) {
	eh._setEventHandler(CHAT_RECEIVE_MESSAGE, callback)
}

func (eh *EventHandler) SetOnClientFlockmlActionHandler(callback func(map[string]interface{})) {
	eh._setEventHandler(CLIENT_FLOCKML_ACTION, callback)
}

func (eh *EventHandler) SetOnClientMessageActionHandler(callback func(map[string]interface{})) {
	eh._setEventHandler(CLIENT_MESSAGE_ACTION, callback)
}

func (eh *EventHandler) SetOnClientOpenAttachmentWidgetHandler(callback func(map[string]interface{})) {
	eh._setEventHandler(CLIENT_OPEN_ATTACHMENT_WIDGET, callback)
}

func (eh *EventHandler) SetOnClientPressButtonHandler(callback func(map[string]interface{})) {
	eh._setEventHandler(CLIENT_PRESS_BUTTON, callback)
}

func (eh *EventHandler) SetOnClientSlashCommandHandler(callback func(map[string]interface{})) {
	eh._setEventHandler(CLIENT_SLASH_COMMAND, callback)
}

func (eh *EventHandler) SetOnClientWidgetActionHandler(callback func(map[string]interface{})) {
	eh._setEventHandler(CLIENT_WIDGET_ACTION, callback)
}
