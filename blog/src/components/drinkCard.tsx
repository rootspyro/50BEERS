import StarsView from "./starsView";

interface drink {
  id: string;
  name: string;
  type: string;
  abv: number;
  country: string;
  date: string;
  challengeNumber: number;
  location: string;
  stars: number;
  pictureUrl: string;
  createdAt: string;
  updatedAt: string;
}

function formatTimestamp(timestamp: string) :string {
  const dateTime = new Date(timestamp.split(" ")[0]);
  const year = dateTime.getFullYear();
  const month = dateTime.toLocaleString('default', { month: 'long' });
  const day = dateTime.getDate();

  return `${year}, ${month} ${day + 1}`;
}

export default function DrinkCard({drink}: {drink:drink}) {

  return(
    <div className="w-full bg-light border border-dark text-dark font-content">
      <div className={`p-3 border-b border-dark flex justify-between items-center gap-2 sm:text-lg text-sm ${drink.challengeNumber > 0 ? "bg-dark text-light" : ""}`}>
        <h3 className="font-title sm:text-nowrap text-center w-full sm:w-auto sm:text-left">{drink.name} : {drink.type}</h3> 
        <hr className={drink.challengeNumber > 0 ? " w-full border border-light border-dashed hidden sm:flex" : "hidden"}></hr>
        <p className={drink.challengeNumber > 0 ? "text-nowrap font-title hidden sm:flex" : "hidden"}>{drink.challengeNumber} / 50</p>
      </div>
      <div className="flex gap-2 flex-col sm:flex-row items-center sm:mt-1 mt-5">
        <div className="w-auto">
          <img alt="beverage draw" src="https://i.imgur.com/0uN9aMd.png" className="sm:max-w-36 max-w-32" />
        </div>
        <div className="w-full text-sm sm:p-0 p-5">
          <p><span className="font-bold">Date</span> {drink.date}</p>
          <p><span className="font-bold">Country</span> {drink.country}</p>
          <p><span className="font-bold">ABV</span> {drink.abv} %</p>

          <p className="mt-5 flex items-center gap-1"><i className="fi fi-ss-marker flex items-center"></i> {drink.location}</p>

          <div className="mt-5">
            <StarsView stars={drink.stars} size="xl" />
          </div>
        </div>
      </div>
      <div className="p-4 sm:p-2 border-t border-dark flex justify-center sm:justify-end items-center">
        <div className="text-nowrap">
          <p className="text-xs"><span className="font-bold">Published on</span> {formatTimestamp(drink.updatedAt)}</p>
        </div> 
      </div>
    </div>
  )
}
