package configs

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type jwtConf struct {
	JwtPublicPath  string
	JwtPrivatePath string
	JwtPassphrase  string
	JwtALGO        string
	JwtTtl         time.Duration
	JwtRefreshTtl  time.Duration
}

var JwtConfInstance jwtConf
var JwtConfOnce sync.Once

func GetJwtConf() *jwtConf {
	JwtConfOnce.Do(func() {
		JwtConfInstance.JwtPublicPath = os.Getenv("JWT_PUBLIC_KEY")
		JwtConfInstance.JwtPrivatePath = os.Getenv("JWT_PRIVATE_KEY")
		JwtConfInstance.JwtPassphrase = os.Getenv("JWT_PASSPHRASE")
		JwtConfInstance.JwtALGO = os.Getenv("JWT_ALGO")
		ttlInt64, err := strconv.ParseInt(os.Getenv("JWT_TTL"), 10, 64)
		if err != nil {
			return
		}
		JwtConfInstance.JwtTtl = time.Duration(ttlInt64)

		ttlInt64, err = strconv.ParseInt(os.Getenv("JWT_REFRESH_TTL"), 10, 64)
		if err != nil {
			return
		}
		JwtConfInstance.JwtRefreshTtl = time.Duration(ttlInt64)
		fmt.Println(JwtConfInstance)
	})
	return &JwtConfInstance
}
