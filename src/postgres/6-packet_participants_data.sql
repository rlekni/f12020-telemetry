CREATE TABLE [IF NOT EXISTS] PacketParticipantsData (
   ID                         SERIAL PRIMARY KEY,
   PacketHeader_ID            SERIAL FOREIGN KEY,
   NumActiveCars              INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] ParticipantData (
   ID                         SERIAL PRIMARY KEY,
   PacketParticipantsData_ID  SERIAL FOREIGN KEY,
   AiControlled               INT NOT NULL,
   DriverID                   INT NOT NULL,
   TeamID                     INT NOT NULL,
   RaceNumber                 INT NOT NULL,
   Nationality                INT NOT NULL,
   Name                       VARCHAR(48) NOT NULL,
   YourTelemetry              INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);