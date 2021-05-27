-- This file contains demo data insertion

INSERT INTO Users (nickname, email, role, biography, profileImage, status)
VALUES ("Nami", "marilou.michel@ynov.com", "Admin", "I love unicorns", "https://urlz.fr/fI83", "");

INSERT INTO Category (name, description)
VALUES ("Video Game", "For video game lovers!");

INSERT INTO UsersCat (id_users, id_category)
VALUES (1, 1);