CREATE TABLE IF NOT EXISTS PacketLobbyInfoData (
   ID                           uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID              uuid NOT NULL,
   NumPlayers                   INT NOT NULL,
   CreatedOn                    TIMESTAMPTZ,

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketHeader_ID)
      REFERENCES PacketHeader (ID)
);

CREATE TABLE IF NOT EXISTS LobbyInfoData (
   ID                           uuid DEFAULT uuid_generate_v4 (),
   PacketLobbyInfoData_ID       uuid NOT NULL,
   AiControlled                 INT NOT NULL,
   TeamID                       INT NOT NULL,
   Nationality                  INT NOT NULL,
   Name                         VARCHAR(48) NOT NULL,
   ReadyStatus                  INT NOT NULL,
   CreatedOn                    TIMESTAMPTZ,

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketLobbyInfoData_ID)
      REFERENCES PacketLobbyInfoData (ID)
);