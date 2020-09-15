CREATE TABLE [IF NOT EXISTS] PacketEventData (
   ID                         SERIAL PRIMARY KEY,
   PacketHeader_ID            SERIAL FOREIGN KEY,
   EventStringCode            VARCHAR(50) NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] FastestLap (
   ID                         SERIAL PRIMARY KEY,
   PacketEventData_ID         SERIAL FOREIGN KEY,
   VehicleIdx                 INT NOT NULL,
   LapTime                    FLOAT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] Retirement (
   ID                         SERIAL PRIMARY KEY,
   PacketEventData_ID         SERIAL FOREIGN KEY,
   VehicleIdx                 INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] TeamMateInPits (
   ID                         SERIAL PRIMARY KEY,
   PacketEventData_ID         SERIAL FOREIGN KEY,
   VehicleIdx                 INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] RaceWinner (
   ID                         SERIAL PRIMARY KEY,
   PacketEventData_ID         SERIAL FOREIGN KEY,
   VehicleIdx                 INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] Penalty (
   ID                         SERIAL PRIMARY KEY,
   PacketEventData_ID         SERIAL FOREIGN KEY,
   PenaltyType                INT NOT NULL,
   InfringementType           INT NOT NULL,
   VehicleIdx                 INT NOT NULL,
   OtherVehicleIdx            INT NOT NULL,
   Time                       INT NOT NULL,
   LapNum                     INT NOT NULL,
   PlacesGained               INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] SpeedTrap (
   ID                         SERIAL PRIMARY KEY,
   PacketEventData_ID         SERIAL FOREIGN KEY,
   VehicleIdx                 INT NOT NULL,
   Speed                      FLOAT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);