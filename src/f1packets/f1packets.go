package f1packets

// PacketHeader each packet has the header
type PacketHeader struct {
	m_packetFormat     uint16  // 2020
	m_gameMajorVersion uint8   // Game major version - "X.00"
	m_gameMinorVersion uint8   // Game minor version - "1.XX"
	m_packetVersion    uint8   // Version of this packet type, all start from 1
	m_packetId         uint8   // Identifier for the packet type, see below
	m_sessionUID       uint64  // Unique identifier for the session
	m_sessionTime      float32 // Session timestamp
	m_frameIdentifier  uint32  // Identifier for the frame the data was retrieved on
	m_playerCarIndex   uint8   // Index of player's car in the array

	// ADDED IN BETA 2:
	m_secondaryPlayerCarIndex uint8 // Index of secondary player's car in the array (splitscreen); 255 if no second player
}

// CarMotionData motion data for a car
// Frequency: Rate as specified in menus
// Size: 1464 bytes (Packet size updated in Beta 3)
// Version: 1
type CarMotionData struct {
	m_worldPositionX     float32 // World space X position
	m_worldPositionY     float32 // World space Y position
	m_worldPositionZ     float32 // World space Z position
	m_worldVelocityX     float32 // Velocity in world space X
	m_worldVelocityY     float32 // Velocity in world space Y
	m_worldVelocityZ     float32 // Velocity in world space Z
	m_worldForwardDirX   int16   // World space forward X direction (normalised)
	m_worldForwardDirY   int16   // World space forward Y direction (normalised)
	m_worldForwardDirZ   int16   // World space forward Z direction (normalised)
	m_worldRightDirX     int16   // World space right X direction (normalised)
	m_worldRightDirY     int16   // World space right Y direction (normalised)
	m_worldRightDirZ     int16   // World space right Z direction (normalised)
	m_gForceLateral      float32 // Lateral G-Force component
	m_gForceLongitudinal float32 // Longitudinal G-Force component
	m_gForceVertical     float32 // Vertical G-Force component
	m_yaw                float32 // Yaw angle in radians
	m_pitch              float32 // Pitch angle in radians
	m_roll               float32 // Roll angle in radians
}

// PacketMotionData packet gives physics data for all the cars being driven
type PacketMotionData struct {
	m_header        PacketHeader      // Header
	m_carMotionData [22]CarMotionData // Data for all cars on track

	// Extra player car ONLY data
	m_suspensionPosition     [4]float32 // Note: All wheel arrays have the following order:
	m_suspensionVelocity     [4]float32 // RL, RR, FL, FR
	m_suspensionAcceleration [4]float32 // RL, RR, FL, FR
	m_wheelSpeed             [4]float32 // Speed of each wheel
	m_wheelSlip              [4]float32 // Slip ratio for each wheel
	m_localVelocityX         float32    // Velocity in local space
	m_localVelocityY         float32    // Velocity in local space
	m_localVelocityZ         float32    // Velocity in local space
	m_angularVelocityX       float32    // Angular velocity x-component
	m_angularVelocityY       float32    // Angular velocity y-component
	m_angularVelocityZ       float32    // Angular velocity z-component
	m_angularAccelerationX   float32    // Angular velocity x-component
	m_angularAccelerationY   float32    // Angular velocity y-component
	m_angularAccelerationZ   float32    // Angular velocity z-component
	m_frontWheelsAngle       float32    // Current front wheels angle in radians
}

// Session Packet
// The session packet includes details about the current session in progress.

// MarshalZone packet
// Frequency: 2 per second
// Size: 251 bytes (Packet size updated in Beta 3)
// Version: 1
type MarshalZone struct {
	m_zoneStart float32 // Fraction (0..1) of way through the lap the marshal zone starts
	m_zoneFlag  int8    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}

//WeatherForecastSample weather forecast data
type WeatherForecastSample struct {
	// 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1
	// 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2
	// 12 = Time Trial
	m_sessionType uint8
	m_timeOffset  uint8 // Time in minutes the forecast is for

	// Weather - 0 = clear, 1 = light cloud, 2 = overcast
	// 3 = light rain, 4 = heavy rain, 5 = storm
	m_weather          uint8
	m_trackTemperature int8 // Track temp. in degrees celsius
	m_airTemperature   int8 // Air temp. in degrees celsius
}

