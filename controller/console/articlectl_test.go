package console

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/88250/pipe/util"
	"github.com/gin-gonic/gin"
)

func TestMarkdownActionAcceptsMarkdownField(t *testing.T) {
	gin.SetMode(gin.TestMode)

	body := bytes.NewBufferString(`{"markdown":"Hello Pipe"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/markdown", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	MarkdownAction(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal response failed: %v", err)
	}

	if resp.Code != util.CodeOk {
		t.Fatalf("expected code %d, got %d (%s)", util.CodeOk, resp.Code, resp.Msg)
	}

	html, ok := resp.Data.(string)
	if !ok {
		t.Fatalf("expected data to be string, got %T", resp.Data)
	}

	if !strings.Contains(html, "Hello Pipe") {
		t.Fatalf("expected html to contain markdown content, got %q", html)
	}
}
