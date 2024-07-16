DROP database film_upload;

     DROP ROLE film_upload;

CREATE ROLE film_upload LOGIN PASSWORD 'salam_news';


CREATE DATABASE film_upload;

       \c film_upload;
          SET client_encoding to 'UTF-8'
              CREATE extension if NOT EXISTS "uuid-ossp";

CREATE TABLE category (
    UUID uuid PRIMARY KEY DEFAULT UUID_GENERATE_V4();
    name_tm VARCHAR(100),
    name_ru VARCHAR(100),
    name_en VARCHAR(100),
    last_update TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);





