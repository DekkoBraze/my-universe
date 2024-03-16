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
          } else if (itemType === 'book') {
            fetch(
              `https://openlibrary.org/search.json?q=${e.target.value}`
            )
            .then(response => response.json())
            .then(data => {
              console.log(data)
              setSearchData(data.docs)
            })
            .catch(error => console.error(error));
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
            ItemArtist: singleResult.artist,
            ItemName: singleResult.name,
            ItemImage: singleResult.image[3]['#text'] 
          })
        } else if (itemType === 'book') {
          setItemJsonData({
            ...itemJsonData,
            Username: loggedUser.username,
            ItemType: itemType,
            ItemId: singleResult.key,
            ItemArtist: singleResult.author_name,
            ItemName: singleResult.title,
            ItemImage: `https://covers.openlibrary.org/b/isbn/${singleResult.isbn[0]}-M.jpg`
          })
        }
      }

      function closeItem() {
        setIsItemOpen(false);
    };

    function searchFields(singleResult) {
      if (itemType === 'game' || itemType === 'series') {
        return [singleResult.id, singleResult.name]
      } else if (itemType === 'movie') {
        return [singleResult.id, singleResult.title]
      } else if (itemType === 'music') {
        return [singleResult.mbid, singleResult.artist + ' - ' + singleResult.name]
      } else if (itemType === 'book') {
        return [singleResult.key, singleResult.author_name + ' - ' + singleResult.title]
      }
    }

    if (loggedUser) {
      return (
        <div className='searchMain'>
        <h2 className="searchTitle">Search new item</h2>
        <div className="searchDiv">
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
                  if (itemType === 'movie' && singleResult.mbid !== '' || 
                  itemType === 'book' && singleResult.isbn !== undefined ||
                  itemType === 'game' || itemType === 'series' || itemType === 'music') {
                    const [keyNum, name] = searchFields(singleResult)
                    return (
                      <div 
                      className="singleResult"
                      key={keyNum}
                      onClick={() =>
                        singleResultView(singleResult)
                      }
                      > 
                      {name}
                      </div>
                    )
                  }
                })
              }
            </div>
          ) 
        }
        </div>
        <div>
          <ItemModal
            isItemOpen={isItemOpen}
            itemInfo={itemJsonData}
            onClose={closeItem}
            isProfilePage={false}
            isUsersPage={true}
            />
        </div>
      </div>
    );
    } else {
      return (
      <h1 className='searchTitle'>You need to login into account to use this page.</h1>
      )
    }
    
}

export default Search;
