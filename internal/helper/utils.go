package helper

import (
	"fmt"
	"strings"

	"github.com/arshamroshannejad/squidshop-backend/config"
)

func BuildMediaURL(cfg *config.Config, mediaPath *string) *string {
	url := fmt.Sprintf("%s/%s", strings.TrimRight(cfg.S3.Domain, "/"), strings.TrimLeft(*mediaPath, "/"))
	return &url
}
