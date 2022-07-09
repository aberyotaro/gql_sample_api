CREATE TABLE shops (
  id serial  PRIMARY KEY,
  name TEXT,
  postcode char(7),
  prefecture_code integer REFERENCES prefectures(id),
  city TEXT,
  address TEXT,
  phone varchar,
  lat REAL,
  lng REAL,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp NULL
);
