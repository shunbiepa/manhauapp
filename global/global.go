package global

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"manhauapp/config"
	"manhauapp/middleware/cache"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *cache.RCache
	GVA_CONFIG *config.Specification
	GVA_LOG    *zap.Logger
)
