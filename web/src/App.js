import React, { useState, useEffect } from 'react'
import './App.css'
import Modal from './Modal';
import NewTask from './NewTask';
import ModalOk from './ModalOk';
import Tasks from './Tasks';

const App = () => {
  const [modalActive, setModalActive] = useState(false)
  const [modalOkActive, setModalOkActive] = useState(false)
  const [message, setMessage] = useState({title: '', message: ''})
  const [tasks, setTasks] = useState([])

  useEffect(() => {
    getFetch();
    // const arr = [
    //   {
    //     id: 0,
    //     text: "Text",
    //     tags: ["tag1", "tag2", "tag2", "tag2"],
    //     due: "18.11.1973"
    //   },
    //   {
    //     id: 1,
    //     text: "Text Task",
    //     tags: ["tag1", "tag2"],
    //     due: "18.11.1973"
    //   },
    //   {
    //     id: 2,
    //     text: "Text",
    //     tags: ["tag1", "tag2", "tag2", "tag2"],
    //     due: "18.11.1973"
    //   },
    //   {
    //     id: 3,
    //     text: "Text Task",
    //     tags: ["tag1", "tag2"],
    //     due: "18.11.1973"
    //   },
    //   {
    //     id: 4,
    //     text: "Text Task",
    //     tags: ["tag1", "tag2"],
    //     due: "18.11.1973"
    //   },
    // ]
    // setTasks(arr)
  }, [])

  const getFetch = () => {
    const url = '/task/'

    fetch(url, {
      headers: {
        'Access-Control-Allow-Origin': '*'
        
      }
    }) 
      .then(res => res.json())
      .then(array => {if (array != null) {
        setTasks(array)
      } else {
        setTasks([])
      }})

    // let headersList = {
    //   "Access-Control-Allow-Origin": "*",
    //  }
     
    //  fetch("/task/", { 
    //    method: "GET",
    //    headers: headersList,
    //  }).then(function(response) {
    //    return response.json();
    //  }).then(function(data) {
    //    console.log(data);
    //  })
  }

  return (
    <div className="app">
      <header className="app-header">
        <span>Task List</span> 
        <a className="createtask__button" onClick={() => setModalActive(true)}>Create Task</a>  
      </header>  

      <Tasks tasks={tasks} getFetch={getFetch} setMessageModal={setMessage} setModalActive={setModalOkActive} />

      <Modal active={modalActive} setActive={setModalActive}>
        <NewTask setActive={setModalActive} setMessage={setMessage} setModalOkActive={setModalOkActive} getFetch={getFetch}/>
      </Modal>
      <ModalOk message={message} active={modalOkActive} setActive={setModalOkActive}/>
    </div>
  );
};

export default App;