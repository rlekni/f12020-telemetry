CREATE TABLE [IF NOT EXISTS] PacketHeader (
   ID                      SERIAL PRIMARY KEY,
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
   CreatedOn               TIMESTAMPTZ
);

CREATE OR REPLACE PROCEDURE insert_packet_header("PacketFormat" INTEGER, "GameMajorVersion" INTEGER, "GameMinorVersion" INTEGER, "PacketVersion" INTEGER, "PacketID" INTEGER, "SessionUID" VARCHAR(250), "SessionTime" FLOAT, "FrameIdentifier" INTEGER, "PlayerCarIndex" INTEGER, "SecondaryPlayerCarIndex" INTEGER)
LANGUAGE PLPGSQL
AS $$
  INSERT INTO PacketHeader 
	VALUES ("PacketFormat", "GameMajorVersion", "GameMinorVersion", "PacketVersion", "PacketID", "SessionUID", "SessionTime", "FrameIdentifier", "PlayerCarIndex", "SecondaryPlayerCarIndex");
$$;