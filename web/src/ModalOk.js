import React from 'react'
import './ModalOk.css'

const ModalOk = ({message, active, setActive}) => {
    return (
        <div className={active ? "modalok active" : "modalok"}>
            <div className={active ? "modalok__content active" : "modalok__content"}>
                <div className="new_modal">
                    <div className="new_modal__entering">
                        <h2>{message.title}</h2>
                        <p>{message.message}</p>
                        <div className="new_modal__buttons">
                            <a onClick={() => setActive(false)}>Confirm</a>
                        </div>
                    </div>    
                </div>
            </div>
        </div>
    );
};

export default ModalOk;