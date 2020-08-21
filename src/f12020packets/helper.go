package f12020packets

import (
	"encoding/binary"
	"fmt"
	"math"
)

const (
	packetHeaderLength                  = 24
	packetMotionDataLength              = 1464 - packetHeaderLength
	packetSessionDataLength             = 251 - packetHeaderLength
	packetLapDataLength                 = 1190 - packetHeaderLength
	packetEventDataLength               = 35 - packetHeaderLength
	packetParticipantsDataLength        = 1213 - packetHeaderLength
	packetCarSetupDataLength            = 1102 - packetHeaderLength
	packetCarTelemetryDataLength        = 1307 - packetHeaderLength
	packetCarStatusDataLength           = 1344 - packetHeaderLength
	packetFinalClassificationDataLength = 839 - packetHeaderLength
	packetLobbyInfoDataLength           = 1169 - packetHeaderLength

	carMotionDataLength         = 60
	marshalZoneLength           = 5
	weatherForecastSampleLength = 5
	lapDataLength               = 53
	fastestLapLength            = 5
	retirementLength            = 1
	teamMateInPitsLength        = 1
	raceWinnerLength            = 1
	penaltyLength               = 7
	speedTrapLength             = 5
	participantDataLength       = 54
	carSetupDataLength          = 49
	carTelemetryDataLength      = 58
	carStatusDataLength         = 60
)

/*
	TODO: Look into gob decoding instead
*/
func convertToFloat32(data []byte) float32 {
	// if len(data) > 4 {
	// 	return 0, fmt.Errorf("Wrong size data provided, expected %d was %d", 4, len(data))
	// }
	return math.Float32frombits(binary.LittleEndian.Uint32(data))
}

func convertTo4LengthFloat32Array(data []byte) [4]float32 {
	var result [4]float32
	for i := 0; i < 4; i++ {
		startIndex := 0 + (i * 4)
		endIndex := startIndex + 4
		value := convertToFloat32(data[startIndex:endIndex])
		result[i] = value
	}
	return result
}

func convertToint16(data []byte) int16 {
	var value int16
	value |= int16(data[0])
	value |= int16(data[1]) << 8
	return value
}

func convertTo4LengthUint16Array(data []byte) [4]uint16 {
	var result [4]uint16
	for i := 0; i < 4; i++ {
		startIndex := 0 + (i * 4)
		endIndex := startIndex + 4
		value := binary.LittleEndian.Uint16(data[startIndex:endIndex])
		result[i] = value
	}
	return result
}

func convertTo4LengthUint8Array(data []byte) [4]uint8 {
	var result [4]uint8
	for i := 0; i < 4; i++ {
		result[i] = uint8(data[i])
	}
	return result
}

// 23 bytes
func ToPacketHeader(data []byte) (*PacketHeader, error) {
	if len(data) != packetHeaderLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", packetHeaderLength, len(data))
	}
	header := &PacketHeader{
		PacketFormat:            binary.LittleEndian.Uint16(data[0:2]),
		GameMajorVersion:        uint8(data[2]),
		GameMinorVersion:        uint8(data[3]),
		PacketVersion:           uint8(data[4]),
		PacketID:                uint8(data[5]),
		SessionUID:              binary.LittleEndian.Uint64(data[6:14]),
		SessionTime:             convertToFloat32(data[14:18]),
		FrameIdentifier:         binary.LittleEndian.Uint32(data[18:22]),
		PlayerCarIndex:          uint8(22),
		SecondaryPlayerCarIndex: uint8(23),
	}

	return header, nil
}

