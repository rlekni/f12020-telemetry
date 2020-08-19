package f12020packets

import (
	"encoding/binary"
	"fmt"
	"math"
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

func convertTo4ByteFloat32Array(data []byte) [4]float32 {
	var result [4]float32
	for i := 0; i < 4; i++ {
		startIndex := 0 + (i * 4)
		endIndex := startIndex + 4
		value := convertToFloat32(data[startIndex:endIndex])
		result[i] = value
	}
	return result
}

// 23 bytes
func ToPacketHeader(data []byte) (*PacketHeader, error) {
	if len(data) != 24 {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", 24, len(data))
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
	if len(data) != 61 {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", 61, len(data))
	}
	motionData := &CarMotionData{}

	return motionData, nil
}

// 1464 bytes
// 23 bytes header
// 1320 bytes car motion data
// 121 bytes the rest
func ToPacketMotionData(data []byte) (*PacketMotionData, error) {
	if len(data) != 1464 {
		return nil, fmt.Errorf("Expected provided data to be %d length, but was %d", 1464, len(data))
	}
	header, err := ToPacketHeader(data[0:24])
	if err != nil {
		return nil, fmt.Errorf("Failed to To Packet Header")
	}

	// 1320 bytes in total
	var motionData [22]CarMotionData
	for i := 0; i < 22; i++ {
		startIndex := 24 + (i * 60)
		// Swallow any exceptions for now
		payload, _ := ToCarMotionData(data[startIndex : startIndex+60])
		motionData[i] = *payload
	}

	// Construct packet and decode the rest of the data
	packet := &PacketMotionData{
		Header:                 header,
		CarMotionData:          motionData,
		SuspensionPosition:     convertTo4ByteFloat32Array(data[1343:1359]),
		SuspensionVelocity:     convertTo4ByteFloat32Array(data[1359:1375]),
		SuspensionAcceleration: convertTo4ByteFloat32Array(data[1375:1391]),
		WheelSpeed:             convertTo4ByteFloat32Array(data[1391:1407]),
		WheelSlip:              convertTo4ByteFloat32Array(data[1407:1423]),
		LocalVelocityX:         convertToFloat32(data[1423:1427]),
		LocalVelocityY:         convertToFloat32(data[1427:1431]),
		LocalVelocityZ:         convertToFloat32(data[1431:1435]),
		AngularVelocityX:       convertToFloat32(data[1435:1439]),
		AngularVelocityY:       convertToFloat32(data[1439:1443]),
		AngularVelocityZ:       convertToFloat32(data[1443:1447]),
		AngularAccelerationX:   convertToFloat32(data[1447:1451]),
		AngularAccelerationY:   convertToFloat32(data[1451:1455]),
		AngularAccelerationZ:   convertToFloat32(data[1455:1459]),
		FrontWheelsAngle:       convertToFloat32(data[1459:1463]),
	}

	return packet, nil
}
