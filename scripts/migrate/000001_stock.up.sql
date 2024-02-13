CREATE TABLE IF NOT EXISTS stock_location (
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(6) NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,

    CONSTRAINT stock_location_pkey PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS stock_item (
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(6) NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,

    CONSTRAINT stock_item_pkey PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS stock_unit (
    id TEXT NOT NULL,

    location_id TEXT NOT NULL,

    CONSTRAINT stock_unit_pkey PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS stock_unit_item (
    id TEXT NOT NULL,

    item_id TEXT NOT NULL,

    CONSTRAINT stock_unit_item_pkey PRIMARY KEY ("id", "item_id")
);