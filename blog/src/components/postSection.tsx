import {useEffect, useState} from "react"
import StarsView from "./starsView"
import DrinkCard from "./drinkCard";

interface country {
  id: string;
  name: string;
  createdAt: string;
  updatedAt: string;
}

interface tag {
  id: string;
  name: string;
  createdAt: string;
  updatedAt: string;
}

interface location {
  id: string;
  name: string;
  url: string;
  comments: string;
  createdAt: string;
  updatedAt: string;
}

interface drink {
  id: string;
  name: string;
  type: string;
  abv: number;
  date: string;
  challengeNumber: number;
  stars: number;
  pictureUrl: string;
  createdAt: string;
  updatedAt: string;
}

interface drinksSearch {
  itemsFound: number;
  items: drink[];
  pagination: {pages: number; page: number, pageSize: number}
}

function PostSection({countries, locations, tags} : {countries: country[], locations: location[], tags: tag[]}) {

  // search Data
  const [direction, SetDirection] = useState("down")

  const [drinks, SetDrinks] = useState<drinksSearch>({itemsFound: 0, items: [], pagination: { pages: 1, page: 1, pageSize: 10 }})

  function handleDirection() {

    if (direction == "down") {
      SetDirection("up")
    } else {
      SetDirection("down")
    }

  }

  function fetchDrinks() {
    const endpoint = import.meta.env.PUBLIC_API_HOST + "/drinks/blog"

      try {

        fetch(endpoint)
          .then(result => result.json())
          .then(drinks => {
              if (drinks.status == "success") {
                let drinksData: drinksSearch = drinks.data
                SetDrinks(drinksData)
              }
          })

      }catch(err) {
        console.log(err)
      }
  }

  useEffect(() => {

    fetchDrinks()

  }, [direction])

  return(
  <>
    <div className="w-full flex justify-between font-content gap-5">

      <div id="content" className="w-full">
        <div id="desktop-filters" className="md:flex hidden w-full justify-between">
          <div className="flex gap-2 items-center">
            <input placeholder="search by name..." className="bg-light border border-dark border-dashed p-2 text-sm rounded-sm outline-none" />
            <button className="p-3 rounded-sm flex items-center justify-center text-white bg-dark"><i className="fi fi-rs-search flex items-center"></i></button>
          </div>
          <div className="flex gap-2 items-center justify-end">
            <button onClick={handleDirection} className="rounded-sm text-white bg-dark p-3">
              <i className={`fi fi-rs-angle-${direction} flex items-center`}></i>
            </button>
            <select className="bg-light border border-dark border-dashed p-2.5 text-sm rounded-sm outline-none min-w-32  max-w-36">
              <option value="">Sort by</option>
              <option value="abv">ABV</option>
              <option value="date">Date</option>
              <option value="name">Name</option>
              <option value="stars">Stars</option>
            </select>
            <select className="bg-light border border-dark border-dashed p-2.5 text-sm rounded-sm outline-none min-w-32  max-w-36">
              <option value="">Country</option>
              {
              countries.map((country: country) => {

                  return <option key={country.id} value={country.id}>{country.name}</option>
              })
              }
            </select>
            <select className="bg-light border border-dark border-dashed p-2.5 text-sm rounded-sm outline-none min-w-32 max-w-36">
              <option value="">Location</option>
              {
              locations.map((location: location) => {

                  return <option key={location.id} value={location.id}>{location.name}</option>
              })
              }
            </select>
          </div>
        </div>

        <div className="mt-5">
          <p className="text-sm"><span className="font-bold">{drinks.itemsFound}</span> items found...</p>
          <div id="posts" className="w-full flex flex-col gap-4 items-start justify-start mt-5">
          {
            drinks.items.map((drink: drink) => {

                return (
                  <DrinkCard drink={drink} />
                )
            })
          }
          </div>
        </div>
      </div> 


      <div id="side-section" className="lg:flex hidden flex-col gap-4 items-center max-w-72">

        <div className="p-5 bg-light border border-dashed border-dark rounded-sm w-full">
          <h3 className="font-title text-xl">Categories</h3>
          <ul className="mt-5">
            <li className="cursor-pointer hover:line-through">{`>`} All</li>
            {
              tags.map((tag: tag) =>{
                return <li key={tag.id} className={`cursor-pointer hover:line-through`}>{`>`} {tag.name}</li>
              })
            }
          </ul>
        </div>

        <div className="p-5 bg-light border border-dashed border-dark rounded-sm w-full">
          <h3 className="font-title text-xl">Popular Drinks</h3>
          <ul className="mt-5 flex flex-col gap-4">
            <li className="flex flex-col gap-1">
              <p>Barlobento Althaia</p>
              <StarsView stars={3} size="sm" /> 
            </li>

            <li className="flex flex-col gap-1">
              <p>Barlobento Althaia</p>
              <StarsView stars={3} size="sm" /> 
            </li>

            <li className="flex flex-col gap-1">
              <p>Barlobento Althaia</p>
              <StarsView stars={3} size="sm" /> 
            </li>
          </ul>
        </div>

        <div className="p-5 bg-light border border-dashed border-dark rounded-sm w-full">
          <h3 className="font-title text-xl">About Me</h3>
          <p className="text-sm mt-5">
            Hi, <span className="font-bold">My name is Spyro!</span>
            <br/>             
            <br/>             
            I am a software developer with a love for beverage tasting who on a trip to Spain was challenged to try 50 different beers and decided to document it.
            <br/>             
            <br/>             
            <span className="font-bold">I’m NOT</span> a professional taster, nor am I an oenologist, or anything like that, this blog is purely a hobby. If you need a developer you can count on me.
          </p>
        </div>

      </div> 
    </div>
  </>
  )
}

export default PostSection;
