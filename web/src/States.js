import { useState, useEffect, useRef } from 'react'

const NewState = () => {
    const [method, setMethod] = useState('');
    const [modalActive, setModalActive] = useState();
    const [modalOkActive, setModalOkActive] = useState();
    const [showCompletedTasks, setShowCompletedTasks] = useState(false);
    const [message, setMessage] = useState({title: '', message: ''});
    const [idTask, setIdTask] = useState(0);
    const [textTask, setTextTask] = useState('');
    const [tagsTask, setTagsTask] = useState('');
    const [dueTask, setDueTask] = useState('');
    const [completedTask, setCompletedTask] = useState(false);
    const [tasks, setTasks] = useState([]);

    const checkFirstRender = useRef(true);

    useEffect(() => {
        if (checkFirstRender.current) {
            checkFirstRender.current = false;
            return;
        }
        setModalOkActive(true);
    }, [message])

    return {
        method, setMethod,
        modalActive, setModalActive,
        modalOkActive, setModalOkActive,
        showCompletedTasks, setShowCompletedTasks,
        message, setMessage,
        idTask, setIdTask,
        textTask, setTextTask,
        tagsTask, setTagsTask,
        dueTask, setDueTask,
        completedTask, setCompletedTask,
        tasks, setTasks,
        checkFirstRender
    }
}

export default NewState;