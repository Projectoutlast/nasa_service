package test_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/Projectoutlast/space_service/auth_service/internal/app"
	"github.com/Projectoutlast/space_service/auth_service/internal/config"
	"github.com/Projectoutlast/space_service/auth_service/internal/logging"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	pb "github.com/Projectoutlast/nasa_proto/gen"
)

func TestMainStart(t *testing.T) {
	os.Chdir("../..")
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}

	_ = logging.New("testing", os.Stdout)
	_ = logging.New("production", os.Stdout)

	newLogger := logging.New(cfg.Environment, os.Stdout)

	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	createTable(t, db)

	application := app.New(newLogger, cfg, db)

	go application.GRPCServer.MustRun()

	connString := fmt.Sprintf("localhost:%v", cfg.AuthConfig.Port)
	conn, err := grpc.NewClient(connString, grpc.WithInsecure())
	require.NoError(t, err)

	authClient := pb.NewAuthClient(conn)

	resReg, err := authClient.Registration(
		context.Background(),
		&pb.RegistrationRequest{
			Email:    "test@example.com",
			Password: "12345678",
		},
	)
	require.NoError(t, err)
	require.Equal(t, int64(1), resReg.GetUserId())

	resLog, err := authClient.Login(
		context.Background(),
		&pb.LoginRequest{
			Email:    "test@example.com",
			Password: "12345678",
		},
	)
	require.NoError(t, err)
	require.NotNil(t, resLog.GetToken())

	_, err = authClient.Registration(
		context.Background(),
		&pb.RegistrationRequest{
			Email:    "",
			Password: "12345678",
		},
	)
	require.Error(t, err)

	_, err = authClient.Login(
		context.Background(),
		&pb.LoginRequest{
			Email:    "",
			Password: "12345678",
		},
	)
	require.Error(t, err)

	_, err = authClient.Login(
		context.Background(),
		&pb.LoginRequest{
			Email:    "test@example",
			Password: "asdfghjgfdghjkgf",
		},
	)
	require.Error(t, err)

	_, err = authClient.Login(
		context.Background(),
		&pb.LoginRequest{
			Email:    "thisisnotanemail",
			Password: "12345678",
		},
	)
	require.Error(t, err)

	_, err = authClient.Registration(
		context.Background(),
		&pb.RegistrationRequest{
			Email:    "test@example.com",
			Password: "12345678",
		},
	)
	require.Error(t, err)
}
