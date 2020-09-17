package clients

import (
	"context"
	"fmt"
	"main/f12020packets"
	"main/helpers"

	uuid "github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (client PostgreClient) insert(ctx context.Context, packetType string, args []interface{}) error {
	logrus.Debugln("Looking up SQL statement for: ", packetType)
	sqlStatement, err := getSQLStatement(packetType)
	if err != nil {
		return err
	}

	_, err = client.Database.ExecContext(ctx, *sqlStatement, args...)

	helpers.LogIfError(err)
	return err
}

func getSQLStatement(packetType string) (*string, error) {
	var sql string
	switch packetType {
	case PacketHeader:
		sql = packetHeaderSQL
	case PacketMotionData:
		sql = packetMotionDataSQL
	case CarMotionData:
		sql = carMotionDataSQL
	case PacketSessionData:
		sql = packetSessionDataSQL
	case MarshalZone:
		sql = marshalZoneSQL
	case WeatherForecastSample:
		sql = weatherForecastSampleSQL
	case PacketLapData:
		sql = packetLapDataSQL
	case LapData:
		sql = lapDataSQL
	case PacketEventData:
		sql = packetEventDataSQL
	case FastestLap:
		sql = fastestLapSQL
	case Retirement:
		sql = retirementSQL
	case TeammateInPits:
		sql = teammateInPitsSQL
	case RaceWinner:
		sql = raceWinnerSQL
	case Penalty:
		sql = penaltySQL
	case SpeedTrap:
		sql = speedTrapSQL
	case PacketParticipantsData:
		sql = packetParticipantsDataSQL
	case ParticipantData:
		sql = participantDataSQL
	case PacketCarSetupData:
		sql = packetCarSetupDataSQL
	case CarSetupData:
		sql = carSetupDataSQL
	case PacketCarTelemetryData:
		sql = packetCarTelemetryDataSQL
	case CarTelemetryData:
		sql = carTelemetryDataSQL
	case PacketCarStatusData:
		sql = packetCarStatusDataSQL
	case CarStatusData:
		sql = carStatusDataSQL
	case PacketFinalClassificationData:
		sql = packetFinalClassificationDataSQL
	case FinalClassificationData:
		sql = finalClassificationDataSQL
	case PacketLobbyInfoData:
		sql = packetLobbyInfoDataSQL
	case LobbyInfoData:
		sql = lobbyInfoDataSQL
	default:
		return nil, fmt.Errorf("PacketType: %s is not valid", packetType)
	}
	return &sql, nil
}

func (client PostgreClient) insertPacketHeader(ctx context.Context, header *f12020packets.PacketHeader) (uuid.UUID, error) {
	id := uuid.New()

	args := []interface{}{
		id,
		header.PacketFormat,
		header.GameMajorVersion,
		header.GameMinorVersion,
		header.PacketVersion,
		header.PacketID,
		header.SessionUID,
		header.SessionTime,
		header.FrameIdentifier,
		header.PlayerCarIndex,
		header.SecondaryPlayerCarIndex,
	}
	return id, client.insert(ctx, PacketHeader, args)
}

func (client PostgreClient) insertPacketMotionData(ctx context.Context, data f12020packets.PacketMotionData, headerID uuid.UUID) (uuid.UUID, error) {
	id := uuid.New()

	args := []interface{}{
		id,
		headerID,
		data.SuspensionPositionRL,
		data.SuspensionPositionRR,
		data.SuspensionPositionFL,
		data.SuspensionPositionFR,
		data.SuspensionVelocityRL,
		data.SuspensionVelocityRR,
		data.SuspensionVelocityFL,
		data.SuspensionVelocityFR,
		data.SuspensionAccelerationRL,
		data.SuspensionAccelerationRR,
		data.SuspensionAccelerationFL,
		data.SuspensionAccelerationFR,
		data.WheelSpeedRL,
		data.WheelSpeedRR,
		data.WheelSpeedFL,
		data.WheelSpeedFR,
		data.WheelSlipRL,
		data.WheelSlipRR,
		data.WheelSlipFL,
		data.WheelSlipFR,
		data.LocalVelocityX,
		data.LocalVelocityY,
		data.LocalVelocityZ,
		data.AngularVelocityX,
		data.AngularVelocityY,
		data.AngularVelocityZ,
		data.AngularAccelerationX,
		data.AngularAccelerationY,
		data.AngularAccelerationZ,
		data.FrontWheelsAngle,
	}
	return id, client.insert(ctx, PacketMotionData, args)
}

