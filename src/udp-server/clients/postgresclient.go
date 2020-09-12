package clients

import (
	"context"
	"database/sql"
	"fmt"
	"main/f12020packets"
	"main/helpers"

	"github.com/sirupsen/logrus"
)

const (
	packetHeaderSQL = `
	INSERT INTO packetHeader (type)
	VALUES ($1)
	RETURNING id`
	packetMotionDataSQL = `
	INSERT INTO packetMotionData (type)
	VALUES ($1)
	RETURNING id`
	packetSessionDataSQL = `
	INSERT INTO packetSessionData (type)
	VALUES ($1)
	RETURNING id`
	packetLapDataSQL = `
	INSERT INTO packetLapData (type)
	VALUES ($1)
	RETURNING id`
	packetEventDataSQL = `
	INSERT INTO packetEventData (type)
	VALUES ($1)
	RETURNING id`
	packetParticipantsDataSQL = `
	INSERT INTO packetParticipantsData (type)
	VALUES ($1)
	RETURNING id`
	packetCarSetupDataSQL = `
	INSERT INTO packetCarSetupData (type)
	VALUES ($1)
	RETURNING id`
	packetCarTelemetryDataSQL = `
	INSERT INTO packetCarTelemetryData (type)
	VALUES ($1)
	RETURNING id`
	packetCarStatusDataSQL = `
	INSERT INTO packetCarStatusData (type)
	VALUES ($1)
	RETURNING id`
	packetFinalClassificationDataSQL = `
	INSERT INTO packetFinalClassificationData (type)
	VALUES ($1)
	RETURNING id`
	packetLobbyInfoDataSQL = `
	INSERT INTO packetLobbyInfoData (type)
	VALUES ($1)
	RETURNING id`
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
	err := client.insert(ctx, "PacketHeader", packetObject.Header.SessionUID)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketMotionData, "")
}

func (client PostgreClient) InsertPacketSessionData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketSessionData)
	err := client.insert(ctx, "PacketHeader", packetObject.Header.SessionUID)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketSessionData, "")
}

func (client PostgreClient) InsertPacketLapData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketLapData)
	err := client.insert(ctx, "PacketHeader", packetObject.Header.SessionUID)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketLapData, "")
}

func (client PostgreClient) InsertPacketEventData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketEventData)
	err := client.insert(ctx, "PacketHeader", packetObject.Header.SessionUID)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketEventData, "")
}

func (client PostgreClient) InsertPacketParticipantsData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketParticipantsData)
	err := client.insert(ctx, "PacketHeader", packetObject.Header.SessionUID)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketParticipantsData, "")
}

func (client PostgreClient) InsertPacketCarSetupData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketCarSetupData)
	err := client.insert(ctx, "PacketHeader", packetObject.Header.SessionUID)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketCarSetupData, "")
}

func (client PostgreClient) InsertPacketCarTelemetryData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketCarTelemetryData)
	err := client.insert(ctx, "PacketHeader", packetObject.Header.SessionUID)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketCarTelemetryData, "")
}

func (client PostgreClient) InsertPacketCarStatusData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketCarStatusData)
	err := client.insert(ctx, "PacketHeader", packetObject.Header.SessionUID)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketCarStatusData, "")
}

func (client PostgreClient) InsertPacketFinalClassificationData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketFinalClassificationData)
	err := client.insert(ctx, "PacketHeader", packetObject.Header.SessionUID)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketFinalClassificationData, "")
}

func (client PostgreClient) InsertPacketLobbyInfoData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketLobbyInfoData)
	err := client.insert(ctx, "PacketHeader", packetObject.Header.SessionUID)
	if err != nil {
		return err
	}
	return client.insert(ctx, PacketLobbyInfoData, "")
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
