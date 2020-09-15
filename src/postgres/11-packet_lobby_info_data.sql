CREATE TABLE [IF NOT EXISTS] PacketLobbyInfoData (
   ID                           SERIAL PRIMARY KEY,
   PacketHeader_ID              SERIAL FOREIGN KEY,
   NumPlayers                   INT NOT NULL,
   CreatedOn                    TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] LobbyInfoData (
   ID                           SERIAL PRIMARY KEY,
   PacketLobbyInfoData_ID       SERIAL FOREIGN KEY,
   AiControlled                 INT NOT NULL,
   TeamID                       INT NOT NULL,
   Nationality                  INT NOT NULL,
   Name                         VARCHAR(48) NOT NULL,
   ReadyStatus                  INT NOT NULL,
   CreatedOn                    TIMESTAMPTZ
);