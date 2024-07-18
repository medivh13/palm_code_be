package usecases

import (
	mediaUC "palm_code_be/src/app/usecases/media"
	pageUC "palm_code_be/src/app/usecases/page"
	teamUC "palm_code_be/src/app/usecases/team"
	uploadUC "palm_code_be/src/app/usecases/upload"
	userUC "palm_code_be/src/app/usecases/user"
)

type AllUseCases struct {
	UserUC   userUC.UserUCInterface
	UpLoadUC uploadUC.UploadUsecase
	PageUC   pageUC.PageUsecase
	TeamUC   teamUC.TeamUsecase
	MediaUC  mediaUC.MediaUsecase
}
