package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/isbik/bank/db/sqlc"
	_ "github.com/lib/pq"
)

func newTestServer(t *testing.T, store db.Store) *Server {

	server := NewServer(store)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
