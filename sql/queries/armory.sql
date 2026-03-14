-- name: UpsertWeapon :exec
INSERT INTO weapons (
    instance_id, hash, name, type, power, kills, level, location, 
    tier, icon_url, slot, damage_type, ammo_type, character_id
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(instance_id) DO UPDATE SET
    power = excluded.power,
    kills = excluded.kills,
    level = excluded.level,
    location = excluded.location,
    character_id = excluded.character_id;

-- name: ClearWeaponStats :exec
DELETE FROM weapon_stats WHERE instance_id = ?;
-- name: ClearWeaponPerks :exec
DELETE FROM weapon_perks WHERE instance_id = ?;

-- name: InsertWeaponStat :exec
INSERT INTO 
  weapon_stats (instance_id, stat_name, value) VALUES (?, ?, ?);

-- name: InsertWeaponPerk :exec
INSERT INTO weapon_perks (instance_id, perk_hash, perk_name, is_equipped, socket_index)
VALUES (?, ?, ?, ?, ?);

-- name: GetGodRollCandidates :many
SELECT 
  w.* FROM weapons w
JOIN weapon_perks p ON w.instance_id = p.instance_id
JOIN weapon_stats s ON w.instance_id = s.instance_id
WHERE p.perk_name = ? AND s.stat_name = 'Range' AND s.value > ?;

-- name: UpsertCharacter :exec
INSERT INTO characters (
    character_id, 
    class_type, 
    light_level, 
    emblem_url, 
    last_played,
    emblem_background_path,
    title_record_hash,
    emblem_color_r,
    emblem_color_g,
    emblem_color_b,
    emblem_color_a
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT(character_id) DO UPDATE SET
    light_level = excluded.light_level,
    last_played = excluded.last_played,
    emblem_url = excluded.emblem_url,
    emblem_background_path = excluded.emblem_background_path,
    title_record_hash = excluded.title_record_hash,
    emblem_color_r = excluded.emblem_color_r,
    emblem_color_g = excluded.emblem_color_g,
    emblem_color_b = excluded.emblem_color_b,
    emblem_color_a = excluded.emblem_color_a;

-- name: UpsertCharacterStat :exec
INSERT INTO character_stats (character_id, stat_hash, value)
VALUES (?, ?, ?)
ON CONFLICT(character_id, stat_hash) DO UPDATE SET
    value = excluded.value;


-- name: ClearAllWeaponsData :exec
DELETE FROM weapons;

-- name: ClearAllWeaponStats :exec
DELETE FROM weapon_stats;

-- name: ClearAllWeaponsPerks :exec
DELETE FROM weapon_perks;

-- name: GetAllWeapons :many
SELECT * FROM weapons;

-- name: GetAllWeaponsWithPerks :many
SELECT w.*, 
       GROUP_CONCAT(p.perk_name, ',') as perks_list
FROM weapons w
LEFT JOIN weapon_perks p ON w.instance_id = p.instance_id
GROUP BY w.instance_id;

-- name: SearchWeaponsByPattern :many
SELECT * FROM weapons 
WHERE name LIKE ?
ORDER BY power DESC;

-- name: GetDuplicatesByHash :many
SELECT * FROM weapons 
WHERE hash = ?
ORDER BY power DESC;

-- name: GetWeaponsByLocation :many
SELECT * FROM weapons 
WHERE location = ?
ORDER BY name ASC;

-- name: GetAllCharacters :many
SELECT * FROM characters
ORDER BY last_played DESC;

-- name: GetWeaponsByType :many
SELECT * FROM weapons 
WHERE type = ?
ORDER BY power DESC, name ASC;

-- name: GetWeaponsByName :many
SELECT * FROM weapons 
WHERE name LIKE ?
ORDER BY power DESC;

-- name: GetWeaponComparison :many
SELECT * FROM weapons 
WHERE instance_id IN (sqlc.slice('instance_ids'))
ORDER BY power DESC;


-- name: GetCharactersWithStats :many
SELECT 
    c.character_id, 
    c.class_type, 
    c.light_level, 
    c.emblem_url, 
    c.last_played,
    c.emblem_background_path,
    c.title_record_hash,
    c.emblem_color_r,
    c.emblem_color_g,
    c.emblem_color_b,
    c.emblem_color_a,
    s.stat_hash,
    s.value as stat_value
FROM characters c
LEFT JOIN character_stats s ON c.character_id = s.character_id
ORDER BY c.last_played DESC;

