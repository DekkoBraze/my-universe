import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom'
import NotFound from '../NotFound';
import "./Profile.css"

function Profile() {
    const [profileData, setProfileData] = useState([]);
    const [requestCompleted, setRequestCompleted] = useState(false)
    const { username } = useParams()
    useEffect(() => {
        fetch('/api/profile/' + username)
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
            <div className="mainProfile">
                <div className="profileInfo">
                    <h1>{profileData.username}</h1>
                    <h3>{profileData.age}</h3>
                    <h3>{profileData.country}</h3>
                    <h3>{profileData.status}</h3>
                    <h3>{profileData.description}</h3>
                </div>
            </div>
        );
    }
}

export default Profile;
