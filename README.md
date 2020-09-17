# go-f1-telemetry

This repository is for easy telemetry capture from F1 2020 game. The whole stack can easily be run on raspberry pi 4.

## F1 2020

Telemetry specification [found here](https://forums.codemasters.com/topic/54423-f1%C2%AE-2020-udp-specification/)

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

## Main packets

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

## Docker Setup

Docker commands:

* `sudo docker-compose up -d --remove-orphans --build` build and deploy
* `sudo docker-compose down` to stop and remove all containers
* `sudo docker-compose stop` will stop containers, but won't remove them
* `sudo docker-compose start` will start containers again