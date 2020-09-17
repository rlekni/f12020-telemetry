CREATE TABLE IF NOT EXISTS PacketFinalClassificationData (
   ID                           uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID              uuid NOT NULL,
   NumCars                      INT NOT NULL,
   CreatedOn                    TIMESTAMPTZ DEFAULT NOW(),

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
   TyreStintsActual                   INT[] DEFAULT NULL,
   TyreStintsVisual                   INT[] DEFAULT NULL,
   CreatedOn                          TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketFinalClassificationData_ID)
      REFERENCES PacketFinalClassificationData (ID)
);

CREATE OR REPLACE PROCEDURE insert_packet_final_classification_data("ID" uuid, "PacketHeader_ID" uuid, "NumCars" integer)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO PacketFinalClassificationData 
		VALUES ("ID", "PacketHeader_ID", "NumCars");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_final_classification_data("ID" uuid, "PacketFinalClassificationData_ID" uuid, "Position" integer, "NumLaps" integer, "GridPosition" integer, "Points" integer, "NumPitStops" integer, "ResultStatus" integer, "BestLapTime" double precision, "TotalRaceTime" double precision, "PenaltiesTime" integer, "NumPenalties" integer, "NumTyreStints" integer)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO FinalClassificationData 
		VALUES ("ID", "PacketFinalClassificationData_ID", "Position", "NumLaps", "GridPosition", "Points", "NumPitStops", "ResultStatus", "BestLapTime", "TotalRaceTime", "PenaltiesTime", "NumPenalties", "NumTyreStints");
	END;
$BODY$;
