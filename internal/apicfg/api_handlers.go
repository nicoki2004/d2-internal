package apicfg

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nicoki2004/d2-internal/internal/auth"
	"github.com/nicoki2004/d2-internal/internal/database"
	"github.com/nicoki2004/d2-internal/internal/destiny"
	"github.com/nicoki2004/d2-internal/internal/destiny/character"
	"github.com/nicoki2004/d2-internal/internal/destiny/store"
	"github.com/nicoki2004/d2-internal/internal/logger"
)

func (cfg *APIConfig) HandlerHome(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code != "" {
		cfg.Log.Info("Code received, exchanging for token...")

		token, err := auth.ExchangeCode(cfg.Client.Cfg, code)
		if err != nil {
			cfg.Log.Error("Exchange error: %v", err)
			http.Error(w, "Failed to exchange code", http.StatusInternalServerError)
			return
		}

		now := time.Now()

		// Add token to APIConfig
		cfg.Client.Token = token
		miID, err := destiny.GetMembershipForCurrentUser(&cfg.Client)
		if err != nil {
			logger.GetLogger().Debug("Error gettin membership: %v", err)
		}
		token.MembershipID = miID.Response.PrimaryMembershipID
		token.DisplayName = miID.Response.DestinyMemberships[0].DisplayName
		cfg.Client.Token = token

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

		cfg.Client.Token.MembershipID = miID.Response.PrimaryMembershipID

		cfg.Log.Info("Usuario %s persistido correctamente", token.MembershipID)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>¡Autenticación exitosa!</h1><p>Puedes volver a la terminal.</p>"))
		return
	}

	user, err := cfg.Repo.GetUser(r.Context(), "")
	if err != nil {
		logger.GetLogger().Debug("Error getting user: %v", err)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("Servidor activo. Ve a <a href='/login'>/login</a> para empezar."))
		return
	}
	cfg.Client.Token.MembershipID = user.MembershipID
	cfg.Client.Token.ExpiresIn = int(user.ExpiresIn)
	cfg.Client.Token.RefreshToken = user.RefreshToken
	cfg.Client.Token.AccessToken = user.AccessToken
	cfg.Client.Token.RefreshToken = user.RefreshToken
	cfg.Client.Token.RefreshExpiresIn = int(user.RefreshExpiresIn)
	cfg.Client.Token.ReceivedAt = user.ReceivedAt

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1>¡Autenticación exitosa!</h1><p>Puedes volver a la terminal.</p>"))
}

func (cfg *APIConfig) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	url, err := auth.AuthURL(*cfg.Client.Cfg)
	if err != nil {
		http.Error(w, "Error generando URL", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (cfg *APIConfig) HandlerGetProfile(w http.ResponseWriter, r *http.Request) {
	log := logger.GetLogger()
	components := []string{}

	dataProfile, err := destiny.GetProfile(&cfg.Client, components...)
	if err != nil {
		log.Fatal("Error getting profile: %v", err)
	}

	stores := store.GetStores(*dataProfile)

	jsonData, err := json.Marshal(stores)
	if err != nil {
		http.Error(w, "Error al procesar JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

	// manifest, err := destiny.LoadManifestItem("items_manifest.json")
	// if err != nil {
	// 	log.Error("Error loading manifest: %v", err)
	// 	return
	// }
	//
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	// defer cancel()

	// if err := destiny.SyncInventory(ctx, cfg.Repo, dataProfile, manifest); err != nil {
	// 	log.Fatal("Error syncing inventory: %v", err)
	// }
}

func (cfg *APIConfig) HandlerGetCharacters(w http.ResponseWriter, r *http.Request) {
	characterID := r.PathValue("id")

	result, err := character.GetCharacterProfile(r.Context(), cfg.Repo)
	if err != nil {
		logger.GetLogger().Error("Error fetching Bungie profile: %v", err)
		http.Error(w, "Error getting characters", http.StatusInternalServerError)
		return
	}

	var response interface{}

	if characterID != "" {
		var selected *character.CharacterDTO
		for _, char := range result {
			if char.ID == characterID {
				c := char
				selected = &c
				break
			}
		}

		if selected == nil {
			http.Error(w, "Character not found", http.StatusNotFound)
			return
		}
		response = selected
	} else {
		response = result
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

// func (cfg *APIConfig) HandlerGetWeapons(w http.ResponseWriter, r *http.Request) {
// 	characterID := r.PathValue("id")
// 	result
// }
