package testUtils

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/CMSgov/bcda-app/conf"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/otiai10/copy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// CtxMatcher allow us to validate that the caller supplied a context.Context argument
// See: https://github.com/stretchr/testify/issues/519
var CtxMatcher = mock.MatchedBy(func(ctx context.Context) bool { return true })

// PrintSeparator prints a line of stars to stdout
func PrintSeparator() {
	fmt.Println("**********************************************************************************")
}

func RandomHexID() string {
	b, err := someRandomBytes(4)
	if err != nil {
		return "not_a_random_client_id"
	}
	return fmt.Sprintf("%x", b)
}

// RandomMBI returns an 11 character string that represents an MBI
func RandomMBI(t *testing.T) string {
	b, err := someRandomBytes(6)
	assert.NoError(t, err)
	return fmt.Sprintf("%x", b)[0:11]
}

func someRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func RandomBase64(n int) string {
	b, err := someRandomBytes(20)
	if err != nil {
		return "not_a_random_base_64_string"
	}
	return base64.StdEncoding.EncodeToString(b)
}

func setEnv(why, key, value string) {
	if err := conf.SetEnv(&testing.T{}, key, value); err != nil {
		log.Printf("Error %s env value %s to %s\n", why, key, value)
	}
}

// SetAndRestoreEnvKey replaces the current value of the env var key,
// returning a function which can be used to restore the original value
func SetAndRestoreEnvKey(key, value string) func() {
	originalValue := conf.GetEnv(key)
	setEnv("setting", key, value)
	return func() {
		setEnv("restoring", key, originalValue)
	}
}

func MakeDirToDelete(s suite.Suite, filePath string) {
	assert := assert.New(s.T())
	_, err := os.Create(filepath.Clean(filepath.Join(filePath, "deleteMe1.txt")))
	assert.Nil(err)
	_, err = os.Create(filepath.Clean(filepath.Join(filePath, "deleteMe2.txt")))
	assert.Nil(err)
	_, err = os.Create(filepath.Clean(filepath.Join(filePath, "deleteMe3.txt")))
	assert.Nil(err)
	_, err = os.Create(filepath.Clean(filepath.Join(filePath, "deleteMe4.txt")))
	assert.Nil(err)
}

// SetPendingDeletionDir sets the PENDING_DELETION_DIR to the supplied "path" and ensures
// that the directory is created
func SetPendingDeletionDir(s suite.Suite, path string) {
	err := conf.SetEnv(s.T(), "PENDING_DELETION_DIR", path)
	if err != nil {
		s.FailNow("failed to set the PENDING_DELETION_DIR env variable,", err)
	}
	cclfDeletion := conf.GetEnv("PENDING_DELETION_DIR")
	err = os.MkdirAll(cclfDeletion, 0744)
	if err != nil {
		s.FailNow("failed to create the pending deletion directory, %s", err.Error())
	}
}

// CopyToTemporaryDirectory copies all of the content found at src into a temporary directory.
// The path to the temporary directory is returned along with a function that can be called to clean up the data.
func CopyToTemporaryDirectory(t *testing.T, src string) (string, func()) {
	newPath, err := ioutil.TempDir("", "*")
	if err != nil {
		t.Fatalf("Failed to create temporary directory %s", err.Error())
	}

	if err = copy.Copy(src, newPath); err != nil {
		t.Fatalf("Failed to copy contents from %s to %s %s", src, newPath, err.Error())
	}

	cleanup := func() {
		err := os.RemoveAll(newPath)
		if err != nil {
			log.Printf("Failed to cleanup data %s", err.Error())
		}
	}

	return newPath, cleanup
}

// GetRandomIPV4Address returns a random IPV4 address using rand.Read() to generate the values.
func GetRandomIPV4Address(t *testing.T) string {
	data := make([]byte, 3)
	if _, err := rand.Read(data); err != nil {
		t.Fatal(err.Error())
	}

	// Use static first byte to ensure that the IP address is valid
	return fmt.Sprintf("146.%d.%d.%d", data[0], data[1], data[2])
}

// GetLogger returns the underlying implementation of the field logger
func GetLogger(logger logrus.FieldLogger) *logrus.Logger {
	if entry, ok := logger.(*logrus.Entry); ok {
		return entry.Logger
	}
	// Must be a *logrus.Logger
	return logger.(*logrus.Logger)
}

//ReadResponseBody will read http.Response and return the body contents as a string.
func ReadResponseBody(r *http.Response) string {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		bodyString, _ := "Error reading the body: %v\n", err
		return bodyString
	}

	bodyString := bytes.NewBuffer(body).String()
	return bodyString
}

//MakeTestServerWithIntrospectEndpoint creates an httptest.Server with an introspect endpoint that will
//return back a response with a json body indicating if "active" is set to true or false (set by active
//token parameter)
func MakeTestServerWithIntrospectEndpoint(activeToken bool) *httptest.Server {
	router := chi.NewRouter()
	router.Post("/introspect", func(w http.ResponseWriter, r *http.Request) {
		var (
			buf   []byte
			input struct {
				Token string `json:"token"`
			}
		)
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Unexpected error creating test server: Error in reading request body: %s", err.Error())
			return
		}

		if unmarshalErr := json.Unmarshal(buf, &input); unmarshalErr != nil {
			fmt.Printf("Unexpected error creating test server: Error in unmarshalling the buffered input to JSON: %s", unmarshalErr.Error())
			return
		}

		body, _ := json.Marshal(struct {
			Active bool `json:"active"`
		}{Active: activeToken})

		_, responseWriterErr := w.Write(body)
		if responseWriterErr != nil {
			fmt.Printf("Unexpected error creating test server: Error reading request body: %s", responseWriterErr.Error())
		}

	})
	return httptest.NewServer(router)
}

//MakeTestServerWithIntrospectTimeout creates an httptest.Server with an introspect endpoint that will sleep for 10 seconds.
//Useful in testing where the env timeout is set to something less (ex. 5 seconds) and you want to ensure *url.Error.Timeout() returns true.
func MakeTestServerWithIntrospectTimeout() *httptest.Server {
	router := chi.NewRouter()
	router.Post("/introspect", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 10)
	})

	return httptest.NewServer(router)
}

//MakeTestServerWithIntrospectReturn502 creates an httptest.Server
//with an introspect endpoint that will return 502 Status Code.
func MakeTestServerWithIntrospectReturn502() *httptest.Server {
	router := chi.NewRouter()
	router.Post("/introspect", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
	})
	return httptest.NewServer(router)
}
