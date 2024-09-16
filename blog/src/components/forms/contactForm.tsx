import {useState} from "react";
import { useForm, type SubmitHandler } from "react-hook-form";
import Notification from "../notification";
import Notify from "../../utils/notify";

type langData = {
  'name': {
    'placeholder': string
    'required': string
  } 
  'email': {
    'placeholder': string
    'required': string
    'invalid': string
  } 
  'message': {
    'placeholder': string
    'required': string
    'minLength': string
    'maxLength': string
  } 
  'send': string
  'response': {
    'success': string
    'error': string
  }
}

export default function ContactForm({content, endpoint}: {content: langData, endpoint: string}) {

  const [notificationView, SetNotificationView] = useState(false)
  const [notificationLabel, SetNotificationLabel] = useState("")
  const [notificationMessage, SetNotificationMessage] = useState("")

  type response = {
    status: string;
    statusCode: number
    data: any
    error: {
      code: string
      message: string
      details: string
    }
  }

  type Inputs = {
    name: string;
    email: string;
    message: string;
  }

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<Inputs>()

  const onSubmit: SubmitHandler<Inputs> = async (data) => {

    try {
      const resp = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
      })

      const respBody : response = await resp.json()
      
      if (respBody.status == "success") {
        SetNotificationLabel("Info") 
        SetNotificationMessage(content.response.success)
        reset()
      } else {
        SetNotificationLabel("Error") 
        SetNotificationMessage(content.response.error)
      }

      Notify(SetNotificationView)

    } catch {
      SetNotificationLabel("Error")
      SetNotificationMessage(content.response.error)
      Notify(SetNotificationView)
    }

  }

  return(
  <>
    <form onSubmit={handleSubmit(onSubmit)} className="w-full shadow-lg max-w-md text-sm bg-light text-dark font-content rounded-sm border border-dark p-5 flex flex-col gap-5">
      <input 
        type="text" 
        placeholder={content.name.placeholder}
        className="outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm" 
        {
        ...register("name", 
          {
            required: {
              value: true,
              message: content.name.required
            }
          })
        }
      /> 
      <p className={errors.name ? "text-xs" : "hidden"}>{errors?.name?.message}</p>
      
      <input 
        type="email" 
        placeholder={content.email.placeholder}
        className="outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm" 
        {
        ...register("email", 
          {
            required: {
              value: true,
              message: content.email.required
            },
            pattern: {
              value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
              message: content.email.invalid 
            }
          })
        }
      /> 
      <p className={errors.email ? "text-xs" : "hidden"}>{errors?.email?.message}</p>

      <textarea
        placeholder={content.message.placeholder}
        className="resize-none h-28 outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm" 
        {
        ...register("message", 
          {
            required: {
              value: true,
              message: content.message.required
            },
            minLength: {
              value: 5,
              message: content.message.minLength
            },
            maxLength: {
              value: 300,
              message: content.message.maxLength
            }
          })
        }
      ></textarea>
      <p className={errors.message ? "text-xs" : "hidden"}>{errors?.message?.message}</p>

      <button className="py-3 text-main bg-dark rounded-sm border border-dark hover:border-dashed hover:bg-light hover:text-dark">{content.send}</button>

    </form>
    <Notification view={notificationView} label={notificationLabel} message={notificationMessage} />
</>
  )
}
