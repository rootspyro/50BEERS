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
      <div className={`p-3 border-b border-dark flex justify-between ${drink.challengeNumber > 0 ? "bg-dark text-light" : ""}`}>
        <h3 className="font-title">{drink.name} : {drink.type}</h3> 
      </div>
    </div>
  )
}
