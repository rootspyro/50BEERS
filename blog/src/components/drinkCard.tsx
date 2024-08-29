import StarsView from "./starsView";

interface drink {
  id: string;
  name: string;
  type: string;
  abv: number;
  date: string;
  challengeNumber: number;
  stars: number;
  pictureUrl: string;
  createdAt: string;
  updatedAt: string;
}

export default function DrinkCard({drink}: {drink:drink}) {
  return(
    <div className="w-full bg-light border border-dark text-dark font-content cursor-pointer">
      <div className={`p-3 border-b border-dark flex justify-between items-center gap-2 text-lg ${drink.challengeNumber > 0 ? "bg-dark text-light" : ""}`}>
        <h3 className="font-title w-auto text-nowrap">{drink.name} : {drink.type}</h3> 
        <hr className={drink.challengeNumber > 0 ? " w-full border border-light border-dashed" : "hidden"}></hr>
        <p className={drink.challengeNumber > 0 ? "text-nowrap font-title" : "hidden"}>{drink.challengeNumber} / 50</p>
      </div>
      <div className="flex gap-2 flex-col md:flex-row md:items-center mt-1">
        <div className="w-auto">
          <img alt="beverage draw" src="https://i.imgur.com/0uN9aMd.png" className="max-w-36" />
        </div>
        <div className="w-full text-sm">
          <p><span className="font-bold">Date</span> {drink.date}</p>
          <p><span className="font-bold">ABV</span> {drink.abv} %</p>

          <div className="mt-5">
            <StarsView stars={drink.stars} size="xl" />
          </div>
        </div>
      </div>
    </div>
  )
}
