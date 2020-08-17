package f12020packets

/*
F1 2020 UDP Telemetry specification
*/

type TestDataPacket struct {
	PacketFormat     uint16
	GameMajorVersion uint8
	GameMinorVersion uint8
}

// PacketHeader each packet has the header
type PacketHeader struct {
	PacketFormat     uint16  // 2020
	GameMajorVersion uint8   // Game major version - "X.00"
	GameMinorVersion uint8   // Game minor version - "1.XX"
	PacketVersion    uint8   // Version of this packet type, all start from 1
	PacketID         uint8   // Identifier for the packet type, see below
	SessionUID       uint64  // Unique identifier for the session
	SessionTime      float32 // Session timestamp
	FrameIdentifier  uint32  // Identifier for the frame the data was retrieved on
	PlayerCarIndex   uint8   // Index of player's car in the array

	// ADDED IN BETA 2:
	SecondaryPlayerCarIndex uint8 // Index of secondary player's car in the array (splitscreen); 255 if no second player
}

// CarMotionData motion data for a car
// Frequency: Rate as specified in menus
// Size: 1464 bytes (Packet size updated in Beta 3)
// Version: 1
type CarMotionData struct {
	WorldPositionX     float32 // World space X position
	WorldPositionY     float32 // World space Y position
	WorldPositionZ     float32 // World space Z position
	WorldVelocityX     float32 // Velocity in world space X
	WorldVelocityY     float32 // Velocity in world space Y
	WorldVelocityZ     float32 // Velocity in world space Z
	WorldForwardDirX   int16   // World space forward X direction (normalised)
	WorldForwardDirY   int16   // World space forward Y direction (normalised)
	WorldForwardDirZ   int16   // World space forward Z direction (normalised)
	WorldRightDirX     int16   // World space right X direction (normalised)
	WorldRightDirY     int16   // World space right Y direction (normalised)
	WorldRightDirZ     int16   // World space right Z direction (normalised)
	GForceLateral      float32 // Lateral G-Force component
	GForceLongitudinal float32 // Longitudinal G-Force component
	GForceVertical     float32 // Vertical G-Force component
	Yaw                float32 // Yaw angle in radians
	Pitch              float32 // Pitch angle in radians
	Roll               float32 // Roll angle in radians
}

// PacketMotionData packet gives physics data for all the cars being driven
type PacketMotionData struct {
	Header        *PacketHeader      // Header
	CarMotionData [22]*CarMotionData // Data for all cars on track

	// Extra player car ONLY data
	SuspensionPosition     [4]float32 // Note: All wheel arrays have the following order:
	SuspensionVelocity     [4]float32 // RL, RR, FL, FR
	SuspensionAcceleration [4]float32 // RL, RR, FL, FR
	WheelSpeed             [4]float32 // Speed of each wheel
	WheelSlip              [4]float32 // Slip ratio for each wheel
	LocalVelocityX         float32    // Velocity in local space
	LocalVelocityY         float32    // Velocity in local space
	LocalVelocityZ         float32    // Velocity in local space
	AngularVelocityX       float32    // Angular velocity x-component
	AngularVelocityY       float32    // Angular velocity y-component
	AngularVelocityZ       float32    // Angular velocity z-component
	AngularAccelerationX   float32    // Angular velocity x-component
	AngularAccelerationY   float32    // Angular velocity y-component
	AngularAccelerationZ   float32    // Angular velocity z-component
	FrontWheelsAngle       float32    // Current front wheels angle in radians
}

// Session Packet
// The session packet includes details about the current session in progress.

// MarshalZone packet
// Frequency: 2 per second
// Size: 251 bytes (Packet size updated in Beta 3)
// Version: 1
type MarshalZone struct {
	ZoneStart float32 // Fraction (0..1) of way through the lap the marshal zone starts
	ZoneFlag  int8    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}

