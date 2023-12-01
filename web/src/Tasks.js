import React from 'react'
import './Tasks.css'
import remove from './Images/remove.png'

const Tasks = ({tasks, getFetch, setMessageModal, setModalActive}) => {

    const deleteClick = (id) => {
        const url = `/task/${id}`
        fetch(url, {
            method: 'DELETE'
        })
        .catch(setMessageModal({title: 'Error!', message: 'Server Internal Error'}))
        .then((res) => {
            if (!res.ok) {
                setMessageModal({title: 'Error!', message: 'Error removing task'});
            } else {
                setMessageModal({title: 'Success!', message: 'Task was removed'});
            }
        })

        setModalActive(true);
        getFetch();
    }

    return(
        <ul className="tasks">
            {tasks.map(x => (
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
                            <a onClick={() => deleteClick(x.id)}>
                                <img className="image" src={remove}></img>
                            </a>
                        </div>
                      
                    </div>
                        
                </li>
            ))}

        </ul> 
    );
};

export default Tasks;