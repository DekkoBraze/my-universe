import React, { useState } from 'react';
import "./Search.css"
import debounce from 'lodash.debounce'

function Search() {
    const [searchData, setSearchData] = useState([]);
    const [singleResult, setSingleResult] = useState([])
    const [itemJsonData, setItemJsonData] = useState({
      username: '',
      itemId: '',
      itemImage: '',
      rating: '',
      comment: ''
    })
    const loggedUser = JSON.parse(localStorage.getItem("user"));

    const handleSearch = async (e) => {
        if (e.target.value.length >= 3) {
          fetch('/api/getRawgKey')
          .then(response => response.json())
          .then(data => {
          fetch(
            `https://api.rawg.io/api/games?key=${data.message}&search=${e.target.value}`
          )
          .then(response => response.json())
          .then(data => {
              setSearchData(data.results)
          })
          .catch(error => console.error(error));
        })
        .catch(error => console.error(error));
        }
      }
    
      const debouncedChange = debounce(handleSearch, 600);
      
      function singleResultView(singleResult) {
        setSingleResult(singleResult)
        setSearchData([])
        setItemJsonData({
          ...itemJsonData,
          username: loggedUser.username,
          itemId: singleResult.id,
          itemImage: singleResult.background_image
        })
      }

      function handleField(e) {
        const nextDataState = {
          ...itemJsonData,
          [e.target.name]: e.target.value,
        };
        setItemJsonData(nextDataState);
      }

      function handleRating(e) {
        var nextValue = Math.max(Number(0), Math.min(Number(10), Number(e.target.value)));
        if (nextValue === 0) {
          nextValue = ''
        }
        const nextDataState = {
          ...itemJsonData,
          [e.target.name]: nextValue,
        };
        setItemJsonData(nextDataState);
      }

      function handleCollecting() {
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
              setSingleResult([])
              setItemJsonData({
                username: '',
                itemId: '',
                itemImage: '',
                rating: '',
                comment: ''
              })
          } else {
            console.log(data.message)
          }})
          .catch((error) => {
            console.log(error)
          })
      }

    return (
        <div className='searchMain'>
        <h2 className="searchTitle">Search new entity</h2>
        <div className='search-bar'>
            <input type="text" name="search" onChange={debouncedChange}></input>
        </div>
        {
          searchData.length > 0 && (
            <div className='searchResults'>
              {
                searchData.map((singleResult) => {
                  return (
                    <div 
                    className="singleResult"
                    key={singleResult.id}
                    onClick={() =>
                      singleResultView(singleResult)
                    }
                    > 
                    {singleResult.name}
                    </div>
                  )
                })
              }
            </div>
          ) 
        }
        {
          singleResult.name && (
            <div className='selectedEntity'>
              <img src={singleResult.background_image} width="250" height="250"/>
              <div className="gameStats">
              <h1>{singleResult.name}</h1>
              <h2>Rating: {singleResult.rating}/5</h2>
              <h2>Ratings count: {singleResult.ratings_count}</h2>
              </div>
              <label>Rating</label>
              <input type="number" name="rating" style={{width: 45}} onChange={handleRating} value={itemJsonData.rating}></input>
              <label>Comment</label>
              <input type="text" name="comment" style={{width: 150}} maxLength={30} onChange={handleField}></input>
              <button type="submit" onClick={handleCollecting}>Add to the collection</button>
            </div>
          )
        }
      </div>
    );
}

export default Search;
