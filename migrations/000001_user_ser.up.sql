CREATE TABLE IF NOT EXISTS customer (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    phone VARCHAR(20) NOT NULL,
    email VARCHAR(30) NOT NULL,
    language VARCHAR(2) CHECK (language IN ('uz', 'ru', 'en')) NOT NULL,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(6) CHECK (gender IN ('male', 'female')) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT customer_unique_phone_deleted_at UNIQUE (phone, deleted_at),
    CONSTRAINT customer_unique_email_deleted_at UNIQUE (email, deleted_at)
);

CREATE TABLE IF NOT EXISTS "system_user" (
    id UUID PRIMARY KEY,
    phone VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    role VARCHAR(7) CHECK (role IN ('admin', 'courier')) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT system_user_unique_phone_deleted_at UNIQUE (phone, deleted_at),
    CONSTRAINT system_user_unique_email_deleted_at UNIQUE (email, deleted_at)
);

CREATE TABLE IF NOT EXISTS shop (
    id UUID PRIMARY KEY,
    slug VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    name_uz VARCHAR(20) DEFAULT '',
    name_ru VARCHAR(20) DEFAULT '',
    name_en VARCHAR(20) DEFAULT '',
    description_uz VARCHAR(500) DEFAULT '',
    description_ru VARCHAR(500) DEFAULT '',
    description_en VARCHAR(500) DEFAULT '',
    location VARCHAR NOT NULL,
    currency VARCHAR(3) CHECK (currency IN ('USD', 'EUR', 'UZS')) NOT NULL,
    payment_types VARCHAR[] NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT shop_unique_phone_deleted_at UNIQUE (phone, deleted_at)
);

CREATE TABLE IF NOT EXISTS seller (
    id UUID PRIMARY KEY,
    phone VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    shop_id UUID REFERENCES shop(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT seller_unique_phone_deleted_at UNIQUE (phone, deleted_at),
    CONSTRAINT seller_unique_email_deleted_at UNIQUE (email, deleted_at)
);

CREATE TABLE IF NOT EXISTS branch (
    id UUID PRIMARY KEY,
    phone VARCHAR NOT NULL,
    name VARCHAR(20) DEFAULT '',
    location POLYGON NOT NULL,
    address VARCHAR NOT NULL,
    open_time TIME NOT NULL,
    close_time TIME NOT NULL,
    active BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT branch_unique_phone_deleted_at UNIQUE (phone, deleted_at)
);


