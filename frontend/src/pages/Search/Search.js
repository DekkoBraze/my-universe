import React, { useState } from 'react';
import "./Search.css"
import debounce from 'lodash.debounce'

function Search() {
    const [searchData, setSearchData] = useState([]);
    const [entityData, setEntityData] = useState([])
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
      
      function singleResultView() {
        
        setSearchData([])
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
                    onClick={singleResultView}
                    > 
                    {singleResult.name}
                    </div>
                  )
                })
              }
            </div>
          ) 
        }
        <div className='selectedEntity'>
            
        </div>
      </div>
    );
}

export default Search;
