export default function Notification({label, message, view}: {label: string, message: string, view: boolean}) {
  return(
    <div className={`w-full fixed z-10 top-0 left-0 p-2 ${view ? "" : "hidden"}`}>
      <div className="w-full flex justify-end">
        <div className="bg-light p-5 w-full max-w-sm rounded-sm border border-dark border-dashed">
          <p className="font-title text-lg">{label}</p>
          <p className="font-content text-sm">{message}</p>
        </div>
      </div>
    </div>
  )
}
