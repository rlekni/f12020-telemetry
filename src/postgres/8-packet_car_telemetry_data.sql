CREATE TABLE [IF NOT EXISTS] PacketCarTelemetryData (
   ID                           SERIAL PRIMARY KEY,
   PacketHeader_ID              SERIAL FOREIGN KEY,
   ButtonStatus                 INT NOT NULL,
   MfdPanelIndex                INT NOT NULL,
   MfdPanelIndexSecondaryPlayer INT NOT NULL,
   SuggestedGear                INT NOT NULL,
   CreatedOn                    TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] CarTelemetryData (
   ID                         SERIAL PRIMARY KEY,
   PacketCarTelemetryData_ID  SERIAL FOREIGN KEY,
   Speed                      INT NOT NULL,
   Throttle                   FLOAT NOT NULL,
   Steer                      FLOAT NOT NULL,
   Brake                      FLOAT NOT NULL,
   Clutch                     INT NOT NULL,
   Gear                       INT NOT NULL,
   EngineRPM                  INT NOT NULL,
   Drs                        INT NOT NULL,
   RevLightsPercent           INT NOT NULL,
   BrakesTemperatureRL        INT NOT NULL,
   BrakesTemperatureRR        INT NOT NULL,
   BrakesTemperatureFL        INT NOT NULL,
   BrakesTemperatureFR        INT NOT NULL,
   TyresSurfaceTemperatureRL  INT NOT NULL,
   TyresSurfaceTemperatureRR  INT NOT NULL,
   TyresSurfaceTemperatureFL  INT NOT NULL,
   TyresSurfaceTemperatureFR  INT NOT NULL,
   TyresInnerTemperatureRL    INT NOT NULL,
   TyresInnerTemperatureRR    INT NOT NULL,
   TyresInnerTemperatureFL    INT NOT NULL,
   TyresInnerTemperatureFR    INT NOT NULL,
   EngineTemperature          INT NOT NULL,
   TyresPressureRL            FLOAT NOT NULL,
   TyresPressureRR            FLOAT NOT NULL,
   TyresPressureFL            FLOAT NOT NULL,
   TyresPressureFR            FLOAT NOT NULL,
   SurfaceTypeRL              INT NOT NULL,
   SurfaceTypeRR              INT NOT NULL,
   SurfaceTypeFL              INT NOT NULL,
   SurfaceTypeFR              INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);