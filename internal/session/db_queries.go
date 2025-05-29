package session

const selectFinalClassificationSQLQuery = `
SELECT 
    starting_position, 
    finishing_position 
FROM final_classifications
WHERE session_id = $1
ORDER BY frame_identifier DESC
LIMIT 1;
`

const selectSessionReferencesSQLQuery = `
SELECT DISTINCT weather,
                total_laps, 
                track_id
FROM session_telemetry
WHERE session_id = $1;
`

const selectSessionDurationSQLQuery = `
SELECT MAX(duration) FROM session_telemetry
WHERE session_id = $1;
`

const selectFastestLapStatsSQLQuery = `
SELECT 
    	best_lap_time_lap_num,
    	laps_history -> 0 -> 'lap_time_ms' AS lap_time_ms,
       	laps_history -> 0 -> 'sector_1_ms' AS sector_1_ms,
       	laps_history -> 0 -> 'sector_2_ms' AS sector_2_ms,
       	laps_history -> 0 -> 'sector_3_ms' AS sector_3_ms
FROM session_history_telemetry
WHERE session_id = $1
  AND best_lap_time_lap_num = (SELECT MAX(best_lap_time_lap_num)
                               FROM session_history_telemetry
                               WHERE session_id = $1)
ORDER BY frame_identifier DESC
LIMIT 1;
`

const insertSessionSQLQuery = `
INSERT INTO sessions (
                      id, 
                      fastest_lap_time_ms, 
                      fastest_lap_sector_1_ms, 
                      fastest_lap_sector_2_ms, 
                      fastest_lap_sector_3_ms, 
                      fastest_lap_number, 
                      total_laps,
                      track_id,
                      weather, 
                      duration, 
                      starting_position, 
                      finishing_position, 
                      created_at
                      )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);
`

const selectSessionExistsSQLQuery = `
SELECT EXISTS(SELECT id FROM sessions WHERE id = $1);
`

const selectSessionSQLQuery = `
SELECT 
	id, 
	fastest_lap_time_ms, 
	fastest_lap_sector_1_ms, 
	fastest_lap_sector_2_ms, 
	fastest_lap_sector_3_ms, 
	fastest_lap_number, 
	total_laps,
	track_id,
	weather, 
	duration, 
	starting_position, 
	finishing_position, 
	created_at
FROM sessions 
WHERE id = $1
`