// PacketSessionData Packet construct for the session
type PacketSessionData struct {
	m_header PacketHeader // Header

	// Weather - 0 = clear, 1 = light cloud, 2 = overcast
	// 3 = light rain, 4 = heavy rain, 5 = storm
	m_weather          uint8
	m_trackTemperature int8   // Track temp. in degrees celsius
	m_airTemperature   int8   // Air temp. in degrees celsius
	m_totalLaps        uint8  // Total number of laps in this race
	m_trackLength      uint16 // Track length in metres

	// 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P
	// 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ
	// 10 = R, 11 = R2, 12 = Time Trial
	m_sessionType uint8
	m_trackId     int8 // -1 for unknown, 0-21 for tracks, see appendix

	// Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2,
	// 3 = F1 Generic
	m_formula             uint8
	m_sessionTimeLeft     uint16          // Time left in session in seconds
	m_sessionDuration     uint16          // Session duration in seconds
	m_pitSpeedLimit       uint8           // Pit speed limit in kilometres per hour
	m_gamePaused          uint8           // Whether the game is paused
	m_isSpectating        uint8           // Whether the player is spectating
	m_spectatorCarIndex   uint8           // Index of the car being spectated
	m_sliProNativeSupport uint8           // SLI Pro support, 0 = inactive, 1 = active
	m_numMarshalZones     uint8           // Number of marshal zones to follow
	m_marshalZones        [21]MarshalZone // List of marshal zones – max 21

	// 0 = no safety car, 1 = full safety car
	// 2 = virtual safety car
	m_safetyCarStatus           uint8
	m_networkGame               uint8                     // 0 = offline, 1 = online
	m_numWeatherForecastSamples uint8                     // Number of weather samples to follow
	m_weatherForecastSamples    [20]WeatherForecastSample // Array of weather forecast samples
}

// Lap Data Packet
// The lap data packet gives details of all the cars in the session.

// LapData provides data on a current and last lap
// Frequency: Rate as specified in menus
// Size: 1190 bytes (Struct updated in Beta 3)
// Version: 1
type LapData struct {
	m_lastLapTime    float32 // Last lap time in seconds
	m_currentLapTime float32 // Current time around the lap in seconds

	//UPDATED in Beta 3:
	m_sector1TimeInMS            uint16  // Sector 1 time in milliseconds
	m_sector2TimeInMS            uint16  // Sector 2 time in milliseconds
	m_bestLapTime                float32 // Best lap time of the session in seconds
	m_bestLapNum                 uint8   // Lap number best time achieved on
	m_bestLapSector1TimeInMS     uint16  // Sector 1 time of best lap in the session in milliseconds
	m_bestLapSector2TimeInMS     uint16  // Sector 2 time of best lap in the session in milliseconds
	m_bestLapSector3TimeInMS     uint16  // Sector 3 time of best lap in the session in milliseconds
	m_bestOverallSector1TimeInMS uint16  // Best overall sector 1 time of the session in milliseconds
	m_bestOverallSector1LapNum   uint8   // Lap number best overall sector 1 time achieved on
	m_bestOverallSector2TimeInMS uint16  // Best overall sector 2 time of the session in milliseconds
	m_bestOverallSector2LapNum   uint8   // Lap number best overall sector 2 time achieved on
	m_bestOverallSector3TimeInMS uint16  // Best overall sector 3 time of the session in milliseconds
	m_bestOverallSector3LapNum   uint8   // Lap number best overall sector 3 time achieved on

	// Distance vehicle is around current lap in metres – could
	// be negative if line hasn’t been crossed yet
	m_lapDistance float32

	// Total distance travelled in session in metres – could
	// be negative if line hasn’t been crossed yet
	m_totalDistance float32

	m_safetyCarDelta    float32 // Delta in seconds for safety car
	m_carPosition       uint8   // Car race position
	m_currentLapNum     uint8   // Current lap number
	m_pitStatus         uint8   // 0 = none, 1 = pitting, 2 = in pit area
	m_sector            uint8   // 0 = sector1, 1 = sector2, 2 = sector3
	m_currentLapInvalid uint8   // Current lap invalid - 0 = valid, 1 = invalid
	m_penalties         uint8   // Accumulated time penalties in seconds to be added

	// Grid position the vehicle started the race in
	// Status of driver - 0 = in garage, 1 = flying lap
	// 2 = in lap, 3 = out lap, 4 = on track
	m_gridPosition uint8
	m_driverStatus uint8

	// Result status - 0 = invalid, 1 = inactive, 2 = active
	// 3 = finished, 4 = disqualified, 5 = not classified
	// 6 = retired
	m_resultStatus uint8
}

