CREATE TABLE IF NOT EXISTS PacketCarStatusData (
   ID                           uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID              uuid NOT NULL,
   CreatedOn                    TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketHeader_ID)
      REFERENCES PacketHeader (ID)
);

CREATE TABLE IF NOT EXISTS CarStatusData (
   ID                           uuid DEFAULT uuid_generate_v4 (),
   PacketCarStatusData_ID       uuid NOT NULL,
   TractionControl              INT NOT NULL,
   AntiLockBrakes               INT NOT NULL,
   FuelMix                      INT NOT NULL,
   FrontBrakeBias               INT NOT NULL,
   PitLimiterStatus             INT NOT NULL,
   FuelInTank                   FLOAT NOT NULL,
   FuelCapacity                 FLOAT NOT NULL,
   FuelRemainingLaps            FLOAT NOT NULL,
   MaxRPM                       INT NOT NULL,
   IdleRPM                      INT NOT NULL,
   MaxGears                     INT NOT NULL,
   DrsAllowed                   INT NOT NULL,
   DrsActivationDistance        INT NOT NULL,
   TyresWearRL                  INT NOT NULL,
   TyresWearRR                  INT NOT NULL,
   TyresWearFL                  INT NOT NULL,
   TyresWearFR                  INT NOT NULL,
   ActualTyreCompound           INT NOT NULL,
   VisualTyreCompound           INT NOT NULL,
   TyresAgeLaps                 INT NOT NULL,
   TyresDamageRL                INT NOT NULL,
   TyresDamageRR                INT NOT NULL,
   TyresDamageFL                INT NOT NULL,
   TyresDamageFR                INT NOT NULL,
   FrontLeftWingDamage          INT NOT NULL,
   FrontRightWingDamage         INT NOT NULL,
   RearWingDamage               INT NOT NULL,
   DrsFault                     INT NOT NULL,
   EngineDamage                 INT NOT NULL,
   GearBoxDamage                INT NOT NULL,
   VehicleFiaFlags              INT NOT NULL,
   ErsStoreEnergy               FLOAT NOT NULL,
   ErsDeployMode                INT NOT NULL,
   ErsHarvestedThisLapMGUK      FLOAT NOT NULL,
   ErsHarvestedThisLapMGUH      FLOAT NOT NULL,
   ErsDeployedThisLap           FLOAT NOT NULL,
   CreatedOn                    TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketCarStatusData_ID)
      REFERENCES PacketCarStatusData (ID)
);

CREATE OR REPLACE PROCEDURE insert_packet_car_status_data("ID" uuid, "PacketHeader_ID" uuid)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO PacketCarStatusData 
		VALUES ("ID", "PacketHeader_ID");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_car_status_data("ID" uuid, "PacketCarStatusData_ID" uuid, "TractionControl" integer, "AntiLockBrakes" integer, "FuelMix" integer, "FrontBrakeBias" integer, "PitLimiterStatus" integer, "FuelInTank" double precision, "FuelCapacity" double precision, "FuelRemainingLaps" double precision, "MaxRPM" integer, "IdleRPM" integer, "MaxGears" integer, "DrsAllowed" integer, "DrsActivationDistance" integer, "TyresWearRL" integer, "TyresWearRR" integer, "TyresWearFL" integer, "TyresWearFR" integer, "ActualTyreCompound" integer, "VisualTyreCompound" integer, "TyresAgeLaps" integer, "TyresDamageRL" integer, "TyresDamageRR" integer, "TyresDamageFL" integer, "TyresDamageFR" integer, "FrontLeftWingDamage" integer, "FrontRightWingDamage" integer, "RearWingDamage" integer, "DrsFault" integer, "EngineDamage" integer, "GearBoxDamage" integer, "VehicleFiaFlags" integer, "ErsStoreEnergy" double precision, "ErsDeployMode" integer, "ErsHarvestedThisLapMGUK" double precision, "ErsHarvestedThisLapMGUH" double precision, "ErsDeployedThisLap" double precision)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO CarStatusData 
		VALUES ("ID", "PacketCarStatusData_ID", "TractionControl", "AntiLockBrakes", "FuelMix", "FrontBrakeBias", "PitLimiterStatus", "FuelInTank", "FuelCapacity", "FuelRemainingLaps", "MaxRPM", "IdleRPM", "MaxGears", "DrsAllowed", "DrsActivationDistance", "TyresWearRL", "TyresWearRR", "TyresWearFL", "TyresWearFR", "ActualTyreCompound", "VisualTyreCompound", "TyresAgeLaps", "TyresDamageRL", "TyresDamageRR", "TyresDamageFL", "TyresDamageFR", "FrontLeftWingDamage", "FrontRightWingDamage", "RearWingDamage", "DrsFault", "EngineDamage", "GearBoxDamage", "VehicleFiaFlags", "ErsStoreEnergy", "ErsDeployMode", "ErsHarvestedThisLapMGUK", "ErsHarvestedThisLapMGUH", "ErsDeployedThisLap");
	END;
$BODY$;
