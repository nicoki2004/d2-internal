# d2-internal

Internal service for a Destiny 2 armory prototype. It handles Bungie OAuth, fetches profile data, syncs weapons and character stats into a local SQLite database, and exposes a small HTTPS API for reading that data.

## What it does

- OAuth login with Bungie and token persistence in SQLite.
- Fetches the Destiny 2 profile from Bungie and optionally caches it locally.
- Syncs characters, character stats, weapons, weapon stats, and weapon perks into SQLite.
- Exposes a minimal HTTPS API for login, profile refresh, and character listing.

## Current API

The server listens on `https://localhost:4200` using the TLS certs in `localhost.pem` and `localhost-key.pem`.

Routes registered in `cmd/destiny/main.go`:

- `GET /`
  - If `?code=` is present, exchanges the OAuth code, stores tokens, and sets the active user.
  - If no code is present, attempts to load a persisted user and sets the token in memory.
- `GET /login`
  - Redirects to Bungie OAuth login.
- `GET /refresh`
  - Fetches the Bungie profile, returns it as JSON, then syncs weapons/characters into SQLite.
- `GET /characters`
  - Returns an array of characters with computed stats from SQLite.
- `GET /characters/{id}`
  - Returns a single character by ID.

There is a weapons service stub in `internal/destiny/weapon/service.go`, but no endpoint is currently wired for it.

## Requirements

- Go 1.25+
- A Bungie API key and OAuth app
- SQLite (used via `github.com/glebarez/go-sqlite`)

## Configuration

The app loads environment variables via `.env` (using `github.com/subosito/gotenv`). Required variables are validated at startup.

Create a `.env` file at the repo root with:

```env
BUNGIE_API_KEY=your_key
BUNGIE_OAUTH_CLIENT_ID=your_client_id
BUNGIE_OAUTH_CLIENT_SECRET=your_client_secret
BUNGIE_OAUTH_REDIRECT_URI=https://localhost:4200/
DB_NAME=./destiny.db
CACHE_FILE_NAME=./profile_cache.json
```

Notes:

- `DB_NAME` points to the SQLite file. The repo already includes `destiny.db` for local testing.
- `CACHE_FILE_NAME` is optional. If not set, it defaults to `profile_cache.json`.
- The redirect URI must match your Bungie OAuth app settings.

## TLS certificates

The server starts with `ListenAndServeTLS("localhost.pem", "localhost-key.pem")`. The repo includes those files for local development.

If you replace them, keep the same filenames or update the code in `cmd/destiny/main.go`.

## Database schema

SQLite schema lives in `sql/schema`:

- `users` (OAuth tokens and membership info)
- `characters` and `character_stats`
- `weapons`, `weapon_stats`, `weapon_perks`

Query code is generated with `sqlc` using `sqlc.yaml` and output to `internal/database`.

## Manifests and cache files

The service relies on local manifest files checked into the repo:

- `items_manifest.json` for item definitions (weapons, tiers, icons, etc.)
- `definition_manifest.json` for title/record definitions

The profile cache (if enabled) is stored in `profile_cache.json` (or `CACHE_FILE_NAME`).

## Running locally

From the repo root:

```bash
go run ./cmd/destiny
```

Then open:

- `https://localhost:4200/login` to start Bungie OAuth
- `https://localhost:4200/refresh` to fetch and sync your profile
- `https://localhost:4200/characters` to list characters from SQLite

## Project layout

- `cmd/destiny` - main server entrypoint
- `internal/apicfg` - HTTP handlers and API wiring
- `internal/auth` - Bungie OAuth flows
- `internal/config` - configuration loading and validation
- `internal/destiny` - Bungie API client + sync logic
- `internal/destiny/character` - character DTO and service
- `internal/database` - sqlc generated DB access layer
- `internal/repository` - repository abstraction over sqlc
- `sql/schema`, `sql/queries` - SQLite schema and queries

## Known gaps / WIP

- Weapons endpoint and DTOs are incomplete (`internal/destiny/weapon/service.go`).
- Error handling around missing users in `/` uses a hardcoded membership ID for fallback.

## License

No license has been specified yet.
