CREATE TABLE IF NOT EXISTS covid19_data(
    ID              UUID PRIMARY KEY UNIQUE NOT NULL,
    Country         VARCHAR(255) NOT NULL,
    CountryCode     VARCHAR(255) NOT NULL,
    Province        VARCHAR(255) NOT NULL,
    City            VARCHAR(255) NOT NULL,
    CityCode        VARCHAR(255) NOT NULL,
    Lat             VARCHAR(255) NOT NULL,
    Lon             VARCHAR(255) NOT NULL,
    Confirmed       BIGINT NOT NULL,
    Deaths          BIGINT NOT NULL,
    Recovered       BIGINT NOT NULL,
    Active          BIGINT NOT NULL,
    Date            TIMESTAMPTZ
);