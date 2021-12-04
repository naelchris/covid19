CREATE TABLE IF NOT EXISTS covid19_data(
    ID              bigserial PRIMARY KEY,
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

ALTER TABLE covid19_data ADD CONSTRAINT covid19_data_unique_date_country UNIQUE(date, country); 

CREATE TABLE IF NOT EXISTS user_data(
    ID              bigserial,
    Name            VARCHAR(255) NOT NULL,
    Email           VARCHAR(255) PRIMARY KEY NOT NULL,
    Password        VARCHAR(255) NOT NULL,
    DateOfBirth     TIMESTAMPTZ NOT NULL,
    VaccineType     VARCHAR(255),
    HealthStatus    VARCHAR (255),
    CreatedAt       TIMESTAMPTZ NOT NULL,
    UpdatedAt       TIMESTAMPTZ
);