package natsmodel

import (
	"fmt"
	"github.com/mynaparrot/plugnmeet-protocol/auth"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	natsservice "github.com/mynaparrot/plugnmeet-server/pkg/services/nats"
	log "github.com/sirupsen/logrus"
	"time"
)

func (m *NatsModel) RenewPNMToken(roomId, userId, token string) {
	token, err := m.authModel.RenewPNMToken(token)
	if err != nil {
		log.Errorln(err)
		return
	}

	err = m.natsService.BroadcastSystemEventToRoom(plugnmeet.NatsMsgServerToClientEvents_RESP_RENEW_PNM_TOKEN, roomId, token, &userId)
	if err != nil {
		log.Errorln(err)
	}
}

func (m *NatsModel) GenerateLivekitToken(roomId string, userInfo *plugnmeet.NatsKvUserInfo) (string, error) {
	c := &plugnmeet.PlugNmeetTokenClaims{
		RoomId:  roomId,
		Name:    userInfo.Name,
		UserId:  userInfo.UserId,
		IsAdmin: userInfo.IsAdmin,
	}

	return auth.GenerateLivekitAccessToken(m.app.LivekitInfo.ApiKey, m.app.LivekitInfo.Secret, m.app.LivekitInfo.TokenValidity, c, userInfo.Metadata)
}

func (m *NatsModel) HandleClientPing(userId string) {
	err := m.natsService.UpdateUserKeyValue(userId, natsservice.UserLastPingAt, fmt.Sprintf("%d", time.Now().UnixMilli()))
	if err != nil {
		log.Errorln(err)
	}
}
