package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lushenle/plam/pkg/db"
	"github.com/lushenle/plam/pkg/log"
	"github.com/lushenle/plam/pkg/token"
	"github.com/lushenle/plam/pkg/util"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		Server: util.Server{
			TokenSymmetricKey:   util.RandomString(32),
			AccessTokenDuration: time.Minute,
		},
	}

	tokenMaker, err := token.NewPasetoMaker(config.Server.TokenSymmetricKey)
	require.NoError(t, err)

	plugin := log.NewStderrPlugin(zapcore.DebugLevel)
	logger := log.NewLogger(plugin)

	server := NewServer(config, WithStore(store), WithLogger(logger), WithTokenMaker(tokenMaker))

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