//WeatherForecastSample weather forecast data
type WeatherForecastSample struct {
	// 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1
	// 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2
	// 12 = Time Trial
	SessionType uint8
	TimeOffset  uint8 // Time in minutes the forecast is for

	// Weather - 0 = clear, 1 = light cloud, 2 = overcast
	// 3 = light rain, 4 = heavy rain, 5 = storm
	Weather          uint8
	TrackTemperature int8 // Track temp. in degrees celsius
	AirTemperature   int8 // Air temp. in degrees celsius
}

// PacketSessionData Packet construct for the session
type PacketSessionData struct {
	Header PacketHeader // Header

	// Weather - 0 = clear, 1 = light cloud, 2 = overcast
	// 3 = light rain, 4 = heavy rain, 5 = storm
	Weather          uint8
	TrackTemperature int8   // Track temp. in degrees celsius
	AirTemperature   int8   // Air temp. in degrees celsius
	TotalLaps        uint8  // Total number of laps in this race
	TrackLength      uint16 // Track length in metres

	// 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P
	// 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ
	// 10 = R, 11 = R2, 12 = Time Trial
	SessionType uint8
	TrackID     int8 // -1 for unknown, 0-21 for tracks, see appendix

	// Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2,
	// 3 = F1 Generic
	Formula             uint8
	SessionTimeLeft     uint16          // Time left in session in seconds
	SessionDuration     uint16          // Session duration in seconds
	PitSpeedLimit       uint8           // Pit speed limit in kilometres per hour
	GamePaused          uint8           // Whether the game is paused
	IsSpectating        uint8           // Whether the player is spectating
	SpectatorCarIndex   uint8           // Index of the car being spectated
	SliProNativeSupport uint8           // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones     uint8           // Number of marshal zones to follow
	MarshalZones        [21]MarshalZone // List of marshal zones – max 21

	// 0 = no safety car, 1 = full safety car
	// 2 = virtual safety car
	SafetyCarStatus           uint8
	NetworkGame               uint8                     // 0 = offline, 1 = online
	NumWeatherForecastSamples uint8                     // Number of weather samples to follow
	WeatherForecastSamples    [20]WeatherForecastSample // Array of weather forecast samples
}

// Lap Data Packet
// The lap data packet gives details of all the cars in the session.

// LapData provides data on a current and last lap
// Frequency: Rate as specified in menus
// Size: 1190 bytes (Struct updated in Beta 3)
// Version: 1
type LapData struct {
	LastLapTime    float32 // Last lap time in seconds
	CurrentLapTime float32 // Current time around the lap in seconds

	//UPDATED in Beta 3:
	Sector1TimeInMS            uint16  // Sector 1 time in milliseconds
	Sector2TimeInMS            uint16  // Sector 2 time in milliseconds
	BestLapTime                float32 // Best lap time of the session in seconds
	BestLapNum                 uint8   // Lap number best time achieved on
	BestLapSector1TimeInMS     uint16  // Sector 1 time of best lap in the session in milliseconds
	BestLapSector2TimeInMS     uint16  // Sector 2 time of best lap in the session in milliseconds
	BestLapSector3TimeInMS     uint16  // Sector 3 time of best lap in the session in milliseconds
	BestOverallSector1TimeInMS uint16  // Best overall sector 1 time of the session in milliseconds
	BestOverallSector1LapNum   uint8   // Lap number best overall sector 1 time achieved on
	BestOverallSector2TimeInMS uint16  // Best overall sector 2 time of the session in milliseconds
	BestOverallSector2LapNum   uint8   // Lap number best overall sector 2 time achieved on
	BestOverallSector3TimeInMS uint16  // Best overall sector 3 time of the session in milliseconds
	BestOverallSector3LapNum   uint8   // Lap number best overall sector 3 time achieved on

	// Distance vehicle is around current lap in metres – could
	// be negative if line hasn’t been crossed yet
	LapDistance float32

	// Total distance travelled in session in metres – could
	// be negative if line hasn’t been crossed yet
	TotalDistance float32

	SafetyCarDelta    float32 // Delta in seconds for safety car
	CarPosition       uint8   // Car race position
	CurrentLapNum     uint8   // Current lap number
	PitStatus         uint8   // 0 = none, 1 = pitting, 2 = in pit area
	Sector            uint8   // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid uint8   // Current lap invalid - 0 = valid, 1 = invalid
	Penalties         uint8   // Accumulated time penalties in seconds to be added

	// Grid position the vehicle started the race in
	// Status of driver - 0 = in garage, 1 = flying lap
	// 2 = in lap, 3 = out lap, 4 = on track
	GridPosition uint8
	DriverStatus uint8

	// Result status - 0 = invalid, 1 = inactive, 2 = active
	// 3 = finished, 4 = disqualified, 5 = not classified
	// 6 = retired
	ResultStatus uint8
}

