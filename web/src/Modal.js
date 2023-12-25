import React from 'react'
import './Modal.css'

const Modal = ({ state }) => {
    if (state.checkFirstRender.current) {
        return null;
    }

    const acceptClick = () => {
        if (!state.textTask || !state.tagsTask || !state.dueTask) {
            state.setMessage({title: 'Error!', message: 'Fill in all the fields'});
            return;
        } 

        if (state.idTask === 0) { 
            state.setMethod('post')
        } else {
            state.setMethod('patch')
        }

        state.setModalActive(false);   
    }

    const declineClick = () => {
        state.setIdTask(0);
        state.setTextTask('');
        state.setTagsTask('');
        state.setDueTask('');
        state.setModalActive(false);
    }

    // if (state.modalActive === undefined) return null; else 
        return (
            <div className={state.modalActive ? "modal active" : "modal"} onClick={() => state.setModalActive(false)}>
                <div className={state.modalActive ? "modal__content active" : "modal__content"} onClick={e => e.stopPropagation()}>
                    <div className="new_task">
                        <form className="new_task__entering">
                            <h2>Enter a task</h2>
                            
                            <p>Text</p>
                            <textarea value={state.textTask} onChange={e => state.setTextTask(e.target.value)}></textarea>

                            <p>Tags</p>
                            <textarea value={state.tagsTask} onChange={e => state.setTagsTask(e.target.value)}></textarea>

                            <p>Date</p>
                            <input type="date" value={state.dueTask} onChange={e => state.setDueTask(e.target.value)}></input>

                        
                            <div className="new_task__buttons">
                                <a onClick={() => declineClick()}>Decline</a>
                                <a onClick={() => acceptClick()}>Accept</a>
                            </div>
                        </form>          
                    </div>
                </div>
            </div>
        );
};

export default Modal;