// 60 bytes
func ToCarMotionData(data []byte) (*CarMotionData, error) {
	if len(data) != carMotionDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", carMotionDataLength, len(data))
	}
	motionData := &CarMotionData{
		WorldPositionX:     convertToFloat32(data[0:4]),
		WorldPositionY:     convertToFloat32(data[4:8]),
		WorldPositionZ:     convertToFloat32(data[8:12]),
		WorldVelocityX:     convertToFloat32(data[12:16]),
		WorldVelocityY:     convertToFloat32(data[16:20]),
		WorldVelocityZ:     convertToFloat32(data[20:24]),
		WorldForwardDirX:   convertToint16(data[24:26]),
		WorldForwardDirY:   convertToint16(data[26:28]),
		WorldForwardDirZ:   convertToint16(data[28:30]),
		WorldRightDirX:     convertToint16(data[30:32]),
		WorldRightDirY:     convertToint16(data[32:34]),
		WorldRightDirZ:     convertToint16(data[34:36]),
		GForceLateral:      convertToFloat32(data[36:40]),
		GForceLongitudinal: convertToFloat32(data[40:44]),
		GForceVertical:     convertToFloat32(data[44:48]),
		Yaw:                convertToFloat32(data[48:52]),
		Pitch:              convertToFloat32(data[52:56]),
		Roll:               convertToFloat32(data[56:60]),
	}

	return motionData, nil
}

// 1464 bytes
// 23 bytes header
// 1320 bytes car motion data
// 121 bytes the rest
func ToPacketMotionData(data []byte, header *PacketHeader) (*PacketMotionData, error) {
	if len(data) != packetMotionDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", packetMotionDataLength, len(data))
	}

	// 1320 bytes in total
	var motionData [22]CarMotionData
	for i := 0; i < 22; i++ {
		startIndex := 0 + (i * carMotionDataLength)
		endIndex := startIndex + carMotionDataLength
		// Swallow any exceptions for now
		payload, _ := ToCarMotionData(data[startIndex:endIndex])
		motionData[i] = *payload
	}

	// Construct packet and decode the rest of the data
	packet := &PacketMotionData{
		Header:                 header,
		CarMotionData:          motionData,
		SuspensionPosition:     convertTo4LengthFloat32Array(data[1319:1335]),
		SuspensionVelocity:     convertTo4LengthFloat32Array(data[1335:1351]),
		SuspensionAcceleration: convertTo4LengthFloat32Array(data[1351:1367]),
		WheelSpeed:             convertTo4LengthFloat32Array(data[1367:1383]),
		WheelSlip:              convertTo4LengthFloat32Array(data[1383:1399]),
		LocalVelocityX:         convertToFloat32(data[1399:1403]),
		LocalVelocityY:         convertToFloat32(data[1403:1407]),
		LocalVelocityZ:         convertToFloat32(data[1407:1411]),
		AngularVelocityX:       convertToFloat32(data[1411:1415]),
		AngularVelocityY:       convertToFloat32(data[1415:1419]),
		AngularVelocityZ:       convertToFloat32(data[1419:1423]),
		AngularAccelerationX:   convertToFloat32(data[1423:1427]),
		AngularAccelerationY:   convertToFloat32(data[1427:1431]),
		AngularAccelerationZ:   convertToFloat32(data[1431:1435]),
		FrontWheelsAngle:       convertToFloat32(data[1435:1439]),
	}

	return packet, nil
}

func ToMarshalZone(data []byte) (*MarshalZone, error) {
	if len(data) != marshalZoneLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", marshalZoneLength, len(data))
	}
	marshalZone := &MarshalZone{
		ZoneStart: convertToFloat32(data[0:4]),
		ZoneFlag:  int8(data[4]),
	}

	return marshalZone, nil
}

func ToWeatherForecastSample(data []byte) (*WeatherForecastSample, error) {
	if len(data) != weatherForecastSampleLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", weatherForecastSampleLength, len(data))
	}
	weatherForecast := &WeatherForecastSample{
		SessionType:      uint8(data[0]),
		TimeOffset:       uint8(data[1]),
		Weather:          uint8(data[2]),
		TrackTemperature: int8(data[3]),
		AirTemperature:   int8(data[4]),
	}

	return weatherForecast, nil
}

