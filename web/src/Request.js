// import React, { useState } from 'react'
import state from './States';

const requestBody = {
    Text: state.textTask,
    Tags: state.tagsTask.split(', '),
    Due: state.dueTask + 'T00:00:00Z',
}

export const fetchGet = () => {
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

export const fetchPost = () => {
    fetch('/task/', {
        headers: {
            'accept': '*',
            'Content-Type': 'application/json',
        },
        method: 'POST',
        body: JSON.stringify(requestBody)
    })
    .catch(state.setMessage({title: 'Error!', message: 'Internal Server Error'}))
    .then((res) => {
        if (!res.ok) {
            state.setMessage({title: 'Error!', message: 'Internal Server Error (!ok)'});
        } else {
            state.setMessage({title: 'Success!', message: 'Task created successfully!'});
        }
    })
}

export const fetchPatch = () => {
    fetch(`/task/${state.idTask}`, {
        headers: {
            'accept': '*',
            'Content-Type': 'application/json',
        },
        method: 'PATCH',
        body: JSON.stringify(requestBody)
    })
    .catch(state.setMessage({title: 'Error!', message: 'Internal Server Error'}))
    .then((res) => { 
        if (!res.ok) {
            state.setMessage({title: 'Error!', message: 'Internal Server Error (!ok)'});
        } else {
            state.setMessage({title: 'Success!', message: 'Task patched successfully!'});
        }
    })
}

export const fetchDelete = () => {
    fetch(`/task/${state.idTask}`, {
        method: 'DELETE'
    })
    .catch(state.setMessage({title: 'Error!', message: 'Server Internal Error'}))
    .then((res) => {
        if (!res.ok) {
            state.setMessage({title: 'Error!', message: 'Error removing task'});
        } else {
            state.setMessage({title: 'Success!', message: 'Task was removed'});
        }
    })
} 
