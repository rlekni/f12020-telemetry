CREATE TABLE [IF NOT EXISTS] PacketFinalClassificationData (
   ID                           SERIAL PRIMARY KEY,
   PacketHeader_ID              SERIAL FOREIGN KEY,
   NumCars                      INT NOT NULL,
   CreatedOn                    TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] FinalClassificationData (
   ID                                 SERIAL PRIMARY KEY,
   PacketFinalClassificationData_ID   SERIAL FOREIGN KEY,
   Position                           INT NOT NULL,
   NumLaps                            INT NOT NULL,
   GridPosition                       INT NOT NULL,
   Points                             INT NOT NULL,
   NumPitStops                        INT NOT NULL,
   ResultStatus                       INT NOT NULL,
   BestLapTime                        FLOAT NOT NULL,
   TotalRaceTime                      FLOAT NOT NULL,
   PenaltiesTime                      INT NOT NULL,
   NumPenalties                       INT NOT NULL,
   NumTyreStints                      INT NOT NULL,
   NumCars                            INT[] NOT NULL,
   NumCars                            INT[] NOT NULL,
   CreatedOn                          TIMESTAMPTZ
);