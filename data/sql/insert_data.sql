INSERT INTO Users (nickname, email, role, biography, profileImage, status, hashedpwd)
VALUES ("Alecs", "alexandre.bruno@ynov.com", "Admin", "Je suis une bio", "", "", "agigcusuyaffa98z4456"),
       ("Testeur", "testeur@ynov.com", "Admin", "Bio de test", "", "", "$2a$14$V5vya0uTIuQXpMChdPopke8F.Ki3hKpTNy0wKC5/8.SaHPF1xkdz2");

INSERT INTO Category (name, description)
VALUES ("Règlement", "règles"),
       ("Présentations", "Présente toi"),
       ("Discussions", "Sujet de discussion ouvert"),
       ("Espace aide / guide", "Guidons nous"),
       ("Divertissement en tout genre / Jeux Video", "Parlons bien, parlons jeux");

INSERT INTO UsersCat (id_users, id_category)
VALUES (1, 1);

INSERT INTO Posts (title, content, creationDate, likes, dislikes, nickname_users, category, status)
VALUES ("Règles de vie", "Bonjour, ici il y a quelques règles a respecter...", '2021-06-14 15:40:00', 0, 0, "Admin", "Règlement", "Actif");