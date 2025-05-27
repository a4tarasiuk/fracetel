CREATE TABLE IF NOT EXISTS lap_data_telemetry
(
    session_id          BIGINT PRIMARY KEY,
    frame_identifier    BIGINT,
    occurred_at         TIMESTAMP,
    last_lap_time_ms    INTEGER,
    current_lap_time_ms INTEGER,
    first_sector_ms     INTEGER,
    second_sector_ms    INTEGER,
    lap_distance        REAL,
    total_distance      REAL,
    car_position        SMALLINT,
    current_lap_num     SMALLINT,
    sector              SMALLINT,
    driver_status       SMALLINT
);

CREATE TABLE IF NOT EXISTS final_classifications
(
    session_id         BIGINT PRIMARY KEY,
    frame_identifier   BIGINT,
    occurred_at        TIMESTAMP,

    finishing_position SMALLINT,
    starting_position  SMALLINT
);

CREATE TABLE IF NOT EXISTS session_telemetry
(
    session_id        BIGINT PRIMARY KEY,
    frame_identifier  BIGINT,
    occurred_at       TIMESTAMP,
    weather           SMALLINT,
    track_temperature SMALLINT,
    air_temperature   SMALLINT,
    total_laps        SMALLINT,
    track_length      SMALLINT,
    track_id          SMALLINT,
    type              INTEGER,
    time_left         INTEGER,
    duration          INTEGER
);

CREATE TABLE IF NOT EXISTS session_history_telemetry
(
    session_id            BIGINT PRIMARY KEY,
    frame_identifier      BIGINT,
    occurred_at           TIMESTAMP,
    num_laps              INTEGER,

    best_lap_time_lap_num INTEGER,

    best_sector_1_lap_num INTEGER,
    best_sector_2_lap_num INTEGER,
    best_sector_3_lap_num INTEGER,

    laps_history           JSON
);
