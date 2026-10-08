package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
)

func TestPing(t *testing.T) {
    // 1. Cria o router
    gin.SetMode(gin.TestMode)
    r := gin.New()
    handler := NewPingHandler()
    r.GET("/ping", handler.Ping)

    // 2. Cria uma requisição de teste
    req := httptest.NewRequest(http.MethodGet, "/ping", nil)

    // 3. Cria um gravador de resposta
    w := httptest.NewRecorder()

    // 4. Executa a requisição
    r.ServeHTTP(w, req)

    // 5. Verifica o resultado
    if w.Code != http.StatusOK {
        t.Errorf("esperava status %d, recebeu %d", http.StatusOK, w.Code)
    }
}
