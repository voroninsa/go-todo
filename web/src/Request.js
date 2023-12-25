import { useEffect } from 'react'

const NewRequest = (state) => {

    useEffect(() => {
        switch (state.method) {
            case 'post':
                fetchPost();
                fetchGet();
                state.setMethod('');
                state.setIdTask(0); 
                break;

            case 'patch':
                fetchPatch();
                fetchGet();
                state.setMethod('');
                state.setIdTask(0); 
                state.setCompletedTask(false);
                break;

            case 'delete':
                fetchDelete();
                fetchGet();
                state.setMethod('');
                state.setIdTask(0);
                break;

            default:
                fetchGet();
                break;
        }
    }, [state.method])

    const requestBody = {
        Text: state.textTask,
        Tags: state.tagsTask.split(', '),
        Due: state.dueTask + 'T00:00:00Z',
        Completed: state.completedTask,
    }

    const fetchGet = () => {
        fetch('/task/', {
        headers: {
            'Access-Control-Allow-Origin': '*'
        }
        }) 
        .then(res => res.json())
        .then(array => {
            array.sort((x, y) => x.id - y.id);
            if (array != null) {
                state.setTasks(array);
            } else {
                state.setTasks([]);
            }
        })
    }

    const fetchPost = () => {
        fetch('/task/', {
            headers: {
                'accept': '*',
                'Content-Type': 'application/json',
            },
            method: 'POST',
            body: JSON.stringify(requestBody)
        })
        // .catch(state.setMessage({title: 'Error!', message: 'Internal Server Error'}))
        .then((res) => {
            if (!res.ok) {
                state.setMessage({title: 'Error!', message: 'Internal Server Error (!ok)'});
            } else {
                state.setMessage({title: 'Success!', message: 'Task created successfully!'});
            }
        })
    }

    const fetchPatch = () => {
        fetch(`/task/${state.idTask}`, {
            headers: {
                'accept': '*',
                'Content-Type': 'application/json',
            },
            method: 'PATCH',
            body: JSON.stringify(requestBody)
        })
        // .catch(state.setMessage({title: 'Error!', message: 'Internal Server Error'}))
        .then((res) => { 
            if (!res.ok) {
                state.setMessage({title: 'Error!', message: 'Internal Server Error (!ok)'});
            } 
            if (requestBody.Completed) {
                state.setMessage({title: 'Success!', message: 'Task completed successfully!'});
            } else {
                state.setMessage({title: 'Success!', message: 'Task patched successfully!'});
            }
        })
    }

    const fetchDelete = () => {
        fetch(`/task/${state.idTask}`, {
            method: 'DELETE'
        })
        // .catch(state.setMessage({title: 'Error!', message: 'Server Internal Error (Delete)'}))
        .then((res) => {
            if (!res.ok) {
                state.setMessage({title: 'Error!', message: 'Error removing task'});
            } else {
                state.setMessage({title: 'Success!', message: 'Task was removed'});
            }
        })
    } 

    return {
        fetchGet,
        fetchPost,
        fetchPatch,
        fetchDelete
    }

}

export default NewRequest;
