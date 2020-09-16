CREATE TABLE IF NOT EXISTS PacketEventData (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID            uuid NOT NULL,
   EventStringCode            VARCHAR(50) NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketHeader_ID)
      REFERENCES PacketHeader (ID)
);

CREATE TABLE IF NOT EXISTS FastestLap (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketEventData_ID         uuid NOT NULL,
   VehicleIdx                 INT NOT NULL,
   LapTime                    FLOAT NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketEventData_ID)
      REFERENCES PacketEventData (ID)
);

CREATE TABLE IF NOT EXISTS Retirement (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketEventData_ID         uuid NOT NULL,
   VehicleIdx                 INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),
   
   PRIMARY KEY (ID),
   FOREIGN KEY (PacketEventData_ID)
      REFERENCES PacketEventData (ID)
);

CREATE TABLE IF NOT EXISTS TeamMateInPits (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketEventData_ID         uuid NOT NULL,
   VehicleIdx                 INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketEventData_ID)
      REFERENCES PacketEventData (ID)
);

CREATE TABLE IF NOT EXISTS RaceWinner (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketEventData_ID         uuid NOT NULL,
   VehicleIdx                 INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketEventData_ID)
      REFERENCES PacketEventData (ID)
);

CREATE TABLE IF NOT EXISTS Penalty (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketEventData_ID         uuid NOT NULL,
   PenaltyType                INT NOT NULL,
   InfringementType           INT NOT NULL,
   VehicleIdx                 INT NOT NULL,
   OtherVehicleIdx            INT NOT NULL,
   Time                       INT NOT NULL,
   LapNum                     INT NOT NULL,
   PlacesGained               INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketEventData_ID)
      REFERENCES PacketEventData (ID)
);

CREATE TABLE IF NOT EXISTS SpeedTrap (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketEventData_ID         uuid NOT NULL,
   VehicleIdx                 INT NOT NULL,
   Speed                      FLOAT NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketEventData_ID)
      REFERENCES PacketEventData (ID)
);