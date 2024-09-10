import {useEffect, useState} from "react"
import FetchUserData from "../utils/auth"
import Notify from "../utils/notify"
import Notification from "./notification"

export default function Navbar() {

  const [notificationView, SetNotificationView] = useState(false)
  const [notificationMsg, SetNotificationMsg] = useState("")
  const [notificationLabel, SetNotificationLabel] = useState("")

  type User = {
    isLogged: boolean
    username: string
    email: string
  }

  const Logout = async() => {
    const endpoint = import.meta.env.PUBLIC_API_HOST + "/auth/blog/logout"

    try {

      const response = await fetch(endpoint, {
        method: "POST",
        credentials: "include"
      })

      const result = await response.json()

        if (result?.status == "success") {

          window.location.href="/"
          return

        } else if (result?.status == "error") {

          SetNotificationLabel("Error")
          SetNotificationMsg(result?.error?.details) 

        } else {
          SetNotificationLabel("Error")
          SetNotificationMsg("something went wrong")
        }

    } catch (err) {
      SetNotificationLabel("Error")
      SetNotificationMsg("something went wrong")
    }

    Notify(SetNotificationView)
  }

  const [user,SetUser] = useState<User>({isLogged: false, username: "", email: ""})

  useEffect(() => {
    FetchUserData().then((user: User) => SetUser(user))
  }, [])

  return(
    <>
    <div className="flex bg-dark justify-between px-5 py-3 items-center text-main">
      <div>
        <a href="/" className="font-title text-xl">50 BEERS</a>
      </div>
      <div className="flex gap-4 font-title items-center justify-end">
        <a href="#" className={`p-2 text-main text-sm hover:line-through duration-200 transition-all ${user.isLogged ? "": "hidden"}`}>{user.username}</a>
        <a onClick={Logout} href="#" className={`p-2 bg-main text-dark text-sm hover:bg-dark hover:text-main duration-200 transition-all ${user.isLogged ? "": "hidden"}`}>Logout</a>
        <a href="/login" className={`p-2 text-main text-sm hover:bg-main hover:text-dark duration-200 transition-all ${user.isLogged ? "hidden": ""}`}>Login</a>
        <a href="/signup" className={`p-2 bg-main text-dark text-sm hover:bg-dark hover:text-main duration-200 transition-all ${user.isLogged ? "hidden": ""}`}>Sign Up</a>
      </div>
    </div>
    <Notification view={notificationView} label={notificationLabel} message={notificationMsg} />
    </>
  )
}
