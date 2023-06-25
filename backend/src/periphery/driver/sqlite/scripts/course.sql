DROP TABLE IF EXISTS courses;
CREATE TABLE IF NOT EXISTS courses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(254) NOT NULL,
    description TEXT NOT NULL,
    age_range TEXT NOT NULL,
    preview_uuid TEXT NOT NULL UNIQUE,
    max_listeners INT NOT NULL,
    timetable TEXT NOT NULL,
    /* format -> integers separated by spaces like "0 2 6" - this is numbers of the week days */
    "from" TIME NOT NULL,
    "to" TIME NOT NULL,
    expires_at TIME NOT NULL
);