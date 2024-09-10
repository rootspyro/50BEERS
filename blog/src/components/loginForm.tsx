import {useState} from "react";
import {useForm, type SubmitHandler} from "react-hook-form";
import Notification from "./notification";
import Notify from "../utils/notify"


export default function Login() {

  const [passwordView, SetPasswordView] = useState(false)
  const [notificationView, SetNotificationView] = useState(false)
  const [notificationMsg, SetNotificationMsg] = useState("")
  const [notificationLabel, SetNotificationLabel] = useState("")

  const handlePasswordView = () => {
    SetPasswordView(!passwordView)
  }

  type Response = {
    status: string;
    statusCode: string;
    data: string,
    error: {
      code: string;
      message: string;
      details: string;
      suggestion: string
      path: string;
      timestamp: string;
    }
  }

  type Inputs = {
    user: string;
    password: string;
  }

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Inputs>()

  const password = watch("password") 

  const onSubmit: SubmitHandler<Inputs> = async(data) => {

    let endpoint = import.meta.env.PUBLIC_API_HOST + "/auth/blog/login"

    try {
      
      const response = await fetch(endpoint, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
      })

      const responseData : Response = await response.json()

      if (responseData.status == "success") {
        window.location.href="/"
        return
      } else if (responseData.status == "error") {
        SetNotificationLabel("Error")
        SetNotificationMsg(responseData.error.details)
      } else {
        SetNotificationLabel("Error")
        SetNotificationMsg("Something went wrong")
      }
      
    } catch(err) {

      SetNotificationLabel("Error")
      SetNotificationMsg("Something went wrong")
    }

    Notify(SetNotificationView)
  }

  return (
  <>
    <form 
    onSubmit={handleSubmit(onSubmit)}
    className="w-full shadow-lg max-w-md text-sm bg-light text-dark font-content rounded-sm border border-dark p-5 flex flex-col gap-5"
    >
      <input 
        type="text"
        placeholder="username or email"
        className="outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm"
        {
          ...register("user",
            {
              required: {
                value: true,
                message: "user is required"
              }
            }
          )
        }
      />       
      <p className={ errors.user ? "text-xs" : "hidden"}><span className="font-bold text-red-500">*</span> {errors.user?.message}</p>
      <div className="w-full flex gap-2">
        <input 
          type={passwordView ? "text" : "password"}
          placeholder="password"
          className="outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm w-full"
          {
            ...register("password",
              {
                required: {
                  value: true,
                  message: "password is required"
                }
              }
            )
          }
        />       
        <button onClick={handlePasswordView} type="button" className="p-2 bg-dark text-main rounded-sm"><i className={`fi fi-${passwordView ? "ss-eye-crossed" : "rs-eye"} flex items-center`}></i></button>
      </div>
      <p className={ errors.password ? "text-xs" : "hidden"}><span className="font-bold text-red-500">*</span> {errors.password?.message}</p>
      <button className="py-3 text-main bg-dark rounded-sm border border-dark hover:border-dashed hover:bg-light hover:text-dark">Login</button>
    </form>
    <Notification view={notificationView} label={notificationLabel} message={notificationMsg} />
    </>
  )
}
