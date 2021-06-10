INSERT INTO Users (nickname, email, role, biography, profileImage, status, hashedpwd)
VALUES ("Alecs", "alexandre.bruno@ynov.com", "Admin", "Je suis une bio", "", "", "agigcusuyaffa98z4456"),
       ("Testeur", "testeur@ynov.com", "Admin", "Bio de test", "", "", "$2a$14$V5vya0uTIuQXpMChdPopke8F.Ki3hKpTNy0wKC5/8.SaHPF1xkdz2");

INSERT INTO Category (name, description)
VALUES ("Presentation", "Presente toi"),
       ("Guide", "Guidons nous"),
       ("Discussion", "Sujet de discussion ouvert"),
       ("Jeux", "Parlons bien, parlons jeux");

INSERT INTO UsersCat (id_users, id_category)
VALUES (1, 1);
