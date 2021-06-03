-- This file contains instructions to create all the tables

-- La table Users commence à la ligne      12
-- La table Category commence à la ligne   23
-- La table UsersCat commence à la ligne   30
-- La table Posts commence à la ligne      37
-- La table PostsCat commence à la ligne   50
-- La table Comments commence à la ligne   57
-- La table Badge commence à la ligne      70
-- La table UsersBadge commence à la ligne 78

CREATE TABLE Users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	nickname TEXT NOT NULL,
	email TEXT NOT NULL,
	role TEXT NOT NULL,
	biography TEXT NOT NULL,
	profileImage TEXT NOT NULL,
	status TEXT NOT NULL,
	hashedpw TEXT NOT NULL -- hashed password
);

CREATE TABLE Category (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL
);

-- Users X Category
CREATE TABLE UsersCat (
	id_users INTEGER,
	id_category INTEGER,
	CONSTRAINT fk_users_id FOREIGN KEY (id_users) REFERENCES Users(id),
	CONSTRAINT fk_category_id FOREIGN KEY (id_category) REFERENCES Category(id)
);

CREATE TABLE Posts (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	creationDate TIMESTAMP,
	modificationDate TIMESTAMP,
	deleteDate TIMESTAMP,
	likes INTEGER, 
	dislikes INTEGER,
	id_users INTEGER,
	CONSTRAINT fk_users_id FOREIGN KEY (id_users) REFERENCES Users(id)
);

-- Posts X Category
CREATE TABLE PostsCat (
	id_posts INTEGER,
	id_category INTEGER,
	CONSTRAINT fk_posts_id FOREIGN KEY (id_posts) REFERENCES Posts(id),
	CONSTRAINT fk_category_id FOREIGN KEY (id_category) REFERENCES Category(id)
);

CREATE TABLE Comments (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	creationDate TIMESTAMP,
	modificationDate TIMESTAMP,
	deleteDate TIMESTAMP,
	likes INTEGER, 
	dislikes INTEGER,
	id_users INTEGER,
	id_posts INTEGER,
	CONSTRAINT fk_users_id FOREIGN KEY (id_users) REFERENCES Users(id),
	CONSTRAINT fk_posts_id FOREIGN KEY (id_posts) REFERENCES Posts(id)
);

CREATE TABLE Badge (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	image TEXT NOT NULL,
	description TEXT NOT NULL
);

-- Users X Badge
CREATE TABLE UsersBadge (
	id_users INTEGER,
	id_badge INTEGER,
	CONSTRAINT fk_users_id FOREIGN KEY (id_users) REFERENCES Users(id),
	CONSTRAINT fk_badge_id FOREIGN KEY (id_badge) REFERENCES Badge(id)
);

