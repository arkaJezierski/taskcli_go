DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS tasks;

CREATE TABLE projects (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    data_to datetime,
    done BOOLEAN DEFAULT FALSE
);
