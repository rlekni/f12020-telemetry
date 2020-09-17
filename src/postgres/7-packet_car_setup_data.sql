CREATE TABLE IF NOT EXISTS PacketCarSetupData (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID            uuid NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketHeader_ID)
      REFERENCES PacketHeader (ID)
);

CREATE TABLE IF NOT EXISTS CarSetupData (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketCarSetupData_ID      uuid NOT NULL,
   FrontWing                  INT NOT NULL,
   RearWing                   INT NOT NULL,
   OnThrottle                 INT NOT NULL,
   OffThrottle                INT NOT NULL,
   FrontCamber                FLOAT NOT NULL,
   RearCamber                 FLOAT NOT NULL,
   FrontToe                   FLOAT NOT NULL,
   RearToe                    FLOAT NOT NULL,
   FrontSuspension            INT NOT NULL,
   RearSuspension             INT NOT NULL,
   FrontAntiRollBar           INT NOT NULL,
   RearAntiRollBar            INT NOT NULL,
   FrontSuspensionHeight      INT NOT NULL,
   RearSuspensionHeight       INT NOT NULL,
   BrakePressure              INT NOT NULL,
   BrakeBias                  INT NOT NULL,
   RearLeftTyrePressure       FLOAT NOT NULL,
   RearRightTyrePressure      FLOAT NOT NULL,
   FrontLeftTyrePressure      FLOAT NOT NULL,
   FrontRightTyrePressure     FLOAT NOT NULL,
   Ballast                    INT NOT NULL,
   FuelLoad                   FLOAT NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketCarSetupData_ID)
      REFERENCES PacketCarSetupData (ID)
);

CREATE OR REPLACE PROCEDURE insert_packet_car_setup_data("ID" uuid, "PacketHeader_ID" uuid)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO PacketCarSetupData 
		VALUES ("ID", "PacketHeader_ID");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_car_setup_data("ID" uuid, "PacketCarSetupData_ID" uuid, "FrontWing" integer, "RearWing" integer, "OnThrottle" integer, "OffThrottle" integer, "FrontCamber" double precision, "RearCamber" double precision, "FrontToe" double precision, "RearToe" double precision, "FrontSuspension" integer, "RearSuspension" integer, "FrontAntiRollBar" integer, "RearAntiRollBar" integer, "FrontSuspensionHeight" integer, "RearSuspensionHeight" integer, "BrakePressure" integer, "BrakeBias" integer, "RearLeftTyrePressure" double precision, "RearRightTyrePressure" double precision, "FrontLeftTyrePressure" double precision, "FrontRightTyrePressure" double precision, "Ballast" integer, "FuelLoad" double precision)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO CarSetupData 
		VALUES ("ID", "PacketCarSetupData_ID", "FrontWing", "RearWing", "OnThrottle", "OffThrottle", "FrontCamber", "RearCamber", "FrontToe", "RearToe", "FrontSuspension", "RearSuspension", "FrontAntiRollBar", "RearAntiRollBar", "FrontSuspensionHeight", "RearSuspensionHeight", "BrakePressure", "BrakeBias", "RearLeftTyrePressure", "RearRightTyrePressure", "FrontLeftTyrePressure", "FrontRightTyrePressure", "Ballast", "FuelLoad");
	END;
$BODY$;
