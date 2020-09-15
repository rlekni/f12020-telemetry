CREATE TABLE IF NOT EXISTS PacketCarTelemetryData (
   ID                           uuid DEFAULT uuid_generate_v4 (),
   PacketHeader_ID              uuid NOT NULL,
   ButtonStatus                 INT NOT NULL,
   MfdPanelIndex                INT NOT NULL,
   MfdPanelIndexSecondaryPlayer INT NOT NULL,
   SuggestedGear                INT NOT NULL,
   CreatedOn                    TIMESTAMPTZ,

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketHeader_ID)
      REFERENCES PacketHeader (ID)
);

CREATE TABLE IF NOT EXISTS CarTelemetryData (
   ID                         uuid DEFAULT uuid_generate_v4 (),
   PacketCarTelemetryData_ID  uuid NOT NULL,
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
   CreatedOn                  TIMESTAMPTZ,

   PRIMARY KEY (ID),
   FOREIGN KEY (PacketCarTelemetryData_ID)
      REFERENCES PacketCarTelemetryData (ID)
);