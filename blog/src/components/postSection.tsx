import {useEffect, useRef, useState} from "react"
import StarsView from "./cards/starsView"
import DrinkCard from "./cards/drinkCard";
import PaginationMenu from "./paginationMenu";

interface lang {
  lang: string
  card: {
    lang: string
    date: string 
    country: string
    published: string
  }
  filters: {
    defaultSearch: string;
    defaultSortBy: string
    defaultCountry: string;
    defaultLocation: string;
    defaultCategory: string;
    categoryTitle: string;
    sort: {
      abv: string;
      date: string;
      name: string;
      stars: string;
   }
  },
  popularDrinks: string
  itemsFound: string
  notFound: {
    title1: string;
    title2: string;
    text: string;
    cta: string;
  },
  aboutMe: {
    title: string;
    part1: string;
    part2: string;
    part3: string;
    part4: string;
    part5: string;
  }
}

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
  coordinates: string[];
  comments: string;
  createdAt: string;
  updatedAt: string;
}

interface drink {
  id: string;
  name: string;
  type: string;
  abv: number;
  country: string;
  date: string;
  challengeNumber: number;
  location: string;
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

function PostSection({countries, locations, tags, lang} : {countries: country[], locations: location[], tags: tag[], lang: lang}) {

  const limit = 5

  const nameRef = useRef(null)

  // search Data
  const [name, SetName] = useState("")
  const [direction, SetDirection] = useState("down")
  const [sortBy, SetSortBy] = useState("created_at")
  const [country, SetCountry] = useState("")
  const [location, SetLocation] = useState("")
  const [category, SetCategory] = useState("")
  const [page, SetPage] = useState(1)
  const [pagination, SetPagination] = useState<number[]>([])

  const [drinks, SetDrinks] = useState<drinksSearch>({itemsFound: 0, items: [], pagination: { pages: 1, page: 1, pageSize: 10 }})

  function handleName(e: any) {
    e.preventDefault()
    SetName(nameRef.current.value)
  }

  function handleDirection() {

    if (direction == "down") {
      SetDirection("up")
    } else {
      SetDirection("down")
    }

  }

  function handleSort(e: any) {
    SetSortBy(e.target.value)
  }

  function handleCountry(e: any) {
    SetCountry(e.target.value)
  }

  function handleLocation(e: any) {
    SetLocation(e.target.value)
  }

  function handleCategoryMobile(e: any) {
    SetCategory(e.target.value)
  }

  function fetchDrinks() {

    const directionFilter = direction == "up" ? "asc" : "desc"

    const endpoint = import.meta.env.PUBLIC_API_HOST + `/drinks/blog?page=${page}&limit=${limit}&direction=${directionFilter}&sortBy=${sortBy}&country=${country}&location=${location}&category=${category}&name=${name}`

      try {

        fetch(endpoint, {
          headers: {
            "Accept-Language": lang.lang
          }
        })
          .then(result => result.json())
          .then(drinks => {
              if (drinks.status == "success") {
                let drinksData: drinksSearch = drinks.data
                
                let pagination = []

                for (let i = 1; i <= drinksData.pagination.pages; i++) {
                  pagination.push(i)
                }

                SetPagination(pagination)
                SetDrinks(drinksData)
              }
          })

      }catch(err) {
        console.log(err)
      }
  }

  useEffect(() => {

    fetchDrinks()

  }, [direction, sortBy, country, location, category, page, name])

  return(
  <>
    <div className="w-full flex justify-between font-content gap-5">

      <div id="content" className="w-full">

        <div id="mobile-filters" className="md:hidden flex flex-col gap-2">
          <form onSubmit={handleName} className="flex gap-2 items-center w-full">
            <input ref={nameRef} placeholder={lang.filters.defaultSearch} className="bg-light w-full border border-dark border-dashed p-2 text-xs rounded-sm outline-none" />
            <button className="p-3 rounded-sm flex items-center text-xs md:text-sm justify-center text-white bg-dark"><i className="fi fi-rs-search flex items-center"></i></button>
          </form>
          <div className="flex gap-2 items-center w-full">
            <select onChange={handleCategoryMobile} className="bg-light border border-dark border-dashed p-2.5 text-xs rounded-sm outline-none w-full">
              <option value="">{lang.filters.categoryTitle}</option>
              {
              tags.map((tag: tag) => {
               return <option key={tag.id} value={tag.id}>{tag.name}</option>
              })
              }
            </select>
            <select onChange={handleSort} defaultValue={"created_at"} className="bg-light border border-dark border-dashed p-2.5 text-xs rounded-sm outline-none w-full">
              <option value="created_at">{lang.filters.defaultSortBy}</option>
              <option value="abv">{lang.filters.sort.abv}</option>
              <option value="date">{lang.filters.sort.abv}</option>
              <option value="name">{lang.filters.sort.name}</option>
              <option value="stars">{lang.filters.sort.stars}</option>
            </select>
          </div>
          <div className="flex gap-2 items-center w-full">
            <select onChange={handleCountry} className="bg-light border border-dark border-dashed p-2.5 text-xs rounded-sm outline-none w-full">
              <option value="">{lang.filters.defaultCountry}</option>
              {
              countries.map((country: country) => {

                  return <option key={country.id} value={country.id}>{country.name}</option>
              })
              }
            </select>
            <select onChange={handleLocation} className="bg-light border border-dark border-dashed p-2.5 text-xs rounded-sm outline-none w-full">
              <option value="">{lang.filters.defaultLocation}</option>
              {
              locations.map((location: location) => {

                  return <option key={location.id} value={location.id}>{location.name}</option>
              })
              }
            </select>
            <button onClick={handleDirection} className="rounded-sm text-xs text-white bg-dark p-3">
              <i className={`fi fi-rs-angle-${direction} flex items-center`}></i>
            </button>
          </div>
        </div>

        <div id="desktop-filters" className="md:flex hidden w-full justify-between">
          <form onSubmit={handleName} className="flex gap-2 items-center">
            <input ref={nameRef} placeholder={lang.filters.defaultSearch} className="bg-light border border-dark border-dashed p-2 text-xs md:text-sm rounded-sm outline-none" />
            <button className="p-3 rounded-sm flex items-center text-xs md:text-sm justify-center text-white bg-dark"><i className="fi fi-rs-search flex items-center"></i></button>
          </form>
          <div className="flex gap-2 items-center justify-end">
            <button onClick={handleDirection} className="rounded-sm text-xs md:text-sm text-white bg-dark p-3">
              <i className={`fi fi-rs-angle-${direction} flex items-center`}></i>
            </button>
            <select onChange={handleSort} className="bg-light border border-dark border-dashed p-2.5 text-xs md:text-sm rounded-sm outline-none min-w-32  max-w-36">
              <option value="">{lang.filters.defaultSortBy}</option>
              <option value="abv">{lang.filters.sort.abv}</option>
              <option value="date">{lang.filters.sort.date}</option>
              <option value="name">{lang.filters.sort.name}</option>
              <option value="stars">{lang.filters.sort.stars}</option>
            </select>
            <select  onChange={handleCountry} className="bg-light border border-dark border-dashed p-2.5 text-xs md:text-sm rounded-sm outline-none min-w-32  max-w-36">
              <option value="">{lang.filters.defaultCountry}</option>
              {
              countries.map((country: country) => {

                  return <option key={country.id} value={country.id}>{country.name}</option>
              })
              }
            </select>
            <select onChange={handleLocation} className="bg-light border border-dark border-dashed p-2.5 text-xs md:text-sm rounded-sm outline-none min-w-32 max-w-36">
              <option value="">{lang.filters.defaultLocation}</option>
              {
              locations.map((location: location) => {

                  return <option key={location.id} value={location.id}>{location.name}</option>
              })
              }
            </select>
          </div>
        </div>
        <ul className="w-full bg-light p-2 hidden md:flex lg:hidden justify-center gap-4 text-sm rounded-sm border border-dashed border-dark mt-4">
          <li onClick={() => SetCategory("")} className={`cursor-pointer hover:line-through ${category == "" ? "line-through" : ""}`}>{lang.filters.defaultCategory}</li>
        {
          tags.map((tag: tag) => {
            return <li onClick={() => SetCategory(tag.id)} key={tag.id} className={`cursor-pointer hover:line-through ${category == tag.id ? "line-through" : ""}`}>{tag.name}</li>
          })
        }
        </ul>

        <div className={`mt-10 sm:mt-5 ${drinks.itemsFound == 0 ? "hidden" : ""}`}>
        <PaginationMenu foundText={lang.itemsFound} setPage={SetPage} found={drinks.itemsFound} page={page} pages={drinks.pagination.pages} pagination={pagination} />
          <div id="posts" className="w-full flex flex-col gap-4 items-start justify-start mt-5">
          {
            drinks.items.map((drink: drink) => {

                return (
                  <DrinkCard key={drink.id} drink={drink} lang={lang.card} />
                )
            })
          }
          </div>
        <PaginationMenu foundText={lang.itemsFound} setPage={SetPage} found={drinks.itemsFound} page={page} pages={drinks.pagination.pages} pagination={pagination} />
        </div>
        <div className={`my-16 sm:mt-20 w-full items-center flex flex-col gap-4 ${drinks.itemsFound == 0 ? "" : "hidden"}`}>
          <img src="/notfound.png" className="max-w-52" />
          <p className="text-center font-title text-3xl">{lang.notFound.title1} <span className="bg-dark text-light px-1">{lang.notFound.title2}</span></p>
          <p className="mt-8 text-center text-sm sm:text-base">{lang.notFound.text}</p>
          <a href="#" className="block mt-3 text-sm bg-dark border border-dark text-main max-w-56 p-3 rounded-sm hover:bg-light hover:text-dark hover:border-dashed active:border-solid">
            {lang.notFound.cta}
          </a>
        </div>

      <div className="mt-10 lg:hidden flex gap-5 flex-col items-start">
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


      <div id="side-section" className="lg:flex hidden flex-col gap-4 items-center max-w-72">

        <div className="p-5 bg-light border border-dashed border-dark rounded-sm w-full">
          <h3 className="font-title text-xl">{lang.filters.categoryTitle}</h3>
          <ul className="mt-5">
            <li onClick={() => SetCategory("")} className={`cursor-pointer hover:line-through ${category == "" ? "line-through" : ""}`}>{lang.filters.defaultCategory}</li>
            {
              tags.map((tag: tag) =>{
                  return <li onClick={() => SetCategory(tag.id)} key={tag.id} className={`cursor-pointer hover:line-through ${category == tag.id ? "line-through" : ""}`}>{tag.name}</li>
              })
            }
          </ul>
        </div>

        <div className="p-5 bg-light border border-dashed border-dark rounded-sm w-full">
          <h3 className="font-title text-xl">{lang.popularDrinks}</h3>
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
          <h3 className="font-title text-xl">{lang.aboutMe.title}</h3>
          <p className="text-sm mt-5">
            {lang.aboutMe.part1} <span className="font-bold">{lang.aboutMe.part2}!</span>
            <br/>             
            <br/>             
              {lang.aboutMe.part3}
            <br/>             
            <br/>             
            <span className="font-bold">{lang.aboutMe.part4}</span> {lang.aboutMe.part5}
          </p>
        </div>

      </div> 
    </div>
  </>
  )
}

export default PostSection;
