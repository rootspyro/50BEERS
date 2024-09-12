import { useForm, type SubmitHandler } from "react-hook-form";

type langData = {
  'name': string
  'email': string
  'message': string
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
        placeholder={content.name}
        className="outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm" 
        {
        ...register("name", 
          {
            required: {
              value: true,
              message: "name is required"
            }
          })
        }
      /> 

      <input 
        type="email" 
        placeholder={content.email}
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
              message: "invalid email format"
            }
          })
        }
      /> 

      <textarea
        placeholder={content.message}
        className="resize-none h-28 outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm" 
        {
        ...register("content", 
          {
            required: {
              value: true,
              message: "content is required"
            },
            minLength: {
              value: 5,
              message: "the message must be at least 5 characters long"
            },
            maxLength: {
              value: 300,
              message: "the message cannot exceed 300 characters"
            }
          })
        }
      ></textarea>

      <button className="py-3 text-main bg-dark rounded-sm border border-dark hover:border-dashed hover:bg-light hover:text-dark">{content.send}</button>

    </form>
  )
}
