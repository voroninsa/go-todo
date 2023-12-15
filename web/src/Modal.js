import React from 'react'
import state from './States'
import * as request from './Request'
import './Modal.css'

const Modal = () => {

    const acceptClick = () => {

        if (!state.textTask || !state.tagsTask || !state.dueTask) {
            state.setMessage({title: 'Error!', message: 'Fill in all the fields'});
            state.setModalOkActive(true);
            return;
        } 

        if (state.idTask === 0) { 
            request.fetchPost();
        } else {
            request.fetchPatch();
        }

        request.fetchGet();
        state.setModalActive(false);  
        state.setIdTask(0);  
        state.setModalOkActive(true);
    }

    const declineClick = () => {
        state.setIdTask(0);
        state.setTextTask('');
        state.setTagsTask('');
        state.setDueTask('');
        state.setModalActive(false);
    }

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
                        <input type="date" value={state.dueTask} onChange={(e) => state.setDueTask(e.target.value)}></input>

                    
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