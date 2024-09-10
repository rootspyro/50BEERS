export default function NavbarComponent({currentPath, currentLang}: {currentPath: string, currentLang: string}) {

  console.log(currentLang)

  return(
    <>
    <div className="flex bg-dark justify-between px-5 py-3 items-center text-main">
      <div>
        <a href="/" className="font-title text-xl">50 BEERS</a>
      </div>
      <div className="flex gap-0 font-title items-center justify-end">
        <a href={currentPath} className={`hover:text-dark hover:bg-main p-2 ${currentLang == "en" ? "text-dark bg-main" : "text-main bg-dark"}`}>EN</a>
        <a href={`/es${currentPath}`} className={`hover:text-dark hover:bg-main p-2 ${currentLang == "es" ? "text-dark bg-main" : "text-main bg-dark"}`}>ES</a>
      </div>
    </div>
    </>
  )
}
