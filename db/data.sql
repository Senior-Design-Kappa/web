CREATE TABLE rooms(id INTEGER PRIMARY KEY ASC, owner_id INTEGER, video_link TEXT);
CREATE TABLE room_invited_users(id INTEGER PRIMARY KEY ASC, room_id INTEGER, user_id INTEGER);
CREATE TABLE friends(id INTEGER PRIMARY KEY ASC, user_id_1 INTEGER, user_id_2 INTEGER);