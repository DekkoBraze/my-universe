import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom'
import { Link } from "react-router-dom";
import NotFound from '../NotFound';
import "./Profile.css"
import avatar from "../../avatarTemp.jpg"

function Profile() {
    const [profileData, setProfileData] = useState([]);
    const [userItems, setUserItems] = useState([])
    const [requestCompleted, setRequestCompleted] = useState(false)
    const { username } = useParams()
    const loggedUser = JSON.parse(localStorage.getItem("user"));

    useEffect(() => {
        fetch('/api/profile/' + username)
        .then(response => response.json())
        .then(data => {
            setProfileData(data)
            fetch('/api/getUserItems/' + username)
            .then(response => response.json())
            .then(data => {
                setUserItems(data)
            })
            .catch(error => console.error(error))
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
                    <img src={avatar} width={250} height={250}/>
                    <h1>{profileData.username}</h1>
                    <h3>{profileData.age}</h3>
                    <h3>{profileData.country}</h3>
                    <h3>{profileData.status}</h3>
                    <h3>{profileData.description}</h3>
                </div>
                {
                    userItems.map((item) => {
                        return (
                            <div 
                            className="item"
                            key={item.ItemId}
                            > 
                            <img src={item.ItemImage} width={125} height={125}></img>
                            </div>
                        )
                    })
                }
                {
                    loggedUser && loggedUser.username === profileData.username && (
                        <Link className='searchLink' to="/search/">Add entities</Link>
                    )
                }
            </div>
        );
    }
}

export default Profile;
