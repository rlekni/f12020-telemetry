package internal

import (
	"context"
	"main/clients"
	"main/f12020packets"
	"main/helpers"

	"github.com/sirupsen/logrus"
)

func DeserialisePacket(ctx context.Context, mongoClient clients.RepositoryClient, data []byte) {
	header, err := f12020packets.ToPacketHeader(data[0:24])
	if err != nil {
		logrus.Errorf("Failed to decode Packet Header. Error: %q", err)
	}

	logrus.Debugln("SessionID: ", header.SessionUID)
	logrus.Debugf("Data length: %d, PacketID: %d\n", len(data), header.PacketID)

	switch header.PacketID {
	case 0:
		packet, err := f12020packets.ToPacketMotionData(data[24:1464], header)
		helpers.LogIfError(err)

		mongoClient.Insert(ctx, "packetMotionData", packet)
	case 1:
		packet, err := f12020packets.ToPacketSessionData(data[24:251], header)
		helpers.LogIfError(err)

		mongoClient.Insert(ctx, "packetSessionData", packet)
	case 2:
		packet, err := f12020packets.ToPacketLapData(data[24:1190], header)
		helpers.LogIfError(err)

		mongoClient.Insert(ctx, "packetLapData", packet)
	case 3:
		packet, err := f12020packets.ToPacketEventData(data[24:35], header)
		helpers.LogIfError(err)

		mongoClient.Insert(ctx, "packetEventData", packet)
	case 4:
		packet, err := f12020packets.ToPacketParticipantsData(data[24:1213], header)
		helpers.LogIfError(err)

		mongoClient.Insert(ctx, "packetParticipantsData", packet)
	case 5:
		packet, err := f12020packets.ToPacketCarSetupData(data[24:1102], header)
		helpers.LogIfError(err)

		mongoClient.Insert(ctx, "packetCarSetupData", packet)
	case 6:
		packet, err := f12020packets.ToPacketCarTelemetryData(data[24:1307], header)
		helpers.LogIfError(err)

		mongoClient.Insert(ctx, "packetCarTelemetryData", packet)
	case 7:
		packet, err := f12020packets.ToPacketCarStatusData(data[24:1344], header)
		helpers.LogIfError(err)

		mongoClient.Insert(ctx, "packetCarStatusData", packet)
	case 8:
		packet, err := f12020packets.ToPacketFinalClassificationData(data[24:839], header)
		helpers.LogIfError(err)

		mongoClient.Insert(ctx, "packetFinalClassificationData", packet)
	case 9:
		packet, err := f12020packets.ToPacketLobbyInfoData(data[24:1169], header)
		helpers.LogIfError(err)

		mongoClient.Insert(ctx, "packetLobbyInfoData", packet)
	default:
		logrus.Warningf("None of the defined PacketIDs matched. Data length: %d, PacketID: %d\n", len(data), header.PacketID)
	}
}
