package validate

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/qri-io/jsonschema"
	"google.golang.org/appengine/v2"
)

func Json(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		ctx := appengine.NewContext(req)
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			encoder := json.NewEncoder(res)
			encoder.Encode(err)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		validator := &jsonschema.Schema{}
		//read in from the file
		//err := json.Unmarshal(registration_schemas.CheckUsernameAvailability, validator)
		pwd, _ := os.Getwd()
		path := req.URL.Path
		if path == "/" {
			next.ServeHTTP(res, req)
			return
		}
		dat, err := ioutil.ReadFile(filepath.Join(pwd, "validate", "schemas", "registration_schemas", path) + ".json")
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			encoder := json.NewEncoder(res)
			encoder.Encode(err)
			return
		}
		err = json.Unmarshal(dat, validator)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			encoder := json.NewEncoder(res)
			encoder.Encode(err.Error())
			return
		}

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			encoder := json.NewEncoder(res)
			encoder.Encode(err.Error())
			return
		}
		errs, err := validator.ValidateBytes(ctx, body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			encoder := json.NewEncoder(res)
			encoder.Encode(err.Error())
			return
		}
		if len(errs) > 0 {
			res.WriteHeader(http.StatusBadRequest)
			encoder := json.NewEncoder(res)
			encoder.Encode(errs)
			return
		}
		next.ServeHTTP(res, req)
	})
}
