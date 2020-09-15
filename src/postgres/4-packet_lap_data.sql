CREATE TABLE [IF NOT EXISTS] PacketLapData (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID            uuid NOT NULL,
   CreatedOn                  TIMESTAMPTZ,

   PRIMARY KEY (ID, PacketHeader_ID),
   FOREIGN KEY (PacketHeader_ID),
      REFERENCES PacketHeader (ID)
);

CREATE TABLE [IF NOT EXISTS] LapData (
  ID                          uuid DEFAULT uuid_generate_v4 (),
  PacketLapData_ID            uuid NOT NULL,
  LastLapTime                 FLOAT NOT NULL,
  CurrentLapTime              FLOAT NOT NULL,
  Sector1TimeInMS             INT NOT NULL,
  Sector2TimeInMS             INT NOT NULL,
  BestLapTime                 FLOAT NOT NULL,
  BestLapNum                  INT NOT NULL,
  BestLapSector1TimeInMS      INT NOT NULL,
  BestLapSector2TimeInMS      INT NOT NULL,
  BestLapSector3TimeInMS      INT NOT NULL,
  BestOverallSector1TimeInMS  INT NOT NULL,
  BestOverallSector1LapNum    INT NOT NULL,
  BestOverallSector2TimeInMS  INT NOT NULL,
  BestOverallSector2LapNum    INT NOT NULL,
  BestOverallSector3TimeInMS  INT NOT NULL,
  BestOverallSector3LapNum    INT NOT NULL,
  LapDistance                 FLOAT NOT NULL,
  TotalDistance               FLOAT NOT NULL,
  SafetyCarDelta              FLOAT NOT NULL,
  CarPosition                 INT NOT NULL,
  CurrentLapNum               FLOAT NOT NULL,
  PitStatus                   FLOAT NOT NULL,
  Sector                      FLOAT NOT NULL,
  CurrentLapInvalid           FLOAT NOT NULL,
  Penalties                   FLOAT NOT NULL,
  GridPosition                FLOAT NOT NULL,
  DriverStatus                FLOAT NOT NULL,
  ResultStatus                FLOAT NOT NULL,
  CreatedOn                   TIMESTAMPTZ,

  PRIMARY KEY (ID, PacketLapData_ID),
   FOREIGN KEY (PacketLapData_ID),
      REFERENCES PacketLapData (ID)      
)