// PacketLapData packet construct for lap data
type PacketLapData struct {
	Header PacketHeader // Header

	LapData [22]LapData // Lap data for all cars on track
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
	Header PacketHeader // Header

	// Event string code, see below
	// Event details - should be interpreted differently
	// for each type
	EventStringCode string

	EventDetails EventDataDetails
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
	AiControlled uint8 // Whether the vehicle is AI (1) or Human (0) controlled
	DriverID     uint8 // Driver id - see appendix
	TeamID       uint8 // Team id - see appendix
	RaceNumber   uint8 // Race number of the car
	// Nationality of the driver
	// Name of participant in UTF-8 format – null terminated
	// Will be truncated with … (U+2026) if too long
	Nationality uint8
	Name        string

	YourTelemetry uint8 // The player's UDP setting, 0 = restricted, 1 = public
}

// PacketParticipantsData packet construct for participants data
type PacketParticipantsData struct {
	Header PacketHeader // Header

	// Number of active cars in the data – should match number of
	// cars on HUD
	NumActiveCars uint8

	Participants [22]ParticipantData
}

// Car Setups Packet
// This packet details the car setups for each vehicle in the session. Note that in multiplayer games, other player cars will appear as blank, you will only be able to see your car setup and AI cars.

// CarSetupData provides information on car setup
// Frequency: 2 per second
// Size: 1102 bytes (Packet size updated in Beta 3)
// Version: 1
type CarSetupData struct {
	FrontWing              uint8   // Front wing aero
	RearWing               uint8   // Rear wing aero
	OnThrottle             uint8   // Differential adjustment on throttle (percentage)
	OffThrottle            uint8   // Differential adjustment off throttle (percentage)
	FrontCamber            float32 // Front camber angle (suspension geometry)
	RearCamber             float32 // Rear camber angle (suspension geometry)
	FrontToe               float32 // Front toe angle (suspension geometry)
	RearToe                float32 // Rear toe angle (suspension geometry)
	FrontSuspension        uint8   // Front suspension
	RearSuspension         uint8   // Rear suspension
	FrontAntiRollBar       uint8   // Front anti-roll bar
	RearAntiRollBar        uint8   // Front anti-roll bar
	FrontSuspensionHeight  uint8   // Front ride height
	RearSuspensionHeight   uint8   // Rear ride height
	BrakePressure          uint8   // Brake pressure (percentage)
	BrakeBias              uint8   // Brake bias (percentage)
	RearLeftTyrePressure   float32 // Rear left tyre pressure (PSI)
	RearRightTyrePressure  float32 // Rear right tyre pressure (PSI)
	FrontLeftTyrePressure  float32 // Front left tyre pressure (PSI)
	FrontRightTyrePressure float32 // Front right tyre pressure (PSI)
	Ballast                uint8   // Ballast
	FuelLoad               float32 // Fuel load
}

// PacketCarSetupData packet construct for session car setup data for all participants
type PacketCarSetupData struct {
	Header PacketHeader // Header

	CarSetups [22]CarSetupData
}

