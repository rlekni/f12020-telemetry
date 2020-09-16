package clients

const (
	PacketHeader                  = "packetHeader"
	CarMotionData                 = "carMotionData"
	PacketMotionData              = "packetMotionData"
	MarshalZone                   = "marshalZone"
	WeatherForecastSample         = "weatherForecastSample"
	PacketSessionData             = "packetSessionData"
	PacketLapData                 = "packetLapData"
	PacketEventData               = "packetEventData"
	PacketParticipantsData        = "packetParticipantsData"
	PacketCarSetupData            = "packetCarSetupData"
	PacketCarTelemetryData        = "packetCarTelemetryData"
	PacketCarStatusData           = "packetCarStatusData"
	PacketFinalClassificationData = "packetFinalClassificationData"
	PacketLobbyInfoData           = "packetLobbyInfoData"
)

const (
	packetHeaderSQL = `CALL public.insert_packet_header($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	carMotionDataSQL    = `CALL public.insert_car_motion_data($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)`
	packetMotionDataSQL = `CALL public.insert_packet_motion_data($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30)`

	marshalZoneSQL           = `CALL public.insert_marshal_zone($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	weatherForecastSampleSQL = `CALL public.insert_weather_forecast_sample($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	packetSessionDataSQL     = `CALL public.insert_packet_session_data($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	packetLapDataSQL                 = `CALL public.insert_packet_lap_data($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	packetEventDataSQL               = `CALL public.insert_packet_event_data($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	packetParticipantsDataSQL        = `CALL public.insert_packet_participants_data($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	packetCarSetupDataSQL            = `CALL public.insert_packet_car_setup_data($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	packetCarTelemetryDataSQL        = `CALL public.insert_packet_car_telemetry_data($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	packetCarStatusDataSQL           = `CALL public.insert_packet_car_status_data($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	packetFinalClassificationDataSQL = `CALL public.insert_packet_final_classification_data($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	packetLobbyInfoDataSQL           = `CALL public.insert_packet_lobby_info_data($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
)