func (client PostgreClient) insertCarMotionData(ctx context.Context, data f12020packets.CarMotionData, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.WorldPositionX,
		data.WorldPositionY,
		data.WorldPositionZ,
		data.WorldVelocityX,
		data.WorldVelocityY,
		data.WorldVelocityZ,
		data.WorldForwardDirX,
		data.WorldForwardDirY,
		data.WorldForwardDirZ,
		data.WorldRightDirX,
		data.WorldRightDirY,
		data.WorldRightDirZ,
		data.GForceLateral,
		data.GForceLongitudinal,
		data.GForceVertical,
		data.Yaw,
		data.Pitch,
		data.Roll,
	}
	return client.insert(ctx, CarMotionData, args)
}

func (client PostgreClient) insertPacketSessionData(ctx context.Context, data f12020packets.PacketSessionData, headerID uuid.UUID) (uuid.UUID, error) {
	id := uuid.New()

	args := []interface{}{
		id,
		headerID,
		data.Weather,
		data.TrackTemperature,
		data.AirTemperature,
		data.TotalLaps,
		data.TrackLength,
		data.SessionType,
		data.TrackID,
		data.Formula,
		data.SessionTimeLeft,
		data.SessionDuration,
		data.PitSpeedLimit,
		data.GamePaused,
		data.IsSpectating,
		data.SpectatorCarIndex,
		data.SliProNativeSupport,
		data.NumMarshalZones,
		data.SafetyCarStatus,
		data.NetworkGame,
		data.NumWeatherForecastSamples,
	}
	return id, client.insert(ctx, PacketSessionData, args)
}

func (client PostgreClient) insertMarshalZone(ctx context.Context, data f12020packets.MarshalZone, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.ZoneStart,
		data.ZoneFlag,
	}
	return client.insert(ctx, MarshalZone, args)
}

func (client PostgreClient) insertWeatherForecastSample(ctx context.Context, data f12020packets.WeatherForecastSample, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.SessionType,
		data.TimeOffset,
		data.Weather,
		data.TrackTemperature,
		data.AirTemperature,
	}
	return client.insert(ctx, WeatherForecastSample, args)
}

func (client PostgreClient) insertPacketLapData(ctx context.Context, data f12020packets.PacketLapData, headerID uuid.UUID) (uuid.UUID, error) {
	id := uuid.New()

	args := []interface{}{
		id,
		headerID,
	}
	return id, client.insert(ctx, PacketLapData, args)
}

func (client PostgreClient) insertLapData(ctx context.Context, data f12020packets.LapData, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.LastLapTime,
		data.CurrentLapTime,
		data.Sector1TimeInMS,
		data.Sector2TimeInMS,
		data.BestLapTime,
		data.BestLapNum,
		data.BestLapSector1TimeInMS,
		data.BestLapSector2TimeInMS,
		data.BestLapSector3TimeInMS,
		data.BestOverallSector1TimeInMS,
		data.BestOverallSector1LapNum,
		data.BestOverallSector2TimeInMS,
		data.BestOverallSector2LapNum,
		data.BestOverallSector3TimeInMS,
		data.BestOverallSector3LapNum,
		data.LapDistance,
		data.TotalDistance,
		data.SafetyCarDelta,
		data.CarPosition,
		data.CurrentLapNum,
		data.PitStatus,
		data.Sector,
		data.CurrentLapInvalid,
		data.Penalties,
		data.GridPosition,
		data.DriverStatus,
		data.ResultStatus,
	}

	return client.insert(ctx, LapData, args)
}

func (client PostgreClient) insertPacketEventData(ctx context.Context, data f12020packets.PacketEventData, headerID uuid.UUID) (uuid.UUID, error) {
	id := uuid.New()

	args := []interface{}{
		id,
		headerID,
		data.EventStringCode,
	}

	return id, client.insert(ctx, PacketEventData, args)
}

