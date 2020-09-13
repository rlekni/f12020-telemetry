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
	packetHeaderSQL = `
	INSERT INTO PacketHeader (PacketFormat, GameMajorVersion, GameMinorVersion, PacketVersion, PacketID, SessionUID, SessionTime, FrameIdentifier, PlayerCarIndex, SecondaryPlayerCarIndex)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id`

	carMotionDataSQL = `
	INSERT INTO CarMotionData (WorldPositionX, WorldPositionY, WorldPositionZ, WorldVelocityX, WorldVelocityY, WorldVelocityZ, WorldForwardDirX, WorldForwardDirY, WorldForwardDirZ, WorldRightDirX, WorldRightDirY, WorldRightDirZ, GForceLateral, GForceLongitudinal, GForceVertical, Yaw, Pitch, Roll)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
	RETURNING id`
	packetMotionDataSQL = `
	INSERT INTO PacketMotionData (SuspensionPositionRL, SuspensionPositionRR, SuspensionPositionFL, SuspensionPositionFR, SuspensionVelocityRL, SuspensionVelocityRR, SuspensionVelocityFL, SuspensionVelocityFR, SuspensionAccelerationRL, SuspensionAccelerationRR, SuspensionAccelerationFL, SuspensionAccelerationFR, WheelSpeedRL, WheelSpeedRR, WheelSpeedFL, WheelSpeedFR, WheelSlipRL, WheelSlipRR, WheelSlipFL, WheelSlipFR, LocalVelocityX, LocalVelocityY, LocalVelocityZ, AngularVelocityX, AngularVelocityY, AngularVelocityZ, AngularAccelerationX, AngularAccelerationY, AngularAccelerationZ, FrontWheelsAngle)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30)
	RETURNING id`

	marshalZoneSQL           = ``
	weatherForecastSampleSQL = ``

	packetSessionDataSQL = `
	INSERT INTO PacketSessionData (type)
	VALUES ($1)
	RETURNING id`
	packetLapDataSQL = `
	INSERT INTO PacketLapData (type)
	VALUES ($1)
	RETURNING id`
	packetEventDataSQL = `
	INSERT INTO PacketEventData (type)
	VALUES ($1)
	RETURNING id`
	packetParticipantsDataSQL = `
	INSERT INTO PacketParticipantsData (type)
	VALUES ($1)
	RETURNING id`
	packetCarSetupDataSQL = `
	INSERT INTO PacketCarSetupData (type)
	VALUES ($1)
	RETURNING id`
	packetCarTelemetryDataSQL = `
	INSERT INTO PacketCarTelemetryData (type)
	VALUES ($1)
	RETURNING id`
	packetCarStatusDataSQL = `
	INSERT INTO PacketCarStatusData (type)
	VALUES ($1)
	RETURNING id`
	packetFinalClassificationDataSQL = `
	INSERT INTO PacketFinalClassificationData (type)
	VALUES ($1)
	RETURNING id`
	packetLobbyInfoDataSQL = `
	INSERT INTO PacketLobbyInfoData (type)
	VALUES ($1)
	RETURNING id`
)
