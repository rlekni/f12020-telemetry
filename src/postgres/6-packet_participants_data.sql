CREATE TABLE IF NOT EXISTS PacketParticipantsData (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID            uuid NOT NULL,
   NumActiveCars              INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ,
   
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
   CreatedOn                  TIMESTAMPTZ,

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketParticipantsData_ID)
      REFERENCES PacketParticipantsData (ID)
);