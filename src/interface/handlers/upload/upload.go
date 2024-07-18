package upload

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	mediaUsecase "palm_code_be/src/app/usecases/media"
	usecases "palm_code_be/src/app/usecases/upload"
	common_error "palm_code_be/src/infra/errors"
	"palm_code_be/src/infra/helper"
	"palm_code_be/src/interface/response"
	dto "palm_code_be/src/app/dto/media"
	_ "github.com/joho/godotenv/autoload"
)

type UploadHandlerInterface interface {
	Upload(w http.ResponseWriter, r *http.Request)
}

type uploadHandler struct {
	response     response.IResponseClient
	usecase      usecases.UploadUsecase
	mediaUsecase mediaUsecase.MediaUsecase
}

func NewUploadHandler(r response.IResponseClient, h usecases.UploadUsecase, m mediaUsecase.MediaUsecase) UploadHandlerInterface {
	return &uploadHandler{
		response:     r,
		usecase:      h,
		mediaUsecase: m,
	}
}

func (h *uploadHandler) Upload(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("missing authorization token")))
		return
	}

	// Verifikasi token
	dataClaim, err := helper.VerifyToken(tokenString)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error receiving file: %v", err)
		return
	}
	defer file.Close()

	ext := filepath.Ext(handler.Filename)

	var fileType string

	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		fileType = "image"
	case ".mp4", ".avi", ".mov", ".wmv":
		fileType = "video"
	default:
		fileType = "unknown"
	}

	tempFile, err := os.Create(filepath.Join("", handler.Filename))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating temporary file: %v", err)
		return
	}
	defer tempFile.Close()

	// Salin data file ke file sementara
	_, err = io.Copy(tempFile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error copying file: %v", err)
		return
	}

	// Mengunggah file ke Firebase Storage
	bucketName := os.Getenv("BUCKET")
	objectName := handler.Filename // Menggunakan nama file dari handler
	filePath := tempFile.Name()

	data, err := h.usecase.Upload(bucketName, objectName, filePath)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error uploading file: %v", err)
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		log.Printf("Failed to delete temporary file %s: %v", filePath, err)
	}

	createDTO := dto.MediaCreateReqDTO{}
	createDTO.CreatedBy = dataClaim.UserID
	createDTO.Type = fileType
	createDTO.URL = data.URL

	err = h.mediaUsecase.Create(&createDTO)

	if err != nil {
		log.Printf("Failed to save file %s: %v to Media", filePath, err)
	}

	h.response.JSON(
		w,
		"Successful Upload",
		data,
		nil,
	)
}