func ToPacketSessionData(data []byte, header *PacketHeader) (*PacketSessionData, error) {
	if len(data) != packetSessionDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", packetSessionDataLength, len(data))
	}

	// 105 bytes in total
	var marshalZones [21]MarshalZone
	for i := 0; i < 21; i++ {
		startIndex := 19 + (i * marshalZoneLength)
		endIndex := startIndex + marshalZoneLength

		payload, _ := ToMarshalZone(data[startIndex:endIndex])
		marshalZones[i] = *payload
	}

	// 100 bytes in total
	var weatherForecastSamples [20]WeatherForecastSample
	for i := 0; i < 20; i++ {
		startIndex := 127 + (i * weatherForecastSampleLength)
		endIndex := startIndex + weatherForecastSampleLength

		payload, _ := ToWeatherForecastSample(data[startIndex:endIndex])
		weatherForecastSamples[i] = *payload
	}

	packet := &PacketSessionData{
		Header:                    header,
		Weather:                   uint8(data[0]),
		TrackTemperature:          int8(data[1]),
		AirTemperature:            int8(data[2]),
		TotalLaps:                 uint8(data[3]),
		TrackLength:               binary.LittleEndian.Uint16(data[4:6]),
		SessionType:               uint8(data[6]),
		TrackID:                   int8(data[7]),
		Formula:                   uint8(data[8]),
		SessionTimeLeft:           binary.LittleEndian.Uint16(data[9:11]),
		SessionDuration:           binary.LittleEndian.Uint16(data[11:13]),
		PitSpeedLimit:             uint8(data[13]),
		GamePaused:                uint8(data[14]),
		IsSpectating:              uint8(data[15]),
		SpectatorCarIndex:         uint8(data[16]),
		SliProNativeSupport:       uint8(data[17]),
		NumMarshalZones:           uint8(data[18]),
		MarshalZones:              marshalZones,
		SafetyCarStatus:           uint8(data[124]),
		NetworkGame:               uint8(data[125]),
		NumWeatherForecastSamples: uint8(data[126]),
		WeatherForecastSamples:    weatherForecastSamples,
	}
	return packet, nil
}

func ToLapData(data []byte) (*LapData, error) {
	if len(data) != lapDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", lapDataLength, len(data))
	}

	lapData := &LapData{
		LastLapTime:                convertToFloat32(data[0:4]),
		CurrentLapTime:             convertToFloat32(data[4:8]),
		Sector1TimeInMS:            binary.LittleEndian.Uint16(data[8:10]),
		Sector2TimeInMS:            binary.BigEndian.Uint16(data[10:12]),
		BestLapTime:                convertToFloat32(data[12:16]),
		BestLapNum:                 uint8(data[16]),
		BestLapSector1TimeInMS:     binary.LittleEndian.Uint16(data[17:19]),
		BestLapSector2TimeInMS:     binary.LittleEndian.Uint16(data[19:21]),
		BestLapSector3TimeInMS:     binary.LittleEndian.Uint16(data[21:23]),
		BestOverallSector1TimeInMS: binary.LittleEndian.Uint16(data[23:25]),
		BestOverallSector1LapNum:   uint8(data[25]),
		BestOverallSector2TimeInMS: binary.LittleEndian.Uint16(data[26:28]),
		BestOverallSector2LapNum:   uint8(data[28]),
		BestOverallSector3TimeInMS: binary.LittleEndian.Uint16(data[29:31]),
		BestOverallSector3LapNum:   uint8(data[32]),
		LapDistance:                convertToFloat32(data[33:37]),
		TotalDistance:              convertToFloat32(data[37:41]),
		SafetyCarDelta:             convertToFloat32(data[41:45]),
		CarPosition:                uint8(data[45]),
		CurrentLapNum:              uint8(data[46]),
		PitStatus:                  uint8(data[47]),
		Sector:                     uint8(data[48]),
		CurrentLapInvalid:          uint8(data[49]),
		Penalties:                  uint8(data[50]),
		GridPosition:               uint8(data[51]),
		DriverStatus:               uint8(data[52]),
		ResultStatus:               uint8(data[53]),
	}

	return lapData, nil
}

