package film

import (
	"bytes"
	"fmt"
	"github.com/Hajymuhammet03/internal/appresult"
	"github.com/Hajymuhammet03/internal/handlers"
	"github.com/Hajymuhammet03/pkg/logging"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	uploadFilm = "/upload-film"
)

type handler struct {
	logger     *logging.Logger
	repository Repository
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc(uploadFilm, appresult.Middleware(h.UploadFilm)).Methods(http.MethodPost)
}

func (h *handler) UploadFilm(w http.ResponseWriter, r *http.Request) error {
	r.ParseMultipartForm(32 << 20)

	file, _, err := r.FormFile("film")
	if err != nil {
		return appresult.ErrMissingParam
	}
	defer file.Close()

	chunkNumber := r.FormValue("chunkNumber")
	totalChunks := r.FormValue("totalChunks")

	totalChunksNum, err := strconv.Atoi(totalChunks)

	chunkNum, err := strconv.Atoi(chunkNumber)
	if err != nil {
		return appresult.ErrMissingParam
	}

	tempFileName := fmt.Sprintf("./temp/movie.part.%d", chunkNum)
	tempFile, err := os.Create(tempFileName)
	if err != nil {
		return appresult.ErrInternalServer
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return appresult.ErrInternalServer
	}

	if chunkNum == totalChunksNum {
		finalFile, err := os.Create("./temp/movie.mp4")
		if err != nil {
			return appresult.ErrInternalServer
		}

		defer finalFile.Close()

		for i := 1; i <= totalChunksNum; i++ {
			chunkFileName := fmt.Sprintf("./temp/movie.part.%d.mp4", i)
			chunkFile, err := os.Open(chunkFileName)
			if err != nil {
				return appresult.ErrInternalServer
			}

			_, err = io.Copy(finalFile, chunkFile)

			if err != nil {
				return appresult.ErrInternalServer
			}
			chunkFile.Close()
			os.Remove(chunkFileName)
		}
		duration, err := GetMovieDuration("./temp/movie.mp4")
		if err != nil {
			return appresult.ErrInternalServer
		}

		fmt.Fprintf(w, "Successfully uploaded file\n")
		fmt.Fprintf(w, "Movie Duration: %s seconds\n", duration)
	} else {
		fmt.Fprintf(w, "Chunk %d uploaded successfully\n", chunkNum)
	}
	return nil
}

func GetMovieDuration(filePath string) (string, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	duration := strings.TrimSpace(out.String())
	return duration, nil
}
