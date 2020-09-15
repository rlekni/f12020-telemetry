CREATE TABLE [IF NOT EXISTS] PacketMotionData (
   ID                         SERIAL PRIMARY KEY,
   PacketHeader_ID            SERIAL FOREIGN KEY,
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
   CreatedOn                  TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] CarMotionData (
   ID                   SERIAL PRIMARY KEY,
   PacketMotionData_ID  SERIAL FOREIGN KEY,
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
   CreatedOn            TIMESTAMPTZ
);
