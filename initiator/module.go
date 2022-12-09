package initiator

import (
	"dating/internal/module"
	"dating/internal/module/oauth"
	"dating/platform/logger"
)

type Module struct {
	// TODO implement
	AuthModule module.AuthModule
}

// if the dating app has its own private key for reading tokens
func InitModule(persistence Persistence, privateKeyPath string, platformLayer PlatformLayer, log logger.Logger) Module {
	// keyFile, err := ioutil.ReadFile(privateKeyPath)
	// if err != nil {
	// 	log.Fatal(context.Background(), "failed to read private key", zap.Error(err))
	// }

	// privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyFile)
	// if err != nil {
	// 	log.Fatal(context.Background(), "failed to parse private key", zap.Error(err))
	// }

	return Module{
		AuthModule: oauth.InitOAuth(log),
	}
}
