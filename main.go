package main

import (
	"context"
	"database/sql"

	usecases "palm_code_be/src/app/usecases"

	"palm_code_be/src/infra/config"

	postgres "palm_code_be/src/infra/persistence/postgres"

	fsInteg "palm_code_be/src/infra/integration/firestorage"
	mediaRepo "palm_code_be/src/infra/persistence/postgres/media"
	pageRepo "palm_code_be/src/infra/persistence/postgres/pages"
	teamRepo "palm_code_be/src/infra/persistence/postgres/team"
	userRepo "palm_code_be/src/infra/persistence/postgres/user"

	rest "palm_code_be/src/interface"

	ms_log "palm_code_be/src/infra/log"

	mediaUC "palm_code_be/src/app/usecases/media"
	pageUC "palm_code_be/src/app/usecases/page"
	teamUC "palm_code_be/src/app/usecases/team"
	uploadUC "palm_code_be/src/app/usecases/upload"
	userUC "palm_code_be/src/app/usecases/user"

	_ "github.com/joho/godotenv/autoload"

	"github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()

	conf := config.Make()

	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	postgresdb, err := postgres.New(conf.SqlDb, logger)

	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.Conn.DB, postgresdb.Conn.DriverName())

	userRepository := userRepo.NewUserRepository(postgresdb.Conn)
	pageRepository := pageRepo.NewPageRepository(postgresdb.Conn)
	mediaRepository := mediaRepo.NewMediaRepository(postgresdb.Conn)
	teamRepository := teamRepo.NewTeamRepository(postgresdb.Conn)
	fsIntegration := fsInteg.NewFireStorage()
	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{
			UserUC:   userUC.NewUserUseCase(userRepository),
			UpLoadUC: uploadUC.NewUploadUseCase(fsIntegration),
			PageUC:   pageUC.NewPageUseCase(pageRepository),
			MediaUC:  mediaUC.NewMediaUseCase(mediaRepository),
			TeamUC:   teamUC.NewTeamUseCase(teamRepository),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
