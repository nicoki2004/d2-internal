-- +goose Up
CREATE TABLE weapons (
    instance_id TEXT PRIMARY KEY,
    hash INTEGER NOT NULL,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    power INTEGER NOT NULL,
    kills INTEGER NOT NULL DEFAULT 0,
    level INTEGER NOT NULL DEFAULT 0,
    location TEXT NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    tier TEXT,
    icon_url TEXT,
    slot TEXT,
    damage_type TEXT,
    ammo_type INTEGER,
    character_id TEXT
);

CREATE TABLE weapon_stats (
    instance_id TEXT NOT NULL,
    stat_name TEXT NOT NULL,
    value INTEGER NOT NULL,
    PRIMARY KEY (instance_id, stat_name),
    FOREIGN KEY (instance_id) REFERENCES weapons(instance_id) ON DELETE CASCADE
);

CREATE TABLE weapon_perks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    instance_id TEXT NOT NULL,
    perk_hash INTEGER NOT NULL,
    perk_name TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT 1,
    is_equipped BOOLEAN NOT NULL DEFAULT 0,
    socket_index INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (instance_id) REFERENCES weapons(instance_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS characters (
    character_id TEXT PRIMARY KEY,
    class_type INTEGER NOT NULL,
    light_level INTEGER NOT NULL,
    emblem_url TEXT,
    last_played DATETIME,
    emblem_background_path TEXT,
    title_record_hash UNSIGNED INTEGER DEFAULT 0,
    emblem_color_r INTEGER DEFAULT 0,
    emblem_color_g INTEGER DEFAULT 0,
    emblem_color_b INTEGER DEFAULT 0,
    emblem_color_a INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS character_stats (
    character_id TEXT NOT NULL,
    stat_hash UNSIGNED INTEGER NOT NULL,
    value INTEGER NOT NULL,
    PRIMARY KEY (character_id, stat_hash),
    FOREIGN KEY (character_id) REFERENCES characters(character_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE weapon_perks;
DROP TABLE weapon_stats;
DROP TABLE weapons;
DROP TABLE characters;
DROP TABLE IF EXISTS character_stats;
