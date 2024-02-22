import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom'
import NotFound from '../NotFound';

function Profile() {
    const [profileData, setProfileData] = useState([]);
    const [requestCompleted, setRequestCompleted] = useState(false)
    const { username } = useParams()
    useEffect(() => {
        fetch('http://localhost:8000/profile/' + username)
        .then(response => response.json())
        .then(data => {
            setProfileData(data)
            setRequestCompleted(true)
        })
        .catch(error => console.error(error));
    }, [username]);
    if(requestCompleted)
    {
        if(!profileData.username) { return <NotFound /> }
        return (
            <div className="main">
                <h1>Welcome home, {profileData.username}!</h1>
            </div>
        );
    }
}

export default Profile;