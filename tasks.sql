INSERT INTO tasks (title, description, links, updated_at, completed) 
VALUES (
    'Improve task writing process', 
    'Need to develop a more efficient method for adding tasks to the database.', 
    '["https://tasks.adjective.workers.dev", "https://github.com/AdjectiveAllison/tasks", "https://AdjectiveAllison.com"]', 
    strftime('%s', 'now'), 
    0
),
(
    'Write better code', 
    'Who wrote this?!', 
    '["https://github.com/AdjectiveAllison"]', 
    1677175342, 
    0
);
