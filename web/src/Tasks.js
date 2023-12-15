import React from 'react'
import './Tasks.css'
import state from './States'
import * as request from './Request'
import remove from './Images/remove.png'
import patch from './Images/patch.png'

const Tasks = () => {

    const deleteClick = (id) => {
        state.setIdTask(id)
        request.fetchDelete();

        state.setModalOkActive(true);
        request.fetchGet();
    }

    const patchClick = (id, text, tags, due) => {
        console.log(state.tasks);
        state.setIdTask(id);
        state.setTextTask(text);
        state.setTagsTask(tags);
        state.setDueTask(due);
        state.setModalActive(true);
    }

    return(
        <ul className="tasks">
            {React.Children.toArray(state.tasks.map(x => (
                <li>
                    <div className="task">
                        <h4>Task:</h4>
                        <p>{ x.text }</p> 

                        <h4>Tags:</h4>
                        <p><i>{ x.tags.join(', ') }</i></p>

                        <h4>Due:</h4>
                        <p>{ x.due.slice(0, 10) }</p>
                    </div>

                    <div className="buttons">
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

                    </div>
                        
                </li>
            )))}

        </ul> 
    );
};

export default Tasks;