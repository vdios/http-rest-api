package apiserver

import(
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T){
  s := New(NewConfig())
  rec := httptest.NewRecorder()
  req,_ := http.NewRequest(http.MethodGet,"/hello",nil)
  s.handleHello().ServeHTTP(rec,req)
  assert.Equal(t,rec.Body.String(),"hello buddy =)")
}
