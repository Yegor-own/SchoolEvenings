DROP TABLE IF EXISTS courses;
CREATE TABLE IF NOT EXISTS courses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    age_range TEXT NOT NULL,
    preview_uuid TEXT NOT NULL UNIQUE,
    max_listeners INT NOT NULL,
    timetable TEXT NOT NULL,
    /* format -> integers separated by spaces like "0 2 6" - this is numbers of the week days */
    "from" DATETIME NOT NULL,
    "to" DATETIME NOT NULL,
    expires_at DATETIME NOT NULL
);