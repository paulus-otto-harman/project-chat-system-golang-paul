package helper

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"homework/domain"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"sync"
)

func Upload(wg *sync.WaitGroup, files []*multipart.FileHeader) ([]domain.CdnResponse, error) {
	var results []domain.CdnResponse
	var err error
	for _, file := range files {
		wg.Add(1)

		go func() {
			defer wg.Done()
			var f multipart.File
			f, err = file.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			var part io.Writer
			part, err = writer.CreateFormFile("image", filepath.Base(""))
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(part, f)

			err = writer.Close()
			if err != nil {
				log.Fatal(err)
			}

			var request *http.Request
			request, err = http.NewRequest("POST", "https://cdn-lumoshive-academy.vercel.app/api/v1/upload", body)
			if err != nil {
				log.Fatal(err)
			}
			request.Header.Add("Content-Type", writer.FormDataContentType())
			client := &http.Client{}
			var response *http.Response
			response, err = client.Do(request)
			if err != nil {
				log.Fatal(err)
			}

			defer response.Body.Close()

			var res []byte
			res, err = io.ReadAll(response.Body)
			if err != nil {
				log.Fatal("Error reading response:", err)
			}

			var result domain.CdnResponse
			json.Unmarshal(res, &result)
			results = append(results, result)
		}()
	}
	wg.Wait()
	return results, err
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func BadResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, Response{
		Status:  false,
		Message: message,
	})
}

func GoodResponseWithData(c *gin.Context, message string, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func GoodResponseWithPage(c *gin.Context, message string, statusCode, total, totalPages, page, Limit int, data interface{}) {
	c.JSON(statusCode, domain.DataPage{
		Status:      true,
		Message:     message,
		Total:       int64(total),
		Pages:       totalPages,
		CurrentPage: uint(page),
		Limit:       uint(Limit),
		Data:        data,
	})
}