CREATE TABLE IF NOT EXISTS PacketMotionData (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID            uuid NOT NULL,
   SuspensionPositionRL       FLOAT NOT NULL,
   SuspensionPositionRR       FLOAT NOT NULL,
   SuspensionPositionFL       FLOAT NOT NULL,
   SuspensionPositionFR       FLOAT NOT NULL,
   SuspensionVelocityRL       FLOAT NOT NULL,
   SuspensionVelocityRR       FLOAT NOT NULL,
   SuspensionVelocityFL       FLOAT NOT NULL,
   SuspensionVelocityFR       FLOAT NOT NULL,
   SuspensionAccelerationRL   FLOAT NOT NULL,
   SuspensionAccelerationRR   FLOAT NOT NULL,
   SuspensionAccelerationFL   FLOAT NOT NULL,
   SuspensionAccelerationFR   FLOAT NOT NULL,
   WheelSpeedRL               FLOAT NOT NULL,
   WheelSpeedRR               FLOAT NOT NULL,
   WheelSpeedFL               FLOAT NOT NULL,
   WheelSpeedFR               FLOAT NOT NULL,
   WheelSlipRL                FLOAT NOT NULL,
   WheelSlipRR                FLOAT NOT NULL,
   WheelSlipFL                FLOAT NOT NULL,
   WheelSlipFR                FLOAT NOT NULL,
   LocalVelocityX             FLOAT NOT NULL,
   LocalVelocityY             FLOAT NOT NULL,
   LocalVelocityZ             FLOAT NOT NULL,
   AngularVelocityX           FLOAT NOT NULL,
   AngularVelocityY           FLOAT NOT NULL,
   AngularVelocityZ           FLOAT NOT NULL,
   AngularAccelerationX       FLOAT NOT NULL,
   AngularAccelerationY       FLOAT NOT NULL,
   AngularAccelerationZ       FLOAT NOT NULL,
   FrontWheelsAngle           FLOAT NOT NULL,
   CreatedOn                  TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketHeader_ID)
      REFERENCES PacketHeader (ID)
);

CREATE TABLE IF NOT EXISTS CarMotionData (
   ID                   uuid DEFAULT uuid_generate_v4 (),
   PacketMotionData_ID  uuid NOT NULL,
   WorldPositionX       FLOAT NOT NULL,
   WorldPositionY       FLOAT NOT NULL,
   WorldPositionZ       FLOAT NOT NULL,
   WorldVelocityX       FLOAT NOT NULL,
   WorldVelocityY       FLOAT NOT NULL,
   WorldVelocityZ       FLOAT NOT NULL,
   WorldForwardDirX     INT NOT NULL,
   WorldForwardDirY     INT NOT NULL,
   WorldForwardDirZ     INT NOT NULL,
   WorldRightDirX       INT NOT NULL,
   WorldRightDirY       INT NOT NULL,
   WorldRightDirZ       INT NOT NULL,
   GForceLateral        FLOAT NOT NULL,
   GForceLongitudinal   FLOAT NOT NULL,
   GForceVertical       FLOAT NOT NULL,
   Yaw                  FLOAT NOT NULL,
   Pitch                FLOAT NOT NULL,
   Roll                 FLOAT NOT NULL,
   CreatedOn            TIMESTAMPTZ DEFAULT NOW(),

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketMotionData_ID)
      REFERENCES PacketMotionData (ID)
);

CREATE OR REPLACE PROCEDURE insert_packet_motion_data("ID" uuid, "PacketHeader_ID" uuid, "SuspensionPositionRL" double precision, "SuspensionPositionRR" double precision, "SuspensionPositionFL" double precision, "SuspensionPositionFR" double precision, "SuspensionVelocityRL" double precision, "SuspensionVelocityRR" double precision, "SuspensionVelocityFL" double precision, "SuspensionVelocityFR" double precision, "SuspensionAccelerationRL" double precision, "SuspensionAccelerationRR" double precision, "SuspensionAccelerationFL" double precision, "SuspensionAccelerationFR" double precision, "WheelSpeedRL" double precision, "WheelSpeedRR" double precision, "WheelSpeedFL" double precision, "WheelSpeedFR" double precision, "WheelSlipRL" double precision, "WheelSlipRR" double precision, "WheelSlipFL" double precision, "WheelSlipFR" double precision, "LocalVelocityX" double precision, "LocalVelocityY" double precision, "LocalVelocityZ" double precision, "AngularVelocityX" double precision, "AngularVelocityY" double precision, "AngularVelocityZ" double precision, "AngularAccelerationX" double precision, "AngularAccelerationY" double precision, "AngularAccelerationZ" double precision, "FrontWheelsAngle" double precision)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO PacketMotionData 
		VALUES ("ID", "PacketHeader_ID", "SuspensionPositionRL", "SuspensionPositionRR", "SuspensionPositionFL", "SuspensionPositionFR", "SuspensionVelocityRL", "SuspensionVelocityRR", "SuspensionVelocityFL", "SuspensionVelocityFR", "SuspensionAccelerationRL", "SuspensionAccelerationRR", "SuspensionAccelerationFL", "SuspensionAccelerationFR", "WheelSpeedRL", "WheelSpeedRR", "WheelSpeedFL", "WheelSpeedFR", "WheelSlipRL", "WheelSlipRR", "WheelSlipFL", "WheelSlipFR", "LocalVelocityX", "LocalVelocityY", "LocalVelocityZ", "AngularVelocityX", "AngularVelocityY", "AngularVelocityZ", "AngularAccelerationX", "AngularAccelerationY", "AngularAccelerationZ", "FrontWheelsAngle");
	END;
$BODY$;

CREATE OR REPLACE PROCEDURE insert_car_motion_data("ID" uuid, "PacketMotionData_ID" uuid, "WorldPositionX" double precision, "WorldPositionY" double precision, "WorldPositionZ" double precision, "WorldVelocityX" double precision, "WorldVelocityY" double precision, "WorldVelocityZ" double precision, "WorldForwardDirX" integer, "WorldForwardDirY" integer, "WorldForwardDirZ" integer, "WorldRightDirX" integer, "WorldRightDirY" integer, "WorldRightDirZ" integer, "GForceLateral" double precision, "GForceLongitudinal" double precision, "GForceVertical" double precision, "Yaw" double precision, "Pitch" double precision, "Roll" double precision)
LANGUAGE 'plpgsql'
AS $BODY$
	BEGIN
		INSERT INTO CarMotionData 
		VALUES ("ID", "PacketMotionData_ID", "WorldPositionX", "WorldPositionY", "WorldPositionZ", "WorldVelocityX", "WorldVelocityY", "WorldVelocityZ", "WorldForwardDirX", "WorldForwardDirY", "WorldForwardDirZ", "WorldRightDirX", "WorldRightDirY", "WorldRightDirZ", "GForceLateral", "GForceLongitudinal", "GForceVertical", "Yaw", "Pitch", "Roll");
	END;
$BODY$;

