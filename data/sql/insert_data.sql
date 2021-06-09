INSERT INTO Users (nickname, email, role, biography, profileImage, status, hashedpwd)
VALUES ("Alecs", "alexandre.bruno@ynov.com", "Admin", "Je suis une bio", "", "", "agigcusuyaffa98z4456");

INSERT INTO Category (name, description)
VALUES ("Video Game", "For video game lovers!");

INSERT INTO UsersCat (id_users, id_category)
VALUES (1, 1);
