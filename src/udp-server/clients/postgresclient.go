package clients

import (
	"context"
	"database/sql"
	"fmt"
	"main/f12020packets"
	"main/helpers"

	"github.com/sirupsen/logrus"
)

type PostgreClient struct {
	Database *sql.DB
}

func (client PostgreClient) Disconnect(ctx context.Context) error {
	logrus.Warningln("Closing PostgreSQL connection!")
	return client.Database.Close()
}

func (client PostgreClient) Insert(ctx context.Context, packetType string, packet interface{}) error {
	return fmt.Errorf("DEPRECATED")
}

func (client PostgreClient) InsertPacketMotionData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketMotionData)
	err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}

	for _, carMotionData := range packetObject.CarMotionData {
		err = client.insertCarMotionData(ctx, carMotionData)
		helpers.LogIfError(err)
	}

	args := []interface{}{
		packetObject.SuspensionPositionRL,
		packetObject.SuspensionPositionRR,
		packetObject.SuspensionPositionFL,
		packetObject.SuspensionPositionFR,
		packetObject.SuspensionVelocityRL,
		packetObject.SuspensionVelocityRR,
		packetObject.SuspensionVelocityFL,
		packetObject.SuspensionVelocityFR,
		packetObject.SuspensionAccelerationRL,
		packetObject.SuspensionAccelerationRR,
		packetObject.SuspensionAccelerationFL,
		packetObject.SuspensionAccelerationFR,
		packetObject.WheelSpeedRL,
		packetObject.WheelSpeedRR,
		packetObject.WheelSpeedFL,
		packetObject.WheelSpeedFR,
		packetObject.WheelSlipRL,
		packetObject.WheelSlipRR,
		packetObject.WheelSlipFL,
		packetObject.WheelSlipFR,
		packetObject.LocalVelocityX,
		packetObject.LocalVelocityY,
		packetObject.LocalVelocityZ,
		packetObject.AngularVelocityX,
		packetObject.AngularVelocityY,
		packetObject.AngularVelocityZ,
		packetObject.AngularAccelerationX,
		packetObject.AngularAccelerationY,
		packetObject.AngularAccelerationZ,
		packetObject.FrontWheelsAngle,
	}
	return client.insert(ctx, PacketMotionData, args...)
}

func (client PostgreClient) InsertPacketSessionData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketSessionData)
	err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketSessionData, "")
}

func (client PostgreClient) InsertPacketLapData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketLapData)
	err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketLapData, "")
}

func (client PostgreClient) InsertPacketEventData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketEventData)
	err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketEventData, "")
}

func (client PostgreClient) InsertPacketParticipantsData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketParticipantsData)
	err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketParticipantsData, "")
}

func (client PostgreClient) InsertPacketCarSetupData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketCarSetupData)
	err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketCarSetupData, "")
}

func (client PostgreClient) InsertPacketCarTelemetryData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketCarTelemetryData)
	err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketCarTelemetryData, "")
}

func (client PostgreClient) InsertPacketCarStatusData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketCarStatusData)
	err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketCarStatusData, "")
}

func (client PostgreClient) InsertPacketFinalClassificationData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketFinalClassificationData)
	err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketFinalClassificationData, "")
}

func (client PostgreClient) InsertPacketLobbyInfoData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketLobbyInfoData)
	err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketLobbyInfoData, "")
}

func (client PostgreClient) insertPacketHeader(ctx context.Context, header *f12020packets.PacketHeader) error {
	args := []interface{}{
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
	return client.insert(ctx, "PacketHeader", args...)
}

func (client PostgreClient) insertCarMotionData(ctx context.Context, data f12020packets.CarMotionData) error {
	args := []interface{}{
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
	return client.insert(ctx, "CarMotionData", args...)
}

func (client PostgreClient) insert(ctx context.Context, packetType string, args ...interface{}) error {
	var id int
	sqlStatement, err := getSQLStatement(packetType)
	if err != nil {
		return err
	}
	err = client.Database.QueryRow(*sqlStatement, args).Scan(&id)
	helpers.LogIfError(err)
	return err
}

func getSQLStatement(packetType string) (*string, error) {
	var sql string
	switch packetType {
	case "PacketHeader":
		sql = packetHeaderSQL
	case "CarMotionData":
		sql = carMotionDataSQL
	case PacketMotionData:
		sql = packetMotionDataSQL
	case PacketSessionData:
		sql = packetSessionDataSQL
	case PacketLapData:
		sql = packetLapDataSQL
	case PacketEventData:
		sql = packetEventDataSQL
	case PacketParticipantsData:
		sql = packetParticipantsDataSQL
	case PacketCarSetupData:
		sql = packetCarSetupDataSQL
	case PacketCarTelemetryData:
		sql = packetCarTelemetryDataSQL
	case PacketCarStatusData:
		sql = packetCarStatusDataSQL
	case PacketFinalClassificationData:
		sql = packetFinalClassificationDataSQL
	case PacketLobbyInfoData:
		sql = packetLobbyInfoDataSQL
	default:
		return nil, fmt.Errorf("PacketType: %s is not valid", packetType)
	}
	return &sql, nil
}
