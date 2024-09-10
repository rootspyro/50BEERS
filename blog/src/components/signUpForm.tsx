import {useEffect, useState} from "react";
import {useForm, type SubmitHandler} from "react-hook-form";
import Notification from "./notification";
import Notify from "../utils/notify"
import FetchUserData from "../utils/auth";


export default function SignUpForm() {

  const [passwordView, SetPasswordView] = useState(false)
  const [confPasswordView, SetConfPasswordView] = useState(false)
  const [notificationView, SetNotificationView] = useState(false)
  const [notificationMsg, SetNotificationMsg] = useState("")
  const [notificationLabel, SetNotificationLabel] = useState("")

  const handlePasswordView = () => {
    SetPasswordView(!passwordView)
  }

  const handleConfPasswordView = () => {
    SetConfPasswordView(!confPasswordView)
  }

  type Response = {
    status: string;
    statusCode: string;
    data: {
      username: string;
      email: string;
      origin: string
    };
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
    username: string;
    email: string;
    password: string;
    confirmPassword: string;
  }

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Inputs>()

  const password = watch("password") 

  const onSubmit: SubmitHandler<Inputs> = async(data) => {

    let endpoint = import.meta.env.PUBLIC_API_HOST + "/auth/blog/signup"

    try {
      
      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
      })

      const responseData : Response = await response.json()

      if (responseData.status == "success") {
        window.location.replace("/login")
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

  useEffect(() => {
    FetchUserData().then((user: {isLogged: boolean}) => {
      if (user.isLogged) {
        window.location.href = "/"
      } 
    })
  }, [])

  return (
  <>
    <form 
    onSubmit={handleSubmit(onSubmit)}
    className="w-full shadow-lg max-w-md text-sm bg-light text-dark font-content rounded-sm border border-dark p-5 flex flex-col gap-5"
    >
      <input 
        type="text"
        placeholder="username"
        className="outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm"
        {
          ...register("username",
            {
              required: {
                value: true,
                message: "username is required"
              }
            }
          )
        }
      />       
      <p className={ errors.username ? "text-xs" : "hidden"}><span className="font-bold text-red-500">*</span> {errors.username?.message}</p>
      <input 
        type="email"
        placeholder="email"
        className="outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm"
        {
          ...register("email",
            {
              required: {
                value: true,
                message: "email is required"
              },
              pattern: {
                value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
                message: "invalid email address"
              }
            }
          )
        }
      />       
      <p className={ errors.email ? "text-xs" : "hidden"}><span className="font-bold text-red-500">*</span> {errors.email?.message}</p>
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
                },
                pattern: {
                  value: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/,
                  message: "Password must be at least 8 characters long and include uppercase letters, lowercase letters, numbers, and special characters.",
                }
              }
            )
          }
        />       
        <button onClick={handlePasswordView} type="button" className="p-2 bg-dark text-main rounded-sm"><i className={`fi fi-${passwordView ? "ss-eye-crossed" : "rs-eye"} flex items-center`}></i></button>
      </div>
      <p className={ errors.password ? "text-xs" : "hidden"}><span className="font-bold text-red-500">*</span> {errors.password?.message}</p>
      <div className="w-full flex gap-2">
        <input 
          type={confPasswordView ? "text" : "password"}
        placeholder="confirm password"
          className="outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm w-full"
          {
            ...register("confirmPassword",
              {
                required: {
                  value: true,
                  message: "password confirmation is required"
                },
                validate: value => value === password || "password do not match"
              }
            )
          }
        />       
        <button onClick={handleConfPasswordView} type="button" className="p-2 bg-dark text-main rounded-sm"><i className={`fi fi-${confPasswordView ? "ss-eye-crossed" : "rs-eye"} flex items-center`}></i></button>
      </div>
      <p className={ errors.confirmPassword ? "text-xs" : "hidden"}><span className="font-bold text-red-500">*</span> {errors.confirmPassword?.message}</p>
      <button className="py-3 text-main bg-dark rounded-sm border border-dark hover:border-dashed hover:bg-light hover:text-dark">Sign Up</button>
    </form>
    <Notification view={notificationView} label={notificationLabel} message={notificationMsg} />
    </>
  )
}
