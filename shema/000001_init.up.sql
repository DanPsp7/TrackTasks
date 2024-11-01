--создание таблиц
CREATE TABLE people (
                        id serial PRIMARY KEY,
                        name varchar(255) NOT NULL,
                        surname varchar(255) NOT NULL,
                        address varchar(255) NOT NULL,
                        passportNumber varchar(255) NOT NULL
);

CREATE TABLE tasks (
                       id serial PRIMARY KEY,
                       description varchar(255) NOT NULL,
                       status varchar(255) NOT NULL
);

CREATE TABLE peopleTasks (
                             personId int NOT NULL,
                             taskId int NOT NULL,
                             startTime timestamp,
                             stopTime timestamp,
                             duration int,
                             PRIMARY KEY (personId, taskId), -- Составной первичный ключ
                             FOREIGN KEY (personId) REFERENCES people(id), -- Внешний ключ на таблицу people
                             FOREIGN KEY (taskId) REFERENCES tasks(id) -- Внешний ключ на таблицу tasks
);