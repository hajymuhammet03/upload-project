DROP database film_upload;

     DROP ROLE film_upload;

CREATE ROLE film_upload LOGIN PASSWORD 'salam_news';


CREATE DATABASE film_upload;

       \c film_upload;
          SET client_encoding to 'UTF-8'
              CREATE extension if NOT EXISTS "uuid-ossp";


-------------Category-----------
CREATE TABLE category (
    UUID uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    name_tm VARCHAR(100),
    name_ru VARCHAR(100),
    name_en VARCHAR(100),
    last_update TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE category ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT NOW();

CREATE TABLE film_category (
    UUID uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    category_id uuid,
    last_update TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT category_id_fk
        FOREIGN KEY (category_id)
                    REFERENCES category (uuid)
                           ON UPDATE CASCADE ON DELETE CASCADE
);

-----------------Film------------------
CREATE TABLE film (
    UUID uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    release_year DATE NOT NULL,
    language_id uuid,
    rental_duration

);


------Language-----
CREATE TABLE language (
    UUID uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    name VARCHAR(100) NOT NULL ,
    last_update TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



