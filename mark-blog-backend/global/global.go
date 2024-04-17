package global

import (
	"mark-blog-backend/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 全局变量，存储所有的重要元素，方便调用
var(
	Config *config.Config
	DB *gorm.DB
	Log *logrus.Logger
)