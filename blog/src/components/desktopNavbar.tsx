import {useEffect, useState} from "react"

export default function DesktopNavbar() {

  const [currentPath, SetPath] = useState("/")

    useEffect(() => {

    }, [currentPath])

  return(
  <>
  <ul className="flex gap-5 font-content text-dark font-bold text-xs w-full justify-center">
    <li ><a href="/" className={currentPath == "/" ? "line-through" : ""}>HOME</a></li>
    <li><a href="#contact" className={currentPath == "/#contact" ? "line-through" : ""}>CONTACT</a></li>
    <li><a href="/" >GITHUB</a></li>
  </ul>
  </>
  )
}
