import {useState} from "react";
import {useForm, type SubmitHandler} from "react-hook-form";
import Notification from "./notification";
import Notify from "../utils/notify";

type lang = {
  text: string
  placeholder: string
  cta: string
  invalid: string
  required: string
  successResponse: string
  alreadySubscribed: string
  errorResponse: string
}

export default function SubscriptionForm({lang, api}: {lang:lang, api: string}) {

  const [notificationView, SetNotificationView] = useState(false)
  const [notificationMessage, SetNotificationMessage] = useState("")
  const [notificationLabel, SetNotificationLabel] = useState("")

  type Inputs = {
    email: string
  }

  type response = {
    status: string
    statusCode: number 
    data: any
    error: {
      code: string
      message: string
    }
  }

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<Inputs>()

  const onSubmit : SubmitHandler<Inputs> = async(data) => {
    const endpoint = api + "/newsletter/subscriber"
    try {

      const result = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data)
      })  
 
      const response : response = await result.json()
      
        if (response.status == "success") {
          SetNotificationLabel("Info")
          SetNotificationMessage(lang.successResponse)
          reset()
        } else if (response.status == "error") {

          SetNotificationLabel('Error')
          if (response.statusCode == 409) {
            SetNotificationMessage(lang.alreadySubscribed) 
          } else {
            SetNotificationMessage(lang.errorResponse)
          }
          
        } else {

          SetNotificationLabel('Error')
          SetNotificationMessage(lang.errorResponse)
        }

    } catch {
      SetNotificationLabel("Error")
      SetNotificationMessage(lang.errorResponse)
    } finally {
      Notify(SetNotificationView)
    }

  } 

  return(
  <>
    <div className="w-full py-7 max-w-md border-y border-light border-dashed flex flex-col items-center">
      <p className="text-light font-bold text-center">{lang.text}</p>
      <form onSubmit={handleSubmit(onSubmit)} className="flex bg-light rounded-sm w-full p-2 flex-between mt-5 text-dark">
        <input 
          placeholder={lang.placeholder}
          className="outline-none w-full text-sm" 
          {...register("email", {
             required: {
              value: true,
              message: lang.required
            },
            pattern: {
              value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
              message: lang.invalid
            }            
          })}
        />
        <button type="submit" className="bg-dark text-main rounded-sm text-sm p-2 border border-dark hover:bg-light hover:text-dark hover:border-dashed">{lang.cta}</button>
      </form>
      <p className="mt-5">{errors.email ? errors.email.message : ""}</p>
    </div>
    <Notification view={notificationView} message={notificationMessage} label={notificationLabel} />
</>
  )
}
