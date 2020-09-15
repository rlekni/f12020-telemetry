CREATE TABLE [IF NOT EXISTS] PacketCarSetupData (
   ID                         SERIAL PRIMARY KEY,
   PacketHeader_ID            SERIAL FOREIGN KEY,
   CreatedOn                  TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] CarSetupData (
   ID                         SERIAL PRIMARY KEY,
   PacketCarSetupData_ID      SERIAL FOREIGN KEY,
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
   CreatedOn                  TIMESTAMPTZ
);