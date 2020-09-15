CREATE TABLE IF NOT EXISTS PacketFinalClassificationData (
   ID                           uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID              uuid NOT NULL,
   NumCars                      INT NOT NULL,
   CreatedOn                    TIMESTAMPTZ,

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketHeader_ID)
      REFERENCES PacketHeader (ID)
);

CREATE TABLE IF NOT EXISTS FinalClassificationData (
   ID                                 uuid DEFAULT uuid_generate_v4 (),
   PacketFinalClassificationData_ID   uuid NOT NULL,
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
   TyreStintsActual                   INT[] NOT NULL,
   TyreStintsVisual                   INT[] NOT NULL,
   CreatedOn                          TIMESTAMPTZ,

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketFinalClassificationData_ID)
      REFERENCES PacketFinalClassificationData (ID)
);