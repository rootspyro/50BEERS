import {useEffect, useRef, useState} from "react"
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet"
import {Icon} from "leaflet"

type location = {
  id: string 
  name: string
  coordinates: number[]
  comments: string
  createdAt: string
  updatedAt: string
}

export default function Map({locations, defaultLocation} : {locations:location[], defaultLocation: location}) {

  const [currentLocation, SetCurrentLocation] = useState<location>(defaultLocation)

  const mapRef = useRef()

  const customIcon = new Icon({
    iconUrl: "https://i.imgur.com/AV4o63w.png", 
    iconSize: [28, 28]
  })

  const handleLocation = (data: location) => {
    SetCurrentLocation(data)
    mapRef.current.flyTo(data.coordinates, 16)
  }

  useEffect(() => {

  }, [currentLocation])

  return(
  <>
  <div className="w-full">
    <div className="relative">
    <MapContainer ref={mapRef} className="z-10 md:h-screen h-96" center={currentLocation.coordinates} zoom={12} scrollWheelZoom={false}>
  <TileLayer
    url="https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png"
  attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
  />
  {
  locations.map(location => {
    return location.coordinates.length > 1
    ? <Marker key={location.id} position={location.coordinates} icon={customIcon} eventHandlers={{click: () => handleLocation(location)}}>
      <Popup className="md:hidden">
        <p className="font-title">{location.name}</p> 
      </Popup>
      </Marker>
    : ""
  })
  }
</MapContainer>

  <div className="md:flex hidden absolute top-3 z-20 right-3 w-full max-w-72 p-4 overflow-y-scroll">
    <div className="flex flex-wrap w-full gap-2 overflow-y relative justify-end" style={{maxHeight: 500}}>
  {
    locations.map(location => {
      return location.coordinates.length > 1 
      ? 
      <div 
      key={location.id}
      className={`p-3 shadow rounded border border-dark bg-light font-title text-sm w-auto cursor-pointer hover:border-solid ${currentLocation.id == location.id ? "border-solid border-2 shadow-lg" : "border-dashed"}`}
      onClick={() => handleLocation(location)}
      >
      {location.name}
      </div>
      : null
    })
  }
    </div>
  </div>

  <div className="md:flex flex-col gap-2 hidden absolute bottom-3 z-20 left-3 w-full max-w-96 p-5 bg-light border border-dashed border-dark rounded-sm">
    <p className="font-title text-xl">{currentLocation.name}</p> 
    <p className="text-xs">{currentLocation.comments}</p>
  </div>

  </div>
  </div>
</>
  )
}
