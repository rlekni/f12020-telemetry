# go-f1-telemetry

[![Build Status](https://rlekni.visualstudio.com/hbi/_apis/build/status/rlekni.go-f1-telemetry?branchName=serialisation)](https://rlekni.visualstudio.com/hbi/_build/latest?definitionId=20&branchName=serialisation)

## F1 2020

Telemetry specification [found here](https://forums.codemasters.com/topic/54423-f1%C2%AE-2020-udp-specification/)

## UDP Server and Client

Server:

* `go run main.go 1234` where 1234 is a PORT

NOTE: If port number is not provided, defaults to use 20777

## Packet IDs

The packets IDs are as follows:

| Packet Name          | Value | Description                                                                      |
| -------------------- | ----- | -------------------------------------------------------------------------------- |
| Motion               | 0     | Contains all motion data for player’s car – only sent while player is in control |
| Session              | 1     | Data about the session – track, time left                                        |
| Lap Data             | 2     | Data about all the lap times of cars in the session                              |
| Event                | 3     | Various notable events that happen during a session                              |
| Participants         | 4     | List of participants in the session, mostly relevant for multiplayer             |
| Car Setups           | 5     | Packet detailing car setups for cars in the race                                 |
| Car Telemetry        | 6     | Telemetry data for all cars                                                      |
| Car Status           | 7     | Status data for all cars such as damage                                          |
| Final Classification | 8     | Final classification confirmation at the end of a race                           |
| Lobby Info           | 9     | Information about players in a multiplayer lobby                                 |


## Event String Codes

| Event                | Code   | Description                                    |
| -------------------- | ------ | ---------------------------------------------- |
| Session Started      | “SSTA” | Sent when the session starts                   |
| Session Ended        | “SEND” | Sent when the session ends                     |
| Fastest Lap          | “FTLP” | When a driver achieves the fastest lap         |
| Retirement           | “RTMT” | When a driver retires                          |
| DRS enabled          | “DRSE” | Race control have enabled DRS                  |
| DRS disabled         | “DRSD” | Race control have disabled DRS                 |
| Team mate in pits    | “TMPT” | Your team mate has entered the pits            |
| Chequered flag       | “CHQF” | The chequered flag has been waved              |
| Race Winner          | “RCWN” | The race winner is announced                   |
| Penalty Issued       | “PENA” | A penalty has been issued – details in event   |
| Speed Trap Triggered | “SPTP” | Speed trap has been triggered by fastest speed |

### Main packets

| Packet Name                   | Size in bytes | Frequency                       |
| ----------------------------- | ------------- | ------------------------------- |
| PacketMotionData              | 1464          | Rate in menus (20 Hz)           |
| PacketSessionData             | 251           | 2 per second                    |
| PacketLapData                 | 1190          | Rate in menus (20 Hz)           |
| PacketEventData               | 35            | When even occurs                |
| PacketParticipantsData        | 1213          | Every 5 seconds                 |
| PacketCarSetupData            | 1102          | 2 per second                    |
| PacketCarTelemetryData        | 1307          | Rate in menus (20 Hz)           |
| PacketCarStatusData           | 1344          | Rate in menus (20 Hz)           |
| PacketFinalClassificationData | 839           | Once at the end of a race       |
| PacketLobbyInfoData           | 1169          | 2 per second, when in the lobby |

## Mongo

Setup:

* `sudo mkdir -p /mongodata`
* `sudo docker run -it -v /data/db:/mongodata -p 27017:27017 --name mongodb -d mongo`
* `sudo docker start mongodb`

To access the databases, download mongo compass

## Docker

To list built images:

* `docker images`

Web:

* `docker build -t f1-telemetry-web .`
* `docker run -it -p 8080:80 --rm --name f1-telemetry-web f1-telemetry-web`