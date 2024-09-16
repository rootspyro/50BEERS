interface props {
  size: string;
  stars: number;
}

export default function StarsView(props: props) {

  let icons = []

  for (let i = 1; i <= 5; i++) {
    if (i <= props.stars) {
      icons.push("fi fi-ss-star")
    } else {
      icons.push("fi fi-rs-star")
    }
  }

  return(
    <div className="flex gap-1">
    {
    icons.map((iconClass: string, index: number) => {
        return <i key={`${iconClass}-${index}`} className={`${iconClass} text-${props.size}`}></i>
    })
    }
    </div>
  )
}
