CREATE TABLE IF NOT EXISTS PacketLobbyInfoData (
   ID                           uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID              uuid NOT NULL,
   NumPlayers                   INT NOT NULL,
   CreatedOn                    TIMESTAMPTZ DEFAULT NOW(),

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
   CreatedOn                    TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketLobbyInfoData_ID)
      REFERENCES PacketLobbyInfoData (ID)
);

CREATE OR REPLACE PROCEDURE insert_packet_lobby_info_data("ID" uuid, "PacketHeader_ID" uuid, "NumPlayers" integer)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO PacketLobbyInfoData 
		VALUES ("ID", "PacketHeader_ID", "NumPlayers");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_lobby_info_data("ID" uuid, "PacketLobbyInfoData_ID" uuid, "AiControlled" integer, "TeamID" integer, "Nationality" integer, "Name" text, "ReadyStatus" integer)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO PacketLobbyInfoData 
		VALUES ("ID", "PacketLobbyInfoData_ID", "AiControlled", "TeamID", "Nationality", "Name", "ReadyStatus");
	END;
$BODY$;
