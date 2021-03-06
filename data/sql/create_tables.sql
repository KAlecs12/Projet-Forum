
CREATE TABLE Users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	nickname TEXT NOT NULL,
	email TEXT NOT NULL,
    hashedpwd TEXT NOT NULL,
    squestion TEXT NOT NULL,
	role TEXT,
	biography TEXT,
	profileImage TEXT,
	status TEXT
);

CREATE TABLE Category (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT
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
	content TEXT NOT NULL,
	creationDate DATETIME,
	modificationDate DATETIME,
	deleteDate DATETIME,
	likes INTEGER, 
	dislikes INTEGER,
    nickname_users INTEGER,
    bio_users       TEXT,
	category TEXT,
	status TEXT,
    CONSTRAINT fk_category FOREIGN KEY (category) REFERENCES Category(name),
    CONSTRAINT fk_users_bio FOREIGN KEY (bio_users) REFERENCES Users(biography),
	CONSTRAINT fk_users_nickname FOREIGN KEY (nickname_users) REFERENCES Users(nickname)
);

CREATE TABLE PostsCat (
	id_posts INTEGER,
	id_category INTEGER,
	CONSTRAINT fk_posts_id FOREIGN KEY (id_posts) REFERENCES Posts(id),
	CONSTRAINT fk_category_id FOREIGN KEY (id_category) REFERENCES Category(id)
);

CREATE TABLE Comments (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	content TEXT,
	creationDate TIMESTAMP,
	modificationDate TIMESTAMP,
	deleteDate TIMESTAMP,
	likes INTEGER, 
	dislikes INTEGER,
	id_users INTEGER,
    bio_users  TEXT,
	id_posts INTEGER,
	CONSTRAINT fk_users_id FOREIGN KEY (id_users) REFERENCES Users(id),
    CONSTRAINT fk_users_bio FOREIGN KEY (bio_users) REFERENCES Users(biography),
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

CREATE TABLE LikesPosts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    id_post INTEGER,
    id_user INTEGER
);

CREATE TABLE DislikesPosts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    id_post INTEGER,
    id_user INTEGER
);

CREATE TABLE LikesComments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    id_post INTEGER,
    id_user INTEGER
);

CREATE TABLE DislikesComments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    id_post INTEGER,
    id_user INTEGER
);