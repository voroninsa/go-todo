import React, { useState } from 'react'
import './NewTask.css'

const NewTask = ({setActive, setMessage, setModalOkActive, getFetch}) => {

    const [text, setText] = useState("")
    const [tags, setTags] = useState([])
    const [date, setDate] = useState("")

    const acceptClick = () => {
        if (!text || !tags.length || !date) {
            setMessage({title: 'Error!', message: 'Fill in all the fields'});
            setModalOkActive(true);
            return;
        }

        const postData = {
            Text: text,
            Tags: tags.split(', '),
            Due: date + 'T00:00:00Z'
        }
        const url = '/task/'

        fetch(url, {
            headers: {
                'accept': '*',
            },
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(postData)
        })
        .catch(setMessage({title: 'Error!', message: 'Server Internal Error'}))
        .then((res) => { console.log(res) 

            // res.text().then((r) => {console.log(r);})
            if (!res.ok) {
                setMessage({title: 'Error!', message: 'Server Internal Error (!ok)'});
            } else {
                setMessage({title: 'Success!', message: 'Task created successfully!'});
            }
        })

        setActive(false);    
        getFetch();
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
            <form className="new_task__entering">
                <h2>Enter a task</h2>
                
                <p>Text</p>
                <textarea  value={text} onChange={e => setText(e.target.value)}></textarea>

                <p>Tags</p>
                <textarea  value={tags} onChange={e => setTags(e.target.value)}></textarea>

                <p>Date</p>
                <input type="date" value={date} onChange={e => setDate(e.target.value)}></input>

            
                <div className="new_task__buttons">
                    <a onClick={() => declineClick()}>Decline</a>
                    <a onClick={() => acceptClick()}>Accept</a>
                </div>
            </form>          
        </div>
    );
};

export default NewTask;