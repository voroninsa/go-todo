import React, { useEffect } from 'react';
import './App.css';
import Modal from './Modal';
import ModalOk from './ModalOk';
import Tasks from './Tasks';
import state from './States';
import * as request from './Request';

const App = () => {
  // const [modalActive, setModalActive] = useState(false)
  // const [modalOkActive, setModalOkActive] = useState(false)
  // const [message, setMessage] = useState({title: '', message: ''})
  // const [idTask, setIdTask] = useState(0)
  // const [textTask, setTextTask] = useState('')
  // const [tagsTask, setTagsTask] = useState('')
  // const [dueTask, setDueTask] = useState('')
  // const [tasks, setTasks] = useState([])

  useEffect(() => {
    request.fetchGet()
    // getFetch();
  }, [])

  // const getFetch = () => {
  //   const url = '/task/'

  //   fetch(url, {
  //     headers: {
  //       'Access-Control-Allow-Origin': '*'
  //     }
  //   }) 
  //     .then(res => res.json())
  //     .then(array => {
  //       array.sort((x, y) => x.id - y.id);
  //       if (array != null) {
  //       setTasks(array);
  //     } else {
  //       setTasks([]);
  //     }})
  // }

  return (
    <div className="app">
      <header className="app-header">
        <span>Task List</span> 
        <a className="createtask__button" onClick={() => state.setModalActive(true)}>Create Task</a>  
      </header>  

      <Tasks />

      <Modal />

      <ModalOk />

    </div>
  );
};

export default App;