func ToPacketLapData(data []byte, header *PacketHeader) (*PacketLapData, error) {
	if len(data) != packetLapDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", packetLapDataLength, len(data))
	}

	// 100 bytes in total
	var lapDatas [22]LapData
	for i := 0; i < 22; i++ {
		startIndex := 0 + (i * lapDataLength)
		endIndex := startIndex + lapDataLength

		payload, _ := ToLapData(data[startIndex:endIndex])
		lapDatas[i] = *payload
	}

	packet := &PacketLapData{
		Header:  header,
		LapData: lapDatas,
	}
	return packet, nil
}

func ToFastestLap(data []byte) (*FastestLap, error) {
	if len(data) != fastestLapLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", fastestLapLength, len(data))
	}

	fastestLap := &FastestLap{
		VehicleIdx: uint8(data[0]),
		LapTime:    convertToFloat32(data[1:5]),
	}

	return fastestLap, nil
}

func ToRetirement(data []byte) (*Retirement, error) {
	if len(data) != retirementLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", retirementLength, len(data))
	}

	retirement := &Retirement{
		VehicleIdx: uint8(data[0]),
	}

	return retirement, nil
}

func ToTeamMateInPits(data []byte) (*TeamMateInPits, error) {
	if len(data) != teamMateInPitsLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", teamMateInPitsLength, len(data))
	}

	teamMateInPits := &TeamMateInPits{
		VehicleIdx: uint8(data[0]),
	}

	return teamMateInPits, nil
}

func ToRaceWinner(data []byte) (*RaceWinner, error) {
	if len(data) != raceWinnerLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", raceWinnerLength, len(data))
	}

	raceWinner := &RaceWinner{
		VehicleIdx: uint8(data[0]),
	}

	return raceWinner, nil
}

func ToPenalty(data []byte) (*Penalty, error) {
	if len(data) != penaltyLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", penaltyLength, len(data))
	}

	penalty := &Penalty{
		PenaltyType:      uint8(data[0]),
		InfringementType: uint8(data[1]),
		VehicleIdx:       uint8(data[2]),
		OtherVehicleIdx:  uint8(data[3]),
		Time:             uint8(data[4]),
		LapNum:           uint8(data[5]),
		PlacesGained:     uint8(data[6]),
	}

	return penalty, nil
}

func ToSpeedTrap(data []byte) (*SpeedTrap, error) {
	if len(data) != speedTrapLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", speedTrapLength, len(data))
	}

	speedTrap := &SpeedTrap{
		VehicleIdx: uint8(data[0]),
		Speed:      convertToFloat32(data[1:5]),
	}

	return speedTrap, nil
}

func ToPacketEventData(data []byte, header *PacketHeader) (*PacketEventData, error) {
	if len(data) != packetEventDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", packetEventDataLength, len(data))
	}
	packet := &PacketEventData{
		Header:          header,
		EventStringCode: string(data[0:4]),
	}

	return packet, nil
}

func ToParticipantData(data []byte) (*ParticipantData, error) {
	if len(data) != participantDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", participantDataLength, len(data))
	}

	participantData := &ParticipantData{
		AiControlled:  uint8(data[0]),
		DriverID:      uint8(data[1]),
		TeamID:        uint8(data[2]),
		RaceNumber:    uint8(data[3]),
		Nationality:   uint8(data[4]),
		Name:          string(data[5:53]),
		YourTelemetry: uint8(data[54]),
	}

	return participantData, nil
}

func ToPacketParticipantsData(data []byte, header *PacketHeader) (*PacketParticipantsData, error) {
	if len(data) != packetParticipantsDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", packetParticipantsDataLength, len(data))
	}

	// 1188 bytes in total
	var participantData [22]ParticipantData
	for i := 0; i < 22; i++ {
		startIndex := 1 + (i * participantDataLength)
		endIndex := startIndex + participantDataLength

		payload, _ := ToParticipantData(data[startIndex:endIndex])
		participantData[i] = *payload
	}

	packet := &PacketParticipantsData{
		Header:        header,
		NumActiveCars: uint8(data[0]),
		Participants:  participantData,
	}
	return packet, nil
}

