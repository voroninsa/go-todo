import React, { useState } from 'react'
import './NewTask.css'


const NewTask = ({setActive, setError, setModalOkActive}) => {

    const [text, setText] = useState("")
    const [tags, setTags] = useState([])
    const [date, setDate] = useState("")

    const acceptClick = () => {
        const postData = {
            Text: text,
            Tags: tags.split(', '),
            Due: date + 'T00:00:00Z'
        }
        const url = 'http://localhost:8080/task/'

        fetch(url, {
            // mode: 'no-cors',
            method: 'post',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(postData)
        })
        .then((res) => {
            if (!res.ok) {
                setError(true);
                console.log("!ok");
            } else {
                setError(false);
                console.log("ok");
            }
        });    

        setActive(false);    

        setModalOkActive(true);
    }

    const declineClick = () => {
        setText("");
        setTags([]);
        setDate("");
        setActive(false)
    }

    return (
        <div className="new_task">
            <div className="new_task__entering">
                <h2>Enter a task</h2>
                
                <p>Text</p>
                <textarea  value={text} onChange={e => setText(e.target.value)} required></textarea>

                <p>Tags</p>
                <textarea  value={tags} onChange={e => setTags(e.target.value)} required></textarea>

                <p>Date</p>
                <input type="date" value={date} onChange={e => setDate(e.target.value)} required></input>

            
                <div className="new_task__buttons">
                    <a onClick={() => declineClick()}>Decline</a>
                    <a onClick={() => acceptClick()}>Accept</a>
                </div>
            </div>   
            
        </div>
    );
};

export default NewTask;