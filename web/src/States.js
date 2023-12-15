import { useState } from 'react'

// export const [modalActive, setModalActive] = useState(false);
// export const [modalOkActive, setModalOkActive] = useState(false);
// export const [message, setMessage] = useState({title: '', message: ''});
// export const [idTask, setIdTask] = useState(0);
// export const [textTask, setTextTask] = useState('');
// export const [tagsTask, setTagsTask] = useState('');
// export const [dueTask, setDueTask] = useState('');
// export const [tasks, setTasks] = useState([]);

const state = () => {
    const [modalActive, setModalActive] = useState(false);
    const [modalOkActive, setModalOkActive] = useState(false);
    const [message, setMessage] = useState({title: '', message: ''});
    const [idTask, setIdTask] = useState(0);
    const [textTask, setTextTask] = useState('');
    const [tagsTask, setTagsTask] = useState('');
    const [dueTask, setDueTask] = useState('');
    const [tasks, setTasks] = useState([]);
}

export default state

// export const state = () => {
//     const [modalActive, setModalActive] = useState(false);
//     const [modalOkActive, setModalOkActive] = useState(false);
//     const [message, setMessage] = useState({title: '', message: ''});
//     const [idTask, setIdTask] = useState(0);
//     const [textTask, setTextTask] = useState('');
//     const [tagsTask, setTagsTask] = useState('');
//     const [dueTask, setDueTask] = useState('');
//     const [tasks, setTasks] = useState([]);
// }