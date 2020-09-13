CREATE TABLE [IF NOT EXISTS] PacketHeader (
   ID                      SERIAL PRIMARY KEY
   PacketFormat            INT NOT NULL,
   GameMajorVersion        INT NOT NULL,
   GameMinorVersion        INT NOT NULL,
   PacketVersion           INT NOT NULL,
   PacketID                INT NOT NULL,
   SessionUID              VARCHAR(250) NOT NULL,
   SessionTime             FLOAT8 NOT NULL,
   FrameIdentifier         INT NOT NULL,
   PlayerCarIndex          INT NOT NULL,
   SecondaryPlayerCarIndex INT NOT NULL,
   created_on              DATETIME
);

CREATE TABLE [IF NOT EXISTS] CarMotionData (
   ID                   SERIAL PRIMARY KEY
   WorldPositionX       FLOAT8 NOT NULL,
   WorldPositionY       FLOAT8 NOT NULL,
   WorldPositionZ       FLOAT8 NOT NULL,
   WorldVelocityX       FLOAT8 NOT NULL,
   WorldVelocityY       FLOAT8 NOT NULL,
   WorldVelocityZ       FLOAT8 NOT NULL,
   WorldForwardDirX     INT NOT NULL,
   WorldForwardDirY     INT NOT NULL,
   WorldForwardDirZ     INT NOT NULL,
   WorldRightDirX       INT NOT NULL,
   WorldRightDirY       INT NOT NULL,
   WorldRightDirZ       INT NOT NULL,
   GForceLateral        FLOAT8 NOT NULL,
   GForceLongitudinal   FLOAT8 NOT NULL,
   GForceVertical       FLOAT8 NOT NULL,
   Yaw                  FLOAT8 NOT NULL,
   Pitch                FLOAT8 NOT NULL,
   Roll                 FLOAT8 NOT NULL,
   CreatedOn            DATETIME
);

CREATE TABLE [IF NOT EXISTS] PacketMotionData (
   ID                         SERIAL PRIMARY KEY
   SuspensionPositionRL       FLOAT8 NOT NULL,
   SuspensionPositionRR       FLOAT8 NOT NULL,
   SuspensionPositionFL       FLOAT8 NOT NULL,
   SuspensionPositionFR       FLOAT8 NOT NULL,
   SuspensionVelocityRL       FLOAT8 NOT NULL,
   SuspensionVelocityRR       FLOAT8 NOT NULL,
   SuspensionVelocityFL       FLOAT8 NOT NULL,
   SuspensionVelocityFR       FLOAT8 NOT NULL,
   SuspensionAccelerationRL   FLOAT8 NOT NULL,
   SuspensionAccelerationRR   FLOAT8 NOT NULL,
   SuspensionAccelerationFL   FLOAT8 NOT NULL,
   SuspensionAccelerationFR   FLOAT8 NOT NULL,
   WheelSpeedRL               FLOAT8 NOT NULL,
   WheelSpeedRR               FLOAT8 NOT NULL,
   WheelSpeedFL               FLOAT8 NOT NULL,
   WheelSpeedFR               FLOAT8 NOT NULL,
   WheelSlipRL                FLOAT8 NOT NULL,
   WheelSlipRR                FLOAT8 NOT NULL,
   WheelSlipFL                FLOAT8 NOT NULL,
   WheelSlipFR                FLOAT8 NOT NULL,
   LocalVelocityX             FLOAT8 NOT NULL,
   LocalVelocityY             FLOAT8 NOT NULL,
   LocalVelocityZ             FLOAT8 NOT NULL,
   AngularVelocityX           FLOAT8 NOT NULL,
   AngularVelocityY           FLOAT8 NOT NULL,
   AngularVelocityZ           FLOAT8 NOT NULL,
   AngularAccelerationX       FLOAT8 NOT NULL,
   AngularAccelerationY       FLOAT8 NOT NULL,
   AngularAccelerationZ       FLOAT8 NOT NULL,
   FrontWheelsAngle           FLOAT8 NOT NULL,
   CreatedOn                  DATETIME
);