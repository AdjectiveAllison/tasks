DROP TABLE IF EXISTS tasks;
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    links TEXT NOT NULL DEFAULT '[]',
    updated_at INTEGER NOT NULL,
    completed INTEGER NOT NULL
);
CREATE INDEX idx_tasks_on_updated_at ON tasks (updated_at DESC);

/*
It could very well be better to give the links it's own table. If we did it would look similar to below:
I'm unsure what would perform better in d1, but I made a decision to move forward with the Json array for now.
TODO: Consider running performance tests for links being a json array or a table that's joined in the code. 
CREATE TABLE IF NOT EXISTS task_links (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id INTEGER NOT NULL,
    link TEXT NOT NULL,
    FOREIGN KEY(task_id) REFERENCES tasks(id) ON DELETE CASCADE
);
CREATE INDEX idx_task_links_on_task_id ON task_links (task_id);

*/