// PacketLapData packet construct for lap data
type PacketLapData struct {
	m_header PacketHeader // Header

	m_lapData [22]LapData // Lap data for all cars on track
}

// Event Packet
// This packet gives details of events that happen during the course of a session.

// Frequency: When the event occurs
// Size: 35 bytes (Packet size updated in Beta 3)
// Version: 1

// EventDataDetails Stub interface for Event data details
// The event details packet is different for each type of event.
// Make sure only the correct type is interpreted.
type EventDataDetails interface{}

// FastestLap Event on fastest lap
type FastestLap struct {
	vehicleIdx uint8   // Vehicle index of car achieving fastest lap
	lapTime    float32 // Lap time is in seconds
}

// Retirement Event on car getting retired from a race
type Retirement struct {
	vehicleIdx uint8 // Vehicle index of car retiring
}

// TeamMateInPits Event
type TeamMateInPits struct {
	vehicleIdx uint8 // Vehicle index of team mate
}

// RaceWinner Event
type RaceWinner struct {
	vehicleIdx uint8 // Vehicle index of the race winner
}

// Penalty event providing details on a penalty
type Penalty struct {
	penaltyType      uint8 // Penalty type – see Appendices
	infringementType uint8 // Infringement type – see Appendices
	vehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
	otherVehicleIdx  uint8 // Vehicle index of the other car involved
	time             uint8 // Time gained, or time spent doing action in seconds
	lapNum           uint8 // Lap the penalty occurred on
	placesGained     uint8 // Number of places gained by this
}

// SpeedTrap Event
type SpeedTrap struct {
	vehicleIdx uint8   // Vehicle index of the vehicle triggering speed trap
	speed      float32 // Top speed achieved in kilometres per hour
}

// PacketEventData event packet construct
type PacketEventData struct {
	m_header PacketHeader // Header

	// Event string code, see below
	// Event details - should be interpreted differently
	// for each type
	m_eventStringCode string

	m_eventDetails EventDataDetails
}

// Participants Packet
// This is a list of participants in the race. If the vehicle is controlled by AI, then the name will be the driver name. If this is a multiplayer game, the names will be the Steam Id on PC, or the LAN name if appropriate.

// N.B. on Xbox One, the names will always be the driver name, on PS4 the name will be the LAN name if playing a LAN game, otherwise it will be the driver name.

// The array should be indexed by vehicle index.

// ParticipantData provides information on participant
// Frequency: Every 5 seconds
// Size: 1213 bytes (Packet size updated in Beta 3)
// Version: 1
type ParticipantData struct {
	m_aiControlled uint8 // Whether the vehicle is AI (1) or Human (0) controlled
	m_driverId     uint8 // Driver id - see appendix
	m_teamId       uint8 // Team id - see appendix
	m_raceNumber   uint8 // Race number of the car
	// Nationality of the driver
	// Name of participant in UTF-8 format – null terminated
	// Will be truncated with … (U+2026) if too long
	m_nationality uint8
	m_name        string

	m_yourTelemetry uint8 // The player's UDP setting, 0 = restricted, 1 = public
}

// PacketParticipantsData packet construct for participants data
type PacketParticipantsData struct {
	m_header PacketHeader // Header

	// Number of active cars in the data – should match number of
	// cars on HUD
	m_numActiveCars uint8

	m_participants [22]ParticipantData
}

// Car Setups Packet
// This packet details the car setups for each vehicle in the session. Note that in multiplayer games, other player cars will appear as blank, you will only be able to see your car setup and AI cars.

// CarSetupData provides information on car setup
// Frequency: 2 per second
// Size: 1102 bytes (Packet size updated in Beta 3)
// Version: 1
type CarSetupData struct {
	m_frontWing              uint8   // Front wing aero
	m_rearWing               uint8   // Rear wing aero
	m_onThrottle             uint8   // Differential adjustment on throttle (percentage)
	m_offThrottle            uint8   // Differential adjustment off throttle (percentage)
	m_frontCamber            float32 // Front camber angle (suspension geometry)
	m_rearCamber             float32 // Rear camber angle (suspension geometry)
	m_frontToe               float32 // Front toe angle (suspension geometry)
	m_rearToe                float32 // Rear toe angle (suspension geometry)
	m_frontSuspension        uint8   // Front suspension
	m_rearSuspension         uint8   // Rear suspension
	m_frontAntiRollBar       uint8   // Front anti-roll bar
	m_rearAntiRollBar        uint8   // Front anti-roll bar
	m_frontSuspensionHeight  uint8   // Front ride height
	m_rearSuspensionHeight   uint8   // Rear ride height
	m_brakePressure          uint8   // Brake pressure (percentage)
	m_brakeBias              uint8   // Brake bias (percentage)
	m_rearLeftTyrePressure   float32 // Rear left tyre pressure (PSI)
	m_rearRightTyrePressure  float32 // Rear right tyre pressure (PSI)
	m_frontLeftTyrePressure  float32 // Front left tyre pressure (PSI)
	m_frontRightTyrePressure float32 // Front right tyre pressure (PSI)
	m_ballast                uint8   // Ballast
	m_fuelLoad               float32 // Fuel load
}

