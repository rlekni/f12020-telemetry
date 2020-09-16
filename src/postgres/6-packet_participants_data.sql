CREATE TABLE IF NOT EXISTS PacketParticipantsData (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID            uuid NOT NULL,
   NumActiveCars              INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),
   
   PRIMARY KEY (ID),
   FOREIGN KEY (PacketHeader_ID)
      REFERENCES PacketHeader (ID)
);

CREATE TABLE IF NOT EXISTS ParticipantData (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketParticipantsData_ID  uuid NOT NULL,
   AiControlled               INT NOT NULL,
   DriverID                   INT NOT NULL,
   TeamID                     INT NOT NULL,
   RaceNumber                 INT NOT NULL,
   Nationality                INT NOT NULL,
   Name                       VARCHAR(48) NOT NULL,
   YourTelemetry              INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketParticipantsData_ID)
      REFERENCES PacketParticipantsData (ID)
);

CREATE OR REPLACE PROCEDURE insert_packet_participants_data("ID" uuid, "PacketHeader_ID" uuid)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO PacketParticipantsData 
		VALUES ("ID", "PacketHeader_ID");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_participant_data("ID" uuid, "PacketParticipantsData_ID" uuid)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO ParticipantData 
		VALUES ("ID", "PacketParticipantsData_ID");
	END;
$BODY$;