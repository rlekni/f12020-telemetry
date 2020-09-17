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
	headerID, err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}

	packetID, err := client.insertPacketMotionData(ctx, packetObject, headerID)
	if err != nil {
		return err
	}

	for _, carMotionData := range packetObject.CarMotionData {
		err = client.insertCarMotionData(ctx, carMotionData, packetID)
		helpers.LogIfError(err)
	}

	return nil
}

func (client PostgreClient) InsertPacketSessionData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketSessionData)
	headerID, err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}

	packetID, err := client.insertPacketSessionData(ctx, packetObject, headerID)
	if err != nil {
		return err
	}

	for _, data := range packetObject.MarshalZones {
		err = client.insertMarshalZone(ctx, data, packetID)
		helpers.LogIfError(err)
	}

	for _, data := range packetObject.WeatherForecastSamples {
		err = client.insertWeatherForecastSample(ctx, data, packetID)
		helpers.LogIfError(err)
	}

	return nil
}

func (client PostgreClient) InsertPacketLapData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketLapData)
	headerID, err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}

	packetID, err := client.insertPacketLapData(ctx, packetObject, headerID)
	if err != nil {
		return err
	}

	for _, data := range packetObject.LapData {
		err = client.insertLapData(ctx, data, packetID)
		helpers.LogIfError(err)
	}
	return nil
}

func (client PostgreClient) InsertPacketEventData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketEventData)
	headerID, err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}

	packetID, err := client.insertPacketEventData(ctx, packetObject, headerID)
	if err != nil {
		return err
	}
	switch packetObject.EventStringCode {
	case "FTLP":
		data := packetObject.EventDetails.(f12020packets.FastestLap)
		err = client.insertFastestLap(ctx, data, packetID)
		if err != nil {
			return err
		}
	case "RTMT":
		data := packetObject.EventDetails.(f12020packets.Retirement)
		err = client.insertRetirement(ctx, data, packetID)
		if err != nil {
			return err
		}
	case "TMPT":
		data := packetObject.EventDetails.(f12020packets.TeamMateInPits)
		err = client.insertTeamMateInPits(ctx, data, packetID)
		if err != nil {
			return err
		}
	case "RCWN":
		data := packetObject.EventDetails.(f12020packets.RaceWinner)
		err = client.insertRaceWinner(ctx, data, packetID)
		if err != nil {
			return err
		}
	case "PENA":
		data := packetObject.EventDetails.(f12020packets.Penalty)
		err = client.insertPenalty(ctx, data, packetID)
		if err != nil {
			return err
		}
	case "SPTP":
		data := packetObject.EventDetails.(f12020packets.SpeedTrap)
		err = client.insertSpeedTrap(ctx, data, packetID)
		if err != nil {
			return err
		}
	default:
		logrus.Warningf("Skipping insert: None of the event Codes matched event code supplied: %q.", packetObject.EventStringCode)
	}

	return nil
}

func (client PostgreClient) InsertPacketParticipantsData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketParticipantsData)
	headerID, err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}
	packetID, err := client.insertPacketParticipantsData(ctx, packetObject, headerID)
	if err != nil {
		return err
	}

	for _, data := range packetObject.Participants {
		err = client.insertParticipantData(ctx, data, packetID)
		helpers.LogIfError(err)
	}

	return nil
}

func (client PostgreClient) InsertPacketCarSetupData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketCarSetupData)
	headerID, err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}
	packetID, err := client.insertPacketCarSetupData(ctx, packetObject, headerID)
	if err != nil {
		return err
	}

	for _, data := range packetObject.CarSetups {
		err = client.insertCarSetupData(ctx, data, packetID)
		helpers.LogIfError(err)
	}

	return nil
}

func (client PostgreClient) InsertPacketCarTelemetryData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketCarTelemetryData)
	headerID, err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}

	packetID, err := client.insertPacketCarTelemetryData(ctx, packetObject, headerID)
	if err != nil {
		return err
	}

	for _, data := range packetObject.CarTelemetryData {
		err = client.insertCarTelemetryData(ctx, data, packetID)
		helpers.LogIfError(err)
	}

	return nil
}

func (client PostgreClient) InsertPacketCarStatusData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketCarStatusData)
	headerID, err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}

	packetID, err := client.insertPacketCarStatusData(ctx, packetObject, headerID)
	if err != nil {
		return err
	}

	for _, data := range packetObject.CarStatusData {
		err = client.insertCarStatusData(ctx, data, packetID)
		helpers.LogIfError(err)
	}

	return nil
}

func (client PostgreClient) InsertPacketFinalClassificationData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketFinalClassificationData)
	headerID, err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}

	packetID, err := client.insertPacketFinalClassificationData(ctx, packetObject, headerID)
	if err != nil {
		return err
	}

	for _, data := range packetObject.ClassificationData {
		err = client.insertFinalClassificationData(ctx, data, packetID)
		helpers.LogIfError(err)
	}

	return nil
}

func (client PostgreClient) InsertPacketLobbyInfoData(ctx context.Context, packet interface{}) error {
	packetObject := packet.(f12020packets.PacketLobbyInfoData)
	headerID, err := client.insertPacketHeader(ctx, packetObject.Header)
	if err != nil {
		return err
	}

	packetID, err := client.insertPacketLobbyInfoData(ctx, packetObject, headerID)
	if err != nil {
		return err
	}

	for _, data := range packetObject.LobbyPlayers {
		err = client.insertLobbyInfoData(ctx, data, packetID)
		helpers.LogIfError(err)
	}

	return nil
}
