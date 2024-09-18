import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet"
export default function Map() {
  return(
    <MapContainer style={{height: 500}} center={[41.3766,2.1465]} zoom={11} scrollWheelZoom={false}>
  <TileLayer
    url="https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png"
  attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
  />
  <Marker position={[41.422483,-71.648708]}>
    <Popup>
      A pretty CSS3 popup. <br /> Easily customizable.
    </Popup>
  </Marker>
</MapContainer>
  )
}
