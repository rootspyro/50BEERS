export default function ContactForm() {
  return(
    <form className="w-full shadow-lg max-w-md text-sm bg-light text-dark font-content rounded-sm border border-dark p-5 flex flex-col gap-5">
      <input 
        type="text" 
        placeholder="Your Name" 
        className="outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm" 
      /> 

      <input 
        type="email" 
        placeholder="Your Email" 
        className="outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm" 
      /> 

      <textarea
        placeholder="Email content."
        className="resize-none h-28 outline-none border border-dark border-dashed focus:border-solid p-2 rounded-sm" 
      ></textarea>

      <button className="py-3 text-main bg-dark rounded-sm border border-dark hover:border-dashed hover:bg-light hover:text-dark">Send</button>

    </form>
  )
}
