import React, { useState } from 'react';
import { MapContainer, Marker, TileLayer, Popup, useMapEvents } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import 'leaflet-defaulticon-compatibility';
import 'leaflet-defaulticon-compatibility/dist/leaflet-defaulticon-compatibility.css';
import { LatLngTuple } from 'leaflet';
import L from 'leaflet'; 
import selectedLocationIcon from '../../public/alert_icon.png';
import liveLocationIcon from '../../public/live_location.png';
import axios from 'axios';
import { useEffect } from 'react';

interface Address {
  ISO3166_2_lvl4: string;
  city: string;
  country: string;
  country_code: string;
  county: string;
  house_number: string;
  municipality: string;
  postcode: string;
  region: string;
  road: string;
  state: string;
  state_district: string;
  suburb: string;
}

interface AddressComponentProps {
  address: Address;
  updateAddress: (newAddress: Address) => void;
}

const MyMap: React.FC<AddressComponentProps> = ({ address, updateAddress }) => {
  const initialPosition: LatLngTuple = [-23.551126208315008, -46.63606901303866];
  const [userPosition, setUserPosition] = useState<LatLngTuple | null>(null);
  const [selectedPosition, setSelectedPosition] = useState<LatLngTuple | null>(null);
  const [firstLoad, setFirstLoad] = useState(true);
  
  const selectedCustomIcon = new L.Icon({
    iconUrl: selectedLocationIcon.src,
    iconSize: [40, 40],
    iconAnchor: [16, 32], 
  });

  const liveLocationCustomIcon = new L.Icon({
    iconUrl: liveLocationIcon.src,
    iconSize: [32, 32],
    iconAnchor: [16, 32], 
  });

  const LocationProvider = () => {
    const map = useMapEvents({
      click: (e) => {
        map.locate()
        setSelectedPosition([e.latlng.lat, e.latlng.lng]);
        fetchData(e.latlng.lat, e.latlng.lng);
        console.log('Location found:', e.latlng);
      }, 
      locationfound(e) {
        setUserPosition([e.latlng.lat, e.latlng.lng])
        if (firstLoad) {
          map.flyTo(e.latlng, map.getZoom())
          setFirstLoad(false);
        }
      },
    });

    return null;
  };

  const fetchData = async (latitude: number, longitude: number) => {
    try {
      const response = await axios.get(`https://nominatim.openstreetmap.org/reverse?lat=${latitude}&lon=${longitude}&format=json`);
      const { address } = response.data;
      updateAddress(address);
    } catch (error) {
      console.error('Error fetching address:', error);
    }
  };

  return (
    <MapContainer style={{ height: '62vh',  width: '50vh' }} center={initialPosition} zoom={16} scrollWheelZoom={false}>
      <LocationProvider />
      <TileLayer
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      {userPosition && <Marker position={userPosition}  icon={liveLocationCustomIcon}>
        <Popup>
          Sua localização atual.
        </Popup>
      </Marker>}
      {selectedPosition && <Marker position={selectedPosition} icon={selectedCustomIcon}>
        <Popup>
        <div>
        <p>Rua: {address.road}</p>
        <p>Bairro: {address.suburb}</p>
        <p>Cidade: {address.city}</p>
        <p>Estado: {address.state}</p>   
        <p>CEP: {address.postcode}</p>
        <p>País: {address.country}</p>
      </div>
        </Popup>
      </Marker>}
    </MapContainer>
  );
}

export default MyMap;
