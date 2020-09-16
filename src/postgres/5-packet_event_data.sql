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

CREATE OR REPLACE PROCEDURE insert_packet_event_data("ID" uuid, "PacketHeader_ID" uuid)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO PacketEventData 
		VALUES ("ID", "PacketHeader_ID");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_fastest_lap("ID" uuid, "PacketEventData_ID" uuid)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO FastestLap 
		VALUES ("ID", "PacketEventData_ID");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_retirement("ID" uuid, "PacketEventData_ID" uuid)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO Retirement 
		VALUES ("ID", "PacketEventData_ID");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_teammate_in_pits("ID" uuid, "PacketEventData_ID" uuid)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO TeamMateInPits 
		VALUES ("ID", "PacketEventData_ID");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_race_winner("ID" uuid, "PacketEventData_ID" uuid)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO RaceWinner 
		VALUES ("ID", "PacketEventData_ID");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_penalty("ID" uuid, "PacketEventData_ID" uuid)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO Penalty 
		VALUES ("ID", "PacketEventData_ID");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_speed_trap("ID" uuid, "PacketEventData_ID" uuid)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO SpeedTrap 
		VALUES ("ID", "PacketEventData_ID");
	END;
$BODY$;