func (client PostgreClient) insertFastestLap(ctx context.Context, data f12020packets.FastestLap, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.VehicleIdx,
		data.LapTime,
	}

	return client.insert(ctx, FastestLap, args)
}

func (client PostgreClient) insertRetirement(ctx context.Context, data f12020packets.Retirement, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.VehicleIdx,
	}

	return client.insert(ctx, Retirement, args)
}

func (client PostgreClient) insertTeamMateInPits(ctx context.Context, data f12020packets.TeamMateInPits, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.VehicleIdx,
	}

	return client.insert(ctx, TeammateInPits, args)
}

func (client PostgreClient) insertRaceWinner(ctx context.Context, data f12020packets.RaceWinner, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.VehicleIdx,
	}

	return client.insert(ctx, RaceWinner, args)
}

func (client PostgreClient) insertPenalty(ctx context.Context, data f12020packets.Penalty, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.PenaltyType,
		data.InfringementType,
		data.VehicleIdx,
		data.OtherVehicleIdx,
		data.Time,
		data.LapNum,
		data.PlacesGained,
	}

	return client.insert(ctx, Penalty, args)
}

func (client PostgreClient) insertSpeedTrap(ctx context.Context, data f12020packets.SpeedTrap, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.VehicleIdx,
		data.Speed,
	}

	return client.insert(ctx, SpeedTrap, args)
}

func (client PostgreClient) insertPacketParticipantsData(ctx context.Context, data f12020packets.PacketParticipantsData, headerID uuid.UUID) (uuid.UUID, error) {
	id := uuid.New()

	args := []interface{}{
		id,
		headerID,
		data.NumActiveCars,
	}

	return id, client.insert(ctx, PacketParticipantsData, args)
}

func (client PostgreClient) insertParticipantData(ctx context.Context, data f12020packets.ParticipantData, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.AiControlled,
		data.DriverID,
		data.TeamID,
		data.RaceNumber,
		data.Nationality,
		data.Name,
		data.YourTelemetry,
	}

	return client.insert(ctx, ParticipantData, args)
}

func (client PostgreClient) insertPacketCarSetupData(ctx context.Context, data f12020packets.PacketCarSetupData, headerID uuid.UUID) (uuid.UUID, error) {
	id := uuid.New()

	args := []interface{}{
		id,
		headerID,
	}

	return id, client.insert(ctx, PacketCarSetupData, args)
}

func (client PostgreClient) insertCarSetupData(ctx context.Context, data f12020packets.CarSetupData, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.FrontWing,
		data.RearWing,
		data.OnThrottle,
		data.OffThrottle,
		data.FrontCamber,
		data.RearCamber,
		data.FrontToe,
		data.RearToe,
		data.FrontSuspension,
		data.RearSuspension,
		data.FrontAntiRollBar,
		data.RearAntiRollBar,
		data.FrontSuspensionHeight,
		data.RearSuspensionHeight,
		data.BrakePressure,
		data.BrakeBias,
		data.RearLeftTyrePressure,
		data.RearRightTyrePressure,
		data.FrontLeftTyrePressure,
		data.FrontRightTyrePressure,
		data.Ballast,
		data.FuelLoad,
	}

	return client.insert(ctx, CarSetupData, args)
}

func (client PostgreClient) insertPacketCarTelemetryData(ctx context.Context, data f12020packets.PacketCarTelemetryData, headerID uuid.UUID) (uuid.UUID, error) {
	id := uuid.New()

	args := []interface{}{
		id,
		headerID,
		data.ButtonStatus,
		data.MfdPanelIndex,
		data.MfdPanelIndexSecondaryPlayer,
		data.SuggestedGear,
	}

	return id, client.insert(ctx, PacketCarTelemetryData, args)
}

