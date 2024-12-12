package bannercontroller_test

import (
	"bytes"
	"dashboard-ecommerce-team2/infra"
	"dashboard-ecommerce-team2/models"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type FormData map[string]interface{}

func createMultipartFormFile(fieldName, fileName string) (*multipart.FileHeader, *os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	// Menulis data ke file
	_, err = file.Write([]byte("test image data"))
	if err != nil {
		return nil, nil, err
	}

	// Membuat file multipart
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		return nil, nil, err
	}
	part.Write([]byte("test image data"))

	writer.Close()

	fileHeader := &multipart.FileHeader{
		Filename: fileName,
		Size:     int64(len("test image data")),
		Header:   make(map[string][]string),
	}
	return fileHeader, file, nil
}

// func SetupTestDB() (*gorm.DB, sqlmock.Sqlmock, error) {
// 	// Setup sqlmock
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("failed to initialize sqlmock: %w", err)
// 	}

// 	// Setup GORM with mock database
// 	dialector := postgres.New(postgres.Config{
// 		Conn: db,
// 	})
// 	gormDB, err := gorm.Open(dialector, &gorm.Config{})
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("failed to open gorm with mock: %w", err)
// 	}

// 	return gormDB, mock, nil
// }
// func SetupTestContext() (*controller.Controller, error) {
// 	// Setup mock database
// 	db, _, err := SetupTestDB()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Instance logger
// 	log, err := helper.InitZapLogger()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Mock configuration
// 	cfg, err := config.ReadConfig()
// 	if err != nil {
// 		return nil, err

// 	}

// 	// Mock cache
// 	rdb := database.NewCacher(cfg, 60*60)

// 	// Repository, service, and controller
// 	repo := repository.NewRepository(db, log)
// 	svc := service.NewService(repo, log)
// 	ctrl := controller.NewController(svc, log, rdb, cfg)

// 	// Setup router
// 	router := gin.Default()
// 	router.POST("/create-banner", ctrl.Banner.CreateBannerController)
// 	router.PUT("/banner", ctrl.Banner.UpdateBannerController)
// 	router.GET("/banner", ctrl.Banner.GetBannerByIDController)
// 	router.DELETE("/banner", ctrl.Banner.DeleteBannerController)

// 	return ctrl, nil
// }

