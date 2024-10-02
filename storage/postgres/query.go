package postgres

const (
	queryCreateTask = `
    INSERT INTO Tasks (description, tags, deadline)
    VALUES ($1, $2, $3)
	RETURNING task_id;
`
	queryReadTask = `
	SELECT * FROM Tasks
	WHERE task_id = $1;
`
	queryReadAllTask = `
	SELECT * FROM Tasks
	ORDER BY task_id ASC;
`
	queryUpdateTask = `
	UPDATE Tasks
	SET description = $1, tags = $2, deadline = $3, updated_at = NOW(), completed = $4
	WHERE task_id = $5;
`
	queryDeleteTask = `
	DELETE FROM Tasks
	WHERE task_id = $1;
`
	queryDeleteAllTasks = `
	DELETE FROM Tasks;
`
	queryReadTasksByTags = `
	SELECT * FROM Tasks
	WHERE $1 = ANY(tags)
	ORDER BY task_id ASC;
`
	queryReadTasksByDeadline = `	
	SELECT * FROM Tasks
	WHERE deadline = $1
	ORDER BY task_id ASC;
`
)