func (client PostgreClient) insertCarTelemetryData(ctx context.Context, data f12020packets.CarTelemetryData, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.Speed,
		data.Throttle,
		data.Steer,
		data.Brake,
		data.Clutch,
		data.Gear,
		data.EngineRPM,
		data.Drs,
		data.RevLightsPercent,
		data.BrakesTemperatureRL,
		data.BrakesTemperatureRR,
		data.BrakesTemperatureFL,
		data.BrakesTemperatureFR,
		data.TyresSurfaceTemperatureRL,
		data.TyresSurfaceTemperatureRR,
		data.TyresSurfaceTemperatureFL,
		data.TyresSurfaceTemperatureFR,
		data.TyresInnerTemperatureRL,
		data.TyresInnerTemperatureRR,
		data.TyresInnerTemperatureFL,
		data.TyresInnerTemperatureFR,
		data.EngineTemperature,
		data.TyresPressureRL,
		data.TyresPressureRR,
		data.TyresPressureFL,
		data.TyresPressureFR,
		data.SurfaceTypeRL,
		data.SurfaceTypeRR,
		data.SurfaceTypeFL,
		data.SurfaceTypeFR,
	}

	return client.insert(ctx, CarTelemetryData, args)
}

func (client PostgreClient) insertPacketCarStatusData(ctx context.Context, data f12020packets.PacketCarStatusData, headerID uuid.UUID) (uuid.UUID, error) {
	id := uuid.New()

	args := []interface{}{
		id,
		headerID,
	}

	return id, client.insert(ctx, PacketCarStatusData, args)
}

func (client PostgreClient) insertCarStatusData(ctx context.Context, data f12020packets.CarStatusData, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.TractionControl,
		data.AntiLockBrakes,
		data.FuelMix,
		data.FrontBrakeBias,
		data.PitLimiterStatus,
		data.FuelInTank,
		data.FuelCapacity,
		data.FuelRemainingLaps,
		data.MaxRPM,
		data.IdleRPM,
		data.MaxGears,
		data.DrsAllowed,
		data.DrsActivationDistance,
		data.TyresWearRL,
		data.TyresWearRR,
		data.TyresWearFL,
		data.TyresWearFR,
		data.ActualTyreCompound,
		data.VisualTyreCompound,
		data.TyresAgeLaps,
		data.TyresDamageRL,
		data.TyresDamageRR,
		data.TyresDamageFL,
		data.TyresDamageFR,
		data.FrontLeftWingDamage,
		data.FrontRightWingDamage,
		data.RearWingDamage,
		data.DrsFault,
		data.EngineDamage,
		data.GearBoxDamage,
		data.VehicleFiaFlags,
		data.ErsStoreEnergy,
		data.ErsDeployMode,
		data.ErsHarvestedThisLapMGUK,
		data.ErsHarvestedThisLapMGUH,
		data.ErsDeployedThisLap,
	}

	return client.insert(ctx, CarStatusData, args)
}

func (client PostgreClient) insertPacketFinalClassificationData(ctx context.Context, data f12020packets.PacketFinalClassificationData, headerID uuid.UUID) (uuid.UUID, error) {
	id := uuid.New()

	args := []interface{}{
		id,
		headerID,
		data.NumCars,
	}

	return id, client.insert(ctx, PacketFinalClassificationData, args)
}

func (client PostgreClient) insertFinalClassificationData(ctx context.Context, data f12020packets.FinalClassificationData, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.Position,
		data.NumLaps,
		data.GridPosition,
		data.Points,
		data.NumPitStops,
		data.ResultStatus,
		data.BestLapTime,
		data.TotalRaceTime,
		data.PenaltiesTime,
		data.NumPenalties,
		data.NumTyreStints,
		// data.TyreStintsActual, // Problematic
		// data.TyreStintsVisual, // Problematic
	}

	return client.insert(ctx, FinalClassificationData, args)
}

func (client PostgreClient) insertPacketLobbyInfoData(ctx context.Context, data f12020packets.PacketLobbyInfoData, headerID uuid.UUID) (uuid.UUID, error) {
	id := uuid.New()

	args := []interface{}{
		id,
		headerID,
		data.NumPlayers,
	}

	return id, client.insert(ctx, PacketLobbyInfoData, args)
}

func (client PostgreClient) insertLobbyInfoData(ctx context.Context, data f12020packets.LobbyInfoData, packetID uuid.UUID) error {
	id := uuid.New()

	args := []interface{}{
		id,
		packetID,
		data.AiControlled,
		data.TeamID,
		data.Nationality,
		data.Name,
		data.ReadyStatus,
	}

	return client.insert(ctx, LobbyInfoData, args)
}
