import React, { useEffect, useState } from 'react';
import "./ItemModal.css"

function ItemModal ({ isItemOpen, itemInfo, onClose }) {
    const [newRating, setNewRating] = useState('')
    const [newComment, setNewComment] = useState('')

    useEffect(() => {
      setNewRating(itemInfo.Rating)
      setNewComment(itemInfo.Comment)
    }, [itemInfo])

    function handleComment(e) {
        setNewComment(e.target.value)
    }

    function handleRating(e) {
      var nextValue = Math.max(Number(0), Math.min(Number(10), Number(e.target.value)));
      if (!Number.isInteger(nextValue)) {
        nextValue = Math.floor(nextValue)
      }
      if (nextValue === 0) {
        nextValue = ''
      }
      setNewRating(nextValue)
    }

    function handleCollecting() {
      var itemJsonData = {
        ...itemInfo,
        Rating: newRating,
        Comment: newComment
      }
      fetch('/api/createItem', {
        method: 'POST',
        body: JSON.stringify(itemJsonData),
        headers: {
          'Content-Type': 'application/json',
        }
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.message === 'OK') {
            onClose()
          } else {
            console.log(data.message)
        }})
        .catch((error) => {
          console.log(error)
        })
    }

    function handleDeletion() {
      fetch('/api/deleteUserItem/' + itemInfo.Username + '/' + itemInfo.ItemId, {
        method: 'DELETE'
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.message === 'OK') {
            onClose()
            window.location.reload();
          } else {
            console.log(data.message)
        }})
        .catch((error) => {
          console.log(error)
        }) 
      onClose()
    }

    if (isItemOpen !== true) {
        return null;
    }

    return (
        <div className="itemModal">
            <div className='itemWindow'>
              <button className="exit" onClick={onClose}>X</button>
              <div className='content'>
              <img className='itemImage' src={itemInfo.ItemImage} width="250" height="250"/>
              <h1 className='itemName'>{itemInfo.ItemName}</h1>
              <div className='inputs'>
                <div className='inputRating'>
                  <label>Rating</label>
                  <input type="number" name="rating" style={{width: 45}} onChange={handleRating} value={newRating || ''}></input>
                </div>
                <div className='inputComment'>
                  <label>Comment</label>
                  <textarea type="text" name="comment" rows="4" maxLength={83} onChange={handleComment} value={newComment}></textarea>
                </div>
              </div>
              <div className='buttons'>
              <button type="submit" onClick={handleCollecting}>Change</button>
              <button type="submit" onClick={handleDeletion}>Delete</button>
              </div>
              </div>
            </div>
        </div>
    )
}

export default ItemModal;