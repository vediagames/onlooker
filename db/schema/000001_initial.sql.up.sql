CREATE TABLE "sessions"
(
    "uuid"        uuid UNIQUE PRIMARY KEY default gen_random_uuid(),
    "client_time" timestamp,
    "server_time" timestamp,
    "ip"          text,
    "url"         text,
    "timezone"    text,
    "metadata"    jsonb
);

CREATE TABLE "levels"
(
    "uuid"         uuid UNIQUE PRIMARY KEY default gen_random_uuid(),
    "session_uuid" uuid,
    "client_time"  timestamp,
    "server_time"  timestamp,
    "level"        int,
    "metadata"     jsonb
);

CREATE TABLE "level_complete_events"
(
    "uuid"        uuid UNIQUE PRIMARY KEY default gen_random_uuid(),
    "level_uuid"  uuid,
    "client_time" timestamp,
    "server_time" timestamp,
    "metadata"    jsonb
);

CREATE TABLE "level_death_events"
(
    "uuid"        uuid UNIQUE PRIMARY KEY default gen_random_uuid(),
    "level_uuid"  uuid,
    "client_time" timestamp,
    "server_time" timestamp,
    "metadata"    jsonb
);

CREATE TABLE "level_grappling_hook_events"
(
    "uuid"        uuid UNIQUE PRIMARY KEY default gen_random_uuid(),
    "level_uuid"  uuid,
    "client_time" timestamp,
    "server_time" timestamp,
    "metadata"    jsonb
);

ALTER TABLE "levels"
    ADD FOREIGN KEY ("session_uuid") REFERENCES "sessions" ("uuid");

ALTER TABLE "level_complete_events"
    ADD FOREIGN KEY ("level_uuid") REFERENCES "levels" ("uuid");

ALTER TABLE "level_death_events"
    ADD FOREIGN KEY ("level_uuid") REFERENCES "levels" ("uuid");

ALTER TABLE "level_grappling_hook_events"
    ADD FOREIGN KEY ("level_uuid") REFERENCES "levels" ("uuid");