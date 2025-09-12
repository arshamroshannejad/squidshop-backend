package service

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/arshamroshannejad/squidshop-backend/config"
	"github.com/arshamroshannejad/squidshop-backend/internal/domain"
	"github.com/arshamroshannejad/squidshop-backend/internal/entity"
)

type smsServiceImpl struct {
	logger *slog.Logger
	cfg    *config.Config
}

func NewSmsService(logger *slog.Logger, cfg *config.Config) domain.SmsService {
	return &smsServiceImpl{
		logger: logger,
		cfg:    cfg,
	}
}

func (s *smsServiceImpl) Send(msg, phone string) error {
	if s.cfg.App.Debug {
		s.logger.Info("debug mode is enabled, code is: ", msg)
		return nil
	}
	smsReq := entity.SmsRequest{
		Recipient: []string{phone},
		Sender:    s.cfg.App.SmsSender,
		Message:   msg,
	}
	payload, err := json.Marshal(&smsReq)
	if err != nil {
		s.logger.Error("failed to marshal sms request", "error:", err)
		return err
	}
	req, err := http.NewRequest(http.MethodPost, s.cfg.App.SmsService, bytes.NewBuffer(payload))
	if err != nil {
		s.logger.Error("failed to create sms request", "error:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", s.cfg.App.SmsApiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		s.logger.Error("failed to send sms", "error:", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		s.logger.Error("failed to send sms", "status code:", resp.StatusCode)
		return err
	}
	return nil
}
