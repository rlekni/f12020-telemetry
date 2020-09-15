CREATE TABLE [IF NOT EXISTS] PacketSessionData (
   ID                         SERIAL PRIMARY KEY,
   PacketHeader_ID            SERIAL FOREIGN KEY,
   Weather                    INT NOT NULL,
   TrackTemperature           INT NOT NULL,
   AirTemperature             INT NOT NULL,
   TotalLaps                  INT NOT NULL,
   TrackLength                INT NOT NULL,
   SessionType                INT NOT NULL,
   TrackID                    INT NOT NULL,
   Formula                    INT NOT NULL,
   SessionTimeLeft            INT NOT NULL,
   SessionDuration            INT NOT NULL,
   PitSpeedLimit              INT NOT NULL,
   GamePaused                 INT NOT NULL,
   IsSpectating               INT NOT NULL,
   SpectatorCarIndex          INT NOT NULL,
   SliProNativeSupport        INT NOT NULL,
   NumMarshalZones            INT NOT NULL,
   SafetyCarStatus            INT NOT NULL,
   NetworkGame                INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] MarshalZone (
   ID                         SERIAL PRIMARY KEY,
   PacketSessionData_ID       SERIAL FOREIGN KEY,
   ZoneStart                  FLOAT NOT NULL,
   ZoneFlag                   FLOAT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);

CREATE TABLE [IF NOT EXISTS] WeatherForecastSample (
   ID                         SERIAL PRIMARY KEY,
   PacketSessionData_ID       SERIAL FOREIGN KEY,
   SessionType                INT NOT NULL,
   TimeOffset                 INT NOT NULL,
   Weather                    INT NOT NULL,
   TrackTemperature           INT NOT NULL,
   AirTemperature             INT NOT NULL,
   CreatedOn                  TIMESTAMPTZ
);