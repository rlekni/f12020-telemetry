CREATE TABLE IF NOT EXISTS PacketHeader (
   ID                      uuid DEFAULT uuid_generate_v4 (),
   PacketFormat            INT NOT NULL,
   GameMajorVersion        INT NOT NULL,
   GameMinorVersion        INT NOT NULL,
   PacketVersion           INT NOT NULL,
   PacketID                INT NOT NULL,
   SessionUID              VARCHAR(250) NOT NULL,
   SessionTime             FLOAT NOT NULL,
   FrameIdentifier         INT NOT NULL,
   PlayerCarIndex          INT NOT NULL,
   SecondaryPlayerCarIndex INT NOT NULL,
   CreatedOn               TIMESTAMPTZ,

   PRIMARY KEY (ID)
);

CREATE OR REPLACE PROCEDURE insert_packet_header("PacketFormat" integer, "GameMajorVersion" integer, "GameMinorVersion" integer, "PacketVersion" integer, "PacketID" integer, "SessionUID" text, "SessionTime" double precision, "FrameIdentifier" integer, "PlayerCarIndex" integer, "SecondaryPlayerCarIndex" integer)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO PacketHeader 
		VALUES ("PacketFormat", "GameMajorVersion", "GameMinorVersion", "PacketVersion", "PacketID", "SessionUID", "SessionTime", "FrameIdentifier", "PlayerCarIndex", "SecondaryPlayerCarIndex");
	END;
$BODY$;