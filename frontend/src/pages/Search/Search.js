import React, { useState } from 'react';
import "./Search.css"
import debounce from 'lodash.debounce'
import ItemModal from '../../components/ItemModal/ItemModal'

function Search() {
    const [itemType, setItemType] = useState('game')
    const [isItemOpen, setIsItemOpen] = useState(false)
    const [searchData, setSearchData] = useState([]);
    const [itemJsonData, setItemJsonData] = useState({})
    const loggedUser = JSON.parse(localStorage.getItem("user"));

    const handleSearch = async (e) => {
        if (e.target.value.length >= 3) {
          if (itemType === 'game') {
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
          } else if (itemType === 'movie') {
            fetch('/api/getTmdbKey')
            .then(response => response.json())
            .then(data => {
              fetch(
                `https://api.themoviedb.org/3/search/movie?query=${e.target.value}&api_key=${data.message}`
              )
              .then(response => response.json())
              .then(data => {
                setSearchData(data.results)
              })
              .catch(error => console.error(error));
            })
          } else if (itemType === 'series') {
            fetch('/api/getTmdbKey')
            .then(response => response.json())
            .then(data => {
              fetch(
                `https://api.themoviedb.org/3/search/tv?query=${e.target.value}&api_key=${data.message}`
              )
              .then(response => response.json())
              .then(data => {
                setSearchData(data.results)
              })
              .catch(error => console.error(error));
            })
          } else if (itemType === 'music') {
            fetch('/api/getLastfmKey')
            .then(response => response.json())
            .then(data => {
              fetch(
                `https://ws.audioscrobbler.com/2.0/?method=album.search&album=${e.target.value}&api_key=${data.message}&format=json`
              )
              .then(response => response.json())
              .then(data => {
                setSearchData(data.results.albummatches.album)
              })
              .catch(error => console.error(error));
            })
          }
        }
      }
    
      const debouncedChange = debounce(handleSearch, 600);
      
      function singleResultView(singleResult) {
        setIsItemOpen(true)
        setSearchData([])
        if (itemType === 'game') {
          setItemJsonData({
            ...itemJsonData,
            Username: loggedUser.username,
            ItemType: itemType,
            ItemId: singleResult.id,
            ItemName: singleResult.name,
            ItemImage: singleResult.background_image 
          })
        } else if (itemType === 'movie') {
          setItemJsonData({
            ...itemJsonData,
            Username: loggedUser.username,
            ItemType: itemType,
            ItemId: singleResult.id,
            ItemName: singleResult.title,
            ItemImage: 'https://image.tmdb.org/t/p/w500' + singleResult.poster_path 
          })
        } else if (itemType === 'series') {
          setItemJsonData({
            ...itemJsonData,
            Username: loggedUser.username,
            ItemType: itemType,
            ItemId: singleResult.id,
            ItemName: singleResult.name,
            ItemImage: 'https://image.tmdb.org/t/p/w500' + singleResult.poster_path 
          })
        } else if (itemType === 'music') {
          setItemJsonData({
            ...itemJsonData,
            Username: loggedUser.username,
            ItemType: itemType,
            ItemId: singleResult.mbid,
            ItemName: singleResult.name,
            ItemImage: singleResult.image[3]['#text'] 
          })
        }
        
      }

      function closeItem() {
        setIsItemOpen(false);
    };

    return (
        <div className='searchMain'>
        <h2 className="searchTitle">Search new entity</h2>
        
        <div className='search-bar'>
        <select name="itemType" onChange={(e) => {setItemType(e.target.value); setSearchData([])}}> 
          <option value="game">Game</option> 
          <option value="movie">Movie</option> 
          <option value="series">Series</option> 
          <option value="music">Music album</option> 
          <option value="book">Literature</option> 
        </select>
            <input type="text" name="search" onChange={debouncedChange}></input>
        </div>
        {
          searchData.length > 0 && (
            <div className='searchResults'>
              {
                searchData.map((singleResult) => {
                  if (singleResult.mbid !== '') {
                    return (
                      <div 
                      className="singleResult"
                      key={itemType === 'music' ? singleResult.mbid : singleResult.id}
                      onClick={() =>
                        singleResultView(singleResult)
                      }
                      > 
                      {itemType === 'game' || itemType === 'series' || itemType === 'music' ? singleResult.name : singleResult.title}
                      </div>
                    )
                  }
                })
              }
            </div>
          ) 
        }
        <div>
          <ItemModal
            isItemOpen={isItemOpen}
            itemInfo={itemJsonData}
            onClose={closeItem}
            isProfilePage={false}
            />
        </div>
      </div>
    );
}

export default Search;
