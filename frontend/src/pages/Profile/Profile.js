import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom'
import { Link } from "react-router-dom";
import NotFound from '../NotFound';
import "./Profile.css"
import avatar from "../../avatarTemp.jpg"
import ItemModal from '../../components/ItemModal/ItemModal'

function Profile() {
    const [isItemOpen, setIsItemOpen] = useState(false)
    const [itemModalInfo, setItemModalInfo] = useState('')

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

    function handleItem(item) {
        setIsItemOpen(true);
        setItemModalInfo(item);
    }

    function closeItem() {
        setIsItemOpen(false);
    };

    function toSearch() {
        window.location.replace('/search');
    }

    if(requestCompleted)
    {   
        if(!profileData.username) { return <NotFound /> }
        return (
            <div className="mainProfile">
                <div className="profileInfo">
                    <img src={avatar} width={250} height={250}/>
                    <h1>{profileData.username}</h1>
                    <h3>{profileData.age}</h3>
                    <h3>{profileData.gender}</h3>
                    <h3>{profileData.country}</h3>
                    <h3>{profileData.status}</h3>
                    <h3>{profileData.description}</h3>
                </div>
                <div className='allItems'>
                {
                    userItems && (
                    userItems.map((item) => {
                        var itemInfo = "Name: " + item.ItemName + "\nRating: " + item.Rating + 
                        "\nComment: " + item.Comment + "\nRated At: " + item.CreatedAt
                        return (
                            <div 
                            className="item"
                            key={item.ItemId}> 
                            <img 
                            src={item.ItemImage} 
                            width={125} 
                            height={125} 
                            title={itemInfo}
                            onClick={() => {handleItem(item)}}>
                            </img>
                            </div>
                        )
                    }))
                }
                </div>
                {  
                loggedUser && loggedUser.username === profileData.username && (
                    <button className='searchButton' onClick={toSearch}>Add items</button>
                )   
   
                }
                <div>
                    <ItemModal
                    isItemOpen={isItemOpen}
                    itemInfo={itemModalInfo}
                    onClose={closeItem}
                    isProfilePage={true}
                    isUsersPage={loggedUser && username === loggedUser.username ? true : false}
                    />
                </div>
            </div>
        );
    }
}

export default Profile;