// PacketCarSetupData packet construct for session car setup data for all participants
type PacketCarSetupData struct {
	m_header PacketHeader // Header

	m_carSetups [22]CarSetupData
}

// Car Telemetry Packet
// This packet details telemetry for all the cars in the race. It details various values that would be recorded on the car such as speed, throttle application, DRS etc.

// CarTelemetryData provides car telemetry data
// Frequency: Rate as specified in menus
// Size: 1307 bytes (Packet size updated in Beta 3)
// Version: 1
type CarTelemetryData struct {
	m_speed                   uint16     // Speed of car in kilometres per hour
	m_throttle                float32    // Amount of throttle applied (0.0 to 1.0)
	m_steer                   float32    // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	m_brake                   float32    // Amount of brake applied (0.0 to 1.0)
	m_clutch                  uint8      // Amount of clutch applied (0 to 100)
	m_gear                    int8       // Gear selected (1-8, N=0, R=-1)
	m_engineRPM               uint16     // Engine RPM
	m_drs                     uint8      // 0 = off, 1 = on
	m_revLightsPercent        uint8      // Rev lights indicator (percentage)
	m_brakesTemperature       [4]uint16  // Brakes temperature (celsius)
	m_tyresSurfaceTemperature [4]uint8   // Tyres surface temperature (celsius)
	m_tyresInnerTemperature   [4]uint8   // Tyres inner temperature (celsius)
	m_engineTemperature       uint16     // Engine temperature (celsius)
	m_tyresPressure           [4]float32 // Tyres pressure (PSI)
	m_surfaceType             [4]uint8   // Driving surface, see appendices
}

// PacketCarTelemetryData packet construct for car telemetry data
type PacketCarTelemetryData struct {
	m_header PacketHeader // Header

	m_carTelemetryData [22]CarTelemetryData

	// Bit flags specifying which buttons are being pressed
	// currently - see appendices
	m_buttonStatus uint32

	// Added in Beta 3:
	// Index of MFD panel open - 255 = MFD closed
	// Single player, race – 0 = Car setup, 1 = Pits
	// 2 = Damage, 3 =  Engine, 4 = Temperatures
	// May vary depending on game mode
	m_mfdPanelIndex uint8

	m_mfdPanelIndexSecondaryPlayer uint8 // See above

	// Suggested gear for the player (1-8)
	// 0 if no gear suggested
	m_suggestedGear int8
}

// Car Status Packet
// This packet details car statuses for all the cars in the race. It includes values such as the damage readings on the car.

