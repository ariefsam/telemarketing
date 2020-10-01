package restapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/ariefsam/telemarketing/restapi/usecaseinterface"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
func Serve(usecase usecaseinterface.Usecase) {
	api := RestAPI{
		Usecase: usecase,
	}
	r := mux.NewRouter()
	r.HandleFunc("/api/login/firebase", api.LoginByFirebase).Methods("POST")
	r.HandleFunc("/api/upload", api.ImportCustomer).Methods("POST")
	r.HandleFunc("/api/customer", api.ListCustomer).Methods("POST")
	r.HandleFunc("/api/customer/closing", api.ClosingCustomer).Methods("POST")
	r.HandleFunc("/api/customer/call", api.Call).Methods("POST")
	r.HandleFunc("/api/customer/create", api.CreateCustomer).Methods("POST")
	r.HandleFunc("/api/call-log/get", api.GetCallLog).Methods("POST")
	r.HandleFunc("/api/customer/assign", api.AssignCustomer).Methods("POST")
	r.HandleFunc("/api/telemarketer/save", api.SaveTelemarketer).Methods("POST")
	r.HandleFunc("/api/telemarketer/create", api.CreateTelemarketer).Methods("POST")
	r.HandleFunc("/api/telemarketer/get", api.GetTelemarketer).Methods("POST")
	r.HandleFunc("/api/telemarketer/report", api.ReportTelemarketer).Methods("POST")
	r.HandleFunc("/api/me", api.Me).Methods("POST")
	r.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{
			"Status": "Ok",
			"URL":    r.URL.String(),
		}
		JSONView(w, resp, http.StatusOK)
	})

	r.PathPrefix("/").Handler(quasarHandler())
	gzip := gziphandler.GzipHandler(r)
	s := &http.Server{
		Addr:           ":8885",
		Handler:        handlers.CORS()(gzip),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func quasarHandler() http.Handler {
	var dir string
	dir = "./frontend/dist/spa"
	return http.StripPrefix("/", http.FileServer(http.Dir(dir)))
}

func getPostModel(r *http.Request) (post RequestModel) {
	json.NewDecoder(r.Body).Decode(&post)
	return
}
