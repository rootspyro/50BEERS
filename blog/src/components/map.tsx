import {useEffect, useState} from "react"
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet"
export default function Map() {

  const [currentLocation, SetCurrentLocation] = useState("Locations")

  useEffect(() => {

  }, [currentLocation])

  return(
  <>
  <div className="w-full">
    <div className="relative">
    <MapContainer className="z-10" style={{height: 550}} center={[41.3766,2.1465]} zoom={11} scrollWheelZoom={false}>
  <TileLayer
    url="https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png"
  attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
  />
  <Marker  position={[41.422483,2.1768421]} eventHandlers={{click: () => SetCurrentLocation("2d2dspuma")}}>
  </Marker>
</MapContainer>
  <div className="text-3xl absolute top-3 right-3 z-20 text-right p-5 bg-light rounded border border-dark border-dashed">
    <h3 className="text-2xl font-title">{currentLocation}</h3>
    <p className="text-xs">Some of the locations where I've been tasting a drink</p>
  </div>
    </div>
  </div>
</>
  )
}
