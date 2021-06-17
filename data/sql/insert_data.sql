INSERT INTO Users (nickname, email, squestion, role, biography, profileImage, status, hashedpwd)
VALUES ("Alecs", "alexandre.bruno@ynov.com", "Queen", "Admin", "Je suis une bio", "", "Actif", "agigcusuyaffa98z4456"),
       ("Admin", "testeur@ynov.com", "admin", "Admin", "Administrateur du site a votre service !", "", "Actif", "$2a$14$V5vya0uTIuQXpMChdPopke8F.Ki3hKpTNy0wKC5/8.SaHPF1xkdz2");

INSERT INTO Category (name, description)
VALUES ("Règlement", "règles"),
       ("Présentations", "Présente toi"),
       ("Discussions", "Sujet de discussion ouvert"),
       ("Espace aide / guide", "Guidons nous"),
       ("Divertissement en tout genre / Jeux Video", "Parlons bien, parlons jeux");

INSERT INTO UsersCat (id_users, id_category)
VALUES (1, 1);

INSERT INTO Posts (title, content, creationDate, likes, dislikes, nickname_users, category, status)
VALUES ("Règles de vie", "Paphius quin etiam et Cornelius senatores, ambo venenorum artibus pravis se polluisse confessi, eodem pronuntiante Maximino sunt interfecti. pari sorte etiam procurator monetae extinctus est. Sericum enim et Asbolium supra dictos, quoniam cum hortaretur passim nominare, quos vellent, adiecta religione firmarat, nullum igni vel ferro se puniri iussurum, plumbi validis ictibus interemit. et post hoe flammis Campensem aruspicem dedit, in negotio eius nullo sacramento constrictus.", '2021-06-12 15:40:00', 0, 0, "Admin", "Règlement", "Actif"),
       ("Savoir faire !", "Paphius quin etiam et Cornelius senatores, ambo venenorum artibus pravis se polluisse confessi, eodem pronuntiante Maximino sunt interfecti. pari sorte etiam procurator monetae extinctus est. Sericum enim et Asbolium supra dictos, quoniam cum hortaretur passim nominare, quos vellent, adiecta religione firmarat, nullum igni vel ferro se puniri iussurum, plumbi validis ictibus interemit. et post hoe flammis Campensem aruspicem dedit, in negotio eius nullo sacramento constrictus.", '2021-06-12 16:32:00', 0, 0, "Admin", "Règlement", "Actif"),
       ("Administration du site", "Paphius quin etiam et Cornelius senatores, ambo venenorum artibus pravis se polluisse confessi, eodem pronuntiante Maximino sunt interfecti. pari sorte etiam procurator monetae extinctus est. Sericum enim et Asbolium supra dictos, quoniam cum hortaretur passim nominare, quos vellent, adiecta religione firmarat, nullum igni vel ferro se puniri iussurum, plumbi validis ictibus interemit. et post hoe flammis Campensem aruspicem dedit, in negotio eius nullo sacramento constrictus.", '2021-06-14 15:40:00', 0, 0, "Admin", "Présentations", "Actif"),
       ("Quelques idées ...", "Paphius quin etiam et Cornelius senatores, ambo venenorum artibus pravis se polluisse confessi, eodem pronuntiante Maximino sunt interfecti. pari sorte etiam procurator monetae extinctus est. Sericum enim et Asbolium supra dictos, quoniam cum hortaretur passim nominare, quos vellent, adiecta religione firmarat, nullum igni vel ferro se puniri iussurum, plumbi validis ictibus interemit. et post hoe flammis Campensem aruspicem dedit, in negotio eius nullo sacramento constrictus.", '2021-06-14 15:40:00', 0, 0, "Alecs", "Espace aide / guide", "Actif"),
       ("Quoi de neuf ?", "Paphius quin etiam et Cornelius senatores, ambo venenorum artibus pravis se polluisse confessi, eodem pronuntiante Maximino sunt interfecti. pari sorte etiam procurator monetae extinctus est. Sericum enim et Asbolium supra dictos, quoniam cum hortaretur passim nominare, quos vellent, adiecta religione firmarat, nullum igni vel ferro se puniri iussurum, plumbi validis ictibus interemit. et post hoe flammis Campensem aruspicem dedit, in negotio eius nullo sacramento constrictus.", '2021-06-14 15:40:00', 0, 0, "Admin", "Discussions", "Actif");
