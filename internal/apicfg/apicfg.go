package apicfg

import (
	"net/http"
	"time"

	"github.com/nicoki2004/d2-internal/internal/auth"
	"github.com/nicoki2004/d2-internal/internal/config"
	"github.com/nicoki2004/d2-internal/internal/database"
	"github.com/nicoki2004/d2-internal/internal/logger"
	"github.com/nicoki2004/d2-internal/internal/repository"
)

type APIConfig struct {
	DB   *database.Queries
	Cfg  *config.Config
	Log  logger.Logger
	Repo repository.DestinyRepository
}

func (cfg *APIConfig) HandlerHome(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code != "" {
		cfg.Log.Info("Code received, exchanging for token...")

		token, err := auth.ExchangeCode(cfg.Cfg, code)
		if err != nil {
			cfg.Log.Error("Exchange error: %v", err)
			http.Error(w, "Failed to exchange code", http.StatusInternalServerError)
			return
		}

		now := time.Now()

		err = cfg.Repo.UpsertUser(r.Context(), database.UpsertUserParams{
			MembershipID:     token.MembershipID,
			DisplayName:      token.DisplayName,
			AccessToken:      token.AccessToken,
			RefreshToken:     token.RefreshToken,
			ExpiresIn:        int64(token.ExpiresIn),
			RefreshExpiresIn: int64(token.RefreshExpiresIn),
			ReceivedAt:       now,
			UpdatedAt:        now,
		})
		if err != nil {
			cfg.Log.Error("Error guardando en DB: %v", err)
			http.Error(w, "Error al guardar sesión", http.StatusInternalServerError)
			return
		}

		cfg.Log.Info("Usuario %s persistido correctamente", token.MembershipID)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>¡Autenticación exitosa!</h1><p>Puedes volver a la terminal.</p>"))
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("Servidor activo. Ve a <a href='/login'>/login</a> para empezar."))
}

func (cfg *APIConfig) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	url, err := auth.AuthURL(*cfg.Cfg)
	if err != nil {
		http.Error(w, "Error generando URL", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
