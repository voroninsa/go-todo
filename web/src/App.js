import React, { useEffect } from 'react';
import './App.css';
import Modal from './Modal';
import ModalOk from './ModalOk';
import Tasks from './Tasks';
import NewState from './States';
import NewRequest from './Request';

const App = () => {
  const state = NewState();
  const request = NewRequest(state);

  useEffect(() => {
    request.fetchGet();
  }, [])

  return (
    <div className="app">
      <header className="app-header">
        <span>Task List</span> 
        <a className="createtask__button" onClick={() => state.setModalActive(true)}>Create Task</a>  
      </header>  

      <Tasks state={state} />

      <Modal state={state} />

      <ModalOk state={state} />

    </div>
  );
};

export default App;