package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestAppConfig_ChatUser(t *testing.T) {
	var appConfig AppConfig
	yamlFile, err := os.ReadFile("../../test/config.yaml")

	if err != nil {
		t.Error(err)
	}

	err = yaml.Unmarshal(yamlFile, &appConfig)
	if err != nil {
		t.Error(err)
	}
	NewAppConfig(&appConfig)

	p := ChatParticipant{
		RoomSid: "RM_test01",
		RoomId:  "room01",
		Name:    "Test",
		UserSid: "PN_test001",
		UserId:  "test001",
	}

	appCnf.AddChatUser("room01", p)
	pp := appCnf.GetChatParticipants("room01")

	if pp[p.UserId].UserSid != p.UserSid {
		t.Errorf("Expected UserSid %s didn't match", p.UserId)
	}

	appCnf.RemoveChatParticipant("room01", p.UserId)
	pp = appCnf.GetChatParticipants("room01")
	if pp[p.UserId].UserId == p.UserSid {
		t.Errorf("Expected UserSid %s shouldn't found", p.UserId)
	}

	appCnf.DeleteChatRoom("room01")
	pp = appCnf.GetChatParticipants("room01")

	if pp != nil {
		t.Error("Expected nil")
	}
}
