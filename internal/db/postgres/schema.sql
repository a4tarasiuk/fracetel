CREATE TABLE IF NOT EXISTS lap_data_telemetry
(
    session_id          bigint PRIMARY KEY,
    frame_identifier    bigint,
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
