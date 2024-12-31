-- Удаление таблиц, если они существуют
DROP TABLE IF EXISTS users_categories;
DROP TABLE IF EXISTS users_tasks;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS categories;
DROP TYPE IF EXISTS category_name;
-- Создание ENUM типа для категорий
CREATE TYPE category_name AS ENUM ('strength', 'intelligence', 'charisma', 'willpower');

-- Таблица пользователей
CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- Таблица категорий
CREATE TABLE categories
(
    id   SERIAL PRIMARY KEY,
    name category_name NOT NULL -- Используем наш ENUM тип
);

-- Таблица задач
CREATE TABLE tasks
(
    id            SERIAL PRIMARY KEY,
    title         VARCHAR(255) NOT NULL,
    description   TEXT,
    difficulty    INT CHECK (difficulty BETWEEN 1 AND 5),  -- Уровень сложности от 1 до 5
    category_id   INT,         -- Внешний ключ на таблицу категорий
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
);

-- Таблица для связи пользователей и задач
CREATE TABLE users_tasks (
                             id SERIAL PRIMARY KEY,
                             user_id INT NOT NULL,
                             task_id INT NOT NULL,
                             FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                             FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
                             CONSTRAINT unique_user_task UNIQUE (user_id, task_id)
);

-- Таблица для связи пользователей и категорий
CREATE TABLE users_categories (
      id          SERIAL PRIMARY KEY,
      user_id     INT NOT NULL,
      category_id INT NOT NULL,
      level       INT CHECK (difficulty BETWEEN 1 AND 5),  -- Уровень от 1 до 5
      FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
      FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE,
      CONSTRAINT unique_user_category UNIQUE (user_id, category_id)
);