func ToCarSetupData(data []byte) (*CarSetupData, error) {
	if len(data) != carSetupDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", carSetupDataLength, len(data))
	}

	carSetupData := &CarSetupData{
		FrontWing:              uint8(data[0]),
		RearWing:               uint8(data[1]),
		OnThrottle:             uint8(data[2]),
		OffThrottle:            uint8(data[3]),
		FrontCamber:            convertToFloat32(data[4:8]),
		RearCamber:             convertToFloat32(data[8:12]),
		FrontToe:               convertToFloat32(data[12:16]),
		RearToe:                convertToFloat32(data[16:20]),
		FrontSuspension:        uint8(data[20]),
		RearSuspension:         uint8(data[21]),
		FrontAntiRollBar:       uint8(data[22]),
		RearAntiRollBar:        uint8(data[23]),
		FrontSuspensionHeight:  uint8(data[24]),
		RearSuspensionHeight:   uint8(data[25]),
		BrakePressure:          uint8(data[26]),
		BrakeBias:              uint8(data[27]),
		RearLeftTyrePressure:   convertToFloat32(data[28:32]),
		RearRightTyrePressure:  convertToFloat32(data[32:36]),
		FrontLeftTyrePressure:  convertToFloat32(data[36:40]),
		FrontRightTyrePressure: convertToFloat32(data[40:44]),
		Ballast:                uint8(data[44]),
		FuelLoad:               convertToFloat32(data[45:49]),
	}

	return carSetupData, nil
}

func ToPacketCarSetupData(data []byte, header *PacketHeader) (*PacketCarSetupData, error) {
	if len(data) != packetCarSetupDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", packetCarSetupDataLength, len(data))
	}

	// 1188 bytes in total
	var carSetupData [22]CarSetupData
	for i := 0; i < 22; i++ {
		startIndex := 0 + (i * carSetupDataLength)
		endIndex := startIndex + carSetupDataLength

		payload, _ := ToCarSetupData(data[startIndex:endIndex])
		carSetupData[i] = *payload
	}

	packet := &PacketCarSetupData{
		Header: header,
	}
	return packet, nil
}

func ToCarTelemetryData(data []byte) (*CarTelemetryData, error) {
	if len(data) != carTelemetryDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", carTelemetryDataLength, len(data))
	}

	carTelemetryData := &CarTelemetryData{
		Speed:                   binary.LittleEndian.Uint16(data[0:2]),
		Throttle:                convertToFloat32(data[2:6]),
		Steer:                   convertToFloat32(data[6:10]),
		Brake:                   convertToFloat32(data[10:14]),
		Clutch:                  uint8(data[14]),
		Gear:                    int8(data[15]),
		EngineRPM:               binary.LittleEndian.Uint16(data[16:18]),
		Drs:                     uint8(data[18]),
		RevLightsPercent:        uint8(data[19]),
		BrakesTemperature:       convertTo4LengthUint16Array(data[20:28]),
		TyresSurfaceTemperature: convertTo4LengthUint8Array(data[28:32]),
		TyresInnerTemperature:   convertTo4LengthUint8Array(data[32:36]),
		EngineTemperature:       binary.LittleEndian.Uint16(data[36:38]),
		TyresPressure:           convertTo4LengthFloat32Array(data[38:54]),
		SurfaceType:             convertTo4LengthUint8Array(data[54:58]),
	}

	return carTelemetryData, nil
}

