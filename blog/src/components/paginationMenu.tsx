export default function PaginationMenu({found, page, pages, setPage, pagination}: {found: number, page: number, pages: number, setPage: Function, pagination: number[]}) {

  function Next() {

    let next = page + 1

    if (next > pages) {
      next = pages
    }

    setPage(next)
  }

  function Back() {
    let back = page - 1

    if (back < 1) {
      back = 1
    }

    setPage(back)
  }

  return(
    <div className="w-full flex justify-between mt-5 items-center text-sm">
      <div className="w-full">
        <p><span className="font-bold">{found}</span> items found</p>
      </div>
      <div className="w-full flex justify-end gap-4">
        <button onClick={Back} className={`${page <= 1 ? "hidden" : ""} font-bold text-main bg-dark rounded-sm p-2 hover:bg-light hover:text-dark border border-dark hover:border-dashed `}>
          <i className="fi fi-rs-angle-small-left flex items-center"></i>
        </button>
      {
      pagination.map(item => {
          return <button key={item} onClick={() => {setPage(item)}} className={item == page ? "line-through" : "hover:line-through"}>{item}</button>
        })
      }
        <button onClick={Next} className={`${page == pages ? "hidden" : ""} font-bold text-main bg-dark rounded-sm p-2 hover:bg-light hover:text-dark border border-dark hover:border-dashed `}>
          <i className="fi fi-rs-angle-small-right flex items-center"></i>
        </button>
      </div>
    </div>
  )
}
