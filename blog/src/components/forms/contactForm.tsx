import { useForm, type SubmitHandler } from "react-hook-form";

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
}

export default function ContactForm({content}: {content: langData}) {

  type Inputs = {
    name: string;
    email: string;
    content: string;
  }

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Inputs>()

  const onSubmit: SubmitHandler<Inputs> = (data) => {
    console.log(data)
  }

  return(
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
        ...register("content", 
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
      <p className={errors.content ? "text-xs" : "hidden"}>{errors?.content?.message}</p>

      <button className="py-3 text-main bg-dark rounded-sm border border-dark hover:border-dashed hover:bg-light hover:text-dark">{content.send}</button>

    </form>
  )
}
