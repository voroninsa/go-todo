import React from 'react'
import './Tasks.css'
import remove from './Images/remove.png'
import patch from './Images/patch.png'

const Tasks = ({ state }) => {

    const deleteClick = (id) => {
        state.setIdTask(id);
        state.setMethod('delete');
    }

    const patchClick = (id, text, tags, due) => {
        state.setIdTask(id);
        state.setTextTask(text);
        state.setTagsTask(tags);
        state.setDueTask(due);
        state.setModalActive(true);
    }

    const completeClick = (id, text, tags, due) => {
        state.setIdTask(id);
        state.setTextTask(text);
        state.setTagsTask(tags);
        state.setDueTask(due);
        state.setCompletedTask(true);
        state.setMethod('patch');
    }

    const getColorByDate = (taskDate) => {
        const date = new Date(taskDate);

        if (date < Date.now()) {
            return 'task__missed';
        } else {
            return 'task__notmissed';
        }
    }
    
    return (
        <ul className="tasks">
            {React.Children.toArray(state.tasks
                .filter((task) => task.completed == state.showCompletedTasks)
                .map(x => (
                    <li className={getColorByDate(x.due.slice(0, 10))}>
                        <div className="task__content">
                            <h4>Task:</h4>
                            <p>{ x.text }</p> 

                            <h4>Tags:</h4>
                            <p><i>{ x.tags.join(', ') }</i></p>

                            <h4>Due:</h4>
                            <p>{ x.due.slice(0, 10) }</p>
                        </div>

                        {!state.showCompletedTasks &&
                            <div className="task__buttons">
                                <div>
                                    <a onClick={() => deleteClick(x.id)} title="Delete task">
                                        <img className="image" src={remove}></img>
                                    </a>
                                </div>
                                <div>                      
                                    <a onClick={() => patchClick(x.id, x.text, x.tags.join(', '), x.due.slice(0, 10))} title="Patch task">
                                        <img className="image" src={patch}></img>
                                    </a>
                                </div> 
                                <div>                      
                                    <a onClick={() => completeClick(x.id, x.text, x.tags.join(', '), x.due.slice(0, 10))} title="Complete task">
                                        <img className="image" src={patch}></img>
                                    </a>
                                </div> 
                                
                            </div>         
                        } 
                    </li>
                )))
            }
        </ul> 
    );
};

export default Tasks;