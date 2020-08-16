# go-f1-telemetry

# F1 2020

Telemetry specification [found here](https://forums.codemasters.com/topic/54423-f1%C2%AE-2020-udp-specification/)

# UDP Server and Client

Server:

* `go run main.go 1234` where 1234 is a PORT

# Packets
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