import React, { useState } from 'react'
import './App.css'
import Modal from './Modal';
import NewTask from './NewTask';
import ModalOk from './ModalOk';

const App = () => {
  const [modalActive, setModalActive] = useState(true)
  const [modalOkActive, setModalOkActive] = useState(false)
  const [error, setError] = useState(false)

  return (
    <div className="app">
      <header className="app-header">
        <span>Task List</span> 
        <a className="createtask__button" onClick={() => setModalActive(true)}>Create Task</a>  
      </header>  

      <div className="tabs">

        <p>
          Task #1
        </p>


      </div>
      <Modal active={modalActive} setActive={setModalActive}>
        <NewTask setActive={setModalActive} setError={setError} setModalOkActive={setModalOkActive}/>
      </Modal>
      <ModalOk error={error} active={modalOkActive} setActive={setModalOkActive}/>
    </div>
  );
};

export default App;