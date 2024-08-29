import {useEffect, useState} from "react"
import StarsView from "./starsView"

interface props {
  apiHost: string
}

function PostSection(props: props) {


  // search Data
  const [direction, SetDirection] = useState("down")

  // async function fetchCountries() {
  //
  //   try {
  //     const endpoint = props.apiHost + "/country/blog"
  //
  //     const result = await fetch(endpoint)
  //     const data = await result.json()
  //
  //     if (data.status == "success") {
  //      
  //     }
  //
  //   } catch(err) {
  //     console.log(err)
  //   }
  // }

  function handleDirection() {

    if (direction == "down") {
      SetDirection("up")
    } else {
      SetDirection("down")
    }

  }

  useEffect(() => {

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
            <select className="bg-light border border-dark border-dashed p-2.5 text-sm rounded-sm outline-none min-w-32">
              <option value="">Sort by</option>
            </select>
            <select className="bg-light border border-dark border-dashed p-2.5 text-sm rounded-sm outline-none min-w-32">
              <option value="">Country</option>
            </select>
            <select className="bg-light border border-dark border-dashed p-2.5 text-sm rounded-sm outline-none min-w-32">
              <option value="">Location</option>
            </select>
          </div>
        </div>
      </div> 

      <div id="side-section" className="md:flex hidden flex-col gap-4 items-center max-w-72">

        <div className="p-5 bg-light border border-dashed border-dark rounded-sm w-full">
          <h3 className="font-title text-xl">Categories</h3>
          <ul className="mt-5">
            <li className="cursor-pointer hover:line-through">{`>`} All</li>
            <li className="cursor-pointer hover:line-through">{`>`} Beer</li>
            <li className="cursor-pointer hover:line-through">{`>`} Whisky</li>
            <li className="cursor-pointer hover:line-through">{`>`} Rum</li>
            <li className="cursor-pointer hover:line-through">{`>`} Wine</li>
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