func ToPacketCarTelemetryData(data []byte, header *PacketHeader) (*PacketCarTelemetryData, error) {
	if len(data) != packetCarTelemetryDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", packetCarTelemetryDataLength, len(data))
	}

	// 1276 bytes in total
	var carTelemetryData [22]CarTelemetryData
	for i := 0; i < 22; i++ {
		startIndex := 0 + (i * carTelemetryDataLength)
		endIndex := startIndex + carTelemetryDataLength

		payload, _ := ToCarTelemetryData(data[startIndex:endIndex])
		carTelemetryData[i] = *payload
	}

	packet := &PacketCarTelemetryData{
		Header:                       header,
		CarTelemetryData:             carTelemetryData,
		ButtonStatus:                 binary.LittleEndian.Uint32(data[1276:1280]),
		MfdPanelIndex:                uint8(data[1280]),
		MfdPanelIndexSecondaryPlayer: uint8(data[1281]),
		SuggestedGear:                int8(data[1282]),
	}
	return packet, nil
}

func ToCarStatusData(data []byte) (*CarStatusData, error) {
	if len(data) != carStatusDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", carStatusDataLength, len(data))
	}

	carStatusData := &CarStatusData{
		TractionControl:         uint8(data[0]),
		AntiLockBrakes:          uint8(data[1]),
		FuelMix:                 uint8(data[2]),
		FrontBrakeBias:          uint8(data[3]),
		PitLimiterStatus:        uint8(data[4]),
		FuelInTank:              convertToFloat32(data[5:9]),
		FuelCapacity:            convertToFloat32(data[9:13]),
		FuelRemainingLaps:       convertToFloat32(data[13:17]),
		MaxRPM:                  binary.LittleEndian.Uint16(data[17:19]),
		IdleRPM:                 binary.LittleEndian.Uint16(data[19:21]),
		MaxGears:                uint8(data[21]),
		DrsAllowed:              uint8(data[22]),
		DrsActivationDistance:   binary.LittleEndian.Uint16(data[23:25]),
		TyresWear:               convertTo4LengthUint8Array(data[25:29]),
		ActualTyreCompound:      uint8(data[29]),
		VisualTyreCompound:      uint8(data[30]),
		TyresAgeLaps:            uint8(data[31]),
		TyresDamage:             convertTo4LengthUint8Array(data[32:36]),
		FrontLeftWingDamage:     uint8(data[36]),
		FrontRightWingDamage:    uint8(data[37]),
		RearWingDamage:          uint8(data[38]),
		DrsFault:                uint8(data[39]),
		EngineDamage:            uint8(data[40]),
		GearBoxDamage:           uint8(data[41]),
		VehicleFiaFlags:         int8(data[42]),
		ErsStoreEnergy:          convertToFloat32(data[43:47]),
		ErsDeployMode:           uint8(data[47]),
		ErsHarvestedThisLapMGUK: convertToFloat32(data[48:52]),
		ErsHarvestedThisLapMGUH: convertToFloat32(data[52:56]),
		ErsDeployedThisLap:      convertToFloat32(data[56:60]),
	}

	return carStatusData, nil
}

func ToPacketCarStatusData(data []byte, header *PacketHeader) (*PacketCarStatusData, error) {
	if len(data) != packetCarStatusDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", packetCarStatusDataLength, len(data))
	}

	// 1320 bytes in total
	var carStatusData [22]CarStatusData
	for i := 0; i < 22; i++ {
		startIndex := 0 + (i * carStatusDataLength)
		endIndex := startIndex + carStatusDataLength

		payload, _ := ToCarStatusData(data[startIndex:endIndex])
		carStatusData[i] = *payload
	}

	packet := &PacketCarStatusData{
		Header:        header,
		CarStatusData: carStatusData,
	}
	return packet, nil
}

func ToPacketFinalClassificationData(data []byte, header *PacketHeader) (*PacketFinalClassificationData, error) {
	if len(data) != packetFinalClassificationDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", packetFinalClassificationDataLength, len(data))
	}

	packet := &PacketFinalClassificationData{
		Header: header,
	}
	return packet, nil
}

func ToPacketLobbyInfoData(data []byte, header *PacketHeader) (*PacketLobbyInfoData, error) {
	if len(data) != packetLobbyInfoDataLength {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", packetLobbyInfoDataLength, len(data))
	}

	packet := &PacketLobbyInfoData{
		Header: header,
	}
	return packet, nil
}
