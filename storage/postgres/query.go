package postgres

const (
	queryCreateTask = `
    INSERT INTO Tasks (description, tags, deadline)
    VALUES ($1, $2, $3)
	RETURNING task_id;
`
	queryReadTask = `
	SELECT * FROM Tasks 
	WHERE id = $1;
`
	queryUpdateTask = `
	UPDATE Tasks
	SET description = $1, tags = $2, deadline = $3, updated_at = NOW(), completed = $4
	WHERE task_id = $5;
`
	queryReadAllTask = `
	SELECT * FROM Tasks;
`
	queryDeleteTask = `
	DELETE FROM Tasks 
	WHERE id = $1;
`
	queryDeleteAllTasks = `
	DELETE FROM Tasks;
`
	queryReadTasksByTags = `
	SELECT * FROM Tasks 
	WHERE $1 = ANY(tags);
`
	queryReadTasksByDeadline = `	
	SELECT * FROM Tasks 
	WHERE date = $1;
`
)
