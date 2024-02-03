CREATE TABLE IF NOT EXISTS candles (
    id SERIAL,
    instrument_uid TEXT,
    open REAL,
    close REAL,
    high REAL,
    low REAL,
    volume INTEGER,
    time BIGINT,
    is_complete BOOLEAN,
    UNIQUE (instrument_uid, time),
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS updates (
    instrument_id TEXT UNIQUE,
    first_time BIGINT,
    last_time BIGINT
);