// CarStatusData provides information on car status
// Frequency: Rate as specified in menus
// Size: 1344 bytes (Packet updated in Beta 3)
// Version: 1
type CarStatusData struct {
	m_tractionControl   uint8   // 0 (off) - 2 (high)
	m_antiLockBrakes    uint8   // 0 (off) - 1 (on)
	m_fuelMix           uint8   // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	m_frontBrakeBias    uint8   // Front brake bias (percentage)
	m_pitLimiterStatus  uint8   // Pit limiter status - 0 = off, 1 = on
	m_fuelInTank        float32 // Current fuel mass
	m_fuelCapacity      float32 // Fuel capacity
	m_fuelRemainingLaps float32 // Fuel remaining in terms of laps (value on MFD)
	m_maxRPM            uint16  // Cars max RPM, point of rev limiter
	m_idleRPM           uint16  // Cars idle RPM
	m_maxGears          uint8   // Maximum number of gears
	m_drsAllowed        uint8   // 0 = not allowed, 1 = allowed, -1 = unknown

	// Added in Beta3:
	// 0 = DRS not available, non-zero - DRS will be available
	// in [X] metres
	m_drsActivationDistance uint16

	m_tyresWear [4]uint8 // Tyre wear percentage

	// F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1
	// 7 = inter, 8 = wet
	// F1 Classic - 9 = dry, 10 = wet
	// F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard
	// 15 = wet
	m_actualTyreCompound uint8

	// F1 visual (can be different from actual compound)
	// 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet
	// F1 Classic – same as above
	// F2 – same as above
	m_visualTyreCompound uint8

	m_tyresAgeLaps         uint8    // Age in laps of the current set of tyres
	m_tyresDamage          [4]uint8 // Tyre damage (percentage)
	m_frontLeftWingDamage  uint8    // Front left wing damage (percentage)
	m_frontRightWingDamage uint8    // Front right wing damage (percentage)
	m_rearWingDamage       uint8    // Rear wing damage (percentage)

	// Added Beta 3:
	m_drsFault uint8 // Indicator for DRS fault, 0 = OK, 1 = fault

	m_engineDamage  uint8 // Engine damage (percentage)
	m_gearBoxDamage uint8 // Gear box damage (percentage)

	// -1 = invalid/unknown, 0 = none, 1 = green
	// 2 = blue, 3 = yellow, 4 = red
	m_vehicleFiaFlags int8
	m_ersStoreEnergy  float32 // ERS energy store in Joules

	// ERS deployment mode, 0 = none, 1 = medium
	// 2 = overtake, 3 = hotlap
	m_ersDeployMode uint8

	m_ersHarvestedThisLapMGUK float32 // ERS energy harvested this lap by MGU-K
	m_ersHarvestedThisLapMGUH float32 // ERS energy harvested this lap by MGU-H
	m_ersDeployedThisLap      float32 // ERS energy deployed this lap
}

// PacketCarStatusData packet construct for car status
type PacketCarStatusData struct {
	m_header PacketHeader // Header

	m_carStatusData [22]CarStatusData
}

// Final Classification Packet
// This packet details the final classification at the end of the race, and the data will match with the post race results screen. This is especially useful for multiplayer games where it is not always possible to send lap times on the final frame because of network delay.

// FinalClassificationData final classification data
// Frequency: Once at the end of a race
// Size: 839 bytes (Packet size updated in Beta 3)
// Version: 1
type FinalClassificationData struct {
	m_position     uint8 // Finishing position
	m_numLaps      uint8 // Number of laps completed
	m_gridPosition uint8 // Grid position of the car
	m_points       uint8 // Number of points scored
	m_numPitStops  uint8 // Number of pit stops made

	// Result status - 0 = invalid, 1 = inactive, 2 = active
	// 3 = finished, 4 = disqualified, 5 = not classified
	// 6 = retired
	m_resultStatus     uint8
	m_bestLapTime      float32  // Best lap time of the session in seconds
	m_totalRaceTime    float32  // Total race time in seconds without penalties
	m_penaltiesTime    uint8    // Total penalties accumulated in seconds
	m_numPenalties     uint8    // Number of penalties applied to this driver
	m_numTyreStints    uint8    // Number of tyres stints up to maximum
	m_tyreStintsActual [8]uint8 // Actual tyres used by this driver
	m_tyreStintsVisual [8]uint8 // Visual tyres used by this driver
}

// PacketFinalClassificationData packet construct for final classification data
type PacketFinalClassificationData struct {
	m_header PacketHeader // Header

	m_numCars            uint8 // Number of cars in the final classification
	m_classificationData [22]FinalClassificationData
}

// Lobby Info Packet
// This packet details the players currently in a multiplayer lobby. It details each player’s selected car, any AI involved in the game and also the ready status of each of the participants.

// LobbyInfoData provides lobby information
// Frequency: Two every second when in the lobby
// Size: 1169 bytes (Packet size updated in Beta 3)
// Version: 1
type LobbyInfoData struct {
	m_aiControlled uint8 // Whether the vehicle is AI (1) or Human (0) controlled
	m_teamId       uint8 // Team id - see appendix (255 if no team currently selected)

	// Nationality of the driver
	// Name of participant in UTF-8 format – null terminated
	// Will be truncated with ... (U+2026) if too long
	m_nationality uint8
	m_name        string
	m_readyStatus uint8 // 0 = not ready, 1 = ready, 2 = spectating
}

// PacketLobbyInfoData packet construct for lobby information
type PacketLobbyInfoData struct {
	m_header PacketHeader // Header

	// Packet specific data
	m_numPlayers   uint8 // Number of players in the lobby data
	m_lobbyPlayers [22]LobbyInfoData
}
