package healthservice

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/subhroacharjee/appname/utils/test_utils"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestHealthService(t *testing.T) {
	ctx := context.Background()
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	defer testutils.FindTestDbAndRemove("test.db")

	healthService := HealthService{
		Logger: logger,
		DB:     db,
	}

	t.Run("Check", func(t *testing.T) {
		err := healthService.Check(ctx)
		assert.Nil(t, err)
	})
}
