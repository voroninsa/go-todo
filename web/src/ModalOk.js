import React from 'react'
import state from './States'
import './ModalOk.css'

const ModalOk = () => {
    return (
        <div className={state.modalOkActive ? "modalok active" : "modalok"}>
            <div className={state.modalOkActive ? "modalok__content active" : "modalok__content"}>
                <div className="new_modal">
                    <div className="new_modal__entering">
                        <h2>{state.message.title}</h2>
                        <p>{state.message.message}</p>
                        <div className="new_modal__buttons">
                            <a onClick={() => state.modalOkActive(false)}>Confirm</a>
                        </div>
                    </div>    
                </div>
            </div>
        </div>
    );
};

export default ModalOk;