func TestBannerController(t *testing.T) {
	// ctx, err := SetupTestContext()
	ctx, err := infra.NewServiceContext()
	if err != nil {
		t.Fatalf("Database Error : %v\n", err)
	}

	// Inisialisasi router
	route := gin.Default()
	route.POST("/create-banner", ctx.Ctl.Banner.CreateBannerController)
	route.PUT("/banner", ctx.Ctl.Banner.UpdateBannerController)
	route.GET("/banner", ctx.Ctl.Banner.GetBannerByIDController)
	route.DELETE("/banner", ctx.Ctl.Banner.DeleteBannerController)

	// imageFileHeader, _, err := createMultipartFormFile("image_path", "test_image.png")
	// if err != nil {
	// 	assert.NoError(t, err, "Failed to create multipartfile")
	// }

	now := time.Now()
	tests := []struct {
		name       string
		method     string
		url        string
		formData   FormData
		expectCode int
		validate   func(t *testing.T, body []byte)
	}{
		// Test CreateBannerController
		// {
		// 	name:   "Success - Create Banner",
		// 	method: http.MethodPost,
		// 	url:    "/create-banner",
		// 	formData: FormData{
		// 		"image_path":   imageFileHeader,
		// 		"title":        "New Banner",
		// 		"type":         []string{"seasonal", "promo"},
		// 		"path_page":    "/spring-sale",
		// 		"release_date": now.Format(time.RFC3339),
		// 		"end_date":     now.Add(30 * 24 * time.Hour).Format(time.RFC3339),
		// 		"published":    true,
		// 	},
		// 	expectCode: http.StatusOK,
		// 	validate: func(t *testing.T, body []byte) {
		// 		assert.Contains(t, string(body), "Successfully created banner", "Response should indicate success")
		// 	},
		// },
		{
			name:   "Failed - Missing Image Path",
			method: http.MethodPost,
			url:    "/create-banner",
			formData: FormData{
				"title":        "Banner Image Missing",
				"type":         []string{"seasonal", "promo"},
				"path_page":    "/spring-sale",
				"release_date": now.Format(time.RFC3339),
				"end_date":     now.Add(30 * 24 * time.Hour).Format(time.RFC3339),
				"published":    true,
			},
			expectCode: http.StatusBadRequest,
			validate: func(t *testing.T, body []byte) {
				assert.Contains(t, string(body), "invalid data input", "Response should indicate invalid data input")
			},
		},

		{
			name:   "Failed - Type missing",
			method: http.MethodPost,
			url:    "/create-banner",
			formData: FormData{
				"title":        "Banner Image Missing",
				"type":         []string{},
				"path_page":    "/spring-sale",
				"release_date": now.Format(time.RFC3339),
				"end_date":     now.Add(30 * 24 * time.Hour).Format(time.RFC3339),
				"published":    true,
			},
			expectCode: http.StatusBadRequest,
			validate: func(t *testing.T, body []byte) {
				assert.Contains(t, string(body), "invalid data input", "Response should indicate invalid data input")
			},
		},

		// {
		// 	name:   "Failed - Database Error",
		// 	method: http.MethodPost,
		// 	url:    "/create-banner",
		// 	formData: FormData{
		// 		"id":           1,
		// 		"image_path":   imageFileHeader,
		// 		"title":        "New Banner",
		// 		"type":         []string{"seasonal", "promo"},
		// 		"path_page":    "/spring-sale",
		// 		"release_date": now.Format(time.RFC3339),
		// 		"end_date":     now.Add(30 * 24 * time.Hour).Format(time.RFC3339),
		// 		"published":    true,
		// 	},
		// 	expectCode: http.StatusBadRequest,
		// 	validate: func(t *testing.T, body []byte) {
		// 		assert.Contains(t, string(body), "invalid data input", "Response should unix id data input")
		// 	},
		// },

		// Test UpdateBannerController
		{
			name:       "Success - Update Banner",
			method:     http.MethodPut,
			url:        "/banner",
			formData:   map[string]interface{}{"id": "2"},
			expectCode: http.StatusOK,
			validate: func(t *testing.T, body []byte) {
				assert.Contains(t, string(body), "Successfully update published banner", "Response should indicate success")
			},
		},
		{
			name:       "Failed - Invalid ID",
			method:     http.MethodPut,
			url:        "/banner",
			formData:   FormData{"id": "abc"},
			expectCode: http.StatusBadRequest,
			validate: func(t *testing.T, body []byte) {
				assert.Contains(t, string(body), "id banner invalid", "Response should indicate invalid ID")
			},
		},

		{
			name:       "Failed - ID Not Found Update",
			method:     http.MethodPut,
			url:        "/banner",
			formData:   FormData{"id": "9999999"},
			expectCode: http.StatusInternalServerError,
			validate: func(t *testing.T, body []byte) {
				assert.Contains(t, string(body), "banner with ID 9999999 not found", "Response should indicate invalid ID")
			},
		},

		// Test GetBannerByIDController
		{
			name:       "Success - Valid ID",
			method:     http.MethodGet,
			url:        "/banner",
			formData:   FormData{"id": "2"},
			expectCode: http.StatusOK,
			validate: func(t *testing.T, body []byte) {
				// Struktur respons umum
				var response struct {
					Data    models.Banner `json:"data"`
					Message string        `json:"message"`
					Status  int           `json:"status"`
				}

				// Log respons untuk debugging
				t.Logf("Response body: %s", string(body))

				// Unmarshal body ke dalam struktur respons
				err := json.Unmarshal(body, &response)
				assert.NoError(t, err, "Response body should be a valid JSON")

				// Validasi data banner
				assert.Equal(t, "Spring Promo 2000", response.Data.Title, "Banner title does not match")
				assert.Equal(t, "Successfully get banner", response.Message, "Response message does not match")
			},
		},
		{
			name:       "Failed - Missing ID",
			method:     http.MethodGet,
			url:        "/banner",
			formData:   FormData{},
			expectCode: http.StatusBadRequest,
			validate: func(t *testing.T, body []byte) {
				assert.Contains(t, string(body), "id banner invalid", "Response should indicate missing ID")
			},
		},
		{
			name:       "Failed - ID Not Found",
			method:     http.MethodGet,
			url:        "/banner",
			formData:   FormData{"id": "99999"},
			expectCode: http.StatusInternalServerError,
			validate: func(t *testing.T, body []byte) {
				assert.Contains(t, string(body), "banner with ID 99999 not found", "Response should indicate invalid ID")
			},
		},

		// Test DeleteBannerController
		{
			name:       "Success - Delete Banner",
			method:     http.MethodDelete,
			url:        "/banner",
			formData:   FormData{"id": "6"},
			expectCode: http.StatusOK,
			validate: func(t *testing.T, body []byte) {
				assert.Contains(t, string(body), "Successfully delete banner", "Response should indicate success")
			},
		},
		{
			name:       "Failed - Invalid ID",
			method:     http.MethodDelete,
			url:        "/banner",
			formData:   FormData{"id": "abc"},
			expectCode: http.StatusBadRequest,
			validate: func(t *testing.T, body []byte) {
				assert.Contains(t, string(body), "id banner invalid", "Response should indicate invalid ID")
			},
		},

		{
			name:       "Failed - ID Not Found",
			method:     http.MethodDelete,
			url:        "/banner",
			formData:   FormData{"id": "99999"},
			expectCode: http.StatusInternalServerError,
			validate: func(t *testing.T, body []byte) {
				assert.Contains(t, string(body), "banner with ID 99999 not found", "Response should indicate invalid ID")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Membuat form-data
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			for key, value := range tt.formData {
				switch v := value.(type) {
				case string:
					writer.WriteField(key, v)
				case []string:
					for _, item := range v {
						writer.WriteField(key, item)
					}
				case bool:
					writer.WriteField(key, strconv.FormatBool(v))
				case int:
					writer.WriteField(key, strconv.Itoa(v))
				case float64:
					writer.WriteField(key, fmt.Sprintf("%f", v))
				case nil:
					// Abaikan field dengan nilai nil
					continue
				case *multipart.FileHeader:
					file, err := v.Open()
					if err != nil {
						t.Fatalf("failed to open file for key %s: %v", key, err)
					}
					defer file.Close()
					part, err := writer.CreateFormFile(key, v.Filename)
					if err != nil {
						t.Fatalf("failed to create form file for key %s: %v", key, err)
					}
					_, err = io.Copy(part, file)
					if err != nil {
						t.Fatalf("failed to copy file data for key %s: %v", key, err)
					}
				default:
					t.Fatalf("unsupported formData value type for key %s", key)
				}
			}
			writer.Close()

			// Membuat request dengan form-data
			r := httptest.NewRequest(tt.method, tt.url, body)
			r.Header.Set("Content-Type", writer.FormDataContentType())
			w := httptest.NewRecorder()

			// Menjalankan request
			route.ServeHTTP(w, r)

			// Membaca response
			res := w.Result()
			defer res.Body.Close()

			b, err := io.ReadAll(res.Body)
			assert.NoError(t, err, "Failed to read response body")

			// Validasi status code
			assert.Equal(t, tt.expectCode, w.Code, "Expected status code %d, got %d", tt.expectCode, w.Code)

			// Validasi body response
			if tt.validate != nil {
				tt.validate(t, b)
			}
		})
	}
}
