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