package helper

import (
	"fmt"
	"strings"

	"github.com/arshamroshannejad/squidshop-backend/config"
)

func BuildMediaURL(cfg *config.Config, mediaPath *string) *string {
	if cfg == nil || cfg.S3.Domain == "" || mediaPath == nil || *mediaPath == "" {
		return nil
	}
	domain := strings.TrimRight(cfg.S3.Domain, "/")
	path := strings.TrimLeft(*mediaPath, "/")
	url := fmt.Sprintf("%s/%s", domain, path)
	return &url
}