// Car Telemetry Packet
// This packet details telemetry for all the cars in the race. It details various values that would be recorded on the car such as speed, throttle application, DRS etc.

// CarTelemetryData provides car telemetry data
// Frequency: Rate as specified in menus
// Size: 1307 bytes (Packet size updated in Beta 3)
// Version: 1
type CarTelemetryData struct {
	Speed                   uint16     // Speed of car in kilometres per hour
	Throttle                float32    // Amount of throttle applied (0.0 to 1.0)
	Steer                   float32    // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake                   float32    // Amount of brake applied (0.0 to 1.0)
	Clutch                  uint8      // Amount of clutch applied (0 to 100)
	Gear                    int8       // Gear selected (1-8, N=0, R=-1)
	EngineRPM               uint16     // Engine RPM
	Drs                     uint8      // 0 = off, 1 = on
	RevLightsPercent        uint8      // Rev lights indicator (percentage)
	BrakesTemperature       [4]uint16  // Brakes temperature (celsius)
	TyresSurfaceTemperature [4]uint8   // Tyres surface temperature (celsius)
	TyresInnerTemperature   [4]uint8   // Tyres inner temperature (celsius)
	EngineTemperature       uint16     // Engine temperature (celsius)
	TyresPressure           [4]float32 // Tyres pressure (PSI)
	SurfaceType             [4]uint8   // Driving surface, see appendices
}

// PacketCarTelemetryData packet construct for car telemetry data
type PacketCarTelemetryData struct {
	Header PacketHeader // Header

	CarTelemetryData [22]CarTelemetryData

	// Bit flags specifying which buttons are being pressed
	// currently - see appendices
	ButtonStatus uint32

	// Added in Beta 3:
	// Index of MFD panel open - 255 = MFD closed
	// Single player, race – 0 = Car setup, 1 = Pits
	// 2 = Damage, 3 =  Engine, 4 = Temperatures
	// May vary depending on game mode
	MfdPanelIndex uint8

	MfdPanelIndexSecondaryPlayer uint8 // See above

	// Suggested gear for the player (1-8)
	// 0 if no gear suggested
	SuggestedGear int8
}

// Car Status Packet
// This packet details car statuses for all the cars in the race. It includes values such as the damage readings on the car.

// CarStatusData provides information on car status
// Frequency: Rate as specified in menus
// Size: 1344 bytes (Packet updated in Beta 3)
// Version: 1
type CarStatusData struct {
	TractionControl   uint8   // 0 (off) - 2 (high)
	AntiLockBrakes    uint8   // 0 (off) - 1 (on)
	FuelMix           uint8   // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias    uint8   // Front brake bias (percentage)
	PitLimiterStatus  uint8   // Pit limiter status - 0 = off, 1 = on
	FuelInTank        float32 // Current fuel mass
	FuelCapacity      float32 // Fuel capacity
	FuelRemainingLaps float32 // Fuel remaining in terms of laps (value on MFD)
	MaxRPM            uint16  // Cars max RPM, point of rev limiter
	IdleRPM           uint16  // Cars idle RPM
	MaxGears          uint8   // Maximum number of gears
	DrsAllowed        uint8   // 0 = not allowed, 1 = allowed, -1 = unknown

	// Added in Beta3:
	// 0 = DRS not available, non-zero - DRS will be available
	// in [X] metres
	DrsActivationDistance uint16

	TyresWear [4]uint8 // Tyre wear percentage

	// F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1
	// 7 = inter, 8 = wet
	// F1 Classic - 9 = dry, 10 = wet
	// F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard
	// 15 = wet
	ActualTyreCompound uint8

	// F1 visual (can be different from actual compound)
	// 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet
	// F1 Classic – same as above
	// F2 – same as above
	VisualTyreCompound uint8

	TyresAgeLaps         uint8    // Age in laps of the current set of tyres
	TyresDamage          [4]uint8 // Tyre damage (percentage)
	FrontLeftWingDamage  uint8    // Front left wing damage (percentage)
	FrontRightWingDamage uint8    // Front right wing damage (percentage)
	RearWingDamage       uint8    // Rear wing damage (percentage)

	// Added Beta 3:
	DrsFault uint8 // Indicator for DRS fault, 0 = OK, 1 = fault

	EngineDamage  uint8 // Engine damage (percentage)
	GearBoxDamage uint8 // Gear box damage (percentage)

	// -1 = invalid/unknown, 0 = none, 1 = green
	// 2 = blue, 3 = yellow, 4 = red
	VehicleFiaFlags int8
	ErsStoreEnergy  float32 // ERS energy store in Joules

	// ERS deployment mode, 0 = none, 1 = medium
	// 2 = overtake, 3 = hotlap
	ErsDeployMode uint8

	ErsHarvestedThisLapMGUK float32 // ERS energy harvested this lap by MGU-K
	ErsHarvestedThisLapMGUH float32 // ERS energy harvested this lap by MGU-H
	ErsDeployedThisLap      float32 // ERS energy deployed this lap
}

