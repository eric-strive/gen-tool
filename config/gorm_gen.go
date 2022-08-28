package config

var GormDefaultImportPackage = `
	"fmt"
	"time"
	"github.com/eric-strive/gen-tool/internal/repository/mysql"
	"gitlab.intsig.net/cs-server2/services/purchase/app_recharge/errno"
	"gorm.io/gorm"
	"context"
`

var GenApiPath = "./gen_file/api"
var GenGormPath = "./gen_file/gorm"
