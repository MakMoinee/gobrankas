package routes

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/MakMoinee/gobrankas/internal/gobrankas/common"
	"github.com/MakMoinee/gobrankas/internal/gobrankas/views"
	"github.com/MakMoinee/gobrankas/internal/pkg/localhttp"
	"github.com/go-chi/cors"
)

type routesHandler struct {
}

type Location struct {
	Name string
	Url  string
}

type authToken struct {
	Token string `json:"auth-token"`
}

func Set(httpService *localhttp.Service) {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-TOKEN"},
		ExposedHeaders:   []string{"Link", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	httpService.Router.Use(cors.Handler)
	svc := routesHandler{}
	initiateRoutes(httpService, svc)
}

func initiateRoutes(httpService *localhttp.Service, handler routesHandler) {
	httpService.Router.Get(common.HomePath, handler.GetHome)
	httpService.Router.Post(common.UploadPath, handler.Upload)
}

func (svc *routesHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes:GetHome()")

	loc := []Location{{Name: "home", Url: "http://localhost:8443/"}}
	tmpl, err := template.New("home").Parse(views.HomeView)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	err = tmpl.Execute(w, loc)
	if err != nil {
		log.Fatal(err)
	}
}

func (svc *routesHandler) Upload(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside routes:Upload()")

	token := r.Header.Get("auth-token")
	n := fmt.Sprintf("%v", time.Now().Unix())
	// Retrieve the file from form data
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	contentType := header.Header.Get("Content-Type")
	log.Println("Content-Type, ", contentType)
	if !(header.Size <= common.MbSize) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("image size exceed 8 mb"))
		return
	}

	if !strings.EqualFold(token, common.AUTH_TOKEN) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("not authorized"))
		return
	}

	if !strings.EqualFold(contentType, "image/png") {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("content type not allowed"))
		return
	}

	path := filepath.Join("../../", "files")
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + n + filepath.Ext(header.Filename)

	filess, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer filess.Close()
	// Copy the file to the destination path
	_, err = io.Copy(filess, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("hello world"))
}