// PacketCarStatusData packet construct for car status
type PacketCarStatusData struct {
	Header PacketHeader // Header

	CarStatusData [22]CarStatusData
}

// Final Classification Packet
// This packet details the final classification at the end of the race, and the data will match with the post race results screen. This is especially useful for multiplayer games where it is not always possible to send lap times on the final frame because of network delay.

// FinalClassificationData final classification data
// Frequency: Once at the end of a race
// Size: 839 bytes (Packet size updated in Beta 3)
// Version: 1
type FinalClassificationData struct {
	Position     uint8 // Finishing position
	NumLaps      uint8 // Number of laps completed
	GridPosition uint8 // Grid position of the car
	Points       uint8 // Number of points scored
	NumPitStops  uint8 // Number of pit stops made

	// Result status - 0 = invalid, 1 = inactive, 2 = active
	// 3 = finished, 4 = disqualified, 5 = not classified
	// 6 = retired
	ResultStatus     uint8
	BestLapTime      float32  // Best lap time of the session in seconds
	TotalRaceTime    float32  // Total race time in seconds without penalties
	PenaltiesTime    uint8    // Total penalties accumulated in seconds
	NumPenalties     uint8    // Number of penalties applied to this driver
	NumTyreStints    uint8    // Number of tyres stints up to maximum
	TyreStintsActual [8]uint8 // Actual tyres used by this driver
	TyreStintsVisual [8]uint8 // Visual tyres used by this driver
}

// PacketFinalClassificationData packet construct for final classification data
type PacketFinalClassificationData struct {
	Header PacketHeader // Header

	NumCars            uint8 // Number of cars in the final classification
	ClassificationData [22]FinalClassificationData
}

// Lobby Info Packet
// This packet details the players currently in a multiplayer lobby. It details each player’s selected car, any AI involved in the game and also the ready status of each of the participants.

// LobbyInfoData provides lobby information
// Frequency: Two every second when in the lobby
// Size: 1169 bytes (Packet size updated in Beta 3)
// Version: 1
type LobbyInfoData struct {
	AiControlled uint8 // Whether the vehicle is AI (1) or Human (0) controlled
	TeamID       uint8 // Team id - see appendix (255 if no team currently selected)

	// Nationality of the driver
	// Name of participant in UTF-8 format – null terminated
	// Will be truncated with ... (U+2026) if too long
	Nationality uint8
	Name        string
	ReadyStatus uint8 // 0 = not ready, 1 = ready, 2 = spectating
}

// PacketLobbyInfoData packet construct for lobby information
type PacketLobbyInfoData struct {
	Header PacketHeader // Header

	// Packet specific data
	NumPlayers   uint8 // Number of players in the lobby data
	LobbyPlayers [22]LobbyInfoData
}
