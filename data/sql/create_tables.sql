
CREATE TABLE Users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	nickname TEXT NOT NULL,
	email TEXT NOT NULL,
    hashedpwd TEXT NOT NULL,
	role TEXT,
	biography TEXT,
	profileImage TEXT,
	status TEXT
);

CREATE TABLE Category (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL
);

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

CREATE TABLE UsersBadge (
	id_users INTEGER,
	id_badge INTEGER,
	CONSTRAINT fk_users_id FOREIGN KEY (id_users) REFERENCES Users(id),
	CONSTRAINT fk_badge_id FOREIGN KEY (id_badge) REFERENCES Badge(id)
);

CREATE TABLE SessionControl (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT,
    id_user INTEGER
);

