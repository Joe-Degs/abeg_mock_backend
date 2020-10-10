-- user_table contains immediate info about a registered user.
CREATE TABLE user_table (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    full_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(10) NOT NULL UNIQUE,
    email VARCHAR(150) UNIQUE NOT NULL,
    password VARCHAR(1024) UNIQUE NOT NULL
);

CREATE UNIQUE INDEX idx_phone_number ON user_table(phone_number);

-- user_data table holds arbitrary data about users as they update their
-- user info. the phone_number field is a foreign key pointing to a phone_number
-- in the user_table.
CREATE TABLE user_data (
    phone_number VARCHAR(10) NOT NULL,
    country VARCHAR(50),
    username VARCHAR(50),
    gender VARCHAR(17),
    date_of_birth DATE,
    image_file_path VARCHAR(15) UNIQUE,
    CONSTRAINT fk_phone_number FOREIGN KEY (phone_number)
        REFERENCES user_table (phone_number)
	ON DELETE CASCADE
	ON UPDATE CASCADE
);

-- next_of_kin table holds data of a users next of kin.
-- user_id is a foreignkey reference to user_table's id field.
CREATE TABLE next_of_kin (
    user_id INT NOT NULL,
    phone_number VARCHAR(10) UNIQUE PRIMARY KEY,
    full_name VARCHAR(50),
    gender VARCHAR(17)
    CONSTRAINT fk_user_id FOREIGN KEY (user_id)
        REFERENCES user_table (id)
	ON UPDATE CASCADE
	ON DELETE CASCADE
);

-- account table holds users monetary data specifically cash balance.
-- and a reference to user_table's phone_number field.
-- every successful registration results in the creation of an account
-- for that user.
CREATE TABLE account (
    phone_number VARCHAR(10) UNIQUE NOT NULL,
    balance REAL DEFAULT(0.0)
    CONSTRAINT fk_phone_number FOREIGN KEY (phone_number)
        REFERENCES user_table (phone_number)
	ON DELETE NO ACTION
	ON UPDATE CASCADE
);
